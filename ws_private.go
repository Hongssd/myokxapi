package myokxapi

import "github.com/bwmarrin/snowflake"

func getOrdersSubscribeArg(instType, instFamily, instId string) WsSubscribeArg {
	return WsSubscribeArg{
		Channel:    "orders",
		InstType:   instType,
		InstFamily: instFamily,
		InstId:     instId,
	}
}

// 订阅订单频道
// > instType	String	是	产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权 ANY：全部
// > instFamily	String	否	交易品种 适用于交割/永续/期权
// > instId	    String	否	产品ID
func (ws *PrivateWsStreamClient) SubscribeOrders(instType, instFamily, instId string) (*Subscription[WsOrders], error) {
	arg := getOrdersSubscribeArg(instType, instFamily, instId)
	args := []WsSubscribeArg{arg}
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}
	log.Infof("SubscribeOrders Success: args:%v", doSub.Args)
	sub := &Subscription[WsOrders]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsOrders, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.ordersSubMap.Store(string(keyData), sub)
	}
	return sub, nil
}

// 取消订阅订单频道
// > instType	String	是	产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权 ANY：全部
// > instFamily	String	否	交易品种 适用于交割/永续/期权
// > instId	    String	否	产品ID
func (ws *PrivateWsStreamClient) UnSubscribeOrders(instType, instFamily, instId string) error {
	arg := getOrdersSubscribeArg(instType, instFamily, instId)
	args := []WsSubscribeArg{arg}
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}
	log.Infof("UnSubscribeOrders Success: args:%v", doSub.Args)

	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.ordersSubMap.Delete(string(keyData))
	}
	return nil
}

type WsOrderArgData struct {
	InstId       string `json:"instId"`                 //String	是	产品ID，如 BTC-USD-190927-5000-C
	TdMode       string `json:"tdMode"`                 //String	是	交易模式 保证金模式 isolated：逐仓 cross： 全仓 非保证金模式 cash：现金 spot_isolated：现货逐仓(仅适用于现货带单)
	Ccy          string `json:"ccy,omitempty"`          //String	否	保证金币种，仅适用于单币种保证金账户下的全仓杠杆订单
	ClOrdId      string `json:"clOrdId,omitempty"`      //String	否	由用户设置的订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag          string `json:"tag,omitempty"`          //String	否	订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-16位之间。
	Side         string `json:"side"`                   //String	是	订单方向，buy sell
	PosSide      string `json:"posSide,omitempty"`      //String	否	持仓方向 在买卖模式下，默认 net 在开平仓模式下必填，且仅可选择 long 或 short，仅适用于交割/永续
	OrdType      string `json:"ordType"`                //String	是	订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单)
	Sz           string `json:"sz"`                     //String	是	委托数量
	Px           string `json:"px,omitempty"`           //String	可选	委托价格，仅适用于limit、post_only、fok、ioc、mmp、mmp_and_post_only类型的订单 期权下单时，px/pxUsd/pxVol 只能填一个
	PxUsd        string `json:"pxUsd,omitempty"`        //String	可选	以USD价格进行期权下单 仅适用于期权 期权下单时 px/pxUsd/pxVol 必填一个，且只能填一个
	PxVol        string `json:"pxVol,omitempty"`        //String	可选	以隐含波动率进行期权下单，例如 1 代表 100% 仅适用于期权 期权下单时 px/pxUsd/pxVol 必填一个，且只能填一个
	ReduceOnly   bool   `json:"reduceOnly,omitempty"`   //Boolean	否	是否只减仓，true 或 false，默认false 仅适用于币币杠杆，以及买卖模式下的交割/永续 仅适用于单币种保证金模式和跨币种保证金模式
	TgtCcy       string `json:"tgtCcy,omitempty"`       //String	否	币币市价单委托数量sz的单位 base_ccy: 交易货币 ；quote_ccy：计价货币 仅适用于币币市价订单 默认买单为quote_ccy，卖单为base_ccy
	BanAmend     bool   `json:"banAmend,omitempty"`     //Boolean	否	是否禁止币币市价改单，true 或 false，默认false 为true时，余额不足时，系统不会改单，下单会失败，仅适用于币币市价单
	QuickMgnType string `json:"quickMgnType,omitempty"` //String	否	一键借币类型，仅适用于杠杆逐仓的一键借币模式： manual：手动，auto_borrow： 自动借币，auto_repay： 自动还币 默认是manual：手动
	StpId        string `json:"stpId,omitempty"`        //String	否	自成交保护ID。来自同一个母账户配着同一个ID的订单不能自成交 用户自定义1<=x<=999999999的整数
	StpMode      string `json:"stpMode,omitempty"`      //String	否	自成交保护模式，需要stpId有值才会生效 默认为 cancel maker cancel_maker,cancel_taker, cancel_both
}

