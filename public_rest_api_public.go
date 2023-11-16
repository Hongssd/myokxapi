package myokxapi

// okx PublicRestPublicInstruments PublicRest接口 GET 获取交易产品基础信息
func (client *PublicRestClient) NewPublicRestPublicInstruments() *PublicRestPublicInstrumentsAPI {
	return &PublicRestPublicInstrumentsAPI{
		client: client,
		req:    &PublicRestPublicInstrumentsReq{},
	}
}
func (api *PublicRestPublicInstrumentsAPI) Do() (*OkxRestRes[PublicRestPublicInstrumentsRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestPublicInstruments])
	return okxCallAPI[PublicRestPublicInstrumentsRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PublicRestPublicTime PublicRest接口 GET 获取系统时间
func (client *PublicRestClient) NewPublicRestPublicTime() *PublicRestPublicTimeAPI {
	return &PublicRestPublicTimeAPI{
		client: client,
		req:    &PublicRestPublicTimeReq{},
	}
}
func (api *PublicRestPublicTimeAPI) Do() (*OkxRestRes[PublicRestPublicTimeRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestPublicTime])
	return okxCallAPI[PublicRestPublicTimeRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PublicRestPublicMarkPrice PublicRest接口 GET 获取标记价格
func (client *PublicRestClient) NewPublicRestPublicMarkPrice() *PublicRestPublicMarkPriceAPI {
	return &PublicRestPublicMarkPriceAPI{
		client: client,
		req:    &PublicRestPublicMarkPriceReq{},
	}
}
func (api *PublicRestPublicMarkPriceAPI) Do() (*OkxRestRes[PublicRestPublicMarkPriceRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestPublicMarkPrice])
	return okxCallAPI[PublicRestPublicMarkPriceRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PublicRestPublicFundingRate PublicRest接口 GET 获取永续合约当前资金费率
func (client *PublicRestClient) NewPublicRestPublicFundingRate() *PublicRestPublicFundingRateAPI {
	return &PublicRestPublicFundingRateAPI{
		client: client,
		req:    &PublicRestPublicFundingRateReq{},
	}
}
func (api *PublicRestPublicFundingRateAPI) Do() (*OkxRestRes[PublicRestPublicFundingRateRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestPublicFundingRate])
	return okxCallAPI[PublicRestPublicFundingRateRes](api.client.c, url, NIL_REQBODY, GET)
}

