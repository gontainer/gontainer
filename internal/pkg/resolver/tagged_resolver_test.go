package resolver_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

func taggedResolverScenarios() []anyResolverStrategyScenario {
	r := resolver.NewTaggedResolver()

	return []anyResolverStrategyScenario{
		{
			Name:       "OK #1",
			Expression: "!tagged\tdoer",
			Supports:   true,
			Arg: resolver.ArgExpr{
				Code:          `dependencyTag("doer")`,
				Raw:           "!tagged\tdoer",
				DependsOnTags: []string{"doer"},
			},
			Resolver: r,
		},
		{
			Name:       "OK #2",
			Expression: `!tagged  doer`,
			Supports:   true,
			Arg: resolver.ArgExpr{
				Code:          `dependencyTag("doer")`,
				Raw:           `!tagged  doer`,
				DependsOnTags: []string{"doer"},
			},
			Resolver: r,
		},
		{
			Name:       "not supported #1",
			Expression: " !tagged doer",
			Supports:   false,
			Resolver:   r,
		},
		{
			Name:       "not supported #2",
			Expression: nil,
			Supports:   false,
			Resolver:   r,
		},
		{
			Name:       "invalid tag #1",
			Expression: `!tagged do..er`,
			Supports:   true,
			Error:      "invalid tag",
			Resolver:   r,
		},
		{
			Name:       "invalid tag #2",
			Expression: "!tagged @!@#$%^&*()",
			Supports:   true,
			Error:      "invalid tag",
			Resolver:   r,
		},
	}
}

func TestTaggedResolver_ResolveArg(t *testing.T) {
	t.Parallel()
	assertScenarios(t, taggedResolverScenarios()...)
}
