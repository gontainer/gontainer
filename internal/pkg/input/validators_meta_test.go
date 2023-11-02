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

	errAssert "github.com/gontainer/gontainer-helpers/v3/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/stretchr/testify/assert"
)

func TestValidateMeta(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name   string
		Meta   input.Meta
		Errors []string
	}{
		{
			Name: "Errors #1",
			Meta: input.Meta{
				Pkg:                  ptr.New("some pkg"),
				ContainerType:        ptr.New("$"),
				ContainerConstructor: ptr.New("New Container"),
				Imports: map[string]string{
					`viper`:  `github.com/spf13/viper`,
					`viper2`: `"github.com/spf13/viper"`,
					`_viper`: `github.com/spf13 viper`,
				},
				Functions: map[string]string{
					"correctFn":    `"pkg".Env`,
					"incorrect Fn": `pkg IncorrectEnv`,
				},
			},
			Errors: []string{
				`meta: pkg: invalid "some pkg"`,
				`meta: container_type: invalid "$"`,
				`meta: container_constructor: invalid "New Container"`,
				`meta: imports: invalid import "github.com/spf13 viper"`,
				`meta: imports: invalid alias "_viper"`,
				`meta: functions: invalid function "incorrect Fn"`,
				`meta: functions: invalid go function "pkg IncorrectEnv"`,
			},
		},
		{
			Name: "OK #1",
			Meta: input.Meta{
				Pkg:                  ptr.New("gontainer"),
				ContainerConstructor: ptr.New("New"),
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			i := input.Input{Meta: s.Meta}

			t.Run("input.ValidateMeta", func(t *testing.T) {
				t.Parallel()
				err := input.ValidateMeta(i)
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

func TestDefaultMetaValidators(t *testing.T) {
	assert.NotEmpty(t, input.DefaultMetaValidators())
}
