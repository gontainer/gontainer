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
