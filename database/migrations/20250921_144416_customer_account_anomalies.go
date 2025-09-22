package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CustomerAccountAnomalies_20250921_144416 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CustomerAccountAnomalies_20250921_144416{}
	m.Created = "20250921_144416"

	migration.Register("CustomerAccountAnomalies_20250921_144416", m)
}

// Run the migrations
func (m *CustomerAccountAnomalies_20250921_144416) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE customer_account_anomalies(`id` int(11) NOT NULL AUTO_INCREMENT,`account_number` varchar(255) NOT NULL,`amount` float NOT NULL,`desc` varchar(100) NOT NULL,`balance` float NOT NULL,`checked_balance` float NOT NULL,`date_created` datetime NOT NULL,`date_modified` datetime NOT NULL,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *CustomerAccountAnomalies_20250921_144416) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `customer_account_anomalies`")
}
