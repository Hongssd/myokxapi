package myokxapi

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
	err = ws.CatchSubscribeReuslt(doSub)
	if err != nil {
		return nil, err
	}
	log.Debugf("SubscribeOrders Success: args:%v", doSub.Args)
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
		ws.ordersSubMap[string(keyData)] = sub
	}
	return sub, nil
}
