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

// okx MarketBooksLite PublicRest接口 GET 获取产品轻量深度
func (client *PublicRestClient) NewPublicRestMarketBooksLite() *PublicRestMarketBooksLiteAPI {
	return &PublicRestMarketBooksLiteAPI{
		client: client,
		req:    &PublicRestMarketBooksLiteReq{},
	}
}
func (api *PublicRestMarketBooksLiteAPI) Do() (*OkxRestRes[PublicRestMarketBooksLiteRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarketBooksLite])
	res, err := okxCallAPI[PublicRestMarketBooksLiteMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	res2 := &OkxRestRes[PublicRestMarketBooksLiteRes]{
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
