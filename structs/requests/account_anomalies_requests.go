package requests

type AccountAnomaliesRequest struct {
	AccountNumber  string
	Amount         float64
	Desc           string
	Balance        float64
	CheckedBalance float64
	CreatedBy      int
	ModifiedBy     int
	Active         int
}

type CustomerAccountAnomaliesRequest struct {
	RequestId      string
	AccountNumber  string
	Amount         float64
	Desc           string
	Balance        float64
	CheckedBalance float64
	CreatedBy      int
	ModifiedBy     int
	Active         int
}
