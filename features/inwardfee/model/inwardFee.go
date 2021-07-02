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
	// var ruleId float32 = 0
	errs := []string{}
	ruleId := validation.Required([]validation.RuleField{
		{FidelName: "RefId", Value: req.RefId},
		{FidelName: "TransDate", Value: req.TransDate},
		{FidelName: "TransTime", Value: req.TransTime},
		{FidelName: "FeeChannel", Value: req.Detail.FeeChannel},
		{FidelName: "TransactionType", Value: req.Detail.TransactionType},
		{FidelName: "ChargeType", Value: req.Detail.ChargeType},
		{FidelName: "FromCCY", Value: req.Detail.FromCCY},
		{FidelName: "ToCCY", Value: req.Detail.ToCCY},
		{FidelName: "EffectiveDate", Value: req.Detail.EffectiveDate},
		{FidelName: "AmountFrom", Value: req.Detail.AmountFrom},
		{FidelName: "AmountTo", Value: req.Detail.AmountTo},
		{FidelName: "ExchangeRate", Value: req.Detail.ExchangeRate},
	}, &errs)
	if len(errs) > 0 {
		return ruleId, errs
	}

	// #region Validate Rule 2 => length => V002
	// ruleId = 2
	// errs := utils.ValidateByRule(req, ruleId, []utils.ValidateRules{
	// 	// //Header
	// 	// {Obj: req, RuleFields: []utils.RuleField{
	// 	// 	{FieldName: "RefId", Length: 15},
	// 	// 	{FieldName: "TransDate", Length: 8},
	// 	// 	{FieldName: "TransTime", Length: 8},
	// 	// }},
	// 	//Detail
	// 	{Obj: req.Detail, RuleFields: []utils.RuleField{
	// 		// {FieldName: "AccountNo", Length: 20},
	// 		// {FieldName: "CifNo", Length: 20},
	// 		// {FieldName: "FeeChannel", Length: 20},
	// 		// {FieldName: "TransactionType", Length: 10},
	// 		// {FieldName: "ChargeType", Length: 3},
	// 		// {FieldName: "OrderingType", Length: 10},
	// 		// {FieldName: "SearchPayInFull", Length: 1},
	// 		// {FieldName: "DepositWithdraw", Length: 10},
	// 		// {FieldName: "BenCountry", Length: 2},
	// 		// {FieldName: "Purpose", Length: 10},
	// 		// {FieldName: "FromCCY", Length: 3},
	// 		// {FieldName: "ToCCY", Length: 3},
	// 		// {FieldName: "EffectiveDate", Length: 8},
	// 		// {FieldName: "FeeType", Length: 20},

	// 		// Todo: เรื่อง decimal
	// 		{FieldName: "AmountFrom", Length: 22},   //17,2
	// 		{FieldName: "AmountTo", Length: 22},     //17,2
	// 		{FieldName: "ExchangeRate", Length: 22}, //17.10

	// 	}},
	// })
	// if errs != nil {
	// 	return ruleId, errs
	// }
	// #endregion

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

	return ruleId, errs
}
