package install

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	hash, err := hashPassword("Test")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}
