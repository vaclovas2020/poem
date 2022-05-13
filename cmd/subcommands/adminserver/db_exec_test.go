package adminserver

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

type newDbObj struct {
	Test string
}

func TestExcecDb(t *testing.T) {
	db, mock := newMock()
	mock.ExpectExec("Test").WillReturnResult(sqlmock.NewResult(0, 1))
	obj := &newDbObj{Test: "Test"}
	err := execDb(db, obj, func(db *sql.DB, obj *newDbObj) (sql.Result, error) { return db.Exec(obj.Test) })
	assert.NoError(t, err)
	mock.ExpectExec("Test").WillReturnError(fmt.Errorf("Testing error handler"))
	err = execDb(db, obj, func(db *sql.DB, obj *newDbObj) (sql.Result, error) { return db.Exec(obj.Test) })
	assert.Error(t, err)
}
