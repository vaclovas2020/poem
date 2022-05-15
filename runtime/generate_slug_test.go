package runtime_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"webimizer.dev/poem/runtime"
)

func TestGenerateSlug(t *testing.T) {
	expected := "test-and-how-it-works"
	result := runtime.GenerateSlug("TeSt And How it +Žworks")
	assert.Equal(t, expected, result)
}
