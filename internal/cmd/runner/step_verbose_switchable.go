package runner

import (
	"fmt"
	"strings"

	"github.com/gontainer/gontainer-helpers/container"
	"github.com/gontainer/gontainer-helpers/errors"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type StepVerboseSwitchable struct {
	parent   Step
	printer  printer
	indenter indenter
	active   bool
}

func DecorateStepVerboseSwitchable(payload container.DecoratorPayload, p printer, i indenter) *StepVerboseSwitchable {
	return NewStepVerboseSwitchable(payload.Service.(Step), p, i)
}

func NewStepVerboseSwitchable(parent Step, printer printer, i indenter) *StepVerboseSwitchable {
	return &StepVerboseSwitchable{
		parent:   parent,
		printer:  printer,
		indenter: i,
		active:   true,
	}
}

func (s *StepVerboseSwitchable) Active(active bool) {
	s.active = active
}

func (s *StepVerboseSwitchable) Run(i *input.Input, o *output.Output) error {
	n := s.name()
	s.printer.PrintAlignedLn(n)
	if !s.active {
		s.printer.PrintAlignedLn(n+" END", "ignored")
		return nil
	}

	var err error
	func() {
		s.indenter.Indent("  ")
		defer s.indenter.EndIndent()

		err = s.parent.Run(i, o)
	}()
	if err == nil {
		s.printer.PrintAlignedLn(n+" END", checkMark)
	} else {
		l := len(errors.Collection(err))
		extra := fmt.Sprintf(" (%d error)", l)
		if l > 1 {
			extra = fmt.Sprintf(" (%d errors)", l)
		}
		s.printer.PrintAlignedLn(n+" END", xMark, extra)
	}
	return err
}

func (s *StepVerboseSwitchable) name() string {
	if v, ok := s.parent.(interface{ Name() string }); ok {
		return v.Name()
	}

	n := strings.Trim(fmt.Sprintf("%T", s.parent), "*")
	n = strings.TrimPrefix(n, "runner.")
	return n
}
