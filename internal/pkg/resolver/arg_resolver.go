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

package resolver

import (
	"fmt"
)

type ArgExpr struct {
	// Code is Go code in on of the following formats:
	// 	* dependencyProvider(%s)
	// 	* dependencyTag(%s)
	// 	* dependencyValue(%s)
	// 	* dependencyService(%s)
	Code              string
	Raw               any
	DependsOnParams   []string
	DependsOnServices []string
	DependsOnTags     []string
}

type ParamExpr struct {
	// Code is Go code in on of the following formats:
	// 	* dependencyProvider(%s)
	// 	* dependencyTag(%s)
	// 	* dependencyValue(%s)
	// 	* dependencyService(%s)
	Code            string
	Raw             any
	DependsOnParams []string
}

type resolverStrategy interface {
	ResolveArg(any) (ArgExpr, error)
	Supports(any) bool
}

type ArgResolver struct {
	strategies []resolverStrategy
}

func NewArgResolver(s ...resolverStrategy) *ArgResolver {
	return &ArgResolver{strategies: s}
}

func (a *ArgResolver) ResolveArg(i any) (e ArgExpr, _ error) {
	for _, s := range a.strategies {
		if s.Supports(i) {
			return s.ResolveArg(i)
		}
	}
	return e, fmt.Errorf("not supported %T", i)
}
