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

type PrivateRestTradeOrderAlgoPostReq struct {
	//策略委托基本参数
	InstId        *string `json:"instId,omitempty"`        //String	是	产品ID
	TdMode        *string `json:"tdMode,omitempty"`        //String	是	交易模式 保证金模式：isolated：逐仓 ；cross：全仓 非保证金模式：cash：非保证金 spot_isolated：现货逐仓(仅适用于现货带单)
	Ccy           *string `json:"ccy,omitempty"`           //String	否	保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
	Side          *string `json:"side,omitempty"`          //String	是	订单方向 buy：买， sell：卖
	PosSide       *string `json:"posSide,omitempty"`       //String	可选	持仓方向 在开平仓模式下必填，且仅可选择 long 或 short。
	OrdType       *string `json:"ordType,omitempty"`       //String	是	订单类型 conditional：单向止盈止损 oco：双向止盈止损 trigger：计划委托 move_order_stop：移动止盈止损 twap：时间加权委托
	Sz            *string `json:"sz,omitempty"`            //String	可选	委托数量 sz和closeFraction必填且只能填其一
	Tag           *string `json:"tag,omitempty"`           //String	否	订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间
	TgtCcy        *string `json:"tgtCcy,omitempty"`        //String	否	委托数量的类型 base_ccy: 交易货币 ；quote_ccy：计价货币 买单默认quote_ccy， 卖单默认base_ccy
	AlgoClOrdId   *string `json:"algoClOrdId,omitempty"`   //String	否	客户自定义策略订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
	CloseFraction *string `json:"closeFraction,omitempty"` //String	可选	策略委托触发时，平仓的百分比。1 代表100% 现在系统只支持全部平仓，唯一接受参数为1 对于同一个仓位，仅支持一笔全部平仓的止盈止损挂单 仅适用于交割或永续 当posSide = net时，reduceOnly必须为true 仅适用于止盈止损 ordType = conditional 或 oco 仅适用于止盈止损市价订单 不支持组合保证金模式 sz和closeFraction必填且只能填其一

	//止盈止损策略参数
	PrivateRestTradeOrderAlgoPostReqConditional

	//条件委托策略参数
	PrivateRestTradeOrderAlgoPostReqTrigger

	//移动止盈止损策略参数
	PrivateRestTradeOrderAlgoPostReqMoveOrderStop

	//时间加权策略参数
	PrivateRestTradeOrderAlgoPostReqTwap
}

type PrivateRestTradeOrderAlgoPostReqConditional struct {
	//止盈止损特有参数
	TpTriggerPx     *string `json:"tpTriggerPx,omitempty"`     //String	否	止盈触发价，如果填写此参数，必须填写止盈委托价
	TpTriggerPxType *string `json:"tpTriggerPxType,omitempty"` //String	否	止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	TpOrdPx         *string `json:"tpOrdPx,omitempty"`         //String	否	止盈委托价，如果填写此参数，必须填写止盈触发价 委托价格为-1时，执行市价止盈
	TpOrdKind       *string `json:"tpOrdKind,omitempty"`       //String	否	止盈订单类型 condition: 条件单 limit: 限价单 默认为condition
	SlTriggerPx     *string `json:"slTriggerPx,omitempty"`     //String	否	止损触发价，如果填写此参数，必须填写止损委托价
	SlTriggerPxType *string `json:"slTriggerPxType,omitempty"` //String	否	止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	SlOrdPx         *string `json:"slOrdPx,omitempty"`         //String	否	止损委托价，如果填写此参数，必须填写止损触发价 委托价格为-1时，执行市价止损
	CxlOnClosePos   *bool   `json:"cxlOnClosePos,omitempty"`   //Boolean	否	决定用户所下的止盈止损订单是否与该交易产品对应的仓位关联。若关联，仓位被全平时，该止盈止损订单会被同时撤销；若不关联，仓位被撤销时，该止盈止损订单不受影响。 有效值： true：下单与仓位关联的止盈止损订单 false：下单与仓位不关联的止盈止损订单 默认值为false。若传入true，用户必须同时传入 reduceOnly = true，说明当下单与仓位关联的止盈止损订单时，必须为只减仓。 适用于单币种保证金模式、跨币种保证金模式。
	ReduceOnly      *bool   `json:"reduceOnly,omitempty"`      //Boolean	否	是否只减仓，true 或 false，默认false 仅适用于币币杠杆，以及买卖模式下的交割/永续 仅适用于单币种保证金模式和跨币种保证金模式
}

type PrivateRestTradeOrderAlgoPostReqTrigger struct {
	//条件委托特有参数
	TriggerPx      *string                                                 `json:"triggerPx,omitempty"`      //String	是	计划委托触发价格
	OrderPx        *string                                                 `json:"orderPx,omitempty"`        //String	是	委托价格 委托价格为-1时，执行市价委托
	TriggerPxType  *string                                                 `json:"triggerPxType,omitempty"`  //String	否	计划委托触发价格类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	AttachAlgoOrds *[]PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd `json:"attachAlgoOrds,omitempty"` //Array of object	否	下单附带止盈止损信息
}

type PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd struct {
	//条件委托附带止盈止损
	AttachAlgoClOrdId *string `json:"attachAlgoClOrdId,omitempty"` //String	否	下单附带止盈止损时，客户自定义的策略订单ID，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 订单完全成交，下止盈止损委托单时，该值会传给algoClOrdId。
	TpTriggerPx       *string `json:"tpTriggerPx,omitempty"`       //String	否	止盈触发价，如果填写此参数，必须填写止盈委托价
	TpTriggerPxType   *string `json:"tpTriggerPxType,omitempty"`   //String	否	止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	TpOrdPx           *string `json:"tpOrdPx,omitempty"`           //String	否	止盈委托价，如果填写此参数，必须填写止盈触发价 委托价格为-1时，执行市价止盈
	SlTriggerPx       *string `json:"slTriggerPx,omitempty"`       //String	否	止损触发价，如果填写此参数，必须填写止损委托价
	SlTriggerPxType   *string `json:"slTriggerPxType,omitempty"`   //String	否	止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	SlOrdPx           *string `json:"slOrdPx,omitempty"`           //String	否	止损委托价，如果填写此参数，必须填写止损触发价 委托价格为-1时，执行市价止损
}
type PrivateRestTradeOrderAlgoPostReqMoveOrderStop struct {
	//移动止盈止损特有参数
	CallbackRatio  *string `json:"callbackRatio,omitempty"`  //String	可选	回调幅度的比例，如 "0.05"代表"5%" callbackRatio和callbackSpread只能传入一个
	CallbackSpread *string `json:"callbackSpread,omitempty"` //String	可选	回调幅度的价距
	ActivePx       *string `json:"activePx,omitempty"`       //String	否	激活价格 激活价格是移动止盈止损的激活条件，当市场最新成交价达到或超过激活价格，委托被激活。激活后系统开始计算止盈止损的实际触发价格。如果不填写激活价格，即下单后就被激活。
	ReduceOnly     *bool   `json:"reduceOnly,omitempty"`     //Boolean	否	是否只减仓，true 或 false，默认false 该参数仅在 交割/永续 的买卖模式下有效，开平模式忽略此参数
}

type PrivateRestTradeOrderAlgoPostReqTwap struct {
	//时间加权委托特有参数
	PxVar        *string `json:"pxVar,omitempty"`        //String	可选	吃单价优于盘口的比例，取值范围在 [0.0001,0.01] 之间，如 "0.01"代表"1%" 以买入为例，市价低于限制价时，策略开始用买一价向上取一定比例的委托价来委托小额买单。当前这个参数就用来确定向上的比例。 pxVar和pxSpread只能传入一个
	PxSpread     *string `json:"pxSpread,omitempty"`     //String	可选	吃单单价优于盘口的价距，取值范围不小于0（无上限） 以买入为例，市价低于限制价时，策略开始用买一价向上取一定价距的委托价来委托小额买单。当前这个参数就用来确定向上的价距。
	SzLimit      *string `json:"szLimit,omitempty"`      //String	是	单笔数量 以买入为例，市价低于 “限制价” 时，策略开始用买一价向上取一定价距 / 比例的委托价来委托 “一定数量” 的买单。当前这个参数用来确定其中的 “一定数量”。
	PxLimit      *string `json:"pxLimit,omitempty"`      //String	是	吃单限制价，取值范围不小于0（无上限） 以买入为例，市价低于 “限制价” 时，策略开始用买一价向上取一定价距 / 比例的委托价来委托小额买单。当前这个参数就是其中的 “限制价”。
	TimeInterval *string `json:"timeInterval,omitempty"` //String	是	下单间隔，单位为秒。 以买入为例，市价低于 “限制价” 时，策略开始按 “时间周期” 用买一价向上取一定价距 / 比例的委托价来委托小额买单。当前这个参数就是其中的 “时间周期”。
}

type PrivateRestTradeOrderAlgoPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderAlgoPostReq
}

