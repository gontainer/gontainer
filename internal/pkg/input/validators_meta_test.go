package input_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/stretchr/testify/assert"
)

func TestValidateMeta(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name   string
		Meta   input.Meta
		Errors []string
	}{
		{
			Name: "Errors #1",
			Meta: input.Meta{
				Pkg:                  ptr.New("some pkg"),
				ContainerType:        ptr.New("$"),
				ContainerConstructor: ptr.New("New Container"),
				Imports: map[string]string{
					`viper`:  `github.com/spf13/viper`,
					`viper2`: `"github.com/spf13/viper"`,
					`_viper`: `github.com/spf13 viper`,
				},
				Functions: map[string]string{
					"correctFn":    `"pkg".Env`,
					"incorrect Fn": `pkg IncorrectEnv`,
				},
			},
			Errors: []string{
				`meta: pkg: invalid "some pkg"`,
				`meta: container_type: invalid "$"`,
				`meta: container_constructor: invalid "New Container"`,
				`meta: imports: invalid import "github.com/spf13 viper"`,
				`meta: imports: invalid alias "_viper"`,
				`meta: functions: invalid function "incorrect Fn"`,
				`meta: functions: invalid go function "pkg IncorrectEnv"`,
			},
		},
		{
			Name: "OK #1",
			Meta: input.Meta{
				Pkg:                  ptr.New("gontainer"),
				ContainerConstructor: ptr.New("New"),
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			i := input.Input{Meta: s.Meta}

			t.Run("input.ValidateMeta", func(t *testing.T) {
				t.Parallel()
				err := input.ValidateMeta(i)
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

func TestDefaultMetaValidators(t *testing.T) {
	assert.NotEmpty(t, input.DefaultMetaValidators())
}
