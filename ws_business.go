package myokxapi

func getCandleSubscribeArg(instId string, interval string) WsSubscribeArg {
	return WsSubscribeArg{
		Channel: "candle" + interval,
		InstId:  instId,
	}
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
	err = ws.CatchSubscribeReuslt(doSub)
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
