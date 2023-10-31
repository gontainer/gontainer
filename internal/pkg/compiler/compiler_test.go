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

package compiler_test

import (
	"fmt"
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/compiler"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/stretchr/testify/assert"
)

type stepFunc func(input.Input, *output.Output) error

func (s stepFunc) Process(i input.Input, o *output.Output) error {
	return s(i, o)
}

func TestNewCompiler(t *testing.T) {
	t.Parallel()

	emptyStep := stepFunc(func(input.Input, *output.Output) error {
		return nil
	})

	scenarios := []struct {
		Name      string
		StepFuncs []stepFunc
		Input     input.Input
		Output    output.Output
		Error     string
	}{
		{
			Name: "Error #1",
			StepFuncs: []stepFunc{func(input.Input, *output.Output) error {
				return fmt.Errorf("my error")
			}},
			Error: "my error",
		},
		{
			Name: "Error #2",
			StepFuncs: []stepFunc{
				emptyStep,
				emptyStep,
				emptyStep,
				func(input.Input, *output.Output) error {
					return fmt.Errorf("another error")
				},
			},
			Error: "another error",
		},
		{
			Name: "OK",
			StepFuncs: []stepFunc{func(i input.Input, o *output.Output) error {
				o.Meta.Pkg = "myPkg"
				return nil
			}},
			Output: output.Output{
				Meta: output.Meta{
					Pkg: "myPkg",
				},
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()

			steps := make([]compiler.Step, len(s.StepFuncs))
			for i, sf := range s.StepFuncs {
				steps[i] = sf
			}

			c := compiler.New(steps...)
			o, err := c.Compile(s.Input)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error, s.Error)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, s.Output, o)
		})
	}
}
