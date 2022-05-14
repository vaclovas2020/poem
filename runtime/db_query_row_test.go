package runtime_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"webimizer.dev/poem/runtime"
)

func TestQueryRowDb(t *testing.T) {
	db, mock := newMock()
	obj := &newDbObj{Test: "Test"}
	mock.ExpectQuery("Test").WillReturnRows(sqlmock.NewRows([]string{"column1"}).AddRow(obj.Test))
	row := runtime.QueryRowDb(db, obj, func(d *sql.DB, t *newDbObj) *sql.Row { return db.QueryRow("Test") })
	assert.NoError(t, row.Err())
	var value string
	err := row.Scan(&value)
	assert.NoError(t, err)
	assert.Equal(t, obj.Test, value)
	mock.ExpectQuery("Test").WillReturnError(fmt.Errorf("Testing error handler"))
	row = runtime.QueryRowDb(db, obj, func(d *sql.DB, t *newDbObj) *sql.Row { return db.QueryRow("Test") })
	assert.Error(t, row.Err())
}
