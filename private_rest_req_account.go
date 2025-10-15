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

// instType	String	是	产品类型
// SPOT：币币
// MARGIN：币币杠杆
// SWAP：永续合约
// FUTURES：交割合约
// OPTION：期权
// instId	String	否	产品ID，如 BTC-USDT
// 仅适用于instType为币币/币币杠杆
// uly	String	否	标的指数
// 适用于交割/永续/期权，如 BTC-USD
// instFamily	String	否	交易品种
// 适用于交割/永续/期权，如 BTC-USD
type PrivateRestAccountTradeFeeReq struct {
	InstType   *string `json:"instType"`   //String	是	产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId     *string `json:"instId"`     //String	否	产品ID，如 BTC-USDT 仅适用于instType为币币/币币杠杆
	Uly        *string `json:"uly"`        //String	否	标的指数 适用于交割/永续/期权，如 BTC-USD
	InstFamily *string `json:"instFamily"` //String	否	交易品种 适用于交割/永续/期权，如 BTC-USD
}

type PrivateRestAccountTradeFeeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountTradeFeeReq
}

// String	是	产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PrivateRestAccountTradeFeeAPI) InstType(instType string) *PrivateRestAccountTradeFeeAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String	否	产品ID，如 BTC-USDT 仅适用于instType为币币/币币杠杆
func (api *PrivateRestAccountTradeFeeAPI) InstId(instId string) *PrivateRestAccountTradeFeeAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String	否	标的指数 适用于交割/永续/期权，如 BTC-USD
func (api *PrivateRestAccountTradeFeeAPI) Uly(uly string) *PrivateRestAccountTradeFeeAPI {
	api.req.Uly = GetPointer(uly)
	return api
}

// String	否	交易品种 适用于交割/永续/期权，如 BTC-USD
func (api *PrivateRestAccountTradeFeeAPI) InstFamily(instFamily string) *PrivateRestAccountTradeFeeAPI {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

type PrivateRestAccountLeverageInfoReq struct {
	InstId  *string `json:"instId"`  //String	是	产品ID 支持多个instId查询，半角逗号分隔。instId个数不超过20个
	MgnMode *string `json:"mgnMode"` //String	是	保证金模式 isolated：逐仓 cross：全仓
}

type PrivateRestAccountLeverageInfoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountLeverageInfoReq
}

// String	是	产品ID 支持多个instId查询，半角逗号分隔。instId个数不超过20个
func (api *PrivateRestAccountLeverageInfoAPI) InstId(instId string) *PrivateRestAccountLeverageInfoAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String	是	保证金模式 isolated：逐仓 cross：全仓
func (api *PrivateRestAccountLeverageInfoAPI) MgnMode(mgnMode string) *PrivateRestAccountLeverageInfoAPI {
	api.req.MgnMode = GetPointer(mgnMode)
	return api
}

type PrivateRestAccountSetLeverageReq struct {
	InstId  *string `json:"instId"`  //String	可选	产品ID：币对、合约 全仓下，instId和ccy至少要传一个；如果两个都传，默认使用instId
	Ccy     *string `json:"ccy"`     //String	可选	保证金币种 仅适用于跨币种保证金模式和组合保证金模式的全仓 币币杠杆。设置自动借币的杠杆倍数时必填
	Lever   *string `json:"lever"`   //String	是	杠杆倍数
	MgnMode *string `json:"mgnMode"` //String	是	保证金模式 isolated：逐仓 cross：全仓 如果ccy有效传值，该参数值只能为cross。
	PosSide *string `json:"posSide"` //String	可选	持仓方向 long：开平仓模式开多 short：开平仓模式开空 仅适用于逐仓交割/永续 在开平仓模式且保证金模式为逐仓条件下必填
}
type PrivateRestAccountSetLeverageAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountSetLeverageReq
}

// String	可选	产品ID：币对、合约 全仓下，instId和ccy至少要传一个；如果两个都传，默认使用instId
func (api *PrivateRestAccountSetLeverageAPI) InstId(instId string) *PrivateRestAccountSetLeverageAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String	可选	保证金币种 仅适用于跨币种保证金模式和组合保证金模式的全仓 币币杠杆。设置自动借币的杠杆倍数时必填
func (api *PrivateRestAccountSetLeverageAPI) Ccy(ccy string) *PrivateRestAccountSetLeverageAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// String	是	杠杆倍数
func (api *PrivateRestAccountSetLeverageAPI) Lever(lever string) *PrivateRestAccountSetLeverageAPI {
	api.req.Lever = GetPointer(lever)
	return api
}

