package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC EditPoem */
func (srv *adminServer) EditPoem(_ context.Context, req *admin.AdminPoemEdit) (result *admin.PoemEditResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err = srv.cmd.editPoem(db, req)
	return
}

/* edit poem on mysql database */
func (p *adminServerCmd) editPoem(db *sql.DB, req *admin.AdminPoemEdit) (result *admin.PoemEditResponse, err error) {
	_, err = runtime.ExecDb(db, req, func(db *sql.DB, req *admin.AdminPoemEdit) (sql.Result, error) {
		return db.Exec("UPDATE poem_poems SET poem_poems.category_id = ?, poem_poems.title = ?, poem_poems.text = ? WHERE poem_poems.poem_id = ? AND poem_poems.user_id = ?;", req.CategoryId, req.Title, req.Text, req.PoemId, req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = new(admin.PoemEditResponse)
	result.Success = true
	result.Poem = req
	return
}
