package myokxapi

type PrivateRestTradeOrderGetRes []PrivateRestTradeOrderGetResRow
type PrivateRestTradeOrderGetResRow struct {
	InstType           string                                        `json:"instType"`           //产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId             string                                        `json:"instId"`             //产品ID
	TgtCcy             string                                        `json:"tgtCcy"`             //币币市价单委托数量sz的单位 base_ccy: 交易货币 ；quote_ccy：计价货币 仅适用于币币市价订单 默认买单为quote_ccy，卖单为base_ccy
	Ccy                string                                        `json:"ccy"`                //保证金币种，仅适用于单币种保证金模式下的全仓币币杠杆订单
	OrdId              string                                        `json:"ordId"`              //订单ID
	ClOrdId            string                                        `json:"clOrdId"`            //客户自定义订单ID
	Tag                string                                        `json:"tag"`                //订单标签
	Px                 string                                        `json:"px"`                 //委托价格，对于期权，以币(如BTC, ETH)为单位
	PxUsd              string                                        `json:"pxUsd"`              //期权价格，以USD为单位 仅适用于期权，其他业务线返回空字符串""
	PxVol              string                                        `json:"pxVol"`              //期权订单的隐含波动率 仅适用于期权，其他业务线返回空字符串""
	PxType             string                                        `json:"pxType"`             //期权的价格类型 px：代表按价格下单，单位为币 (请求参数 px 的数值单位是BTC或ETH) pxVol：代表按pxVol下单 pxUsd：代表按照pxUsd下单，单位为USD (请求参数px 的数值单位是USD)
	Sz                 string                                        `json:"sz"`                 //委托数量
	Pnl                string                                        `json:"pnl"`                //收益，适用于有成交的平仓订单，其他情况均为0
	OrdType            string                                        `json:"ordType"`            //订单类型 market：市价单 limit：限价单 post_only：只做maker单 fok：全部成交或立即取消 ioc：立即成交并取消剩余 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单)
	Side               string                                        `json:"side"`               //订单方向
	PosSide            string                                        `json:"posSide"`            //持仓方向
	TdMode             string                                        `json:"tdMode"`             //交易模式
	AccFillSz          string                                        `json:"accFillSz"`          //累计成交数量 对于币币和杠杆，单位为交易货币，如 BTC-USDT, 单位为 BTC；对于市价单，无论tgtCcy是base_ccy，还是quote_ccy，单位均为交易货币； 对于交割、永续以及期权，单位为张。
	FillPx             string                                        `json:"fillPx"`             //最新成交价格，如果成交数量为0，该字段为""
	TradeId            string                                        `json:"tradeId"`            //最新成交ID
	FillSz             string                                        `json:"fillSz"`             //最新成交数量 对于币币和杠杆，单位为交易货币，如 BTC-USDT, 单位为 BTC；对于市价单，无论tgtCcy是base_ccy，还是quote_ccy，单位均为交易货币； 对于交割、永续以及期权，单位为张。
	FillTime           string                                        `json:"fillTime"`           //最新成交时间
	AvgPx              string                                        `json:"avgPx"`              //成交均价，如果成交数量为0，该字段也为""
	State              string                                        `json:"state"`              //订单状态 canceled：撤单成功 live：等待成交 partially_filled：部分成交 filled：完全成交 mmp_canceled：做市商保护机制导致的自动撤单
	Lever              string                                        `json:"lever"`              //杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	AttachAlgoClOrdId  string                                        `json:"attachAlgoClOrdId"`  //下单附带止盈止损时，客户自定义的策略订单ID
	TpTriggerPx        string                                        `json:"tpTriggerPx"`        //止盈触发价
	TpTriggerPxType    string                                        `json:"tpTriggerPxType"`    //止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格
	TpOrdPx            string                                        `json:"tpOrdPx"`            //止盈委托价
	SlTriggerPx        string                                        `json:"slTriggerPx"`        //止损触发价
	SlTriggerPxType    string                                        `json:"slTriggerPxType"`    //止损触发价类型 last：最新价格 index：指数价格 mark：标记价格
	SlOrdPx            string                                        `json:"slOrdPx"`            //止损委托价
	AttachAlgoOrds     []PrivateRestTradeOrderGetResRowattachAlgoOrd `json:"attachAlgoOrds"`     //下单附带止盈止损信息
	StpId              string                                        `json:"stpId"`              //自成交保护ID 如果自成交保护不适用则返回""
	StpMode            string                                        `json:"stpMode"`            //自成交保护模式 如果自成交保护不适用则返回""
	FeeCcy             string                                        `json:"feeCcy"`             //交易手续费币种
	Fee                string                                        `json:"fee"`                //手续费与返佣 对于币币和杠杆，为订单交易累计的手续费，平台向用户收取的交易手续费，为负数。如： -0.01 对于交割、永续和期权，为订单交易累计的手续费和返佣
	RebateCcy          string                                        `json:"rebateCcy"`          //返佣金币种
	Source             string                                        `json:"source"`             //订单来源 6：计划委托策略触发后的生成的普通单 7：止盈止损策略触发后的生成的普通单 13：策略委托单触发后的生成的普通单 24：移动止盈止损策略触发后的生成的普通单
	Rebate             string                                        `json:"rebate"`             //返佣金额，仅适用于币币和杠杆，平台向达到指定lv交易等级的用户支付的挂单奖励（返佣），如果没有返佣金，该字段为“”。手续费返佣为正数，如：0.01
	Category           string                                        `json:"category"`           //订单种类 normal：普通委托 twap：TWAP自动换币 adl：ADL自动减仓 full_liquidation：强制平仓 partial_liquidation：强制减仓 delivery：交割 ddh：对冲减仓类型订单
	ReduceOnly         string                                        `json:"reduceOnly"`         //是否只减仓，true 或 false
	CancelSource       string                                        `json:"cancelSource"`       //订单取消来源的原因枚举值代码
	CancelSourceReason string                                        `json:"cancelSourceReason"` //订单取消来源的对应具体原因
	QuickMgnType       string                                        `json:"quickMgnType"`       //一键借币类型，仅适用于杠杆逐仓的一键借币模式 manual：手动，auto_borrow： 自动借币，auto_repay： 自动还币
	AlgoClOrdId        string                                        `json:"algoClOrdId"`        //客户自定义策略订单ID。策略订单触发，且策略单有algoClOrdId时有值，否则为"",
	AlgoId             string                                        `json:"algoId"`             //策略委托单ID，策略订单触发时有值，否则为""
	UTime              string                                        `json:"uTime"`              //订单状态更新时间，Unix时间戳的毫秒数格式，如：1597026383085
	CTime              string                                        `json:"cTime"`              //订单创建时间，Unix时间戳的毫秒数格式， 如 ：1597026383085
}

