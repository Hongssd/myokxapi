package myokxapi

// Account
// okx PrivateRestAccountBalance PrivateRest接口 GET 查看账户余额
func (client *PrivateRestClient) NewPrivateRestAccountBalance() *PrivateRestAccountBalanceAPI {
	return &PrivateRestAccountBalanceAPI{
		client: client,
		req:    &PrivateRestAccountBalanceReq{},
	}
}
func (api *PrivateRestAccountBalanceAPI) Do() (*OkxRestRes[PrivateRestAccountBalanceRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountBalance])
	return okxCallAPIWithSecret[PrivateRestAccountBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestAccountPosition PrivateRest接口 GET 查看持仓信息
func (client *PrivateRestClient) NewPrivateRestAccountPosition() *PrivateRestAccountPositionAPI {
	return &PrivateRestAccountPositionAPI{
		client: client,
		req:    &PrivateRestAccountPositionReq{},
	}
}
func (api *PrivateRestAccountPositionAPI) Do() (*OkxRestRes[PrivateRestAccountPositionRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountPosition])
	return okxCallAPIWithSecret[PrivateRestAccountPositionRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestAccountConfig PrivateRest接口 GET 查看账户配置
func (client *PrivateRestClient) NewPrivateRestAccountConfig() *PrivateRestAccountConfigAPI {
	return &PrivateRestAccountConfigAPI{
		client: client,
		req:    &PrivateRestAccountConfigReq{},
	}
}
func (api *PrivateRestAccountConfigAPI) Do() (*OkxRestRes[PrivateRestAccountConfigRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountConfig])
	return okxCallAPIWithSecret[PrivateRestAccountConfigRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestAccountTradeFee PrivateRest接口 GET 查看账户手续费费率
func (client *PrivateRestClient) NewPrivateRestAccountTradeFee() *PrivateRestAccountTradeFeeAPI {
	return &PrivateRestAccountTradeFeeAPI{
		client: client,
		req:    &PrivateRestAccountTradeFeeReq{},
	}
}
func (api *PrivateRestAccountTradeFeeAPI) Do() (*OkxRestRes[PrivateRestAccountTradeFeeRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountTradeFee])
	return okxCallAPIWithSecret[PrivateRestAccountTradeFeeRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestAccountSetLeverage PrivateRest接口 POST 设置杠杆倍数
func (client *PrivateRestClient) NewPrivateRestAccountSetLeverage() *PrivateRestAccountSetLeverageAPI {
	return &PrivateRestAccountSetLeverageAPI{
		client: client,
		req:    &PrivateRestAccountSetLeverageReq{},
	}
}
func (api *PrivateRestAccountSetLeverageAPI) Do() (*OkxRestRes[PrivateRestAccountSetLeverageRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountSetLeverage])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestAccountSetLeverageRes](api.client.c, url, reqBody, POST)
}
