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
