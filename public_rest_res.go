package myokxapi

type PublicRestInstrumentsRes []PublicRestInstrumentsResRow
type PublicRestInstrumentsResRow struct {
	InstType     string `json:"instType"`     //产品类型  SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId       string `json:"instId"`       //产品id， 如 BTC-USD-SWAP
	Uly          string `json:"uly"`          //标的指数，如 BTC-USD，仅适用于交割/永续/期权
	InstFamily   string `json:"instFamily"`   //交易品种，如 BTC-USD，仅适用于交割/永续/期权
	Category     string `json:"category"`     //币种类别（已废弃）
	BaseCcy      string `json:"baseCcy"`      //交易货币币种，如 BTC-USDT 中的 BTC ，仅适用于币币/币币杠杆
	QuoteCcy     string `json:"quoteCcy"`     //计价货币币种，如 BTC-USDT 中的USDT ，仅适用于币币/币币杠杆
	SettleCcy    string `json:"settleCcy"`    //盈亏结算和保证金币种，如 BTC 仅适用于交割/永续/期权
	CtVal        string `json:"ctVal"`        //合约面值，仅适用于交割/永续/期权
	CtMult       string `json:"ctMult"`       //合约乘数，仅适用于交割/永续/期权
	CtValCcy     string `json:"ctValCcy"`     //合约面值计价币种，仅适用于交割/永续/期权
	OptType      string `json:"optType"`      //期权类型，C或P 仅适用于期权
	Stk          string `json:"stk"`          //行权价格，仅适用于期权
	ListTime     string `json:"listTime"`     //上线时间  Unix时间戳的毫秒数格式，如 1597026383085
	ExpTime      string `json:"expTime"`      //产品下线时间  适用于币币/杠杆/交割/永续/期权，对于 交割/期权，为交割/行权日期；亦可以为产品下线时间，有变动就会推送。
	Lever        string `json:"lever"`        //该instId支持的最大杠杆倍数，不适用于币币、期权
	TickSz       string `json:"tickSz"`       //下单价格精度，如 0.0001 对于期权来说，是梯度中的最小下单价格精度，如果想要获取期权价格梯度，请使用"获取期权价格梯度"接口
	LotSz        string `json:"lotSz"`        //下单数量精度，如 BTC-USDT-SWAP：1
	MinSz        string `json:"minSz"`        //最小下单数量, 合约的数量单位是“张”，现货的数量单位是“交易货币”
	CtType       string `json:"ctType"`       //linear：正向合约 inverse：反向合约 仅适用于交割/永续
	Alias        string `json:"alias"`        //合约日期别名 this_week：本周 next_week：次周 quarter：季度 next_quarter：次季度 仅适用于交割
	State        string `json:"state"`        //产品状态 live：交易中 suspend：暂停中 preopen：预上线，如：交割和期权的新合约在 live 之前，会有 preopen 状态 test：测试中（测试产品，不可交易）
	MaxLmtSz     string `json:"maxLmtSz"`     //合约或现货限价单的单笔最大委托数量, 合约的数量单位是“张”，现货的数量单位是“交易货币”
	MaxMktSz     string `json:"maxMktSz"`     //合约或现货市价单的单笔最大委托数量, 合约的数量单位是“张”，现货的数量单位是“USDT”
	MaxTwapSz    string `json:"maxTwapSz"`    //合约或现货时间加权单的单笔最大委托数量, 合约的数量单位是“张”，现货的数量单位是“交易货币”
	MaxIcebergSz string `json:"maxIcebergSz"` //合约或现货冰山委托的单笔最大委托数量, 合约的数量单位是“张”，现货的数量单位是“交易货币”
	MaxTriggerSz string `json:"maxTriggerSz"` //合约或现货计划委托委托的单笔最大委托数量, 合约的数量单位是“张”，现货的数量单位是“交易货币”
	MaxStopSz    string `json:"maxStopSz"`    //合约或现货止盈止损市价委托的单笔最大委托数量, 合约的数量单位是“张”，现货的数量单位是“USDT”
}

type PublicRestTimeRes []PublicRestTimeResRow
type PublicRestTimeResRow struct {
	Ts string `json:"ts"` //系统时间，Unix时间戳的毫秒数格式，如 1597026383085
}

