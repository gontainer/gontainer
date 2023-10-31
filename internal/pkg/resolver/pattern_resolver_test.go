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
