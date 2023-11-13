package myokxapi

type PrivateRestAPI int

const (
	//Account
	PrivateRestAccountBalance  PrivateRestAPI = iota //查看账户余额
	PrivateRestAccountPosition                       //查看持仓信息
	PrivateRestAccountConfig                         //查看账户配置

	//Trade
	PrivateRestTradeOrderGet  //查看订单信息
	PrivateRestTradeOrderPost //下单
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	//Account
	PrivateRestAccountBalance:  "/api/v5/account/balance",   //GET 查看账户余额
	PrivateRestAccountPosition: "/api/v5/account/positions", //GET 查看持仓信息
	PrivateRestAccountConfig:   "/api/v5/account/config",    //GET 查看账户配置

	//Trade
	PrivateRestTradeOrderGet:  "/api/v5/trade/order", //GET 查看订单信息
	PrivateRestTradeOrderPost: "/api/v5/trade/order", //POST 下单
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
