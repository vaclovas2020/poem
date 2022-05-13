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

type EmptyTemplate struct {
}

func TestTemplateParse(t *testing.T) {
	obj := &HelloTemplate{Name: "World"}
	output, err := runtime.TemplateParse(content, "generic_template_test.tmpl", obj)
	require.NoError(t, err)
	assert.Contains(t, output, "Hello World!")
	_, err = runtime.TemplateParse(content, "not_exists.tmpl", obj)
	assert.Error(t, err)
	emptyObj := new(EmptyTemplate)
	_, err = runtime.TemplateParse(content, "generic_template_test.tmpl", emptyObj)
	assert.Error(t, err)
}
