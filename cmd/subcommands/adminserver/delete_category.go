package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC DeleteCategory */
func (srv *adminServer) DeleteCategory(_ context.Context, req *admin.DeleteCategoryRequest) (result *admin.DeleteCategoryResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err = srv.cmd.deleteCategory(db, req)
	return
}

/* delete category from mysql database */
func (p *adminServerCmd) deleteCategory(db *sql.DB, req *admin.DeleteCategoryRequest) (result *admin.DeleteCategoryResponse, err error) {
	_, err = runtime.ExecDb(db, req, func(db *sql.DB, req *admin.DeleteCategoryRequest) (sql.Result, error) {
		return db.Exec("DELETE FROM poem_categories WHERE poem_categories.category_id = ? AND poem_categories.user_id = ?;", req.CategoryId, req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = new(admin.DeleteCategoryResponse)
	result.Success = true
	return
}
