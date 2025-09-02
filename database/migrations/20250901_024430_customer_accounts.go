package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CustomerAccounts_20250901_024430 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CustomerAccounts_20250901_024430{}
	m.Created = "20250901_024430"

	migration.Register("CustomerAccounts_20250901_024430", m)
}

// Run the migrations
func (m *CustomerAccounts_20250901_024430) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE customer_accounts(`customer_account_id` int(11) NOT NULL AUTO_INCREMENT,`customer_id` int, `account_number` varchar(100) NOT NULL,`account_alias` varchar(255) NOT NULL,`balance` float DEFAULT 0.0,`frozen_amount` float DEFAULT 0.0,`balance_before` float DEFAULT 0.0,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 0,PRIMARY KEY (`customer_account_id`), FOREIGN KEY (customer_id) REFERENCES customers(customer_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *CustomerAccounts_20250901_024430) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `customer_accounts`")
}
