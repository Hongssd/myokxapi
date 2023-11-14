package myokxapi

type PrivateRestAccountBalanceRes []PrivateRestAccountBalanceResRow
type PrivateRestAccountBalanceResRow struct {
	UTime       string                                  `json:"uTime"`       //账户信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	TotalEq     string                                  `json:"totalEq"`     //美金层面权益
	IsoEq       string                                  `json:"isoEq"`       //美金层面逐仓仓位权益 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	AdjEq       string                                  `json:"adjEq"`       //美金层面有效保证金 适用于跨币种保证金模式和组合保证金模式
	OrdFroz     string                                  `json:"ordFroz"`     //美金层面全仓挂单占用保证金 仅适用于跨币种保证金模式
	Imr         string                                  `json:"imr"`         //美金层面占用保证金 适用于跨币种保证金模式和组合保证金模式
	Mmr         string                                  `json:"mmr"`         //美金层面维持保证金 适用于跨币种保证金模式和组合保证金模式
	BorrowFroz  string                                  `json:"borrowFroz"`  //账户美金层面潜在借币占用保证金 仅适用于跨币种保证金模式和组合保证金模式. 在其他账户模式下为"".
	MgnRatio    string                                  `json:"mgnRatio"`    //美金层面保证金率 适用于跨币种保证金模式 和组合保证金模式
	NotionalUsd string                                  `json:"notionalUsd"` //以美金价值为单位的持仓数量，即仓位美金价值 适用于跨币种保证金模式和组合保证金模式
	Details     []PrivateRestAccountBalanceResRowDetail `json:"details"`     //各币种资产详细信息
}

type PrivateRestAccountBalanceResRowDetail struct {
	Ccy           string `json:"ccy"`           //币种
	Eq            string `json:"eq"`            //币种总权益
	CashBal       string `json:"cashBal"`       //币种余额
	UTime         string `json:"uTime"`         //币种余额信息的更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	IsoEq         string `json:"isoEq"`         //币种逐仓仓位权益 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	AvailEq       string `json:"availEq"`       //可用保证金 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	DisEq         string `json:"disEq"`         //美金层面币种折算权益
	FixedBal      string `json:"fixedBal"`      //币种冻结金额
	AvailBal      string `json:"availBal"`      //可用余额 适用于简单交易模式、单币种保证金模式、跨币种保证金模式和组合保证金模式
	FrozenBal     string `json:"frozenBal"`     //币种占用金额
	OrdFrozen     string `json:"ordFrozen"`     //挂单冻结数量
	Liab          string `json:"liab"`          //币种负债额 值为正数，如："21625.64"，适用于跨币种保证金模式和组合保证金模式
	Upl           string `json:"upl"`           //未实现盈亏 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	UplLiab       string `json:"uplLiab"`       //由于仓位未实现亏损导致的负债 适用于跨币种保证金模式和组合保证金模式
	CrossLiab     string `json:"crossLiab"`     //币种全仓负债额 适用于跨币种保证金模式和组合保证金模式
	IsoLiab       string `json:"isoLiab"`       //币种逐仓负债额 适用于跨币种保证金模式和组合保证金模式
	MgnRatio      string `json:"mgnRatio"`      //保证金率 适用于单币种保证金模式
	Interest      string `json:"interest"`      //计息，应扣未扣利息。 值为正数，如："9.01"，适用于跨币种保证金模式和组合保证金模式
	Twap          string `json:"twap"`          //当前负债币种触发系统自动换币的风险 0、1、2、3、4、5其中之一，数字越大代表您的负债币种触发自动换币概率越高 适用于跨币种保证金模式和组合保证金模式
	MaxLoan       string `json:"maxLoan"`       //币种最大可借 适用于跨币种保证金模式和组合保证金模式 的全仓
	EqUsd         string `json:"eqUsd"`         //币种权益美金价值
	BorrowFroz    string `json:"borrowFroz"`    //币种美金层面潜在借币占用保证金 仅适用于跨币种保证金模式和组合保证金模式. 在其他账户模式下为"".
	NotionalLever string `json:"notionalLever"` //币种杠杆倍数 适用于单币种保证金模式
	StgyEq        string `json:"stgyEq"`        //策略权益
	IsoUpl        string `json:"isoUpl"`        //逐仓未实现盈亏 适用于单币种保证金模式和跨币种保证金模式和组合保证金模式
	SpotInUseAmt  string `json:"spotInUseAmt"`  //现货对冲占用数量 适用于组合保证金模式
	SpotIsoBal    string `json:"spotIsoBal"`    //现货逐仓余额，仅适用于现货带单/跟单
}

