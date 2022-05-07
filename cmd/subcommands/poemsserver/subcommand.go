/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

/* poems server subcommand */
package poemsserver

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

/* Poems server command struct */
type poemsServerCmd struct {
	host string
	port int
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
	f.StringVar(&p.host, "host", "localhost", "server hostname")
	f.IntVar(&p.port, "host", 8044, "server port")
}

/* Execute subcommand */
func (p *poemsServerCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	return subcommands.ExitSuccess
}

/* Register poems-server subcommand */
func RegisterSubcommand() {
	subcommands.Register(&poemsServerCmd{}, "Server")
}
