/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

/* poems server subcommand */
package poemsserver

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/google/subcommands"
)

/* Poems server command struct */
type poemsServerCmd struct {
	host          string // server hostname
	port          int    // server port number
	mysqlHost     string // mysql hostname
	mysqlPort     int    // mysql port
	mysqlUser     string // mysql username
	mysqlPassword string // mysql password
	mysqlDatabase string // mysql database name
}

func (*poemsServerCmd) Name() string     { return "poems-server" }
func (*poemsServerCmd) Synopsis() string { return "Start poems gRPC server" }
func (*poemsServerCmd) Usage() string {
	return `poems-server [-host] [-port]:
  Start poems gRPC server
`
}

/* Set subcommand flags */
func (p *poemsServerCmd) SetFlags(f *flag.FlagSet) {
	port, err := strconv.Atoi(os.Getenv("RPC_PORT"))
	if err != nil {
		port = 0
	}
	f.StringVar(&p.host, "host", os.Getenv("RPC_HOST"), "server hostname")
	f.IntVar(&p.port, "port", port, "server port")
	mysql_port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		port = 0
	}
	f.StringVar(&p.mysqlHost, "mysql-host", os.Getenv("MYSQL_HOST"), "mysql hostname")
	f.IntVar(&p.mysqlPort, "mysql-port", mysql_port, "mysql port")
	f.StringVar(&p.mysqlUser, "mysql-user", os.Getenv("MYSQL_USER"), "mysql user")
	f.StringVar(&p.mysqlPassword, "mysql-password", os.Getenv("MYSQL_PASSWORD"), "mysql password")
	f.StringVar(&p.mysqlDatabase, "mysql-database", os.Getenv("MYSQL_DATABASE"), "mysql database name")
}

/* Execute subcommand */
func (p *poemsServerCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("Starting server...")
	if p.port > 0 && p.host != "" {
		p.runServer()
	}
	return subcommands.ExitSuccess
}

/* Register poems-server subcommand */
func RegisterSubcommand() {
	subcommands.Register(&poemsServerCmd{}, "Server")
}