// String 是 产品ID
func (api *PrivateRestTradeOrderAlgoPostAPI) InstId(instId string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 是 交易模式 保证金模式：isolated：逐仓 ；cross：全仓 非保证金模式：cash：非保证金 spot_isolated：现货逐仓(仅适用于现货带单)
func (api *PrivateRestTradeOrderAlgoPostAPI) TdMode(tdMode string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.TdMode = GetPointer(tdMode)
	return api
}

// String 否 保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
func (api *PrivateRestTradeOrderAlgoPostAPI) Ccy(ccy string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// String 否 订单方向 buy：买， sell：卖
func (api *PrivateRestTradeOrderAlgoPostAPI) Side(side string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.Side = GetPointer(side)
	return api
}

// String 可选 持仓方向 在开平仓模式下必填，且仅可选择 long 或 short。 仅适用交割、永续。
func (api *PrivateRestTradeOrderAlgoPostAPI) PosSide(posSide string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}

// String 是 订单类型 conditional：单向止盈止损 oco：双向止盈止损 trigger：计划委托 move_order_stop：移动止盈止损 twap：时间加权委托
func (api *PrivateRestTradeOrderAlgoPostAPI) OrdType(ordType string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.OrdType = GetPointer(ordType)
	return api
}

// String 可选 委托数量 sz和closeFraction必填且只能填其一
func (api *PrivateRestTradeOrderAlgoPostAPI) Sz(sz string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.Sz = GetPointer(sz)
	return api
}

// String 否 订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间
func (api *PrivateRestTradeOrderAlgoPostAPI) Tag(tag string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// String 否 委托数量的类型 base_ccy: 交易货币 ；quote_ccy：计价货币 买单默认quote_ccy， 卖单默认base_ccy
func (api *PrivateRestTradeOrderAlgoPostAPI) TgtCcy(tgtCcy string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.TgtCcy = GetPointer(tgtCcy)
	return api
}

// String 否 客户自定义策略订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
func (api *PrivateRestTradeOrderAlgoPostAPI) AlgoClOrdId(algoClOrdId string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.AlgoClOrdId = GetPointer(algoClOrdId)
	return api
}

// String 可选 策略委托触发时，平仓的百分比。1 代表100% 现在系统只支持全部平仓，唯一接受参数为1 对于同一个仓位，仅支持一笔全部平仓的止盈止损挂单 仅适用于交割或永续 当posSide = net时，reduceOnly必须为true 仅适用于止盈止损 ordType = conditional 或 oco 仅适用于止盈止损市价订单 不支持组合保证金模式 sz和closeFraction必填且只能填其一
func (api *PrivateRestTradeOrderAlgoPostAPI) CloseFraction(closeFraction string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.CloseFraction = GetPointer(closeFraction)
	return api
}

// 止盈止损特有参数
// String 否 止盈触发价，如果填写此参数，必须填写止盈委托价
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalTpTriggerPx(tpTriggerPx string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.TpTriggerPx = GetPointer(tpTriggerPx)
	return api
}

// String 否 止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalTpTriggerPxType(tpTriggerPxType string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.TpTriggerPxType = GetPointer(tpTriggerPxType)
	return api
}

// String 否 止盈委托价，如果填写此参数，必须填写止盈触发价 委托价格为-1时，执行市价止盈
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalTpOrdPx(tpOrdPx string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.TpOrdPx = GetPointer(tpOrdPx)
	return api
}

// String 否 止盈订单类型 condition: 条件单 limit: 限价单 默认为condition
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalTpOrdKind(tpOrdKind string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.TpOrdKind = GetPointer(tpOrdKind)
	return api
}

// String 否 止损触发价，如果填写此参数，必须填写止损委托价
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalSlTriggerPx(slTriggerPx string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.SlTriggerPx = GetPointer(slTriggerPx)
	return api
}

// String 否 止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalSlTriggerPxType(slTriggerPxType string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.SlTriggerPxType = GetPointer(slTriggerPxType)
	return api
}

// String 否 止损委托价，如果填写此参数，必须填写止损触发价 委托价格为-1时，执行市价止损
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalSlOrdPx(slOrdPx string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.SlOrdPx = GetPointer(slOrdPx)
	return api
}

// Boolean 否 决定用户所下的止盈止损订单是否与该交易产品对应的仓位关联。若关联，仓位被全平时，该止盈止损订单会被同时撤销；若不关联，仓位被撤销时，该止盈止损订单不受影响。 有效值： true：下单与仓位关联的止盈止损订单 false：下单与仓位不关联的止盈止损订单 默认值为false。若传入true，用户必须同时传入 reduceOnly = true，说明当下单与仓位关联的止盈止损订单时，必须为只减仓。 适用于单币种保证金模式、跨币种保证金模式。
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalCxlOnClosePos(cxlOnClosePos bool) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.CxlOnClosePos = GetPointer(cxlOnClosePos)
	return api
}

// Boolean 否 是否只减仓，true 或 false，默认false 仅适用于币币杠杆，以及买卖模式下的交割/永续 仅适用于单币种保证金模式和跨币种保证金模式
func (api *PrivateRestTradeOrderAlgoPostAPI) ConditionalReduceOnly(reduceOnly bool) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqConditional.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// 计划委托特有参数
// String 是 计划委托触发价格
func (api *PrivateRestTradeOrderAlgoPostAPI) TriggerTriggerPx(triggerPx string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTrigger.TriggerPx = GetPointer(triggerPx)
	return api
}

// String 是 委托价格 委托价格为-1时，执行市价委托
func (api *PrivateRestTradeOrderAlgoPostAPI) TriggerOrderPx(orderPx string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTrigger.OrderPx = GetPointer(orderPx)
	return api
}

// String 否 计划委托触发价格类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
func (api *PrivateRestTradeOrderAlgoPostAPI) TriggerTriggerPxType(triggerPxType string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTrigger.TriggerPxType = GetPointer(triggerPxType)
	return api
}

// 下单附带止盈止损信息
func (api *PrivateRestTradeOrderAlgoPostAPI) TriggerAttachAlgoOrds(attachAlgoOrds []PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTrigger.AttachAlgoOrds = &attachAlgoOrds
	return api
}

func (api *PrivateRestTradeOrderAlgoPostAPI) TriggerNewAttachAlgoOrd() *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd {
	attachAlgoOrd := PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd{}
	return &attachAlgoOrd
}

// String 可选 下单附带止盈止损时，客户自定义的策略订单ID，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。 订单完全成交，下止盈止损委托单时，该值会传给algoClOrdId。
func (algoOrds *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd) SetAttachAlgoClOrdId(attachAlgoClOrdId string) *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd {
	algoOrds.AttachAlgoClOrdId = GetPointer(attachAlgoClOrdId)
	return algoOrds
}

// String 可选 止盈触发价，如果填写此参数，必须填写止盈委托价
func (algoOrds *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd) SetTpTriggerPx(tpTriggerPx string) *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd {
	algoOrds.TpTriggerPx = GetPointer(tpTriggerPx)
	return algoOrds
}

// String 可选 止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
func (algoOrds *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd) SetTpTriggerPxType(tpTriggerPxType string) *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd {
	algoOrds.TpTriggerPxType = GetPointer(tpTriggerPxType)
	return algoOrds
}

// String 可选 止盈委托价，如果填写此参数，必须填写止盈触发价 委托价格为-1时，执行市价止盈
func (algoOrds *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd) SetTpOrdPx(tpOrdPx string) *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd {
	algoOrds.TpOrdPx = GetPointer(tpOrdPx)
	return algoOrds
}

// String 可选 止损触发价，如果填写此参数，必须填写止损委托价
func (algoOrds *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd) SetSlTriggerPx(slTriggerPx string) *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd {
	algoOrds.SlTriggerPx = GetPointer(slTriggerPx)
	return algoOrds
}

// String 可选 止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
func (algoOrds *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd) SetSlTriggerPxType(slTriggerPxType string) *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd {
	algoOrds.SlTriggerPxType = GetPointer(slTriggerPxType)
	return algoOrds
}

// String 可选 止损委托价，如果填写此参数，必须填写止损触发价 委托价格为-1时，执行市价止损
func (algoOrds *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd) SetSlOrdPx(slOrdPx string) *PrivateRestTradeOrderAlgoPostReqTriggerAttachAlgoOrd {
	algoOrds.SlOrdPx = GetPointer(slOrdPx)
	return algoOrds
}

// 移动止盈止损特有参数
// String 可选 回调幅度的比例，如 "0.05"代表"5%" callbackRatio和callbackSpread只能传入一个
func (api *PrivateRestTradeOrderAlgoPostAPI) MoveOrderStopCallbackRatio(callbackRatio string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqMoveOrderStop.CallbackRatio = GetPointer(callbackRatio)
	return api
}

// String 可选 回调幅度的价距，取值范围不小于0（无上限） 以买入为例，市价低于限制价时，策略开始用买一价向上取一定价距的委托价来委托小额买单。当前这个参数就用来确定向上的价距。
func (api *PrivateRestTradeOrderAlgoPostAPI) MoveOrderStopCallbackSpread(callbackSpread string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqMoveOrderStop.CallbackSpread = GetPointer(callbackSpread)
	return api
}

// String 可选 激活价格 激活价格是移动止盈止损的激活条件，当市场最新成交价达到或超过激活价格，委托被激活。激活后系统开始计算止盈止损的实际触发价格。如果不填写激活价格，即下单后就被激活。
func (api *PrivateRestTradeOrderAlgoPostAPI) MoveOrderStopActivePx(activePx string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqMoveOrderStop.ActivePx = GetPointer(activePx)
	return api
}

