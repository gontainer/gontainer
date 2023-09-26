package input_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestTag_UnmarshalYAML(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name     string
		Input    string
		Expected input.Tag
		Error    string
	}{
		{
			Name:  "int",
			Input: "5",
			Error: "unexpected type `int`",
		},
		{
			Name:     "string",
			Input:    `"my-tag"`,
			Expected: input.Tag{Name: "my-tag", Priority: 0},
		},
		{
			Name:     "object ok #1",
			Input:    `{"name": "my-tag-with-priority", "priority": 50}`,
			Expected: input.Tag{Name: "my-tag-with-priority", Priority: 50},
		},
		{
			Name:     "object ok #2",
			Input:    `{"name": "my-tag-with-priority"}`,
			Expected: input.Tag{Name: "my-tag-with-priority", Priority: 0},
		},
		{
			Name:  "object err #1",
			Input: `{"name": 50, "priority": 50}`,
			Error: "name must be an instance of string",
		},
		{
			Name:  "object err #2",
			Input: `{"priority": 50}`,
			Error: "missing tag name",
		},
		{
			Name:  "object err #3",
			Input: `{"name": "my-tag", "priority": "50"}`,
			Error: "priority must be an int",
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()

			var tag input.Tag
			err := yaml.Unmarshal([]byte(s.Input), &tag)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, s.Expected, tag)
		})
	}
}
