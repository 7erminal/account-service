package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type PasswordResetTokens struct {
	CreatedAt time.Time `orm:"column(created_at);type(timestamp);null"`
	Id        int       `orm:"column(email);pk"`
	Token     string    `orm:"column(token);size(255)"`
}

func (t *PasswordResetTokens) TableName() string {
	return "password_reset_tokens"
}

func init() {
	orm.RegisterModel(new(PasswordResetTokens))
}

// AddPasswordResetTokens insert a new PasswordResetTokens into database and returns
// last inserted Id on success.
func AddPasswordResetTokens(m *PasswordResetTokens) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPasswordResetTokensById retrieves PasswordResetTokens by Id. Returns error if
// Id doesn't exist
func GetPasswordResetTokensById(id int) (v *PasswordResetTokens, err error) {
	o := orm.NewOrm()
	v = &PasswordResetTokens{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPasswordResetTokens retrieves all PasswordResetTokens matches certain condition. Returns empty list if
// no records exist
func GetAllPasswordResetTokens(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PasswordResetTokens))
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

	var l []PasswordResetTokens
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

// UpdatePasswordResetTokens updates PasswordResetTokens by Id and returns error if
// the record to be updated doesn't exist
func UpdatePasswordResetTokensById(m *PasswordResetTokens) (err error) {
	o := orm.NewOrm()
	v := PasswordResetTokens{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePasswordResetTokens deletes PasswordResetTokens by Id and returns error if
// the record to be deleted doesn't exist
func DeletePasswordResetTokens(id int) (err error) {
	o := orm.NewOrm()
	v := PasswordResetTokens{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PasswordResetTokens{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
