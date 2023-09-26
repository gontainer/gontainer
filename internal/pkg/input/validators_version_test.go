package input_test

import (
	"fmt"
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/stretchr/testify/assert"
)

func TestVersionValidator_ValidateVersion(t *testing.T) {
	t.Parallel()

	scenarios := []struct {
		libraryVer string
		givenVer   *input.Version
		error      string
	}{
		{
			libraryVer: "0.1.0",
			givenVer:   ptr.New(input.Version("v0.1.1")),
			error:      "",
		},
		{
			libraryVer: "1.0.0-alpha.1",
			givenVer:   ptr.New(input.Version("v1.0.0")),
			error:      "",
		},
		{
			libraryVer: "dev-master",
			givenVer:   ptr.New(input.Version("v999.999.999")),
			error:      "",
		},
		{
			libraryVer: "0.3.0",
			givenVer:   ptr.New(input.Version("v0.2.0")),
			error:      "version: current: v0.3.0, given: v0.2.0: possibly incompatible versions",
		},
		{
			libraryVer: "1.3.0",
			givenVer:   ptr.New(input.Version("v1.2.0")),
			error:      "",
		},
		{
			libraryVer: "1.2.0",
			givenVer:   ptr.New(input.Version("v1.3.0")),
			error:      "version: current: v1.2.0, given: v1.3.0: update Gontainer to use all new features",
		},
		{
			libraryVer: "1.0.0",
			givenVer:   ptr.New(input.Version("v2.0.0")),
			error:      "version: current: v1.0.0, given: v2.0.0: incompatible versions",
		},
		{
			libraryVer: "1.0.0",
			givenVer:   ptr.New(input.Version("v0.999.0")),
			error:      "version: current: v1.0.0, given: v0.999.0: incompatible versions",
		},
	}

	for i, tmp := range scenarios {
		s := tmp
		t.Run(fmt.Sprintf("scenario #%d", i), func(t *testing.T) {
			t.Parallel()
			i := input.Input{
				Version: s.givenVer,
			}

			t.Run("input.NewVersionValidator", func(t *testing.T) {
				t.Parallel()
				err := input.NewVersionValidator(s.libraryVer).ValidateVersion(i)
				if s.error == "" {
					assert.NoError(t, err)
					return
				}
				assert.EqualError(t, err, s.error)
			})
			t.Run("input.NewDefaultValidator", func(t *testing.T) {
				t.Parallel()
				err := input.NewDefaultValidator(s.libraryVer).Validate(i)
				if s.error == "" {
					assert.NoError(t, err)
					return
				}
				assert.EqualError(t, err, s.error)
			})
		})
	}
}
