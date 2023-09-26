package input_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
)

func TestCall_UnmarshalYAML(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name     string
		Input    string
		Expected input.Call
		Error    string
	}{
		{
			Name:  "string",
			Input: "call",
			Error: "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `call` into []interface {}",
		},
		{
			Name:  "empty array",
			Input: "[]",
			Error: "the object Call must contain 1 - 3 values, 0 given",
		},
		{
			Name:  "invalid method name",
			Input: "[~]",
			Error: "first element of the object Call must be a string, `<nil>` given",
		},
		{
			Name:  "invalid args",
			Input: "[Call, arg]",
			Error: "second element of the object Call must be an array, `string` given",
		},
		{
			Name:  "invalid immutable",
			Input: "[Call, [], 5]",
			Error: "third element of the object Call must be a bool, `int` given",
		},
		{
			Name:  "OK #1",
			Input: "[SetTimeout, [3600]]",
			Expected: input.Call{
				Method:    "SetTimeout",
				Args:      []any{3600},
				Immutable: false,
			},
		},
		{
			Name:  "OK #2",
			Input: "[SetTimeout, [3600], false]",
			Expected: input.Call{
				Method:    "SetTimeout",
				Args:      []any{3600},
				Immutable: false,
			},
		},
		{
			Name:  "OK #3 (immutable)",
			Input: "[SetTimeout, [3600], true]",
			Expected: input.Call{
				Method:    "SetTimeout",
				Args:      []any{3600},
				Immutable: true,
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()

			var c input.Call
			err := yaml.Unmarshal([]byte(s.Input), &c)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, s.Expected, c)
		})
	}
}
