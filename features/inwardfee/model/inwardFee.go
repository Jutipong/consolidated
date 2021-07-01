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
	rule := []utils.ValidateRules{
		//config validate header
		{
			Obj: h,
			RuleFields: []utils.RuleField{
				{FieldName: "RefId"},
				{FieldName: "TransDate"},
				{FieldName: "TransTime"},
			},
		},
		//config validate detail
		{
			Obj: h.Detail,
			RuleFields: []utils.RuleField{
				{FieldName: "AccountNo"},
			},
		},
	}
	errs := utils.ValidateByRule(h, ruleId, rule)
	if errs != nil {
		return ruleId, errs
	}

	// //## Validate Rule 1
	// ruleId = 1
	// rule := []utils.Rules{
	// 	{
	// 		RuleId: ruleId,
	// 		FieldRules: []utils.ValidateRules{
	// 			{Obj: h, Name: []string{"RefId", "TransDate", "transTime"}},
	// 			{Obj: h.ReqDetail, Name: []string{"AccountNo"}},
	// 		},
	// 	},
	// }
	// errs := utils.ValidateByRule(h, ruleId, rule)
	// if errs != nil {
	// 	return ruleId, errs
	// }

	//## Validate Rule 2
	// ruleId = 2
	// rule = []utils.Rules{
	// 	{
	// 		RuleId: ruleId,
	// 		FieldRules: []utils.FieldRule{
	// 			{Obj: h, Name: []string{"RefId", "TransDate", "transTime"}},
	// 			{Obj: h.ReqDetail, Name: []string{"AccountNo"}},
	// 		},
	// 	},
	// }
	// errs = utils.ValidateByRule(h, ruleId, rule)
	// if errs != nil {
	// 	return ruleId, errs
	// }

	return ruleId, errs
}

//1, 2, 3.2, 3.3, 4, 5.1
