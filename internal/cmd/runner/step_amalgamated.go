package runner

import (
	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type StepAmalgamated struct {
	name  string
	steps []Step
}

func NewStepAmalgamated(name string, steps ...Step) *StepAmalgamated {
	return &StepAmalgamated{
		name:  name,
		steps: steps,
	}
}

func (s *StepAmalgamated) Name() string {
	return s.name
}

func (s *StepAmalgamated) Run(i *input.Input, o *output.Output) error {
	errs := make([]error, 0, len(s.steps))
	for _, st := range s.steps {
		errs = append(errs, st.Run(i, o))
	}
	return grouperror.Join(errs...)
}
