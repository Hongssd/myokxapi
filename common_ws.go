package myokxapi

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
)

const (
	OKX_API_WS_PUBLIC   = "/ws/v5/public"   //公共频道
	OKX_API_WS_PRIVATE  = "/ws/v5/private"  //私有频道
	OKX_API_WS_BUSINESS = "/ws/v5/business" //业务频道
)

const (
	LOGIN              = "login"               //登陆
	SUBSCRIBE          = "subscribe"           //订阅
	UNSUBSCRIBE        = "unsubscribe"         //取消订阅
	ORDER              = "order"               //下单
	BATCH_ORDER        = "batch-orders"        //批量下单
	CANCEL_ORDER       = "cancel-order"        //撤单
	BATCH_CANCEL_ORDER = "batch-cancel-orders" //批量撤单
	AMEND_ORDER        = "amend-order"         //改单
	BATCH_AMEND_ORDER  = "batch-amend-orders"  //批量改单
)

var (
	// WebsocketTimeout is the text ping interval. The connection is closed (and then
	// reconnected) after 3*WebsocketTimeout without any text frame, pong included.
	WebsocketTimeout        = time.Second * 10
	WebsocketKeepalive      = true
	SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔

	// WsActionTimeout bounds the wait for a login/subscribe/order ack, so a lost ack
	// cannot hold the action lock forever.
	WsActionTimeout = 10 * time.Second
	// wsWriteTimeout keeps a stuck socket from holding the write lock.
	wsWriteTimeout = 10 * time.Second
	// wsPushSendTimeout bounds a push hand-off to a stalled consumer. Without it one
	// full consumer channel would stop all reads on the connection.
	wsPushSendTimeout = 5 * time.Second
	// wsResubscribeMaxAttempts is the retry count of the post-reconnect resubscribe.
	wsResubscribeMaxAttempts = 5
)

// 数据流订阅基础客户端
type WsStreamClient struct {
	client          *RestClient
	lastLogin       *WsLoginReq
	apiType         APIType
	conn            *websocket.Conn
	connId          string
	commonSubMap    MySyncMap[string, *Subscription[WsActionResult]] //订阅的返回结果
	waitSubResult   *Subscription[WsActionResult]                    //订阅结果
	waitSubResultMu *sync.Mutex                                      //订阅结果锁

	waitOrderResult   *Subscription[WsOrderResult] //订单操作结果
	waitOrderResultMu *sync.Mutex                  //订单操作结果锁

	candleSubMap       MySyncMap[string, *Subscription[WsCandles]]      //K线推送订阅频道
	booksSubMap        MySyncMap[string, *Subscription[WsBooks]]        //深度推送订阅频道
	tradesSubMap       MySyncMap[string, *Subscription[WsTrades]]       //成交流推送订阅频道
	allTradesSubMap    MySyncMap[string, *Subscription[WsAllTrades]]    //全部成交流推送订阅频道
	optSummarySubMap   MySyncMap[string, *Subscription[WsOptSummary]]   //期权定价频道
	markPriceSubMap    MySyncMap[string, *Subscription[WsMarkPrice]]    //标记价格频道
	indexTickersSubMap MySyncMap[string, *Subscription[WsIndexTickers]] //指数行情频道
	tickersSubMap      MySyncMap[string, *Subscription[WsTickers]]      //ticker行情频道
	ordersSubMap       MySyncMap[string, *Subscription[WsOrders]]       //订单推送订阅频道
	ordersAlgoSubMap   MySyncMap[string, *Subscription[WsOrdersAlgo]]   //策略委托订单频道

	resultChan chan []byte
	errChan    chan error
	isClose    bool

	reSubscribeMu      *sync.Mutex
	AutoReConnectTimes int //自动重连次数

	writeMu      sync.Mutex   // gorilla/websocket allows a single concurrent writer
	lastMsgNano  atomic.Int64 // last text frame read from the current connection
	reconnectGen atomic.Int64 // a newer reconnect cancels the older relogin/resubscribe
}

// 登陆请求相关
type WsLoginReq struct {
	Op   string       `json:"op"`   //String 是操作
	Args []WsLoginArg `json:"args"` //Array 是请求订阅的频道列表
}