// Boolean 否 是否只减仓，true 或 false，默认false 该参数仅在 交割/永续 的买卖模式下有效，开平模式忽略此参数
func (api *PrivateRestTradeOrderAlgoPostAPI) MoveOrderStopReduceOnly(reduceOnly bool) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqMoveOrderStop.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// 时间加权委托特有参数
// String 可选 吃单价优于盘口的比例，取值范围在 [0.0001,0.01] 之间，如 "0.01"代表"1%" 以买入为例，市价低于限制价时，策略开始用买一价向上取一定比例的委托价来委托小额买单。当前这个参数就用来确定向上的比例。 pxVar和pxSpread只能传入一个
func (api *PrivateRestTradeOrderAlgoPostAPI) TwapPxVar(pxVar string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTwap.PxVar = GetPointer(pxVar)
	return api
}

// String 可选 吃单单价优于盘口的价距，取值范围不小于0（无上限） 以买入为例，市价低于限制价时，策略开始用买一价向上取一定价距的委托价来委托小额买单。当前这个参数就用来确定向上的价距。
func (api *PrivateRestTradeOrderAlgoPostAPI) TwapPxSpread(pxSpread string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTwap.PxSpread = GetPointer(pxSpread)
	return api
}

// String 是 单笔数量 以买入为例，市价低于 “限制价” 时，策略开始用买一价向上取一定价距 / 比例的委托价来委托 “一定数量” 的买单。当前这个参数用来确定其中的 “一定数量”。
func (api *PrivateRestTradeOrderAlgoPostAPI) TwapSzLimit(szLimit string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTwap.SzLimit = GetPointer(szLimit)
	return api
}

// String 是 吃单限制价，取值范围不小于0（无上限） 以买入为例，市价低于 “限制价” 时，策略开始用买一价向上取一定价距 / 比例的委托价来委托小额买单。当前这个参数就是其中的 “限制价”。
func (api *PrivateRestTradeOrderAlgoPostAPI) TwapPxLimit(pxLimit string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTwap.PxLimit = GetPointer(pxLimit)
	return api
}

// String 是 下单间隔，单位为秒。 以买入为例，市价低于 “限制价” 时，策略开始按 “时间周期” 用买一价向上取一定价距 / 比例的委托价来委托小额买单。当前这个参数就是其中的 “时间周期”。
func (api *PrivateRestTradeOrderAlgoPostAPI) TwapTimeInterval(timeInterval string) *PrivateRestTradeOrderAlgoPostAPI {
	api.req.PrivateRestTradeOrderAlgoPostReqTwap.TimeInterval = GetPointer(timeInterval)
	return api
}

type PrivateRestTradeCancelOrderAlgoReq []PrivateRestTradeCancelOrderAlgoReqRow

type PrivateRestTradeCancelOrderAlgoReqRow struct {
	AlgoId string `json:"algoId"` //String	是	策略委托ID
	InstId string `json:"instId"` //String	是	产品ID 如 BTC-USDT
}

type PrivateRestTradeCancelOrderAlgoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeCancelOrderAlgoReq
}

func (api *PrivateRestTradeCancelOrderAlgoAPI) AddCancelOrderAlgo(row ...*PrivateRestTradeCancelOrderAlgoReqRow) *PrivateRestTradeCancelOrderAlgoAPI {
	if api.req == nil {
		api.req = &PrivateRestTradeCancelOrderAlgoReq{}
	}

	for _, r := range row {
		*api.req = append(*api.req, *r)
	}

	return api
}

func (api *PrivateRestTradeCancelOrderAlgoAPI) NewCancelOrderAlgo() *PrivateRestTradeCancelOrderAlgoReqRow {
	return &PrivateRestTradeCancelOrderAlgoReqRow{}
}

// String	是	策略委托ID
func (row *PrivateRestTradeCancelOrderAlgoReqRow) SetAlgoId(algoId string) *PrivateRestTradeCancelOrderAlgoReqRow {
	row.AlgoId = algoId
	return row
}

// String	是	产品ID 如 BTC-USDT
func (row *PrivateRestTradeCancelOrderAlgoReqRow) SetInstId(instId string) *PrivateRestTradeCancelOrderAlgoReqRow {
	row.InstId = instId
	return row
}

type PrivateRestTradeAmendOrderAlgoAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeAmendOrderAlgoReq
}

type PrivateRestTradeAmendOrderAlgoReq struct {
	InstId             *string `json:"instId,omitempty"`             //String	是	产品ID
	AlgoId             *string `json:"algoId,omitempty"`             //String	可选	策略委托单ID algoId和algoClOrdId必须传一个，若传两个，以algoId为主
	AlgoClOrdId        *string `json:"algoClOrdId,omitempty"`        //String	可选	客户自定义策略订单ID algoId和algoClOrdId必须传一个，若传两个，以algoId为主
	CxlOnFail          *bool   `json:"cxlOnFail,omitempty"`          //Boolean	否	当订单修改失败时，该订单是否需要自动撤销。默认为false
	ReqId              *string `json:"reqId,omitempty"`              //String	否	用户自定义修改事件ID，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
	NewSz              *string `json:"newSz,omitempty"`              //String	可选	修改的新数量，必须大于0。
	NewTpTriggerPx     *string `json:"newTpTriggerPx,omitempty"`     //String	可选	止盈触发价 如果止盈触发价或者委托价为0，那代表删除止盈。只适用于交割和永续合约。
	NewTpOrdPx         *string `json:"newTpOrdPx,omitempty"`         //String	可选	止盈委托价 委托价格为-1时，执行市价止盈
	NewSlTriggerPx     *string `json:"newSlTriggerPx,omitempty"`     //String	可选	止损触发价 如果止损触发价或者委托价为0，那代表删除止损。只适用于交割和永续合约。
	NewSlOrdPx         *string `json:"newSlOrdPx,omitempty"`         //String	可选	止损委托价 委托价格为-1时，执行市价止损
	NewTpTriggerPxType *string `json:"newTpTriggerPxType,omitempty"` //String	可选	止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
	NewSlTriggerPxType *string `json:"newSlTriggerPxType,omitempty"` //String	可选	止损触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
}

// instid String 是 产品ID
func (api *PrivateRestTradeAmendOrderAlgoAPI) InstId(instId string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// algoId String 可选 策略委托单ID algoId和algoClOrdId必须传一个，若传两个，以algoId为主
func (api *PrivateRestTradeAmendOrderAlgoAPI) AlgoId(algoId string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.AlgoId = GetPointer(algoId)
	return api
}

// algoClOrdId String 可选 客户自定义策略订单ID algoId和algoClOrdId必须传一个，若传两个，以algoId为主
func (api *PrivateRestTradeAmendOrderAlgoAPI) AlgoClOrdId(algoClOrdId string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.AlgoClOrdId = GetPointer(algoClOrdId)
	return api
}

// cxlOnFail Boolean 否 当订单修改失败时，该订单是否需要自动撤销。默认为false
func (api *PrivateRestTradeAmendOrderAlgoAPI) CxlOnFail(cxlOnFail bool) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.CxlOnFail = GetPointer(cxlOnFail)
	return api
}

// reqId String 否 用户自定义修改事件ID，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
func (api *PrivateRestTradeAmendOrderAlgoAPI) ReqId(reqId string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.ReqId = GetPointer(reqId)
	return api
}

// newSz String 可选 修改的新数量，必须大于0。
func (api *PrivateRestTradeAmendOrderAlgoAPI) NewSz(newSz string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.NewSz = GetPointer(newSz)
	return api
}

// newTpTriggerPx String 可选 止盈触发价 如果止盈触发价或者委托价为0，那代表删除止盈。只适用于交割和永续合约。
func (api *PrivateRestTradeAmendOrderAlgoAPI) NewTpTriggerPx(newTpTriggerPx string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.NewTpTriggerPx = GetPointer(newTpTriggerPx)
	return api
}

// newTpOrdPx String 可选 止盈委托价 委托价格为-1时，执行市价止盈
func (api *PrivateRestTradeAmendOrderAlgoAPI) NewTpOrdPx(newTpOrdPx string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.NewTpOrdPx = GetPointer(newTpOrdPx)
	return api
}

// newSlTriggerPx String 可选 止损触发价 如果止损触发价或者委托价为0，那代表删除止损。只适用于交割和永续合约。
func (api *PrivateRestTradeAmendOrderAlgoAPI) NewSlTriggerPx(newSlTriggerPx string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.NewSlTriggerPx = GetPointer(newSlTriggerPx)
	return api
}

// newSlOrdPx String 可选 止损委托价 委托价格为-1时，执行市价止损
func (api *PrivateRestTradeAmendOrderAlgoAPI) NewSlOrdPx(newSlOrdPx string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.NewSlOrdPx = GetPointer(newSlOrdPx)
	return api
}

// newTpTriggerPxType String 可选 止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格 默认为last
func (api *PrivateRestTradeAmendOrderAlgoAPI) NewTpTriggerPxType(newTpTriggerPxType string) *PrivateRestTradeAmendOrderAlgoAPI {
	api.req.NewTpTriggerPxType = GetPointer(newTpTriggerPxType)
	return api
}

type PrivateRestTradeOrderAlgoGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderAlgoGetReq
}

// algoId和algoClOrdId必须传一个，若传两个，以algoId为主
// algoClOrdId	String	可选	客户自定义策略订单ID
// 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
type PrivateRestTradeOrderAlgoGetReq struct {
	AlgoId      *string `json:"algoId"`      //String	可选	策略委托单ID
	AlgoClOrdId *string `json:"algoClOrdId"` //String	可选	客户自定义策略订单ID
}

func (api *PrivateRestTradeOrderAlgoGetAPI) AlgoId(algoId string) *PrivateRestTradeOrderAlgoGetAPI {
	api.req.AlgoId = GetPointer(algoId)
	return api
}
func (api *PrivateRestTradeOrderAlgoGetAPI) AlgoClOrdId(algoClOrdId string) *PrivateRestTradeOrderAlgoGetAPI {
	api.req.AlgoClOrdId = GetPointer(algoClOrdId)
	return api
}

type PrivateRestTradeOrderAlgoPendingAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderAlgoPendingReq
}
type PrivateRestTradeOrderAlgoPendingReq struct {
	AlgoId      *string `json:"algoId"`      //String	否	策略委托单ID
	InstType    *string `json:"instType"`    //String	否	产品类型
	InstId      *string `json:"instId"`      //String	否	产品ID，如 BTC-USDT
	OrdType     *string `json:"ordType"`     //String	是	订单类型
	After       *string `json:"after"`       //String	否	请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的algoId
	Before      *string `json:"before"`      //String	否	请求此ID之后（更新的数据）的分页内容，传的值为对应接口的algoId
	Limit       *string `json:"limit"`       //String	否	返回结果的数量，最大为100，默认100条
	AlgoClOrdId *string `json:"algoClOrdId"` //String	否	客户自定义策略订单ID
}

// String	否	策略委托单ID
func (api *PrivateRestTradeOrderAlgoPendingAPI) AlgoId(algoId string) *PrivateRestTradeOrderAlgoPendingAPI {
	api.req.AlgoId = GetPointer(algoId)
	return api
}

// String	否	产品类型
func (api *PrivateRestTradeOrderAlgoPendingAPI) InstType(instType string) *PrivateRestTradeOrderAlgoPendingAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String	否	产品ID，如 BTC-USDT
func (api *PrivateRestTradeOrderAlgoPendingAPI) InstId(instId string) *PrivateRestTradeOrderAlgoPendingAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String	是	订单类型
func (api *PrivateRestTradeOrderAlgoPendingAPI) OrdType(ordType string) *PrivateRestTradeOrderAlgoPendingAPI {
	api.req.OrdType = GetPointer(ordType)
	return api
}

// String	否	请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的algoId
func (api *PrivateRestTradeOrderAlgoPendingAPI) After(after string) *PrivateRestTradeOrderAlgoPendingAPI {
	api.req.After = GetPointer(after)
	return api
}

// String	否	请求此ID之后（更新的数据）的分页内容，传的值为对应接口的algoId
func (api *PrivateRestTradeOrderAlgoPendingAPI) Before(before string) *PrivateRestTradeOrderAlgoPendingAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String	否	返回结果的数量，最大为100，默认100条
func (api *PrivateRestTradeOrderAlgoPendingAPI) Limit(limit string) *PrivateRestTradeOrderAlgoPendingAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// String	否	客户自定义策略订单ID
func (api *PrivateRestTradeOrderAlgoPendingAPI) AlgoClOrdId(algoClOrdId string) *PrivateRestTradeOrderAlgoPendingAPI {
	api.req.AlgoClOrdId = GetPointer(algoClOrdId)
	return api
}

type PrivateRestTradeOrderAlgoHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderAlgoHistoryReq
}
type PrivateRestTradeOrderAlgoHistoryReq struct {
	OrdType  *string `json:"ordType"`  //String	是	订单类型
	State    *string `json:"state"`    //String	可选	订单状态
	AlgoId   *string `json:"algoId"`   //String	可选	策略委托单ID
	InstType *string `json:"instType"` //String	否	产品类型
	InstId   *string `json:"instId"`   //String	否	产品ID，BTC-USDT
	After    *string `json:"after"`    //String	否	请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的algoId
	Before   *string `json:"before"`   //String	否	请求此ID之后（更新的数据）的分页内容，传的值为对应接口的algoId
	Limit    *string `json:"limit"`    //String	否	返回结果的数量，最大为100，默认100条
}

// String	是	订单类型
func (api *PrivateRestTradeOrderAlgoHistoryAPI) OrdType(ordType string) *PrivateRestTradeOrderAlgoHistoryAPI {
	api.req.OrdType = GetPointer(ordType)
	return api
}

// String	可选	订单状态
func (api *PrivateRestTradeOrderAlgoHistoryAPI) State(state string) *PrivateRestTradeOrderAlgoHistoryAPI {
	api.req.State = GetPointer(state)
	return api
}

// String	可选	策略委托单ID
func (api *PrivateRestTradeOrderAlgoHistoryAPI) AlgoId(algoId string) *PrivateRestTradeOrderAlgoHistoryAPI {
	api.req.AlgoId = GetPointer(algoId)
	return api
}

// String	否	产品类型
func (api *PrivateRestTradeOrderAlgoHistoryAPI) InstType(instType string) *PrivateRestTradeOrderAlgoHistoryAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String	否	产品ID，BTC-USDT
func (api *PrivateRestTradeOrderAlgoHistoryAPI) InstId(instId string) *PrivateRestTradeOrderAlgoHistoryAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String	否	请求此ID之前（更旧的数据）的分页内容，传的值为对应接口的algoId
func (api *PrivateRestTradeOrderAlgoHistoryAPI) After(after string) *PrivateRestTradeOrderAlgoHistoryAPI {
	api.req.After = GetPointer(after)
	return api
}

// String	否	请求此ID之后（更新的数据）的分页内容，传的值为对应接口的algoId
func (api *PrivateRestTradeOrderAlgoHistoryAPI) Before(before string) *PrivateRestTradeOrderAlgoHistoryAPI {
	api.req.Before = GetPointer(before)
	return api
}

// String	否	返回结果的数量，最大为100，默认100条
func (api *PrivateRestTradeOrderAlgoHistoryAPI) Limit(limit string) *PrivateRestTradeOrderAlgoHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
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
	orderApi.Tag(BrokerCode)
	*api.req = append(*api.req, *orderApi.req)
	return api
}

func (api *PrivateRestTradeBatchOrdersAPI) SetOrderList(orderApiList []PrivateRestTradeOrderPostAPI) *PrivateRestTradeBatchOrdersAPI {
	api.req = &PrivateRestTradeBatchOrdersReq{}
	for _, v := range orderApiList {
		v.Tag(BrokerCode)
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

type PrivateRestTradeClosePostionReq struct {
	InstId  *string `json:"instId"`  //String	是	产品ID
	PosSide *string `json:"posSide"` //String	可选	持仓方向
	MgnMode *string `json:"mgnMode"` //String	是	保证金模式 cross：全仓 ； isolated：逐仓
	Ccy     *string `json:"ccy"`     //String	可选	保证金币种，现货和合约模式下的全仓币币杠杆平仓必填
	AutoCxl *bool   `json:"autoCxl"` //Boolean	否	当市价全平时，平仓单是否需要自动撤销,默认为false. false：不自动撤单 true：自动撤单
	ClOrdId *string `json:"clOrdId"` //String	否	客户自定义ID
	Tag     *string `json:"tag"`     //String	否	订单标签
}

type PrivateRestTradeClosePostionAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeClosePostionReq
}

// String 是 产品ID
func (api *PrivateRestTradeClosePostionAPI) InstId(instId string) *PrivateRestTradeClosePostionAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String 可选 持仓方向
func (api *PrivateRestTradeClosePostionAPI) PosSide(posSide string) *PrivateRestTradeClosePostionAPI {
	api.req.PosSide = GetPointer(posSide)
	return api
}

// String 是 保证金模式 cross：全仓 ； isolated：逐仓
func (api *PrivateRestTradeClosePostionAPI) MgnMode(mgnMode string) *PrivateRestTradeClosePostionAPI {
	api.req.MgnMode = GetPointer(mgnMode)
	return api
}

// String 可选 保证金币种，现货和合约模式下的全仓币币杠杆平仓必填
func (api *PrivateRestTradeClosePostionAPI) Ccy(ccy string) *PrivateRestTradeClosePostionAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// Boolean 否 当市价全平时，平仓单是否需要自动撤销,默认为false. false：不自动撤单 true：自动撤单
func (api *PrivateRestTradeClosePostionAPI) AutoCxl(autoCxl bool) *PrivateRestTradeClosePostionAPI {
	api.req.AutoCxl = GetPointer(autoCxl)
	return api
}

// String 否 客户自定义ID
func (api *PrivateRestTradeClosePostionAPI) ClOrdId(clOrdId string) *PrivateRestTradeClosePostionAPI {
	api.req.ClOrdId = GetPointer(clOrdId)
	return api
}

// String 否 订单标签
func (api *PrivateRestTradeClosePostionAPI) Tag(tag string) *PrivateRestTradeClosePostionAPI {
	api.req.Tag = GetPointer(tag)
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

// Asset
// 获取币种列表
type PrivateRestAssetCurrenciesReq struct {
	Ccy *string `json:"ccy"` //String	否	币种，如BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
}
type PrivateRestAssetCurrenciesAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetCurrenciesReq
}

// String	否	币种，如BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
func (api *PrivateRestAssetCurrenciesAPI) Ccy(ccy string) *PrivateRestAssetCurrenciesAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// 获取资金账户余额
type PrivateRestAssetBalancesReq struct {
	Ccy *string `json:"ccy"` //String	否	币种，如BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
}
type PrivateRestAssetBalancesAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetBalancesReq
}

