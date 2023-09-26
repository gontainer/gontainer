package resolver_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

func fixedValueResolverScenarios() []anyResolverStrategyScenario {
	r := resolver.NewFixedValueResolver("$gontainer", "rootGontainer")

	return []anyResolverStrategyScenario{
		{
			Name:       "OK",
			Expression: "$gontainer",
			Supports:   true,
			Arg: resolver.ArgExpr{
				Code: "dependencyValue(rootGontainer)",
				Raw:  "$gontainer",
			},
			Resolver: r,
		},
		{
			Name:       "Not supported",
			Expression: "some expression",
			Supports:   false,
			Resolver:   r,
		},
	}
}

func TestFixedValueResolver_ResolveArg(t *testing.T) {
	t.Parallel()
	assertScenarios(t, fixedValueResolverScenarios()...)
}
