package compiler

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/grouperror"
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