type WsLoginArg struct {
	ApiKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

// 订阅请求相关
type WsSubscribeReq struct {
	Op   string           `json:"op"`   //String 是操作
	Args []WsSubscribeArg `json:"args"` //Array 是请求订阅的频道列表
}
type WsSubscribeArg struct {
	Channel    string `json:"channel"`              //频道名
	InstType   string `json:"instType,omitempty"`   //String 否 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION ： 期权 ANY： 全部
	InstFamily string `json:"instFamily,omitempty"` //String 否 交易品种 适用于交割/永续/期权
	InstId     string `json:"instId,omitempty"`     //String 是 产品ID 交易对
}

// 登陆及订阅返回结果
type WsActionResult struct {
	Event  string         `json:"event"`         //事件，subscribe error
	Arg    WsSubscribeArg `json:"arg,omitempty"` //订阅参数
	Code   string         `json:"code"`          //错误码
	Msg    string         `json:"msg"`           //错误消息
	ConnId string         `json:"connId"`        //WebSocket连接ID
}

// 订单请求相关
type WsOrderArg[T any] struct {
	Id      string `json:"id"`                //String	是	消息的唯一标识 用户提供，返回参数中会返回以便于找到相应的请求。 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度必须要在1-32位之间。
	Op      string `json:"op"`                //String	是	支持的业务操作，如 order
	Args    []T    `json:"args"`              //Array	是	请求参数
	ExpTime string `json:"expTime,omitempty"` //String	否	请求有效截止时间。Unix时间戳的毫秒数格式，如 1597026383085
}

// 订单操作返回结果
type WsOrderResult struct {
	Id      string              `json:"id"`      //消息的唯一标识
	Op      string              `json:"op"`      //业务操作
	Code    string              `json:"code"`    //代码
	Msg     string              `json:"msg"`     //消息
	Data    []WsOrderResultData `json:"data"`    //请求成功后返回的数据
	InTime  string              `json:"inTime"`  //WebSocket 网关接收请求时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123
	OutTime string              `json:"outTime"` //WebSocket 网关发送响应时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123
}

type WsOrderResultData struct {
	OrdId   string `json:"ordId"`   //订单ID
	ClOrdId string `json:"clOrdId"` //用户提供的订单ID
	Tag     string `json:"tag"`     //订单标签
	ReqId   string `json:"reqId"`   //用户提供的reqId
	SCode   string `json:"sCode"`   //订单状态码，0 代表成功
	SMsg    string `json:"sMsg"`    //订单状态消息
}

// 数据流订阅标准结构体
type Subscription[T any] struct {
	SubId        int64            //订阅ID
	Ws           *WsStreamClient  //订阅的连接
	Op           string           //订阅方法
	Args         []WsSubscribeArg //订阅参数
	resultChan   chan T           //接收订阅结果的通道
	errChan      chan error       //接收订阅错误的通道
	closeChan    chan struct{}    //接收订阅关闭的通道
	subResultMap map[string]bool  //订阅结果
}

// 获取订阅结果
func (sub *Subscription[T]) ResultChan() chan T {
	return sub.resultChan
}

// 获取错误订阅
func (sub *Subscription[T]) ErrChan() chan error {
	return sub.errChan
}

// 获取关闭订阅信号
func (sub *Subscription[T]) CloseChan() chan struct{} {
	return sub.closeChan
}

func (ws *WsStreamClient) GetConn() *websocket.Conn {
	return ws.conn
}

func (ws *WsStreamClient) login(op string, arg WsLoginArg) (*Subscription[WsActionResult], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if ws.waitSubResult != nil {
		return nil, fmt.Errorf("websocket is busy")
	}
	loginReq := WsLoginReq{
		Op:   op,
		Args: []WsLoginArg{arg},
	}
	data, err := json.Marshal(loginReq)
	if err != nil {
		return nil, err
	}
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	result := &Subscription[WsActionResult]{
		SubId:        id,
		Op:           op,
		Args:         []WsSubscribeArg{},
		resultChan:   make(chan WsActionResult, 50),
		errChan:      make(chan error, 1),
		closeChan:    make(chan struct{}),
		Ws:           ws,
		subResultMap: map[string]bool{},
	}

	ws.waitSubResultMu.Lock()
	// Register the waiter before the write so a fast ack is not dropped.
	ws.waitSubResult = result
	log.Debugf("send msg: %s", string(data))
	if err := ws.writeText(data); err != nil {
		ws.waitSubResult = nil
		ws.waitSubResultMu.Unlock()
		return nil, err
	}
	ws.lastLogin = &loginReq

	return result, nil
}

func orderAction[T any](ws *PrivateWsStreamClient, op string, reqId string, args []T, expTimestamp int64) (*WsOrderResult, error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if ws.waitOrderResult != nil {
		return nil, fmt.Errorf("websocket is busy")
	}
	orderReq := WsOrderArg[T]{
		Id:      reqId,
		Op:      op,
		Args:    args,
		ExpTime: strconv.FormatInt(expTimestamp, 10),
	}
	data, err := json.Marshal(orderReq)
	if err != nil {
		return nil, err
	}
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	resultSub := &Subscription[WsOrderResult]{
		SubId:        id,
		Op:           op,
		Args:         []WsSubscribeArg{},
		resultChan:   make(chan WsOrderResult, 50),
		errChan:      make(chan error, 1),
		closeChan:    make(chan struct{}),
		Ws:           &ws.WsStreamClient,
		subResultMap: map[string]bool{},
	}

	ws.waitOrderResultMu.Lock()
	ws.waitOrderResult = resultSub
	log.Debugf("send msg: %s", string(data))
	if err := ws.writeText(data); err != nil {
		ws.waitOrderResult = nil
		ws.waitOrderResultMu.Unlock()
		return nil, err
	}

	return ws.CatchOrderReuslt(resultSub)
}

func subscribe[T any](ws *WsStreamClient, op string, args []WsSubscribeArg) (*Subscription[T], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if ws.waitSubResult != nil {
		return nil, fmt.Errorf("websocket is busy")
	}

	subscribeReq := WsSubscribeReq{
		Op:   op,
		Args: args,
	}
	data, err := json.Marshal(subscribeReq)
	if err != nil {
		return nil, err
	}
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	result := &Subscription[T]{
		SubId:        id,
		Op:           op,
		Args:         args,
		resultChan:   make(chan T, 50),
		errChan:      make(chan error, 1),
		closeChan:    make(chan struct{}),
		Ws:           ws,
		subResultMap: map[string]bool{},
	}

	// The lock must be released on every failure after this point, or the
	// connection stays "busy" forever.
	ws.waitSubResultMu.Lock()
	// Register the waiter before the write so a fast ack is not dropped.
	if waiter, ok := any(result).(*Subscription[WsActionResult]); ok {
		ws.waitSubResult = waiter
	}
	log.Debugf("send msg: %s", string(data))
	if err := ws.writeText(data); err != nil {
		ws.waitSubResult = nil
		ws.waitSubResultMu.Unlock()
		return nil, err
	}
	return result, nil
}

func (ws *WsStreamClient) Close() error {
	ws.isClose = true
	ws.connId = ""

	err := ws.conn.Close()
	if err != nil {
		return err
	}
	//手动关闭成功，给所有订阅发送关闭信号
	ws.sendWsCloseToAllSub()

	//初始化连接状态
	ws.conn = nil
	close(ws.resultChan)
	close(ws.errChan)
	ws.resultChan = nil
	ws.errChan = nil
	ws.lastLogin = nil
	ws.commonSubMap = NewMySyncMap[string, *Subscription[WsActionResult]]()
	ws.candleSubMap = NewMySyncMap[string, *Subscription[WsCandles]]()
	ws.booksSubMap = NewMySyncMap[string, *Subscription[WsBooks]]()
	ws.tradesSubMap = NewMySyncMap[string, *Subscription[WsTrades]]()
	ws.allTradesSubMap = NewMySyncMap[string, *Subscription[WsAllTrades]]()
	ws.optSummarySubMap = NewMySyncMap[string, *Subscription[WsOptSummary]]()
	ws.markPriceSubMap = NewMySyncMap[string, *Subscription[WsMarkPrice]]()
	ws.indexTickersSubMap = NewMySyncMap[string, *Subscription[WsIndexTickers]]()
	ws.tickersSubMap = NewMySyncMap[string, *Subscription[WsTickers]]()
	ws.ordersSubMap = NewMySyncMap[string, *Subscription[WsOrders]]()
	ws.ordersAlgoSubMap = NewMySyncMap[string, *Subscription[WsOrdersAlgo]]()

	// The waiter owns the action lock and releases it in its deferred cleanup, so
	// closing must only wake it. Taking the lock here would deadlock.
	if w := ws.waitSubResult; w != nil {
		select {
		case w.errChan <- fmt.Errorf("websocket is closed"):
		default:
		}
	}
	if w := ws.waitOrderResult; w != nil {
		select {
		case w.errChan <- fmt.Errorf("websocket is closed"):
		default:
		}
	}
	return nil
}

// writeText is the only way to write a text frame. It serializes writers and
// bounds the time a dead socket can hold the write lock.
func (ws *WsStreamClient) writeText(data []byte) error {
	conn := ws.conn
	if conn == nil {
		return fmt.Errorf("websocket is close")
	}
	return ws.writeTextTo(conn, data)
}

func (ws *WsStreamClient) writeTextTo(conn *websocket.Conn, data []byte) error {
	ws.writeMu.Lock()
	defer ws.writeMu.Unlock()
	_ = conn.SetWriteDeadline(time.Now().Add(wsWriteTimeout))
	return conn.WriteMessage(websocket.TextMessage, data)
}

// LastMessageTime is the arrival time of the last text frame on the current
// connection. Pong replies count, so it advances on an idle but healthy link.
func (ws *WsStreamClient) LastMessageTime() time.Time {
	return time.Unix(0, ws.lastMsgNano.Load())
}

// ForceReconnect drops the current connection. The read loop then fails and the
// normal reconnect path logs in again and resubscribes every channel.
func (ws *WsStreamClient) ForceReconnect(reason string) error {
	conn := ws.conn
	if conn == nil || ws.isClose {
		return fmt.Errorf("websocket is close")
	}
	log.Warnf("force reconnect okx ws: %s", reason)
	return conn.Close()
}

func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan == nil {
		ws.resultChan = make(chan []byte)
	}
	if ws.errChan == nil {
		ws.errChan = make(chan error)
	}
	apiUrl := handlerWsStreamRequestApi(ws.apiType)
	if ws.conn == nil {
		conn, err := wsStreamServe(ws, apiUrl, ws.resultChan, ws.errChan)
		if err != nil {
			return err
		}
		ws.conn = conn
		ws.isClose = false
		ws.connId = ""
		log.Info("OpenConn success to ", apiUrl)
		ws.handleResult(ws.resultChan, ws.errChan)
		return nil
	}
	// A failed redial must keep the old conn. Clearing it would make the next
	// attempt look like a first open and start a second handleResult loop.
	conn, err := wsStreamServe(ws, apiUrl, ws.resultChan, ws.errChan)
	if err != nil {
		return err
	}
	ws.conn = conn
	ws.connId = ""
	log.Info("Auto ReOpenConn success to ", apiUrl)
	return nil
}

