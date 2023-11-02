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
	"strings"

	"github.com/gontainer/gontainer-helpers/v3/grouperror"
)

type argResolver interface {
	ResolveArg(any) (ArgExpr, error)
}

// ParamResolver resolves input parameters. In opposition to arguments,
// parameters may depend o on other parameters only.
type ParamResolver struct {
	resolver argResolver
}

func NewParamResolver(resolver argResolver) *ParamResolver {
	return &ParamResolver{resolver: resolver}
}

func (p ParamResolver) ResolveParam(i any) (ParamExpr, error) {
	a, err := p.resolver.ResolveArg(i)
	// params cannot depend on services or tags
	// if they do, most likely a wrong ArgResolver has been injected
	// to the constructor NewParamResolver
	var errs []error
	if len(a.DependsOnServices) > 0 {
		errs = append(
			errs,
			fmt.Errorf("param cannot depend on any service: %s", strings.Join(a.DependsOnServices, ", ")),
		)
	}
	if len(a.DependsOnTags) > 0 {
		errs = append(
			errs,
			fmt.Errorf("param cannot depend on any tag: %s", strings.Join(a.DependsOnTags, ", ")),
		)
	}
	if len(errs) > 0 {
		return ParamExpr{}, grouperror.Join(errs...)
	}
	return ParamExpr{
		Code:            a.Code,
		Raw:             a.Raw,
		DependsOnParams: a.DependsOnParams,
	}, err
}
