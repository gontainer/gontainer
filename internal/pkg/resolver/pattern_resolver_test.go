package resolver_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
	"github.com/gontainer/gontainer/internal/pkg/token"
)

type mockTokenizer func(string) (token.Tokens, error)

func (m mockTokenizer) Tokenize(pattern string) (token.Tokens, error) {
	return m(pattern)
}

func patternResolverScenarios() []anyResolverStrategyScenario {
	f := token.NewStrategyFactory(
		token.FactoryPercentMark{},
		token.FactoryReference{},
		token.FactoryUnexpectedFunction{},
		token.FactoryUnexpectedToken{},
		token.FactoryString{},
	)
	r := resolver.NewPatternResolver(token.NewTokenizer(token.NewChunker(), f))
	return []anyResolverStrategyScenario{
		{
			Name:       "not supported",
			Expression: nil,
			Supports:   false,
			Resolver:   r,
		},
		{
			Name:       "error #1",
			Expression: "%name",
			Supports:   true,
			Error:      `not closed token: "%name"`,
			Resolver:   r,
		},
		{
			Name:       "error #2",
			Expression: "my expression",
			Supports:   true,
			Error:      "unexpected error: len(tokens) == 0",
			Resolver: resolver.NewPatternResolver(mockTokenizer(func(string) (token.Tokens, error) {
				return nil, nil
			})),
		},
		{
			Name:       "OK #1",
			Expression: "%name%",
			Supports:   true,
			Arg: resolver.ArgExpr{
				Code:            `dependencyProvider(func() (interface{}, error) { return getParam("name") })`,
				Raw:             "%name%",
				DependsOnParams: []string{"name"},
			},
			Resolver: r,
		},
	}
}

func TestPatternResolver_ResolveArg(t *testing.T) {
	t.Parallel()
	assertScenarios(t, patternResolverScenarios()...)
}