// String	否	币种，如BTC 支持多币种查询（不超过20个），币种之间半角逗号分隔
func (api *PrivateRestAssetBalancesAPI) Ccy(ccy string) *PrivateRestAssetBalancesAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// 资金划转
type PrivateRestAssetTransferReq struct {
	Type        *string `json:"type"`        //String	否	划转类型 0：账户内划转 1：母账户转子账户(仅适用于母账户APIKey) 2：子账户转母账户(仅适用于母账户APIKey) 3：子账户转母账户(仅适用于子账户APIKey) 4：子账户转子账户(仅适用于子账户APIKey，且目标账户需要是同一母账户下的其他子账户。子账户主动转出权限默认是关闭的，权限调整参考 设置子账户主动转出权限。) 默认是0 如果您希望通过母账户API Key控制子账户之间的划转，参考接口 子账户间资金划转
	Ccy         *string `json:"ccy"`         //String	是	划转币种，如 USDT
	Amt         *string `json:"amt"`         //String	是	划转数量
	From        *string `json:"from"`        //String	是	转出账户 6：资金账户 18：交易账户
	To          *string `json:"to"`          //String	是	转入账户 6：资金账户 18：交易账户
	SubAcct     *string `json:"subAcct"`     //String	可选	子账户名称 当type为1/2/4时，该字段必填
	LoanTrans   *bool   `json:"loanTrans"`   //Boolean	否	是否支持跨币种保证金模式或组合保证金模式下的借币转出 true：支持借币转出 false：不支持借币转出 默认为false
	OmitPosRisk *string `json:"omitPosRisk"` //String	否	是否忽略仓位风险 默认为false 仅适用于组合保证金模式
	ClientId    *string `json:"clientId"`    //String	否	客户自定义ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
}

type PrivateRestAssetTransferAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetTransferReq
}

// String	否	划转类型 0：账户内划转 1：母账户转子账户(仅适用于母账户APIKey) 2：子账户转母账户(仅适用于母账户APIKey) 3：子账户转母账户(仅适用于子账户APIKey) 4：子账户转子账户(仅适用于子账户APIKey，且目标账户需要是同一母账户下的其他子账户。子账户主动转出权限默认是关闭的，权限调整参考 设置子账户主动转出权限。) 默认是0 如果您希望通过母账户API Key控制子账户之间的划转，参考接口 子账户间资金划转
func (api *PrivateRestAssetTransferAPI) Type(t string) *PrivateRestAssetTransferAPI {
	api.req.Type = GetPointer(t)
	return api
}

// String	是	划转币种，如 USDT
func (api *PrivateRestAssetTransferAPI) Ccy(ccy string) *PrivateRestAssetTransferAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// String	是	划转数量
func (api *PrivateRestAssetTransferAPI) Amt(amt string) *PrivateRestAssetTransferAPI {
	api.req.Amt = GetPointer(amt)
	return api
}

// String	是	转出账户 6：资金账户 18：交易账户
func (api *PrivateRestAssetTransferAPI) From(from string) *PrivateRestAssetTransferAPI {
	api.req.From = GetPointer(from)
	return api
}

// String	是	转入账户 6：资金账户 18：交易账户
func (api *PrivateRestAssetTransferAPI) To(to string) *PrivateRestAssetTransferAPI {
	api.req.To = GetPointer(to)
	return api
}

// String	可选	子账户名称 当type为1/2/4时，该字段必填
func (api *PrivateRestAssetTransferAPI) SubAcct(subAcct string) *PrivateRestAssetTransferAPI {
	api.req.SubAcct = GetPointer(subAcct)
	return api
}

// Boolean	否	是否支持跨币种保证金模式或组合保证金模式下的借币转出 true：支持借币转出 false：不支持借币转出 默认为false
func (api *PrivateRestAssetTransferAPI) LoanTrans(loanTrans bool) *PrivateRestAssetTransferAPI {
	api.req.LoanTrans = GetPointer(loanTrans)
	return api
}

// String	否	是否忽略仓位风险 默认为false 仅适用于组合保证金模式
func (api *PrivateRestAssetTransferAPI) OmitPosRisk(omitPosRisk string) *PrivateRestAssetTransferAPI {
	api.req.OmitPosRisk = GetPointer(omitPosRisk)
	return api
}

// String	否	客户自定义ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateRestAssetTransferAPI) ClientId(clientId string) *PrivateRestAssetTransferAPI {
	api.req.ClientId = GetPointer(clientId)
	return api
}

// 闪兑预估询价
type PrivateRestAssetConvertEstimateQuoteReq struct {
	BaseCcy  *string `json:"baseCcy"`  //String	是	交易货币币种，如 BTC-USDT中的BTC
	QuoteCcy *string `json:"quoteCcy"` //String	是	计价货币币种，如 BTC-USDT中的USDT
	Side     *string `json:"side"`     //String	是	交易方向 买：buy 卖：sell 描述的是对于baseCcy的交易方向
	RfqSz    *string `json:"rfqSz"`    //String	是	询价数量
	RfqSzCcy *string `json:"rfqSzCcy"` //String	是	询价币种
	ClQReqId *string `json:"clQReqId"` //String	否	客户端自定义的订单标识 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag      *string `json:"tag"`      //String	否	订单标签 适用于broker用户
}
type PrivateRestAssetConvertEstimateQuoteAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetConvertEstimateQuoteReq
}

// String	是	交易货币币种，如 BTC-USDT中的BTC
func (api *PrivateRestAssetConvertEstimateQuoteAPI) BaseCcy(baseCcy string) *PrivateRestAssetConvertEstimateQuoteAPI {
	api.req.BaseCcy = GetPointer(baseCcy)
	return api
}

// String	是	计价货币币种，如 BTC-USDT中的USDT
func (api *PrivateRestAssetConvertEstimateQuoteAPI) QuoteCcy(quoteCcy string) *PrivateRestAssetConvertEstimateQuoteAPI {
	api.req.QuoteCcy = GetPointer(quoteCcy)
	return api
}

// String	是	交易方向 买：buy 卖：sell 描述的是对于baseCcy的交易方向
func (api *PrivateRestAssetConvertEstimateQuoteAPI) Side(side string) *PrivateRestAssetConvertEstimateQuoteAPI {
	api.req.Side = GetPointer(side)
	return api
}

// String	是	询价数量
func (api *PrivateRestAssetConvertEstimateQuoteAPI) RfqSz(rfqSz string) *PrivateRestAssetConvertEstimateQuoteAPI {
	api.req.RfqSz = GetPointer(rfqSz)
	return api
}

// String	是	询价币种
func (api *PrivateRestAssetConvertEstimateQuoteAPI) RfqSzCcy(rfqSzCcy string) *PrivateRestAssetConvertEstimateQuoteAPI {
	api.req.RfqSzCcy = GetPointer(rfqSzCcy)
	return api
}

// String	否	客户端自定义的订单标识 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateRestAssetConvertEstimateQuoteAPI) ClQReqId(clQReqId string) *PrivateRestAssetConvertEstimateQuoteAPI {
	api.req.ClQReqId = GetPointer(clQReqId)
	return api
}

// String	否	订单标签 适用于broker用户
func (api *PrivateRestAssetConvertEstimateQuoteAPI) Tag(tag string) *PrivateRestAssetConvertEstimateQuoteAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// 闪兑下单
// quoteId	String	是	报价ID
// baseCcy	String	是	交易货币币种，如 BTC-USDT中的BTC
// quoteCcy	String	是	计价货币币种，如 BTC-USDT中的USDT
// side	String	是	交易方向
// buy：买
// sell：卖
// 描述的是对于baseCcy的交易方向
// sz	String	是	用户报价数量
// 报价数量应不大于预估询价中的询价数量
// szCcy	String	是	用户报价币种
// clTReqId	String	否	用户自定义的订单标识
// 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
// tag	String	否	订单标签
// 适用于broker用户
type PrivateRestAssetConvertTradeReq struct {
	QuoteId  *string `json:"quoteId"`  //String	是	报价ID
	BaseCcy  *string `json:"baseCcy"`  //String	是	交易货币币种，如 BTC-USDT中的BTC
	QuoteCcy *string `json:"quoteCcy"` //String	是	计价货币币种，如 BTC-USDT中的USDT
	Side     *string `json:"side"`     //String	是	交易方向 buy：买 sell：卖 描述的是对于baseCcy的交易方向
	Sz       *string `json:"sz"`       //String	是	用户报价数量 报价数量应不大于预估询价中的询价数量
	SzCcy    *string `json:"szCcy"`    //String	是	用户报价币种
	ClTReqId *string `json:"clTReqId"` //String	否	用户自定义的订单标识 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag      *string `json:"tag"`      //String	否	订单标签 适用于broker用户
}
type PrivateRestAssetConvertTradeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetConvertTradeReq
}

