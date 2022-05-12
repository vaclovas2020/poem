/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package adminserver

import (
	"context"

	"webimizer.dev/poem/admin"
)

/* gRPC AddCategory */
func (srv *adminServer) AddCategory(_ context.Context, category *admin.AdminCategory) (response *admin.CategoryResponse, err error) {
	err = srv.cmd.addCategory(category)
	if err != nil {
		return nil, err
	}
	response = &admin.CategoryResponse{Success: true, Category: category}
	return response, nil
}

/* Insert category to database */
func (p *adminServerCmd) addCategory(category *admin.AdminCategory) error {
	db, err := p.openDBConnection()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO `poem_categories`(name,slug,status) VALUES (?,?,?);", category.Name, category.Slug, category.Status.String())
	if err != nil {
		return err
	}
	return nil
}
