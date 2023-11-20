package myokxapi

func getCandleSubscribeArg(instId string, interval string) SubscribeArg {
	return SubscribeArg{
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
	args := []SubscribeArg{}
	for _, s := range instId {
		for _, i := range interval {
			arg := getCandleSubscribeArg(s, i)
			args = append(args, arg)
		}
	}
	doSub, err := sendMsg[SubscribeResult](&ws.WsStreamClient, SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.CatchSubscribeReuslt(doSub)
	if err != nil {
		return nil, err
	}
	log.Debugf("SubscribeCandle Success: args:%v", doSub.Args)
	sub := &Subscription[WsCandles]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsCandles, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.candleSubMap[string(keyData)] = sub
	}
	return sub, nil
}