type PublicWsStreamClient struct {
	WsStreamClient
}
type PrivateWsStreamClient struct {
	WsStreamClient
}
type BusinessWsStreamClient struct {
	WsStreamClient
}

func (*MyOkx) NewPublicWsStreamClient() *PublicWsStreamClient {
	return &PublicWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:            WS_PUBLIC,
			commonSubMap:       NewMySyncMap[string, *Subscription[WsActionResult]](),
			candleSubMap:       NewMySyncMap[string, *Subscription[WsCandles]](),
			booksSubMap:        NewMySyncMap[string, *Subscription[WsBooks]](),
			tradesSubMap:       NewMySyncMap[string, *Subscription[WsTrades]](),
			optSummarySubMap:   NewMySyncMap[string, *Subscription[WsOptSummary]](),
			markPriceSubMap:    NewMySyncMap[string, *Subscription[WsMarkPrice]](),
			indexTickersSubMap: NewMySyncMap[string, *Subscription[WsIndexTickers]](),
			tickersSubMap:      NewMySyncMap[string, *Subscription[WsTickers]](),
			ordersSubMap:       NewMySyncMap[string, *Subscription[WsOrders]](),
			ordersAlgoSubMap:   NewMySyncMap[string, *Subscription[WsOrdersAlgo]](),
			waitSubResult:      nil,
			waitSubResultMu:    &sync.Mutex{},
			waitOrderResult:    nil,
			waitOrderResultMu:  &sync.Mutex{},
			reSubscribeMu:      &sync.Mutex{},
		},
	}
}
func (*MyOkx) NewPrivateWsStreamClient() *PrivateWsStreamClient {
	return &PrivateWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:            WS_PRIVATE,
			commonSubMap:       NewMySyncMap[string, *Subscription[WsActionResult]](),
			candleSubMap:       NewMySyncMap[string, *Subscription[WsCandles]](),
			booksSubMap:        NewMySyncMap[string, *Subscription[WsBooks]](),
			tradesSubMap:       NewMySyncMap[string, *Subscription[WsTrades]](),
			optSummarySubMap:   NewMySyncMap[string, *Subscription[WsOptSummary]](),
			markPriceSubMap:    NewMySyncMap[string, *Subscription[WsMarkPrice]](),
			indexTickersSubMap: NewMySyncMap[string, *Subscription[WsIndexTickers]](),
			tickersSubMap:      NewMySyncMap[string, *Subscription[WsTickers]](),
			ordersSubMap:       NewMySyncMap[string, *Subscription[WsOrders]](),
			ordersAlgoSubMap:   NewMySyncMap[string, *Subscription[WsOrdersAlgo]](),
			waitSubResult:      nil,
			waitSubResultMu:    &sync.Mutex{},
			waitOrderResult:    nil,
			waitOrderResultMu:  &sync.Mutex{},
			reSubscribeMu:      &sync.Mutex{},
		},
	}
}
func (*MyOkx) NewBusinessWsStreamClient() *BusinessWsStreamClient {
	return &BusinessWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:            WS_BUSINESS,
			commonSubMap:       NewMySyncMap[string, *Subscription[WsActionResult]](),
			candleSubMap:       NewMySyncMap[string, *Subscription[WsCandles]](),
			booksSubMap:        NewMySyncMap[string, *Subscription[WsBooks]](),
			tradesSubMap:       NewMySyncMap[string, *Subscription[WsTrades]](),
			allTradesSubMap:    NewMySyncMap[string, *Subscription[WsAllTrades]](),
			optSummarySubMap:   NewMySyncMap[string, *Subscription[WsOptSummary]](),
			markPriceSubMap:    NewMySyncMap[string, *Subscription[WsMarkPrice]](),
			indexTickersSubMap: NewMySyncMap[string, *Subscription[WsIndexTickers]](),
			tickersSubMap:      NewMySyncMap[string, *Subscription[WsTickers]](),
			ordersSubMap:       NewMySyncMap[string, *Subscription[WsOrders]](),
			ordersAlgoSubMap:   NewMySyncMap[string, *Subscription[WsOrdersAlgo]](),
			waitSubResult:      nil,
			waitSubResultMu:    &sync.Mutex{},
			waitOrderResult:    nil,
			waitOrderResultMu:  &sync.Mutex{},
			reSubscribeMu:      &sync.Mutex{},
		},
	}
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result WsActionResult) {
	if ws.connId == "" && result.ConnId != "" {
		ws.connId = result.ConnId
	}
	// Connection-level events carry no arg. Routing them to a waiter would be
	// counted as a subscribe ack, so they are handled here.
	switch result.Event {
	case "notice":
		log.Warnf("okx ws notice: code=%s msg=%s connId=%s", result.Code, result.Msg, result.ConnId)
		go func() { _ = ws.ForceReconnect("server notice " + result.Code) }()
		return
	case "channel-conn-count":
		log.Debugf("okx ws channel-conn-count: connId=%s", result.ConnId)
		return
	case "channel-conn-count-error":
		log.Errorf("okx ws channel-conn-count-error: code=%s msg=%s connId=%s", result.Code, result.Msg, result.ConnId)
		return
	}

	w := ws.waitSubResult
	if w == nil {
		if result.Event == "error" || (result.Code != "" && result.Code != "0") {
			log.Warnf("okx ws event without waiter: %+v", result)
		}
		return
	}
	// Non-blocking: a waiter that already timed out must not stall the read path.
	if result.Code != "" && result.Code != "0" {
		select {
		case w.errChan <- fmt.Errorf("errHandler: %+v", result):
		default:
			log.Warnf("okx ws drop error event, waiter busy: %+v", result)
		}
		return
	}
	select {
	case w.resultChan <- result:
	default:
		log.Warnf("okx ws drop event, waiter full: %+v", result)
	}
}

