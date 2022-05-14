package runtime

import "database/sql"

/* Execute database query */
func QueryRowDb[T any](db *sql.DB, obj T, queryRowHandler func(*sql.DB, T) *sql.Row) *sql.Row {
	defer db.Close()
	return queryRowHandler(db, obj)
}
