package controllers

import (
	"account_service/functions"
	"account_service/models"
	"account_service/structs/requests"
	"account_service/structs/responses"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Customer_accountsController operations for Customer_accounts
type Customer_accountsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Customer_accountsController) URLMapping() {
	c.Mapping("AddCustomerAccount", c.AddCustomerAccount)
	c.Mapping("DebitAccount", c.DebitAccount)
	c.Mapping("CreditAccount", c.CreditAccount)
	c.Mapping("GetAccountByAccountNumber", c.GetAccountByAccountNumber)
	c.Mapping("GetAccountByCustomerId", c.GetAccountsByCustomerId)
	c.Mapping("Delete", c.Delete)
	c.Mapping("AccountHistory", c.AccountHistory)
	c.Mapping("UpdateBalance", c.UpdateBalance)
}

// AddCustomerAccount ...
// @Title Add CustomerAccount
// @Description create Customer_accounts
// @Param	body		body 	requests.CreateCustomerAccountRequest	true		"body for Customer_accounts content"
// @Success 201 {int} models.Customer_accounts
// @Failure 403 body is empty
// @router /add-account [post]
func (c *Customer_accountsController) AddCustomerAccount() {
	var v requests.CreateCustomerAccountRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if customer_, err := models.GetCustomerById(v.CustomerId); err != nil {
		logs.Error("Error fetching customer: ", err)
		var resp = responses.CustomerAccountResponse{StatusCode: "500", StatusMessage: "Error creating customer account: " + err.Error(), Result: nil}
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = resp
	} else {

		customerJson, _ := json.Marshal(customer_)
		logs.Info("Customer fetched is ", string(customerJson))
		customer := models.Customer_accounts{
			AccountNumber: v.AccountNumber,
			AccountAlias:  v.AccountAlias,
			Customer:      customer_,
			Balance:       0,
			FrozenAmount:  0,
			BalanceBefore: 0,
			AccountType:   v.AccountType,
			Reference:     v.Reference,
			DateCreated:   time.Now(),
			DateModified:  time.Now(),
			CreatedBy:     v.CreatedBy,
			ModifiedBy:    v.CreatedBy,
			Active:        v.Active,
		}
		if _, err := models.AddCustomer_accounts(&customer); err == nil {
			c.Ctx.Output.SetStatus(200)
			customerObj := responses.CustomerAccountResponseObj{
				CustomerAccountId: customer.CustomerAccountId,
				AccountNumber:     customer.AccountNumber,
				AccountAlias:      customer.AccountAlias,
				Balance:           customer.Balance,
				FrozenAmount:      customer.FrozenAmount,
				BalanceBefore:     customer.BalanceBefore,
				DateCreated:       customer.DateCreated.Format("2006-01-02 15:04:05"),
				Active:            customer.Active,
			}

			resp := responses.CustomerAccountResponse{
				StatusCode:    "200",
				StatusMessage: "Customer account created successfully",
				Result:        &customerObj,
			}
			c.Data["json"] = resp
		} else {
			logs.Error("Error adding customer account: ", err)
			var resp = responses.CustomerAccountResponse{StatusCode: "500", StatusMessage: "Error creating customer account: " + err.Error(), Result: nil}
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = resp
		}
	}
	c.ServeJSON()
}

