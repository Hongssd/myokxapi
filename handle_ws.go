package myokxapi

type WsCandles struct {
	WsSubscribeArg        //订阅信息
	Interval       string `json:"interval"`    //K线周期，如 1m，5m，15m，30m，1h，2h，4h，6h，12h，1d，1w，1M
	Ts             string `json:"ts"`          //开始时间，Unix时间戳的毫秒数格式，如 1597026383085
	O              string `json:"o"`           //开盘价格
	H              string `json:"h"`           //最高价格
	L              string `json:"l"`           //最低价格
	C              string `json:"c"`           //收盘价格
	Vol            string `json:"vol"`         //交易量，以张为单位
	VolCcy         string `json:"volCcy"`      //交易量，以币为单位
	VolCcyQuote    string `json:"volCcyQuote"` //交易量，以计价货币为单位
	Confirm        string `json:"confirm"`     //K线状态 0 代表 K 线未完结，1 代表 K 线已完结。
}

type WsCandlesMiddle struct {
	Arg  WsSubscribeArg   `json:"arg"`  //订阅信息
	Data [][9]interface{} `json:"data"` //K线数据
}

func handleWsCandle(data []byte) (*WsCandles, error) {

	wsCandlesMiddle := WsCandlesMiddle{}
	err := json.Unmarshal(data, &wsCandlesMiddle)
	if err != nil {
		return nil, err
	}
	candelData := wsCandlesMiddle.Data[0]
	candle := WsCandles{
		WsSubscribeArg: wsCandlesMiddle.Arg,
		Interval:       wsCandlesMiddle.Arg.Channel[6:],
		Ts:             candelData[0].(string),
		O:              candelData[1].(string),
		H:              candelData[2].(string),
		L:              candelData[3].(string),
		C:              candelData[4].(string),
		Vol:            candelData[5].(string),
		VolCcy:         candelData[6].(string),
		VolCcyQuote:    candelData[7].(string),
		Confirm:        candelData[8].(string),
	}
	return &candle, nil
}

type WsBooks struct {
	WsSubscribeArg         //订阅信息
	Action         string  `json:"action"`    //推送数据动作，增量推送数据还是全量推送数据 snapshot：全量 update：增量
	Asks           []Books `json:"asks"`      //卖方深度
	Bids           []Books `json:"bids"`      //买方深度
	Ts             string  `json:"ts"`        //深度产生的时间
	CheckSum       int64   `json:"checksum"`  //检验和
	PrevSeqId      int64   `json:"prevSeqId"` //上一个推送的序列号。仅适用 books，books-l2-tbt，books50-l2-tbt
	SeqId          int64   `json:"seqId"`     //推送的序列号
}

type WsBooksMiddle struct {
	Arg    WsSubscribeArg `json:"arg"`    //订阅信息
	Action string         `json:"action"` //推送数据动作，增量推送数据还是全量推送数据 snapshot：全量 update：增量
	Data   []struct {
		Asks      [][4]interface{} `json:"asks"`      //卖方深度
		Bids      [][4]interface{} `json:"bids"`      //买方深度
		Ts        string           `json:"ts"`        //深度产生的时间
		CheckSum  int64            `json:"checksum"`  //检验和
		PrevSeqId int64            `json:"prevSeqId"` //上一个推送的序列号。仅适用 books，books-l2-tbt，books50-l2-tbt
		SeqId     int64            `json:"seqId"`     //推送的序列号
	} `json:"data"` //深度数据
}

func handleWsBooks(data []byte) (*WsBooks, error) {

	wsBooksMiddle := WsBooksMiddle{}
	err := json.Unmarshal(data, &wsBooksMiddle)
	if err != nil {
		return nil, err
	}
	middleRow := wsBooksMiddle.Data[0]

	wsBook := WsBooks{
		WsSubscribeArg: wsBooksMiddle.Arg,
		Action:         wsBooksMiddle.Action,
		Asks:           []Books{},
		Bids:           []Books{},
		Ts:             middleRow.Ts,
		CheckSum:       middleRow.CheckSum,
		PrevSeqId:      middleRow.PrevSeqId,
		SeqId:          middleRow.SeqId,
	}

	for _, ask := range middleRow.Asks {
		wsBook.Asks = append(wsBook.Asks, Books{
			Price:      ask[0].(string),
			Quantity:   ask[1].(string),
			OrderCount: ask[3].(string),
		})
	}
	for _, bid := range middleRow.Bids {
		wsBook.Bids = append(wsBook.Bids, Books{
			Price:      bid[0].(string),
			Quantity:   bid[1].(string),
			OrderCount: bid[3].(string),
		})
	}

	return &wsBook, nil
}

