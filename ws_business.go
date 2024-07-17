package myokxapi

func getCandleSubscribeArg(instId string, interval string) WsSubscribeArg {
	return WsSubscribeArg{
		Channel: "candle" + interval,
		InstId:  instId,
	}
}

func getOrdersAlgoSubscribeArg(instType, instFamily, instId string) WsSubscribeArg {
	return WsSubscribeArg{
		Channel:    "orders-algo",
		InstType:   instType,
		InstFamily: instFamily,
		InstId:     instId,
	}
}

// 订阅策略订单频道
// > instType	String	是	产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权 ANY：全部
// > instFamily	String	否	交易品种 适用于交割/永续/期权
// > instId	    String	否	产品ID
func (ws *BusinessWsStreamClient) SubscribeOrdersAlgo(instType, instFamily, instId string) (*Subscription[WsOrdersAlgo], error) {
	arg := getOrdersAlgoSubscribeArg(instType, instFamily, instId)
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
	sub := &Subscription[WsOrdersAlgo]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsOrdersAlgo, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.ordersAlgoSubMap.Store(string(keyData), sub)
	}
	return sub, nil
}

// 取消订阅策略订单频道
func (ws *BusinessWsStreamClient) UnSubscribeOrdersAlgo(instType, instFamily, instId string) error {
	arg := getOrdersAlgoSubscribeArg(instType, instFamily, instId)
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
		ws.ordersAlgoSubMap.Delete(string(keyData))
	}
	return nil
}

// 订阅单个K线 如: "BTCUSDT","1m"
func (ws *BusinessWsStreamClient) SubscribeCandle(instId string, interval string) (*Subscription[WsCandles], error) {
	return ws.SubscribeCandleMultiple([]string{instId}, []string{interval})
}

// 批量订阅K线 如: ["BTC-USDT","ETH-USDT"],["1m","5m"]
func (ws *BusinessWsStreamClient) SubscribeCandleMultiple(instId []string, interval []string) (*Subscription[WsCandles], error) {
	args := []WsSubscribeArg{}
	for _, s := range instId {
		for _, i := range interval {
			arg := getCandleSubscribeArg(s, i)
			args = append(args, arg)
		}
	}
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return nil, err
	}
	log.Infof("SubscribeCandle Success: args:%v", doSub.Args)
	sub := &Subscription[WsCandles]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsCandles, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.candleSubMap.Store(string(keyData), sub)
	}
	return sub, nil
}

// 取消订阅单个K线 如: "BTCUSDT","1m"
func (ws *BusinessWsStreamClient) UnSubscribeCandle(instId string, interval string) error {
	return ws.UnSubscribeCandleMultiple([]string{instId}, []string{interval})
}

// 批量取消订阅K线 如: ["BTC-USDT","ETH-USDT"],["1m","5m"]
func (ws *BusinessWsStreamClient) UnSubscribeCandleMultiple(instId []string, interval []string) error {
	args := []WsSubscribeArg{}
	for _, s := range instId {
		for _, i := range interval {
			arg := getCandleSubscribeArg(s, i)
			args = append(args, arg)
		}
	}
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, UNSUBSCRIBE, args)
	if err != nil {
		return err
	}
	err = ws.catchSubscribeResult(doSub)
	if err != nil {
		return err
	}
	log.Infof("UnSubscribeCandle Success: args:%v", doSub.Args)

	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.candleSubMap.Delete(string(keyData))
	}
	return nil
}
