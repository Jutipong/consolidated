package model

//## Request Model
type Request struct {
	ChannelID string `json:"channelID"`
	RefId     string `json:"refId"`
	TransDate string `json:"transDate"`
	TransTime string `json:"transTime"`
	// reqDetail reqDetail
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
