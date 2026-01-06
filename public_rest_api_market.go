package myokxapi

// okx MarketTickers PublicRest接口 GET 获取指数行情
func (client *PublicRestClient) NewPublicRestMarketTickers() *PublicRestMarketTickersAPI {
	return &PublicRestMarketTickersAPI{
		client: client,
		req:    &PublicRestMarketTickersReq{},
	}
}
func (api *PublicRestMarketTickersAPI) Do() (*OkxRestRes[PublicRestMarketTickersRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketTickers])
	return okxCallAPI[PublicRestMarketTickersRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx MarketBooks PublicRest接口 GET 获取产品轻量深度
func (client *PublicRestClient) NewPublicRestMarketBooks() *PublicRestMarketBooksAPI {
	return &PublicRestMarketBooksAPI{
		client: client,
		req:    &PublicRestMarketBooksReq{},
	}
}
func (api *PublicRestMarketBooksAPI) Do() (*OkxRestRes[PublicRestMarketBooksRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketBooks])
	res, err := okxCallAPI[PublicRestMarketBooksMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	res2 := &OkxRestRes[PublicRestMarketBooksRes]{
		OkxErrorRes: res.OkxErrorRes,
		OkxTimeRes:  res.OkxTimeRes,
		Data:        *res.Data.ConvertToRes(),
	}

	return res2, nil
}

// okx MarketCandles PublicRest接口 GET 获取K线数据
func (client *PublicRestClient) NewPublicRestMarketCandles() *PublicRestMarketCandlesAPI {
	return &PublicRestMarketCandlesAPI{
		client: client,
		req:    &PublicRestMarketCandlesReq{},
	}
}
func (api *PublicRestMarketCandlesAPI) Do() (*OkxRestRes[PublicRestMarketCandlesRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketCandles])
	res, err := okxCallAPI[PublicRestMarketCandlesMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	res2 := &OkxRestRes[PublicRestMarketCandlesRes]{
		OkxErrorRes: res.OkxErrorRes,
		OkxTimeRes:  res.OkxTimeRes,
		Data:        *res.Data.ConvertToRes(),
	}
	return res2, nil
}

// okx MarketHistoryCandles PublicRest接口 GET 获取历史K线数据
func (client *PublicRestClient) NewPublicRestMarketHistoryCandles() *PublicRestMarketHistoryCandlesAPI {
	return &PublicRestMarketHistoryCandlesAPI{
		client: client,
		req:    &PublicRestMarketHistoryCandlesReq{},
	}
}
func (api *PublicRestMarketHistoryCandlesAPI) Do() (*OkxRestRes[PublicRestMarketCandlesRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketHistoryCandles])
	res, err := okxCallAPI[PublicRestMarketCandlesMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	res2 := &OkxRestRes[PublicRestMarketCandlesRes]{
		OkxErrorRes: res.OkxErrorRes,
		OkxTimeRes:  res.OkxTimeRes,
		Data:        *res.Data.ConvertToRes(),
	}
	return res2, nil
}

// okx PublicRestMarketTrades PublicRest接口 GET 获取交易产品公共成交数据
func (client *PublicRestClient) NewPublicRestMarketTrades() *PublicRestMarketTradesAPI {
	return &PublicRestMarketTradesAPI{
		client: client,
		req:    &PublicRestMarketTradesReq{},
	}
}
func (api *PublicRestMarketTradesAPI) Do() (*OkxRestRes[PublicRestMarketTradesRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketTrades])
	return okxCallAPI[PublicRestMarketTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx MarketHistoryTrades PublicRest接口 GET 获取交易产品公共历史成交数据 可以分页获取最近3个月的数据。
func (client *PublicRestClient) NewPublicRestMarketHistoryTrades() *PublicRestMarketHistoryTradesAPI {
	return &PublicRestMarketHistoryTradesAPI{
		client: client,
		req:    &PublicRestMarketHistoryTradesReq{},
	}
}
func (api *PublicRestMarketHistoryTradesAPI) Do() (*OkxRestRes[PublicRestMarketHistoryTradesRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketHistoryTrades])
	return okxCallAPI[PublicRestMarketHistoryTradesRes](api.client.c, url, NIL_REQBODY, GET)
}
