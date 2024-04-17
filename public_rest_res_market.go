package myokxapi

type PublicRestMarketTickersRes []PublicRestMarketTickersResRow
type PublicRestMarketTickersResRow struct {
	InstId  string `json:"instId"`  //指数
	IdxPx   string `json:"idxPx"`   //最新指数价格
	High24h string `json:"high24h"` //24小时指数最高价格
	SodUtc0 string `json:"sodUtc0"` //UTC 0 时开盘价
	Open24h string `json:"open24h"` //24小时指数开盘价格
	Low24h  string `json:"low24h"`  //24小时指数最低价格
	SodUtc8 string `json:"sodUtc8"` //UTC+8 时开盘价
	Ts      string `json:"ts"`      //指数价格更新时间，Unix时间戳的毫秒数格式，如1597026383085
}

type PublicRestMarketBooksRes []PublicRestMarketBooksResRow
type PublicRestMarketBooksResRow struct {
	Asks []Books `json:"asks"` //卖方深度
	Bids []Books `json:"bids"` //买方深度
	Ts   string  `json:"ts"`   //深度产生的时间
}
type Books struct {
	Price      string `json:"price"`       //价格
	Quantity   string `json:"quantity"`    //合约张数或交易币的数量
	OrderCount string `json:"order_count"` //订单数量
}

type PublicRestMarketBooksMiddle []PublicRestMarketBooksMiddleRow
type PublicRestMarketBooksMiddleRow struct {
	Asks []interface{} `json:"asks"` //卖方深度
	Bids []interface{} `json:"bids"` //买方深度
	Ts   string        `json:"ts"`   //深度产生的时间
}

func (middle *PublicRestMarketBooksMiddle) ConvertToRes() *PublicRestMarketBooksRes {
	resList := PublicRestMarketBooksRes{}
	for _, v := range *middle {
		resList = append(resList, *v.ConvertToRes())
	}
	return &resList
}

func (middleRow *PublicRestMarketBooksMiddleRow) ConvertToRes() *PublicRestMarketBooksResRow {
	res := PublicRestMarketBooksResRow{
		Ts: middleRow.Ts,
	}
	res.Bids = []Books{}
	res.Asks = []Books{}
	for _, bid := range middleRow.Bids {
		res.Bids = append(res.Bids, Books{
			Price:      bid.([]interface{})[0].(string),
			Quantity:   bid.([]interface{})[1].(string),
			OrderCount: bid.([]interface{})[3].(string),
		})
	}
	for _, ask := range middleRow.Asks {
		res.Asks = append(res.Asks, Books{
			Price:      ask.([]interface{})[0].(string),
			Quantity:   ask.([]interface{})[1].(string),
			OrderCount: ask.([]interface{})[3].(string),
		})
	}
	return &res
}

type PublicRestMarketCandlesRes []PublicRestMarketCandlesResRow
type PublicRestMarketCandlesResRow struct {
	Ts          string `json:"ts"`          //开始时间，Unix时间戳的毫秒数格式，如 1597026383085
	O           string `json:"o"`           //开盘价格
	H           string `json:"h"`           //最高价格
	L           string `json:"l"`           //最低价格
	C           string `json:"c"`           //收盘价格
	Vol         string `json:"vol"`         //交易量，以张为单位
	VolCcy      string `json:"volCcy"`      //交易量，以币为单位
	VolCcyQuote string `json:"volCcyQuote"` //交易量，以计价货币为单位
	Confirm     string `json:"confirm"`     //K线状态 0 代表 K 线未完结，1 代表 K 线已完结。
}
type PublicRestMarketCandlesMiddle []PublicRestMarketCandlesMiddleRow
type PublicRestMarketCandlesMiddleRow [9]interface{}

func (middle *PublicRestMarketCandlesMiddle) ConvertToRes() *PublicRestMarketCandlesRes {
	resList := PublicRestMarketCandlesRes{}
	for _, v := range *middle {
		resList = append(resList, *v.ConvertToRes())
	}
	return &resList
}

func (middleRow *PublicRestMarketCandlesMiddleRow) ConvertToRes() *PublicRestMarketCandlesResRow {
	res := PublicRestMarketCandlesResRow{
		Ts:          middleRow[0].(string),
		O:           middleRow[1].(string),
		H:           middleRow[2].(string),
		L:           middleRow[3].(string),
		C:           middleRow[4].(string),
		Vol:         middleRow[5].(string),
		VolCcy:      middleRow[6].(string),
		VolCcyQuote: middleRow[7].(string),
		Confirm:     middleRow[8].(string),
	}
	return &res
}
