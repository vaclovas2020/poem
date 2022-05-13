package adminserver

import "database/sql"

/* Execute database query */
func execDb[T any](db *sql.DB, obj T, execHandler func(*sql.DB, T) (sql.Result, error)) error {
	defer db.Close()
	_, err := execHandler(db, obj)
	if err != nil {
		return err
	}
	return nil
}