type WsTrades struct {
	WsSubscribeArg //订阅信息
	Trades
}

type Trades struct {
	InstId  string `json:"instId"`  //产品ID，如 BTC-USD-180216
	TradeId string `json:"tradeId"` //聚合的多笔交易中最新一笔交易的成交ID
	Px      string `json:"px"`      //成交价格
	Sz      string `json:"sz"`      //成交数量
	Side    string `json:"side"`    //成交方向，buy sell
	Ts      string `json:"ts"`      //成交时间，Unix时间戳的毫秒数格式，如 1597026383085
	Count   string `json:"count"`   //聚合的订单匹配数量
}

type WsTradesMiddle struct {
	Arg  WsSubscribeArg `json:"arg"`
	Data []Trades       `json:"data"`
}

func handleWsTrades(data []byte) (*WsTrades, error) {

	wsTradesMiddle := WsTradesMiddle{}
	err := json.Unmarshal(data, &wsTradesMiddle)
	if err != nil {
		return nil, err
	}
	trades := wsTradesMiddle.Data[0]
	wsTrades := WsTrades{
		WsSubscribeArg: wsTradesMiddle.Arg,
		Trades: Trades{
			InstId:  trades.InstId,
			TradeId: trades.TradeId,
			Px:      trades.Px,
			Sz:      trades.Sz,
			Side:    trades.Side,
			Ts:      trades.Ts,
			Count:   trades.Count,
		},
	}
	return &wsTrades, nil
}

type WsOrders struct {
	WsSubscribeArg
	Orders
}

