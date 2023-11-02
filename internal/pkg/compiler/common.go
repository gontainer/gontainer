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

package compiler

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/v3/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

func argExprToArg(e resolver.ArgExpr) output.Arg {
	return output.Arg{
		Code:              e.Code,
		Raw:               e.Raw,
		DependsOnParams:   e.DependsOnParams,
		DependsOnServices: e.DependsOnServices,
		DependsOnTags:     e.DependsOnTags,
	}
}

func resolveArgs(resolver argResolver, args []any) (r []output.Arg, _ error) {
	var errs []error
	if len(args) > 0 {
		r = make([]output.Arg, len(args))
	}
	for i, arg := range args {
		argExpr, err := resolver.ResolveArg(arg)
		errs = append(errs, grouperror.Prefix(fmt.Sprintf("%d: ", i), err))
		r[i] = argExprToArg(argExpr)
	}
	return r, grouperror.Prefix("args: ", errs...)
}
