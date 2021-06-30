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
	// var req Request

	// // //## RefId
	// errs := utils.ValidField(h, "RefId", []utils.Rule{
	// 	{Id: 1, Value: 0},
	// 	{Id: 2, Value: 15},
	// })
	// if errs != nil {
	// 	return errs
	// }

	// //## TransDate
	// errs = utils.ValidField(h, "TransDate", []utils.Rule{
	// 	{Id: 1, Value: 0},
	// 	{Id: 2, Value: 8},
	// 	{Id: 5.1},
	// })
	// if errs != nil {
	// 	return errs
	// }

	//## Validate Rule 1
	statusCode, message, errs := utils.ValidateByRule(h, 1, 0, []string{"RefId", "TransDate"})
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
