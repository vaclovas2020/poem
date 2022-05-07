/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights eserved.
*/

/* Package for CLI subcommands registration */
package subcommands

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"webimizer.dev/poem/cmd/subcommands/poemsserver"
)

/* Register all application's subcommands and execute choosen subcommand */
func RegisterSubcommands() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	poemsserver.RegisterSubcommand()
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
