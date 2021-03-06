/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

/* Poem CMS main package */
package poem

import "webimizer.dev/poem/cmd/subcommands"

/*

This package implements Poem CMS based on gRPC services.
You need to start each service indvidually using Docker or using Poem CLI subcommands

*/

/* Initialize Poem CLI application */
func InitApplication() {
	subcommands.RegisterSubcommands()
}
