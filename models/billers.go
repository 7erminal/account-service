package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Billers struct {
	Active            int        `orm:"column(active);null"`
	BillerCode        string     `orm:"column(biller_code);size(80)"`
	Id                int        `orm:"column(biller_id);auto"`
	BillerName        string     `orm:"column(biller_name);size(80)"`
	BillerReferenceId string     `orm:"column(biller_reference_id);size(250);null"`
	CreatedBy         int        `orm:"column(created_by);null"`
	DateCreated       time.Time  `orm:"column(date_created);type(datetime);null;auto_now_add"`
	DateModified      time.Time  `orm:"column(date_modified);type(datetime);null"`
	Description       string     `orm:"column(description);size(255);null"`
	ModifiedBy        int        `orm:"column(modified_by);null"`
	OperatorId        *Operators `orm:"column(operator_id);rel(fk)"`
}

func (t *Billers) TableName() string {
	return "billers"
}

func init() {
	orm.RegisterModel(new(Billers))
}

// AddBillers insert a new Billers into database and returns
// last inserted Id on success.
func AddBillers(m *Billers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBillersById retrieves Billers by Id. Returns error if
// Id doesn't exist
func GetBillersById(id int) (v *Billers, err error) {
	o := orm.NewOrm()
	v = &Billers{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBillers retrieves all Billers matches certain condition. Returns empty list if
// no records exist
func GetAllBillers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Billers))
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

	var l []Billers
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

// UpdateBillers updates Billers by Id and returns error if
// the record to be updated doesn't exist
func UpdateBillersById(m *Billers) (err error) {
	o := orm.NewOrm()
	v := Billers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBillers deletes Billers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBillers(id int) (err error) {
	o := orm.NewOrm()
	v := Billers{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Billers{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
