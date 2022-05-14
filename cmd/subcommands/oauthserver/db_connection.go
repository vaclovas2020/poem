package oauthserver

import (
	"database/sql"
	"fmt"
)

/* Connect to mysql database */
func (p *oauthServerCmd) openDBConnection() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", p.mysqlUser, p.mysqlPassword, p.mysqlHost, p.mysqlPort, p.mysqlDatabase))
}