// String	是	保证金模式 isolated：逐仓 cross：全仓 如果ccy有效传值，该参数值只能为cross。
func (api *PrivateRestAccountSetLeverageAPI) MgnMode(mgnMode string) *PrivateRestAccountSetLeverageAPI {
	api.req.MgnMode = GetPointer(mgnMode)
	return api
}

// String	可选	持仓方向 long：开平仓模式开多 short：开平仓模式开空 仅适用于逐仓交割/永续 在开平仓模式且保证金模式为逐仓条件下必填
func (api *PrivateRestAccountSetLeverageAPI) PosSide(posSide string) *PrivateRestAccountSetLeverageAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}

// posMode	String	是	持仓方式
// long_short_mode：开平仓模式 net_mode：买卖模式
// 仅适用交割/永续
type PrivateRestAccountSetPositionModeReq struct {
	PosMode *string `json:"posMode"` //String	是	持仓方式 long_short_mode：开平仓模式 net_mode：买卖模式 仅适用交割/永续
}

type PrivateRestAccountSetPositionModeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountSetPositionModeReq
}

// String	是	持仓方式 long_short_mode：开平仓模式 net_mode：买卖模式 仅适用交割/永续
func (api *PrivateRestAccountSetPositionModeAPI) PosMode(posMode string) *PrivateRestAccountSetPositionModeAPI {
	api.req.PosMode = GetPointer(posMode)
	return api
}

// acctLv	String	是	账户模式
// 1: 简单交易模式
// 2: 单币种保证金模式
// 3: 跨币种保证金模式
// 4: 组合保证金模式
type PrivateRestAccountSetAccountLevelReq struct {
	AcctLv *string `json:"acctLv"` //String	是	账户模式 1: 简单交易模式 2: 单币种保证金模式 3: 跨币种保证金模式 4: 组合保证金模式
}

type PrivateRestAccountSetAccountLevelAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountSetAccountLevelReq
}

// String	是	账户模式 1: 简单交易模式 2: 单币种保证金模式 3: 跨币种保证金模式 4: 组合保证金模式
func (api *PrivateRestAccountSetAccountLevelAPI) AcctLv(acctLv string) *PrivateRestAccountSetAccountLevelAPI {
	api.req.AcctLv = GetPointer(acctLv)
	return api
}

// instType	String	否	产品类型
// MARGIN：币币杠杆
// SWAP：永续合约
// FUTURES：交割合约
// OPTION：期权
// instId	String	否	交易产品ID，如：BTC-USD-SWAP
// mgnMode	String	否	保证金模式
// cross：全仓，isolated：逐仓
// type	String	否	最近一次平仓的类型
// 1：部分平仓;2：完全平仓;3：强平;4：强减; 5：ADL自动减仓 - 仓位未完全平仓; 6：ADL自动减仓 - 仓位完全平仓
// 状态叠加时，以最新的平仓类型为准状态为准。
// posId	String	否	持仓ID。存在有效期的属性，自最近一次完全平仓算起，满30天 posId 会失效，之后的仓位，会使用新的 posId。
// after	String	否	查询仓位更新 (uTime) 之前的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
// before	String	否	查询仓位更新 (uTime) 之后的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
// limit	String	否	分页返回结果的数量，最大为100，默认100条，uTime 相同的记录均会在当前请求中全部返回

type PrivateRestAccountPositionsHistoryReq struct {
	InstType *string `json:"instType"` //String	否	产品类型 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId   *string `json:"instId"`   //String	否	交易产品ID，如：BTC-USD-SWAP
	MgnMode  *string `json:"mgnMode"`  //String	否	保证金模式 cross：全仓，isolated：逐仓
	Type     *string `json:"type"`     //String	否	最近一次平仓的类型
	PosId    *string `json:"posId"`    //String	否	持仓ID。存在有效期的属性，自最近一次完全平仓算起，满30天 posId 会失效，之后的仓位，会使用新的 posId。
	After    *string `json:"after"`    //String	否	查询仓位更新 (uTime) 之前的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
	Before   *string `json:"before"`   //String	否	查询仓位更新 (uTime) 之后的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
	Limit    *string `json:"limit"`    //String	否	分页返回结果的数量，最大为100，默认100条，uTime 相同的记录均会在当前请求中全部返回
}

type PrivateRestAccountPositionsHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountPositionsHistoryReq
}

