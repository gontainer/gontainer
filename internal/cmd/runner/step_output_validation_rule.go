package runner

import (
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type StepOutputValidationRule struct {
	validator func(output.Output) error
	ruleName  string
}

func NewStepOutputValidationRule(v func(output.Output) error, ruleName string) *StepOutputValidationRule {
	return &StepOutputValidationRule{
		validator: v,
		ruleName:  ruleName,
	}
}

func (s *StepOutputValidationRule) Name() string {
	return s.ruleName
}

func (s *StepOutputValidationRule) Run(_ *input.Input, o *output.Output) error {
	return s.validator(*o)
}
