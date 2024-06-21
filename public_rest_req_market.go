package myokxapi

type PublicRestPublicMarkPriceReq struct {
	InstType   *string `json:"instType"`   //String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String	否 标的指数，如 BTC-USD，仅适用于交割/永续/期权
	InstFamily *string `json:"instFamily"` //String	否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String	否 产品ID，如 BTC-USD-200213
}

type PublicRestPublicMarkPriceAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicMarkPriceReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PublicRestPublicMarkPriceAPI) InstType(instType string) *PublicRestPublicMarkPriceAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 标的指数，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestPublicMarkPriceAPI) Uly(uly string) *PublicRestPublicMarkPriceAPI {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestPublicMarkPriceAPI) InstFamily(instFamily string) *PublicRestPublicMarkPriceAPI {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如 BTC-USD-200213
func (api *PublicRestPublicMarkPriceAPI) InstId(instId string) *PublicRestPublicMarkPriceAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

type PublicRestMarketTickersReq struct {
	QuoteCcy *string `json:"quoteCcy"` //String	可选 指数计价单位， 目前只有 USD/USDT/BTC/USDC为计价单位的指数，quoteCcy和instId必须填写一个
	InstId   *string `json:"instId"`   //String	可选 指数，如 BTC-USD
}

type PublicRestMarketTickersAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketTickersReq
}

// String 可选 指数计价单位， 目前只有 USD/USDT/BTC/USDC为计价单位的指数，quoteCcy和instId必须填写一个
func (api *PublicRestMarketTickersAPI) QuoteCcy(quoteCcy string) *PublicRestMarketTickersAPI {
	api.req.QuoteCcy = GetPointer(quoteCcy)
	return api
}

// String 可选 指数，如 BTC-USD
func (api *PublicRestMarketTickersAPI) InstId(instId string) *PublicRestMarketTickersAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// instId	String	是	产品ID，如 BTC-USDT
// sz	String	否	深度档位数量，最大值可传400，即买卖深度共800条 不填写此参数，默认返回1档深度数据
type PublicRestMarketBooksReq struct {
	InstId *string `json:"instId"` //String	是	产品ID，如 BTC-USDT
	Sz     *string `json:"sz"`     //String	否	返回深度档位，最多返回150档
}

type PublicRestMarketBooksAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketBooksReq
}

// String 是 产品ID，如 BTC-USDT
func (api *PublicRestMarketBooksAPI) InstId(instId string) *PublicRestMarketBooksAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 返回深度档位，最多返回150档
func (api *PublicRestMarketBooksAPI) Sz(sz string) *PublicRestMarketBooksAPI {
	api.req.Sz = GetPointer(sz)
	return api
}

type PublicRestMarketCandlesReq struct {
	InstId *string `json:"instId"` //String	是	产品ID，如BTC-USD-190927-5000-C
	Bar    *string `json:"bar"`    //String	否	时间粒度，默认值1m
	After  *int64  `json:"after"`  //String	否	请求此时间戳之前（更旧的数据）的分页内容，传的值为对应接口的ts
	Before *int64  `json:"before"` //String	否	请求此时间戳之后（更新的数据）的分页内容，传的值为对应接口的ts, 单独使用时，会返回最新的数据。
	Limit  *int    `json:"limit"`  //String	否	分页返回的结果集数量，最大为300，不填默认返回100条
}

type PublicRestMarketCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketCandlesReq
}

// String 是 产品ID，如BTC-USD-190927-5000-C
func (api *PublicRestMarketCandlesAPI) InstId(instId string) *PublicRestMarketCandlesAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 时间粒度，默认值1m
// 如 [1m/3m/5m/15m/30m/1H/2H/4H]
// 香港时间开盘价k线：[6H/12H/1D/2D/3D/1W/1M/3M]
// UTC时间开盘价k线：[/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc]
func (api *PublicRestMarketCandlesAPI) Bar(bar string) *PublicRestMarketCandlesAPI {
	api.req.Bar = GetPointer(bar)
	return api
}

// String 否 请求此时间戳之前（更旧的数据）的分页内容，传的值为对应接口的ts
func (api *PublicRestMarketCandlesAPI) After(after int64) *PublicRestMarketCandlesAPI {
	api.req.After = GetPointer(after)
	return api
}

// String 否 请求此时间戳之后（更新的数据）的分页内容，传的值为对应接口的ts, 单独使用时，会返回最新的数据。
func (api *PublicRestMarketCandlesAPI) Before(before int64) *PublicRestMarketCandlesAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String 否 分页返回的结果集数量，最大为300，不填默认返回100条
func (api *PublicRestMarketCandlesAPI) Limit(limit int) *PublicRestMarketCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestMarketHistoryCandlesReq struct {
	InstId *string `json:"instId"` //String	是	产品ID，如BTC-USD-190927-5000-C
	Bar    *string `json:"bar"`    //String	否	时间粒度，默认值1m
	After  *int64  `json:"after"`  //String	否	请求此时间戳之前（更旧的数据）的分页内容，传的值为对应接口的ts
	Before *int64  `json:"before"` //String	否	请求此时间戳之后（更新的数据）的分页内容，传的值为对应接口的ts, 单独使用时，会返回最新的数据。
	Limit  *int    `json:"limit"`  //String	否	分页返回的结果集数量，最大为300，不填默认返回100条
}

type PublicRestMarketHistoryCandlesAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketHistoryCandlesReq
}

// String 是 产品ID，如BTC-USD-190927-5000-C
func (api *PublicRestMarketHistoryCandlesAPI) InstId(instId string) *PublicRestMarketHistoryCandlesAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 时间粒度，默认值1m
func (api *PublicRestMarketHistoryCandlesAPI) Bar(bar string) *PublicRestMarketHistoryCandlesAPI {
	api.req.Bar = GetPointer(bar)
	return api
}

// String 否 请求此时间戳之前（更旧的数据）的分页内容，传的值为对应接口的ts
func (api *PublicRestMarketHistoryCandlesAPI) After(after int64) *PublicRestMarketHistoryCandlesAPI {
	api.req.After = GetPointer(after)
	return api
}

// String 否 请求此时间戳之后（更新的数据）的分页内容，传的值为对应接口的ts, 单独使用时，会返回最新的数据。
func (api *PublicRestMarketHistoryCandlesAPI) Before(before int64) *PublicRestMarketHistoryCandlesAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String 否 分页返回的结果集数量，最大为100，不填默认返回100条
func (api *PublicRestMarketHistoryCandlesAPI) Limit(limit int) *PublicRestMarketHistoryCandlesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}
