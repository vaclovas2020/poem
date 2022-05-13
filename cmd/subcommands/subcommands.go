/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

/* Package for CLI subcommands registration */
package subcommands

import (
	"context"
	"flag"

	"github.com/google/subcommands"
	"webimizer.dev/poem/cmd/subcommands/adminfrontend"
	"webimizer.dev/poem/cmd/subcommands/adminserver"
	"webimizer.dev/poem/cmd/subcommands/install"
	"webimizer.dev/poem/cmd/subcommands/poemsserver"
)

/* Register all application's subcommands and execute choosen subcommand */
func RegisterSubcommands() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	poemsserver.RegisterSubcommand()
	adminserver.RegisterSubcommand()
	install.RegisterSubcommand()
	adminfrontend.RegisterSubcommand()
	flag.Parse()
	ctx := context.Background()
	subcommands.Execute(ctx)
}
