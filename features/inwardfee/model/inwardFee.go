package inwardfree

//## Request Model
type Request struct {
	reqHeader reqHeader
	reqDetail reqDetail
}

type reqHeader struct {
	refId     string `validate:"required|maxLen:7"`
	transDate string `validate:"required|maxLen:7"`
	transTime string `validate:"required|maxLen:7"`
}

type reqDetail struct {
	accountNo       string
	cifNo           string
	feeChannel      string
	transactionType string
	chargeType      string
	orderingType    string
	searchPayInFull string
}

//## Response Model
// type Response struct {
// }
