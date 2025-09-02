package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type BilInsTransactions struct {
	Active                 int              `orm:"column(active);null"`
	Amount                 float32          `orm:"column(amount);null"`
	Id                     int              `orm:"column(bil_ins_transaction_id);auto"`
	BilTransactionId       *BilTransactions `orm:"column(bil_transaction_id);rel(fk)"`
	BillerId               *Billers         `orm:"column(biller_id);rel(fk)"`
	CreatedBy              int              `orm:"column(created_by);null"`
	DateCreated            time.Time        `orm:"column(date_created);type(datetime);null;auto_now_add"`
	DateModified           time.Time        `orm:"column(date_modified);type(datetime);null"`
	ModifiedBy             int              `orm:"column(modified_by);null"`
	Network                string           `orm:"column(network);size(150);null"`
	RecipientAccountNumber string           `orm:"column(recipient_account_number);size(255);null"`
	Request                string           `orm:"column(request);null"`
	Response               string           `orm:"column(response);null"`
	SenderAccountNumber    string           `orm:"column(sender_account_number);size(255);null"`
}

func (t *BilInsTransactions) TableName() string {
	return "bil_ins_transactions"
}

func init() {
	orm.RegisterModel(new(BilInsTransactions))
}

// AddBilInsTransactions insert a new BilInsTransactions into database and returns
// last inserted Id on success.
func AddBilInsTransactions(m *BilInsTransactions) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBilInsTransactionsById retrieves BilInsTransactions by Id. Returns error if
// Id doesn't exist
func GetBilInsTransactionsById(id int) (v *BilInsTransactions, err error) {
	o := orm.NewOrm()
	v = &BilInsTransactions{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBilInsTransactions retrieves all BilInsTransactions matches certain condition. Returns empty list if
// no records exist
func GetAllBilInsTransactions(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BilInsTransactions))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []BilInsTransactions
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBilInsTransactions updates BilInsTransactions by Id and returns error if
// the record to be updated doesn't exist
func UpdateBilInsTransactionsById(m *BilInsTransactions) (err error) {
	o := orm.NewOrm()
	v := BilInsTransactions{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBilInsTransactions deletes BilInsTransactions by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBilInsTransactions(id int) (err error) {
	o := orm.NewOrm()
	v := BilInsTransactions{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BilInsTransactions{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
