package oauthserver

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
	"webimizer.dev/poem/oauth"
	"webimizer.dev/poem/runtime"
)

/* gRPC NewUser */
func (srv *oAuthServer) NewUser(_ context.Context, request *oauth.AuthRequest) (response *oauth.AuthResponse, err error) {
	db, err := srv.cmd.openDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	row := runtime.QueryRowDb(db, request, func(db *sql.DB, request *oauth.AuthRequest) *sql.Row {
		return db.QueryRow("SELECT user_email FROM `poem_users` WHERE user_email = ?;", request.Email)
	})
	response = new(oauth.AuthResponse)
	var email string
	err = row.Scan(&email)
	if err == nil { // user already exists
		response.Success = false
		return response, nil
	}
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		response.Success = false
		return response, err
	}
	result, err := runtime.ExecDb(db, request, func(db *sql.DB, request *oauth.AuthRequest) (sql.Result, error) {
		return db.Exec("INSERT INTO `poem_users`(user_email, password_hash, user_role) VALUES(?,?,?);", request.Email, string(hashBytes), request.Role.String())
	})
	if err != nil {
		response.Success = false
		return response, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		response.Success = false
		return response, err
	}
	response.Success = true
	response.User = &oauth.User{UserId: userId, Email: request.Email, Role: request.Role}
	return response, nil
}
