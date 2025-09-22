package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToCustomerAccountHistoriesTable_20250922_194643 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToCustomerAccountHistoriesTable_20250922_194643{}
	m.Created = "20250922_194643"

	migration.Register("AddColumnToCustomerAccountHistoriesTable_20250922_194643", m)
}

// Run the migrations
func (m *AddColumnToCustomerAccountHistoriesTable_20250922_194643) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE customer_account_history ADD COLUMN reason VARCHAR(255) NULL AFTER credit_amount;")

}

// Reverse the migrations
func (m *AddColumnToCustomerAccountHistoriesTable_20250922_194643) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
