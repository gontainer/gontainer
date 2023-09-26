package compiler

import (
	"github.com/gontainer/gontainer-helpers/errors"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type inputValidator interface {
	Validate(input input.Input) error
}

type StepValidateInput struct {
	validator inputValidator
}

func NewStepValidateInput(v inputValidator) *StepValidateInput {
	return &StepValidateInput{validator: v}
}

func (s StepValidateInput) Process(i input.Input, _ *output.Output) error {
	return errors.PrefixedGroup("compiler.StepValidateInput: ", s.validator.Validate(i))
}
