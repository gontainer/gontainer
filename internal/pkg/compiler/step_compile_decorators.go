package compiler

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/regex"
	"github.com/gontainer/gontainer/internal/pkg/syntax"
)

var (
	regexDecoratorMethod = regex.MustCompileAz(regex.DecoratorMethod)
)

type StepCompileDecorators struct {
	aliaser     aliaser
	argResolver argResolver
}

func NewStepCompileDecorators(a aliaser, ar argResolver) *StepCompileDecorators {
	return &StepCompileDecorators{
		aliaser:     a,
		argResolver: ar,
	}
}

func (s StepCompileDecorators) Process(i input.Input, d *output.Output) error {
	errs := make([]error, len(i.Decorators))
	d.Decorators = make([]output.Decorator, len(i.Decorators))
	for j, curr := range i.Decorators {
		dec, err := s.processDecorator(curr)
		errs[j] = grouperror.Prefix(fmt.Sprintf("#%d %+q: ", j, curr.Decorator), err)
		d.Decorators[j] = dec
	}
	return grouperror.Prefix("compiler.StepCompileDecorators: ", errs...)
}

func (s StepCompileDecorators) processDecorator(d input.Decorator) (output.Decorator, error) {
	_, m := regex.Match(regexDecoratorMethod, d.Decorator)

	method := m["fn"]
	import_ := syntax.SanitizeImport(m["import"])
	if import_ != "" {
		method = s.aliaser.Alias(import_) + "." + method
	}

	args, err := resolveArgs(s.argResolver, d.Args)

	return output.Decorator{
		Tag:       d.Tag,
		Decorator: method,
		Args:      args,
		Raw:       d.Decorator,
	}, err
}
