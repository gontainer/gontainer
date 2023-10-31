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

package runner

import (
	"fmt"
	"strings"

	"github.com/gontainer/gontainer-helpers/v2/container"
	"github.com/gontainer/gontainer-helpers/v2/grouperror"
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
		l := len(grouperror.Collection(err))
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
