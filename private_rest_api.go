package myokxapi

type PrivateRestAPI int

const (
	//Account
	PrivateRestAccountBalance         PrivateRestAPI = iota //查看账户余额
	PrivateRestAccountPosition                              //查看持仓信息
	PrivateRestAccountConfig                                //查看账户配置
	PrivateRestAccountTradeFee                              //查看账户手续费费率
	PrivateRestAccountLeverageInfo                          //获取杠杆倍数
	PrivateRestAccountSetLeverage                           //设置杠杆倍数
	PrivateRestAccountSetPositionMode                       //设置持仓模式
	PrivateRestAccountSetAccountLevel                       //设置账户模式

	//Trade
	PrivateRestTradeOrderGet      //查看订单信息
	PrivateRestTradeOrdersPending //查看未成交订单列表
	PrivateRestTradeOrderPost     //下单

	PrivateRestTradeOrderAlgoPost    //策略委托下单
	PrivateRestTradeCancelOrderAlgo  //撤销策略委托订单
	PrivateRestTradeAmendOrderAlgo   //修改策略委托订单
	PrivateRestTradeOrderAlgoGet     //获取策略委托单信息
	PrivateRestTradeOrderAlgoPending // 获取未完成策略委托单列表
	PrivateRestTradeOrderAlgoHistory //获取历史策略委托单记录

	PrivateRestTradeCancelOrder         //撤单
	PrivateRestTradeAmendOrder          //修改订单
	PrivateRestTradeBatchOrders         //批量下单
	PrivateRestTradeCancelBatchOrders   //批量撤单
	PrivateRestTradeAmendBatchOrders    //批量修改订单
	PrivateRestTradeOrderHistory        //获取历史订单记录（近七天）
	PrivateRestTradeOrderHistoryArchive //获取历史订单记录（近三个月）

	PrivateRestTradeFills        //获取成交明细（近三天）
	PrivateRestTradeFillsHistory //获取成交明细（近三个月）
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	//Account
	PrivateRestAccountBalance:         "/api/v5/account/balance",           //GET 查看账户余额
	PrivateRestAccountPosition:        "/api/v5/account/positions",         //GET 查看持仓信息
	PrivateRestAccountConfig:          "/api/v5/account/config",            //GET 查看账户配置
	PrivateRestAccountTradeFee:        "/api/v5/account/trade-fee",         //GET 查看账户手续费费率
	PrivateRestAccountLeverageInfo:    "/api/v5/account/leverage-info",     //GET 获取杠杆倍数
	PrivateRestAccountSetLeverage:     "/api/v5/account/set-leverage",      //POST 设置杠杆倍数
	PrivateRestAccountSetPositionMode: "/api/v5/account/set-position-mode", //POST 设置持仓模式
	PrivateRestAccountSetAccountLevel: "/api/v5/account/set-account-level", //POST 设置账户模式

	//Trade
	PrivateRestTradeOrderGet:      "/api/v5/trade/order",          //GET 查看订单信息
	PrivateRestTradeOrdersPending: "/api/v5/trade/orders-pending", //GET 查看未成交订单列表
	PrivateRestTradeOrderPost:     "/api/v5/trade/order",          //POST 下单

	PrivateRestTradeOrderAlgoPost:    "/api/v5/trade/order-algo",          //POST  策略委托下单
	PrivateRestTradeCancelOrderAlgo:  "/api/v5/trade/cancel-algos",        // POST 撤销策略委托订单
	PrivateRestTradeAmendOrderAlgo:   "/api/v5/trade/amend-algos",         // POST 修改策略委托订单
	PrivateRestTradeOrderAlgoGet:     "/api/v5/trade/order-algo",          //GET 获取策略委托单信息
	PrivateRestTradeOrderAlgoPending: "/api/v5/trade/orders-algo-pending", //GET  获取未完成策略委托单列表
	PrivateRestTradeOrderAlgoHistory: "/api/v5/trade/orders-algo-history", //GET 获取历史策略委托单记录

	PrivateRestTradeCancelOrder:         "/api/v5/trade/cancel-order",           //POST 撤单
	PrivateRestTradeAmendOrder:          "/api/v5/trade/amend-order",            //POST 修改订单
	PrivateRestTradeBatchOrders:         "/api/v5/trade/batch-orders",           //POST 批量下单
	PrivateRestTradeCancelBatchOrders:   "/api/v5/trade/cancel-batch-orders",    //POST 批量撤单
	PrivateRestTradeAmendBatchOrders:    "/api/v5/trade/amend-batch-orders",     //POST 批量修改订单
	PrivateRestTradeOrderHistory:        "/api/v5/trade/orders-history",         //GET 获取历史订单记录（近七天）
	PrivateRestTradeOrderHistoryArchive: "/api/v5/trade/orders-history-archive", //GET 获取历史订单记录（近三个月）

	PrivateRestTradeFills:        "/api/v5/trade/fills",         //GET 获取成交明细（近三天）
	PrivateRestTradeFillsHistory: "/api/v5/trade/fills-history", //GET 获取成交明细（近三个月）

}
