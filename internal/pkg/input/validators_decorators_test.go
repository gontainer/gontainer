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

package input_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
)

func TestDefaultDecoratorsValidators(t *testing.T) {
	assert.NotEmpty(t, input.DefaultDecoratorsValidators())
}

func TestValidateDecorators(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name      string
		Decorator []input.Decorator
		Errors    []string
	}{
		{
			Name: "Errors #1",
			Decorator: []input.Decorator{
				{
					Tag:       "my tag",
					Decorator: "my func",
					Args:      []any{func() {}, nil, (*int)(nil), 17},
				},
			},
			Errors: []string{
				`decorators: 0 "my func": tag: invalid "my tag"`,
				`decorators: 0 "my func": method: invalid "my func"`,
				`decorators: 0 "my func": arguments: 0: unsupported type func()`,
				`decorators: 0 "my func": arguments: 2: unsupported type *int`,
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			i := input.Input{Decorators: s.Decorator}

			t.Run("input.ValidateDecorators", func(t *testing.T) {
				t.Parallel()
				err := input.ValidateDecorators(i)
				errAssert.EqualErrorGroup(t, err, s.Errors)
			})
			t.Run("input.NewDefaultValidator", func(t *testing.T) {
				t.Parallel()
				err := input.NewDefaultValidator("dev-main").Validate(i)
				errAssert.EqualErrorGroup(t, err, s.Errors)
			})
		})
	}
}
