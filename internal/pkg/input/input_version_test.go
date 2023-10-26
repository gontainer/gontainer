package input_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestVersion_UnmarshalYAML(t *testing.T) {
	t.Parallel()

	scenarios := []struct {
		Input    string
		Expected input.Version
		Error    string
	}{
		{
			Input:    `"0.1.0"`,
			Expected: input.Version("0.1.0"),
		},
		{
			Input:    `0.1.0`,
			Expected: input.Version("0.1.0"),
		},
		{
			Input: `v0.1.0`,
			Error: `version must follow the semver scheme, and it must not be Prefix by "v", see https://semver.org/`,
		},
		{
			Input: `5`,
			Error: `it must be a string`,
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Input, func(t *testing.T) {
			t.Parallel()

			var ver input.Version
			err := yaml.Unmarshal([]byte(s.Input), &ver)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, s.Expected, ver)
		})
	}
}