func (ws *WsStreamClient) sendOrderResultToChan(result WsOrderResult) {
	w := ws.waitOrderResult
	if w == nil {
		return
	}
	if result.Code != "" && result.Code != "0" {
		select {
		case w.errChan <- fmt.Errorf("errHandler: %+v", result):
		default:
		}
	}
	select {
	case w.resultChan <- result:
	default:
	}
}

// signalClose delivers a close signal without blocking the caller. The channels
// are unbuffered, so a blocking send made Close() hang forever whenever the
// consumer had stopped reading.
func signalClose(ch chan struct{}) {
	if ch == nil {
		return
	}
	go func() {
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()
		select {
		case ch <- struct{}{}:
		case <-timer.C:
		}
	}()
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(args []WsSubscribeArg) {
	for _, arg := range args {
		data, _ := json.Marshal(&arg)
		key := string(data)
		if sub, ok := ws.candleSubMap.Load(key); ok {
			ws.candleSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.booksSubMap.Load(key); ok {
			ws.booksSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.tradesSubMap.Load(key); ok {
			ws.tradesSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.allTradesSubMap.Load(key); ok {
			ws.allTradesSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.ordersSubMap.Load(key); ok {
			ws.ordersSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.ordersAlgoSubMap.Load(key); ok {
			ws.ordersAlgoSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.optSummarySubMap.Load(key); ok {
			ws.optSummarySubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.markPriceSubMap.Load(key); ok {
			ws.markPriceSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.indexTickersSubMap.Load(key); ok {
			ws.indexTickersSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.tickersSubMap.Load(key); ok {
			ws.tickersSubMap.Delete(key)
			if sub.closeChan != nil {
				signalClose(sub.closeChan)
				sub.closeChan = nil
			}
		}
	}
}

func (ws *WsStreamClient) sendWsCloseToAllSub() {
	args := []WsSubscribeArg{}
	ws.commonSubMap.Range(func(key string, _ *Subscription[WsActionResult]) bool {
		arg := WsSubscribeArg{}
		err := json.Unmarshal([]byte(key), &arg)
		if err != nil {
			return false
		}
		args = append(args, arg)
		return true
	})
	ws.sendUnSubscribeSuccessToCloseChan(args)
}

// reSubscribeForReconnect resubscribes every recorded channel. done carries the
// channels already restored, so a retry after a partial failure skips them.
// A successful subscribe replaces the commonSubMap entry with the new
// subscription, so both pointers are recorded.
func (ws *WsStreamClient) reSubscribeForReconnect(done map[*Subscription[WsActionResult]]struct{}) error {
	ws.reSubscribeMu.Lock()
	defer ws.reSubscribeMu.Unlock()
	var wErr error
	ws.commonSubMap.Range(func(_ string, sub *Subscription[WsActionResult]) bool {
		if _, ok := done[sub]; ok {
			return true
		}

		reSub, err := subscribe[WsActionResult](ws, sub.Op, sub.Args)
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}
		err = ws.catchSubscribeResult(reSub)
		if err != nil {
			log.Error(err)
			wErr = err
			return false
		}
		log.Infof("reSubscribe Success: args:%v", reSub.Args)

		sub.SubId = reSub.SubId
		done[sub] = struct{}{}
		done[reSub] = struct{}{}
		time.Sleep(500 * time.Millisecond)
		return true
	})
	return wErr
}

// restoreAfterReconnect logs in again and resubscribes. A newer reconnect
// (gen moved on) or a manual Close cancels it. A resubscribe that keeps failing
// would leave a silently dead subscription, so it drops the connection and
// lets the reconnect cycle start over.
func (ws *WsStreamClient) restoreAfterReconnect(gen int64) {
	superseded := func() bool {
		return ws.isClose || ws.reconnectGen.Load() != gen
	}
	if ws.lastLogin != nil && ws.client != nil {
		err := ws.Login(ws.client)
		for err != nil {
			log.Error(err)
			time.Sleep(1500 * time.Millisecond)
			if superseded() {
				return
			}
			err = ws.Login(ws.client)
		}
	}

	done := map[*Subscription[WsActionResult]]struct{}{}
	for attempt := 1; ; attempt++ {
		if superseded() {
			return
		}
		err := ws.reSubscribeForReconnect(done)
		if err == nil {
			return
		}
		log.Errorf("reSubscribe failed, attempt %d/%d: %v", attempt, wsResubscribeMaxAttempts, err)
		if attempt >= wsResubscribeMaxAttempts {
			_ = ws.ForceReconnect("resubscribe exhausted")
			return
		}
		time.Sleep(time.Duration(attempt) * 1500 * time.Millisecond)
	}
}

// loadOrdersSub finds the subscription for an orders push. The exact arg is tried
// first. The push may carry instId/instFamily that the subscribe request did not
// have (it is subscribed per instType), so coarser keys are tried next instead
// of dropping the message.
func (ws *WsStreamClient) loadOrdersSub(arg WsSubscribeArg) (*Subscription[WsOrders], bool) {
	candidates := []WsSubscribeArg{
		arg,
		{Channel: arg.Channel, InstType: arg.InstType, InstFamily: arg.InstFamily},
		{Channel: arg.Channel, InstType: arg.InstType},
	}
	for _, c := range candidates {
		keyData, _ := json.Marshal(c)
		if sub, ok := ws.ordersSubMap.Load(string(keyData)); ok {
			return sub, true
		}
	}
	return nil, false
}

// deliverOrders hands a push to the consumer. A consumer that stops reading
// would otherwise block the read loop for every channel on this connection.
func deliverOrders(sub *Subscription[WsOrders], order WsOrders) {
	select {
	case sub.resultChan <- order:
		return
	default:
	}
	timer := time.NewTimer(wsPushSendTimeout)
	defer timer.Stop()
	select {
	case sub.resultChan <- order:
	case <-timer.C:
		log.Errorf("orders consumer stalled, push dropped: clOrdId=%s ordId=%s state=%s",
			order.ClOrdId, order.OrdId, order.State)
	}
}

func (ws *WsStreamClient) handleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			select {
			case err, ok := <-errChan:
				if !ok {
					log.Error("errChan is closed")
					return
				}
				log.Error(err)
				// The read loop exits on any read error, so the connection is dead
				// whatever the message says. Matching on EOF/close/reset missed
				// others (broken pipe, i/o timeout) and left it dead for good.
				if ws.isClose {
					continue
				}
				gen := ws.reconnectGen.Add(1)
				err = ws.OpenConn()
				for err != nil {
					log.Error(err)
					time.Sleep(1500 * time.Millisecond)
					if ws.isClose {
						return
					}
					err = ws.OpenConn()
				}
				ws.AutoReConnectTimes += 1
				go ws.restoreAfterReconnect(gen)
			case data, ok := <-resultChan:
				if !ok {
					log.Error("resultChan is closed")
					return
				}
				// log.Debug("receive result: ", string(data))
				//处理订阅或查询订阅列表请求返回结果
				if strings.Contains(string(data), "event") || strings.Contains(string(data), "connId") {
					result := WsActionResult{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.sendSubscribeResultToChan(result)
					continue
				}

				if strings.Contains(string(data), "id") && strings.Contains(string(data), "op") && strings.Contains(string(data), "code") {
					result := WsOrderResult{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.sendOrderResultToChan(result)
					continue
				}

				//处理正常数据的返回结果
				//K线处理
				if strings.Contains(string(data), "candle") {
					c, err := handleWsCandle(data)

					arg := c.WsSubscribeArg
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.candleSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}
				if strings.Contains(string(data), "books") || strings.Contains(string(data), "bbo-tbt") {
					b, err := handleWsBooks(data)

					arg := b.WsSubscribeArg
					keyData, _ := json.Marshal(arg)

					sub, ok := ws.booksSubMap.Load(string(keyData))

					if ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *b
					}
					continue
				}
				if strings.Contains(string(data), "trades-all") {
					t, err := handleWsAllTrades(data)
					arg := t.WsSubscribeArg
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.allTradesSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *t
					}
					continue
				}
				if strings.Contains(string(data), "trades") {
					t, err := handleWsTrades(data)
					arg := t.WsSubscribeArg
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.tradesSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *t
					}
					continue
				}
				if strings.Contains(string(data), "channel\":\"orders\"") {
					ordersList, err := handleWsOrders(data)
					if err != nil {
						// ordersList is nil here, so it must not be dereferenced.
						log.Errorf("parse orders push failed: %v", err)
						continue
					}
					if len(*ordersList) == 0 {
						continue
					}
					arg := (*ordersList)[0].WsSubscribeArg
					sub, ok := ws.loadOrdersSub(arg)
					if !ok {
						log.Warnf("orders push has no matching subscription, dropped: arg=%+v", arg)
						continue
					}
					for _, order := range *ordersList {
						deliverOrders(sub, order)
					}
					continue
				}

				if strings.Contains(string(data), "channel\":\"orders-algo\"") {
					ordersAlgoList, err := handleWsOrdersAlgo(data)
					if err != nil {
						log.Error(err)
						continue
					}
					if len(*ordersAlgoList) == 0 {
						continue
					}
					arg := (*ordersAlgoList)[0].WsSubscribeArg
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.ordersAlgoSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, order := range *ordersAlgoList {
							//d, _ := json.Marshal(order)
							//log.Debug("orders-algo: ", string(d))
							sub.resultChan <- order
						}
					}
					continue
				}

				if strings.Contains(string(data), "channel\":\"opt-summary\"") {
					optSummaryList, err := handleWsOptSummary(data)
					if len(*optSummaryList) == 0 {
						continue
					}
					arg := (*optSummaryList)[0].Args
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.optSummarySubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, optSummary := range *optSummaryList {
							sub.resultChan <- optSummary
						}
					}
					continue
				}

				if strings.Contains(string(data), "channel\":\"mark-price\"") {
					markPriceList, err := handleWsMarkPrice(data)
					if len(*markPriceList) == 0 {
						continue
					}
					arg := (*markPriceList)[0].Args
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.markPriceSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, markPrice := range *markPriceList {
							sub.resultChan <- markPrice
						}
					}
					continue
				}

				if strings.Contains(string(data), "channel\":\"index-tickers\"") {
					indexTickersList, err := handleWsIndexTickers(data)
					if len(*indexTickersList) == 0 {
						continue
					}
					arg := (*indexTickersList)[0].Args
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.indexTickersSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, indexTickers := range *indexTickersList {
							sub.resultChan <- indexTickers
						}
					}
					continue
				}

				if strings.Contains(string(data), "channel\":\"tickers\"") {
					tickersList, err := handleWsTickers(data)
					if len(*tickersList) == 0 {
						continue
					}
					arg := (*tickersList)[0].Args
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.tickersSubMap.Load(string(keyData)); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, tickers := range *tickersList {
							sub.resultChan <- tickers
						}
					}
					continue
				}
			}
		}
	}()
}