type Orders struct {
	InstType          string          `json:"instType"`          //String	产品类型
	InstId            string          `json:"instId"`            //String	产品ID
	Ccy               string          `json:"ccy"`               //String	保证金币种，仅适用于单币种保证金账户下的全仓币币杠杆订单
	OrdId             string          `json:"ordId"`             //String	订单ID
	ClOrdId           string          `json:"clOrdId"`           //String	由用户设置的订单ID来识别您的订单
	Tag               string          `json:"tag"`               //String	订单标签
	Px                string          `json:"px"`                //String	委托价格，对于期权，以币(如BTC, ETH)为单位
	PxUsd             string          `json:"pxUsd"`             //String	期权价格，以USD为单位 仅适用于期权，其他业务线返回空字符串""
	PxVol             string          `json:"pxVol"`             //String	期权订单的隐含波动率 仅适用于期权，其他业务线返回空字符串""
	PxType            string          `json:"pxType"`            //String	期权的价格类型 px：代表按价格下单，单位为币 (请求参数 px 的数值单位是BTC或ETH) pxVol：代表按pxVol下单 pxUsd：代表按照pxUsd下单，单位为USD (请求参数px 的数值单位是USD)
	Sz                string          `json:"sz"`                //String	原始委托数量，币币/币币杠杆，以币为单位；交割/永续/期权 ，以张为单位
	NotionalUsd       string          `json:"notionalUsd"`       //String	委托单预估美元价值
	FillNotionalUsd   string          `json:"fillNotionalUsd"`   //String	委托单已成交的美元价值
	OrdType           string          `json:"ordType"`           //String	订单类型 market：市价单 limit：限价单 post_only： 只做maker单 fok：全部成交或立即取消单 ioc：立即成交并取消剩余单 optimal_limit_ioc：市价委托立即成交并取消剩余（仅适用交割、永续） mmp：做市商保护(仅适用于组合保证金账户模式下的期权订单) mmp_and_post_only：做市商保护且只做maker单(仅适用于组合保证金账户模式下的期权订单) op_fok：期权简选（全部成交或立即取消）
	Side              string          `json:"side"`              //String	订单方向，buy sell
	PosSide           string          `json:"posSide"`           //String	持仓方向 long：开平仓模式开多 short：开平仓模式开空 net：买卖模式
	TdMode            string          `json:"tdMode"`            //String	交易模式 保证金模式 isolated：逐仓 cross：全仓 非保证金模式 cash：现金
	TgtCcy            string          `json:"tgtCcy"`            //String	市价单委托数量sz的单位 base_ccy: 交易货币 quote_ccy：计价货币
	FillPx            string          `json:"fillPx"`            //String	最新成交价格
	TradeId           string          `json:"tradeId"`           //String	最新成交ID
	FillSz            string          `json:"fillSz"`            //String	最新成交数量 对于币币和杠杆，单位为交易货币，如 BTC-USDT, 单位为 BTC；对于市价单，无论tgtCcy是base_ccy，还是quote_ccy，单位均为交易货币； 对于交割、永续以及期权，单位为张。
	FillPnl           string          `json:"fillPnl"`           //String	最新成交收益，适用于有成交的平仓订单。其他情况均为0。
	FillTime          string          `json:"fillTime"`          //String	最新成交时间
	FillFee           string          `json:"fillFee"`           //String	最新一笔成交的手续费金额或者返佣金额： 手续费扣除 为 ‘负数’，如 -0.01 ； 手续费返佣 为 ‘正数’，如 0.01
	FillFeeCcy        string          `json:"fillFeeCcy"`        //String	最新一笔成交的手续费币种或者返佣币种。 如果fillFee小于0，为手续费币种；如果fillFee大于等于0，为返佣币种
	FillPxVol         string          `json:"fillPxVol"`         //String	成交时的隐含波动率仅适用于期权，其他业务线返回空字符串""
	FillPxUsd         string          `json:"fillPxUsd"`         //String	成交时的期权价格，以USD为单位仅适用于期权，其他业务线返回空字符串""
	FillMarkVol       string          `json:"fillMarkVol"`       //String	成交时的标记波动率，仅适用于期权，其他业务线返回空字符串""
	FillFwdPx         string          `json:"fillFwdPx"`         //String	成交时的远期价格，仅适用于期权，其他业务线返回空字符串""
	FillMarkPx        string          `json:"fillMarkPx"`        //String	成交时的标记价格，仅适用于 交割/永续/期权
	ExecType          string          `json:"execType"`          //String	最新一笔成交的流动性方向 T：taker M：maker
	AccFillSz         string          `json:"accFillSz"`         //String	累计成交数量 对于币币和杠杆，单位为交易货币，如 BTC-USDT, 单位为 BTC；对于市价单，无论tgtCcy是base_ccy，还是quote_ccy，单位均为交易货币； 对于交割、永续以及期权，单位为张。
	AvgPx             string          `json:"avgPx"`             //String	成交均价，如果成交数量为0，该字段也为0
	State             string          `json:"state"`             //String	订单状态 canceled：撤单成功 live：等待成交 partially_filled： 部分成交 filled：完全成交 mmp_canceled：做市商保护机制导致的自动撤单
	Lever             string          `json:"lever"`             //String	杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	AttachAlgoClOrdId string          `json:"attachAlgoClOrdId"` //String	下单附带止盈止损时，客户自定义的策略订单ID
	TpTriggerPx       string          `json:"tpTriggerPx"`       //String	止盈触发价
	TpTriggerPxType   string          `json:"tpTriggerPxType"`   //String	止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格
	TpOrdPx           string          `json:"tpOrdPx"`           //String	止盈委托价，止盈委托价格为-1时，执行市价止盈
	SlTriggerPx       string          `json:"slTriggerPx"`       //String	止损触发价
	SlTriggerPxType   string          `json:"slTriggerPxType"`   //String	止损触发价类型 last：最新价格 index：指数价格 mark：标记价格
	SlOrdPx           string          `json:"slOrdPx"`           //String	止损委托价，止损委托价格为-1时，执行市价止损
	AttachAlgoOrds    []AttachAlgoOrd `json:"attachAlgoOrds"`    //Array	下单附带止盈止损信息
	StpId             string          `json:"stpId"`             //String	自成交保护ID 如果自成交保护不适用则返回""
	StpMode           string          `json:"stpMode"`           //String	自成交保护模式 如果自成交保护不适用则返回""
	FeeCcy            string          `json:"feeCcy"`            //String	交易手续费币种 币币/币币杠杆：如果是买的话，收取的就是BTC；如果是卖的话，收取的就是USDT 交割/永续/期权 收取的就是保证金
	Fee               string          `json:"fee"`               //String	订单交易累计的手续费与返佣 对于币币和杠杆，为订单交易累计的手续费，平台向用户收取的交易手续费，为负数。如： -0.01 对于交割、永续和期权，为订单交易累计的手续费和返佣
	RebateCcy         string          `json:"rebateCcy"`         //String	返佣金币种 ，如果没有返佣金，该字段为“”
	Rebate            string          `json:"rebate"`            //String	返佣累计金额，仅适用于币币和杠杆，平台向达到指定lv交易等级的用户支付的挂单奖励（返佣），如果没有返佣金，该字段为“”
	Pnl               string          `json:"pnl"`               //String	收益，适用于有成交的平仓订单，其他情况均为0
	Source            string          `json:"source"`            //String	订单来源 6：计划委托策略触发后的生成的普通单 7：止盈止损策略触发后的生成的普通单 13：策略委托单触发后的生成的普通单 24：移动止盈止损策略触发后的生成的普通单
	CancelSource      string          `json:"cancelSource"`      //String	订单取消的来源 有效值及对应的含义是： 0: 已撤单：系统撤单 1: 用户主动撤单 2: 已撤单：预减仓撤单，用户保证金不足导致挂单被撤回 3: 已撤单：风控撤单，用户保证金不足有爆仓风险，导致挂单被撤回 4: 已撤单：币种借币量达到平台硬顶，系统已撤回该订单 6: 已撤单：触发 ADL 撤单，用户保证金率较低且有爆仓风险，导致挂单被撤回 7: 已撤单：交割合约到期 9: 已撤单：扣除资金费用后可用余额不足，系统已撤回该订单 13: 已撤单：FOK 委托订单未完全成交，导致挂单被完全撤回 14: 已撤单：IOC 委托订单未完全成交，仅部分成交，导致部分挂单被撤回 17: 已撤单：平仓单被撤单，由于仓位已被市价全平 20: 系统倒计时撤单 21: 已撤单：相关仓位被完全平仓，系统已撤销该止盈止损订单 22, 23: 已撤单：只减仓订单仅允许减少仓位数量，系统已撤销该订单 27: 成交滑点超过5%，触发成交差价保护导致系统撤单 31: 当前只挂单订单 (Post only) 将会吃掉挂单深度 32: 自成交保护 33: 当前 taker 订单匹配的订单数量超过最大限制
	AmendSource       string          `json:"amendSource"`       //String	订单修改的来源 1: 用户主动改单，改单成功 2: 用户主动改单，并且当前这笔订单被只减仓修改，改单成功 3: 用户主动下单，并且当前这笔订单被只减仓修改，改单成功 4: 用户当前已存在的挂单（非当前操作的订单），被只减仓修改，改单成功 5. 期权 px, pxVol 或 pxUsd 的跟随变动导致的改单，比如 iv=60，usd，px 锚定iv=60 时，usd, px 产生变动时的改单
	Category          string          `json:"category"`          //String	订单种类分类 normal：普通委托订单种类 twap：TWAP订单种类 adl：ADL订单种类 full_liquidation：爆仓订单种类 partial_liquidation：减仓订单种类 delivery：交割 ddh：对冲减仓类型订单
	UTime             string          `json:"uTime"`             //String	订单更新时间，Unix时间戳的毫秒数格式，如 1597026383085
	CTime             string          `json:"cTime"`             //String	订单创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	ReqId             string          `json:"reqId"`             //String	修改订单时使用的request ID，如果没有修改，该字段为""
	AmendResult       string          `json:"amendResult"`       //String	修改订单的结果 -1： 失败 0：成功 1：自动撤单（因为修改成功导致订单自动撤销） 2: 自动改单成功，仅适用于期权pxUsd和pxVol订单的自动改单 通过API修改订单时，如果cxlOnFail设置为true且修改返回结果为失败时，则返回 "" 通过API修改订单时，如果修改返回结果为成功但修改最终失败后，当cxlOnFail设置为false时返回 -1;当cxlOnFail设置为true时则返回1 通过Web/APP修改订单时，如果修改失败后，则返回-1
	ReduceOnly        string          `json:"reduceOnly"`        //String	是否只减仓，true 或 false
	QuickMgnType      string          `json:"quickMgnType"`      //String	一键借币类型，仅适用于杠杆逐仓的一键借币模式 manual：手动，auto_borrow： 自动借币，auto_repay： 自动还币
	AlgoClOrdId       string          `json:"algoClOrdId"`       //String	客户自定义策略订单ID。策略订单触发，且策略单有algoClOrdId时有值，否则为"",
	AlgoId            string          `json:"algoId"`            //String	策略委托单ID，策略订单触发时有值，否则为""
	LastPx            string          `json:"lastPx"`            //String	最新成交价
	Code              string          `json:"code"`              //String	错误码，默认为0
	Msg               string          `json:"msg"`               //String	错误消息，默认为""
}

