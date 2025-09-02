package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UsersBkp struct {
	Branchid        int64     `orm:"column(branchid)"`
	Corpid          int64     `orm:"column(corpid)"`
	CreatedAt       time.Time `orm:"column(created_at);type(timestamp);null"`
	Email           string    `orm:"column(email);size(255)"`
	EmailVerifiedAt time.Time `orm:"column(email_verified_at);type(timestamp);null"`
	Id              int       `orm:"column(id);auto"`
	Name            string    `orm:"column(name);size(255)"`
	Password        string    `orm:"column(password);size(255)"`
	RememberToken   string    `orm:"column(remember_token);size(100);null"`
	UpdatedAt       time.Time `orm:"column(updated_at);type(timestamp);null"`
	Username        string    `orm:"column(username);size(255)"`
}

func (t *UsersBkp) TableName() string {
	return "users_bkp"
}

func init() {
	orm.RegisterModel(new(UsersBkp))
}

// AddUsersBkp insert a new UsersBkp into database and returns
// last inserted Id on success.
func AddUsersBkp(m *UsersBkp) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersBkpById retrieves UsersBkp by Id. Returns error if
// Id doesn't exist
func GetUsersBkpById(id int) (v *UsersBkp, err error) {
	o := orm.NewOrm()
	v = &UsersBkp{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsersBkp retrieves all UsersBkp matches certain condition. Returns empty list if
// no records exist
func GetAllUsersBkp(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UsersBkp))
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

	var l []UsersBkp
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

// UpdateUsersBkp updates UsersBkp by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersBkpById(m *UsersBkp) (err error) {
	o := orm.NewOrm()
	v := UsersBkp{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsersBkp deletes UsersBkp by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsersBkp(id int) (err error) {
	o := orm.NewOrm()
	v := UsersBkp{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UsersBkp{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
