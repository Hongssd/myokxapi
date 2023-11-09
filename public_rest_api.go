package myokxapi

type PublicRestApi int

const (

	//GET
	PublicRestInstruments     PublicRestApi = iota //获取交易产品基础信息
	PublicRestTime                                 //获取系统时间
	PublicRestMarkPrice                            //获取标记价格
	PublicRestMarketTickers                        //获取指数行情
	PublicRestMarketBooksLite                      // 获取产品轻量深度
	PublicRestMarketCandles                        // 获取K线数据

)

var PublicRestApiMap = map[PublicRestApi]string{
	PublicRestInstruments:   "/api/v5/public/instruments",   //GET 获取交易产品基础信息
	PublicRestTime:          "/api/v5/public/time",          //GET 获取系统时间
	PublicRestMarkPrice:     "/api/v5/public/mark-price",    //GET 获取标记价格
	PublicRestMarketTickers: "/api/v5/market/index-tickers", //GET 获取指数行情

	PublicRestMarketBooksLite: "/api/v5/market/books-lite", // 获取产品轻量深度
	PublicRestMarketCandles:   "/api/v5/market/candles",    // 获取K线数据

}

// okx PublicRestInstruments PublicRest接口 GET 获取交易产品基础信息
func (client *PublicRestClient) NewPublicRestInstruments() *PublicRestInstrumentsApi {
	return &PublicRestInstrumentsApi{
		PublicRestClient: *client,
		req:              &PublicRestInstrumentsReq{},
	}
}
func (api *PublicRestInstrumentsApi) Do() (*PublicRestInstrumentsRes, error) {
	url := okxHandlerRequestApi(REST, api.req, PublicRestApiMap[PublicRestInstruments])
	return okxCallApiWithSecret[PublicRestInstrumentsRes](api.PublicRestClient.c, url, GET)
}

// okx PublicRestTime PublicRest接口 GET 获取系统时间
func (client *PublicRestClient) NewPublicRestTime() *PublicRestTimeApi {
	return &PublicRestTimeApi{
		PublicRestClient: *client,
		req:              &PublicRestTimeReq{},
	}
}
func (api *PublicRestTimeApi) Do() (*PublicRestTimeRes, error) {
	url := okxHandlerRequestApi(REST, api.req, PublicRestApiMap[PublicRestTime])
	return okxCallApiWithSecret[PublicRestTimeRes](api.PublicRestClient.c, url, GET)
}

// okx PublicRestMarkPrice PublicRest接口 GET 获取标记价格
func (client *PublicRestClient) NewPublicRestMarkPrice() *PublicRestMarkPriceApi {
	return &PublicRestMarkPriceApi{
		PublicRestClient: *client,
		req:              &PublicRestMarkPriceReq{},
	}
}
func (api *PublicRestMarkPriceApi) Do() (*PublicRestMarkPriceRes, error) {
	url := okxHandlerRequestApi(REST, api.req, PublicRestApiMap[PublicRestMarkPrice])
	return okxCallApiWithSecret[PublicRestMarkPriceRes](api.PublicRestClient.c, url, GET)
}

// okx MarketTickers PublicRest接口 GET 获取指数行情
func (client *PublicRestClient) NewPublicRestMarketTickers() *PublicRestMarketTickersApi {
	return &PublicRestMarketTickersApi{
		PublicRestClient: *client,
		req:              &PublicRestMarketTickersReq{},
	}
}
func (api *PublicRestMarketTickersApi) Do() (*PublicRestMarketTickersRes, error) {
	url := okxHandlerRequestApi(REST, api.req, PublicRestApiMap[PublicRestMarketTickers])
	return okxCallApiWithSecret[PublicRestMarketTickersRes](api.PublicRestClient.c, url, GET)
}

// okx MarketBooksLite PublicRest接口 GET 获取产品轻量深度
func (client *PublicRestClient) NewPublicRestMarketBooksLite() *PublicRestMarketBooksLiteApi {
	return &PublicRestMarketBooksLiteApi{
		PublicRestClient: *client,
		req:              &PublicRestMarketBooksLiteReq{},
	}
}
func (api *PublicRestMarketBooksLiteApi) Do() (*PublicRestMarketBooksLiteRes, error) {
	url := okxHandlerRequestApi(REST, api.req, PublicRestApiMap[PublicRestMarketBooksLite])
	res, err := okxCallApiWithSecret[PublicRestMarketBooksLiteMiddle](api.PublicRestClient.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}

// okx MarketCandles PublicRest接口 GET 获取K线数据
func (client *PublicRestClient) NewPublicRestMarketCandles() *PublicRestMarketCandlesApi {
	return &PublicRestMarketCandlesApi{
		PublicRestClient: *client,
		req:              &PublicRestMarketCandlesReq{},
	}
}
func (api *PublicRestMarketCandlesApi) Do() (*PublicRestMarketCandlesRes, error) {
	url := okxHandlerRequestApi(REST, api.req, PublicRestApiMap[PublicRestMarketCandles])
	res, err := okxCallApiWithSecret[PublicRestMarketCandlesMiddle](api.PublicRestClient.c, url, GET)
	if err != nil {
		return nil, err
	}
	return res.ConvertToRes(), nil
}
