package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokens_GoCode(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		c, err := Tokens{}.GoCode()
		assert.EqualError(t, err, "unexpected error: len(tokens) == 0")
		assert.Empty(t, c)
	})
}