// String	是	报价ID
func (api *PrivateRestAssetConvertTradeAPI) QuoteId(quoteId string) *PrivateRestAssetConvertTradeAPI {
	api.req.QuoteId = GetPointer(quoteId)
	return api
}

// String	是	交易货币币种，如 BTC-USDT中的BTC
func (api *PrivateRestAssetConvertTradeAPI) BaseCcy(baseCcy string) *PrivateRestAssetConvertTradeAPI {
	api.req.BaseCcy = GetPointer(baseCcy)
	return api
}

// String	是	计价货币币种，如 BTC-USDT中的USDT
func (api *PrivateRestAssetConvertTradeAPI) QuoteCcy(quoteCcy string) *PrivateRestAssetConvertTradeAPI {
	api.req.QuoteCcy = GetPointer(quoteCcy)
	return api
}

// String	是	交易方向 buy：买 sell：卖 描述的是对于baseCcy的交易方向
func (api *PrivateRestAssetConvertTradeAPI) Side(side string) *PrivateRestAssetConvertTradeAPI {
	api.req.Side = GetPointer(side)
	return api
}

// String	是	用户报价数量 报价数量应不大于预估询价中的询价数量
func (api *PrivateRestAssetConvertTradeAPI) Sz(sz string) *PrivateRestAssetConvertTradeAPI {
	api.req.Sz = GetPointer(sz)
	return api
}

// String	是	用户报价币种
func (api *PrivateRestAssetConvertTradeAPI) SzCcy(szCcy string) *PrivateRestAssetConvertTradeAPI {
	api.req.SzCcy = GetPointer(szCcy)
	return api
}

// String	否	用户自定义的订单标识 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateRestAssetConvertTradeAPI) ClTReqId(clTReqId string) *PrivateRestAssetConvertTradeAPI {
	api.req.ClTReqId = GetPointer(clTReqId)
	return api
}

// String	否	订单标签 适用于broker用户
func (api *PrivateRestAssetConvertTradeAPI) Tag(tag string) *PrivateRestAssetConvertTradeAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// 资金划转状态查询
type PrivateRestAssetTransferStateReq struct {
	TransId  *string `json:"transId"`  //String	可选	划转ID
	ClientId *string `json:"clientId"` //String	可选	客户自定义ID
	Type     *string `json:"type"`     //String	否	划转类型 0：账户内划转 1：母账户转子账户(仅适用于母账户APIKey) 2：子账户转母账户(仅适用于母账户APIKey) 3：子账户转母账户(仅适用于子账户APIKey) 4：子账户转子账户(仅适用于子账户APIKey，且目标账户需要是同一母账户下的其他子账户) 默认是0
}
type PrivateRestAssetTransferStateAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAssetTransferStateReq
}

// String	可选	划转ID
func (api *PrivateRestAssetTransferStateAPI) TransId(transId string) *PrivateRestAssetTransferStateAPI {
	api.req.TransId = GetPointer(transId)
	return api
}

// String	可选	客户自定义ID
func (api *PrivateRestAssetTransferStateAPI) ClientId(clientId string) *PrivateRestAssetTransferStateAPI {
	api.req.ClientId = GetPointer(clientId)
	return api
}

// String	否	划转类型 0：账户内划转 1：母账户转子账户(仅适用于母账户APIKey) 2：子账户转母账户(仅适用于母账户APIKey) 3：子账户转母账户(仅适用于子账户APIKey) 4：子账户转子账户(仅适用于子账户APIKey，且目标账户需要是同一母账户下的其他子账户) 默认是0
func (api *PrivateRestAssetTransferStateAPI) Type(t string) *PrivateRestAssetTransferStateAPI {
	api.req.Type = GetPointer(t)
	return api
}

type PrivateTradingBotGridOrderAlgoTriggerParam struct {
	TriggerAction   *string `json:"triggerAction"`   //String	是	触发行为 start：网格启动 stop：网格停止
	TriggerStrategy *string `json:"triggerStrategy"` //String	是	触发策略 instant：立即触发 price：价格触发 rsi：rsi指标触发 默认为instant
	DelaySeconds    *string `json:"delaySeconds"`    //String	否	延迟触发时间，单位为秒，默认为0
	Timeframe       *string `json:"timeframe"`       //String	否	K线种类 3m, 5m, 15m, 30m (m代表分钟) 1H, 4H (H代表小时) 1D (D代表天) 该字段只在triggerStrategy为rsi时有效
	Thold           *string `json:"thold"`           //String	否	阈值 取值[1,100]的整数 该字段只在triggerStrategy为rsi时有效
	TriggerCond     *string `json:"triggerCond"`     //String	否	触发条件 cross_up：上穿 cross_down：下穿 above：上方 below：下方 cross：交叉 该字段只在triggerStrategy为rsi时有效
	TimePeriod      *string `json:"timePeriod"`      //String	否	周期 14 该字段只在triggerStrategy为rsi下有效
	TriggerPx       *string `json:"triggerPx"`       //String	否	触发价格 该字段只在triggerStrategy为price下有效
	StopType        *string `json:"stopType"`        //String	否	策略停止类型 现货 1：卖出交易币，2：不卖出交易币 合约网格 1：停止平仓，2：停止不平仓 该字段只在triggerAction为stop时有效
}
type PrivateTradingBotGridOrderAlgoPostReq struct {
	InstId             *string                                       `json:"instId"`             //String	是	产品ID，如BTC-USDT
	AlgoOrdType        *string                                       `json:"algoOrdType"`        //String	是	策略订单类型 grid：现货网格委托 contract_grid：合约网格委托
	MaxPx              *string                                       `json:"maxPx"`              //String	是	区间最高价格
	MinPx              *string                                       `json:"minPx"`              //String	是	区间最低价格
	GridNum            *string                                       `json:"gridNum"`            //String	是	网格数量
	RunType            *string                                       `json:"runType"`            //String	否	网格类型 1：等差，2：等比 默认为等差
	TpTriggerPx        *string                                       `json:"tpTriggerPx"`        //String	否	止盈触发价 适用于现货网格/合约网格
	SlTriggerPx        *string                                       `json:"slTriggerPx"`        //String	否	止损触发价 适用于现货网格/合约网格
	AlgoClOrdId        *string                                       `json:"algoClOrdId"`        //String	否	用户自定义策略ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag                *string                                       `json:"tag"`                //String	否	订单标签
	ProfitSharingRatio *string                                       `json:"profitSharingRatio"` //String	否	带单员分润比例，仅支持固定比例分润 0,0.1,0.2,0.3
	TriggerParams      *[]PrivateTradingBotGridOrderAlgoTriggerParam `json:"triggerParams"`      //Array of object	否	信号触发参数 适用于现货网格/合约网格

	// 现货网格
	QuoteSz *string `json:"quoteSz"` //String	可选	计价币投入数量 quoteSz和baseSz至少指定一个
	BaseSz  *string `json:"baseSz"`  //String	可选	交易币投入数量 quoteSz和baseSz至少指定一个

	// 合约网格
	Sz        *string `json:"sz"`        //String	是	投入保证金,单位为USDT
	Direction *string `json:"direction"` //String	是	合约网格类型 long：做多，short：做空，neutral：中性
	Lever     *string `json:"lever"`     //String	是	杠杆倍数
	BasePos   *bool   `json:"basePos"`   //Boolean	否	是否开底仓 默认为false 中性合约网格忽略该参数
	TpRatio   *string `json:"tpRatio"`   //String	否	止盈比率，0.1 代表 10%
	SlRatio   *string `json:"slRatio"`   //String	否	止损比率，0.1 代表 10%

}
type PrivateTradingBotGridOrderAlgoPostAPI struct {
	client *PrivateRestClient
	req    *PrivateTradingBotGridOrderAlgoPostReq
}

// String	是	产品ID，如BTC-USDT
func (api *PrivateTradingBotGridOrderAlgoPostAPI) InstId(instId string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

// String	是	策略订单类型 grid：现货网格委托 contract_grid：合约网格委托
func (api *PrivateTradingBotGridOrderAlgoPostAPI) AlgoOrdType(algoOrdType string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.AlgoOrdType = GetPointer(algoOrdType)
	return api
}

// String	是	区间最高价格
func (api *PrivateTradingBotGridOrderAlgoPostAPI) MaxPx(maxPx string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.MaxPx = GetPointer(maxPx)
	return api
}

// String	是	区间最低价格
func (api *PrivateTradingBotGridOrderAlgoPostAPI) MinPx(minPx string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.MinPx = GetPointer(minPx)
	return api
}

// String	是	网格数量
func (api *PrivateTradingBotGridOrderAlgoPostAPI) GridNum(gridNum string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.GridNum = GetPointer(gridNum)
	return api
}

// String	否	网格类型 1：等差，2：等比 默认为等差
func (api *PrivateTradingBotGridOrderAlgoPostAPI) RunType(runType string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.RunType = GetPointer(runType)
	return api
}

// String	否	止盈触发价 适用于现货网格/合约网格
func (api *PrivateTradingBotGridOrderAlgoPostAPI) TpTriggerPx(tpTriggerPx string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.TpTriggerPx = GetPointer(tpTriggerPx)
	return api
}

// String	否	止损触发价 适用于现货网格/合约网格
func (api *PrivateTradingBotGridOrderAlgoPostAPI) SlTriggerPx(slTriggerPx string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.SlTriggerPx = GetPointer(slTriggerPx)
	return api
}

// String	否	用户自定义策略ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateTradingBotGridOrderAlgoPostAPI) AlgoClOrdId(algoClOrdId string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.AlgoClOrdId = GetPointer(algoClOrdId)
	return api
}

