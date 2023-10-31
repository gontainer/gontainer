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