// DebitAccount ...
// @Title DebitAccount
// @Description update the Customer_accounts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.DebitAccountRequest	true		"body for Customer_accounts content"
// @Success 200 {object} models.Customer_accounts
// @Failure 403 :id is not int
// @router /debit-account/:id [put]
func (c *Customer_accountsController) DebitAccount() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.DebitAccountRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	amount := c.Ctx.Input.Query("amount")
	modifiedBy := c.Ctx.Input.Query("modified_by")
	// reason := c.Ctx.Input.Query("reason")

	statusCode := "500"
	statusDesc := "Error debiting account"
	result := responses.CustomerAccountResponseObj{}

	if custAccount, err := models.GetCustomer_accountsById(id); err == nil {
		logs.Info("About to validate amount ", amount, " against balance ", custAccount.Balance)
		amountFloat, _ := strconv.ParseFloat(amount, 64)
		modifiedByInt, _ := strconv.Atoi(modifiedBy)
		if status, message := functions.ValidateAmount(amountFloat, custAccount.Balance); !status {
			logs.Error("Validation error: ", message)
			statusCode = "400"
			statusDesc = message
		} else {

			currentBalance := custAccount.Balance
			custAccount.BalanceBefore = currentBalance
			custAccount.Balance = currentBalance - amountFloat
			custAccount.DateModified = time.Now()
			custAccount.ModifiedBy = modifiedByInt
			logs.Info("Amount to debit is ", amountFloat)
			logs.Info("Customer account before update is ", custAccount)
			logs.Info("Current balance before debit is ", currentBalance)
			logs.Info("New balance after debit will be ", custAccount.Balance)
			if err := models.UpdateCustomer_accountsById(custAccount); err == nil {
				accountHistory := models.Customer_account_history{
					CustomerAccount: custAccount,
					DebitAmount:     amountFloat,
					CreditAmount:    0,
					Reason:          v.Reason,
					DateCreated:     time.Now(),
					DateModified:    time.Now(),
					CreatedBy:       modifiedByInt,
					ModifiedBy:      modifiedByInt,
				}

				if _, err := models.AddCustomer_account_history(&accountHistory); err != nil {
					logs.Error("Error adding to account history: ", err)
					statusCode = "500"
					statusDesc = "Error debiting account: " + err.Error()
				} else {
					logs.Info("Account history added successfully: ", accountHistory)

					statusCode = "200"
					statusDesc = "Customer account debited successfully"

					result = responses.CustomerAccountResponseObj{
						CustomerAccountId: custAccount.CustomerAccountId,
						AccountNumber:     custAccount.AccountNumber,
						AccountAlias:      custAccount.AccountAlias,
						Balance:           custAccount.Balance,
						FrozenAmount:      custAccount.FrozenAmount,
						BalanceBefore:     custAccount.BalanceBefore,
						DateCreated:       custAccount.DateCreated.Format("2006-01-02 15:04:05"),
						Active:            custAccount.Active,
					}
				}

			} else {
				logs.Error("Error updating customer account: ", err)
				statusCode = "500"
				statusDesc = "Error updating customer account: " + err.Error()
			}
		}
	} else {
		logs.Error("Error fetching customer account: ", err)
		statusCode = "500"
		statusDesc = "Error fetching customer account: " + err.Error()
	}
	resp := responses.CustomerAccountResponse{
		StatusCode:    statusCode,
		StatusMessage: statusDesc,
		Result:        &result,
	}

	c.Data["json"] = resp
	c.Ctx.Output.SetStatus(200)

	c.ServeJSON()
}

// CreditAccount ...
// @Title CreditAccount
// @Description update the Customer_accounts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.CreditAccountRequest	true		"body for Customer_accounts content"
// @Success 200 {object} models.Customer_accounts
// @Failure 403 :id is not int
// @router /credit-account/:id [put]
func (c *Customer_accountsController) CreditAccount() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.CreditAccountRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	amount := c.Ctx.Input.Query("amount")
	modifiedBy := c.Ctx.Input.Query("modified_by")
	// reason := c.Ctx.Input.Query("reason")

	statusCode := "500"
	statusDesc := "Error crediting account"
	result := responses.CustomerAccountResponseObj{}

	if custAccount, err := models.GetCustomer_accountsById(id); err == nil {
		logs.Info("About to validate amount ", amount, " against balance ", custAccount.Balance)
		amountFloat, _ := strconv.ParseFloat(amount, 64)
		modifiedByInt, _ := strconv.Atoi(modifiedBy)
		if status, message := functions.ValidateAmount(amountFloat, custAccount.Balance); !status {
			logs.Error("Validation error: ", message)
			statusCode = "400"
			statusDesc = message
		} else {

			currentBalance := custAccount.Balance
			custAccount.BalanceBefore = currentBalance
			custAccount.Balance = currentBalance + amountFloat
			custAccount.DateModified = time.Now()
			custAccount.ModifiedBy = modifiedByInt

			logs.Info("Amount to credit is ", amountFloat)
			logs.Info("Customer account before update is ", custAccount)
			logs.Info("Current balance before credit is ", currentBalance)
			logs.Info("New balance after credit will be ", custAccount.Balance)
			if err := models.UpdateCustomer_accountsById(custAccount); err == nil {
				accountHistory := models.Customer_account_history{
					CustomerAccount: custAccount,
					DebitAmount:     amountFloat,
					CreditAmount:    0,
					Reason:          v.Reason,
					DateCreated:     time.Now(),
					DateModified:    time.Now(),
					CreatedBy:       modifiedByInt,
					ModifiedBy:      modifiedByInt,
				}

				if _, err := models.AddCustomer_account_history(&accountHistory); err != nil {
					logs.Error("Error adding to account history: ", err)
					statusCode = "500"
					statusDesc = "Error crediting account: " + err.Error()
				} else {
					logs.Info("Account history added successfully: ", accountHistory)
					statusCode = "200"
					statusDesc = "Customer account credited successfully"

					result = responses.CustomerAccountResponseObj{
						CustomerAccountId: custAccount.CustomerAccountId,
						AccountNumber:     custAccount.AccountNumber,
						AccountAlias:      custAccount.AccountAlias,
						Balance:           custAccount.Balance,
						FrozenAmount:      custAccount.FrozenAmount,
						BalanceBefore:     custAccount.BalanceBefore,
						DateCreated:       custAccount.DateCreated.Format("2006-01-02 15:04:05"),
						Active:            custAccount.Active,
					}
				}

			} else {
				logs.Error("Error updating customer account: ", err)
				statusCode = "500"
				statusDesc = "Error updating customer account: " + err.Error()
			}
		}
	} else {
		logs.Error("Error fetching customer account: ", err)
		statusCode = "500"
		statusDesc = "Error fetching customer account: " + err.Error()
	}
	resp := responses.CustomerAccountResponse{
		StatusCode:    statusCode,
		StatusMessage: statusDesc,
		Result:        &result,
	}

	c.Data["json"] = resp
	c.Ctx.Output.SetStatus(200)

	c.ServeJSON()
}

