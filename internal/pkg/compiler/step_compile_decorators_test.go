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
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/compiler"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/resolver"
	"github.com/stretchr/testify/assert"
)

func TestStepCompileDecorators_Process(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name     string
		Input    []input.Decorator
		Output   []output.Decorator
		Compiler *compiler.StepCompileDecorators
		Errors   []string
	}{
		{
			Name:     "OK",
			Compiler: compiler.NewStepCompileDecorators(simpleAliaserFunc, resolver.NewArgResolver(resolver.NewServiceResolver())),
			Input: []input.Decorator{
				{
					Tag:       "endpoint",
					Decorator: "http.TraceableMiddleware",
					Args:      []any{"@logger"},
				},
			},
			Output: []output.Decorator{
				{
					Tag:       "endpoint",
					Decorator: "alias_http.TraceableMiddleware",
					Raw:       "http.TraceableMiddleware",
					Args: []output.Arg{{
						Code:              `dependencyService("logger")`,
						Raw:               "@logger",
						DependsOnParams:   nil,
						DependsOnServices: []string{"logger"},
						DependsOnTags:     nil,
					}},
				},
			},
		},
		{
			Name: "Error",
			Compiler: compiler.NewStepCompileDecorators(simpleAliaserFunc, argResolverFunc(func(any) (resolver.ArgExpr, error) {
				return resolver.ArgExpr{}, fmt.Errorf("could not resolve param")
			})),
			Input: []input.Decorator{
				{
					Tag:       "endpoint",
					Decorator: "http.TraceableMiddleware",
					Args:      []any{"@logger"},
				},
			},
			Output: []output.Decorator{
				{
					Tag:       "endpoint",
					Decorator: "alias_http.TraceableMiddleware",
					Raw:       "http.TraceableMiddleware",
					Args:      []output.Arg{{}},
				},
			},
			Errors: []string{
				`compiler.StepCompileDecorators: #0 "http.TraceableMiddleware": args: 0: could not resolve param`,
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			o := output.Output{}
			err := s.Compiler.Process(input.Input{Decorators: s.Input}, &o)
			errAssert.EqualErrorGroup(t, err, s.Errors)
			assert.Equal(t, s.Output, o.Decorators)
		})
	}
}
