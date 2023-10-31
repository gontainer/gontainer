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

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/compiler"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

func TestStepValidateInput_Process(t *testing.T) {
	t.Parallel()

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		step := compiler.NewStepValidateInput(inputValidatorFunc(func(input.Input) error {
			return grouperror.Join(fmt.Errorf("my error #1"), fmt.Errorf("my error #2"))
		}))

		errAssert.EqualErrorGroup(
			t,
			step.Process(input.Input{}, &output.Output{}),
			[]string{
				"compiler.StepValidateInput: my error #1",
				"compiler.StepValidateInput: my error #2",
			},
		)
	})
}
