package resolver_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

func serviceResolverScenarios() []anyResolverStrategyScenario {
	r := resolver.NewServiceResolver()

	return []anyResolverStrategyScenario{
		{
			Name:       "OK",
			Expression: `@logger.debug`,
			Supports:   true,
			Arg: resolver.ArgExpr{
				Code:              `dependencyService("logger.debug")`,
				Raw:               `@logger.debug`,
				DependsOnServices: []string{`logger.debug`},
			},
			Resolver: r,
		},
		{
			Name:       "not supported #1",
			Expression: " @logger",
			Supports:   false,
			Resolver:   r,
		},
		{
			Name:       "not supported #2",
			Expression: 5,
			Supports:   false,
			Resolver:   r,
		},
		{
			Name:       "invalid service #1",
			Expression: "@logger..debug",
			Supports:   true,
			Error:      "invalid service",
			Resolver:   r,
		},
		{
			Name:       "invalid service #2",
			Expression: "@!@#$%^&*()",
			Supports:   true,
			Error:      "invalid service",
			Resolver:   r,
		},
	}
}

func TestServiceResolver_ResolveArg(t *testing.T) {
	t.Parallel()
	assertScenarios(t, serviceResolverScenarios()...)
}
