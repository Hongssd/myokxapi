package myokxapi

type PublicRestPublicInstrumentsRes []PublicRestPublicInstrumentsResRow
type PublicRestPublicInstrumentsResRow struct {
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

type PublicRestPublicTimeRes []PublicRestPublicTimeResRow
type PublicRestPublicTimeResRow struct {
	Ts string `json:"ts"` //系统时间，Unix时间戳的毫秒数格式，如 1597026383085
}

type PublicRestPublicMarkPriceRes []PublicRestPublicMarkPriceResRow
type PublicRestPublicMarkPriceResRow struct {
	InstType string `json:"instType"` //产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId   string `json:"instId"`   //产品ID，如 BTC-USD-200214
	MarkPx   string `json:"markPx"`   //标记价格
	Ts       string `json:"ts"`       //接口数据返回时间，Unix时间戳的毫秒数格式，如1597026383085
}

type PublicRestPublicFundingRateRes []PublicRestPublicFundingRateResRow
type PublicRestPublicFundingRateResRow struct {
	InstType        string `json:"instType"`        //产品类型 SWAP：永续合约
	InstId          string `json:"instId"`          //产品ID，如BTC-USD-SWAP
	FundingRate     string `json:"fundingRate"`     //资金费率
	NextFundingRate string `json:"nextFundingRate"` //下一期预测资金费率
	FundingTime     string `json:"fundingTime"`     //资金费时间 ，Unix时间戳的毫秒数格式，如 1597026383085
	NextFundingTime string `json:"nextFundingTime"` //下一期资金费时间 ，Unix时间戳的毫秒数格式，如 1622851200000
}

func (r *PublicRestMarketCandlesRes) Reverse() {
	for i, j := 0, len(*r)-1; i < j; i, j = i+1, j-1 {
		(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
	}
}
