/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
)

/* gRPC AddCategory */
func (srv *adminServer) AddCategory(_ context.Context, category *admin.AdminCategory) (response *admin.CategoryResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	err = execDb(db, category, func(db *sql.DB, category *admin.AdminCategory) (sql.Result, error) {
		return db.Exec("INSERT INTO `poem_categories`(name,slug,status) VALUES (?,?,?);", category.Name, category.Slug, category.Status.String())
	})
	if err != nil {
		return nil, err
	}
	response = &admin.CategoryResponse{Success: true, Category: category}
	return response, nil
}
