package myokxapi

type WsBooksType string

const (
	WS_BOOKS_SNAPSHOT_1_10MS  WsBooksType = "bbo-tbt"        //首次推1档快照数据，以后定量推送，每10毫秒当1档快照数据有变化推送一次1档数据
	WS_BOOKS_SNAPSHOT_5_100MS WsBooksType = "books5"         //首次推5档快照数据，以后定量推送，每100毫秒当5档快照数据有变化推送一次5档数据
	WS_BOOKS_UPDATE_400_100MS WsBooksType = "books"          //首次推400档快照数据，以后增量推送，每100毫秒推送一次变化的数据
	WS_BOOKS_UPDATE_400_10MS  WsBooksType = "books-l2-tbt"   //首次推400档快照数据，以后增量推送，每10毫秒推送一次变化的数据
	WS_BOOKS_UPDATE_50_10MS   WsBooksType = "books50-l2-tbt" //首次推50档快照数据，以后增量推送，每10毫秒推送一次变化的数据
)

func (wsBooksType WsBooksType) String() string {
	return string(wsBooksType)
}

// 批量订阅深度 如: ["BTC-USDT","ETH-USDT"], WS_BOOKS_SNAPSHOT_5_100MS
func (ws *PublicWsStreamClient) SubscribeBooksMultiple(instIds []string, wsBooksType WsBooksType) (*Subscription[WsBooks], error) {
	args := []WsSubscribeArg{}
	for _, s := range instIds {
		arg := getBooksSubscribeArg(s, wsBooksType)
		args = append(args, arg)
	}
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.CatchSubscribeReuslt(doSub)
	if err != nil {
		return nil, err
	}
	log.Debugf("SubscribeBooks Success: args:%v", doSub.Args)
	sub := &Subscription[WsBooks]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsBooks, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.booksSubMap[string(keyData)] = sub
	}
	return sub, nil
}

// 订阅深度 如: "BTC-USDT", WS_BOOKS_SNAPSHOT_5_100MS
func (ws *PublicWsStreamClient) SubscribeBooks(instIds string, wsBooksType WsBooksType) (*Subscription[WsBooks], error) {
	return ws.SubscribeBooksMultiple([]string{instIds}, wsBooksType)
}

func getBooksSubscribeArg(instId string, wsBooksType WsBooksType) WsSubscribeArg {
	return WsSubscribeArg{
		Channel: wsBooksType.String(),
		InstId:  instId,
	}
}

// 批量订阅交易 如: ["BTC-USDT","ETH-USDT"]
func (ws *PublicWsStreamClient) SubscribeTradesMultiple(instIds []string) (*Subscription[WsTrades], error) {
	args := []WsSubscribeArg{}
	for _, s := range instIds {
		arg := getTradesSubscribeArg(s)
		args = append(args, arg)
	}
	doSub, err := subscribe[WsActionResult](&ws.WsStreamClient, SUBSCRIBE, args)
	if err != nil {
		return nil, err
	}
	err = ws.CatchSubscribeReuslt(doSub)
	if err != nil {
		return nil, err
	}
	log.Debugf("SubscribeTrades Success: args:%v", doSub.Args)
	sub := &Subscription[WsTrades]{
		SubId:      doSub.SubId,
		Op:         SUBSCRIBE,
		Args:       doSub.Args,
		resultChan: make(chan WsTrades, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
		Ws:         &ws.WsStreamClient,
	}
	for _, arg := range args {
		keyData, _ := json.Marshal(&arg)
		ws.tradesSubMap[string(keyData)] = sub
	}
	return sub, nil
}

func getTradesSubscribeArg(instId string) WsSubscribeArg {
	return WsSubscribeArg{
		Channel: "trades",
		InstId:  instId,
	}
}
