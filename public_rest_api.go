package myokxapi

type PublicRestAPI int

const (

	//GET
	PublicRestInstruments     PublicRestAPI = iota //获取交易产品基础信息
	PublicRestTime                                 //获取系统时间
	PublicRestMarkPrice                            //获取标记价格
	PublicRestMarketTickers                        //获取指数行情
	PublicRestMarketBooksLite                      // 获取产品轻量深度
	PublicRestMarketCandles                        // 获取K线数据

)

var PublicRestAPIMap = map[PublicRestAPI]string{
	PublicRestInstruments:     "/api/v5/public/instruments",   //GET 获取交易产品基础信息
	PublicRestTime:            "/api/v5/public/time",          //GET 获取系统时间
	PublicRestMarkPrice:       "/api/v5/public/mark-price",    //GET 获取标记价格
	PublicRestMarketTickers:   "/api/v5/market/index-tickers", //GET 获取指数行情
	PublicRestMarketBooksLite: "/api/v5/market/books-lite",    //GET 获取产品轻量深度
	PublicRestMarketCandles:   "/api/v5/market/candles",       //GET 获取K线数据

}

// okx PublicRestInstruments PublicRest接口 GET 获取交易产品基础信息
func (client *PublicRestClient) NewPublicRestInstruments() *PublicRestInstrumentsAPI {
	return &PublicRestInstrumentsAPI{
		client: client,
		req:    &PublicRestInstrumentsReq{},
	}
}
func (api *PublicRestInstrumentsAPI) Do() (*OkxRestRes[PublicRestInstrumentsRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestInstruments])
	return okxCallAPI[PublicRestInstrumentsRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PublicRestTime PublicRest接口 GET 获取系统时间
func (client *PublicRestClient) NewPublicRestTime() *PublicRestTimeAPI {
	return &PublicRestTimeAPI{
		client: client,
		req:    &PublicRestTimeReq{},
	}
}
func (api *PublicRestTimeAPI) Do() (*OkxRestRes[PublicRestTimeRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestTime])
	return okxCallAPI[PublicRestTimeRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PublicRestMarkPrice PublicRest接口 GET 获取标记价格
func (client *PublicRestClient) NewPublicRestMarkPrice() *PublicRestMarkPriceAPI {
	return &PublicRestMarkPriceAPI{
		client: client,
		req:    &PublicRestMarkPriceReq{},
	}
}
func (api *PublicRestMarkPriceAPI) Do() (*OkxRestRes[PublicRestMarkPriceRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarkPrice])
	return okxCallAPI[PublicRestMarkPriceRes](api.client.c, url, NIL_REQBODY, GET)
}

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
