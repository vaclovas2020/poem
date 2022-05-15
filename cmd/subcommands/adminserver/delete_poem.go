package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC DeletePoem */
func (srv *adminServer) DeletePoem(_ context.Context, req *admin.DeletePoemRequest) (result *admin.DeletePoemResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err = srv.cmd.deletePoem(db, req)
	return
}

/* delete poem from mysql database */
func (p *adminServerCmd) deletePoem(db *sql.DB, req *admin.DeletePoemRequest) (result *admin.DeletePoemResponse, err error) {
	_, err = runtime.ExecDb(db, req, func(db *sql.DB, req *admin.DeletePoemRequest) (sql.Result, error) {
		return db.Exec("DELETE FROM poem_poems WHERE poem_poems.poem_id = ? AND poem_poems.user_id = ?;", req.PoemId, req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = new(admin.DeletePoemResponse)
	result.Success = true
	return
}