type PublicRestMarkPriceRes []PublicRestMarkPriceResRow
type PublicRestMarkPriceResRow struct {
	InstType string `json:"instType"` //产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId   string `json:"instId"`   //产品ID，如 BTC-USD-200214
	MarkPx   string `json:"markPx"`   //标记价格
	Ts       string `json:"ts"`       //接口数据返回时间，Unix时间戳的毫秒数格式，如1597026383085
}

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

// asks	Array	卖方深度
// bids	Array	买方深度
// ts	String	深度产生的时间
// 合约的asks和bids值数组举例说明： ["411.8","10", "0","4"] 411.8为深度价格，10为此价格的合约张数，0该字段已弃用(始终为0)，4为此价格的订单数量
// 现货/币币杠杆的asks和bids值数组举例说明： ["411.8","10", "0","4"] 411.8为深度价格，10为此价格的交易币的数量，0该字段已弃用(始终为0)，4为此价格的订单数量 asks和bids值数组举例说明： ["411.8", "10", "0", "4"]
// - 411.8为深度价格
// - 10为此价格的数量 （合约交易为张数，现货/币币杠杆为交易币的数量）
// - 0该字段已弃用(始终为0)
// - 4为此价格的订单数量

type BooksLite struct {
	Price      string `json:"price"`       //价格
	Quantity   string `json:"quantity"`    //合约张数或交易币的数量
	OrderCount string `json:"order_count"` //订单数量
}

type PublicRestMarketBooksLiteRes []PublicRestMarketBooksLiteResRow
type PublicRestMarketBooksLiteResRow struct {
	Asks []BooksLite `json:"asks"` //卖方深度
	Bids []BooksLite `json:"bids"` //买方深度
	Ts   string      `json:"ts"`   //深度产生的时间
}
type PublicRestMarketBooksLiteMiddle []PublicRestMarketBooksLiteMiddleRow
type PublicRestMarketBooksLiteMiddleRow struct {
	Asks []interface{} `json:"asks"` //卖方深度
	Bids []interface{} `json:"bids"` //买方深度
	Ts   string        `json:"ts"`   //深度产生的时间
}

func (middle *PublicRestMarketBooksLiteMiddle) ConvertToRes() *PublicRestMarketBooksLiteRes {
	resList := PublicRestMarketBooksLiteRes{}
	for _, v := range *middle {
		res := PublicRestMarketBooksLiteResRow{
			Ts: v.Ts,
		}
		res.Bids = []BooksLite{}
		res.Asks = []BooksLite{}
		for _, bid := range v.Bids {
			res.Bids = append(res.Bids, BooksLite{
				Price:      bid.([]interface{})[0].(string),
				Quantity:   bid.([]interface{})[1].(string),
				OrderCount: bid.([]interface{})[3].(string),
			})
		}
		for _, ask := range v.Asks {
			res.Asks = append(res.Asks, BooksLite{
				Price:      ask.([]interface{})[0].(string),
				Quantity:   ask.([]interface{})[1].(string),
				OrderCount: ask.([]interface{})[3].(string),
			})
		}
		resList = append(resList, res)
	}
	return &resList
}

// ts	String	开始时间，Unix时间戳的毫秒数格式，如 1597026383085
// o	String	开盘价格
// h	String	最高价格
// l	String	最低价格
// c	String	收盘价格
// vol	String	交易量，以张为单位
// 如果是衍生品合约，数值为合约的张数。
// 如果是币币/币币杠杆，数值为交易货币的数量。
// volCcy	String	交易量，以币为单位
// 如果是衍生品合约，数值为交易货币的数量。
// 如果是币币/币币杠杆，数值为计价货币的数量。
// volCcyQuote	String	交易量，以计价货币为单位
// 如：BTC-USDT 和 BTC-USDT-SWAP, 单位均是 USDT；
// BTC-USD-SWAP 单位是 USD
// confirm	String	K线状态
// 0 代表 K 线未完结，1 代表 K 线已完结。

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
		res := PublicRestMarketCandlesResRow{
			Ts:          v[0].(string),
			O:           v[1].(string),
			H:           v[2].(string),
			L:           v[3].(string),
			C:           v[4].(string),
			Vol:         v[5].(string),
			VolCcy:      v[6].(string),
			VolCcyQuote: v[7].(string),
			Confirm:     v[8].(string),
		}
		resList = append(resList, res)
	}
	return &resList
}
