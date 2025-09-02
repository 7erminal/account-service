package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type AuthUsers struct {
	Active              int               `orm:"column(active);null"`
	Address             string            `orm:"column(address);size(255);null"`
	CreatedBy           int               `orm:"column(created_by);null"`
	DateCreated         time.Time         `orm:"column(date_created);type(datetime);null;auto_now_add"`
	DateModified        time.Time         `orm:"column(date_modified);type(datetime);null"`
	Dob                 time.Time         `orm:"column(dob);type(date);null"`
	Email               string            `orm:"column(email);size(255);null"`
	FullName            string            `orm:"column(full_name);size(255)"`
	Gender              string            `orm:"column(gender);size(10);null"`
	IdNumber            string            `orm:"column(id_number);size(100);null"`
	IdType              string            `orm:"column(id_type);size(5);null"`
	ImagePath           string            `orm:"column(image_path);size(200);null"`
	IsVerified          int8              `orm:"column(is_verified);null"`
	LoginFailedAttempts int               `orm:"column(login_failed_attempts);null"`
	MaritalStatus       string            `orm:"column(marital_status);size(20);null"`
	ModifiedBy          int               `orm:"column(modified_by);null"`
	Password            string            `orm:"column(password);size(255)"`
	PhoneNumber         string            `orm:"column(phone_number);size(255);null"`
	Role                *Roles            `orm:"column(role);rel(fk)"`
	UserDetailsId       *UserExtraDetails `orm:"column(user_details_id);rel(fk)"`
	Id                  int               `orm:"column(user_id);auto"`
	UserType            int               `orm:"column(user_type);null"`
	Username            string            `orm:"column(username);size(40);null"`
}

func (t *AuthUsers) TableName() string {
	return "auth_users"
}

func init() {
	orm.RegisterModel(new(AuthUsers))
}

// AddAuthUsers insert a new AuthUsers into database and returns
// last inserted Id on success.
func AddAuthUsers(m *AuthUsers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAuthUsersById retrieves AuthUsers by Id. Returns error if
// Id doesn't exist
func GetAuthUsersById(id int) (v *AuthUsers, err error) {
	o := orm.NewOrm()
	v = &AuthUsers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAuthUsers retrieves all AuthUsers matches certain condition. Returns empty list if
// no records exist
func GetAllAuthUsers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AuthUsers))
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

	var l []AuthUsers
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

// UpdateAuthUsers updates AuthUsers by Id and returns error if
// the record to be updated doesn't exist
func UpdateAuthUsersById(m *AuthUsers) (err error) {
	o := orm.NewOrm()
	v := AuthUsers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAuthUsers deletes AuthUsers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAuthUsers(id int) (err error) {
	o := orm.NewOrm()
	v := AuthUsers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AuthUsers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
