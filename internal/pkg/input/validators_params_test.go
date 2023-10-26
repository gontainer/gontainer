package input_test

import (
	"math"
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
)

func TestDefaultParamsValidators(t *testing.T) {
	assert.NotEmpty(t, input.DefaultParamsValidators())
}

func TestValidateParams(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name   string
		Params map[string]any
		Errors []string
	}{
		{
			Name: "OK",
			Params: map[string]any{
				"math.Pi": math.Pi,
				"nil":     nil,
			},
		},
		{
			Name: "Errors",
			Params: map[string]any{
				"math..Pi":   math.Pi,
				"slice":      []int{1, 2, 3},
				"emptySLice": []int(nil),
			},
			Errors: []string{
				`parameters: "emptySLice": unsupported type []int`,
				`parameters: "math..Pi": invalid name`,
				`parameters: "slice": unsupported type []int`,
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			i := input.Input{Params: s.Params}

			t.Run("input.ValidateParams", func(t *testing.T) {
				t.Parallel()
				err := input.ValidateParams(i)
				errAssert.EqualErrorGroup(t, err, s.Errors)
			})
			t.Run("input.NewDefaultValidator", func(t *testing.T) {
				t.Parallel()
				err := input.NewDefaultValidator("dev-main").Validate(i)
				errAssert.EqualErrorGroup(t, err, s.Errors)
			})
		})
	}
}
