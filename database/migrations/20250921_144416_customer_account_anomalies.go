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
	m.SQL("CREATE TABLE customer_account_anomalies(`customer_account_anomaly_id` int(11) NOT NULL AUTO_INCREMENT,`customer_account_id` int(11) NOT NULL,`amount` float DEFAULT 0.0,`desc` varchar(100) DEFAULT NULL,`statement` varchar(100) DEFAULT NULL,`balance` float DEFAULT 0.0,`checked_balance` float DEFAULT 0.0,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`customer_account_anomaly_id`), FOREIGN KEY (customer_account_id) REFERENCES customer_accounts(customer_account_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *CustomerAccountAnomalies_20250921_144416) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `customer_account_anomalies`")
}
