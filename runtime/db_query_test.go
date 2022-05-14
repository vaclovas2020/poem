package runtime_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"webimizer.dev/poem/runtime"
)

func TestQueryDb(t *testing.T) {
	db, mock := newMock()
	obj := &newDbObj{Test: "Test"}
	mock.ExpectQuery("Test").WillReturnRows(sqlmock.NewRows([]string{"column1"}).AddRow(obj.Test))
	rows, err := runtime.QueryDb(db, obj, func(d *sql.DB, t *newDbObj) (*sql.Rows, error) { return db.Query("Test") })
	assert.NoError(t, err)
	var value string
	rows.Next()
	err = rows.Scan(&value)
	assert.NoError(t, err)
	assert.Equal(t, obj.Test, value)
	mock.ExpectQuery("Test").WillReturnError(fmt.Errorf("Testing error handler"))
	_, err = runtime.QueryDb(db, obj, func(d *sql.DB, t *newDbObj) (*sql.Rows, error) { return db.Query("Test") })
	assert.Error(t, err)
}
