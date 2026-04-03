package responses

type CustomerAccountResponseObj struct {
	CustomerAccountId int64   `json:"customer_account_id"`
	CustomerName      string  `json:"customer_name"`
	AccountNumber     string  `json:"account_number"`
	AccountAlias      string  `json:"account_alias"`
	AccountType       string  `json:"account_type,omitempty"`
	Reference         string  `json:"reference,omitempty"`
	Balance           float64 `json:"balance"`
	FrozenAmount      float64 `json:"frozen_amount"`
	BalanceBefore     float64 `json:"balance_before"`
	DateCreated       string  `json:"date_created"`
	Active            int     `json:"active"`
}

type CustomerAccountListResponse struct {
	TotalCount int64                     `json:"total_count"`
	Accounts   []CustomerAccountResponse `json:"accounts"`
}

type CustomerAccountResponse struct {
	StatusCode    string                      `json:"statusCode"`
	StatusMessage string                      `json:"statusMessage"`
	Result        *CustomerAccountResponseObj `json:"result,omitempty"`
}

type CustomerAccountsResponse struct {
	StatusCode    string                        `json:"statusCode"`
	StatusMessage string                        `json:"statusMessage"`
	Result        []*CustomerAccountResponseObj `json:"result,omitempty"`
}

type CustomerAccountHistoryData struct {
	CustomerAccountHistoryId int64
	CustomerAccount          string
	DebitAmount              float64
	CreditAmount             float64
	TransactionDate          string
	Reference                string
	CreatedBy                int
	ModifiedBy               int
}

type CustomerAccountHistoryResponse struct {
	StatusCode    string                        `json:"statusCode"`
	StatusMessage string                        `json:"statusMessage"`
	Result        []*CustomerAccountHistoryData `json:"result,omitempty"`
}
