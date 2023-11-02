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
	"math"
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v3/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
)

func TestDefaultParamsValidators(t *testing.T) {
	assert.NotEmpty(t, input.DefaultParamsValidators())
}

func TestValidateParams(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name   string
		Params map[string]any
		Errors []string
	}{
		{
			Name: "OK",
			Params: map[string]any{
				"math.Pi": math.Pi,
				"nil":     nil,
			},
		},
		{
			Name: "Errors",
			Params: map[string]any{
				"math..Pi":   math.Pi,
				"slice":      []int{1, 2, 3},
				"emptySLice": []int(nil),
			},
			Errors: []string{
				`parameters: "emptySLice": unsupported type []int`,
				`parameters: "math..Pi": invalid name`,
				`parameters: "slice": unsupported type []int`,
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			i := input.Input{Params: s.Params}

			t.Run("input.ValidateParams", func(t *testing.T) {
				t.Parallel()
				err := input.ValidateParams(i)
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
