package model

import (
	"consolidated/validation"
)

//## Request Model
type Request struct {
	ChannelID string    `json:"channelID"`
	RefId     string    `json:"refId"`
	TransDate string    `json:"transDate"`
	TransTime string    `json:"transTime"`
	Detail    ReqDetail `json:"reqDetail"`
}

type ReqDetail struct {
	AccountNo       string `json:"accountNo"`
	CifNo           string `json:"cifNo"`
	FeeChannel      string `json:"feeChannel"`
	TransactionType string `json:"transactionType"`
	ChargeType      string `json:"chargeType"`
	OrderingType    string `json:"orderingType"`
	SearchPayInFull string `json:"searchPayInFull"`
	DepositWithdraw string `json:"depositWithdraw"`
	BenCountry      string `json:"benCountry"`
	Purpose         string `json:"purpose"`
	FromCCY         string `json:"fromCCY"`
	ToCCY           string `json:"toCCY"`

	AmountFrom   float64 `json:"amountFrom"`
	AmountTo     float64 `json:"amountTo"`
	ExchangeRate float64 `json:"exchangeRate"`

	EffectiveDate string `json:"effectiveDate"`
	FeeType       string `json:"feeType"`
}

type Response struct {
	FeeType     string  `json:"feeType"`
	FeeAmount   float64 `json:"feeAmount"` //Numeric 17,2
	FeeCCY      string  `json:"feeCCY"`
	FeeCategory string  `json:"feeCategory"`
	FeeRefID    string  `json:"feeRefID"`
}

func (h *Request) Validate() (float32, []string) {
	ruleId, errs := validateHeader(h)
	if len(errs) != 0 {
		return ruleId, errs
	} else {
		return ruleId, errs
	}
}

