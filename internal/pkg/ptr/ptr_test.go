package ptr_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/stretchr/testify/assert"
)

func TestDereference(t *testing.T) {
	assert.Equal(t, ptr.Dereference((*string)(nil), "default"), "default")
	assert.Equal(
		t,
		ptr.Dereference(ptr.New("first"), "default"),
		"first",
	)
}

func TestNew(t *testing.T) {
	assert.Equal(t, *ptr.New(uint(5)), uint(5))
}