type WsCancelOrderArgData struct {
	InstId  string `json:"instId"`  //String	是	产品ID
	OrdId   string `json:"ordId"`   //String	可选	订单ID ordId和clOrdId必须传一个，若传两个，以 ordId 为主
	ClOrdId string `json:"clOrdId"` //String	可选	用户提供的订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度要在1-32位之间
}

type WsAmendOrderArgData struct {
	InstId    string `json:"instId"`    //String	是	产品ID
	CxlOnFail bool   `json:"cxlOnFail"` //Boolean	否	当订单修改失败时，该订单是否需要自动撤销 false：不自动撤单 true：自动撤单 ，默认为 false
	OrdId     string `json:"ordId"`     //String	可选	订单ID ordId和clOrdId必须传一个，若传两个，以 ordId 为主
	ClOrdId   string `json:"clOrdId"`   //String	可选	用户提供的订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	ReqId     string `json:"reqId"`     //String	否	用户提供的reqId 如果提供，那在返回参数中返回reqId，方便找到相应的修改请求。 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	NewSz     string `json:"newSz"`     //String	可选	请求修改的新数量，newSz和newPx不可同时为空。对于部分成交订单，该数量应包含已成交数量。
	NewPx     string `json:"newPx"`     //String	可选	修改后的新价格 修改的新价格期权改单时，newPx/newPxUsd/newPxVol 只能填一个，且必须与下单参数保持一致，如下单用px，改单时需使用newPx
	NewPxUsd  string `json:"newPxUsd"`  //String	可选	以USD价格进行期权改单 仅适用于期权，期权改单时，newPx/newPxUsd/newPxVol 只能填一个
	NewPxVol  string `json:"newPxVol"`  //String	可选	以隐含波动率进行期权改单，例如 1 代表 100% 仅适用于期权，期权改单时，newPx/newPxUsd/newPxVol 只能填一个
}

// 下单
func (ws *PrivateWsStreamClient) Order(orderArgData WsOrderArgData, expTimestamp int64) (*WsOrderResult, error) {
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	reqId := node.Generate().String()
	return orderAction[WsOrderArgData](ws, ORDER, reqId, []WsOrderArgData{orderArgData}, expTimestamp)
}

// 批量下单
func (ws *PrivateWsStreamClient) BatchOrder(orderArgDataList []WsOrderArgData, expTimestamp int64) (*WsOrderResult, error) {
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	reqId := node.Generate().String()
	return orderAction(ws, BATCH_ORDER, reqId, orderArgDataList, expTimestamp)
}

// 撤单
func (ws *PrivateWsStreamClient) CancelOrder(cancelOrderArgData WsCancelOrderArgData, expTimestamp int64) (*WsOrderResult, error) {
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	reqId := node.Generate().String()
	return orderAction(ws, CANCEL_ORDER, reqId, []WsCancelOrderArgData{cancelOrderArgData}, expTimestamp)
}

// 批量撤单
func (ws *PrivateWsStreamClient) BatchCancelOrder(cancelOrderArgDataList []WsCancelOrderArgData, expTimestamp int64) (*WsOrderResult, error) {
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	reqId := node.Generate().String()
	return orderAction(ws, BATCH_CANCEL_ORDER, reqId, cancelOrderArgDataList, expTimestamp)
}

// 改单
func (ws *PrivateWsStreamClient) AmendOrder(amendOrderArgData WsAmendOrderArgData, expTimestamp int64) (*WsOrderResult, error) {
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	reqId := node.Generate().String()
	return orderAction(ws, AMEND_ORDER, reqId, []WsAmendOrderArgData{amendOrderArgData}, expTimestamp)
}

// 批量改单
func (ws *PrivateWsStreamClient) BatchAmendOrder(amendOrderArgDataList []WsAmendOrderArgData, expTimestamp int64) (*WsOrderResult, error) {
	node, err := snowflake.NewNode(2)
	if err != nil {
		return nil, err
	}
	reqId := node.Generate().String()
	return orderAction(ws, BATCH_AMEND_ORDER, reqId, amendOrderArgDataList, expTimestamp)
}
