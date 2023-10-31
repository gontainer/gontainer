package slices_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/slices"
	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	t.Run("append vs copy", func(t *testing.T) {
		slice := make([]int, 5)
		clone := append(slice) //nolint:all
		cp := slices.Copy(slice)

		assert.NotSame(t, slice, clone)
		assert.Equal(t, slice, clone)
		assert.Equal(t, slice, cp)

		slice[0] = 5                  // changing a value in `slice`
		assert.Equal(t, slice, clone) // changes the corresponding value in `clone`
		assert.NotEqual(t, slice, cp) // but not in `cp`
	})
	t.Run("nil", func(t *testing.T) {
		var nil_ []int
		notNil := make([]int, 0)

		assert.Nil(t, nil_)
		assert.Nil(t, slices.Copy(nil_))

		assert.NotNil(t, notNil)
		assert.NotNil(t, slices.Copy(notNil))
	})
}
