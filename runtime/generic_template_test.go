package runtime_test

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"webimizer.dev/poem/runtime"
)

//go:embed generic_template_test.tmpl
var content embed.FS

type HelloTemplate struct {
	Name string
}

func TestTemplateParse(t *testing.T) {
	obj := &HelloTemplate{Name: "World"}
	output, err := runtime.TemplateParse(content, "generic_template_test.tmpl", obj)
	require.NoError(t, err)
	assert.Contains(t, output, "Hello World!")
}