// DeferSub records a fully acknowledged subscription and always releases the
// action lock. Releasing only on success left the connection "busy" after the
// first failed or timed-out subscribe.
func (ws *WsStreamClient) DeferSub() {
	if w := ws.waitSubResult; w != nil && len(w.subResultMap) == len(w.Args) {
		for _, arg := range w.Args {
			keyData, _ := json.Marshal(&arg)
			ws.commonSubMap.Store(string(keyData), w)
		}
	}
	ws.waitSubResult = nil
	ws.waitSubResultMu.Unlock()
}

func (ws *WsStreamClient) DeferOrder() {
	ws.waitOrderResult = nil
	ws.waitOrderResultMu.Unlock()
}

// 取消订阅
func (sub *Subscription[T]) Unsubscribe() error {

	unSub, err := subscribe[WsActionResult](sub.Ws, UNSUBSCRIBE, sub.Args)
	if err != nil {
		return err
	}
	err = sub.Ws.catchSubscribeResult(unSub)
	if err != nil {
		return err
	}
	log.Debugf("Unsubscribe Success args:%v ", unSub.Args)

	//取消订阅成功，给所有订阅消息的通道发送关闭信号
	sub.Ws.sendUnSubscribeSuccessToCloseChan(unSub.Args)
	//删除当前订阅列表中已存在的记录
	for _, arg := range unSub.Args {
		data, _ := json.Marshal(&arg)
		key := string(data)
		sub.Ws.commonSubMap.Delete(key)
	}
	return nil
}

