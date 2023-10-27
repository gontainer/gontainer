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
