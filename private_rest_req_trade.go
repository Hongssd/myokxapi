package myokxapi

type PrivateRestTradeOrderGetReq struct {
	InstId  *string `json:"instId"`  //String	是	产品ID ，如BTC-USD-190927
	OrdId   *string `json:"ordId"`   //String	可选	订单ID ， ordId和clOrdId必须传一个，若传两个，以ordId为主
	ClOrdId *string `json:"clOrdId"` //String	可选	用户自定义ID
}

type PrivateRestTradeOrderGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderGetReq
}

// String 是 产品ID ，如BTC-USD-190927
func (api *PrivateRestTradeOrderGetAPI) InstId(instId string) *PrivateRestTradeOrderGetAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 可选 订单ID ， ordId和clOrdId必须传一个，若传两个，以ordId为主
func (api *PrivateRestTradeOrderGetAPI) OrdId(ordId string) *PrivateRestTradeOrderGetAPI {
	api.req.OrdId = GetPointer(ordId)
	return api
}

// String 可选 用户自定义ID
func (api *PrivateRestTradeOrderGetAPI) ClOrdId(clOrdId string) *PrivateRestTradeOrderGetAPI {
	api.req.ClOrdId = GetPointer(clOrdId)
	return api
}

type PrivateRestTradeOrdersPendingReq struct {
	InstType   *string `json:"instType"`   //String	否	产品类型  SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String	否	标的指数
	InstFamily *string `json:"instFamily"` //String	否	交易品种 适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String	否	产品ID，如BTC-USD-200927
	OrdType    *string `json:"ordType"`    //String	否	订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单)
	State      *string `json:"state"`      //String	否	订单状态 live：等待成交 partially_filled：部分成交
	After      *string `json:"after"`      //String	否	请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
	Before     *string `json:"before"`     //String	否	请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
	Limit      *string `json:"limit"`      //String	否	返回结果的数量，最大为100，默认100条
}

type PrivateRestTradeOrdersPendingAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrdersPendingReq
}

