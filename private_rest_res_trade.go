package myokxapi

type PrivateRestTradeOrderGetRes []PrivateRestTradeOrderGetResRow
type PrivateRestTradeOrderGetResRow struct {
	InstType           string                                        `json:"instType"`           //产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId             string                                        `json:"instId"`             //产品ID
	TgtCcy             string                                        `json:"tgtCcy"`             //币币市价单委托数量sz的单位 base_ccy: 交易货币 ；quote_ccy：计价货币 仅适用于币币市价订单 默认买单为quote_ccy，卖单为base_ccy
	Ccy                string                                        `json:"ccy"`                //保证金币种，仅适用于单币种保证金模式下的全仓币币杠杆订单
	OrdId              string                                        `json:"ordId,omitempty"`    //订单ID
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

type PrivateRestTradeOrderAlgoPostRes []PrivateRestTradeOrderAlgoPostResRow
type PrivateRestTradeOrderAlgoPostResRow struct {
	AlgoId      string `json:"algoId"`      //String	策略委托单ID
	AlgoClOrdId string `json:"algoClOrdId"` //String	客户自定义策略订单ID
	SCode       string `json:"sCode"`       //String	事件执行结果的code，0代表成功
	SMsg        string `json:"sMsg"`        //String	事件执行失败时的msg
	Tag         string `json:"tag"`         //String	订单标签
}

type PrivateRestTradeCancelOrderAlgoRes []PrivateRestTradeCancelOrderAlgoResRow
type PrivateRestTradeCancelOrderAlgoResRow struct {
	AlgoId      string `json:"algoId"`      //String	策略委托单ID
	AlgoClOrdId string `json:"algoClOrdId"` //String	客户自定义策略订单ID
	SCode       string `json:"sCode"`       //String	事件执行结果的code，0代表成功
}

type PrivateRestTradeAmendOrderAlgoRes []PrivateRestTradeAmendOrderAlgoResRow
type PrivateRestTradeAmendOrderAlgoResRow struct {
	AlgoId      string `json:"algoId"`      //String	策略委托单ID
	AlgoClOrdId string `json:"algoClOrdId"` //String	客户自定义策略订单ID
	SCode       string `json:"sCode"`       //String	事件执行结果的code，0代表成功
	SMsg        string `json:"sMsg"`        //String	事件执行失败时的msg
}

type PrivateRestTradeOrderAlgoGetResRowAttachAlgoOrdRow struct {
	AttachAlgoClOrdId string `json:"attachAlgoClOrdId"` //String	下单附带止盈止损时，客户自定义的策略订单ID
	TpTriggerPx       string `json:"tpTriggerPx"`       //String	止盈触发价
	TpTriggerPxType   string `json:"tpTriggerPxType"`   //String	止盈触发价类型
	TpOrdPx           string `json:"tpOrdPx"`           //String	止盈委托价
	SlTriggerPx       string `json:"slTriggerPx"`       //String	止损触发价
	SlTriggerPxType   string `json:"slTriggerPxType"`   //String	止损触发价类型
	SlOrdPx           string `json:"slOrdPx"`           //String	止损委托价
}
type PrivateRestTradeOrderAlgoGetResRowLinkedOrd struct {
	OrdId string `json:"ordId"` //String	订单ID
}
type PrivateRestTradeOrderAlgoGetRes []PrivateRestTradeOrderAlgoGetResRow
type PrivateRestTradeOrderAlgoGetResRow struct {
	InstType             string                                               `json:"instType"`             //产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId               string                                               `json:"instId"`               //产品ID
	Ccy                  string                                               `json:"ccy"`                  //保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
	OrdId                string                                               `json:"ordId"`                //最新一笔订单ID，即将废弃。
	OrdIdList            []string                                             `json:"ordIdList,omitempty"`  //订单ID列表，当止盈止损存在市价拆单时，会有多个。
	AlgoId               string                                               `json:"algoId"`               //策略委托单ID
	ClOrdId              string                                               `json:"clOrdId"`              //客户自定义订单ID
	Sz                   string                                               `json:"sz"`                   //委托数量
	CloseFraction        string                                               `json:"closeFraction"`        //策略委托触发时，平仓的百分比。1 代表100%
	OrdType              string                                               `json:"ordType"`              //订单类型
	Side                 string                                               `json:"side"`                 //订单方向
	PosSide              string                                               `json:"posSide"`              //持仓方向
	TdMode               string                                               `json:"tdMode"`               //交易模式
	TgtCcy               string                                               `json:"tgtCcy"`               //币币市价单委托数量sz的单位
	State                string                                               `json:"state"`                //订单状态
	Lever                string                                               `json:"lever"`                //杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	TpTriggerPx          string                                               `json:"tpTriggerPx"`          //止盈触发价
	TpTriggerPxType      string                                               `json:"tpTriggerPxType"`      //止盈触发价类型
	TpOrdPx              string                                               `json:"tpOrdPx"`              //止盈委托价
	SlTriggerPx          string                                               `json:"slTriggerPx"`          //止损触发价
	SlTriggerPxType      string                                               `json:"slTriggerPxType"`      //止损触发价类型
	SlOrdPx              string                                               `json:"slOrdPx"`              //止损委托价
	TriggerPx            string                                               `json:"triggerPx"`            //计划委托触发价格
	TriggerPxType        string                                               `json:"triggerPxType"`        //计划委托触发价格类型
	OrdPx                string                                               `json:"ordPx"`                //订单委托价格
	ActualSz             string                                               `json:"actualSz"`             //实际委托量
	ActualPx             string                                               `json:"actualPx"`             //实际委托价
	ActualSide           string                                               `json:"actualSide"`           //实际触发方向
	TriggerTime          string                                               `json:"triggerTime"`          //策略委托触发时间，Unix时间戳的毫秒数格式，如 1597026383085
	PxVar                string                                               `json:"pxVar"`                //价格比例
	PxSpread             string                                               `json:"pxSpread"`             //价距
	SzLimit              string                                               `json:"szLimit"`              //单笔数量
	PxLimit              string                                               `json:"pxLimit"`              //挂单限制价
	Tag                  string                                               `json:"tag"`                  //订单标签
	TimeInterval         string                                               `json:"timeInterval"`         //下单间隔
	CallbackRatio        string                                               `json:"callbackRatio"`        //回调幅度的比例
	CallbackSpread       string                                               `json:"callbackSpread"`       //回调幅度的价距
	ActivePx             string                                               `json:"activePx"`             //移动止盈止损激活价格
	MoveTriggerPx        string                                               `json:"moveTriggerPx"`        //移动止盈止损触发价格
	ReduceOnly           string                                               `json:"reduceOnly"`           //是否只减仓
	QuickMgnType         string                                               `json:"quickMgnType"`         //一键借币类型，仅适用于杠杆逐仓的一键借币模式
	Last                 string                                               `json:"last"`                 //下单时的最新成交价
	FailCode             string                                               `json:"failCode"`             //代表策略触发失败的原因，已撤销和已生效时为""，委托失败时有值，如 51008；
	AlgoClOrdId          string                                               `json:"algoClOrdId"`          //客户自定义策略订单ID
	AmendPxOnTriggerType string                                               `json:"amendPxOnTriggerType"` //是否启用开仓价止损
	AttachAlgoOrds       []PrivateRestTradeOrderAlgoGetResRowAttachAlgoOrdRow `json:"attachAlgoOrds"`       //附带止盈止损信息
	LinkedOrds           PrivateRestTradeOrderAlgoGetResRowLinkedOrd          `json:"linkedOrds"`           //附带关联订单信息
	CTime                string                                               `json:"cTime"`                //订单创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime                string                                               `json:"uTime"`                //订单状态更新时间，Unix时间戳的毫秒数格式，如 1597026383085
}

type PrivateRestTradeOrderPendingRowAttachAlgoOrd struct {
	AttachAlgoClOrdId string `json:"attachAlgoClOrdId"` //String	下单附带止盈止损时，客户自定义的策略订单ID
	TpTriggerPx       string `json:"tpTriggerPx"`       //String	止盈触发价
	TpTriggerPxType   string `json:"tpTriggerPxType"`   //String	止盈触发价类型
	TpOrdPx           string `json:"tpOrdPx"`           //String	止盈委托价
	SlTriggerPx       string `json:"slTriggerPx"`       //String	止损触发价
	SlTriggerPxType   string `json:"slTriggerPxType"`   //String	止损触发价类型
	SlOrdPx           string `json:"slOrdPx"`           //String	止损委托价
}
type PrivateRestTradeOrderPendingRowLinkedOrd struct {
	OrdId string `json:"ordId"` //String	订单ID
}
type PrivateRestTradeOrderAlgoPendingRes []PrivateRestTradeOrderAlgoPendingResRow
type PrivateRestTradeOrderAlgoPendingResRow struct {
	InstType             string                                         `json:"instType"`             //产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId               string                                         `json:"instId"`               //产品ID
	Ccy                  string                                         `json:"ccy"`                  //保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
	OrdId                string                                         `json:"ordId"`                //最新一笔订单ID，即将废弃。
	OrdIdList            []string                                       `json:"ordIdList,omitempty"`  //订单ID列表，当止盈止损存在市价拆单时，会有多个。
	AlgoId               string                                         `json:"algoId"`               //策略委托单ID
	ClOrdId              string                                         `json:"clOrdId"`              //客户自定义订单ID
	Sz                   string                                         `json:"sz"`                   //委托数量
	CloseFraction        string                                         `json:"closeFraction"`        //策略委托触发时，平仓的百分比。1 代表100%
	OrdType              string                                         `json:"ordType"`              //订单类型
	Side                 string                                         `json:"side"`                 //订单方向
	PosSide              string                                         `json:"posSide"`              //持仓方向
	TdMode               string                                         `json:"tdMode"`               //交易模式
	TgtCcy               string                                         `json:"tgtCcy"`               //币币市价单委托数量sz的单位
	State                string                                         `json:"state"`                //订单状态
	Lever                string                                         `json:"lever"`                //杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	TpTriggerPx          string                                         `json:"tpTriggerPx"`          //止盈触发价
	TpTriggerPxType      string                                         `json:"tpTriggerPxType"`      //止盈触发价类型
	TpOrdPx              string                                         `json:"tpOrdPx"`              //止盈委托价
	SlTriggerPx          string                                         `json:"slTriggerPx"`          //止损触发价
	SlTriggerPxType      string                                         `json:"slTriggerPxType"`      //止损触发价类型
	SlOrdPx              string                                         `json:"slOrdPx"`              //止损委托价
	TriggerPx            string                                         `json:"triggerPx"`            //计划委托触发价格
	TriggerPxType        string                                         `json:"triggerPxType"`        //计划委托触发价格类型
	OrdPx                string                                         `json:"ordPx"`                //订单委托价格
	ActualSz             string                                         `json:"actualSz"`             //实际委托量
	ActualPx             string                                         `json:"actualPx"`             //实际委托价
	ActualSide           string                                         `json:"actualSide"`           //实际触发方向
	TriggerTime          string                                         `json:"triggerTime"`          //策略委托触发时间，Unix时间戳的毫秒数格式，如 1597026383085
	PxVar                string                                         `json:"pxVar"`                //价格比例
	PxSpread             string                                         `json:"pxSpread"`             //价距
	SzLimit              string                                         `json:"szLimit"`              //单笔数量
	PxLimit              string                                         `json:"pxLimit"`              //挂单限制价
	Tag                  string                                         `json:"tag"`                  //订单标签
	TimeInterval         string                                         `json:"timeInterval"`         //下单间隔
	CallbackRatio        string                                         `json:"callbackRatio"`        //回调幅度的比例
	CallbackSpread       string                                         `json:"callbackSpread"`       //回调幅度的价距
	ActivePx             string                                         `json:"activePx"`             //移动止盈止损激活价格
	MoveTriggerPx        string                                         `json:"moveTriggerPx"`        //移动止盈止损触发价格
	ReduceOnly           string                                         `json:"reduceOnly"`           //是否只减仓
	QuickMgnType         string                                         `json:"quickMgnType"`         //一键借币类型，仅适用于杠杆逐仓的一键借币模式
	Last                 string                                         `json:"last"`                 //下单时的最新成交价
	FailCode             string                                         `json:"failCode"`             //代表策略触发失败的原因，已撤销和已生效时为""，委托失败时有值，如 51008；
	AlgoClOrdId          string                                         `json:"algoClOrdId"`          //客户自定义策略订单ID
	AmendPxOnTriggerType string                                         `json:"amendPxOnTriggerType"` //是否启用开仓价止损
	AttachAlgoOrds       []PrivateRestTradeOrderPendingRowAttachAlgoOrd `json:"attachAlgoOrds"`       //附带止盈止损信息
	LinkedOrds           PrivateRestTradeOrderPendingRowLinkedOrd       `json:"linkedOrds"`           //附带关联订单信息
	CTime                string                                         `json:"cTime"`                //订单创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime                string                                         `json:"uTime"`                //订单状态更新时间，Unix时间戳的毫秒数格式，如 1597026383085
}

type PrivateRestTradeOrderAlgoHistoryRowAttachAlgoOrd struct {
	AttachAlgoClOrdId string `json:"attachAlgoClOrdId"` //String	下单附带止盈止损时，客户自定义的策略订单ID
	TpTriggerPx       string `json:"tpTriggerPx"`       //String	止盈触发价
	TpTriggerPxType   string `json:"tpTriggerPxType"`   //String	止盈触发价类型
	TpOrdPx           string `json:"tpOrdPx"`           //String	止盈委托价
	SlTriggerPx       string `json:"slTriggerPx"`       //String	止损触发价
	SlTriggerPxType   string `json:"slTriggerPxType"`   //String	止损触发价类型
	SlOrdPx           string `json:"slOrdPx"`           //String	止损委托价
}
type PrivateRestTradeOrderAlgoHistoryRowLinkedOrd struct {
	OrdId string `json:"ordId"` //String	订单ID
}
type PrivateRestTradeOrderAlgoHistoryRes []PrivateRestTradeOrderAlgoPendingResRow
type PrivateRestTradeOrderAlgoHistoryResRow struct {
	InstType             string                                             `json:"instType"`             //产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	InstId               string                                             `json:"instId"`               //产品ID
	Ccy                  string                                             `json:"ccy"`                  //保证金币种，仅适用于单币种保证金模式下的全仓杠杆订单
	OrdId                string                                             `json:"ordId"`                //最新一笔订单ID，即将废弃。
	OrdIdList            []string                                           `json:"ordIdList,omitempty"`  //订单ID列表，当止盈止损存在市价拆单时，会有多个。
	AlgoId               string                                             `json:"algoId"`               //策略委托单ID
	ClOrdId              string                                             `json:"clOrdId"`              //客户自定义订单ID
	Sz                   string                                             `json:"sz"`                   //委托数量
	CloseFraction        string                                             `json:"closeFraction"`        //策略委托触发时，平仓的百分比。1 代表100%
	OrdType              string                                             `json:"ordType"`              //订单类型
	Side                 string                                             `json:"side"`                 //订单方向
	PosSide              string                                             `json:"posSide"`              //持仓方向
	TdMode               string                                             `json:"tdMode"`               //交易模式
	TgtCcy               string                                             `json:"tgtCcy"`               //币币市价单委托数量sz的单位
	State                string                                             `json:"state"`                //订单状态
	Lever                string                                             `json:"lever"`                //杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	TpTriggerPx          string                                             `json:"tpTriggerPx"`          //止盈触发价
	TpTriggerPxType      string                                             `json:"tpTriggerPxType"`      //止盈触发价类型
	TpOrdPx              string                                             `json:"tpOrdPx"`              //止盈委托价
	SlTriggerPx          string                                             `json:"slTriggerPx"`          //止损触发价
	SlTriggerPxType      string                                             `json:"slTriggerPxType"`      //止损触发价类型
	SlOrdPx              string                                             `json:"slOrdPx"`              //止损委托价
	TriggerPx            string                                             `json:"triggerPx"`            //计划委托触发价格
	TriggerPxType        string                                             `json:"triggerPxType"`        //计划委托触发价格类型
	OrdPx                string                                             `json:"ordPx"`                //订单委托价格
	ActualSz             string                                             `json:"actualSz"`             //实际委托量
	ActualPx             string                                             `json:"actualPx"`             //实际委托价
	ActualSide           string                                             `json:"actualSide"`           //实际触发方向
	TriggerTime          string                                             `json:"triggerTime"`          //策略委托触发时间，Unix时间戳的毫秒数格式，如 1597026383085
	PxVar                string                                             `json:"pxVar"`                //价格比例
	PxSpread             string                                             `json:"pxSpread"`             //价距
	SzLimit              string                                             `json:"szLimit"`              //单笔数量
	PxLimit              string                                             `json:"pxLimit"`              //挂单限制价
	Tag                  string                                             `json:"tag"`                  //订单标签
	TimeInterval         string                                             `json:"timeInterval"`         //下单间隔
	CallbackRatio        string                                             `json:"callbackRatio"`        //回调幅度的比例
	CallbackSpread       string                                             `json:"callbackSpread"`       //回调幅度的价距
	ActivePx             string                                             `json:"activePx"`             //移动止盈止损激活价格
	MoveTriggerPx        string                                             `json:"moveTriggerPx"`        //移动止盈止损触发价格
	ReduceOnly           string                                             `json:"reduceOnly"`           //是否只减仓
	QuickMgnType         string                                             `json:"quickMgnType"`         //一键借币类型，仅适用于杠杆逐仓的一键借币模式
	Last                 string                                             `json:"last"`                 //下单时的最新成交价
	FailCode             string                                             `json:"failCode"`             //代表策略触发失败的原因，已撤销和已生效时为""，委托失败时有值，如 51008；
	AlgoClOrdId          string                                             `json:"algoClOrdId"`          //客户自定义策略订单ID
	AmendPxOnTriggerType string                                             `json:"amendPxOnTriggerType"` //是否启用开仓价止损
	AttachAlgoOrds       []PrivateRestTradeOrderAlgoHistoryRowAttachAlgoOrd `json:"attachAlgoOrds"`       //附带止盈止损信息
	LinkedOrds           PrivateRestTradeOrderAlgoHistoryRowLinkedOrd       `json:"linkedOrds"`           //附带关联订单信息
	CTime                string                                             `json:"cTime"`                //订单创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime                string                                             `json:"uTime"`                //订单状态更新时间，Unix时间戳的毫秒数格式，如 1597026383085
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

type PrivateRestTradeClosePostionResRow struct {
	InstId  string `json:"instId"`  //String	产品ID
	PosSide string `json:"posSide"` //String	持仓方向
	ClOrdId string `json:"clOrdId"` //String	客户自定义ID
	Tag     string `json:"tag"`     //String	订单标签
}
type PrivateRestTradeClosePostionRes []PrivateRestTradeClosePostionResRow

type PrivateRestTradeOrderHistoryRes []PrivateRestTradeOrderGetResRow
type PrivateRestTradeOrderHistoryArchiveRes []PrivateRestTradeOrderGetResRow

type PrivateRestTradeFillsRes []PrivateRestTradeFillsResRow
type PrivateRestTradeFillsResRow struct {
	InstType    string `json:"instType"`    //String	产品类型
	InstId      string `json:"instId"`      //String	产品 ID
	TradeId     string `json:"tradeId"`     //String	最新成交 ID
	OrdId       string `json:"ordId"`       //String	订单 ID
	ClOrdId     string `json:"clOrdId"`     //String	用户自定义订单ID
	BillId      string `json:"billId"`      //String	账单 ID
	SubType     string `json:"subType"`     //String	成交类型
	Tag         string `json:"tag"`         //String	订单标签
	FillPx      string `json:"fillPx"`      //String	最新成交价格
	FillSz      string `json:"fillSz"`      //String	最新成交数量
	FillIdxPx   string `json:"fillIdxPx"`   //String	交易执行时的指数价格
	FillPnl     string `json:"fillPnl"`     //String	最新成交收益，适用于有成交的平仓订单。其他情况均为0。
	FillPxVol   string `json:"fillPxVol"`   //String	成交时的隐含波动率，仅适用于期权，其他业务线返回空字符串""
	FillPxUsd   string `json:"fillPxUsd"`   //String	成交时的期权价格，以USD为单位，仅适用于期权，其他业务线返回空字符串""
	FillMarkVol string `json:"fillMarkVol"` //String	成交时的标记波动率，仅适用于期权，其他业务线返回空字符串""
	FillFwdPx   string `json:"fillFwdPx"`   //String	成交时的远期价格，仅适用于期权，其他业务线返回空字符串""
	FillMarkPx  string `json:"fillMarkPx"`  //String	成交时的标记价格，仅适用于 交割/永续/期权
	Side        string `json:"side"`        //String	订单方向
	PosSide     string `json:"posSide"`     //String	持仓方向
	ExecType    string `json:"execType"`    //String	流动性方向
	FeeCcy      string `json:"feeCcy"`      //String	交易手续费币种或者返佣金币种
	Fee         string `json:"fee"`         //String	手续费金额或者返佣金额
	Ts          string `json:"ts"`          //String	成交明细产生时间，Unix时间戳的毫秒数格式，如 1597026383085
	FillTime    string `json:"fillTime"`    //String	成交时间，与订单频道的fillTime相同
}

type PrivateRestTradeFillsHistoryRes PrivateRestTradeFillsRes

// Asset
// 获取币种列表
type PrivateRestAssetCurrenciesRes []PrivateRestAssetCurrenciesResRow
type PrivateRestAssetCurrenciesResRow struct {
	Ccy                  string `json:"ccy"`                  //String	币种名称，如 BTC
	Name                 string `json:"name"`                 //String	币种名称，不显示则无对应名称
	LogoLink             string `json:"logoLink"`             //String	币种Logo链接
	Chain                string `json:"chain"`                //String	币种链信息
	CanDep               bool   `json:"canDep"`               //Boolean	当前是否可充值
	CanWd                bool   `json:"canWd"`                //Boolean	当前是否可提币
	CanInternal          bool   `json:"canInternal"`          //Boolean	当前是否可内部转账
	MinDep               string `json:"minDep"`               //String	币种单笔最小充值量
	MinWd                string `json:"minWd"`                //String	币种单笔最小链上提币量
	MaxWd                string `json:"maxWd"`                //String	币种单笔最大链上提币量
	WdTickSz             string `json:"wdTickSz"`             //String	提币精度,表示小数点后的位数。提币手续费精度与提币精度保持一致。
	WdQuota              string `json:"wdQuota"`              //String	过去24小时内提币额度（包含链上提币和内部转账），单位为USD
	UsedWdQuota          string `json:"usedWdQuota"`          //String	过去24小时内已用提币额度，单位为USD
	MinFee               string `json:"minFee"`               //String	普通地址最小提币手续费数量
	MaxFee               string `json:"maxFee"`               //String	普通地址最大提币手续费数量
	MinFeeForCtAddr      string `json:"minFeeForCtAddr"`      //String	合约地址最小提币手续费数量
	MaxFeeForCtAddr      string `json:"maxFeeForCtAddr"`      //String	合约地址最大提币手续费数量
	MainNet              bool   `json:"mainNet"`              //Boolean	当前链是否为主链
	NeedTag              bool   `json:"needTag"`              //Boolean	当前链是否需要标签（tag/memo）信息，如 EOS该字段为true
	MinDepArrivalConfirm string `json:"minDepArrivalConfirm"` //String	充值到账最小网络确认数。币已到账但不可提。
	MinWdUnlockConfirm   string `json:"minWdUnlockConfirm"`   //String	提现解锁最小网络确认数
	DepQuotaFixed        string `json:"depQuotaFixed"`        //String	充币固定限额，单位为USD
	UsedDepQuotaFixed    string `json:"usedDepQuotaFixed"`    //String	已用充币固定额度，单位为USD
	DepQuoteDailyLayer2  string `json:"depQuoteDailyLayer2"`  //String	Layer2网络每日充值上限
}

// 获取资金账户余额
type PrivateRestAssetBalancesRes []PrivateRestAssetBalancesResRow
type PrivateRestAssetBalancesResRow struct {
	Ccy       string `json:"ccy"`       //String	币种，如 BTC
	Bal       string `json:"bal"`       //String	余额
	FrozenBal string `json:"frozenBal"` //String	冻结余额
	AvailBal  string `json:"availBal"`  //String	可用余额
}

// 资金划转
type PrivateRestAssetTransferRes []PrivateRestAssetTransferResRow
type PrivateRestAssetTransferResRow struct {
	TransId  string `json:"transId"`  //String	划转ID
	Ccy      string `json:"ccy"`      //String	划转币种
	From     string `json:"from"`     //String	转出账户
	Amt      string `json:"amt"`      //String	划转量
	To       string `json:"to"`       //String	转入账户
	ClientId string `json:"clientId"` //String	客户自定义ID
}

// 资金划转状态查询
type PrivateRestAssetTransferStateRes []PrivateRestAssetTransferStateResRow
type PrivateRestAssetTransferStateResRow struct {
	TransId  string `json:"transId"`  //String	划转 ID
	ClientId string `json:"clientId"` //String	客户自定义 ID
	Ccy      string `json:"ccy"`      //String	划转币种
	Amt      string `json:"amt"`      //String	划转量
	Type     string `json:"type"`     //String	划转类型 0：账户内划转 1：母账户转子账户(仅适用于母账户APIKey) 2：子账户转母账户(仅适用于母账户APIKey) 3：子账户转母账户(仅适用于子账户APIKey) 4：子账户转子账户(仅适用于子账户APIKey，且目标账户需要是同一母账户下的其他子账户)
	From     string `json:"from"`     //String	转出账户 6：资金账户 18：交易账户
	To       string `json:"to"`       //String	转入账户 6：资金账户 18：交易账户
	SubAcct  string `json:"subAcct"`  //String	子账户名称
	State    string `json:"state"`    //String	转账状态 success：成功 pending：处理中 failed：失败
}

// 网格策略委托下单
type PrivateTradingBotGridOrderAlgoPostResRow struct {
	AlgoId      string `json:"algoId"`      //String	策略订单ID
	AlgoClOrdId string `json:"algoClOrdId"` //String	用户自定义策略ID
	SCode       string `json:"sCode"`       //String	事件执行结果的code，0代表成功
	SMsg        string `json:"sMsg"`        //String	事件执行失败时的msg
	Tag         string `json:"tag"`         //String	订单标签
}
type PrivateTradingBotGridOrderAlgoPostRes []PrivateTradingBotGridOrderAlgoPostResRow

// 定投策略委托下单
// algoId	String	策略订单ID
// algoClOrdId	String	客户自定义订单ID
// sCode	String	事件执行结果的code，0代表成功
// sMsg	String	事件执行失败时的msg
type PrivateTradingBotRecurringOrderAlgoPostResRow struct {
	AlgoId      string `json:"algoId"`      //String	策略订单ID
	AlgoClOrdId string `json:"algoClOrdId"` //String	客户自定义订单ID
	SCode       string `json:"sCode"`       //String	事件执行结果的code，0代表成功
	SMsg        string `json:"sMsg"`        //String	事件执行失败时的msg
}
type PrivateTradingBotRecurringOrderAlgoPostRes []PrivateTradingBotRecurringOrderAlgoPostResRow

type PrivateRestRfqCounterPartiesResRow struct {
	TraderName string `json:"traderName"` //String	报价方名称
	TraderCode string `json:"traderCode"` //String	报价方唯一标识代码，公开可见；报价和询价的相关接口都使用该代码代表报价方。
	Type       string `json:"type"`       //String	报价方类型。LP指通过API连接的自动做市商。
}
type PrivateRestRfqCounterPartiesRes []PrivateRestRfqCounterPartiesResRow

type PrivateRestRfqCreateRfqResRow struct {
	CTime                 string                       `json:"cTime"`                 //String	询价单创建时间，Unix时间戳的毫秒数格式。
	UTime                 string                       `json:"uTime"`                 //String	询价单状态更新时间，Unix时间戳的毫秒数格式。
	State                 string                       `json:"state"`                 //String	询价单的状态
	Counterparties        []string                     `json:"counterparties"`        //Array of strings	报价方列表
	ValidUntil            string                       `json:"validUntil"`            //String	询价单的过期时间，Unix时间戳的毫秒数格式。
	ClRfqId               string                       `json:"clRfqId"`               //String	询价单自定义ID，为客户端敏感信息，不会公开，对报价方返回""。
	Tag                   string                       `json:"tag"`                   //String	RFQ标签，与此RFQ关联的大宗交易将有相同的标签。
	AllowPartialExecution bool                         `json:"allowPartialExecution"` //Boolean	RFQ是否可以被部分执行，如果腿的比例和原RFQ一致。有效值为true或false。未指定时，默认为false。
	TraderCode            string                       `json:"traderCode"`            //String	询价方唯一标识代码。
	RfqId                 string                       `json:"rfqId"`                 //String	询价单ID
	Legs                  []PrivateRestRfqCreateRfqLeg `json:"legs"`                  //Array of objects	组合交易，每个请求最多可放置15条腿
}
type PrivateRestRfqCreateRfqRes []PrivateRestRfqCreateRfqResRow

type PrivateRestSprdOrderPostResRow struct {
	OrdId   string `json:"ordId"`   //String	订单ID
	ClOrdId string `json:"clOrdId"` //String	客户自定义订单ID
	Tag     string `json:"tag"`     //String	订单标签
	SCode   string `json:"sCode"`   //String	事件执行结果的code，0代表成功
	SMsg    string `json:"sMsg"`    //String	事件执行失败或成功时的msg
}
type PrivateRestSprdOrderPostRes []PrivateRestSprdOrderPostResRow
