package functions

func ValidateAmount(amount float64, accountBalance float64) (status bool, message string) {
	if amount <= 0 {
		return false, "Amount must be greater than zero"
	}

	if amount > accountBalance {
		return false, "Insufficient account balance"
	}

	return true, "Amount is valid"
}
