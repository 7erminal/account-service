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
	m.SQL("CREATE TABLE account_anomalies(`account_anomaly_id` int(11) NOT NULL AUTO_INCREMENT,`account_id` varchar(255) NOT NULL,`amount` float DEFAULT 0.0,`desc` varchar(100) DEFAULT NULL,`statement` varchar(100) DEFAULT NULL,`balance` float DEFAULT 0.0,`checked_balance` float DEFAULT 0.0,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`account_anomaly_id`))")
}

// Reverse the migrations
func (m *AccountAnomalies_20250921_141712) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `account_anomalies`")
}
