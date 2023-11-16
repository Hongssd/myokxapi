package myokxapi

type PrivateRestAPI int

const (
	//Account
	PrivateRestAccountBalance  PrivateRestAPI = iota //查看账户余额
	PrivateRestAccountPosition                       //查看持仓信息
	PrivateRestAccountConfig                         //查看账户配置

	//Trade
	PrivateRestTradeOrderGet          //查看订单信息
	PrivateRestTradeOrdersPending     //查看未成交订单列表
	PrivateRestTradeOrderPost         //下单
	PrivateRestTradeCancelOrder       //撤单
	PrivateRestTradeAmendOrder        //修改订单
	PrivateRestTradeBatchOrders       //批量下单
	PrivateRestTradeCancelBatchOrders //批量撤单
	PrivateRestTradeAmendBatchOrders  //批量修改订单

)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	//Account
	PrivateRestAccountBalance:  "/api/v5/account/balance",   //GET 查看账户余额
	PrivateRestAccountPosition: "/api/v5/account/positions", //GET 查看持仓信息
	PrivateRestAccountConfig:   "/api/v5/account/config",    //GET 查看账户配置

	//Trade
	PrivateRestTradeOrderGet:          "/api/v5/trade/order",               //GET 查看订单信息
	PrivateRestTradeOrdersPending:     "/api/v5/trade/orders-pending",      //GET 查看未成交订单列表
	PrivateRestTradeOrderPost:         "/api/v5/trade/order",               //POST 下单
	PrivateRestTradeCancelOrder:       "/api/v5/trade/cancel-order",        //POST 撤单
	PrivateRestTradeAmendOrder:        "/api/v5/trade/amend-order",         //POST 修改订单
	PrivateRestTradeBatchOrders:       "/api/v5/trade/batch-orders",        //POST 批量下单
	PrivateRestTradeCancelBatchOrders: "/api/v5/trade/cancel-batch-orders", //POST 批量撤单
	PrivateRestTradeAmendBatchOrders:  "/api/v5/trade/amend-batch-orders",  //POST 批量修改订单
}
