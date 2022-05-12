/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package poem_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"webimizer.dev/poem"
)

/* Testing Poem CLI application output */
func TestInitApplication(t *testing.T) {
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stderr = w
	subcommands.DefaultCommander.Error = w
	poem.InitApplication()
	_ = w.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r)
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "Usage:")
	assert.Contains(t, buf.String(), "Subcommands:")
	assert.Contains(t, buf.String(), "commands")
	assert.Contains(t, buf.String(), "flags")
	assert.Contains(t, buf.String(), "help")
	assert.Contains(t, buf.String(), "install")
	assert.Contains(t, buf.String(), "poems-server")
	assert.Contains(t, buf.String(), "admin-server")
}

/* Testing Poems server output */
func TestInitPoemsServer(t *testing.T) {
	os.Args = []string{"poem", "poems-server"}
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	subcommands.DefaultCommander.Output = w
	poem.InitApplication()
	_ = w.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r)
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "Starting server...")
}

/* Testing Admin server output */
func TestInitAdminServer(t *testing.T) {
	os.Args = []string{"poem", "admin-server"}
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	subcommands.DefaultCommander.Output = w
	poem.InitApplication()
	_ = w.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r)
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "Starting server...")
}

/* Testing install subcommand output */
func TestInstallSubcommand(t *testing.T) {
	os.Args = []string{"poem", "install"}
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	subcommands.DefaultCommander.Output = w
	poem.InitApplication()
	_ = w.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r)
	require.NoError(t, err)
	assert.Contains(t, buf.String(), "Installing CMS database...")
}