// 捕获订阅结果
func (ws *WsStreamClient) CatchLoginReuslt(sub *Subscription[WsActionResult]) error {
	ws.waitSubResult = sub
	defer sub.Ws.DeferSub()

	timer := time.NewTimer(WsActionTimeout)
	defer timer.Stop()
	select {
	case err := <-sub.ErrChan():
		// log.Error(err)
		return fmt.Errorf("Login Error: %v", err)
	case loginResult := <-sub.ResultChan():
		if loginResult.Code != "" && loginResult.Code != "0" {
			log.Error(loginResult.Code, ":", loginResult.Msg)
			return fmt.Errorf("%s:%s", loginResult.Code, loginResult.Msg)
		}
		log.Debug("catchLoginResults: ", loginResult)
		return nil
	case <-timer.C:
		return fmt.Errorf("Login Error: wait result timeout")
	}
}

// 捕获订阅结果
func (ws *WsStreamClient) catchSubscribeResult(sub *Subscription[WsActionResult]) error {
	ws.waitSubResult = sub
	defer sub.Ws.DeferSub()
	timer := time.NewTimer(WsActionTimeout)
	defer timer.Stop()
	for {
		select {
		case err := <-sub.ErrChan():
			log.Error(err)
			return fmt.Errorf("SubAction Error: %v", err)
		case subResult := <-sub.ResultChan():
			if subResult.Code != "" && subResult.Code != "0" {
				log.Error(subResult.Code, ":", subResult.Msg)
				return fmt.Errorf("%s:%s", subResult.Code, subResult.Msg)
			}
			keyData, _ := json.Marshal(subResult.Arg)
			sub.subResultMap[string(keyData)] = true
			if len(sub.subResultMap) == len(sub.Args) {
				log.Debug("catchResults: ", sub.subResultMap)
				return nil
			}
		case <-timer.C:
			return fmt.Errorf("SubAction Error: wait subscribe result timeout, op:%s args:%v", sub.Op, sub.Args)
		}
	}
}

