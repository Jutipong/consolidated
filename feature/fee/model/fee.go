package model

import (
	"consolidated/base"
	"consolidated/validation"
)

//## Request Model
type Request struct {
	RefId     string    `json:"refId"`
	TransDate string    `json:"transDate"`
	TransTime string    `json:"transTime"`
	Detail    ReqDetail `json:"reqDetail"`
}

type ReqDetail struct {
	PromotionCode    string  `json:"promotionCode"`
	CifNo            string  `json:"cifNo"`
	Channel          string  `json:"channel"`
	Product          string  `json:"product"`
	ChargeType       string  `json:"chargeType"`
	OrderingType     string  `json:"orderingType"`
	Sof              string  `json:"sof"`
	FromCCY          string  `json:"fromCCY"`
	ToCCY            string  `json:"toCCY"`
	FeeOutwardAmount float64 `json:"feeOutwardAmount"`
	FeeComAmount     float64 `json:"FeeComAmount"`
	ExchangeRate     float64 `json:"exchangeRate"`
	EffectiveDate    string  `json:"effectiveDate"`
}

type Response struct {
	PromotionType string  `json:"promotionType"`
	Amount        float64 `json:"amount"`
	CCY           string  `json:"ccy"`
	ProRefID      string  `json:"proRefID"`
}

func (h *Request) Validate() (float32, []string) {
	var errs []string
	ruleId := validate(h, &errs)
	if len(errs) != 0 {
		return ruleId, errs
	} else {
		return ruleId, errs
	}
}

func validate(req *Request, errs *[]string) float32 {

	//Rule 1 => Required
	ruleId := validation.Required(&[]validation.RequiredRule{
		{FieldName: "RefId", Value: req.RefId},
		{FieldName: "TransDate", Value: req.TransDate},
		{FieldName: "TransTime", Value: req.TransTime},
		{FieldName: "PromotionCode", Value: req.Detail.PromotionCode},
		{FieldName: "Channel", Value: req.Detail.Channel},
		{FieldName: "Product", Value: req.Detail.Product},
		{FieldName: "ChargeType", Value: req.Detail.ChargeType},
		{FieldName: "OrderingType", Value: req.Detail.OrderingType},
		{FieldName: "Sof", Value: req.Detail.Sof},
		{FieldName: "FromCCY", Value: req.Detail.FromCCY},
		{FieldName: "ToCCY", Value: req.Detail.ToCCY},
		{FieldName: "ExchangeRate", Value: req.Detail.ExchangeRate},
		{FieldName: "EffectiveDate", Value: req.Detail.EffectiveDate},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 2 => MaxLength
	ruleId = validation.MaxLength(&[]validation.MaxLengthRule{
		{FieldName: "RefId", Value: req.RefId, Length: 15},
		{FieldName: "TransDate", Value: req.TransDate, Length: 8},
		{FieldName: "TransTime", Value: req.TransTime, Length: 8},
		{FieldName: "PromotionCode", Value: req.Detail.PromotionCode, Length: 20},
		{FieldName: "CifNo", Value: req.Detail.CifNo, Length: 20},
		{FieldName: "Channel", Value: req.Detail.Channel, Length: 20},
		{FieldName: "Product", Value: req.Detail.Product, Length: 20},
		{FieldName: "ChargeType", Value: req.Detail.ChargeType, Length: 3},
		{FieldName: "OrderingType", Value: req.Detail.OrderingType, Length: 10},
		{FieldName: "Sof", Value: req.Detail.Sof, Length: 3},
		{FieldName: "FromCCY", Value: req.Detail.FromCCY, Length: 3},
		{FieldName: "ToCCY", Value: req.Detail.ToCCY, Length: 3},
		{FieldName: "FeeOutwardAmount", Value: req.Detail.FeeOutwardAmount, Length: 18},
		{FieldName: "FeeComAmount", Value: req.Detail.FeeComAmount, Length: 18},
		{FieldName: "ExchangeRate", Value: req.Detail.ExchangeRate, Length: 18},
		{FieldName: "EffectiveDate", Value: req.Detail.EffectiveDate, Length: 8},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 2.1 => Digit length => .2
	ruleId = validation.DigitLength(&[]validation.DigitLengthRule{
		{FieldName: "FeeOutwardAmount", Value: req.Detail.FeeOutwardAmount},
		{FieldName: "FeeComAmount", Value: req.Detail.FeeComAmount},
		{FieldName: "ExchangeRate", Value: req.Detail.ExchangeRate},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 2.2 => MinValue
	ruleId = validation.MinValue(&[]validation.MinValueRule{
		{FieldName: "ExchangeRate", Min: 0, Value: req.Detail.ExchangeRate},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 4 => Fix value
	ruleId = validation.FixValue(&[]validation.FixValueRule{
		{FieldName: "Channel", Value: req.Detail.Channel,
			Conditions: []string{"Branch", "Online-Corporate", "Online-NEXT", "OWRM", "IWRM"}},
		{FieldName: "Product", Value: req.Detail.Product,
			Conditions: []string{"FCD", "Outward", "Inward"}},
		{FieldName: "ChargeType", Value: req.Detail.ChargeType,
			Conditions: []string{"BEN", "SHA", "OUR"}},
		{FieldName: "OrderingType", Value: req.Detail.OrderingType,
			Conditions: []string{"corp", "retail"}},
		{FieldName: "Sof", Value: req.Detail.Sof,
			Conditions: []string{"THB", "FCD", "e-Money", "CASH"}},
		{FieldName: "FromCCY", Value: req.Detail.FromCCY,
			Conditions: base.MsCurrency()},
		{FieldName: "ToCCY", Value: req.Detail.ToCCY,
			Conditions: base.MsCurrency()},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 5.1 => YYYYMMDD
	ruleId = validation.YYYYMMDD(&[]validation.DateTimeRule{
		{FieldName: "TransDate", Value: req.TransDate},
		{FieldName: "EffectiveDate", Value: req.Detail.EffectiveDate},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 5.2 => YYYYMMDD
	ruleId = validation.HHMMSS(&[]validation.DateTimeRule{
		{FieldName: "TransTime", Value: req.TransTime},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	return base.Successfully
}
