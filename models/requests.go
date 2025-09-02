package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Requests struct {
	ApiRequestId     int        `orm:"column(api_request_id);null" description:"API Request ID"`
	CallbackResponse string     `orm:"column(callback_response);null"`
	CreatedAt        time.Time  `orm:"column(created_at);type(datetime);auto_now_add"`
	CustId           *Customers `orm:"column(cust_id);rel(fk)"`
	Request          string     `orm:"column(request)"`
	RequestAmount    float64    `orm:"column(request_amount);digits(10);decimals(2)"`
	RequestDate      time.Time  `orm:"column(request_date);type(datetime)"`
	Id               int        `orm:"column(request_id);auto"`
	RequestResponse  string     `orm:"column(request_response);null"`
	RequestStatus    string     `orm:"column(request_status);size(255)"`
	RequestType      string     `orm:"column(request_type);size(255)"`
	UpdatedAt        time.Time  `orm:"column(updated_at);type(datetime);auto_now_add"`
}

func (t *Requests) TableName() string {
	return "requests"
}

func init() {
	orm.RegisterModel(new(Requests))
}

// AddRequests insert a new Requests into database and returns
// last inserted Id on success.
func AddRequests(m *Requests) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRequestsById retrieves Requests by Id. Returns error if
// Id doesn't exist
func GetRequestsById(id int) (v *Requests, err error) {
	o := orm.NewOrm()
	v = &Requests{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRequests retrieves all Requests matches certain condition. Returns empty list if
// no records exist
func GetAllRequests(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Requests))
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

	var l []Requests
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

// UpdateRequests updates Requests by Id and returns error if
// the record to be updated doesn't exist
func UpdateRequestsById(m *Requests) (err error) {
	o := orm.NewOrm()
	v := Requests{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRequests deletes Requests by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRequests(id int) (err error) {
	o := orm.NewOrm()
	v := Requests{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Requests{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
