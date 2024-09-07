package myokxapi

type PublicRestPublicInstrumentsReq struct {
	InstType   *string `json:"instType"`   //String 是   产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
	Uly        *string `json:"uly"`        //String 可选 标的指数，如 BTC-USD，仅适用于交割/永续/期权，期权必填
	InstFamily *string `json:"instFamily"` //String 否   交易品种，如 BTC-USD，仅适用于交割/永续/期权
	InstId     *string `json:"instId"`     //String 否   产品ID，如 BTC-USD-200213
}

type PublicRestPublicInstrumentsAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicInstrumentsReq
}

// String 是 产品类型 SPOT：币币 MARGIN：币币杠杆 SWAP：永续合约 FUTURES：交割合约 OPTION：期权
func (api *PublicRestPublicInstrumentsAPI) InstType(instType string) *PublicRestPublicInstrumentsAPI {
	api.req.InstType = GetPointer(instType)
	return api
}

// String 可选 标的指数，如 BTC-USD，仅适用于交割/永续/期权，期权必填
func (api *PublicRestPublicInstrumentsAPI) Uly(uly string) *PublicRestPublicInstrumentsAPI {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 否 交易品种，如 BTC-USD，仅适用于交割/永续/期权
func (api *PublicRestPublicInstrumentsAPI) InstFamily(instFamily string) *PublicRestPublicInstrumentsAPI {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 产品ID，如 BTC-USD-200213
func (api *PublicRestPublicInstrumentsAPI) InstId(instId string) *PublicRestPublicInstrumentsAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

type PublicRestPublicTimeReq struct {
}

type PublicRestPublicTimeAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicTimeReq
}

type PublicRestPublicFundingRateReq struct {
	InstId *string `json:"instId"` //String	否	产品ID ，如 BTC-USD-SWAP 仅适用于永续
}

type PublicRestPublicFundingRateAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicFundingRateReq
}

// String	否	产品ID ，如 BTC-USD-SWAP 仅适用于永续
func (api *PublicRestPublicFundingRateAPI) InstId(instId string) *PublicRestPublicFundingRateAPI {
	api.req.InstId = GetPointer(instId)
	return api
}

type PublicRestPublicOptSummaryReq struct {
	Uly        *string `json:"uly"`        //String 可选 标的指数，仅适用于期权 uly与instFamily必须传一个,若传两个，以instFamily为主
	InstFamily *string `json:"instFamily"` //String 可选 交易品种，仅适用于期权 uly与instFamily必须传一个,若传两个，以instFamily为主
	ExpTime    *string `json:"expTime"`    //String 否 合约到期日，格式为"YYMMDD"，如 "200527"
}

type PublicRestPublicOptSummaryAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicOptSummaryReq
}

// String 可选 标的指数，仅适用于期权 uly与instFamily必须传一个,若传两个，以instFamily为主
func (api *PublicRestPublicOptSummaryAPI) Uly(uly string) *PublicRestPublicOptSummaryAPI {
	api.req.Uly = GetPointer(uly)
	return api
}

// String 可选 交易品种，仅适用于期权 uly与instFamily必须传一个,若传两个，以instFamily为主
func (api *PublicRestPublicOptSummaryAPI) InstFamily(instFamily string) *PublicRestPublicOptSummaryAPI {
	api.req.InstFamily = GetPointer(instFamily)
	return api
}

// String 否 合约到期日，格式为"YYMMDD"，如 "200527"
func (api *PublicRestPublicOptSummaryAPI) ExpTime(expTime string) *PublicRestPublicOptSummaryAPI {
	api.req.ExpTime = GetPointer(expTime)
	return api
}
