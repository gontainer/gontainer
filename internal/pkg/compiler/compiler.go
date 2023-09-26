package compiler

import (
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type Step interface {
	Process(input.Input, *output.Output) error
}

type Compiler struct {
	steps []Step
}

func New(steps ...Step) *Compiler {
	return &Compiler{steps: steps}
}

func (c Compiler) Compile(i input.Input) (o output.Output, _ error) {
	for _, s := range c.steps {
		if err := s.Process(i, &o); err != nil {
			return o, err
		}
	}
	return o, nil
}
