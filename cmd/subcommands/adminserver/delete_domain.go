package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC DeleteDomain */
func (srv *adminServer) DeleteDomain(_ context.Context, req *admin.AdminDomain) (result *admin.DomainResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err = srv.cmd.deleteDomain(db, req)
	return
}

/* delete domain from mysql database */
func (p *adminServerCmd) deleteDomain(db *sql.DB, req *admin.AdminDomain) (result *admin.DomainResponse, err error) {
	_, err = runtime.ExecDb(db, req, func(db *sql.DB, req *admin.AdminDomain) (sql.Result, error) {
		return db.Exec("DELETE FROM poem_domain WHERE poem_domain.user_id = ?;", req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = new(admin.DomainResponse)
	result.Success = true
	return
}
