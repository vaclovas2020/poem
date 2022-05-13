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
	gRPChost      string // admin gRPC hostname
	gRPCport      int    // admin gRPC port number
	mysqlHost     string // mysql hostname
	mysqlPort     int    // mysql port
	mysqlUser     string // mysql username
	mysqlPassword string // mysql password
	mysqlDatabase string // mysql database name
}

func (*adminFrontendCmd) Name() string     { return "admin-frontend" }
func (*adminFrontendCmd) Synopsis() string { return "Start admin frontend server" }
func (*adminFrontendCmd) Usage() string {
	return `admin-frontend [-host] [-port] [-grpc-host] [-grpc-port] [-mysql-host] [-mysql-port] [-mysql-user] [-mysql-password] [-mysql-database]:
	Start admin frontend server
`
}

/* Set subcommand flags */
func (p *adminFrontendCmd) SetFlags(f *flag.FlagSet) {
	grpc_port, err := strconv.Atoi(os.Getenv("RPC_PORT"))
	if err != nil {
		grpc_port = 0
	}
	f.StringVar(&p.gRPChost, "grpc-host", os.Getenv("RPC_HOST"), "admin gRPC hostname")
	f.IntVar(&p.gRPCport, "grpc-port", grpc_port, "admin gRPC port")
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
