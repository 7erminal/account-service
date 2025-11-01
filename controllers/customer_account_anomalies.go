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

// Customer_account_anomaliesController operations for Customer_account_anomalies
type Customer_account_anomaliesController struct {
	beego.Controller
}

// URLMapping ...
func (c *Customer_account_anomaliesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Customer_account_anomalies
// @Param	body		body 	requests.CustomerAccountAnomaliesRequest	true		"body for Customer_account_anomalies content"
// @Success 201 {int} models.Customer_account_anomalies
// @Failure 403 body is empty
// @router / [post]
func (c *Customer_account_anomaliesController) Post() {
	var v requests.CustomerAccountAnomaliesRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	response := responses.CustomerAccountAnomalyResponse{
		StatusCode:    "502",
		StatusMessage: "Account anomaly creation failed",
		Result:        nil,
	}

	if account, err := models.GetCustomer_accountsByAccountNumber(v.AccountNumber); err != nil {
		c.Data["json"] = err.Error()

		response = responses.CustomerAccountAnomalyResponse{
			StatusCode:    "502",
			StatusMessage: "Account not found" + err.Error(),
			Result:        nil,
		}
		c.Data["json"] = response
	} else {
		statement := v.Desc

		accountAnomaly := models.Customer_account_anomalies{
			Account:        account,
			Amount:         v.Amount,
			Desc:           v.Desc,
			Statement:      statement,
			Balance:        v.Balance,
			CheckedBalance: v.CheckedBalance,
			DateCreated:    time.Now(),
			DateModified:   time.Now(),
			CreatedBy:      v.CreatedBy,
			ModifiedBy:     v.ModifiedBy,
			Active:         v.Active,
		}
		if _, err := models.AddCustomer_account_anomalies(&accountAnomaly); err == nil {
			response = responses.CustomerAccountAnomalyResponse{
				StatusCode:    "200",
				StatusMessage: "Account anomaly created successfully",
				Result:        &accountAnomaly,
			}

			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = response
		} else {
			// c.Data["json"] = err.Error()
			logs.Error("Error creating account anomaly: ", err)

			response = responses.CustomerAccountAnomalyResponse{
				StatusCode:    "502",
				StatusMessage: "Account anomaly creation failed" + err.Error(),
				Result:        nil,
			}
			c.Data["json"] = response
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Customer_account_anomalies by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customer_account_anomalies
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Customer_account_anomaliesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	response := responses.CustomerAccountAnomalyResponse{}
	v, err := models.GetCustomer_account_anomaliesById(id)
	if err != nil {
		logs.Error("Error fetching account anomaly: ", err)
		response = responses.CustomerAccountAnomalyResponse{
			StatusCode:    "502",
			StatusMessage: "Error fetching account anomaly" + err.Error(),
			Result:        nil,
		}
		c.Data["json"] = response
	} else {
		logs.Info("Account anomaly fetched successfully")

		response = responses.CustomerAccountAnomalyResponse{
			StatusCode:    "200",
			StatusMessage: "Account anomaly fetched successfully",
			Result:        v,
		}
		c.Data["json"] = response
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Customer_account_anomalies
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customer_account_anomalies
// @Failure 403
// @router / [get]
func (c *Customer_account_anomaliesController) GetAll() {
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

	response := responses.CustomerAccountAnomaliesResponse{}

	l, err := models.GetAllCustomer_account_anomalies(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("Error fetching account anomalies: ", err)
		response = responses.CustomerAccountAnomaliesResponse{
			StatusCode:    "502",
			StatusMessage: "Error fetching account anomalies" + err.Error(),
			Result:        nil,
		}
		c.Data["json"] = response
	} else {
		logs.Info("Account anomalies fetched successfully")
		accountAnomaly := []models.Customer_account_anomalies{}
		for _, v := range l {
			accountAnomaly = append(accountAnomaly, v.(models.Customer_account_anomalies))
		}

		response = responses.CustomerAccountAnomaliesResponse{
			StatusCode:    "200",
			StatusMessage: "Account anomalies fetched successfully",
			Result:        &accountAnomaly,
		}
		c.Data["json"] = response
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Customer_account_anomalies
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Customer_account_anomalies	true		"body for Customer_account_anomalies content"
// @Success 200 {object} models.Customer_account_anomalies
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Customer_account_anomaliesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Customer_account_anomalies{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateCustomer_account_anomaliesById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customer_account_anomalies
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Customer_account_anomaliesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCustomer_account_anomalies(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
