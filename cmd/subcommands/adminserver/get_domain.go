package adminserver

import (
	"context"
	"database/sql"

	"webimizer.dev/poem/admin"
	"webimizer.dev/poem/runtime"
)

/* gRPC EditDomain */
func (srv *adminServer) GetDomain(_ context.Context, in *admin.GetAdminDomain) (result *admin.DomainResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err = srv.cmd.getDomain(db, in)
	return
}

/* edit domain on mysql database */
func (p *adminServerCmd) getDomain(db *sql.DB, req *admin.GetAdminDomain) (result *admin.DomainResponse, err error) {
	row := runtime.QueryRowDb(db, req, func(db *sql.DB, req *admin.GetAdminDomain) *sql.Row {
		return db.QueryRow("SELECT poem_domain.user_domain FROM poem_domain WHERE poem_domain.user_id = ?;", req.UserId)
	})
	if err != nil {
		return nil, err
	}
	result = new(admin.DomainResponse)
	var domainStr string
	err = row.Scan(&domainStr)
	if err != nil {
		if err == sql.ErrNoRows {
			result.Success = false
			return result, nil
		}
		return nil, err
	}
	result.Success = true
	result.Domain = new(admin.AdminDomain)
	result.Domain.UserId = req.UserId
	result.Domain.Domain = domainStr
	return
}
