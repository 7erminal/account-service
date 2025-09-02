package controllers

import (
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
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("DebitAccount", c.DebitAccount)
	c.Mapping("CreditAccount", c.CreditAccount)
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

	customer := models.Customer_accounts{
		AccountNumber: v.AccountNumber,
		AccountAlias:  v.AccountAlias,
		Balance:       v.Balance,
		FrozenAmount:  0,
		BalanceBefore: v.Balance,
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

	statusCode := "500"
	statusDesc := "Error debiting account"
	result := responses.CustomerAccountResponseObj{}

	if custAccount, err := models.GetCustomer_accountsById(id); err == nil {
		currentBalance := custAccount.Balance
		custAccount.BalanceBefore = currentBalance
		custAccount.Balance = currentBalance - v.Amount
		custAccount.DateModified = time.Now()
		custAccount.ModifiedBy = v.ModifiedBy
		if err := models.UpdateCustomer_accountsById(custAccount); err == nil {
			accountHistory := models.Customer_account_history{
				CustomerAccount: custAccount,
				DebitAmount:     v.Amount,
				CreditAmount:    0,
				DateCreated:     time.Now(),
				DateModified:    time.Now(),
				CreatedBy:       v.ModifiedBy,
				ModifiedBy:      v.ModifiedBy,
			}

			if _, err := models.AddCustomer_account_history(&accountHistory); err != nil {
				logs.Error("Error adding to account history: ", err)
				statusCode = "500"
				statusDesc = "Error debiting account: " + err.Error()
			} else {
				logs.Info("Account history added successfully: ", accountHistory)

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
	v := requests.DebitAccountRequest{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	statusCode := "500"
	statusDesc := "Error crediting account"
	result := responses.CustomerAccountResponseObj{}

	if custAccount, err := models.GetCustomer_accountsById(id); err == nil {
		currentBalance := custAccount.Balance
		custAccount.BalanceBefore = currentBalance
		custAccount.Balance = currentBalance + v.Amount
		custAccount.DateModified = time.Now()
		custAccount.ModifiedBy = v.ModifiedBy
		if err := models.UpdateCustomer_accountsById(custAccount); err == nil {
			accountHistory := models.Customer_account_history{
				CustomerAccount: custAccount,
				DebitAmount:     v.Amount,
				CreditAmount:    0,
				DateCreated:     time.Now(),
				DateModified:    time.Now(),
				CreatedBy:       v.ModifiedBy,
				ModifiedBy:      v.ModifiedBy,
			}

			if _, err := models.AddCustomer_account_history(&accountHistory); err != nil {
				logs.Error("Error adding to account history: ", err)
				statusCode = "500"
				statusDesc = "Error crediting account: " + err.Error()
			} else {
				logs.Info("Account history added successfully: ", accountHistory)

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

// GetOne ...
// @Title Get One
// @Description get Customer_accounts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customer_accounts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Customer_accountsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCustomer_accountsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Customer_accounts
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customer_accounts
// @Failure 403
// @router / [get]
func (c *Customer_accountsController) GetAll() {
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

	l, err := models.GetAllCustomer_accounts(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Customer_accounts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Customer_accounts	true		"body for Customer_accounts content"
// @Success 200 {object} models.Customer_accounts
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Customer_accountsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Customer_accounts{CustomerAccountId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateCustomer_accountsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
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
