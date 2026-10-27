package myokxapi

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

const fakeOrdersPush = `{"arg":{"channel":"orders","instType":"ANY"},"data":[{"clOrdId":"c1","ordId":"1","state":"live"}]}`

// fakeOkxServer speaks just enough of the OKX private ws protocol: login ack,
// subscribe ack, text ping/pong and order pushes.
type fakeOkxServer struct {
	srv *httptest.Server

	mu    sync.Mutex
	conns []*websocket.Conn

	logins atomic.Int32
	subs   atomic.Int32

	answerPing atomic.Bool
	ackSub     atomic.Bool
}

func newFakeOkxServer(t *testing.T) *fakeOkxServer {
	t.Helper()
	f := &fakeOkxServer{}
	f.answerPing.Store(true)
	f.ackSub.Store(true)

	up := websocket.Upgrader{}
	f.srv = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := up.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		f.mu.Lock()
		f.conns = append(f.conns, c)
		f.mu.Unlock()
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				return
			}
			text := string(msg)
			switch {
			case text == "ping":
				if f.answerPing.Load() {
					_ = c.WriteMessage(websocket.TextMessage, []byte("pong"))
				}
			case strings.Contains(text, `"op":"login"`):
				f.logins.Add(1)
				_ = c.WriteMessage(websocket.TextMessage, []byte(`{"event":"login","code":"0","msg":"","connId":"fake"}`))
			case strings.Contains(text, `"op":"subscribe"`):
				f.subs.Add(1)
				if f.ackSub.Load() {
					_ = c.WriteMessage(websocket.TextMessage,
						[]byte(`{"event":"subscribe","arg":{"channel":"orders","instType":"ANY"},"connId":"fake"}`))
				}
			}
		}
	}))
	t.Cleanup(f.srv.Close)

	// The client always dials wss://ws.okx.com, so redirect the dial.
	d := websocket.DefaultDialer
	oldDial, oldTLS, oldProxy := d.NetDialContext, d.TLSClientConfig, d.Proxy
	addr := f.srv.Listener.Addr().String()
	d.NetDialContext = func(ctx context.Context, network, _ string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, network, addr)
	}
	d.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	d.Proxy = nil
	t.Cleanup(func() { d.NetDialContext, d.TLSClientConfig, d.Proxy = oldDial, oldTLS, oldProxy })
	return f
}

func (f *fakeOkxServer) conn(i int) *websocket.Conn {
	f.mu.Lock()
	defer f.mu.Unlock()
	if i >= len(f.conns) {
		return nil
	}
	return f.conns[i]
}

func (f *fakeOkxServer) push(t *testing.T, i int) {
	t.Helper()
	c := f.conn(i)
	if c == nil {
		t.Fatalf("server conn %d not established", i)
	}
	if err := c.WriteMessage(websocket.TextMessage, []byte(fakeOrdersPush)); err != nil {
		t.Fatalf("push on conn %d: %v", i, err)
	}
}