// UpdateBalance ...
// @Title UpdateBalance
// @Description update the Customer_accounts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateBalanceRequest	true		"body for Customer_accounts content"
// @Success 200 {object} models.Customer_accounts
// @Failure 403 :id is not int
// @router /update-balance/:id [put]
func (c *Customer_accountsController) UpdateBalance() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.UpdateBalanceRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logs.Error("Failed to pretty-print request body: ", err)
	} else {
		logs.Info("Request body:\n", string(pretty))
	}

	statusCode := "500"
	statusDesc := "Error debiting account"
	result := responses.CustomerAccountResponseObj{}

	if custAccount, err := models.GetCustomer_accountsById(id); err == nil {
		custAccount.BalanceBefore = custAccount.Balance
		custAccount.Balance = v.Balance
		custAccount.DateModified = time.Now()
		custAccount.ModifiedBy = v.ModifiedBy

		if err := models.UpdateCustomer_accountsById(custAccount); err == nil {
			result = responses.CustomerAccountResponseObj{
				CustomerAccountId: custAccount.CustomerAccountId,
				AccountNumber:     custAccount.AccountNumber,
				AccountAlias:      custAccount.AccountAlias,
				Balance:           custAccount.Balance,
				FrozenAmount:      custAccount.FrozenAmount,
				BalanceBefore:     custAccount.BalanceBefore,
				DateCreated:       custAccount.DateCreated.Format("2006-01-02 15:04:05"),
				Active:            custAccount.Active,
			}

			statusCode = "200"
			statusDesc = "Customer account balance updated successfully"

		} else {
			logs.Error("Error updating customer account: ", err)
			statusCode = "500"
			statusDesc = "Error updating customer account: " + err.Error()
		}
	} else {
		logs.Error("Error fetching customer account: ", err)
		statusCode = "500"
	}
	resp := responses.CustomerAccountResponse{
		StatusCode:    statusCode,
		StatusMessage: statusDesc,
		Result:        &result,
	}

	c.Data["json"] = resp
	c.Ctx.Output.SetStatus(200)

	c.ServeJSON()
}

