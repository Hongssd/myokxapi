package myokxapi

type PublicRestInstrumentsReq struct {
	InstType   *string `json:"instType"`   //String 是   产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String 可选 标的指数，如 BTC-USD，仅适用于交割/永续/期权，期权必填
	InstFamily *string `json:"instFamily"` //String 否   交易品种，如 BTC-USD，仅适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String 否   产品ID，如 BTC-USD-200213
}

type PublicRestInstrumentsApi struct {
	PublicRestClient
	req *PublicRestInstrumentsReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PublicRestInstrumentsApi) InstType(instType string) *PublicRestInstrumentsApi {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 可选 标的指数，如 BTC-USD，仅适用于交割/永续/期权，期权必填
func (api *PublicRestInstrumentsApi) Uly(uly string) *PublicRestInstrumentsApi {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestInstrumentsApi) InstFamily(instFamily string) *PublicRestInstrumentsApi {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如 BTC-USD-200213
func (api *PublicRestInstrumentsApi) InstId(instId string) *PublicRestInstrumentsApi {
	api.req.InstId = GetPointer(instId)
	return api
}

type PublicRestTimeReq struct {
}

type PublicRestTimeApi struct {
	PublicRestClient
	req *PublicRestTimeReq
}

type PublicRestMarkPriceReq struct {
	InstType   *string `json:"instType"`   //String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String	否 标的指数，如 BTC-USD，仅适用于交割/永续/期权
	InstFamily *string `json:"instFamily"` //String	否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String	否 产品ID，如 BTC-USD-200213
}

type PublicRestMarkPriceApi struct {
	PublicRestClient
	req *PublicRestMarkPriceReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PublicRestMarkPriceApi) InstType(instType string) *PublicRestMarkPriceApi {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 否 标的指数，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestMarkPriceApi) Uly(uly string) *PublicRestMarkPriceApi {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestMarkPriceApi) InstFamily(instFamily string) *PublicRestMarkPriceApi {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如 BTC-USD-200213
func (api *PublicRestMarkPriceApi) InstId(instId string) *PublicRestMarkPriceApi {
	api.req.InstId = GetPointer(instId)
	return api
}

type PublicRestMarketTickersReq struct {
	QuoteCcy *string `json:"quoteCcy"` //String	可选 指数计价单位， 目前只有 USD/USDT/BTC/USDC为计价单位的指数，quoteCcy和instId必须填写一个
	InstId   *string `json:"instId"`   //String	可选 指数，如 BTC-USD
}

type PublicRestMarketTickersApi struct {
	PublicRestClient
	req *PublicRestMarketTickersReq
}

// String 可选 指数计价单位， 目前只有 USD/USDT/BTC/USDC为计价单位的指数，quoteCcy和instId必须填写一个
func (api *PublicRestMarketTickersApi) QuoteCcy(quoteCcy string) *PublicRestMarketTickersApi {
	api.req.QuoteCcy = GetPointer(quoteCcy)
	return api
}

// String 可选 指数，如 BTC-USD
func (api *PublicRestMarketTickersApi) InstId(instId string) *PublicRestMarketTickersApi {
	api.req.InstId = GetPointer(instId)
	return api
}
