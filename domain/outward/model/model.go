package model

import (
	"consolidated/base"
	"consolidated/validation"
)

type Request struct {
	ChannelID string    `json:"channelID"`
	RefId     string    `json:"refId"`
	TransDate string    `json:"transDate"`
	TransTime string    `json:"transTime"`
	Detail    ReqDetail `json:"reqDetail"`
}

type ReqDetail struct {
	AccountNo       string  `json:"accountNo"`
	CifNo           string  `json:"cifNo"`
	FeeChannel      string  `json:"feeChannel"`
	TransactionType string  `json:"transactionType"`
	ChargeType      string  `json:"chargeType"`
	OrderingType    string  `json:"orderingType"`
	SearchPayInFull string  `json:"searchPayInFull"`
	BenCountry      string  `json:"benCountry"`
	Purpose         string  `json:"purpose"`
	FromCCY         string  `json:"fromCCY"`
	ToCCY           string  `json:"toCCY"`
	AmountFrom      float64 `json:"amountFrom"`
	AmountTo        float64 `json:"amountTo"`
	ExchangeRate    float64 `json:"exchangeRate"`
	EffectiveDate   string  `json:"effectiveDate"`
	FeeType         string  `json:"feeType"`
}

type Response struct {
	FeeType     string  `json:"feeType"`
	FeeAmount   float64 `json:"feeAmount"`
	FeeCCY      string  `json:"feeCCY"`
	FeeCategory string  `json:"feeCategory"`
	FeeRefID    string  `json:"feeRefID"`
} //@name Outward.Response

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
		{FieldName: "FeeChannel", Value: req.Detail.FeeChannel},
		{FieldName: "TransactionType", Value: req.Detail.TransactionType},
		{FieldName: "ChargeType", Value: req.Detail.ChargeType},
		{FieldName: "FromCCY", Value: req.Detail.FromCCY},
		{FieldName: "ToCCY", Value: req.Detail.ToCCY},
		{FieldName: "EffectiveDate", Value: req.Detail.EffectiveDate},
		{FieldName: "AmountFrom", Value: req.Detail.AmountFrom},
		{FieldName: "AmountTo", Value: req.Detail.AmountTo},
		{FieldName: "ExchangeRate", Value: req.Detail.ExchangeRate},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 2 => MaxLength
	ruleId = validation.MaxLength(&[]validation.MaxLengthRule{
		{FieldName: "RefId", Value: req.RefId, Length: 15},
		{FieldName: "TransDate", Value: req.TransDate, Length: 8},
		{FieldName: "TransTime", Value: req.TransTime, Length: 8},
		{FieldName: "AccountNo", Value: req.Detail.AccountNo, Length: 20},
		{FieldName: "CifNo", Value: req.Detail.CifNo, Length: 20},
		{FieldName: "FeeChannel", Value: req.Detail.FeeChannel, Length: 20},
		{FieldName: "TransactionType", Value: req.Detail.TransactionType, Length: 10},
		{FieldName: "ChargeType", Value: req.Detail.ChargeType, Length: 3},
		{FieldName: "OrderingType", Value: req.Detail.OrderingType, Length: 10},
		{FieldName: "SearchPayInFull", Value: req.Detail.SearchPayInFull, Length: 1},
		{FieldName: "BenCountry", Value: req.Detail.BenCountry, Length: 2},
		{FieldName: "Purpose", Value: req.Detail.Purpose, Length: 10},
		{FieldName: "FromCCY", Value: req.Detail.FromCCY, Length: 3},
		{FieldName: "ToCCY", Value: req.Detail.ToCCY, Length: 3},
		{FieldName: "EffectiveDate", Value: req.Detail.EffectiveDate, Length: 8},
		{FieldName: "FeeType", Value: req.Detail.FeeType, Length: 20},
		{FieldName: "AmountFrom", Value: req.Detail.AmountFrom, Length: 18},
		{FieldName: "AmountTo", Value: req.Detail.AmountTo, Length: 18},
		{FieldName: "ExchangeRate", Value: req.Detail.ExchangeRate, Length: 18},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 2.1 => Digit length => .2
	ruleId = validation.DigitLength(&[]validation.DigitLengthRule{
		{FieldName: "AmountFrom", Value: req.Detail.AmountFrom},
		{FieldName: "AmountTo", Value: req.Detail.AmountTo},
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

	//Rule 3 => Number
	ruleId = validation.Number(&[]validation.NumberRule{
		{FieldName: "AccountNo", Value: req.Detail.AccountNo},
	}, errs)
	if len(*errs) > 0 {
		return ruleId
	}

	//Rule 4 => Fix value
	ruleId = validation.FixValue(&[]validation.FixValueRule{
		{FieldName: "FeeChannel", Value: req.Detail.FeeChannel,
			Conditions: []string{"SWIFT"}},
		{FieldName: "TransactionType", Value: req.Detail.TransactionType,
			Conditions: []string{"THB", "FCD", "CASH", "Multi", "e-Money"}},
		{FieldName: "ChargeType", Value: req.Detail.ChargeType,
			Conditions: []string{"BEN", "SHA", "OUR"}},
		{FieldName: "OrderingType", Value: req.Detail.OrderingType,
			Conditions: []string{"corp", "retail"}},
		{FieldName: "SearchPayInFull", Value: req.Detail.SearchPayInFull,
			Conditions: []string{"Y", "N"}},
		{FieldName: "FeeType", Value: req.Detail.FeeType,
			Conditions: []string{"Inward Fee", "Cable Fee", "Bahtnet Fee", "Investigate Fee"}},
		{FieldName: "FromCCY", Value: req.Detail.FromCCY,
			Conditions: base.MsCurrency()},
		{FieldName: "ToCCY", Value: req.Detail.ToCCY,
			Conditions: base.MsCurrency()},
		{FieldName: "BenCountry", Value: req.Detail.BenCountry,
			Conditions: base.MsCountry()},
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

	return 0000
}
