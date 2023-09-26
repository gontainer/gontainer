package runner

import (
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type compiler interface {
	Compile(input.Input) (output.Output, error)
}

type StepCompile struct {
	compiler compiler
}

func NewStepCompile(c compiler) *StepCompile {
	return &StepCompile{compiler: c}
}

func (s *StepCompile) Name() string {
	return "Compile"
}

func (s *StepCompile) Run(i *input.Input, o *output.Output) error {
	c, err := s.compiler.Compile(*i)
	*o = c
	return err
}
