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

func validateHeader(req *Request) (float32, []string) {
	var ruleId float32 = 0

	// #region Validate Rule 1 => Mandatory field => V001
	ruleId = 1
	errs := utils.ValidateByRule(req, ruleId, []utils.ValidateRules{
		//config validate header
		{Obj: req, RuleFields: []utils.RuleField{
			{FieldName: "RefId"},
			{FieldName: "TransDate"},
			{FieldName: "TransTime"},
		}},
		//config validate detail
		{Obj: req.Detail, RuleFields: []utils.RuleField{
			{FieldName: "FeeChannel"},
			{FieldName: "TransactionType"},
			{FieldName: "ChargeType"},
			{FieldName: "FromCCY"},
			{FieldName: "ToCCY"},
			{FieldName: "AmountFrom"},   //Numeric
			{FieldName: "AmountTo"},     //Numeric
			{FieldName: "ExchangeRate"}, //Numeric
			{FieldName: "EffectiveDate"},
		}},
	})
	if errs != nil {
		return ruleId, errs
	}
	// #endregion

	// #region Validate Rule 2 => length => V002
	ruleId = 2
	errs = utils.ValidateByRule(req, ruleId, []utils.ValidateRules{
		//config validate header
		{Obj: req, RuleFields: []utils.RuleField{
			{FieldName: "RefId", Length: 15},
			{FieldName: "TransDate", Length: 8},
			{FieldName: "TransTime", Length: 8},
		}},
		//config validate detail
		{Obj: req.Detail, RuleFields: []utils.RuleField{
			{FieldName: "AccountNo", Length: 20},
			{FieldName: "CifNo", Length: 20},
			{FieldName: "FeeChannel", Length: 20},
			{FieldName: "TransactionType", Length: 10},
			{FieldName: "ChargeType", Length: 3},
			{FieldName: "OrderingType", Length: 10},
			{FieldName: "SearchPayInFull", Length: 1},
			{FieldName: "DepositWithdraw", Length: 10},
			{FieldName: "BenCountry", Length: 2},
			{FieldName: "Purpose", Length: 10},
			{FieldName: "FromCCY", Length: 3},
			{FieldName: "ToCCY", Length: 3},
			{FieldName: "EffectiveDate", Length: 8},
			{FieldName: "FeeType", Length: 20},

			// Todo: เรื่อง decimal
			// {FieldName: "AmountFrom", Length: 17}, //17,2
			// {FieldName: "AmountTo", Length: 17}, //17,2
			// {FieldName: "ExchangeRate", Length: 17}, //17.10

		}},
	})
	if errs != nil {
		return ruleId, errs
	}
	// #endregion

	// #region Validate Rule 4 => Fix value => V004
	ruleId = 4
	errs = utils.ValidateByRule(req, ruleId, []utils.ValidateRules{
		//config validate detail
		{Obj: req.Detail, RuleFields: []utils.RuleField{
			{FieldName: "FeeChannel", Value: req.Detail.FeeChannel,
				Condition: []string{"SWIFT"}},
			{FieldName: "TransactionType", Value: req.Detail.TransactionType,
				Condition: []string{"THB", "FCD", "CASH", "Multi", "e-Money"}},
			{FieldName: "ChargeType", Value: req.Detail.ChargeType,
				Condition: []string{"BEN", "SHA", "OUR"}},
			{FieldName: "OrderingType", Value: req.Detail.OrderingType,
				Condition: []string{"corp", "retail"}},
			{FieldName: "SearchPayInFull", Value: req.Detail.SearchPayInFull,
				Condition: []string{"Y", "N"}},
			{FieldName: "DepositWithdraw", Value: req.Detail.DepositWithdraw,
				Condition: []string{"Deposit", "Withdraw"}},
			{FieldName: "FeeType", Value: req.Detail.FeeType,
				Condition: []string{"Inward Fee", "Cable Fee", "Bahtnet Fee", "Investigate Fee"}},
		}},
	})
	if errs != nil {
		return ruleId, errs
	}
	// #endregion

	return ruleId, errs
}

//1 	=> Mandatory field => V001
//2 	=> length => V002
//3.x 	=> Character set => V003
//4, 	=> Fix value => V004
//5.x 	=> Date pattern => V005
