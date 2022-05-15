package install

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

/* Create mysql database mock */
func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

var p = installCmd{}

func TestCreateUserDb(t *testing.T) {
	db, mock := newMock()
	sql := "CREATE TABLE IF NOT EXISTS `poem_users` \\(user_id INT NOT NULL AUTO_INCREMENT, user_email VARCHAR\\(255\\) NOT NULL, password_hash VARCHAR\\(255\\) NOT NULL, user_role VARCHAR\\(20\\) NOT NULL, PRIMARY KEY \\(user_id\\), UNIQUE KEY \\(user_email\\) \\);"
	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 0))
	err := p.createUserDb(db)
	assert.NoError(t, err)
	mock.ExpectExec(sql).WillReturnError(fmt.Errorf("Testing error handler"))
	err = p.createUserDb(db)
	assert.Error(t, err)
}

func TestCreateCategoriesDb(t *testing.T) {
	db, mock := newMock()
	sql := "CREATE TABLE IF NOT EXISTS `poem_categories` \\(category_id INT NOT NULL AUTO_INCREMENT, user_id INT NOT NULL, name VARCHAR\\(100\\) NOT NULL, slug VARCHAR\\(100\\) NOT NULL, status VARCHAR\\(10\\) NOT NULL, PRIMARY KEY \\(category_id\\), UNIQUE KEY \\(slug\\) \\);"
	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 0))
	err := p.createCategoriesDb(db)
	assert.NoError(t, err)
	mock.ExpectExec(sql).WillReturnError(fmt.Errorf("Testing error handler"))
	err = p.createCategoriesDb(db)
	assert.Error(t, err)
}

func TestCreatePoemDb(t *testing.T) {
	db, mock := newMock()
	sql := "CREATE TABLE IF NOT EXISTS `poem_poems` \\(poem_id INT NOT NULL AUTO_INCREMENT, category_id INT NOT NULL, user_id INT NOT NULL, title VARCHAR\\(100\\) NOT NULL, text TEXT NOT NULL, PRIMARY KEY \\(poem_id\\) \\);"
	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 0))
	err := p.createPoemsDb(db)
	assert.NoError(t, err)
	mock.ExpectExec(sql).WillReturnError(fmt.Errorf("Testing error handler"))
	err = p.createPoemsDb(db)
	assert.Error(t, err)
}

func TestDoInstall(t *testing.T) {
	db, _ := newMock()
	err := p.doInstall(db, []doInstallHandler{
		doInstallHandler(func(db *sql.DB) error { return nil }),
		doInstallHandler(func(db *sql.DB) error { return nil }),
		doInstallHandler(func(db *sql.DB) error { return nil }),
	})
	assert.NoError(t, err)

	err = p.doInstall(db, []doInstallHandler{
		doInstallHandler(func(db *sql.DB) error { return nil }),
		doInstallHandler(func(db *sql.DB) error { return fmt.Errorf("Test error") }),
		doInstallHandler(func(db *sql.DB) error { return nil }),
	})
	assert.Error(t, err)
}

func TestInstallDatabase(t *testing.T) {
	db, mock := newMock()
	sql := "CREATE TABLE IF NOT EXISTS `poem_users` \\(user_id INT NOT NULL AUTO_INCREMENT, user_email VARCHAR\\(255\\) NOT NULL, password_hash VARCHAR\\(255\\) NOT NULL, user_role VARCHAR\\(20\\) NOT NULL, PRIMARY KEY \\(user_id\\), UNIQUE KEY \\(user_email\\) \\);"
	mock.ExpectExec(sql).WillReturnResult(sqlmock.NewResult(0, 0))
	sql3 := "CREATE TABLE IF NOT EXISTS `poem_categories` \\(category_id INT NOT NULL AUTO_INCREMENT, user_id INT NOT NULL, name VARCHAR\\(100\\) NOT NULL, slug VARCHAR\\(100\\) NOT NULL, status VARCHAR\\(10\\) NOT NULL, PRIMARY KEY \\(category_id\\), UNIQUE KEY \\(slug\\) \\);"
	mock.ExpectExec(sql3).WillReturnResult(sqlmock.NewResult(0, 0))
	sql4 := "CREATE TABLE IF NOT EXISTS `poem_poems` \\(poem_id INT NOT NULL AUTO_INCREMENT, category_id INT NOT NULL, user_id INT NOT NULL, title VARCHAR\\(100\\) NOT NULL, text TEXT NOT NULL, PRIMARY KEY \\(poem_id\\) \\);"
	mock.ExpectExec(sql4).WillReturnResult(sqlmock.NewResult(0, 0))
	err := p.installDatabase(db)
	assert.NoError(t, err)
	mock.ExpectExec(sql).WillReturnError(fmt.Errorf("Testing error handler"))
	err = p.installDatabase(db)
	assert.Error(t, err)
}

func TestOpenDBConnection(t *testing.T) {
	p.mysqlHost = "fake"
	p.mysqlPort = 45
	p.mysqlUser = "user"
	p.mysqlPassword = "password"
	p.mysqlDatabase = "test"
	_, err := p.openDBConnection()
	assert.NoError(t, err)
}
