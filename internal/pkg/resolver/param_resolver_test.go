package resolver_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
	"github.com/gontainer/gontainer/internal/pkg/token"
	"github.com/stretchr/testify/assert"
)

type argResolver interface {
	ResolveArg(any) (resolver.ArgExpr, error)
}

type argResolverMock func(any) (resolver.ArgExpr, error)

func (a argResolverMock) ResolveArg(i any) (resolver.ArgExpr, error) {
	return a(i)
}

func TestParamResolver_ResolveParam(t *testing.T) {
	t.Parallel()

	scenarios := []struct {
		Name        string
		ArgResolver argResolver
		Input       any
		Param       resolver.ParamExpr
		Error       string
	}{
		{
			Name:        "Integer",
			ArgResolver: resolver.NewArgResolver(resolver.NewNonStringPrimitiveResolver()),
			Input:       5,
			Param: resolver.ParamExpr{
				Code:            `dependencyValue(int(5))`,
				Raw:             5,
				DependsOnParams: nil,
			},
			Error: "",
		},
		{
			Name: "DependsOnParams",
			ArgResolver: resolver.NewArgResolver(
				resolver.NewPatternResolver(
					token.NewTokenizer(
						token.NewChunker(),
						token.NewStrategyFactory(token.FactoryReference{}),
					),
				),
			),
			Input: `%name%`,
			Param: resolver.ParamExpr{
				Code:            `dependencyProvider(func() (interface{}, error) { return getParam("name") })`,
				Raw:             `%name%`,
				DependsOnParams: []string{"name"},
			},
			Error: "",
		},
		{
			Name:        "Error #1 service",
			ArgResolver: resolver.NewArgResolver(resolver.NewServiceResolver()),
			Input:       "@db",
			Param:       resolver.ParamExpr{},
			Error:       "param cannot depend on any service: db",
		},
		{
			Name: "Error #2",
			ArgResolver: argResolverMock(func(any) (resolver.ArgExpr, error) {
				return resolver.ArgExpr{
					Code:              "some code",
					Raw:               nil,
					DependsOnParams:   nil,
					DependsOnServices: []string{"service1", "service2"},
					DependsOnTags:     []string{"param1"},
				}, nil
			}),
			Input: "some input",
			Param: resolver.ParamExpr{},
			Error: `param cannot depend on any service: service1, service2` + "\n" +
				`param cannot depend on any tag: param1`,
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			param, err := resolver.NewParamResolver(s.ArgResolver).
				ResolveParam(s.Input)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, s.Param, param)
		})
	}
}
