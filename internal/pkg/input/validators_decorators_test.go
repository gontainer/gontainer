package input_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
)

func TestDefaultDecoratorsValidators(t *testing.T) {
	assert.NotEmpty(t, input.DefaultDecoratorsValidators())
}

func TestValidateDecorators(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name      string
		Decorator []input.Decorator
		Errors    []string
	}{
		{
			Name: "Errors #1",
			Decorator: []input.Decorator{
				{
					Tag:       "my tag",
					Decorator: "my func",
					Args:      []any{func() {}, nil, (*int)(nil), 17},
				},
			},
			Errors: []string{
				`decorators: 0 "my func": tag: invalid "my tag"`,
				`decorators: 0 "my func": method: invalid "my func"`,
				`decorators: 0 "my func": arguments: 0: unsupported type func()`,
				`decorators: 0 "my func": arguments: 2: unsupported type *int`,
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			i := input.Input{Decorators: s.Decorator}

			t.Run("input.ValidateDecorators", func(t *testing.T) {
				t.Parallel()
				err := input.ValidateDecorators(i)
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
