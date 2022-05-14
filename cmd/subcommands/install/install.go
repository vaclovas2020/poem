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
type doInstallHandler func(db *sql.DB) error

/* Install CMS database */
func (p *installCmd) installDatabase(db *sql.DB) error {
	defer db.Close()
	err := p.doInstall(db, []doInstallHandler{
		doInstallHandler(p.createUserDb),
		doInstallHandler(p.createCategoriesDb),
		doInstallHandler(p.createPoemsDb),
	})
	if err != nil {
		return err
	}
	return nil
}

/* Execute installer functions */
func (p *installCmd) doInstall(db *sql.DB, handlers []doInstallHandler) error {
	for _, handler := range handlers {
		err := handler(db)
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
func (p *installCmd) createUserDb(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `poem_users` (user_id INT NOT NULL AUTO_INCREMENT, user_email VARCHAR(255) NOT NULL, password_hash VARCHAR(255) NOT NULL, user_role VARCHAR(20) NOT NULL, PRIMARY KEY (user_id), UNIQUE KEY (user_name) );")
	if err != nil {
		return err
	}
	role := "admin"
	_, err = db.Exec("REPLACE INTO `poem_users` (user_email, password_hash, user_role) VALUES (?,?,?);", p.cmsUser, p.cmsPasswordHash, role)
	if err != nil {
		return err
	}
	fmt.Println("CMS user database installed!")
	return nil
}

/* Create categories database schema */
func (p *installCmd) createCategoriesDb(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `poem_categories` (category_id INT NOT NULL AUTO_INCREMENT, name VARCHAR(100) NOT NULL, slug VARCHAR(100) NOT NULL, status VARCHAR(10) NOT NULL, PRIMARY KEY (category_id), UNIQUE KEY (slug) );")
	if err != nil {
		return err
	}
	fmt.Println("CMS categories database installed!")
	return nil
}

/* Create poems database schema */
func (p *installCmd) createPoemsDb(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `poem_poems` (poem_id INT NOT NULL AUTO_INCREMENT, category_id INT NOT NULL, title VARCHAR(100) NOT NULL, text TEXT NOT NULL, PRIMARY KEY (poem_id) );")
	if err != nil {
		return err
	}
	fmt.Println("CMS poems database installed!")
	return nil
}
