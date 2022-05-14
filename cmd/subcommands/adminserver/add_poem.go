/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC AddPoem */
func (srv *adminServer) AddPoem(_ context.Context, poem *admin.AdminPoem) (response *admin.PoemResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	_, err = runtime.ExecDb(db, poem, func(db *sql.DB, poem *admin.AdminPoem) (sql.Result, error) {
		return db.Exec("INSERT INTO `poem_poems`(category_id,title,text) VALUES (?,?,?);", poem.CategoryId, poem.Title, poem.Text)
	})
	if err != nil {
		return nil, err
	}
	response = &admin.PoemResponse{Success: true, Poem: poem}
	return response, nil
}