// 捕获订单操作结果
func (ws *PrivateWsStreamClient) CatchOrderReuslt(sub *Subscription[WsOrderResult]) (*WsOrderResult, error) {
	ws.waitOrderResult = sub
	defer sub.Ws.DeferOrder()
	timer := time.NewTimer(WsActionTimeout)
	defer timer.Stop()
	select {
	case err := <-sub.ErrChan():
		select {
		case orderResult := <-sub.ResultChan():
			return &orderResult, fmt.Errorf("[op:%s] Action Error: %v", sub.Op, err)
		case <-timer.C:
			return nil, fmt.Errorf("[op:%s] Action Error: %v", sub.Op, err)
		}
	case orderResult := <-sub.ResultChan():
		return &orderResult, nil
	case <-timer.C:
		return nil, fmt.Errorf("[op:%s] Action Error: wait order result timeout", sub.Op)
	}
}

// okx websocket登陆功能
func (ws *WsStreamClient) Login(client *RestClient) error {
	ws.client = client
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	method := "GET"
	requestPath := "/users/self/verify"
	hmacSha256Data := timestamp + method + requestPath
	sign := base64.StdEncoding.EncodeToString(HmacSha256(client.c.SecretKey, hmacSha256Data))
	loginArg := WsLoginArg{
		ApiKey:     client.c.APIKey,
		Passphrase: client.c.Passphrase,
		Timestamp:  timestamp,
		Sign:       sign,
	}
	loginSub, err := ws.login(LOGIN, loginArg)
	if err != nil {
		return err
	}

	err = ws.CatchLoginReuslt(loginSub)
	if err != nil {
		return err
	}
	log.Infof("Login Success args:%v ", ws.lastLogin.Args)
	return nil
}

