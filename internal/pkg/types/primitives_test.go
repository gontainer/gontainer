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

package types_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/types"
	"github.com/stretchr/testify/assert"
)

type MyInt int

func TestIsPrimitive(t *testing.T) {
	scenarios := []struct {
		Name     string
		Value    any
		Expected bool
	}{
		{
			Name:     "MyInt(5)",
			Value:    MyInt(5),
			Expected: true,
		},
		{
			Name:     "nil",
			Value:    nil,
			Expected: true,
		},
		{
			Name:     "(*error)(nil)",
			Value:    (*error)(nil),
			Expected: false,
		},
		{
			Name:     "(error)(nil)",
			Value:    (error)(nil),
			Expected: true,
		},
		{
			Name:     "(*int)(nil)",
			Value:    (*int)(nil),
			Expected: false,
		},
		{
			Name:     "(*interface{})(nil)",
			Value:    (*interface{})(nil),
			Expected: false,
		},
		{
			Name:     "(interface{})(nil)",
			Value:    (interface{})(nil),
			Expected: true,
		},
		{
			Name:     "(interface{ Foo() })(nil)",
			Value:    (interface{ Foo() })(nil),
			Expected: true,
		},
		{
			Name:     "struct{}{}",
			Value:    struct{}{},
			Expected: false,
		},
		{
			Name:     "int",
			Value:    5,
			Expected: true,
		},
		{
			Name:     "empty slice",
			Value:    ([]int)(nil),
			Expected: false,
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, s.Expected, types.IsPrimitive(s.Value))
		})
	}
}
