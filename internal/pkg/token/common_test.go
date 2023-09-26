package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toExpr(t *testing.T) {
	expr, ok := toExpr("%name")
	assert.Empty(t, expr)
	assert.False(t, ok)

	expr, ok = toExpr("%")
	assert.Empty(t, expr)
	assert.False(t, ok)

	expr, ok = toExpr("%host%")
	assert.Equal(t, "host", expr)
	assert.True(t, ok)
}
