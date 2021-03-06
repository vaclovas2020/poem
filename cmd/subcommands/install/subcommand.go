/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

/* CMS install subcommand */
package install

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/google/subcommands"
)

/* install subcommand struct */
type installCmd struct {
	mysqlHost     string // mysql hostname
	mysqlPort     int    // mysql port
	mysqlUser     string // mysql username
	mysqlPassword string // mysql password
	mysqlDatabase string // mysql database name
}

func (*installCmd) Name() string     { return "install" }
func (*installCmd) Synopsis() string { return "Install CMS database" }
func (*installCmd) Usage() string {
	return `install [-mysql-host] [-mysql-port] [-mysql-user] [-mysql-password] [-mysql-database]:
	Install CMS database
`
}

/* Set subcommand flags */
func (p *installCmd) SetFlags(f *flag.FlagSet) {
	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		port = 0
	}
	f.StringVar(&p.mysqlHost, "mysql-host", os.Getenv("MYSQL_HOST"), "mysql hostname")
	f.IntVar(&p.mysqlPort, "mysql-port", port, "mysql port")
	f.StringVar(&p.mysqlUser, "mysql-user", os.Getenv("MYSQL_USER"), "mysql user")
	f.StringVar(&p.mysqlPassword, "mysql-password", os.Getenv("MYSQL_PASSWORD"), "mysql password")
	f.StringVar(&p.mysqlDatabase, "mysql-database", os.Getenv("MYSQL_DATABASE"), "mysql database name")
}

/* Execute subcommand */
func (p *installCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Println("Installing CMS database...")
	if p.mysqlHost != "" && p.mysqlPort > 0 && p.mysqlUser != "" &&
		p.mysqlPassword != "" && p.mysqlDatabase != "" {
		db, err := p.openDBConnection()
		if err != nil {
			fmt.Println(err.Error())
			return subcommands.ExitFailure
		}
		p.installDatabase(db)
	}
	return subcommands.ExitSuccess
}

/* Register install subcommand */
func RegisterSubcommand() {
	subcommands.Register(&installCmd{}, "Setup")
}
