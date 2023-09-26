package token_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestChunker_Chunks(t *testing.T) {
	t.Parallel()

	scenarios := []struct {
		Name   string
		Input  string
		Output []string
		Error  string
	}{
		{
			Name:   "%%",
			Input:  "%%",
			Output: []string{"%%"},
		},
		{
			Name:   "UTF-8",
			Input:  "%✓%",
			Output: []string{"%✓%"},
		},
		{
			Name:   "Name",
			Input:  "%firstname% %lastname%",
			Output: []string{"%firstname%", " ", "%lastname%"},
		},
		{
			Name:  "Error",
			Input: "%firstname% %lastname",
			Error: `not closed token: "%lastname"`,
		},
		{
			Name:  "Single delimiter",
			Input: "%",
			Error: `not closed token: "%"`,
		},
		{
			Name:  "Advanced",
			Input: `mysql://%env("USERNAME")%:%env("PASSWORD")%@%host%:%port%/%db_name%`,
			Output: []string{
				`mysql://`,
				`%env("USERNAME")%`,
				`:`,
				`%env("PASSWORD")%`,
				`@`,
				`%host%`,
				`:`,
				`%port%`,
				`/`,
				`%db_name%`,
			},
		},
		{
			Name:   "Trailing string",
			Input:  "%% text",
			Output: []string{"%%", " text"},
		},
	}

	ch := token.NewChunker()

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()

			o, err := ch.Chunks(s.Input)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error)
				assert.Empty(t, o)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, s.Output, o)
		})
	}
}
