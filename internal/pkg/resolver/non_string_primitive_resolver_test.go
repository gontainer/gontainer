package resolver_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

func nonStringPrimitiveResolverScenarios() []anyResolverStrategyScenario {
	r := resolver.NewNonStringPrimitiveResolver()

	return []anyResolverStrategyScenario{
		{
			Name:       "OK #1",
			Expression: 5,
			Supports:   true,
			Arg: resolver.ArgExpr{
				Code: `dependencyValue(int(5))`,
				Raw:  5,
			},
			Resolver: r,
		},
		{
			Name:       "OK #2",
			Expression: nil,
			Supports:   true,
			Arg: resolver.ArgExpr{
				Code: `dependencyValue(nil)`,
				Raw:  nil,
			},
			Resolver: r,
		},
		{
			Name:       "not supported #1",
			Expression: struct{}{},
			Supports:   false,
			Resolver:   r,
		},
		{
			Name:       "not supported #2",
			Expression: "hello",
			Supports:   false,
			Resolver:   r,
		},
	}
}

func TestNonStringPrimitiveResolver_ResolveArg(t *testing.T) {
	t.Parallel()
	assertScenarios(t, nonStringPrimitiveResolverScenarios()...)
}
