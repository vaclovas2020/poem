package runtime

import "database/sql"

/* Execute database query */
func ExecDb[T any](db *sql.DB, obj T, execHandler func(*sql.DB, T) (sql.Result, error)) (sql.Result, error) {
	return execHandler(db, obj)
}