type AttachAlgoOrd struct {
	OrdId                string `json:"ordId"`                //String	附带止盈止损的订单ID，订单完全成交，下止盈止损委托单时，该值不会传给 algoId
	AttachAlgoClOrdId    string `json:"attachAlgoClOrdId"`    //String	下单附带止盈止损时，客户自定义的策略订单ID
	TpTriggerPx          string `json:"tpTriggerPx"`          //String	止盈触发价
	TpTriggerPxType      string `json:"tpTriggerPxType"`      //String	止盈触发价类型 last：最新价格 index：指数价格 mark：标记价格
	TpOrdPx              string `json:"tpOrdPx"`              //String	止盈委托价
	SlTriggerPx          string `json:"slTriggerPx"`          //String	止损触发价
	SlTriggerPxType      string `json:"slTriggerPxType"`      //String	止损触发价类型 last：最新价格 index：指数价格 mark：标记价格
	SlOrdPx              string `json:"slOrdPx"`              //String	止损委托价
	Sz                   string `json:"sz"`                   //String	张数。仅适用于“多笔止盈”的止盈订单且必填
	AmendPxOnTriggerType string `json:"amendPxOnTriggerType"` //String	是否启用开仓价止损，仅适用于分批止盈的止损订单 0. 不开启，默认值 1. 开启“开仓价止损”
}

