package adminserver

import "database/sql"

/* Execute database query */
func execDb[T any](p *adminServerCmd, obj T, execHandler func(*sql.DB, T) (sql.Result, error)) error {
	db, err := p.openDBConnection()
	if err != nil {
		return err
	}
	_, err = execHandler(db, obj)
	if err != nil {
		return err
	}
	return nil
}