// GetAccountByAccountNumber ...
// @Title Get Account with Account Number
// @Description get Customer_accounts by account number
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customer_accounts
// @Failure 403 :id is empty
// @router /account/:accountNumber [get]
func (c *Customer_accountsController) GetAccountByAccountNumber() {
	accountNumberStr := c.Ctx.Input.Param(":accountNumber")
	v, err := models.GetCustomer_accountsByAccountNumber(accountNumberStr)

	statusCode := "500"
	statusDesc := "Error crediting account"
	result := responses.CustomerAccountResponseObj{}

	if err != nil {
		logs.Error("Error fetching customer account: ", err)
		statusCode = "500"
		statusDesc = "Error fetching customer account: " + err.Error()

	} else {
		result = responses.CustomerAccountResponseObj{
			CustomerAccountId: v.CustomerAccountId,
			AccountNumber:     v.AccountNumber,
			AccountAlias:      v.AccountAlias,
			Balance:           v.Balance,
			FrozenAmount:      v.FrozenAmount,
			BalanceBefore:     v.BalanceBefore,
			DateCreated:       v.DateCreated.Format("2006-01-02 15:04:05"),
			Active:            v.Active,
		}
		statusCode = "200"
		statusDesc = "Customer account fetched successfully"
	}

	resp := responses.CustomerAccountResponse{
		StatusCode:    statusCode,
		StatusMessage: statusDesc,
		Result:        &result,
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// GetAccountsByCustomerId ...
// @Title Get Account with Customer Id
// @Description get Customer_accounts by account number
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customer_accounts
// @Failure 403 :id is empty
// @router /customer/:id [get]
func (c *Customer_accountsController) GetAccountsByCustomerId() {
	customerId := c.Ctx.Input.Param(":id")

	logs.Info("Fetching accounts for customer id: ", customerId)

	statusCode := "500"
	statusDesc := "Error crediting account"
	result := []*responses.CustomerAccountResponseObj{}

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	customerSearch := "Customer__CustomerId:" + customerId

	if v := customerSearch; v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	logs.Info("Query map to get customer accounts is ", query)

	if custAccounts, err := models.GetAllCustomer_accounts(query, fields, sortby, order, offset, limit); err != nil {
		logs.Error("Error fetching customer account: ", err)
		statusCode = "500"
		statusDesc = "Error fetching customer accounts: " + err.Error()
	} else {
		logs.Info("Customer accounts fetched are ", custAccounts)
		if len(custAccounts) > 0 {
			logs.Info("Number of customer accounts found is ", len(custAccounts))
			for _, v := range custAccounts {
				custAccount := v.(models.Customer_accounts)
				account := responses.CustomerAccountResponseObj{
					CustomerAccountId: custAccount.CustomerAccountId,
					AccountNumber:     custAccount.AccountNumber,
					AccountAlias:      custAccount.AccountAlias,
					AccountType:       custAccount.AccountType,
					Reference:         custAccount.Reference,
					Balance:           custAccount.Balance,
					FrozenAmount:      custAccount.FrozenAmount,
					BalanceBefore:     custAccount.BalanceBefore,
					DateCreated:       custAccount.DateCreated.Format("2006-01-02 15:04:05"),
					Active:            custAccount.Active,
				}
				result = append(result, &account)
			}
			statusCode = "200"
			statusDesc = "Customer account fetched successfully"
		} else {
			statusCode = "200"
			statusDesc = "No customer accounts found"
		}
	}
	resp := responses.CustomerAccountsResponse{
		StatusCode:    statusCode,
		StatusMessage: statusDesc,
		Result:        result,
	}
	c.Data["json"] = resp
	c.ServeJSON()

}

// AccountHistory ...
// @Title Account History
// @Description get Customer_account_anomalies
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customer_account_anomalies
// @Failure 403 :accountNumber is empty
// @router /account-history/:accountNumber [get]
func (c *Customer_accountsController) AccountHistory() {
	accountNumber := c.Ctx.Input.Param(":accountNumber")

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	query["CustomerAccount__AccountNumber"] = accountNumber

	logs.Info("Query map to get customer account history is ", query)

	response := responses.CustomerAccountHistoryResponse{}

	l, err := models.GetAllCustomer_account_history(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("Error fetching account anomalies: ", err)
		response = responses.CustomerAccountHistoryResponse{
			StatusCode:    "502",
			StatusMessage: "Error fetching account anomalies" + err.Error(),
			Result:        nil,
		}
		c.Data["json"] = response
	} else {
		logs.Info("Account anomalies fetched successfully")
		accountHistory := []*responses.CustomerAccountHistoryData{}
		for _, v := range l {
			accountHistModel := v.(models.Customer_account_history)
			accountHistData := responses.CustomerAccountHistoryData{
				CustomerAccountHistoryId: accountHistModel.CustomerAccountHistoryId,
				CustomerAccount:          accountHistModel.CustomerAccount.AccountNumber,
				DebitAmount:              accountHistModel.DebitAmount,
				CreditAmount:             accountHistModel.CreditAmount,
				TransactionDate:          accountHistModel.DateCreated.Format("2006-01-02 15:04:05"),
				Reference:                accountHistModel.Reason,
			}
			accountHistory = append(accountHistory, &accountHistData)
			//
		}

		response = responses.CustomerAccountHistoryResponse{
			StatusCode:    "200",
			StatusMessage: "Account anomalies fetched successfully",
			Result:        accountHistory,
		}
		c.Data["json"] = response
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customer_accounts
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Customer_accountsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCustomer_accounts(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
