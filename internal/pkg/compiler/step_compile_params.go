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

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/maps"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type StepCompileParams struct {
	resolver paramResolver
}

func NewStepCompileParams(r paramResolver) *StepCompileParams {
	return &StepCompileParams{resolver: r}
}

func (s StepCompileParams) Process(i input.Input, d *output.Output) error {
	var errs []error
	if d.Params == nil && len(i.Params) > 0 {
		d.Params = make([]output.Param, 0, len(i.Params))
	}
	maps.Iterate(i.Params, func(k string, v any) {
		expr, err := s.resolver.ResolveParam(v)
		if err != nil {
			errs = append(errs, grouperror.Prefix(fmt.Sprintf("%+q: ", k), err))
			d.Params = append(d.Params, output.Param{
				Name:      k,
				Code:      "",
				Raw:       nil,
				DependsOn: nil,
			})
			return
		}

		d.Params = append(d.Params, output.Param{
			Name:      k,
			Code:      expr.Code,
			Raw:       expr.Raw,
			DependsOn: expr.DependsOnParams,
		})
	})
	return grouperror.Prefix("compiler.StepCompileParams: ", errs...)
}
