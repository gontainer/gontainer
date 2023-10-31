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

package compiler_test

import (
	"fmt"
	"math"
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/compiler"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/resolver"
	"github.com/stretchr/testify/assert"
)

func TestStepCompileParams_Process(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name           string
		Resolver       paramResolveFunc
		Params         map[string]any
		Errors         []string
		CompiledParams []output.Param
	}{
		{
			Name: "Errors",
			Resolver: func(a any) (p resolver.ParamExpr, err error) {
				err = fmt.Errorf("could not resolve param %T", a)
				return
			},
			Params: map[string]any{
				"intValue":   5,
				"floatValue": math.Pi,
				"slice":      []any{},
			},
			Errors: []string{
				`compiler.StepCompileParams: "floatValue": could not resolve param float64`,
				`compiler.StepCompileParams: "intValue": could not resolve param int`,
				`compiler.StepCompileParams: "slice": could not resolve param []interface {}`,
			},
			CompiledParams: []output.Param{
				{Name: "floatValue"},
				{Name: "intValue"},
				{Name: "slice"},
			},
		},
		{
			Name:           "Empty",
			Resolver:       nil,
			Params:         nil,
			Errors:         nil,
			CompiledParams: nil,
		},
		{
			Name: "OK",
			Resolver: func(a any) (resolver.ParamExpr, error) {
				return resolver.ParamExpr{
					Code:            "my-code",
					Raw:             "my-raw-value",
					DependsOnParams: []string{"paramA", "paramB"},
				}, nil
			},
			Params: map[string]any{
				"someParam":    "someValue",
				"anotherParam": "anotherValue",
			},
			CompiledParams: []output.Param{
				{
					Name:      "anotherParam",
					Code:      "my-code",
					Raw:       "my-raw-value",
					DependsOn: []string{"paramA", "paramB"},
				},
				{
					Name:      "someParam",
					Code:      "my-code",
					Raw:       "my-raw-value",
					DependsOn: []string{"paramA", "paramB"},
				},
			},
		},
	}
	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			i := input.Input{Params: s.Params}
			o := output.Output{}
			err := compiler.NewStepCompileParams(s.Resolver).Process(i, &o)
			errAssert.EqualErrorGroup(t, err, s.Errors)
			assert.Equal(t, s.CompiledParams, o.Params)
		})
	}
}
