// Copyright (c) 2023 Bartłomiej Krukowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is furnished
// to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