type PrivateRestAccountPositionRes []PrivateRestAccountPositionResRow
type PrivateRestAccountPositionResRow struct {
	InstType       string                                           `json:"instType"`       //产品类型
	MgnMode        string                                           `json:"mgnMode"`        //保证金模式 cross：全仓 isolated：逐仓
	PosId          string                                           `json:"posId"`          //持仓ID
	PosSide        string                                           `json:"posSide"`        //持仓方向 long：开平仓模式开多，pos为正 short：开平仓模式开空，pos为正 net：买卖模式（交割/永续/期权：pos为正代表开多，pos为负代表开空。币币杠杆时，pos均为正，posCcy为交易货币时，代表开多；posCcy为计价货币时，代表开空。）
	Pos            string                                           `json:"pos"`            //持仓数量，逐仓自主划转模式下，转入保证金后会产生pos为0的仓位
	BaseBal        string                                           `json:"baseBal"`        //交易币余额，适用于 币币杠杆（逐仓自主划转模式 和 一键借币模式）
	QuoteBal       string                                           `json:"quoteBal"`       //计价币余额 ，适用于 币币杠杆（逐仓自主划转模式 和 一键借币模式）
	BaseBorrowed   string                                           `json:"baseBorrowed"`   //交易币已借，适用于 币币杠杆（逐仓一键借币模式）
	BaseInterest   string                                           `json:"baseInterest"`   //交易币计息，适用于 币币杠杆（逐仓一键借币模式）
	QuoteBorrowed  string                                           `json:"quoteBorrowed"`  //计价币已借，适用于 币币杠杆（逐仓一键借币模式）
	QuoteInterest  string                                           `json:"quoteInterest"`  //计价币计息，适用于 币币杠杆（逐仓一键借币模式）
	PosCcy         string                                           `json:"posCcy"`         //仓位资产币种，仅适用于币币杠杆仓位
	AvailPos       string                                           `json:"availPos"`       //可平仓数量，适用于 币币杠杆,交割/永续（开平仓模式），期权 对于杠杆仓位，平仓时，杠杆还清负债后，余下的部分会视为币币交易，如果想要减少币币交易的数量，可通过"获取最大可用数量"接口获取只减仓的可用数量。
	AvgPx          string                                           `json:"avgPx"`          //开仓平均价
	Upl            string                                           `json:"upl"`            //未实现收益（以标记价格计算）
	UplRatio       string                                           `json:"uplRatio"`       //未实现收益率（以标记价格计算
	UplLastPx      string                                           `json:"uplLastPx"`      //以最新成交价格计算的未实现收益，主要做展示使用，实际值还是 upl
	UplRatioLastPx string                                           `json:"uplRatioLastPx"` //以最新成交价格计算的未实现收益率
	InstId         string                                           `json:"instId"`         //产品ID，如 BTC-USD-180216
	Lever          string                                           `json:"lever"`          //杠杆倍数，不适用于期权
	LiqPx          string                                           `json:"liqPx"`          //预估强平价 不适用于期权
	MarkPx         string                                           `json:"markPx"`         //最新标记价格
	Imr            string                                           `json:"imr"`            //初始保证金，仅适用于全仓
	Margin         string                                           `json:"margin"`         //保证金余额，可增减，仅适用于逐仓
	MgnRatio       string                                           `json:"mgnRatio"`       //保证金率
	Mmr            string                                           `json:"mmr"`            //维持保证金
	Liab           string                                           `json:"liab"`           //负债额，仅适用于币币杠杆
	LiabCcy        string                                           `json:"liabCcy"`        //负债币种，仅适用于币币杠杆
	Interest       string                                           `json:"interest"`       //利息，已经生成的未扣利息
	TradeId        string                                           `json:"tradeId"`        //最新成交ID
	OptVal         string                                           `json:"optVal"`         //期权市值，仅适用于期权
	NotionalUsd    string                                           `json:"notionalUsd"`    //以美金价值为单位的持仓数量
	Adl            string                                           `json:"adl"`            //信号区 分为5档，从1到5，数字越小代表adl强度越弱
	Ccy            string                                           `json:"ccy"`            //占用保证金的币种
	Last           string                                           `json:"last"`           //最新成交价
	IdxPx          string                                           `json:"idxPx"`          //最新指数价格
	UsdPx          string                                           `json:"usdPx"`          //美金价格
	BePx           string                                           `json:"bePx"`           //盈亏平衡价
	DeltaBS        string                                           `json:"deltaBS"`        //美金本位持仓仓位delta，仅适用于期权
	DeltaPA        string                                           `json:"deltaPA"`        //币本位持仓仓位delta，仅适用于期权
	GammaBS        string                                           `json:"gammaBS"`        //美金本位持仓仓位gamma，仅适用于期权
	GammaPA        string                                           `json:"gammaPA"`        //币本位持仓仓位gamma，仅适用于期权
	ThetaBS        string                                           `json:"thetaBS"`        //美金本位持仓仓位theta，仅适用于期权
	ThetaPA        string                                           `json:"thetaPA"`        //币本位持仓仓位theta，仅适用于期权
	VegaBS         string                                           `json:"vegaBS"`         //美金本位持仓仓位vega，仅适用于期权
	VegaPA         string                                           `json:"vegaPA"`         //币本位持仓仓位vega，仅适用于期权
	CTime          string                                           `json:"cTime"`          //持仓创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime          string                                           `json:"uTime"`          //最近一次持仓更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	SpotInUseAmt   string                                           `json:"spotInUseAmt"`   //现货对冲占用数量 适用于组合保证金模式
	SpotInUseCcy   string                                           `json:"spotInUseCcy"`   //现货对冲占用币种，如 BTC 适用于组合保证金模式
	RealizedPnl    string                                           `json:"realizedPnl"`    //已实现收益
	Pnl            string                                           `json:"pnl"`            //平仓订单累计收益额
	Fee            string                                           `json:"fee"`            //累计手续费金额，正数代表平台返佣 ，负数代表平台扣除
	FundingFee     string                                           `json:"fundingFee"`     //累计资金费用
	LiqPenalty     string                                           `json:"liqPenalty"`     //累计爆仓罚金，有值时为负数。
	CloseOrderAlgo []PrivateRestAccountPositionResRowCloseOrderAlgo `json:"closeOrderAlgo"` //平仓策略委托订单。调用策略委托下单，且closeFraction=1 时，该数组才会有值。
	BizRefId       string                                           `json:"bizRefId"`       //外部业务id，e.g. 体验券id
	BizRefType     string                                           `json:"bizRefType"`     //外部业务类型

}

