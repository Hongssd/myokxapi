package myokxapi

type RestInstrumentsReq struct {
	InstType   *string `json:"instType"`   //String 是   产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String 可选 标的指数，如 BTC-USD，仅适用于交割/永续/期权，期权必填
	InstFamily *string `json:"instFamily"` //String 否   交易品种，如 BTC-USD，仅适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String 否   产品ID，如 BTC-USD-200213
}

type RestInstrumentsApi struct {
	RestClient
	req *RestInstrumentsReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *RestInstrumentsApi) InstType(instType string) *RestInstrumentsApi {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 可选 标的指数，如 BTC-USD，仅适用于交割/永续/期权，期权必填
func (api *RestInstrumentsApi) Uly(uly string) *RestInstrumentsApi {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
func (api *RestInstrumentsApi) InstFamily(instFamily string) *RestInstrumentsApi {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如 BTC-USD-200213
func (api *RestInstrumentsApi) InstId(instId string) *RestInstrumentsApi {
	api.req.InstId = GetPointer(instId)
	return api
}

type RestTimeReq struct {
}

type RestTimeApi struct {
	RestClient
	req *RestTimeReq
}

type RestMarkPriceReq struct {
	InstType   *string `json:"instType"`   //String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String	否 标的指数，如 BTC-USD，仅适用于交割/永续/期权
	InstFamily *string `json:"instFamily"` //String	否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String	否 产品ID，如 BTC-USD-200213
}

type RestMarkPriceApi struct {
	RestClient
	req *RestMarkPriceReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *RestMarkPriceApi) InstType(instType string) *RestMarkPriceApi {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 标的指数，如 BTC-USD，仅适用于交割/永续/期权
func (api *RestMarkPriceApi) Uly(uly string) *RestMarkPriceApi {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
func (api *RestMarkPriceApi) InstFamily(instFamily string) *RestMarkPriceApi {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如 BTC-USD-200213
func (api *RestMarkPriceApi) InstId(instId string) *RestMarkPriceApi {
	api.req.InstId = GetPointer(instId)
	return api
}

type RestMarketTickersReq struct {
	QuoteCcy *string `json:"quoteCcy"` //String	可选 指数计价单位， 目前只有 USD/USDT/BTC/USDC为计价单位的指数，quoteCcy和instId必须填写一个
	InstId   *string `json:"instId"`   //String	可选 指数，如 BTC-USD
}

type RestMarketTickersApi struct {
	RestClient
	req *RestMarketTickersReq
}

// String 可选 指数计价单位， 目前只有 USD/USDT/BTC/USDC为计价单位的指数，quoteCcy和instId必须填写一个
func (api *RestMarketTickersApi) QuoteCcy(quoteCcy string) *RestMarketTickersApi {
	api.req.QuoteCcy = GetPointer(quoteCcy)
	return api
}

// String 可选 指数，如 BTC-USD
func (api *RestMarketTickersApi) InstId(instId string) *RestMarketTickersApi {
	api.req.InstId = GetPointer(instId)
	return api
}
