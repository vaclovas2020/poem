/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package install

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/* Install function handler */
type doInstallHandler func() error

/* Install CMS database */
func (p *installCmd) installDatabase() {
	err := p.doInstall([]doInstallHandler{
		doInstallHandler(p.createUserDb),
	})
	if err != nil {
		panic(err)
	}
}

/* Execute installer functions */
func (p *installCmd) doInstall(handlers []doInstallHandler) error {
	for _, handler := range handlers {
		err := handler()
		if err != nil {
			return err
		}
	}
	return nil
}

/* Connect to mysql database */
func (p *installCmd) openDBConnection() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", p.mysqlUser, p.mysqlPassword, p.mysqlHost, p.mysqlPort, p.mysqlDatabase))
}

/* Create user database schema */
func (p *installCmd) createUserDb() error {
	db, err := p.openDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `poem_users` (user_id INT NOT NULL AUTO_INCREMENT, user_name VARCHAR(100), password_hash VARCHAR(255), PRIMARY KEY (user_id) );")
	if err != nil {
		return err
	}
	return nil
}