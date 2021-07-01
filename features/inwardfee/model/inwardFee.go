package model

import (
	"consolidated/utils"
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

	//
	AmountFrom   string `json:"amountFrom"`
	AmountTo     string `json:"amountTo"`
	ExchangeRate string `json:"exchangeRate"`
	//

	EffectiveDate string `json:"effectiveDate"`
	FeeType       string `json:"feeType"`
}

func (h *Request) Validate() (float32, []string) {
	ruleId, errs := validateHeader(h)
	if len(errs) != 0 {
		return ruleId, errs
	} else {
		return ruleId, errs
	}
}

func validateHeader(h *Request) (float32, []string) {
	var ruleId float32 = 0

	//## Validate Rule 1
	ruleId = 1
	errs := utils.ValidateByRule(h, ruleId, []utils.ValidateRules{
		//config validate header
		{Obj: h, RuleFields: []utils.RuleField{
			{FieldName: "RefId"},
			{FieldName: "TransDate"},
			{FieldName: "TransTime"},
		}},
		//config validate detail
		{Obj: h.Detail, RuleFields: []utils.RuleField{
			{FieldName: "FeeChannel"},
			{FieldName: "TransactionType"},
			{FieldName: "ChargeType"},
			{FieldName: "FromCCY"},
			{FieldName: "ToCCY"},
			{FieldName: "AmountFrom"},
			{FieldName: "AmountTo"},
			{FieldName: "ExchangeRate"},
			{FieldName: "EffectiveDate"},
		}},
	})
	if errs != nil {
		return ruleId, errs
	}

	return ruleId, errs
}

//1, 2, 3.2, 3.3, 4, 5.1
