package input_test

import (
	"fmt"
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
)

func TestScope_UnmarshalYAML(t *testing.T) {
	t.Parallel()

	scenarios := []struct {
		Name     string
		Input    string
		Expected input.Scope
		Error    string
	}{
		{
			Name:     "shared",
			Input:    "shared",
			Expected: input.ScopeShared,
		},
		{
			Name:     "contextual",
			Input:    "contextual",
			Expected: input.ScopeContextual,
		},
		{
			Name:     "private",
			Input:    "private",
			Expected: input.ScopePrivate,
		},
		{
			Name:  "error #1",
			Input: "my_scope",
			Error: `invalid value for input.Scope: "my_scope"`,
		},
		{
			Name:  "error #2",
			Input: "[",
			Error: `yaml: line 1: did not find expected node content`,
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()

			var sc input.Scope
			err := yaml.Unmarshal([]byte(s.Input), &sc)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, s.Expected, sc)
		})
	}
}

func TestScope_String(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		input  input.Scope
		output string
	}{
		{
			input:  input.ScopeShared,
			output: "shared",
		},
		{
			input:  input.Scope(0),
			output: "invalid (0)",
		},
	}

	for i, tmp := range scenarios {
		s := tmp
		t.Run(fmt.Sprintf("scenario #%d", i), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, s.output, s.input.String())
		})
	}
}
