package oauthserver

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"webimizer.dev/poem/oauth"
	"webimizer.dev/poem/runtime"
)

/* gRPC AuthUser */
func (srv *oAuthServer) AuthUser(_ context.Context, request *oauth.AuthRequest) (response *oauth.AuthResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	row := runtime.QueryRowDb(db, request, func(db *sql.DB, request *oauth.AuthRequest) *sql.Row {
		return db.QueryRow("SELECT user_name, password_hash, user_role FROM `poem_users` WHERE user_name = ? AND user_role = ?;", request.Username, request.Role.String())
	})
	response = new(oauth.AuthResponse)
	var username string
	var role string
	var password_hash string
	err = row.Scan(&username, &password_hash, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			response.Success = false
			return response, nil
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(request.Password))
	if err != nil {
		response.Success = false
		return response, nil
	}
	response.Success = true
	response.User = &oauth.User{Name: username, Role: request.Role}
	return response, nil
}
