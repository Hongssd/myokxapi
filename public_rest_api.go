package myokxapi

type PublicRestAPI int

const (

	//Public
	PublicRestPublicInstruments PublicRestAPI = iota //获取交易产品基础信息
	PublicRestPublicTime                             //获取系统时间
	PublicRestPublicMarkPrice                        //获取标记价格
	PublicRestPublicFundingRate                      //获取永续合约当前资金费率

	//Market
	PublicRestMarketTickers        //获取指数行情
	PublicRestMarketBooks          // 获取产品轻量深度
	PublicRestMarketCandles        // 获取K线数据 最近1440条
	PublicRestMarketHistoryCandles // 获取K线数据 最近几年 1sK线为最近3个月
)

var PublicRestAPIMap = map[PublicRestAPI]string{
	//Public
	PublicRestPublicInstruments: "/api/v5/public/instruments",  //GET 获取交易产品基础信息
	PublicRestPublicTime:        "/api/v5/public/time",         //GET 获取系统时间
	PublicRestPublicMarkPrice:   "/api/v5/public/mark-price",   //GET 获取标记价格
	PublicRestPublicFundingRate: "/api/v5/public/funding-rate", //GET 获取永续合约当前资金费率

	//Market
	PublicRestMarketTickers:        "/api/v5/market/index-tickers",   //GET 获取指数行情
	PublicRestMarketBooks:          "/api/v5/market/books",           //GET 获取产品轻量深度
	PublicRestMarketCandles:        "/api/v5/market/candles",         //GET 获取K线数据
	PublicRestMarketHistoryCandles: "/api/v5/market/history-candles", //GET 获取K线数据

}