type WsOrdersMiddle struct {
	Arg  WsSubscribeArg `json:"arg"`
	Data []Orders       `json:"data"`
}

func handleWsOrders(data []byte) (*[]WsOrders, error) {

	wsOrdersMiddle := WsOrdersMiddle{}
	err := json.Unmarshal(data, &wsOrdersMiddle)
	if err != nil {
		return nil, err
	}

	orders := wsOrdersMiddle.Data
	wsOrdersList := []WsOrders{}
	for _, order := range orders {
		wsOrder := WsOrders{
			WsSubscribeArg: wsOrdersMiddle.Arg,
			Orders:         order,
		}
		wsOrdersList = append(wsOrdersList, wsOrder)
	}
	return &wsOrdersList, nil
}

type LinkedOrd struct {
	OrdId string `json:"ordId"` //String	关联订单ID
}
type WsOrdersAlgo struct {
	WsSubscribeArg
	OrdersAlgo
}
type OrdersAlgo struct {
	InstType             string          `json:"instType"`             //String	产品类型
	InstId               string          `json:"instId"`               //String	产品ID
	Ccy                  string          `json:"ccy"`                  //String	保证金币种，仅单币种保证金账户下的全仓币币杠杆需要选择保证金币种
	OrdId                string          `json:"ordId"`                //String	最新一笔订单ID，与策略委托订单关联的订单ID，即将废弃。
	OrdIdList            []string        `json:"ordIdList"`            //String	订单ID列表，当止盈止损存在市价拆单时，会有多个。
	AlgoId               string          `json:"algoId"`               //String	策略委托单ID
	ClOrdId              string          `json:"clOrdId"`              //String	客户自定义订单ID
	Sz                   string          `json:"sz"`                   //String	委托数量，币币/币币杠杆 以币为单位；交割/永续/期权 以张为单位
	OrdType              string          `json:"ordType"`              //String	订单类型
	Side                 string          `json:"side"`                 //String	订单方向，buy sell
	PosSide              string          `json:"posSide"`              //String	持仓方向
	TdMode               string          `json:"tdMode"`               //String	交易模式
	TgtCcy               string          `json:"tgtCcy"`               //String	币币市价单委托数量sz的单位
	Lever                string          `json:"lever"`                //String	杠杆倍数，0.01到125之间的数值，仅适用于 币币杠杆/交割/永续
	State                string          `json:"state"`                //String	订单状态
	TpTriggerPx          string          `json:"tpTriggerPx"`          //String	止盈触发价
	TpTriggerPxType      string          `json:"tpTriggerPxType"`      //String	止盈触发价类型
	TpOrdPx              string          `json:"tpOrdPx"`              //String	止盈委托价，委托价格为-1时，执行市价止盈
	SlTriggerPx          string          `json:"slTriggerPx"`          //String	止损触发价
	SlTriggerPxType      string          `json:"slTriggerPxType"`      //String	止损触发价类型
	SlOrdPx              string          `json:"slOrdPx"`              //String	止损委托价委托价格为-1时，执行市价止损
	TriggerPx            string          `json:"triggerPx"`            //String	计划委托单的触发价格
	TriggerPxType        string          `json:"triggerPxType"`        //String	计划委托单的触发价类型
	OrdPx                string          `json:"ordPx"`                //String	计划委托单的委托价格
	Last                 string          `json:"last"`                 //String	下单时的最新成交价
	ActualSz             string          `json:"actualSz"`             //String	实际委托量
	ActualPx             string          `json:"actualPx"`             //String	实际委价
	Tag                  string          `json:"tag"`                  //String	订单标签
	NotionalUsd          string          `json:"notionalUsd"`          //String	委托单预估美元价值
	ActualSide           string          `json:"actualSide"`           //String	实际触发方向
	TriggerTime          string          `json:"triggerTime"`          //String	策略委托触发时间，Unix时间戳的毫秒数格式，如 1597026383085
	ReduceOnly           string          `json:"reduceOnly"`           //String	是否只减仓，true 或 false
	FailCode             string          `json:"failCode"`             //String	代表策略触发失败的原因，已撤销和已生效时为""，委托失败时有值，如 51008；
	AlgoClOrdId          string          `json:"algoClOrdId"`          //String	客户自定义策略订单ID
	ReqId                string          `json:"reqId"`                //String	修改订单时使用的request ID，如果没有修改，该字段为""
	AmendResult          string          `json:"amendResult"`          //String	修改订单的结果
	AmendPxOnTriggerType string          `json:"amendPxOnTriggerType"` //String	是否启用开仓价止损，仅适用于分批止盈的止损订单
	AttachAlgoOrds       []AttachAlgoOrd `json:"attachAlgoOrds"`       //Array	附带止盈止损信息
	LinkedOrd            LinkedOrd       `json:"linkedOrd"`            //Object	止盈订单信息，仅适用于双向止盈止损的限价止盈单
	CTime                string          `json:"cTime"`                //String	订单创建时间，Unix时间戳的毫秒数格式，如 1597026383085
	UTime                string          `json:"uTime"`                //String	订单更新时间，Unix时间戳的毫秒数格式，如 1597026383085
}

type WsOrdersAlgoMiddle struct {
	Arg  WsSubscribeArg `json:"arg"`
	Data []OrdersAlgo   `json:"data"`
}

func handleWsOrdersAlgo(data []byte) (*[]WsOrdersAlgo, error) {
	wsOrdersAlgoMiddle := WsOrdersAlgoMiddle{}
	err := json.Unmarshal(data, &wsOrdersAlgoMiddle)
	if err != nil {
		return nil, err
	}

	ordersAlgo := wsOrdersAlgoMiddle.Data
	wsOrdersAlgoList := []WsOrdersAlgo{}
	for _, orderAlgo := range ordersAlgo {
		wsOrderAlgo := WsOrdersAlgo{
			WsSubscribeArg: wsOrdersAlgoMiddle.Arg,
			OrdersAlgo:     orderAlgo,
		}
		wsOrdersAlgoList = append(wsOrdersAlgoList, wsOrderAlgo)
	}

	return &wsOrdersAlgoList, nil
}