// String	否	订单标签
func (api *PrivateTradingBotGridOrderAlgoPostAPI) Tag(tag string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// String	否	带单员分润比例，仅支持固定比例分润 0,0.1,0.2,0.3
func (api *PrivateTradingBotGridOrderAlgoPostAPI) ProfitSharingRatio(profitSharingRatio string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.ProfitSharingRatio = GetPointer(profitSharingRatio)
	return api
}

// Array of object	否	信号触发参数 适用于现货网格/合约网格
func (api *PrivateTradingBotGridOrderAlgoPostAPI) TriggerParams(triggerParams []PrivateTradingBotGridOrderAlgoTriggerParam) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.TriggerParams = &triggerParams
	return api
}

// String	可选	计价币投入数量 quoteSz和baseSz至少指定一个 现货网格
func (api *PrivateTradingBotGridOrderAlgoPostAPI) QuoteSz(quoteSz string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.QuoteSz = GetPointer(quoteSz)
	return api
}

// String	可选	交易币投入数量 quoteSz和baseSz至少指定一个 现货网格
func (api *PrivateTradingBotGridOrderAlgoPostAPI) BaseSz(baseSz string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.BaseSz = GetPointer(baseSz)
	return api
}

// String	是	投入保证金,单位为USDT 合约网格
func (api *PrivateTradingBotGridOrderAlgoPostAPI) Sz(sz string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.Sz = GetPointer(sz)
	return api
}

// String	是	合约网格类型 long：做多，short：做空，neutral：中性
func (api *PrivateTradingBotGridOrderAlgoPostAPI) Direction(direction string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.Direction = GetPointer(direction)
	return api
}

// String	是	杠杆倍数
func (api *PrivateTradingBotGridOrderAlgoPostAPI) Lever(lever string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.Lever = GetPointer(lever)
	return api
}

// Boolean	否	是否开底仓 默认为false 中性合约网格忽略该参数
func (api *PrivateTradingBotGridOrderAlgoPostAPI) BasePos(basePos bool) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.BasePos = GetPointer(basePos)
	return api
}

// String	否	止盈比率，0.1 代表 10%
func (api *PrivateTradingBotGridOrderAlgoPostAPI) TpRatio(tpRatio string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.TpRatio = GetPointer(tpRatio)
	return api
}

// String	否	止损比率，0.1 代表 10%
func (api *PrivateTradingBotGridOrderAlgoPostAPI) SlRatio(slRatio string) *PrivateTradingBotGridOrderAlgoPostAPI {
	api.req.SlRatio = GetPointer(slRatio)
	return api
}

type PrivateRestRfqCounterPartiesReq struct{}
type PrivateRestRfqCounterPartiesAPI struct {
	client *PrivateRestClient
	req    *PrivateRestRfqCounterPartiesReq
}

type PrivateTradingBotRecurringOrderAlgoPostRecurring struct {
	Ccy   *string `json:"ccy"`   //String	是	定投币种，如 BTC
	Ratio *string `json:"ratio"` //String	是	定投币种资产占比，如 "0.2"代表占比20%
}
type PrivateTradingBotRecurringOrderAlgoPostReq struct {
	StgyName      *string                                             `json:"stgyName"`      //String	是	策略自定义名称，不超过40个字符
	RecurringList *[]PrivateTradingBotRecurringOrderAlgoPostRecurring `json:"recurringList"` //Array of object	是	定投信息
	Period        *string                                             `json:"period"`        //String	是	周期类型 monthly：月 weekly：周 daily：日 hourly：小时
	RecurringDay  *string                                             `json:"recurringDay"`  //String	可选	投资日 当周期类型为monthly，则取值范围是 [1,28] 的整数 当周期类型为weekly，则取值范围是 [1,7] 的整数 当周期类型为daily/hourly，该参数可不填。
	RecurringHour *string                                             `json:"recurringHour"` //String	可选	小时级别定投的间隔 1/4/8/12 如：1代表每隔1个小时定投 当周期类型选择hourly，该字段必填。
	RecurringTime *string                                             `json:"recurringTime"` //String	是	投资时间，取值范围是 [0,23] 的整数 当周期类型选择hourly代表首次定投发生的时间
	TimeZone      *string                                             `json:"timeZone"`      //String	是	时区（UTC），取值范围是 [-12,14] 的整数 如 8表示UTC+8（东8区），北京时间
	Amt           *string                                             `json:"amt"`           //String	是	每期投入数量
	InvestmentCcy *string                                             `json:"investmentCcy"` //String	是	投入数量单位，只能是USDT/USDC
	TdMode        *string                                             `json:"tdMode"`        //String	是	交易模式 跨币种保证金模式/组合保证金模式下选择 cross：全仓 现货模式/现货和合约模式下选择 cash：非保证金
	AlgoClOrdId   *string                                             `json:"algoClOrdId"`   //String	否	客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag           *string                                             `json:"tag"`           //String	否	订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
}
type PrivateTradingBotRecurringOrderAlgoPostAPI struct {
	client *PrivateRestClient
	req    *PrivateTradingBotRecurringOrderAlgoPostReq
}

// String	是	策略自定义名称，不超过40个字符
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) StgyName(stgyName string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.StgyName = GetPointer(stgyName)
	return api
}

// Array of object	是	定投信息
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) RecurringList(recurringList []PrivateTradingBotRecurringOrderAlgoPostRecurring) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.RecurringList = &recurringList
	return api
}

// String	是	周期类型 monthly：月 weekly：周 daily：日 hourly：小时
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) Period(period string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.Period = GetPointer(period)
	return api
}

// String	可选	投资日 当周期类型为monthly，则取值范围是 [1,28] 的整数 当周期类型为weekly，则取值范围是 [1,7] 的整数 当周期类型为daily/hourly，该参数可不填。
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) RecurringDay(recurringDay string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.RecurringDay = GetPointer(recurringDay)
	return api
}

// String	可选	小时级别定投的间隔 1/4/8/12 如：1代表每隔1个小时定投 当周期类型选择hourly，该字段必填。
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) RecurringHour(recurringHour string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.RecurringHour = GetPointer(recurringHour)
	return api
}

// String	是	投资时间，取值范围是 [0,23] 的整数 当周期类型选择hourly代表首次定投发生的时间
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) RecurringTime(recurringTime string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.RecurringTime = GetPointer(recurringTime)
	return api
}

// String	是	时区（UTC），取值范围是 [-12,14] 的整数 如 8表示UTC+8（东8区），北京时间
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) TimeZone(timeZone string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.TimeZone = GetPointer(timeZone)
	return api
}

// String	是	每期投入数量
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) Amt(amt string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.Amt = GetPointer(amt)
	return api
}

// String	是	投入数量单位，只能是USDT/USDC
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) InvestmentCcy(investmentCcy string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.InvestmentCcy = GetPointer(investmentCcy)
	return api
}

// String	是	交易模式 跨币种保证金模式/组合保证金模式下选择 cross：全仓 现货模式/现货和合约模式下选择 cash：非保证金
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) TdMode(tdMode string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.TdMode = GetPointer(tdMode)
	return api
}

// String	否	客户自定义订单ID 字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) AlgoClOrdId(algoClOrdId string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.AlgoClOrdId = GetPointer(algoClOrdId)
	return api
}

