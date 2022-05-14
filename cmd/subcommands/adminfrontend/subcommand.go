/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

/* admin frontend subcommand */
package adminfrontend

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/google/subcommands"
)

/* Admin frontend command struct */
type adminFrontendCmd struct {
	host          string // server hostname
	port          int    // server port number
	gRPCAdminHost string // admin gRPC admin hostname
	gRPCAdminPort int    // admin gRPC admin port number
	gRPCOauthHost string // admin gRPC oauth hostname
	gRPCOauthPort int    // admin gRPC oauth port number
	mysqlHost     string // mysql hostname
	mysqlPort     int    // mysql port
	mysqlUser     string // mysql username
	mysqlPassword string // mysql password
	mysqlDatabase string // mysql database name
	hashKey       string // session hash key
	cryptoKey     string // session crypto key
}

func (*adminFrontendCmd) Name() string     { return "admin-frontend" }
func (*adminFrontendCmd) Synopsis() string { return "Start admin frontend server" }
func (*adminFrontendCmd) Usage() string {
	return `admin-frontend [-host] [-port] [-admin-grpc-host] [-admin-grpc-port] [-mysql-host] [-mysql-port] [-mysql-user] [-mysql-password] [-mysql-database] [-hash-key] [-crypto-key]:
	Start admin frontend server
`
}

/* Set subcommand flags */
func (p *adminFrontendCmd) SetFlags(f *flag.FlagSet) {
	grpc_port, err := strconv.Atoi(os.Getenv("ADMIN_RPC_PORT"))
	if err != nil {
		grpc_port = 0
	}
	f.StringVar(&p.gRPCAdminHost, "admin-grpc-host", os.Getenv("ADMIN_RPC_HOST"), "admin gRPC hostname")
	f.IntVar(&p.gRPCAdminPort, "admin-grpc-port", grpc_port, "admin gRPC port")
	grpc_oauth_port, err := strconv.Atoi(os.Getenv("OAUTH_RPC_PORT"))
	if err != nil {
		grpc_port = 0
	}
	f.StringVar(&p.gRPCOauthHost, "oauth-grpc-host", os.Getenv("OAUTH_RPC_HOST"), "oauth gRPC hostname")
	f.IntVar(&p.gRPCOauthPort, "oauth-grpc-port", grpc_oauth_port, "oauth gRPC port")
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		port = 0
	}
	mysql_port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		mysql_port = 0
	}
	f.StringVar(&p.host, "host", os.Getenv("SERVER_HOST"), "server hostname")
	f.IntVar(&p.port, "port", port, "server port")
	f.StringVar(&p.mysqlHost, "mysql-host", os.Getenv("MYSQL_HOST"), "mysql hostname")
	f.IntVar(&p.mysqlPort, "mysql-port", mysql_port, "mysql port")
	f.StringVar(&p.mysqlUser, "mysql-user", os.Getenv("MYSQL_USER"), "mysql user")
	f.StringVar(&p.mysqlPassword, "mysql-password", os.Getenv("MYSQL_PASSWORD"), "mysql password")
	f.StringVar(&p.mysqlDatabase, "mysql-database", os.Getenv("MYSQL_DATABASE"), "mysql database name")
	f.StringVar(&p.hashKey, "hash-key", os.Getenv("HASH_KEY"), "session hash key")
	f.StringVar(&p.cryptoKey, "crypto-key", os.Getenv("CRYPTO_KEY"), "session crypto key")
}

/* Execute subcommand */
func (p *adminFrontendCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("Starting server...")
	if p.port > 0 && p.host != "" {
		p.runServer()
	}
	return subcommands.ExitSuccess
}

/* Register admin-frontend subcommand */
func RegisterSubcommand() {
	subcommands.Register(&adminFrontendCmd{}, "Server")
}
