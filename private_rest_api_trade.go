package myokxapi

// Trade
// okx PrivateRestTradeOrderGet PrivateRest接口 GET 查看订单信息
func (client *PrivateRestClient) NewPrivateRestTradeOrderGet() *PrivateRestTradeOrderGetAPI {
	return &PrivateRestTradeOrderGetAPI{
		client: client,
		req:    &PrivateRestTradeOrderGetReq{},
	}
}
func (api *PrivateRestTradeOrderGetAPI) Do() (*OkxRestRes[PrivateRestTradeOrderGetRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrderGet])
	return okxCallAPIWithSecret[PrivateRestTradeOrderGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestTradeOrdersPending PrivateRest接口 GET 查看未成交订单列表
func (client *PrivateRestClient) NewPrivateRestTradeOrdersPending() *PrivateRestTradeOrdersPendingAPI {
	return &PrivateRestTradeOrdersPendingAPI{
		client: client,
		req:    &PrivateRestTradeOrdersPendingReq{},
	}
}
func (api *PrivateRestTradeOrdersPendingAPI) Do() (*OkxRestRes[PrivateRestTradeOrdersPendingRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrdersPending])
	return okxCallAPIWithSecret[PrivateRestTradeOrdersPendingRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestTradeOrderPost PrivateRest接口 POST 下单
func (client *PrivateRestClient) NewPrivateRestTradeOrderPost() *PrivateRestTradeOrderPostAPI {
	return &PrivateRestTradeOrderPostAPI{
		client: client,
		req:    &PrivateRestTradeOrderPostReq{},
	}
}
func (api *PrivateRestTradeOrderPostAPI) Do() (*OkxRestRes[PrivateRestTradeOrderPostRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeOrderPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeOrderPostRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeOrderAlgoPost PrivateRest接口 POST 策略委托下单
func (client *PrivateRestClient) NewPrivateRestTradeOrderAlgoPost() *PrivateRestTradeOrderAlgoPostAPI {
	return &PrivateRestTradeOrderAlgoPostAPI{
		client: client,
		req:    &PrivateRestTradeOrderAlgoPostReq{},
	}
}
func (api *PrivateRestTradeOrderAlgoPostAPI) Do() (*OkxRestRes[PrivateRestTradeOrderAlgoPostRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeOrderAlgoPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeOrderAlgoPostRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeCancelOrderAlgo PrivateRest接口 POST 撤销策略委托订单
func (client *PrivateRestClient) NewPrivateRestTradeCancelOrderAlgo() *PrivateRestTradeCancelOrderAlgoAPI {
	return &PrivateRestTradeCancelOrderAlgoAPI{
		client: client,
		req:    &PrivateRestTradeCancelOrderAlgoReq{},
	}
}

func (api *PrivateRestTradeCancelOrderAlgoAPI) Do() (*OkxRestRes[PrivateRestTradeCancelOrderAlgoRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeCancelOrderAlgo])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeCancelOrderAlgoRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeAmendOrderAlgo PrivateRest接口 POST 修改策略委托订单
func (client *PrivateRestClient) NewPrivateRestTradeAmendOrderAlgo() *PrivateRestTradeAmendOrderAlgoAPI {
	return &PrivateRestTradeAmendOrderAlgoAPI{
		client: client,
		req:    &PrivateRestTradeAmendOrderAlgoReq{},
	}
}
func (api *PrivateRestTradeAmendOrderAlgoAPI) Do() (*OkxRestRes[PrivateRestTradeAmendOrderAlgoRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeAmendOrderAlgo])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeAmendOrderAlgoRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeOrderAlgoGet PrivateRest接口 GET 获取策略委托单信息
func (client *PrivateRestClient) NewPrivateRestTradeOrderAlgoGet() *PrivateRestTradeOrderAlgoGetAPI {
	return &PrivateRestTradeOrderAlgoGetAPI{
		client: client,
		req:    &PrivateRestTradeOrderAlgoGetReq{},
	}
}
func (api *PrivateRestTradeOrderAlgoGetAPI) Do() (*OkxRestRes[PrivateRestTradeOrderAlgoGetRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrderAlgoGet])
	return okxCallAPIWithSecret[PrivateRestTradeOrderAlgoGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestTradeCancelOrder PrivateRest接口 POST 撤单
func (client *PrivateRestClient) NewPrivateRestTradeCancelOrder() *PrivateRestTradeCancelOrderAPI {
	return &PrivateRestTradeCancelOrderAPI{
		client: client,
		req:    &PrivateRestTradeCancelOrderReq{},
	}
}
func (api *PrivateRestTradeCancelOrderAPI) Do() (*OkxRestRes[PrivateRestTradeCancelOrderRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeCancelOrder])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeCancelOrderRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeAmendOrder PrivateRest接口 POST 修改订单
func (client *PrivateRestClient) NewPrivateRestTradeAmendOrder() *PrivateRestTradeAmendOrderAPI {
	return &PrivateRestTradeAmendOrderAPI{
		client: client,
		req:    &PrivateRestTradeAmendOrderReq{},
	}
}
func (api *PrivateRestTradeAmendOrderAPI) Do() (*OkxRestRes[PrivateRestTradeAmendOrderRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeAmendOrder])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeAmendOrderRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeOrderAlgoPending PrivateRest接口 GET 获取未完成策略委托单列表
func (client *PrivateRestClient) NewPrivateRestTradeOrderAlgoPending() *PrivateRestTradeOrderAlgoPendingAPI {
	return &PrivateRestTradeOrderAlgoPendingAPI{
		client: client,
		req:    &PrivateRestTradeOrderAlgoPendingReq{},
	}
}
func (api *PrivateRestTradeOrderAlgoPendingAPI) Do() (*OkxRestRes[PrivateRestTradeOrderAlgoPendingRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrderAlgoPending])
	return okxCallAPIWithSecret[PrivateRestTradeOrderAlgoPendingRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestTradeOrderAlgoHistory PrivateRest接口 GET 获取历史策略委托单记录
func (client *PrivateRestClient) NewPrivateRestTradeOrderAlgoHistory() *PrivateRestTradeOrderAlgoHistoryAPI {
	return &PrivateRestTradeOrderAlgoHistoryAPI{
		client: client,
		req:    &PrivateRestTradeOrderAlgoHistoryReq{},
	}
}
func (api *PrivateRestTradeOrderAlgoHistoryAPI) Do() (*OkxRestRes[PrivateRestTradeOrderAlgoHistoryRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrderAlgoHistory])
	return okxCallAPIWithSecret[PrivateRestTradeOrderAlgoHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestTradeBatchOrders PrivateRest接口 POST 批量下单
func (client *PrivateRestClient) NewPrivateRestTradeBatchOrders() *PrivateRestTradeBatchOrdersAPI {
	return &PrivateRestTradeBatchOrdersAPI{
		client: client,
		req:    &PrivateRestTradeBatchOrdersReq{},
	}
}
func (api *PrivateRestTradeBatchOrdersAPI) Do() (*OkxRestRes[PrivateRestTradeBatchOrdersRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeBatchOrders])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeBatchOrdersRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeCancelBatchOrders PrivateRest接口 POST 批量撤单
func (client *PrivateRestClient) NewPrivateRestTradeCancelBatchOrders() *PrivateRestTradeCancelBatchOrdersAPI {
	return &PrivateRestTradeCancelBatchOrdersAPI{
		client: client,
		req:    &PrivateRestTradeCancelBatchOrdersReq{},
	}
}
func (api *PrivateRestTradeCancelBatchOrdersAPI) Do() (*OkxRestRes[PrivateRestTradeCancelBatchOrdersRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeCancelBatchOrders])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeCancelBatchOrdersRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeAmendBatchOrders PrivateRest接口 POST 批量修改订单
func (client *PrivateRestClient) NewPrivateRestTradeAmendBatchOrders() *PrivateRestTradeAmendBatchOrdersAPI {
	return &PrivateRestTradeAmendBatchOrdersAPI{
		client: client,
		req:    &PrivateRestTradeAmendBatchOrdersReq{},
	}
}
func (api *PrivateRestTradeAmendBatchOrdersAPI) Do() (*OkxRestRes[PrivateRestTradeAmendBatchOrdersRes], error) {
	url := okxHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestTradeAmendBatchOrders])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return okxCallAPIWithSecret[PrivateRestTradeAmendBatchOrdersRes](api.client.c, url, reqBody, POST)
}

// okx PrivateRestTradeOrderHistory PrivateRest接口 GET 获取历史订单记录（近七天）
func (client *PrivateRestClient) NewPrivateRestTradeOrderHistory() *PrivateRestTradeOrderHistoryAPI {
	return &PrivateRestTradeOrderHistoryAPI{
		client: client,
		req:    &PrivateRestTradeOrderHistoryReq{},
	}
}
func (api *PrivateRestTradeOrderHistoryAPI) Do() (*OkxRestRes[PrivateRestTradeOrderHistoryRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrderHistory])
	return okxCallAPIWithSecret[PrivateRestTradeOrderHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestTradeOrderHistoryArchive PrivateRest接口 GET 获取历史订单记录（近三个月）
func (client *PrivateRestClient) NewPrivateRestTradeOrderHistoryArchive() *PrivateRestTradeOrderHistoryArchiveAPI {
	return &PrivateRestTradeOrderHistoryArchiveAPI{

		client: client,
		req:    &PrivateRestTradeOrderHistoryArchiveReq{},
	}
}
func (api *PrivateRestTradeOrderHistoryArchiveAPI) Do() (*OkxRestRes[PrivateRestTradeOrderHistoryArchiveRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeOrderHistoryArchive])
	return okxCallAPIWithSecret[PrivateRestTradeOrderHistoryArchiveRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestTradeFills PrivateRest接口 GET 获取成交明细（近三天）
func (client *PrivateRestClient) NewPrivateRestTradeFills() *PrivateRestTradeFillsAPI {
	return &PrivateRestTradeFillsAPI{
		client: client,
		req:    &PrivateRestTradeFillsReq{},
	}
}
func (api *PrivateRestTradeFillsAPI) Do() (*OkxRestRes[PrivateRestTradeFillsRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeFills])
	return okxCallAPIWithSecret[PrivateRestTradeFillsRes](api.client.c, url, NIL_REQBODY, GET)
}

// okx PrivateRestTradeFillsHistory PrivateRest接口 GET 获取成交明细（近三个月）
func (client *PrivateRestClient) NewPrivateRestTradeFillsHistory() *PrivateRestTradeFillsHistoryAPI {
	return &PrivateRestTradeFillsHistoryAPI{
		PrivateRestTradeFillsAPI{
			client: client,
			req:    &PrivateRestTradeFillsReq{},
		},
	}
}
func (api *PrivateRestTradeFillsHistoryAPI) Do() (*OkxRestRes[PrivateRestTradeFillsHistoryRes], error) {
	url := okxHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestTradeFillsHistory])
	return okxCallAPIWithSecret[PrivateRestTradeFillsHistoryRes](api.client.c, url, NIL_REQBODY, GET)
}
