package runtime

import "database/sql"

/* Execute database query */
func QueryDb[T any](db *sql.DB, obj T, queryHandler func(*sql.DB, T) (*sql.Rows, error)) (*sql.Rows, error) {
	defer db.Close()
	return queryHandler(db, obj)
}