type PrivateRestAccountPositionResRowCloseOrderAlgo struct {
	AlgoId          string `json:"algoId"`          //策略委托单ID
	SlTriggerPx     string `json:"slTriggerPx"`     //止损触发价
	SlTriggerPxType string `json:"slTriggerPxType"` //止损触发价类型 last：最新价格 index：指数价格 mark：标记价格
	TpTriggerPx     string `json:"tpTriggerPx"`     //止盈委托价
	TpTriggerPxType string `json:"tpTriggerPxType"` //止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格
	CloseFraction   string `json:"closeFraction"`   //策略委托触发时，平仓的百分比。1 代表100%
}

type PrivateRestAccountConfigRes []PrivateRestAccountConfigResRow
type PrivateRestAccountConfigResRow struct {
	Uid             string   `json:"uid"`             //当前请求的账户ID，账户uid和app上的一致
	MainUid         string   `json:"mainUid"`         //当前请求的母账户ID 如果 uid = mainUid，代表当前账号为母账户；如果 uid != mainUid，代表当前账户为子账户。
	AcctLv          string   `json:"acctLv"`          //账户层级 1：简单交易模式 2：单币种保证金模式 3：跨币种保证金模式 4：组合保证金模式
	PosMode         string   `json:"posMode"`         //持仓方式 long_short_mode：开平仓模式 net_mode：买卖模式 仅适用交割/永续
	AutoLoan        bool     `json:"autoLoan"`        //是否自动借币 true：自动借币 false：非自动借币
	GreeksType      string   `json:"greeksType"`      //当前希腊字母展示方式 PA：币本位 BS：美元本位
	Level           string   `json:"level"`           //当前在平台上真实交易量的用户等级，例如 lv1
	LevelTmp        string   `json:"levelTmp"`        //特约用户的临时体验用户等级，例如 lv3
	CtIsoMode       string   `json:"ctIsoMode"`       //衍生品的逐仓保证金划转模式 automatic：开仓划转 autonomy：自主划转
	MgnIsoMode      string   `json:"mgnIsoMode"`      //币币杠杆的逐仓保证金划转模式 automatic：开仓划转 quick_margin：一键借币（对于新的账户，包括新的子账户，有些默认是开仓划转，另外的默认是一键借币）
	SpotOffsetType  string   `json:"spotOffsetType"`  //现货对冲类型 1：现货对冲模式U模式 2：现货对冲模式币模式 3：非现货对冲模式 适用于组合保证金模式
	RoleType        string   `json:"roleType"`        //用户角色 0：普通用户 1：带单者 2：跟单者
	TraderInsts     []string `json:"traderInsts"`     //当前账号已经设置的带单合约，仅适用于带单者
	SpotRoleType    string   `json:"spotRoleType"`    //现货跟单角色。 0：普通用户；1：带单者；2：跟单者
	SpotTraderInsts []string `json:"spotTraderInsts"` //当前账号已经设置的带单合约，仅适用于带单者
	OpAuth          string   `json:"opAuth"`          //是否开通期权交易 0：未开通 1：已经开通
	KycLv           string   `json:"kycLv"`           //母账户KYC等级 0: 未认证 1: 已完成 level 1 认证 2: 已完成 level 2 认证 3: 已完成 level 3认证 如果请求来自子账户, kycLv 为其母账户的等级 如果请求来自母账户, kycLv 为当前请求的母账户等级
	Label           string   `json:"label"`           //当前请求API key的备注名，不超过50位字母（区分大小写）或数字，可以是纯字母或纯数字。
	Ip              string   `json:"ip"`              //当前请求API key绑定的ip地址，多个ip用半角逗号隔开，如：
	Perm            string   `json:"perm"`            //当前请求的 API key权限 read_only：读取 trade：交易 withdraw：提币
}

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
