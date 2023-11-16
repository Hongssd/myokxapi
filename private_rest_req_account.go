package myokxapi

type PrivateRestAccountBalanceReq struct {
	Ccy *string `json:"ccy"` //String 否 币种 如 BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
}

type PrivateRestAccountBalanceAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountBalanceReq
}

// String 否 币种 如 BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
func (api *PrivateRestAccountBalanceAPI) Ccy(ccy string) *PrivateRestAccountBalanceAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

type PrivateRestAccountPositionReq struct {
	InstType *string `json:"instType"` //String 否 产品类型 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权 instType和instId同时传入的时候会校验instId与instType是否一致。
	InstId   *string `json:"instId"`   //String	否 交易产品ID，如：BTC-USD-190927-5000-C 支持多个instId查询（不超过10个），半角逗号分隔
	PosId    *string `json:"posId"`    //String	否 持仓ID 支持多个posId查询（不超过20个）。 存在有效期的属性，自最近一次完全平仓算起，满30天 posId 以及整个仓位会被清除。
}

type PrivateRestAccountPositionAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountPositionReq
}

// String 否 产品类型 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权 instType和instId同时传入的时候会校验instId与instType是否一致。
func (api *PrivateRestAccountPositionAPI) InstType(instType string) *PrivateRestAccountPositionAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 交易产品ID，如：BTC-USD-190927-5000-C 支持多个instId查询（不超过10个），半角逗号分隔
func (api *PrivateRestAccountPositionAPI) InstId(instId string) *PrivateRestAccountPositionAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 持仓ID 支持多个posId查询（不超过20个）。 存在有效期的属性，自最近一次完全平仓算起，满30天 posId 以及整个仓位会被清除。
func (api *PrivateRestAccountPositionAPI) PosId(posId string) *PrivateRestAccountPositionAPI {
	api.req.PosId = GetPointer(posId)
	return api
}

type PrivateRestAccountConfigReq struct {
}

type PrivateRestAccountConfigAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountConfigReq
}
