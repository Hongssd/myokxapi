package myokxapi

type PublicRestInstrumentsReq struct {
	InstType   *string `json:"instType"`   //String 是   产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String 可选 标的指数，如 BTC-USD，仅适用于交割/永续/期权，期权必填
	InstFamily *string `json:"instFamily"` //String 否   交易品种，如 BTC-USD，仅适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String 否   产品ID，如 BTC-USD-200213
}

type PublicRestInstrumentsApi struct {
	PublicRestClient
	req *PublicRestInstrumentsReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PublicRestInstrumentsApi) InstType(instType string) *PublicRestInstrumentsApi {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 可选 标的指数，如 BTC-USD，仅适用于交割/永续/期权，期权必填
func (api *PublicRestInstrumentsApi) Uly(uly string) *PublicRestInstrumentsApi {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestInstrumentsApi) InstFamily(instFamily string) *PublicRestInstrumentsApi {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如 BTC-USD-200213
func (api *PublicRestInstrumentsApi) InstId(instId string) *PublicRestInstrumentsApi {
	api.req.InstId = GetPointer(instId)
	return api
}

type PublicRestTimeReq struct {
}

type PublicRestTimeApi struct {
	PublicRestClient
	req *PublicRestTimeReq
}

type PublicRestMarkPriceReq struct {
	InstType   *string `json:"instType"`   //String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String	否 标的指数，如 BTC-USD，仅适用于交割/永续/期权
	InstFamily *string `json:"instFamily"` //String	否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String	否 产品ID，如 BTC-USD-200213
}

type PublicRestMarkPriceApi struct {
	PublicRestClient
	req *PublicRestMarkPriceReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PublicRestMarkPriceApi) InstType(instType string) *PublicRestMarkPriceApi {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 标的指数，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestMarkPriceApi) Uly(uly string) *PublicRestMarkPriceApi {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestMarkPriceApi) InstFamily(instFamily string) *PublicRestMarkPriceApi {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如 BTC-USD-200213
func (api *PublicRestMarkPriceApi) InstId(instId string) *PublicRestMarkPriceApi {
	api.req.InstId = GetPointer(instId)
	return api
}

type PublicRestMarketTickersReq struct {
	QuoteCcy *string `json:"quoteCcy"` //String	可选 指数计价单位， 目前只有 USD/USDT/BTC/USDC为计价单位的指数，quoteCcy和instId必须填写一个
	InstId   *string `json:"instId"`   //String	可选 指数，如 BTC-USD
}

type PublicRestMarketTickersApi struct {
	PublicRestClient
	req *PublicRestMarketTickersReq
}

// String 可选 指数计价单位， 目前只有 USD/USDT/BTC/USDC为计价单位的指数，quoteCcy和instId必须填写一个
func (api *PublicRestMarketTickersApi) QuoteCcy(quoteCcy string) *PublicRestMarketTickersApi {
	api.req.QuoteCcy = GetPointer(quoteCcy)
	return api
}

// String 可选 指数，如 BTC-USD
func (api *PublicRestMarketTickersApi) InstId(instId string) *PublicRestMarketTickersApi {
	api.req.InstId = GetPointer(instId)
	return api
}

// instId	String	是	产品ID，如 BTC-USDT
type PublicRestMarketBooksLiteReq struct {
	InstId *string `json:"instId"` //String	是	产品ID，如 BTC-USDT
}

type PublicRestMarketBooksLiteApi struct {
	PublicRestClient
	req *PublicRestMarketBooksLiteReq
}

// String 是 产品ID，如 BTC-USDT
func (api *PublicRestMarketBooksLiteApi) InstId(instId string) *PublicRestMarketBooksLiteApi {
	api.req.InstId = GetPointer(instId)
	return api
}

// instId	String	是	产品ID，如BTC-USD-190927-5000-C
// bar	String	否	时间粒度，默认值1m
// 	如 [1m/3m/5m/15m/30m/1H/2H/4H]
// 	香港时间开盘价k线：[6H/12H/1D/2D/3D/1W/1M/3M]
// 	UTC时间开盘价k线：[/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc]
// after	String	否	请求此时间戳之前（更旧的数据）的分页内容，传的值为对应接口的ts
// before	String	否	请求此时间戳之后（更新的数据）的分页内容，传的值为对应接口的ts, 单独使用时，会返回最新的数据。
// limit	String	否	分页返回的结果集数量，最大为300，不填默认返回100条

type PublicRestMarketCandlesReq struct {
	InstId *string `json:"instId"` //String	是	产品ID，如BTC-USD-190927-5000-C
	Bar    *string `json:"bar"`    //String	否	时间粒度，默认值1m
	After  *int64  `json:"after"`  //String	否	请求此时间戳之前（更旧的数据）的分页内容，传的值为对应接口的ts
	Before *int64  `json:"before"` //String	否	请求此时间戳之后（更新的数据）的分页内容，传的值为对应接口的ts, 单独使用时，会返回最新的数据。
	Limit  *int    `json:"limit"`  //String	否	分页返回的结果集数量，最大为300，不填默认返回100条
}

type PublicRestMarketCandlesApi struct {
	PublicRestClient
	req *PublicRestMarketCandlesReq
}

// String 是 产品ID，如BTC-USD-190927-5000-C
func (api *PublicRestMarketCandlesApi) InstId(instId string) *PublicRestMarketCandlesApi {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 时间粒度，默认值1m
// 如 [1m/3m/5m/15m/30m/1H/2H/4H]
// 香港时间开盘价k线：[6H/12H/1D/2D/3D/1W/1M/3M]
// UTC时间开盘价k线：[/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc]
func (api *PublicRestMarketCandlesApi) Bar(bar string) *PublicRestMarketCandlesApi {
	api.req.Bar = GetPointer(bar)
	return api
}

// String 否 请求此时间戳之前（更旧的数据）的分页内容，传的值为对应接口的ts
func (api *PublicRestMarketCandlesApi) After(after int64) *PublicRestMarketCandlesApi {
	api.req.After = GetPointer(after)
	return api
}

// String 否 请求此时间戳之后（更新的数据）的分页内容，传的值为对应接口的ts, 单独使用时，会返回最新的数据。
func (api *PublicRestMarketCandlesApi) Before(before int64) *PublicRestMarketCandlesApi {
	api.req.Before = GetPointer(before)
	return api
}

// String 否 分页返回的结果集数量，最大为300，不填默认返回100条
func (api *PublicRestMarketCandlesApi) Limit(limit int) *PublicRestMarketCandlesApi {
	api.req.Limit = GetPointer(limit)
	return api
}