// String	否	订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
func (api *PrivateTradingBotRecurringOrderAlgoPostAPI) Tag(tag string) *PrivateTradingBotRecurringOrderAlgoPostAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// 大宗交易询价
type PrivateRestRfqCreateRfqLeg struct {
	InstId  *string `json:"instId"`  //String	是	产品ID
	TdMode  *string `json:"tdMode"`  //String	否	交易模式 保证金模式：cross全仓 isolated逐仓 非保证金模式：cash非保证金. 如未提供，tdMode 将继承系统设置的默认值： 现货和合约模式 & 现货: cash 现货和合约模式和跨币种保证金模式下买入期权： isolated 其他情况: cross
	Ccy     *string `json:"ccy"`     //String	否	保证金币种，仅适用于现货和合约模式下的全仓杠杆订单 在其他情况下该参数将被忽略。
	Sz      *string `json:"sz"`      //String	是	委托数量
	Side    *string `json:"side"`    //String	是	询价单方向
	PosSide *string `json:"posSide"` //String	否	持仓方向 买卖模式下默认为net。在开平仓模式下仅可选择long或short。 如未指定，则处于开平仓模式下的用户始终会开新仓位。 仅适用交割、永续。
	TgtCcy  *string `json:"tgtCcy"`  //String	否	委托数量的类型 定义sz属性的单位。仅适用于 instType=SPOT。有效值为base_ccy和quote_ccy。未指定时，默认为base_ccy。
}
type PrivateRestRfqCreateRfqReq struct {
	Counterparties        *[]string                     `json:"counterparties"`        //Array of strings	是	希望收到询价的报价方列表，可通过/api/v5/rfq/counterparties/获取。
	Anonymous             *bool                         `json:"anonymous"`             //Boolean	否	是否匿名询价，true表示匿名询价，false表示公开询价，默认值为 false，为true时，即使在交易执行之后，身份也不会透露给报价方。
	ClRfqId               *string                       `json:"clRfqId"`               //String	否	询价单自定义ID，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag                   *string                       `json:"tag"`                   //String	否	询价单标签，与此询价单关联的大宗交易将有相同的标签。
	AllowPartialExecution *bool                         `json:"allowPartialExecution"` //Boolean	否	RFQ是否可以被部分执行，如果腿的比例和原RFQ一致。有效值为true或false。默认为false。
	Legs                  *[]PrivateRestRfqCreateRfqLeg `json:"legs"`                  //Array of objects	是	组合交易，每次最多可以提交15组交易信息
}
type PrivateRestRfqCreateRfqAPI struct {
	client *PrivateRestClient
	req    *PrivateRestRfqCreateRfqReq
}

// Array of strings	是	希望收到询价的报价方列表，可通过/api/v5/rfq/counterparties/获取。
func (api *PrivateRestRfqCreateRfqAPI) Counterparties(counterparties []string) *PrivateRestRfqCreateRfqAPI {
	api.req.Counterparties = &counterparties
	return api
}

// Boolean	否	是否匿名询价，true表示匿名询价，false表示公开询价，默认值为 false，为true时，即使在交易执行之后，身份也不会透露给报价方。
func (api *PrivateRestRfqCreateRfqAPI) Anonymous(anonymous bool) *PrivateRestRfqCreateRfqAPI {
	api.req.Anonymous = GetPointer(anonymous)
	return api
}

// String	否	询价单自定义ID，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateRestRfqCreateRfqAPI) ClRfqId(clRfqId string) *PrivateRestRfqCreateRfqAPI {
	api.req.ClRfqId = GetPointer(clRfqId)
	return api
}

// String	否	询价单标签，与此询价单关联的大宗交易将有相同的标签。
func (api *PrivateRestRfqCreateRfqAPI) Tag(tag string) *PrivateRestRfqCreateRfqAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// Boolean	否	RFQ是否可以被部分执行，如果腿的比例和原RFQ一致。有效值为true或false。默认为false。
func (api *PrivateRestRfqCreateRfqAPI) AllowPartialExecution(allowPartialExecution bool) *PrivateRestRfqCreateRfqAPI {
	api.req.AllowPartialExecution = GetPointer(allowPartialExecution)
	return api
}

// Array of objects	是	组合交易，每次最多可以提交15组交易信息
func (api *PrivateRestRfqCreateRfqAPI) Legs(legs []PrivateRestRfqCreateRfqLeg) *PrivateRestRfqCreateRfqAPI {
	api.req.Legs = &legs
	return api
}

// 价差交易下单
type PrivateRestSprdOrderPostReq struct {
	SprdId  *string `json:"sprdId"`  //String	是	spread ID，如 BTC-USDT_BTC-USDT-SWAP
	ClOrdId *string `json:"clOrdId"` //String	否	客户自定义订单ID字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
	Tag     *string `json:"tag"`     //String	否	订单标签字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
	Side    *string `json:"side"`    //String	是	订单方向 buy：买，sell：卖
	OrdType *string `json:"ordType"` //String	是	订单类型 limit：限价单 post_only：只做maker单 ioc：立即成交并取消剩余
	Sz      *string `json:"sz"`      //String	是	委托数量。反向价差的数量单位为USD，正向及混合价差为其对应baseCcy
	Px      *string `json:"px"`      //String	是	委托价格，仅适用于limit, post_only, ioc类型的订单
}
type PrivateRestSprdOrderPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestSprdOrderPostReq
}

// String	是	spread ID，如 BTC-USDT_BTC-USDT-SWAP
func (api *PrivateRestSprdOrderPostAPI) SprdId(sprdId string) *PrivateRestSprdOrderPostAPI {
	api.req.SprdId = GetPointer(sprdId)
	return api
}

// String	否	客户自定义订单ID字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间。
func (api *PrivateRestSprdOrderPostAPI) ClOrdId(clOrdId string) *PrivateRestSprdOrderPostAPI {
	api.req.ClOrdId = GetPointer(clOrdId)
	return api
}

// String	否	订单标签字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间。
func (api *PrivateRestSprdOrderPostAPI) Tag(tag string) *PrivateRestSprdOrderPostAPI {
	api.req.Tag = GetPointer(tag)
	return api
}

// String	是	订单方向 buy：买，sell：卖
func (api *PrivateRestSprdOrderPostAPI) Side(side string) *PrivateRestSprdOrderPostAPI {
	api.req.Side = GetPointer(side)
	return api
}

// String	是	订单类型 limit：限价单 post_only：只做maker单 ioc：立即成交并取消剩余
func (api *PrivateRestSprdOrderPostAPI) OrdType(ordType string) *PrivateRestSprdOrderPostAPI {
	api.req.OrdType = GetPointer(ordType)
	return api
}

// String	是	委托数量。反向价差的数量单位为USD，正向及混合价差为其对应baseCcy
func (api *PrivateRestSprdOrderPostAPI) Sz(sz string) *PrivateRestSprdOrderPostAPI {
	api.req.Sz = GetPointer(sz)
	return api
}

// String	是	委托价格，仅适用于limit, post_only, ioc类型的订单
func (api *PrivateRestSprdOrderPostAPI) Px(px string) *PrivateRestSprdOrderPostAPI {
	api.req.Px = GetPointer(px)
	return api
}

// 金融产品链上赚币查看项目
type PrivateRestFinanceStakingDefiOffersReq struct {
	ProductId    *string `json:"productId"`    //String	否	项目ID
	ProtocolType *string `json:"protocolType"` //String	否	项目类型 defi：链上赚币
	Ccy          *string `json:"ccy"`          //String	否	投资币种，如 BTC
}
type PrivateRestFinanceStakingDefiOffersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFinanceStakingDefiOffersReq
}

// String	否	项目ID
func (api *PrivateRestFinanceStakingDefiOffersAPI) ProductId(productId string) *PrivateRestFinanceStakingDefiOffersAPI {
	api.req.ProductId = GetPointer(productId)
	return api
}

// String	否	项目类型 defi：链上赚币
func (api *PrivateRestFinanceStakingDefiOffersAPI) ProtocolType(protocolType string) *PrivateRestFinanceStakingDefiOffersAPI {
	api.req.ProtocolType = GetPointer(protocolType)
	return api
}

// String	否	投资币种，如 BTC
func (api *PrivateRestFinanceStakingDefiOffersAPI) Ccy(ccy string) *PrivateRestFinanceStakingDefiOffersAPI {
	api.req.Ccy = GetPointer(ccy)
	return api
}

// 金融产品链上赚币申购项目
type PrivateRestFinanceStakingDefiPurchaseInvestData struct {
	Ccy *string `json:"ccy"` //String	是	投资币种，如 BTC
	Amt *string `json:"amt"` //String	是	投资数量
}
type PrivateRestFinanceStakingDefiPurchaseReq struct {
	ProductId  *string                                            `json:"productId"`  //String	是	项目ID
	InvestData *[]PrivateRestFinanceStakingDefiPurchaseInvestData `json:"investData"` //Array	是	投资信息
	Term       *string                                            `json:"term"`       //String	可选	投资期限 定期项目必须指定投资期限
	Tag        *string                                            `json:"tag"`        //String	否	订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间
}
type PrivateRestFinanceStakingDefiPurchaseAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFinanceStakingDefiPurchaseReq
}

// String	是	项目ID
func (api *PrivateRestFinanceStakingDefiPurchaseAPI) ProductId(productId string) *PrivateRestFinanceStakingDefiPurchaseAPI {
	api.req.ProductId = GetPointer(productId)
	return api
}

// Array	是	投资信息
func (api *PrivateRestFinanceStakingDefiPurchaseAPI) InvestData(investData []PrivateRestFinanceStakingDefiPurchaseInvestData) *PrivateRestFinanceStakingDefiPurchaseAPI {
	api.req.InvestData = &investData
	return api
}

// String	可选	投资期限 定期项目必须指定投资期限
func (api *PrivateRestFinanceStakingDefiPurchaseAPI) Term(term string) *PrivateRestFinanceStakingDefiPurchaseAPI {
	api.req.Term = GetPointer(term)
	return api
}

// String	否	订单标签 字母（区分大小写）与数字的组合，可以是纯字母、纯数字，且长度在1-16位之间
func (api *PrivateRestFinanceStakingDefiPurchaseAPI) Tag(tag string) *PrivateRestFinanceStakingDefiPurchaseAPI {
	api.req.Tag = GetPointer(tag)
	return api
}
