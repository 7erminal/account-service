package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CustomerAccountHistory_20250902_101848 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CustomerAccountHistory_20250902_101848{}
	m.Created = "20250902_101848"

	migration.Register("CustomerAccountHistory_20250902_101848", m)
}

// Run the migrations
func (m *CustomerAccountHistory_20250902_101848) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE customer_account_history(`customer_account_history_id` int(11) NOT NULL AUTO_INCREMENT,`customer_account_id` int NOT NULL,`debit_amount` float DEFAULT 0.0,`credit_amount` float DEFAULT 0.0,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`customer_account_history_id`), FOREIGN KEY (customer_account_id) REFERENCES customer_accounts(customer_account_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *CustomerAccountHistory_20250902_101848) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `customer_account_history`")
}