// String 否 产品类型  SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PrivateRestTradeOrdersPendingAPI) InstType(instType string) *PrivateRestTradeOrdersPendingAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 标的指数
func (api *PrivateRestTradeOrdersPendingAPI) Uly(uly string) *PrivateRestTradeOrdersPendingAPI {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种 适用于交割/永续/期权
func (api *PrivateRestTradeOrdersPendingAPI) InstFamily(instFamily string) *PrivateRestTradeOrdersPendingAPI {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如BTC-USD-200927
func (api *PrivateRestTradeOrdersPendingAPI) InstId(instId string) *PrivateRestTradeOrdersPendingAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单)
func (api *PrivateRestTradeOrdersPendingAPI) OrdType(ordType string) *PrivateRestTradeOrdersPendingAPI {
	api.req.OrdType = GetPointer(ordType)
	return api
}

// String 否 订单状态 live：等待成交 partially_filled：部分成交
func (api *PrivateRestTradeOrdersPendingAPI) State(state string) *PrivateRestTradeOrdersPendingAPI {
	api.req.State = GetPointer(state)
	return api
}

// String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
func (api *PrivateRestTradeOrdersPendingAPI) After(after string) *PrivateRestTradeOrdersPendingAPI {
	api.req.After = GetPointer(after)
	return api
}

// String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
func (api *PrivateRestTradeOrdersPendingAPI) Before(before string) *PrivateRestTradeOrdersPendingAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String 否 返回结果的数量，最大为100，默认100条
func (api *PrivateRestTradeOrdersPendingAPI) Limit(limit string) *PrivateRestTradeOrdersPendingAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestTradeOrderPostReq struct {
	InstId         *string                                      `json:"instId,omitempty"`         //String	是	产品ID，如 BTC-USD-190927-5000-C
	TdMode         *string                                      `json:"tdMode,omitempty"`         //String	是	交易模式 保证金模式：isolated：逐仓 ；cross：全仓 非保证金模式：cash：非保证金 spot_isolated：现货逐仓(仅适用于现货带单)
	Ccy            *string                                      `json:"ccy,omitempty"`            //String	否	保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
	ClOrdId        *string                                      `json:"clOrdId,omitempty"`        //String	否	客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag            *string                                      `json:"tag,omitempty"`            //String	否	订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
	Side           *string                                      `json:"side,omitempty"`           //String	是	订单方向 buy：买， sell：卖
	PosSide        *string                                      `json:"posSide,omitempty"`        //String	可选	持仓方向 在开平仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续。
	OrdType        *string                                      `json:"ordType,omitempty"`        //String	是	订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单)
	Sz             *string                                      `json:"sz,omitempty"`             //String	是	委托数量
	Px             *string                                      `json:"px,omitempty"`             //String	可选	委托价格，仅适用于limit、post_only、fok、ioc、mmp、mmp_and_post_only类型的订单 期权下单时，px/pxUsd/pxVol 只能填一个
	PxUsd          *string                                      `json:"pxUsd,omitempty"`          //String	可选	以USD价格进行期权下单 仅适用于期权 期权下单时 px/pxUsd/pxVol 必填一个，且只能填一个
	PxVol          *string                                      `json:"pxVol,omitempty"`          //String	可选	以隐含波动率进行期权下单，例如 1 代表 100% 仅适用于期权 期权下单时 px/pxUsd/pxVol 必填一个，且只能填一个
	ReduceOnly     *bool                                        `json:"reduceOnly,omitempty"`     //Boolean	否	是否只减仓，true 或 false，默认false 仅适用于币币杠杆，以及买卖模式下的交割/永续 仅适用于单币种保证金模式和跨币种保证金模式
	TgtCcy         *string                                      `json:"tgtCcy,omitempty"`         //String	否	市价单委托数量sz的单位，仅适用于币币市价订单 base_ccy: 交易货币 ；quote_ccy：计价货币 买单默认quote_ccy， 卖单默认base_ccy
	BanAmend       *bool                                        `json:"banAmend,omitempty"`       //Boolean	否	是否禁止币币市价改单，true 或 false，默认false 为true时，余额不足时，系统不会改单，下单会失败，仅适用于币币市价单
	QuickMgnType   *string                                      `json:"quickMgnType,omitempty"`   //String	否	一键借币类型，仅适用于杠杆逐仓的一键借币模式： manual：手动，auto_borrow： 自动借币，auto_repay： 自动还币 默认是manual：手动
	StpId          *string                                      `json:"stpId,omitempty"`          //String	否	自成交保护ID。来自同一个母账户配着同一个ID的订单不能自成交 用户自定义1<=x<=999999999的整数
	StpMode        *string                                      `json:"stpMode,omitempty"`        //String	否	自成交保护模式，需要stpId有值才会生效 默认为 cancel maker cancel_maker,cancel_taker, cancel_both Cancel both不支持FOK
	AttachAlgoOrds *[]PrivateRestTradeOrderPostReqAttachAlgoOrd `json:"attachAlgoOrds,omitempty"` //Array of object	否	下单附带止盈止损信息
}

type PrivateRestTradeOrderPostReqAttachAlgoOrd struct {
	AttachAlgoClOrdId    *string `json:"attachAlgoClOrdId,omitempty"`    //String	否	下单附带止盈止损时，客户自定义的策略订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 订单完全成交，下止盈止损委托单时，该值会传给algoClOrdId
	TpTriggerPx          *string `json:"tpTriggerPx,omitempty"`          //String	可选	止盈触发价，如果填写此参数，必须填写 止盈委托价
	TpOrdPx              *string `json:"tpOrdPx,omitempty"`              //String	可选	止盈委托价，如果填写此参数，必须填写 止盈触发价 委托价格为-1时，执行市价止盈
	SlTriggerPx          *string `json:"slTriggerPx,omitempty"`          //String	可选	止损触发价，如果填写此参数，必须填写 止损委托价
	SlOrdPx              *string `json:"slOrdPx,omitempty"`              //String	可选	止损委托价，如果填写此参数，必须填写 止损触发价 委托价格为-1时，执行市价止损
	TpTriggerPxType      *string `json:"tpTriggerPxType,omitempty"`      //String	否	止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	SlTriggerPxType      *string `json:"slTriggerPxType,omitempty"`      //String	否	止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	Sz                   *string `json:"sz,omitempty"`                   //String	可选	数量。仅适用于“多笔止盈”的止盈订单，且对于“多笔止盈”的止盈订单必填
	AmendPxOnTriggerType *string `json:"amendPxOnTriggerType,omitempty"` //String	否	是否启用开仓价止损，仅适用于分批止盈的止损订单，第一笔止盈触发时，止损触发价格是否移动到开仓均价止损 0. 不开启，默认值 1. 开启“开仓价止损”，且止损触发价不能为空
}

type PrivateRestTradeOrderPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderPostReq
}

// String 是 产品ID，如 BTC-USD-190927-5000-C
func (api *PrivateRestTradeOrderPostAPI) InstId(instId string) *PrivateRestTradeOrderPostAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 是 交易模式 保证金模式：isolated：逐仓 ；cross：全仓 非保证金模式：cash：非保证金 spot_isolated：现货逐仓(仅适用于现货带单)
func (api *PrivateRestTradeOrderPostAPI) TdMode(tdMode string) *PrivateRestTradeOrderPostAPI {
	api.req.TdMode = GetPointer(tdMode)
	return api
}

// String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
func (api *PrivateRestTradeOrderPostAPI) Ccy(ccy string) *PrivateRestTradeOrderPostAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// String 否 客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateRestTradeOrderPostAPI) ClOrdId(clOrdId string) *PrivateRestTradeOrderPostAPI {
	api.req.ClOrdId = GetPointer(clOrdId)
	return api
}

// String 否 订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
func (api *PrivateRestTradeOrderPostAPI) Tag(tag string) *PrivateRestTradeOrderPostAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// String 是 订单方向 buy：买， sell：卖
func (api *PrivateRestTradeOrderPostAPI) Side(side string) *PrivateRestTradeOrderPostAPI {
	api.req.Side = GetPointer(side)
	return api
}

// String 可选 持仓方向 在开平仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续。
func (api *PrivateRestTradeOrderPostAPI) PosSide(posSide string) *PrivateRestTradeOrderPostAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}

// String 是 订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单)
func (api *PrivateRestTradeOrderPostAPI) OrdType(ordType string) *PrivateRestTradeOrderPostAPI {
	api.req.OrdType = GetPointer(ordType)
	return api
}

// String 是 委托数量
func (api *PrivateRestTradeOrderPostAPI) Sz(sz string) *PrivateRestTradeOrderPostAPI {
	api.req.Sz = GetPointer(sz)
	return api
}

