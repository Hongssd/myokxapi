package myokxapi

type WsCandles struct {
	WsSubscribeArg        //订阅信息
	Interval       string `json:"interval"`    //K线周期，如 1m，5m，15m，30m，1h，2h，4h，6h，12h，1d，1w，1M
	Ts             string `json:"ts"`          //开始时间，Unix时间戳的毫秒数格式，如 1597026383085
	O              string `json:"o"`           //开盘价格
	H              string `json:"h"`           //最高价格
	L              string `json:"l"`           //最低价格
	C              string `json:"c"`           //收盘价格
	Vol            string `json:"vol"`         //交易量，以张为单位
	VolCcy         string `json:"volCcy"`      //交易量，以币为单位
	VolCcyQuote    string `json:"volCcyQuote"` //交易量，以计价货币为单位
	Confirm        string `json:"confirm"`     //K线状态 0 代表 K 线未完结，1 代表 K 线已完结。
}

type WsCandlesMiddle struct {
	Arg  WsSubscribeArg   `json:"arg"`  //订阅信息
	Data [][9]interface{} `json:"data"` //K线数据
}

type WsBooks struct {
	WsSubscribeArg             //订阅信息
	Action         string      `json:"action"`    //推送数据动作，增量推送数据还是全量推送数据 snapshot：全量 update：增量
	Asks           []BooksLite `json:"asks"`      //卖方深度
	Bids           []BooksLite `json:"bids"`      //买方深度
	Ts             string      `json:"ts"`        //深度产生的时间
	CheckSum       int64       `json:"checksum"`  //检验和
	PrevSeqId      int64       `json:"prevSeqId"` //上一个推送的序列号。仅适用 books，books-l2-tbt，books50-l2-tbt
	SeqId          int64       `json:"seqId"`     //推送的序列号
}

type WsBooksMiddle struct {
	Arg    WsSubscribeArg `json:"arg"`    //订阅信息
	Action string         `json:"action"` //推送数据动作，增量推送数据还是全量推送数据 snapshot：全量 update：增量
	Data   []struct {
		Asks      [][4]interface{} `json:"asks"`      //卖方深度
		Bids      [][4]interface{} `json:"bids"`      //买方深度
		Ts        string           `json:"ts"`        //深度产生的时间
		CheckSum  int64            `json:"checksum"`  //检验和
		PrevSeqId int64            `json:"prevSeqId"` //上一个推送的序列号。仅适用 books，books-l2-tbt，books50-l2-tbt
		SeqId     int64            `json:"seqId"`     //推送的序列号
	} `json:"data"` //深度数据
}

// arg	Object	订阅成功的频道
// > channel	String	频道名
// > instId	String	产品ID
// data	Array	订阅的数据
// > instId	String	产品ID，如 BTC-USD-180216
// > tradeId	String	聚合的多笔交易中最新一笔交易的成交ID
// > px	String	成交价格
// > sz	String	成交数量
// > side	String	成交方向，buy sell
// > ts	String	成交时间，Unix时间戳的毫秒数格式，如 1597026383085
// > count	String	聚合的订单匹配数量
type WsTrades struct {
	WsSubscribeArg //订阅信息
	Trades
}

type Trades struct {
	InstId  string `json:"instId"`  //产品ID，如 BTC-USD-180216
	TradeId string `json:"tradeId"` //聚合的多笔交易中最新一笔交易的成交ID
	Px      string `json:"px"`      //成交价格
	Sz      string `json:"sz"`      //成交数量
	Side    string `json:"side"`    //成交方向，buy sell
	Ts      string `json:"ts"`      //成交时间，Unix时间戳的毫秒数格式，如 1597026383085
	Count   string `json:"count"`   //聚合的订单匹配数量
}

type WsTradesMiddle struct {
	Arg  WsSubscribeArg `json:"arg"`
	Data []Trades       `json:"data"`
}

func handleWsCandle(data []byte) (*WsCandles, error) {

	wsCandlesMiddle := WsCandlesMiddle{}
	err := json.Unmarshal(data, &wsCandlesMiddle)
	if err != nil {
		return nil, err
	}
	candelData := wsCandlesMiddle.Data[0]
	candle := WsCandles{
		WsSubscribeArg: wsCandlesMiddle.Arg,
		Interval:       wsCandlesMiddle.Arg.Channel[6:],
		Ts:             candelData[0].(string),
		O:              candelData[1].(string),
		H:              candelData[2].(string),
		L:              candelData[3].(string),
		C:              candelData[4].(string),
		Vol:            candelData[5].(string),
		VolCcy:         candelData[6].(string),
		VolCcyQuote:    candelData[7].(string),
		Confirm:        candelData[8].(string),
	}
	return &candle, nil
}

func handleWsBooks(data []byte) (*WsBooks, error) {

	wsBooksMiddle := WsBooksMiddle{}
	err := json.Unmarshal(data, &wsBooksMiddle)
	if err != nil {
		return nil, err
	}
	middleRow := wsBooksMiddle.Data[0]

	wsBook := WsBooks{
		WsSubscribeArg: wsBooksMiddle.Arg,
		Action:         wsBooksMiddle.Action,
		Asks:           []BooksLite{},
		Bids:           []BooksLite{},
		Ts:             middleRow.Ts,
		CheckSum:       middleRow.CheckSum,
		PrevSeqId:      middleRow.PrevSeqId,
		SeqId:          middleRow.SeqId,
	}

	for _, ask := range middleRow.Asks {
		wsBook.Asks = append(wsBook.Asks, BooksLite{
			Price:      ask[0].(string),
			Quantity:   ask[1].(string),
			OrderCount: ask[3].(string),
		})
	}
	for _, bid := range middleRow.Bids {
		wsBook.Bids = append(wsBook.Bids, BooksLite{
			Price:      bid[0].(string),
			Quantity:   bid[1].(string),
			OrderCount: bid[3].(string),
		})
	}

	return &wsBook, nil
}

func handleWsTrades(data []byte) (*WsTrades, error) {

	wsTradesMiddle := WsTradesMiddle{}
	err := json.Unmarshal(data, &wsTradesMiddle)
	if err != nil {
		return nil, err
	}
	trades := wsTradesMiddle.Data[0]
	wsTrades := WsTrades{
		WsSubscribeArg: wsTradesMiddle.Arg,
		Trades: Trades{
			InstId:  trades.InstId,
			TradeId: trades.TradeId,
			Px:      trades.Px,
			Sz:      trades.Sz,
			Side:    trades.Side,
			Ts:      trades.Ts,
			Count:   trades.Count,
		},
	}
	return &wsTrades, nil
}
