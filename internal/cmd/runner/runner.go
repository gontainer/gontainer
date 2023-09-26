package runner

import (
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type Step interface {
	Run(*input.Input, *output.Output) error
}

type Runner struct {
	steps []Step
}

func NewRunner(steps ...Step) *Runner {
	return &Runner{
		steps: steps,
	}
}

func (r *Runner) Run() error {
	i := input.Input{}
	o := output.Output{}
	for _, sr := range r.steps {
		if err := sr.Run(&i, &o); err != nil {
			return err
		}
	}
	return nil
}