// String 可选 委托价格，仅适用于limit、post_only、fok、ioc、mmp、mmp_and_post_only类型的订单 期权下单时，px/pxUsd/pxVol 只能填一个
func (api *PrivateRestTradeOrderPostAPI) Px(px string) *PrivateRestTradeOrderPostAPI {
	api.req.Px = GetPointer(px)
	return api
}

// String 可选 以USD价格进行期权下单 仅适用于期权 期权下单时 px/pxUsd/pxVol 必填一个，且只能填一个
func (api *PrivateRestTradeOrderPostAPI) PxUsd(pxUsd string) *PrivateRestTradeOrderPostAPI {
	api.req.PxUsd = GetPointer(pxUsd)
	return api
}

// String 可选 以隐含波动率进行期权下单，例如 1 代表 100% 仅适用于期权 期权下单时 px/pxUsd/pxVol 必填一个，且只能填一个
func (api *PrivateRestTradeOrderPostAPI) PxVol(pxVol string) *PrivateRestTradeOrderPostAPI {
	api.req.PxVol = GetPointer(pxVol)
	return api
}

// Boolean 否 是否只减仓，true 或 false，默认false 仅适用于币币杠杆，以及买卖模式下的交割/永续 仅适用于单币种保证金模式和跨币种保证金模式
func (api *PrivateRestTradeOrderPostAPI) ReduceOnly(reduceOnly bool) *PrivateRestTradeOrderPostAPI {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// String 否 市价单委托数量sz的单位，仅适用于币币市价订单 base_ccy: 交易货币 ；quote_ccy：计价货币 买单默认quote_ccy， 卖单默认base_ccy
func (api *PrivateRestTradeOrderPostAPI) TgtCcy(tgtCcy string) *PrivateRestTradeOrderPostAPI {
	api.req.TgtCcy = GetPointer(tgtCcy)
	return api
}

// Boolean 否 是否禁止币币市价改单，true 或 false，默认false 为true时，余额不足时，系统不会改单，下单会失败，仅适用于币币市价单
func (api *PrivateRestTradeOrderPostAPI) BanAmend(banAmend bool) *PrivateRestTradeOrderPostAPI {
	api.req.BanAmend = GetPointer(banAmend)
	return api
}

// String 否 一键借币类型，仅适用于杠杆逐仓的一键借币模式： manual：手动，auto_borrow： 自动借币，auto_repay： 自动还币 默认是manual：手动
func (api *PrivateRestTradeOrderPostAPI) QuickMgnType(quickMgnType string) *PrivateRestTradeOrderPostAPI {
	api.req.QuickMgnType = GetPointer(quickMgnType)
	return api
}

// String 否 自成交保护ID。来自同一个母账户配着同一个ID的订单不能自成交 用户自定义1<=x<=999999999的整数
func (api *PrivateRestTradeOrderPostAPI) StpId(stpId string) *PrivateRestTradeOrderPostAPI {
	api.req.StpId = GetPointer(stpId)
	return api
}

// String 否 自成交保护模式，需要stpId有值才会生效 默认为 cancel maker cancel_maker,cancel_taker, cancel_both Cancel both不支持FOK
func (api *PrivateRestTradeOrderPostAPI) StpMode(stpMode string) *PrivateRestTradeOrderPostAPI {
	api.req.StpMode = GetPointer(stpMode)
	return api
}

// Array of object 否 下单附带止盈止损信息
func (api *PrivateRestTradeOrderPostAPI) AttachAlgoOrds(attachAlgoOrds []PrivateRestTradeOrderPostReqAttachAlgoOrd) *PrivateRestTradeOrderPostAPI {
	api.req.AttachAlgoOrds = &attachAlgoOrds
	return api
}

// 新建止盈止损信息
func (api *PrivateRestTradeOrderPostAPI) NewAttachAlgoOrd() *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	return &PrivateRestTradeOrderPostReqAttachAlgoOrd{}
}

// String 否 下单附带止盈止损时，客户自定义的策略订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 订单完全成交，下止盈止损委托单时，该值会传给algoClOrdId
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetAttachAlgoClOrdId(attachAlgoClOrdId string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.AttachAlgoClOrdId = GetPointer(attachAlgoClOrdId)
	return algoOrds
}

// String 可选 止盈触发价，如果填写此参数，必须填写 止盈委托价
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetTpTriggerPx(tpTriggerPx string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.TpTriggerPx = GetPointer(tpTriggerPx)
	return algoOrds
}

// String 可选 止盈委托价，如果填写此参数，必须填写 止盈触发价 委托价格为-1时，执行市价止盈
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetTpOrdPx(tpOrdPx string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.TpOrdPx = GetPointer(tpOrdPx)
	return algoOrds
}

// String 可选 止损触发价，如果填写此参数，必须填写 止损委托价
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetSlTriggerPx(slTriggerPx string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.SlTriggerPx = GetPointer(slTriggerPx)
	return algoOrds
}

// String 可选 止损委托价，如果填写此参数，必须填写 止损触发价 委托价格为-1时，执行市价止损
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetSlOrdPx(slOrdPx string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.SlOrdPx = GetPointer(slOrdPx)
	return algoOrds
}

// String 否 止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetTpTriggerPxType(tpTriggerPxType string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.TpTriggerPxType = GetPointer(tpTriggerPxType)
	return algoOrds
}