func validateHeader(req *Request) (float32, []string) {
	errs := []string{}

	//Rule 1 => Required
	ruleId := validation.Required([]validation.RequiredRule{
		// {FieldName: "RefId", Value: req.RefId},
		// {FieldName: "TransDate", Value: req.TransDate},
		// {FieldName: "TransTime", Value: req.TransTime},
		//Detail
		// {FieldName: "FeeChannel", Value: req.Detail.FeeChannel},
		// {FieldName: "TransactionType", Value: req.Detail.TransactionType},
		// {FieldName: "ChargeType", Value: req.Detail.ChargeType},
		// {FieldName: "FromCCY", Value: req.Detail.FromCCY},
		// {FieldName: "ToCCY", Value: req.Detail.ToCCY},
		// {FieldName: "EffectiveDate", Value: req.Detail.EffectiveDate},
		// {FieldName: "AmountFrom", Value: req.Detail.AmountFrom},
		// {FieldName: "AmountTo", Value: req.Detail.AmountTo},
		// {FieldName: "ExchangeRate", Value: req.Detail.ExchangeRate},
	}, &errs)
	if len(errs) > 0 {
		return ruleId, errs
	}

	//Rule 2 => Required
	ruleId = validation.MaxLength([]validation.MaxLengthRule{
		// {FieldName: "RefId", Value: req.RefId, Length: 15},
		// {FieldName: "TransDate", Value: req.TransDate, Length: 8},
		// {FieldName: "TransTime", Value: req.TransTime, Length: 8},
		// {FieldName: "AccountNo", Value: req.Detail.AccountNo, Length: 20},
		// {FieldName: "CifNo", Value: req.Detail.CifNo, Length: 20},
		// {FieldName: "FeeChannel", Value: req.Detail.FeeChannel, Length: 20},
		// {FieldName: "TransactionType", Value: req.Detail.TransactionType, Length: 10},
		// {FieldName: "ChargeType", Value: req.Detail.ChargeType, Length: 3},
		// {FieldName: "OrderingType", Value: req.Detail.OrderingType, Length: 10},
		// {FieldName: "SearchPayInFull", Value: req.Detail.SearchPayInFull, Length: 1},
		// {FieldName: "DepositWithdraw", Value: req.Detail.DepositWithdraw, Length: 10},
		// {FieldName: "BenCountry", Value: req.Detail.BenCountry, Length: 2},
		// {FieldName: "Purpose", Value: req.Detail.Purpose, Length: 10},
		// {FieldName: "FromCCY", Value: req.Detail.FromCCY, Length: 3},
		// {FieldName: "ToCCY", Value: req.Detail.ToCCY, Length: 3},
		// {FieldName: "EffectiveDate", Value: req.Detail.EffectiveDate, Length: 8},
		// {FieldName: "FeeType", Value: req.Detail.FeeType, Length: 20},
		// // Todo: เรื่อง decimal
		// {FieldName: "AmountFrom", Value: req.Detail.AmountFrom, Length: 18},     //17,2
		// {FieldName: "AmountTo", Value: req.Detail.AmountTo, Length: 18},         //17,2
		// {FieldName: "ExchangeRate", Value: req.Detail.ExchangeRate, Length: 18}, //17.10
	}, &errs)
	if len(errs) > 0 {
		return ruleId, errs
	}

	//Rule 2. => Digit length => .2
	ruleId = validation.DigitLength([]validation.DigitLengthRule{
		{FieldName: "AmountFrom", Value: req.Detail.AmountFrom},     //17,2
		{FieldName: "AmountTo", Value: req.Detail.AmountTo},         //17,2
		{FieldName: "ExchangeRate", Value: req.Detail.ExchangeRate}, //17.10
	}, &errs)
	if len(errs) > 0 {
		return ruleId, errs
	}

	// #region Validate Rule 2.1 => 2 digit => V002
	// ruleId = 2.1
	// errs := utils.ValidateByRule(req, ruleId, []utils.ValidateRules{
	// 	//Detail
	// 	{Obj: req.Detail, RuleFields: []utils.RuleField{
	// 		// Todo: เรื่อง decimal
	// 		{FieldName: "AmountFrom", ValueFloat: req.Detail.AmountFrom},     //17,2
	// 		{FieldName: "AmountTo", ValueFloat: req.Detail.AmountTo},         //17,2
	// 		{FieldName: "ExchangeRate", ValueFloat: req.Detail.ExchangeRate}, //17.10

	// 	}},
	// })
	// if errs != nil {
	// 	return ruleId, errs
	// }
	// #endregio

	// // #region Validate Rule 4 => Fix value => V004
	// ruleId = 4
	// errs = utils.ValidateByRule(req, ruleId, []utils.ValidateRules{
	// 	//Detail
	// 	{Obj: req.Detail, RuleFields: []utils.RuleField{
	// 		{FieldName: "FeeChannel", Value: req.Detail.FeeChannel,
	// 			Condition: []string{"SWIFT"}},
	// 		{FieldName: "TransactionType", Value: req.Detail.TransactionType,
	// 			Condition: []string{"THB", "FCD", "CASH", "Multi", "e-Money"}},
	// 		{FieldName: "ChargeType", Value: req.Detail.ChargeType,
	// 			Condition: []string{"BEN", "SHA", "OUR"}},
	// 		{FieldName: "OrderingType", Value: req.Detail.OrderingType,
	// 			Condition: []string{"corp", "retail"}},
	// 		{FieldName: "SearchPayInFull", Value: req.Detail.SearchPayInFull,
	// 			Condition: []string{"Y", "N"}},
	// 		{FieldName: "DepositWithdraw", Value: req.Detail.DepositWithdraw,
	// 			Condition: []string{"Deposit", "Withdraw"}},
	// 		{FieldName: "FeeType", Value: req.Detail.FeeType,
	// 			Condition: []string{"Inward Fee", "Cable Fee", "Bahtnet Fee", "Investigate Fee"}},
	// 		{FieldName: "FromCCY", Value: req.Detail.FromCCY,
	// 			Condition: base.MsCurrency()},
	// 		{FieldName: "ToCCY", Value: req.Detail.ToCCY,
	// 			Condition: base.MsCurrency()},
	// 		{FieldName: "BenCountry", Value: req.Detail.BenCountry,
	// 			Condition: base.MsCountry()},
	// 	}},
	// })
	// if errs != nil {
	// 	return ruleId, errs
	// }
	// // #endregion

	// // #region Validate Rule 5.1 => Date pattern (format: YYYYMMDD) => V005
	// ruleId = 5.1
	// errs = utils.ValidateByRule(req, ruleId, []utils.ValidateRules{
	//	//Header
	// 	{Obj: req, RuleFields: []utils.RuleField{
	// 		{FieldName: "TransDate"},
	// 	}},
	// 	//Detail
	// 	{Obj: req.Detail, RuleFields: []utils.RuleField{
	// 		{FieldName: "EffectiveDate"},
	// 	}},
	// })
	// if errs != nil {
	// 	return ruleId, errs
	// }
	// // #endregion

	return 0000, errs
}
