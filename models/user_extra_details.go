package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UserExtraDetails struct {
	Active       int       `orm:"column(active);null"`
	Branch       int       `orm:"column(branch);null"`
	CreatedBy    int       `orm:"column(created_by);null"`
	DateCreated  time.Time `orm:"column(date_created);type(datetime);null;auto_now_add"`
	DateModified time.Time `orm:"column(date_modified);type(datetime);null"`
	ModifiedBy   int       `orm:"column(modified_by);null"`
	Nickname     string    `orm:"column(nickname);size(100);null"`
	ShopId       int       `orm:"column(shop_id);null"`
	Id           int       `orm:"column(user_details_id);auto"`
	UserId       int       `orm:"column(user_id);null"`
}

func (t *UserExtraDetails) TableName() string {
	return "user_extra_details"
}

func init() {
	orm.RegisterModel(new(UserExtraDetails))
}

// AddUserExtraDetails insert a new UserExtraDetails into database and returns
// last inserted Id on success.
func AddUserExtraDetails(m *UserExtraDetails) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserExtraDetailsById retrieves UserExtraDetails by Id. Returns error if
// Id doesn't exist
func GetUserExtraDetailsById(id int) (v *UserExtraDetails, err error) {
	o := orm.NewOrm()
	v = &UserExtraDetails{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserExtraDetails retrieves all UserExtraDetails matches certain condition. Returns empty list if
// no records exist
func GetAllUserExtraDetails(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserExtraDetails))
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

	var l []UserExtraDetails
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

// UpdateUserExtraDetails updates UserExtraDetails by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserExtraDetailsById(m *UserExtraDetails) (err error) {
	o := orm.NewOrm()
	v := UserExtraDetails{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserExtraDetails deletes UserExtraDetails by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserExtraDetails(id int) (err error) {
	o := orm.NewOrm()
	v := UserExtraDetails{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserExtraDetails{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
