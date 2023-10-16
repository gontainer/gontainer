package compiler_test

import (
	"fmt"
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/grouperror/assert"
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
