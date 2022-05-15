package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC EditCategory */
func (srv *adminServer) EditCategory(_ context.Context, req *admin.AdminCategoryEdit) (result *admin.CategoryEditResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err = srv.cmd.editCategory(db, req)
	return
}

/* edit category on mysql database */
func (p *adminServerCmd) editCategory(db *sql.DB, req *admin.AdminCategoryEdit) (result *admin.CategoryEditResponse, err error) {
	_, err = runtime.ExecDb(db, req, func(db *sql.DB, req *admin.AdminCategoryEdit) (sql.Result, error) {
		return db.Exec("UPDATE poem_categories SET poem_categories.name = ?, poem_categories.slug = ? WHERE poem_categories.category_id = ? AND poem_categories.user_id = ?;", req.Name, req.Slug, req.CategoryId, req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = new(admin.CategoryEditResponse)
	result.Success = true
	result.Category = req
	return
}
