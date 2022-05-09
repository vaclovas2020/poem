/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package install

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/* Install CMS database */
func (p *installCmd) installDatabase() {

}

/* Connect to mysql database */
func (p *installCmd) openDBConnection() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", p.mysqlUser, p.mysqlPassword, p.mysqlHost, p.mysqlPort, p.mysqlDatabase))
}
