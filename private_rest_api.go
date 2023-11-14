package myokxapi

type PrivateRestAPI int

const (
	//Account
	PrivateRestAccountBalance  PrivateRestAPI = iota //查看账户余额
	PrivateRestAccountPosition                       //查看持仓信息
	PrivateRestAccountConfig                         //查看账户配置

	//Trade
	PrivateRestTradeOrderGet          //查看订单信息
	PrivateRestTradeOrdersPending     //查看未成交订单列表
	PrivateRestTradeOrderPost         //下单
	PrivateRestTradeCancelOrder       //撤单
	PrivateRestTradeAmendOrder        //修改订单
	PrivateRestTradeBatchOrders       //批量下单
	PrivateRestTradeCancelBatchOrders //批量撤单
	PrivateRestTradeAmendBatchOrders  //批量修改订单

)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	//Account
	PrivateRestAccountBalance:  "/api/v5/account/balance",   //GET 查看账户余额
	PrivateRestAccountPosition: "/api/v5/account/positions", //GET 查看持仓信息
	PrivateRestAccountConfig:   "/api/v5/account/config",    //GET 查看账户配置

	//Trade
	PrivateRestTradeOrderGet:          "/api/v5/trade/order",               //GET 查看订单信息
	PrivateRestTradeOrdersPending:     "/api/v5/trade/orders-pending",      //GET 查看未成交订单列表
	PrivateRestTradeOrderPost:         "/api/v5/trade/order",               //POST 下单
	PrivateRestTradeCancelOrder:       "/api/v5/trade/cancel-order",        //POST 撤单
	PrivateRestTradeAmendOrder:        "/api/v5/trade/amend-order",         //POST 修改订单
	PrivateRestTradeBatchOrders:       "/api/v5/trade/batch-orders",        //POST 批量下单
	PrivateRestTradeCancelBatchOrders: "/api/v5/trade/cancel-batch-orders", //POST 批量撤单
	PrivateRestTradeAmendBatchOrders:  "/api/v5/trade/amend-batch-orders",  //POST 批量修改订单
}

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