type PrivateRestTradeOrdersPendingRes []PrivateRestTradeOrderGetResRow

type PrivateRestTradeOrderGetResRowattachAlgoOrd struct {
	AttachAlgoId         string `json:"attachAlgoId"`         //附带止盈止损的订单ID，订单完全成交，下止盈止损委托单时，该值不会传给 algoId
	AttachAlgoClOrdId    string `json:"attachAlgoClOrdId"`    //下单附带止盈止损时，客户自定义的策略订单ID
	TpTriggerPx          string `json:"tpTriggerPx"`          //止盈触发价
	TpTriggerPxType      string `json:"tpTriggerPxType"`      //止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格
	TpOrdPx              string `json:"tpOrdPx"`              //止盈委托价
	SlTriggerPx          string `json:"slTriggerPx"`          //止损触发价
	SlTriggerPxType      string `json:"slTriggerPxType"`      //止损触发价类型 last：最新价格 index：指数价格 mark：标记价格
	SlOrdPx              string `json:"slOrdPx"`              //止损委托价
	Sz                   string `json:"sz"`                   //张数。仅适用于“多笔止盈”的止盈订单且必填
	AmendPxOnTriggerType string `json:"amendPxOnTriggerType"` //是否启用开仓价止损，仅适用于分批止盈的止损订单 0. 不开启，默认值 1. 开启“开仓价止损”
}

type PrivateRestTradeOrderPostRes []PrivateRestTradeOrderPostResRow
type PrivateRestTradeOrderPostResRow struct {
	OrdId   string `json:"ordId"`   //String	订单ID
	ClOrdId string `json:"clOrdId"` //String	客户自定义订单ID
	Tag     string `json:"tag"`     //String	订单标签
	SCode   string `json:"sCode"`   //String	订单状态码
	SMsg    string `json:"sMsg"`    //String	订单状态消息
}

type PrivateRestTradeCancelOrderRes []PrivateRestTradeCancelOrderResRow
type PrivateRestTradeCancelOrderResRow struct {
	OrdId   string `json:"ordId"`   //String	订单ID
	ClOrdId string `json:"clOrdId"` //String	客户自定义订单ID
	SCode   string `json:"sCode"`   //String	事件执行结果的code，0代表成功
	SMsg    string `json:"sMsg"`    //String	事件执行失败时的msg
}

type PrivateRestTradeAmendOrderRes []PrivateRestTradeAmendOrderResRow
type PrivateRestTradeAmendOrderResRow struct {
	OrdId   string `json:"ordId"`   //String	订单ID
	ClOrdId string `json:"clOrdId"` //String	客户自定义订单ID
	ReqId   string `json:"reqId"`   //String	用户自定义修改事件ID
	SCode   string `json:"sCode"`   //String	事件执行结果的code，0代表成功
	SMsg    string `json:"sMsg"`    //String	事件执行失败时的msg
}

type PrivateRestTradeBatchOrdersRes []PrivateRestTradeOrderPostResRow

type PrivateRestTradeCancelBatchOrdersRes []PrivateRestTradeCancelOrderResRow

type PrivateRestTradeAmendBatchOrdersRes []PrivateRestTradeAmendOrderResRow


type PrivateRestTradeOrderHistoryRes []PrivateRestTradeOrderGetResRow
type PrivateRestTradeOrderHistoryArchiveRes []PrivateRestTradeOrderGetResRow