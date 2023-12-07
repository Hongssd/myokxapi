package myokxapi

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
)

const (
	OKX_API_WS_HOST     = "ws.okx.com:8443" //WebSocket请求地址
	OKX_API_WS_PUBLIC   = "/ws/v5/public"   //公共频道
	OKX_API_WS_PRIVATE  = "/ws/v5/private"  //私有频道
	OKX_API_WS_BUSINESS = "/ws/v5/business" //业务频道
)

const (
	LOGIN       = "login"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

var (
	WebsocketTimeout        = time.Second * 60
	WebsocketKeepalive      = true
	SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔
)

// 数据流订阅基础客户端
type WsStreamClient struct {
	client          *RestClient
	lastLogin       *WsLoginReq
	apiType         APIType
	conn            *websocket.Conn
	connId          string
	commonSubMap    map[string]*Subscription[WsActionResult] //订阅的返回结果
	waitSubResult   *Subscription[WsActionResult]
	waitSubResultMu *sync.Mutex
	candleSubMap    map[string]*Subscription[WsCandles]
	booksSubMap     map[string]*Subscription[WsBooks]
	tradesSubMap    map[string]*Subscription[WsTrades]
	ordersSubMap    map[string]*Subscription[WsOrders]
	resultChan      chan []byte
	errChan         chan error
	isClose         bool
}

type WsLoginArg struct {
	ApiKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

type WsLoginReq struct {
	Op   string       `json:"op"`   //String 是操作
	Args []WsLoginArg `json:"args"` //Array 是请求订阅的频道列表
}

type WsSubscribeArg struct {
	Channel    string `json:"channel"`              //频道名
	InstType   string `json:"instType,omitempty"`   //String 否 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION ： 期权 ANY： 全部
	InstFamily string `json:"instFamily,omitempty"` //String 否 交易品种 适用于交割/永续/期权
	InstId     string `json:"instId"`               //String 是 产品ID 交易对
}

// 订阅请求结构体
type WsSubscribeReq struct {
	Op   string           `json:"op"`   //String 是操作
	Args []WsSubscribeArg `json:"args"` //Array 是请求订阅的频道列表
}

// 订阅结果结构体
type WsActionResult struct {
	Event  string         `json:"event"`         //事件，subscribe error
	Arg    WsSubscribeArg `json:"arg,omitempty"` //订阅参数
	Code   string         `json:"code"`          //错误码
	Msg    string         `json:"msg"`           //错误消息
	ConnId string         `json:"connId"`        //WebSocket连接ID
}

// 数据流订阅标准结构体
type Subscription[T any] struct {
	SubId        int64                     //订阅ID
	Ws           *WsStreamClient           //订阅的连接
	Op           string                    //订阅方法
	Args         []WsSubscribeArg          //订阅参数
	resultChan   chan T                    //接收订阅结果的通道
	errChan      chan error                //接收订阅错误的通道
	closeChan    chan struct{}             //接收订阅关闭的通道
	subResultMap map[string]WsActionResult //订阅结果
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

func (ws *WsStreamClient) login(op string, arg WsLoginArg) (*Subscription[WsActionResult], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if ws.waitSubResult != nil {
		return nil, fmt.Errorf("websocket is busy")
	}
	ws.waitSubResultMu.Lock()
	loginReq := WsLoginReq{
		Op:   op,
		Args: []WsLoginArg{arg},
	}
	data, err := json.Marshal(loginReq)
	if err != nil {
		return nil, err
	}
	log.Debugf("send msg: %s", string(data))
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return nil, err
	}
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	id := node.Generate().Int64()
	ws.lastLogin = &loginReq
	result := &Subscription[WsActionResult]{
		SubId:        id,
		Op:           op,
		Args:         []WsSubscribeArg{},
		resultChan:   make(chan WsActionResult, 50),
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		Ws:           ws,
		subResultMap: map[string]WsActionResult{},
	}
	return result, nil
}

func subscribe[T any](ws *WsStreamClient, op string, args []WsSubscribeArg) (*Subscription[T], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}
	if ws.waitSubResult != nil {
		return nil, fmt.Errorf("websocket is busy")
	}

	ws.waitSubResultMu.Lock()

	subscribeReq := WsSubscribeReq{
		Op:   op,
		Args: args,
	}
	data, err := json.Marshal(subscribeReq)
	if err != nil {
		return nil, err
	}
	log.Debugf("send msg: %s", string(data))
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
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
		errChan:      make(chan error),
		closeChan:    make(chan struct{}),
		Ws:           ws,
		subResultMap: map[string]WsActionResult{},
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
	ws.commonSubMap = make(map[string]*Subscription[WsActionResult])
	ws.candleSubMap = make(map[string]*Subscription[WsCandles])
	ws.booksSubMap = make(map[string]*Subscription[WsBooks])
	ws.tradesSubMap = make(map[string]*Subscription[WsTrades])
	ws.ordersSubMap = make(map[string]*Subscription[WsOrders])
	if ws.waitSubResult != nil {
		//给当前等待订阅结果的请求返回错误
		ws.waitSubResultMu.Lock()
		ws.waitSubResult.errChan <- fmt.Errorf("websocket is closed")
		ws.waitSubResult = nil
		ws.waitSubResultMu.Unlock()
	}
	return nil
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
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.isClose = false
		ws.connId = ""
		log.Debug("OpenConn success to ", apiUrl)
		ws.handleResult(ws.resultChan, ws.errChan)
		return err
	} else {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.connId = ""
		log.Debug("Auto ReOpenConn success to ", apiUrl)
		return err
	}

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
			apiType:         WS_PUBLIC,
			commonSubMap:    make(map[string]*Subscription[WsActionResult]),
			candleSubMap:    make(map[string]*Subscription[WsCandles]),
			booksSubMap:     make(map[string]*Subscription[WsBooks]),
			tradesSubMap:    make(map[string]*Subscription[WsTrades]),
			ordersSubMap:    make(map[string]*Subscription[WsOrders]),
			waitSubResult:   nil,
			waitSubResultMu: &sync.Mutex{},
		},
	}
}
func (*MyOkx) NewPrivateWsStreamClient() *PrivateWsStreamClient {
	return &PrivateWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:         WS_PRIVATE,
			commonSubMap:    make(map[string]*Subscription[WsActionResult]),
			candleSubMap:    make(map[string]*Subscription[WsCandles]),
			booksSubMap:     make(map[string]*Subscription[WsBooks]),
			tradesSubMap:    make(map[string]*Subscription[WsTrades]),
			ordersSubMap:    make(map[string]*Subscription[WsOrders]),
			waitSubResult:   nil,
			waitSubResultMu: &sync.Mutex{},
		},
	}
}
func (*MyOkx) NewBusinessWsStreamClient() *BusinessWsStreamClient {
	return &BusinessWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:         WS_BUSINESS,
			commonSubMap:    make(map[string]*Subscription[WsActionResult]),
			candleSubMap:    make(map[string]*Subscription[WsCandles]),
			booksSubMap:     make(map[string]*Subscription[WsBooks]),
			tradesSubMap:    make(map[string]*Subscription[WsTrades]),
			ordersSubMap:    make(map[string]*Subscription[WsOrders]),
			waitSubResult:   nil,
			waitSubResultMu: &sync.Mutex{},
		},
	}
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result WsActionResult) {
	if ws.connId == "" && result.ConnId != "" {
		ws.connId = result.ConnId
	}
	if result.Code != "" && result.Code != "0" {
		ws.waitSubResult.errChan <- fmt.Errorf("errHandler: %+v", result)
	} else {
		ws.waitSubResult.resultChan <- result
	}
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(args []WsSubscribeArg) {
	for _, arg := range args {
		data, _ := json.Marshal(&arg)
		key := string(data)
		if sub, ok := ws.candleSubMap[key]; ok {
			delete(ws.candleSubMap, key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.booksSubMap[key]; ok {
			delete(ws.booksSubMap, key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.tradesSubMap[key]; ok {
			delete(ws.tradesSubMap, key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
		if sub, ok := ws.ordersSubMap[key]; ok {
			delete(ws.ordersSubMap, key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
	}
}

func (ws *WsStreamClient) sendWsCloseToAllSub() {
	args := []WsSubscribeArg{}
	for key := range ws.commonSubMap {
		arg := WsSubscribeArg{}
		json.Unmarshal([]byte(key), &arg)
		args = append(args, arg)
	}
	ws.sendUnSubscribeSuccessToCloseChan(args)
}

func (ws *WsStreamClient) reSubscribeForReconnect() {
	isDoReSubscribe := map[int64]bool{}
	for _, sub := range ws.commonSubMap {
		if _, ok := isDoReSubscribe[sub.SubId]; ok {
			continue
		}

		reSub, err := subscribe[WsActionResult](ws, sub.Op, sub.Args)
		if err != nil {
			log.Error(err)
			ws.reSubscribeForReconnect()
			return
		}
		err = ws.CatchSubscribeReuslt(reSub)
		if err != nil {
			log.Error(err)
			ws.reSubscribeForReconnect()
			return
		}
		log.Debugf("reSubscribe Success: args:%v", reSub.Args)
		sub.SubId = reSub.SubId
		isDoReSubscribe[sub.SubId] = true
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
				//错误处理 重连等
				//ws标记为非关闭 且返回错误包含EOF、close、reset时自动重连
				if !ws.isClose && (strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "close") ||
					strings.Contains(err.Error(), "reset")) {
					//重连
					err := ws.OpenConn()
					for err != nil {
						time.Sleep(500 * time.Millisecond)
						err = ws.OpenConn()
					}
					go func() {
						//重新登陆
						if ws.lastLogin != nil && ws.client != nil {
							err = ws.Login(ws.client)
							for err != nil {
								time.Sleep(500 * time.Millisecond)
								err = ws.Login(ws.client)
							}
						}
						//重新订阅
						go func() {
							ws.reSubscribeForReconnect()
						}()
					}()

				} else {
					continue
				}
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

				//处理正常数据的返回结果
				//K线处理
				if strings.Contains(string(data), "candle") {
					c, err := handleWsCandle(data)

					arg := c.WsSubscribeArg
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.candleSubMap[string(keyData)]; ok {
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
					if sub, ok := ws.booksSubMap[string(keyData)]; ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *b
					}
					continue
				}
				if strings.Contains(string(data), "trades") {
					t, err := handleWsTrades(data)
					arg := t.WsSubscribeArg
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.tradesSubMap[string(keyData)]; ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *t
					}
					continue
				}
				if strings.Contains(string(data), "orders") {
					ordersList, err := handleWsOrders(data)
					if len(*ordersList) == 0 {
						continue
					}
					arg := (*ordersList)[0].WsSubscribeArg
					keyData, _ := json.Marshal(arg)
					if sub, ok := ws.ordersSubMap[string(keyData)]; ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						for _, order := range *ordersList {
							sub.resultChan <- order
						}
					}
					continue
				}

			}
		}
	}()
}

func (ws *WsStreamClient) DeferSub() {
	if len(ws.waitSubResult.subResultMap) == len(ws.waitSubResult.Args) {
		for _, arg := range ws.waitSubResult.Args {
			keyData, _ := json.Marshal(&arg)
			ws.commonSubMap[string(keyData)] = ws.waitSubResult
		}
		ws.waitSubResult = nil
		ws.waitSubResultMu.Unlock()
	}
}

// 取消订阅
func (sub *Subscription[T]) Unsubscribe() error {

	unSub, err := subscribe[WsActionResult](sub.Ws, UNSUBSCRIBE, sub.Args)
	if err != nil {
		return err
	}
	err = sub.Ws.CatchSubscribeReuslt(unSub)
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
		delete(sub.Ws.commonSubMap, key)
	}
	return nil
}

// 捕获订阅结果
func (ws *WsStreamClient) CatchLoginReuslt(sub *Subscription[WsActionResult]) error {
	ws.waitSubResult = sub
	defer sub.Ws.DeferSub()

	select {
	case err := <-sub.ErrChan():
		log.Error(err)
		return fmt.Errorf("Login Error: %v", err)
	case loginResult := <-sub.ResultChan():
		if loginResult.Code != "" && loginResult.Code != "0" {
			log.Error(loginResult.Code, ":", loginResult.Msg)
			return fmt.Errorf(loginResult.Code, ":", loginResult.Msg)
		}
		log.Debug("catchLoginResults: ", loginResult)
		return nil
	}

}

// 捕获订阅结果
func (ws *WsStreamClient) CatchSubscribeReuslt(sub *Subscription[WsActionResult]) error {
	ws.waitSubResult = sub
	defer sub.Ws.DeferSub()
	isBreak := false
	for {
		select {
		case err := <-sub.ErrChan():
			log.Error(err)
			return fmt.Errorf("Unsubscribe Error: %v", err)
		case subResult := <-sub.ResultChan():
			if subResult.Code != "" && subResult.Code != "0" {
				log.Error(subResult.Code, ":", subResult.Msg)
				return fmt.Errorf(subResult.Code, ":", subResult.Msg)
			}
			keyData, _ := json.Marshal(subResult.Arg)
			sub.subResultMap[string(keyData)] = subResult
			if len(sub.subResultMap) == len(sub.Args) {
				isBreak = true
			}
		}
		if isBreak {
			break
		}
	}
	log.Debug("catchResults: ", sub.subResultMap)
	return nil
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
	log.Debugf("Login Success args:%v ", ws.lastLogin.Args)
	return nil
}

// 标准订阅方法
func wsStreamServe(api string, resultChan chan []byte, errChan chan error) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(655350)
	go func() {
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- message
		}
	}()
	return c, err
}

// 获取数据流请求URL
func handlerWsStreamRequestApi(apiType APIType) string {
	u := url.URL{
		Scheme:   "wss",
		Host:     OKX_API_WS_HOST,
		Path:     getWsApi(apiType),
		RawQuery: "",
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

// 发送ping/pong消息以检查连接稳定性
func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}
