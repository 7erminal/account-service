package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AccountAnomalies_20250921_141712 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AccountAnomalies_20250921_141712{}
	m.Created = "20250921_141712"

	migration.Register("AccountAnomalies_20250921_141712", m)
}

// Run the migrations
func (m *AccountAnomalies_20250921_141712) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE account_anomalies(`id` int(11) NOT NULL AUTO_INCREMENT,`account_number` varchar(255) NOT NULL,`amount` float NOT NULL,`desc` varchar(100) NOT NULL,`balance` float NOT NULL,`checked_balance` float NOT NULL,`date_created` datetime NOT NULL,`date_modified` datetime NOT NULL,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *AccountAnomalies_20250921_141712) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `account_anomalies`")
}
