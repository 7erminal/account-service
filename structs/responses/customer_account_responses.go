package responses

type CustomerAccountResponseObj struct {
	CustomerAccountId int64   `json:"customer_account_id"`
	AccountNumber     string  `json:"account_number"`
	AccountAlias      string  `json:"account_alias"`
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
