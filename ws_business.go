package myokxapi

func getCandleSubscribeArg(instId string, interval string) WsSubscribeArg {
	return WsSubscribeArg{
		Channel: "candle" + interval,
		InstId:  instId,
	}
}

// 订阅K线 如: "BTCUSDT","1m"
func (ws *BusinessWsStreamClient) SubscribeCandle(instId string, interval string) (*Subscription[WsCandles], error) {
	return ws.SubscribeCandleMultiple([]string{instId}, []string{interval})
}

// 订阅K线 如: ["BTC-USDT","ETH-USDT"],["1m","5m"]
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
	err = ws.CatchSubscribeReuslt(doSub)
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
		ws.candleSubMap[string(keyData)] = sub
	}
	return sub, nil
}
