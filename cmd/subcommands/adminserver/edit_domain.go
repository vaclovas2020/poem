package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC EditDomain */
func (srv *adminServer) EditDomain(_ context.Context, in *admin.AdminDomain) (result *admin.DomainResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err = srv.cmd.editDomain(db, in)
	return
}

/* edit domain on mysql database */
func (p *adminServerCmd) editDomain(db *sql.DB, req *admin.AdminDomain) (result *admin.DomainResponse, err error) {
	_, err = runtime.ExecDb(db, req, func(db *sql.DB, req *admin.AdminDomain) (sql.Result, error) {
		return db.Exec("UPDATE poem_domain SET poem_domain.user_domain = ? WHERE poem_domain.user_id = ?;", req.Domain, req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = new(admin.DomainResponse)
	result.Success = true
	result.Domain = req
	return
}