// String 否 止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetSlTriggerPxType(slTriggerPxType string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.SlTriggerPxType = GetPointer(slTriggerPxType)
	return algoOrds
}

// String 可选 数量。仅适用于“多笔止盈”的止盈订单，且对于“多笔止盈”的止盈订单必填
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetSz(sz string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.Sz = GetPointer(sz)
	return algoOrds
}

// String 否 是否启用开仓价止损，仅适用于分批止盈的止损订单，第一笔止盈触发时，止损触发价格是否移动到开仓均价止损 0. 不开启，默认值 1. 开启“开仓价止损”，且止损触发价不能为空
func (algoOrds *PrivateRestTradeOrderPostReqAttachAlgoOrd) SetAmendPxOnTriggerType(amendPxOnTriggerType string) *PrivateRestTradeOrderPostReqAttachAlgoOrd {
	algoOrds.AmendPxOnTriggerType = GetPointer(amendPxOnTriggerType)
	return algoOrds
}

// instId	String	是	产品ID，如 BTC-USD-190927
// ordId	String	可选	订单ID， ordId和clOrdId必须传一个，若传两个，以ordId为主
// clOrdId	String	可选	用户自定义ID
type PrivateRestTradeCancelOrderReq struct {
	InstId  *string `json:"instId"`  //String	是	产品ID，如 BTC-USD-190927
	OrdId   *string `json:"ordId"`   //String	可选	订单ID， ordId和clOrdId必须传一个，若传两个，以ordId为主
	ClOrdId *string `json:"clOrdId"` //String	可选	用户自定义ID
}

type PrivateRestTradeCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeCancelOrderReq
}

// String 是 产品ID，如 BTC-USD-190927
func (api *PrivateRestTradeCancelOrderAPI) InstId(instId string) *PrivateRestTradeCancelOrderAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 可选 订单ID， ordId和clOrdId必须传一个，若传两个，以ordId为主
func (api *PrivateRestTradeCancelOrderAPI) OrdId(ordId string) *PrivateRestTradeCancelOrderAPI {
	api.req.OrdId = GetPointer(ordId)
	return api
}

// String 可选 用户自定义ID
func (api *PrivateRestTradeCancelOrderAPI) ClOrdId(clOrdId string) *PrivateRestTradeCancelOrderAPI {
	api.req.ClOrdId = GetPointer(clOrdId)
	return api
}

type PrivateRestTradeAmendOrderReq struct {
	InstId         *string                                       `json:"instId,omitempty"`         //String	是	产品ID
	CxlOnFail      *bool                                         `json:"cxlOnFail,omitempty"`      //Boolean	否	false：不自动撤单 true：自动撤单 当订单修改失败时，该订单是否需要自动撤销。默认为false
	OrdId          *string                                       `json:"ordId,omitempty"`          //String	可选	订单ID， ordId和clOrdId必须传一个，若传两个，以ordId为主
	ClOrdId        *string                                       `json:"clOrdId,omitempty"`        //String	可选	用户自定义order ID
	ReqId          *string                                       `json:"reqId,omitempty"`          //String	否	用户自定义修改事件ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	NewSz          *string                                       `json:"newSz,omitempty"`          //String	可选	修改的新数量，对于部分成交订单，该数量应包含已成交数量。
	NewPx          *string                                       `json:"newPx,omitempty"`          //String	可选	修改后的新价格 修改的新价格期权改单时，newPx/newPxUsd/newPxVol 只能填一个，且必须与下单参数保持一致，如下单用px，改单时需使用newPx
	NewPxUsd       *string                                       `json:"newPxUsd,omitempty"`       //String	可选	以USD价格进行期权改单 仅适用于期权，期权改单时，newPx/newPxUsd/newPxVol 只能填一个
	NewPxVol       *string                                       `json:"newPxVol,omitempty"`       //String	可选	以隐含波动率进行期权改单，例如 1 代表 100% 仅适用于期权，期权改单时，newPx/newPxUsd/newPxVol 只能填一个
	AttachAlgoOrds *[]PrivateRestTradeAmendOrderReqAttachAlgoOrd `json:"attachAlgoOrds,omitempty"` //Array of object	否	下单附带止盈止损信息
}

type PrivateRestTradeAmendOrderReqAttachAlgoOrd struct {
	AttachAlgoId         *string `json:"attachAlgoId,omitempty"`         //String	可选	附带止盈止损的订单ID，由系统生成，改单时，可用来标识该笔附带止盈止损订单
	AttachAlgoClOrdId    *string `json:"attachAlgoClOrdId,omitempty"`    //String	可选	下单附带止盈止损时，客户自定义的策略订单ID
	NewTpTriggerPx       *string `json:"newTpTriggerPx,omitempty"`       //String	可选	止盈触发价 如果止盈触发价或者委托价为0，那代表删除止盈。只适用于交割和永续合约。
	NewTpOrdPx           *string `json:"newTpOrdPx,omitempty"`           //String	可选	止盈委托价 委托价格为-1时，执行市价止盈。只适用于交割和永续合约。
	NewSlTriggerPx       *string `json:"newSlTriggerPx,omitempty"`       //String	可选	止损触发价 如果止损触发价或者委托价为0，那代表删除止损。只适用于交割和永续合约。
	NewSlOrdPx           *string `json:"newSlOrdPx,omitempty"`           //String	可选	止损委托价 委托价格为-1时，执行市价止损。 只适用于交割和永续合约。
	NewTpTriggerPxType   *string `json:"newTpTriggerPxType,omitempty"`   //String	可选	止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 只适用于交割/永续 如果要新增止盈，该参数必填
	NewSlTriggerPxType   *string `json:"newSlTriggerPxType,omitempty"`   //String	可选	止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 只适用于交割/永续 如果要新增止损，该参数必填
	Sz                   *string `json:"sz,omitempty"`                   //String	可选	新的张数。仅适用于“多笔止盈”的止盈订单且必填
	AmendPxOnTriggerType *string `json:"amendPxOnTriggerType,omitempty"` //String	否	是否启用开仓价止损，仅适用于分批止盈的止损订单 0. 不开启，默认值 1. 开启“开仓价止损”
}

type PrivateRestTradeAmendOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeAmendOrderReq
}

// String 是 产品ID
func (api *PrivateRestTradeAmendOrderAPI) InstId(instId string) *PrivateRestTradeAmendOrderAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// Boolean 否 false：不自动撤单 true：自动撤单 当订单修改失败时，该订单是否需要自动撤销。默认为false
func (api *PrivateRestTradeAmendOrderAPI) CxlOnFail(cxlOnFail bool) *PrivateRestTradeAmendOrderAPI {
	api.req.CxlOnFail = GetPointer(cxlOnFail)
	return api
}

// String 可选 订单ID， ordId和clOrdId必须传一个，若传两个，以ordId为主
func (api *PrivateRestTradeAmendOrderAPI) OrdId(ordId string) *PrivateRestTradeAmendOrderAPI {
	api.req.OrdId = GetPointer(ordId)
	return api
}

// String 可选 用户自定义order ID
func (api *PrivateRestTradeAmendOrderAPI) ClOrdId(clOrdId string) *PrivateRestTradeAmendOrderAPI {
	api.req.ClOrdId = GetPointer(clOrdId)
	return api
}

// String 否 用户自定义修改事件ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateRestTradeAmendOrderAPI) ReqId(reqId string) *PrivateRestTradeAmendOrderAPI {
	api.req.ReqId = GetPointer(reqId)
	return api
}

// String 可选 修改的新数量，对于部分成交订单，该数量应包含已成交数量。
func (api *PrivateRestTradeAmendOrderAPI) NewSz(newSz string) *PrivateRestTradeAmendOrderAPI {
	api.req.NewSz = GetPointer(newSz)
	return api
}

// String 可选 修改后的新价格 修改的新价格期权改单时，newPx/newPxUsd/newPxVol 只能填一个，且必须与下单参数保持一致，如下单用px，改单时需使用newPx
func (api *PrivateRestTradeAmendOrderAPI) NewPx(newPx string) *PrivateRestTradeAmendOrderAPI {
	api.req.NewPx = GetPointer(newPx)
	return api
}

// String 可选 以USD价格进行期权改单 仅适用于期权，期权改单时，newPx/newPxUsd/newPxVol 只能填一个
func (api *PrivateRestTradeAmendOrderAPI) NewPxUsd(newPxUsd string) *PrivateRestTradeAmendOrderAPI {
	api.req.NewPxUsd = GetPointer(newPxUsd)
	return api
}

// String 可选 以隐含波动率进行期权改单，例如 1 代表 100% 仅适用于期权，期权改单时，newPx/newPxUsd/newPxVol 只能填一个
func (api *PrivateRestTradeAmendOrderAPI) NewPxVol(newPxVol string) *PrivateRestTradeAmendOrderAPI {
	api.req.NewPxVol = GetPointer(newPxVol)
	return api
}

// Array of object 否 下单附带止盈止损信息
func (api *PrivateRestTradeAmendOrderAPI) AttachAlgoOrds(attachAlgoOrds []PrivateRestTradeAmendOrderReqAttachAlgoOrd) *PrivateRestTradeAmendOrderAPI {
	api.req.AttachAlgoOrds = &attachAlgoOrds
	return api
}

// 新建止盈止损信息
func (api *PrivateRestTradeAmendOrderAPI) NewAttachAlgoOrd() *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	return &PrivateRestTradeAmendOrderReqAttachAlgoOrd{}
}

// String 可选 附带止盈止损的订单ID，由系统生成，改单时，可用来标识该笔附带止盈止损订单
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetAttachAlgoId(attachAlgoId string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.AttachAlgoId = GetPointer(attachAlgoId)
	return algoOrds
}

// String 可选 下单附带止盈止损时，客户自定义的策略订单ID
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetAttachAlgoClOrdId(attachAlgoClOrdId string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.AttachAlgoClOrdId = GetPointer(attachAlgoClOrdId)
	return algoOrds
}

// String 可选 止盈触发价 如果止盈触发价或者委托价为0，那代表删除止盈。只适用于交割和永续合约。
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetNewTpTriggerPx(newTpTriggerPx string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.NewTpTriggerPx = GetPointer(newTpTriggerPx)
	return algoOrds
}

// String 可选 止盈委托价 委托价格为-1时，执行市价止盈。只适用于交割和永续合约。
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetNewTpOrdPx(newTpOrdPx string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.NewTpOrdPx = GetPointer(newTpOrdPx)
	return algoOrds
}

// String 可选 止损触发价 如果止损触发价或者委托价为0，那代表删除止损。只适用于交割和永续合约。
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetNewSlTriggerPx(newSlTriggerPx string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.NewSlTriggerPx = GetPointer(newSlTriggerPx)
	return algoOrds
}