// 标准订阅方法
func wsStreamServe(ws *WsStreamClient, api string, resultChan chan []byte, errChan chan error) (*websocket.Conn, error) {
	dialer := websocket.DefaultDialer
	if WsUseProxy {
		proxy, _ := getBestProxyAndWeight()
		if proxy == nil {
			return nil, errors.New("no proxy available")
		}
		proxyUrl, err := url.Parse(proxy.ProxyUrl)
		if err != nil {
			return nil, err
		}
		dialer = &websocket.Dialer{
			Proxy: http.ProxyURL(proxyUrl),
		}
	}
	c, _, err := dialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(655350)
	ws.lastMsgNano.Store(time.Now().UnixNano())
	if WebsocketKeepalive {
		keepAlive(ws, c, WebsocketTimeout)
	}
	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			ws.lastMsgNano.Store(time.Now().UnixNano())
			// The reply to the text "ping" only proves liveness.
			if string(message) == "pong" {
				continue
			}
			resultChan <- message
		}
	}()
	return c, err
}

// 获取数据流请求URL
func handlerWsStreamRequestApi(apiType APIType) string {
	host := OKX_API_WEBSOCKET
	switch SERVER_TYPE {
	case BASE:
		host = OKX_API_WEBSOCKET
	case AWS:
		host = OKX_API_WEBSOCKET_AWS
	case COLO:
		switch apiType {
		case WS_PRIVATE, WS_PUBLIC:
			host = OKX_API_WEBSOCKET_COLO_PUBLIC_AND_PRIVATE
		case WS_BUSINESS:
			host = OKX_API_WEBSOCKET_COLO_BUSINESS
		}
	default:
	}

	query := ""
	if NowNetType == TEST_NET {
		// query = "brokerId=9999"
		host = OKX_API_WEBSOCKET_TESTNET
	}
	u := url.URL{
		Scheme:   "wss",
		Host:     host,
		Path:     getWsApi(apiType),
		RawQuery: query,
	}
	return u.String()
}

// 获取数据流请求Path
func getWsApi(apiType APIType) string {
	switch apiType {
	case WS_PUBLIC:
		return OKX_API_WS_PUBLIC
	case WS_PRIVATE:
		return OKX_API_WS_PRIVATE
	case WS_BUSINESS:
		return OKX_API_WS_BUSINESS
	default:
		log.Error("apiType Error is ", apiType)
		return ""
	}
}

// keepAlive sends OKX's text "ping" when the link has been quiet and closes the
// connection when no text frame (data or "pong") arrived for 3*timeout. The
// protocol-level pong is not trusted: it can be answered by an intermediate hop
// while the upstream is gone. Closing makes the read loop fail, which triggers
// the reconnect path.
func keepAlive(ws *WsStreamClient, c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	go func() {
		defer ticker.Stop()
		for {
			idle := time.Since(time.Unix(0, ws.lastMsgNano.Load()))
			if idle > 3*timeout {
				log.Warnf("okx ws silent for %s, closing to reconnect", idle.Round(time.Second))
				if err := c.Close(); err != nil {
					log.Error(err)
				}
				return
			}
			if idle >= timeout/2 {
				if err := ws.writeTextTo(c, []byte("ping")); err != nil {
					log.Warnf("okx ws ping failed, closing to reconnect: %v", err)
					_ = c.Close()
					return
				}
			}
			// Keep the protocol-level ping as well for intermediaries that expect it.
			_ = c.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second))
			<-ticker.C
		}
	}()
}
