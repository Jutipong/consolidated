package model

import (
	"consolidated/utils"
)

//## Request Model
type Request struct {
	ChannelID string `json:"channelID"`
	RefId     string `json:"refId"`
	TransDate string `json:"transDate"`
	TransTime string `json:"transTime"`
	// reqDetail reqDetail
}

func (h *Request) Validate() (string, string, interface{}) {
	statusCode, message, res := validateHeader(h)
	return statusCode, message, res
}

func validateHeader(h *Request) (string, string, interface{}) {
	//## Validate Rule 1
	statusCode, message, errs := utils.ValidateByRule(h, 1, 0, []utils.ValidateRule{
		{FieldName: "RefId"},
		{FieldName: "TransDate"},
	})
	if errs != nil {
		return statusCode, message, errs
	}

	// ValidateByRule

	return statusCode, message, nil
}

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
