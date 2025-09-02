package requests

type CreateCustomerAccountRequest struct {
	AccountNumber string  `json:"account_number" validate:"required"`
	AccountAlias  string  `json:"account_alias" validate:"required"`
	Balance       float64 `json:"balance" validate:"required,gte=0"`
	CreatedBy     int     `json:"created_by" validate:"required"`
	Active        int     `json:"active" validate:"required,oneof=0 1"`
}

type UpdateCustomerAccountRequest struct {
	AccountAlias string  `json:"account_alias" validate:"omitempty"`
	Balance      float64 `json:"balance" validate:"omitempty,gte=0"`
	ModifiedBy   int     `json:"modified_by" validate:"required"`
	Active       int     `json:"active" validate:"omitempty,oneof=0 1"`
}

type DebitAccountRequest struct {
	Amount     float64 `json:"amount" validate:"required,gt=0"`
	ModifiedBy int     `json:"modified_by" validate:"required"`
	Reason     string  `json:"reason" validate:"omitempty"`
}

type CreditAccountRequest struct {
	Amount     float64 `json:"amount" validate:"required,gt=0"`
	ModifiedBy int     `json:"modified_by" validate:"required"`
	Reason     string  `json:"reason" validate:"omitempty"`
}
