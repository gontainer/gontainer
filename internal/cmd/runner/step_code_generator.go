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
	"os"
	"path/filepath"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type codeBuilder interface {
	Build(output.Output) (string, error)
}

type StepCodeGenerator struct {
	printer    printer
	builder    codeBuilder
	outputFile string
}

func NewStepCodeGenerator(printer printer, builder codeBuilder, outputFile string) *StepCodeGenerator {
	return &StepCodeGenerator{
		printer:    printer,
		builder:    builder,
		outputFile: outputFile,
	}
}

func (s *StepCodeGenerator) Name() string {
	return "Generate code"
}

func (s *StepCodeGenerator) Run(_ *input.Input, o *output.Output) error {
	s.printer.Println("Generating source code")
	tpl, err := s.builder.Build(*o)
	if err != nil {
		return err
	}
	s.printer.Println(fmt.Sprintf("Printing to the file `%s`", s.outputFile))
	of := filepath.Clean(s.outputFile)
	if err := os.WriteFile(of, []byte(tpl), 0644); err != nil {
		return err
	}
	return nil
}
