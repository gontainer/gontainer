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