func waitFor(t *testing.T, what string, d time.Duration, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if cond() {
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("timeout waiting for %s", what)
}

func setTestTimings(t *testing.T, keepalive, action time.Duration) {
	t.Helper()
	oldKA, oldAct := WebsocketTimeout, WsActionTimeout
	WebsocketTimeout, WsActionTimeout = keepalive, action
	t.Cleanup(func() { WebsocketTimeout, WsActionTimeout = oldKA, oldAct })
}

// connectOrders opens, logs in and subscribes against the fake server.
func connectOrders(t *testing.T) (*PrivateWsStreamClient, *Subscription[WsOrders]) {
	t.Helper()
	ws := (&MyOkx{}).NewPrivateWsStreamClient()
	if err := ws.OpenConn(); err != nil {
		t.Fatalf("OpenConn: %v", err)
	}
	if err := ws.Login((&MyOkx{}).NewRestClient("k", "s", "p")); err != nil {
		t.Fatalf("Login: %v", err)
	}
	sub, err := ws.SubscribeOrders("ANY", "", "")
	if err != nil {
		t.Fatalf("SubscribeOrders: %v", err)
	}
	return ws, sub
}

func expectOrderPush(t *testing.T, sub *Subscription[WsOrders]) {
	t.Helper()
	select {
	case o := <-sub.ResultChan():
		if o.ClOrdId != "c1" {
			t.Fatalf("unexpected order push: %+v", o)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("order push not delivered")
	}
}

func TestPrivateWs_InitialPushDelivered(t *testing.T) {
	setTestTimings(t, 10*time.Second, 2*time.Second)
	f := newFakeOkxServer(t)
	ws, sub := connectOrders(t)
	defer ws.Close()

	f.push(t, 0)
	expectOrderPush(t, sub)
}

func TestPrivateWs_ReconnectAfterServerDrop(t *testing.T) {
	setTestTimings(t, 10*time.Second, 2*time.Second)
	f := newFakeOkxServer(t)
	ws, sub := connectOrders(t)
	defer ws.Close()

	_ = f.conn(0).UnderlyingConn().Close()

	waitFor(t, "relogin and resubscribe", 10*time.Second, func() bool {
		return f.logins.Load() >= 2 && f.subs.Load() >= 2
	})
	f.push(t, 1)
	expectOrderPush(t, sub)
}

func TestPrivateWs_ForceReconnectRestoresSubscription(t *testing.T) {
	setTestTimings(t, 10*time.Second, 2*time.Second)
	f := newFakeOkxServer(t)
	ws, sub := connectOrders(t)
	defer ws.Close()

	if err := ws.ForceReconnect("test"); err != nil {
		t.Fatalf("ForceReconnect: %v", err)
	}
	waitFor(t, "relogin and resubscribe", 10*time.Second, func() bool {
		return f.logins.Load() >= 2 && f.subs.Load() >= 2
	})
	f.push(t, 1)
	expectOrderPush(t, sub)
}

func TestPrivateWs_SilentConnectionIsReconnected(t *testing.T) {
	setTestTimings(t, 200*time.Millisecond, 2*time.Second)
	f := newFakeOkxServer(t)
	ws, _ := connectOrders(t)
	defer ws.Close()

	f.answerPing.Store(false)
	waitFor(t, "reconnect of a silent connection", 10*time.Second, func() bool {
		return f.logins.Load() >= 2
	})
}

func TestPrivateWs_PongKeepsConnectionAlive(t *testing.T) {
	setTestTimings(t, 200*time.Millisecond, 2*time.Second)
	f := newFakeOkxServer(t)
	ws, _ := connectOrders(t)
	defer ws.Close()

	time.Sleep(1500 * time.Millisecond)
	if n := f.logins.Load(); n != 1 {
		t.Fatalf("connection was reconnected %d times although pongs arrived", n)
	}
}

func TestPrivateWs_FailedSubscribeReleasesLock(t *testing.T) {
	setTestTimings(t, 10*time.Second, 300*time.Millisecond)
	f := newFakeOkxServer(t)
	ws := (&MyOkx{}).NewPrivateWsStreamClient()
	if err := ws.OpenConn(); err != nil {
		t.Fatalf("OpenConn: %v", err)
	}
	defer ws.Close()
	if err := ws.Login((&MyOkx{}).NewRestClient("k", "s", "p")); err != nil {
		t.Fatalf("Login: %v", err)
	}

	f.ackSub.Store(false)
	if _, err := ws.SubscribeOrders("ANY", "", ""); err == nil {
		t.Fatal("expected a timeout error when the server never acks")
	}

	f.ackSub.Store(true)
	if _, err := ws.SubscribeOrders("ANY", "", ""); err != nil {
		t.Fatalf("second subscribe must not be blocked by the failed one: %v", err)
	}
}
