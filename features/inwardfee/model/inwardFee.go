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
	ReqDetail ReqDetail `json:"reqDetail"`
}

type ReqDetail struct {
	AccountNo       string `json:"reqDetail"`
	CifNo           string `json:"cifNo"`
	FeeChannel      string `json:"feeChannel"`
	TransactionType string `json:"transactionType"`
	ChargeType      string `json:"chargeType"`
	OrderingType    string `json:"orderingType"`
	SearchPayInFull string `json:"searchPayInFull"`
}

func (h *Request) Validate() (float32, []string) {
	ruleId, errs := validateHeader(h)
	if len(errs) != 0 {
		return ruleId, errs
	}
	// ruleId, errs = validateDetail(&h.ReqDetail)
	return ruleId, errs
}

func validateHeader(h *Request) (float32, []string) {
	var ruleId float32 = 0

	//## Validate Rule 1

	rule := []utils.Rules{
		{
			RuleId: 1,
			FieldRules: []utils.FieldRule{
				{Obj: h, Name: []string{"RefId", "TransDate"}},
				{Obj: h.ReqDetail, Name: []string{"AccountNo"}},
			},
		},
	}

	// {Obj: h, FieldName: []string{"RefId", "TransDate"}},
	// {Obj: h.ReqDetail, FieldName: []string{"AccountNo"}},

	ruleId = 1
	errs := utils.ValidateByRule(h, ruleId, rule)
	if errs != nil {
		return ruleId, errs
	}

	// ruleId = 1
	// errs := utils.ValidateByRule(h, ruleId, []utils.ValidateRule{
	// 	{FieldName: "RefId"},
	// 	{FieldName: "TransDate"},
	// })
	// if errs != nil {
	// 	return ruleId, errs
	// }
	return ruleId, errs
}

// func validateDetail(h *ReqDetail) (float32, []string) {
// 	var ruleId float32 = 0

// 	//## Validate Rule 1
// 	ruleId = 1
// 	errs := utils.ValidateByRule(h, ruleId, 0, []utils.ValidateRule{
// 		{FieldName: "RefId"},
// 		{FieldName: "TransDate"},
// 	})
// 	if errs != nil {
// 		return ruleId, errs
// 	}
// 	return ruleId, errs
// }

// type reqHeader struct {
// 	refId     string
// 	transDate string
// 	transTime string
// }

//## Validate
// func ValidateHeader() {

// }

// type reqDetail struct {
// 	accountNo       string
// 	cifNo           string
// 	feeChannel      string
// 	transactionType string
// 	chargeType      string
// 	orderingType    string
// 	searchPayInFull string
// }

//## Response Model
// type Response struct {
// }