// String	否	产品类型 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PrivateRestAccountPositionsHistoryAPI) InstType(instType string) *PrivateRestAccountPositionsHistoryAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String	否	交易产品ID，如：BTC-USD-SWAP
func (api *PrivateRestAccountPositionsHistoryAPI) InstId(instId string) *PrivateRestAccountPositionsHistoryAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String	否	保证金模式 cross：全仓，isolated：逐仓
func (api *PrivateRestAccountPositionsHistoryAPI) MgnMode(mgnMode string) *PrivateRestAccountPositionsHistoryAPI {
	api.req.MgnMode = GetPointer(mgnMode)
	return api
}

// String	否	最近一次平仓的类型
func (api *PrivateRestAccountPositionsHistoryAPI) Type(type_ string) *PrivateRestAccountPositionsHistoryAPI {
	api.req.Type = GetPointer(type_)
	return api
}

// String	否	持仓ID。存在有效期的属性，自最近一次完全平仓算起，满30天 posId 会失效，之后的仓位，会使用新的 posId。
func (api *PrivateRestAccountPositionsHistoryAPI) PosId(posId string) *PrivateRestAccountPositionsHistoryAPI {
	api.req.PosId = GetPointer(posId)
	return api
}

// String	否	查询仓位更新 (uTime) 之前的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
func (api *PrivateRestAccountPositionsHistoryAPI) After(after string) *PrivateRestAccountPositionsHistoryAPI {
	api.req.After = GetPointer(after)
	return api
}

// String	否	查询仓位更新 (uTime) 之后的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
func (api *PrivateRestAccountPositionsHistoryAPI) Before(before string) *PrivateRestAccountPositionsHistoryAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String	否	分页返回结果的数量，最大为100，默认100条，uTime 相同的记录均会在当前请求中全部返回
func (api *PrivateRestAccountPositionsHistoryAPI) Limit(limit string) *PrivateRestAccountPositionsHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// Asset
type PrivateRestAssetBillsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetBillsReq
}
type PrivateRestAssetBillsReq struct {
	Ccy      *string `json:"ccy"`      //String	否	币种
	Type     *string `json:"type"`     //String	否	账单类型
	ClientId *string `json:"clientId"` //String	否	转账或提币的客户自定义ID
	After    *string `json:"after"`    //String	否	查询在此之前的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
	Before   *string `json:"before"`   //String	否	查询在此之后的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
	Limit    *string `json:"limit"`    //String	否	分页返回的结果集数量，最大为 100，不填默认返回 100 条
}

// String	否	币种
func (api *PrivateRestAssetBillsAPI) Ccy(ccy string) *PrivateRestAssetBillsAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// String	否	账单类型
func (api *PrivateRestAssetBillsAPI) Type(type_ string) *PrivateRestAssetBillsAPI {
	api.req.Type = GetPointer(type_)
	return api
}

// String	否	转账或提币的客户自定义ID
func (api *PrivateRestAssetBillsAPI) ClientId(clientId string) *PrivateRestAssetBillsAPI {
	api.req.ClientId = GetPointer(clientId)
	return api
}

// String	否	查询在此之前的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
func (api *PrivateRestAssetBillsAPI) After(after string) *PrivateRestAssetBillsAPI {
	api.req.After = GetPointer(after)
	return api
}

// String	否	查询在此之后的内容，值为时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
func (api *PrivateRestAssetBillsAPI) Before(before string) *PrivateRestAssetBillsAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String	否	分页返回的结果集数量，最大为 100，不填默认返回 100 条
func (api *PrivateRestAssetBillsAPI) Limit(limit string) *PrivateRestAssetBillsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// SubAccount

type PrivateRestSubAccountSetTransferOutReq struct {
	SubAcct     *string `json:"subAcct"`     //String	是	子账户名称，支持设置多个（不超过20个），子账户名称之间半角逗号分隔
	CanTransOut *bool   `json:"canTransOut"` //Boolean	否	是否可以主动转出，默认为true false：不可转出 true：可以转出
}
type PrivateRestSubAccountSetTransferOutAPI struct {
	client *PrivateRestClient
	req    *PrivateRestSubAccountSetTransferOutReq
}

// String	是	子账户名称，支持设置多个（不超过20个），子账户名称之间半角逗号分隔
func (api *PrivateRestSubAccountSetTransferOutAPI) SubAcct(subAcct string) *PrivateRestSubAccountSetTransferOutAPI {
	api.req.SubAcct = GetPointer(subAcct)
	return api
}

// Boolean	否	是否可以主动转出，默认为true false：不可转出 true：可以转出
func (api *PrivateRestSubAccountSetTransferOutAPI) CanTransOut(canTransOut bool) *PrivateRestSubAccountSetTransferOutAPI {
	api.req.CanTransOut = GetPointer(canTransOut)
	return api
}