// String 可选 止损委托价 委托价格为-1时，执行市价止损。 只适用于交割和永续合约。
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetNewSlOrdPx(newSlOrdPx string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.NewSlOrdPx = GetPointer(newSlOrdPx)
	return algoOrds
}

// String 可选 止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 只适用于交割/永续 如果要新增止盈，该参数必填
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetNewTpTriggerPxType(newTpTriggerPxType string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.NewTpTriggerPxType = GetPointer(newTpTriggerPxType)
	return algoOrds
}

// String 可选 止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 只适用于交割/永续 如果要新增止损，该参数必填
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetNewSlTriggerPxType(newSlTriggerPxType string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.NewSlTriggerPxType = GetPointer(newSlTriggerPxType)
	return algoOrds
}

// String 可选 新的张数。仅适用于“多笔止盈”的止盈订单且必填
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetSz(sz string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.Sz = GetPointer(sz)
	return algoOrds
}

// String 否 是否启用开仓价止损，仅适用于分批止盈的止损订单 0. 不开启，默认值 1. 开启“开仓价止损”
func (algoOrds *PrivateRestTradeAmendOrderReqAttachAlgoOrd) SetAmendPxOnTriggerType(amendPxOnTriggerType string) *PrivateRestTradeAmendOrderReqAttachAlgoOrd {
	algoOrds.AmendPxOnTriggerType = GetPointer(amendPxOnTriggerType)
	return algoOrds
}

type PrivateRestTradeBatchOrdersReq []PrivateRestTradeOrderPostReq

type PrivateRestTradeBatchOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeBatchOrdersReq
}

func (api *PrivateRestTradeBatchOrdersAPI) AddNewOrderReq(orderApi *PrivateRestTradeOrderPostAPI) *PrivateRestTradeBatchOrdersAPI {
	if api.req == nil {
		api.req = &PrivateRestTradeBatchOrdersReq{}
	}
	*api.req = append(*api.req, *orderApi.req)
	return api
}

func (api *PrivateRestTradeBatchOrdersAPI) SetOrderList(orderApiList []PrivateRestTradeOrderPostAPI) *PrivateRestTradeBatchOrdersAPI {
	api.req = &PrivateRestTradeBatchOrdersReq{}
	for _, v := range orderApiList {
		*api.req = append(*api.req, *v.req)
	}
	return api
}

type PrivateRestTradeCancelBatchOrdersReq []PrivateRestTradeCancelOrderReq

type PrivateRestTradeCancelBatchOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeCancelBatchOrdersReq
}

func (api *PrivateRestTradeCancelBatchOrdersAPI) AddNewOrderReq(orderApi *PrivateRestTradeCancelOrderAPI) *PrivateRestTradeCancelBatchOrdersAPI {
	if api.req == nil {
		api.req = &PrivateRestTradeCancelBatchOrdersReq{}
	}
	*api.req = append(*api.req, *orderApi.req)
	return api
}

func (api *PrivateRestTradeCancelBatchOrdersAPI) SetOrderList(orderApiList []PrivateRestTradeCancelOrderAPI) *PrivateRestTradeCancelBatchOrdersAPI {
	api.req = &PrivateRestTradeCancelBatchOrdersReq{}
	for _, v := range orderApiList {
		*api.req = append(*api.req, *v.req)
	}
	return api
}

type PrivateRestTradeAmendBatchOrdersReq []PrivateRestTradeAmendOrderReq

type PrivateRestTradeAmendBatchOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeAmendBatchOrdersReq
}

func (api *PrivateRestTradeAmendBatchOrdersAPI) AddNewOrderReq(orderApi *PrivateRestTradeAmendOrderAPI) *PrivateRestTradeAmendBatchOrdersAPI {
	if api.req == nil {
		api.req = &PrivateRestTradeAmendBatchOrdersReq{}
	}
	*api.req = append(*api.req, *orderApi.req)
	return api
}

func (api *PrivateRestTradeAmendBatchOrdersAPI) SetOrderList(orderApiList []PrivateRestTradeAmendOrderAPI) *PrivateRestTradeAmendBatchOrdersAPI {
	api.req = &PrivateRestTradeAmendBatchOrdersReq{}
	for _, v := range orderApiList {
		*api.req = append(*api.req, *v.req)
	}
	return api
}

type PrivateRestTradeOrderHistoryReq struct {
	InstType   *string `json:"instType"`   //String	是	产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String	否	标的指数
	InstFamily *string `json:"instFamily"` //String	否	交易品种 适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String	否	产品ID，如BTC-USD-190927
	OrdType    *string `json:"ordType"`    //String	否	订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单) op_fok：期权简选（全部成交或立即取消）
	State      *string `json:"state"`      //String	否	订单状态 canceled：撤单成功 filled：完全成交 mmp_canceled：做市商保护机制导致的自动撤单
	Category   *string `json:"category"`   //String	否	订单种类 twap：TWAP自动换币 adl：ADL自动减仓 full_liquidation：强制平仓 partial_liquidation：强制减仓 delivery：交割 ddh：对冲减仓类型订单
	After      *string `json:"after"`      //String	否	请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
	Before     *string `json:"before"`     //String	否	请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
	Begin      *string `json:"begin"`      //String	否	筛选的开始时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
	End        *string `json:"end"`        //String	否	筛选的结束时间戳，Unix 时间戳为毫秒数格式，如 1597027383085
	Limit      *string `json:"limit"`      //String	否	返回结果的数量，最大为100，默认100条
}

type PrivateRestTradeOrderHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderHistoryReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PrivateRestTradeOrderHistoryAPI) InstType(instType string) *PrivateRestTradeOrderHistoryAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 标的指数
func (api *PrivateRestTradeOrderHistoryAPI) Uly(uly string) *PrivateRestTradeOrderHistoryAPI {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种 适用于交割/永续/期权
func (api *PrivateRestTradeOrderHistoryAPI) InstFamily(instFamily string) *PrivateRestTradeOrderHistoryAPI {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如BTC-USD-190927
func (api *PrivateRestTradeOrderHistoryAPI) InstId(instId string) *PrivateRestTradeOrderHistoryAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单) op_fok：期权简选（全部成交或立即取消）
func (api *PrivateRestTradeOrderHistoryAPI) OrdType(ordType string) *PrivateRestTradeOrderHistoryAPI {
	api.req.OrdType = GetPointer(ordType)
	return api
}

// String 否 订单状态 canceled：撤单成功 filled：完全成交 mmp_canceled：做市商保护机制导致的自动撤单
func (api *PrivateRestTradeOrderHistoryAPI) State(state string) *PrivateRestTradeOrderHistoryAPI {
	api.req.State = GetPointer(state)
	return api
}

// String 否 订单种类 twap：TWAP自动换币 adl：ADL自动减仓 full_liquidation：强制平仓 partial_liquidation：强制减仓 delivery：交割 ddh：对冲减仓类型订单
func (api *PrivateRestTradeOrderHistoryAPI) Category(category string) *PrivateRestTradeOrderHistoryAPI {
	api.req.Category = GetPointer(category)
	return api
}

// String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
func (api *PrivateRestTradeOrderHistoryAPI) After(after string) *PrivateRestTradeOrderHistoryAPI {
	api.req.After = GetPointer(after)
	return api
}

// String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
func (api *PrivateRestTradeOrderHistoryAPI) Before(before string) *PrivateRestTradeOrderHistoryAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String 否 筛选的开始时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
func (api *PrivateRestTradeOrderHistoryAPI) Begin(begin string) *PrivateRestTradeOrderHistoryAPI {
	api.req.Begin = GetPointer(begin)
	return api
}

// String 否 筛选的结束时间戳，Unix 时间戳为毫秒数格式，如 1597027383085
func (api *PrivateRestTradeOrderHistoryAPI) End(end string) *PrivateRestTradeOrderHistoryAPI {
	api.req.End = GetPointer(end)
	return api
}

// String 否 返回结果的数量，最大为100，默认100条
func (api *PrivateRestTradeOrderHistoryAPI) Limit(limit string) *PrivateRestTradeOrderHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestTradeOrderHistoryArchiveReq struct {
	InstType   *string `json:"instType"`   //String	是	产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String	否	标的指数
	InstFamily *string `json:"instFamily"` //String	否	交易品种 适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String	否	产品ID，如BTC-USD-190927
	OrdType    *string `json:"ordType"`    //String	否	订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单) op_fok：期权简选（全部成交或立即取消）
	State      *string `json:"state"`      //String	否	订单状态 canceled：撤单成功 filled：完全成交 mmp_canceled：做市商保护机制导致的自动撤单
	Category   *string `json:"category"`   //String	否	订单种类 twap：TWAP自动换币 adl：ADL自动减仓 full_liquidation：强制平仓 partial_liquidation：强制减仓 delivery：交割 ddh：对冲减仓类型订单
	After      *string `json:"after"`      //String	否	请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
	Before     *string `json:"before"`     //String	否	请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
	Begin      *string `json:"begin"`      //String	否	筛选的开始时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
	End        *string `json:"end"`        //String	否	筛选的结束时间戳，Unix 时间戳为毫秒数格式，如 1597027383085
	Limit      *string `json:"limit"`      //String	否	返回结果的数量，最大为100，默认100条
}

type PrivateRestTradeOrderHistoryArchiveAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderHistoryArchiveReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PrivateRestTradeOrderHistoryArchiveAPI) InstType(instType string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 标的指数
func (api *PrivateRestTradeOrderHistoryArchiveAPI) Uly(uly string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种 适用于交割/永续/期权
func (api *PrivateRestTradeOrderHistoryArchiveAPI) InstFamily(instFamily string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如BTC-USD-190927
func (api *PrivateRestTradeOrderHistoryArchiveAPI) InstId(instId string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单) op_fok：期权简选（全部成交或立即取消）
func (api *PrivateRestTradeOrderHistoryArchiveAPI) OrdType(ordType string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.OrdType = GetPointer(ordType)
	return api
}

// String 否 订单状态 canceled：撤单成功 filled：完全成交 mmp_canceled：做市商保护机制导致的自动撤单
func (api *PrivateRestTradeOrderHistoryArchiveAPI) State(state string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.State = GetPointer(state)
	return api
}

// String 否 订单种类 twap：TWAP自动换币 adl：ADL自动减仓 full_liquidation：强制平仓 partial_liquidation：强制减仓 delivery：交割 ddh：对冲减仓类型订单
func (api *PrivateRestTradeOrderHistoryArchiveAPI) Category(category string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.Category = GetPointer(category)
	return api
}

// String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的ordId
func (api *PrivateRestTradeOrderHistoryArchiveAPI) After(after string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.After = GetPointer(after)
	return api
}

// String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的ordId
func (api *PrivateRestTradeOrderHistoryArchiveAPI) Before(before string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String 否 筛选的开始时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
func (api *PrivateRestTradeOrderHistoryArchiveAPI) Begin(begin string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.Begin = GetPointer(begin)
	return api
}

// String 否 筛选的结束时间戳，Unix 时间戳为毫秒数格式，如 1597027383085
func (api *PrivateRestTradeOrderHistoryArchiveAPI) End(end string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.End = GetPointer(end)
	return api
}

// String 否 返回结果的数量，最大为100，默认100条
func (api *PrivateRestTradeOrderHistoryArchiveAPI) Limit(limit string) *PrivateRestTradeOrderHistoryArchiveAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestTradeFillsReq struct {
	InstType   *string `json:"instType"`   //String	是	产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String	否	标的指数
	InstFamily *string `json:"instFamily"` //String	否	交易品种 适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String	否	产品ID，如BTC-USD-190927
	OrdId      *string `json:"ordId"`      //String	否	订单ID
	SubType    *string `json:"subType"`    //String	否	成交类型 1：买入 2：卖出 3：开多 4：开空 5：平多 6：平空 100：强减平多 101：强减平空 102：强减买入 103：强减卖出 104：强平平多 105：强平平空 106：强平买入 107：强平卖出 110：强平换币转入 111：强平换币转出 118：系统换币转入 119：系统换币转出 125：自动减仓平多 126：自动减仓平空 127：自动减仓买入 128：自动减仓卖出 212：一键借币的自动借币 213：一键借币的自动还币 204：大宗交易买 205：大宗交易卖 206：大宗交易开多 207：大宗交易开空 208：大宗交易平多 209：大宗交易平空 270：价差交易买 271：价差交易卖 272：价差交易开多 273：价差交易开空 274：价差交易平多 275：价差交易平空
	After      *string `json:"after"`      //String	否	请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的billId
	Before     *string `json:"before"`     //String	否	请求此ID之后（更新的数据）的分页内容，传的值为对应接口的billId
	Begin      *string `json:"begin"`      //String	否	筛选的开始时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
	End        *string `json:"end"`        //String	否	筛选的结束时间戳，Unix 时间戳为毫秒数格式，如 1597027383085
	Limit      *string `json:"limit"`      //String	否	返回结果的数量，最大为100，默认100条
}

type PrivateRestTradeFillsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeFillsReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PrivateRestTradeFillsAPI) InstType(instType string) *PrivateRestTradeFillsAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 标的指数
func (api *PrivateRestTradeFillsAPI) Uly(uly string) *PrivateRestTradeFillsAPI {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种 适用于交割/永续/期权
func (api *PrivateRestTradeFillsAPI) InstFamily(instFamily string) *PrivateRestTradeFillsAPI {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如BTC-USD-190927
func (api *PrivateRestTradeFillsAPI) InstId(instId string) *PrivateRestTradeFillsAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 否 订单ID
func (api *PrivateRestTradeFillsAPI) OrdId(ordId string) *PrivateRestTradeFillsAPI {
	api.req.OrdId = GetPointer(ordId)
	return api
}

// String 否 成交类型 1：买入 2：卖出 3：开多 4：开空 5：平多 6：平空 100：强减平多 101：强减平空 102：强减买入 103：强减卖出 104：强平平多 105：强平平空 106：强平买入 107：强平卖出 110：强平换币转入 111：强平换币转出 118：系统换币转入 119：系统换币转出 125：自动减仓平多 126：自动减仓平空 127：自动减仓买入 128：自动减仓卖出 212：一键借币的自动借币 213：一键借币的自动还币 204：大宗交易买 205：大宗交易卖 206：大宗交易开多 207：大宗交易开空 208：大宗交易平多 209：大宗交易平空 270：价差交易买 271：价差交易卖 272：价差交易开多 273：价差交易开空 274：价差交易平多 275：价差交易平空
func (api *PrivateRestTradeFillsAPI) SubType(subType string) *PrivateRestTradeFillsAPI {
	api.req.SubType = GetPointer(subType)
	return api
}

// String 否 请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的billId
func (api *PrivateRestTradeFillsAPI) After(after string) *PrivateRestTradeFillsAPI {
	api.req.After = GetPointer(after)
	return api
}

// String 否 请求此ID之后（更新的数据）的分页内容，传的值为对应接口的billId
func (api *PrivateRestTradeFillsAPI) Before(before string) *PrivateRestTradeFillsAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String 否 筛选的开始时间戳，Unix 时间戳为毫秒数格式，如 1597026383085
func (api *PrivateRestTradeFillsAPI) Begin(begin string) *PrivateRestTradeFillsAPI {
	api.req.Begin = GetPointer(begin)
	return api
}

// String 否 筛选的结束时间戳，Unix 时间戳为毫秒数格式，如 1597027383085
func (api *PrivateRestTradeFillsAPI) End(end string) *PrivateRestTradeFillsAPI {
	api.req.End = GetPointer(end)
	return api
}

// String 否 返回结果的数量，最大为100，默认100条
func (api *PrivateRestTradeFillsAPI) Limit(limit string) *PrivateRestTradeFillsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestTradeFillsHistoryAPI struct {
	PrivateRestTradeFillsAPI
}
