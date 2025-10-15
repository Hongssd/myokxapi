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

// okx PrivateRestAccountLeverageInfo PrivateRest接口 GET 获取杠杆倍数
func (client *PrivateRestClient) NewPrivateRestAccountLeverageInfo() *PrivateRestAccountLeverageInfoAPI {
	return &PrivateRestAccountLeverageInfoAPI{
		client: client,
		req:    &PrivateRestAccountLeverageInfoReq{},
	}
}
func (api *PrivateRestAccountLeverageInfoAPI) Do() (*OkxRestRes[PrivateRestAccountLeverageInfoRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountLeverageInfo])
	return okxCallAPIWithSecret[PrivateRestAccountLeverageInfoRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestAccountSetLeverage PrivateRest接口 POST 设置杠杆倍数
func (client *PrivateRestClient) NewPrivateRestAccountSetLeverage() *PrivateRestAccountSetLeverageAPI {
	return &PrivateRestAccountSetLeverageAPI{
		client: client,
		req:    &PrivateRestAccountSetLeverageReq{},
	}
}
func (api *PrivateRestAccountSetLeverageAPI) Do() (*OkxRestRes[PrivateRestAccountSetLeverageRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestAccountSetLeverage])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestAccountSetLeverageRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestAccountSetPositionMode PrivateRest接口 POST 设置持仓模式
func (client *PrivateRestClient) NewPrivateRestAccountSetPositionMode() *PrivateRestAccountSetPositionModeAPI {
	return &PrivateRestAccountSetPositionModeAPI{
		client: client,
		req:    &PrivateRestAccountSetPositionModeReq{},
	}
}
func (api *PrivateRestAccountSetPositionModeAPI) Do() (*OkxRestRes[PrivateRestAccountSetPositionModeRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestAccountSetPositionMode])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestAccountSetPositionModeRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestAccountSetAccountLevel PrivateRest接口 POST 设置账户模式
func (client *PrivateRestClient) NewPrivateRestAccountSetAccountLevel() *PrivateRestAccountSetAccountLevelAPI {
	return &PrivateRestAccountSetAccountLevelAPI{
		client: client,
		req:    &PrivateRestAccountSetAccountLevelReq{},
	}
}
func (api *PrivateRestAccountSetAccountLevelAPI) Do() (*OkxRestRes[PrivateRestAccountSetAccountLevelRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestAccountSetAccountLevel])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestAccountSetAccountLevelRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestAccountPositionsHistory PrivateRest接口 GET 查看持仓历史
func (client *PrivateRestClient) NewPrivateRestAccountPositionsHistory() *PrivateRestAccountPositionsHistoryAPI {
	return &PrivateRestAccountPositionsHistoryAPI{
		client: client,
		req:    &PrivateRestAccountPositionsHistoryReq{},
	}
}
func (api *PrivateRestAccountPositionsHistoryAPI) Do() (*OkxRestRes[PrivateRestAccountPositionsHistoryRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountPositionsHistory])
	return okxCallAPIWithSecret[PrivateRestAccountPositionsHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// Asset

// okx PrivateRestAssetBills PrivateRest接口 GET 查询最近一个月内资金账户账单流水
func (client *PrivateRestClient) NewPrivateRestAssetBills() *PrivateRestAssetBillsAPI {
	return &PrivateRestAssetBillsAPI{
		client: client,
		req:    &PrivateRestAssetBillsReq{},
	}
}
func (api *PrivateRestAssetBillsAPI) Do() (*OkxRestRes[PrivateRestAssetBillsRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAssetBills])
	return okxCallAPIWithSecret[PrivateRestAssetBillsRes](api.client.c, url, NIL_REQBODY, GET)
}

// SubAccount

// okx PrivateRestSubAccountSetTransferOut PrivateRest接口 POST 设置子账户主动转出权限
func (client *PrivateRestClient) NewPrivateRestSubAccountSetTransferOut() *PrivateRestSubAccountSetTransferOutAPI {
	return &PrivateRestSubAccountSetTransferOutAPI{
		client: client,
		req:    &PrivateRestSubAccountSetTransferOutReq{},
	}
}
func (api *PrivateRestSubAccountSetTransferOutAPI) Do() (*OkxRestRes[PrivateRestSubAccountSetTransferOutRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestSubAccountSetTransferOut])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestSubAccountSetTransferOutRes](api.client.c, url, reqBody, POST)
}
