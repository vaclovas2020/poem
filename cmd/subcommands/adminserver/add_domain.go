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

/* gRPC AddDomain */
func (srv *adminServer) AddDomain(_ context.Context, in *admin.AdminDomain) (response *admin.DomainResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	_, err = runtime.ExecDb(db, in, func(db *sql.DB, in *admin.AdminDomain) (sql.Result, error) {
		return db.Exec("REPLACE INTO `poem_domain`(user_id, user_domain) VALUES (?,?);", in.UserId, in.Domain)
	})
	if err != nil {
		return nil, err
	}
	response = &admin.DomainResponse{Success: true, Domain: in}
	return response, nil
}
