package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToAccountsTable_20250918_160754 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToAccountsTable_20250918_160754{}
	m.Created = "20250918_160754"

	migration.Register("AddColumnToAccountsTable_20250918_160754", m)
}

// Run the migrations
func (m *AddColumnToAccountsTable_20250918_160754) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE `customer_accounts` ADD COLUMN `account_type` VARCHAR(255) DEFAULT NULL, ADD COLUMN `reference` VARCHAR(100) DEFAULT NULL;")

}

// Reverse the migrations
func (m *AddColumnToAccountsTable_20250918_160754) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
