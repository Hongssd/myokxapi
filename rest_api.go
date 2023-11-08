package myokxapi

type RestApi int

const (

	//GET
	RestInstruments   RestApi = iota //获取交易产品基础信息
	RestTime                         //获取系统时间
	RestMarkPrice                    //获取标记价格
	RestMarketTickers                //获取指数行情

)

var RestApiMap = map[RestApi]string{
	RestInstruments:   "/api/v5/public/instruments",   //GET 获取交易产品基础信息
	RestTime:          "/api/v5/public/time",          //GET 获取系统时间
	RestMarkPrice:     "/api/v5/public/mark-price",    //GET 获取标记价格
	RestMarketTickers: "/api/v5/market/index-tickers", //GET 获取指数行情

}

// okx RestInstruments Rest接口 GET 获取交易产品基础信息
func (client *RestClient) NewRestInstruments() *RestInstrumentsApi {
	return &RestInstrumentsApi{
		RestClient: *client,
		req:        &RestInstrumentsReq{},
	}
}
func (api *RestInstrumentsApi) Do() (*RestInstrumentsRes, error) {
	url := okxHandlerRequestApi(REST, api.req, RestApiMap[RestInstruments])
	return okxCallApiWithSecret[RestInstrumentsRes](api.RestClient.c, url, GET)
}


// okx RestTime Rest接口 GET 获取系统时间
func (client *RestClient) NewRestTime() *RestTimeApi {
	return &RestTimeApi{
		RestClient: *client,
		req:        &RestTimeReq{},
	}
}
func (api *RestTimeApi) Do() (*RestTimeRes, error) {
	url := okxHandlerRequestApi(REST, api.req, RestApiMap[RestTime])
	return okxCallApiWithSecret[RestTimeRes](api.RestClient.c, url, GET)
}

// okx RestMarkPrice Rest接口 GET 获取标记价格
func (client *RestClient) NewRestMarkPrice() *RestMarkPriceApi {
	return &RestMarkPriceApi{
		RestClient: *client,
		req:        &RestMarkPriceReq{},
	}
}
func (api *RestMarkPriceApi) Do() (*RestMarkPriceRes, error) {
	url := okxHandlerRequestApi(REST, api.req, RestApiMap[RestMarkPrice])
	return okxCallApiWithSecret[RestMarkPriceRes](api.RestClient.c, url, GET)
}

// okx MarketTickers Rest接口 GET 获取指数行情
func (client *RestClient) NewRestMarketTickers() *RestMarketTickersApi {
	return &RestMarketTickersApi{
		RestClient: *client,
		req:        &RestMarketTickersReq{},
	}
}
func (api *RestMarketTickersApi) Do() (*RestMarketTickersRes, error) {
	url := okxHandlerRequestApi(REST, api.req, RestApiMap[RestMarketTickers])
	return okxCallApiWithSecret[RestMarketTickersRes](api.RestClient.c, url, GET)
}
