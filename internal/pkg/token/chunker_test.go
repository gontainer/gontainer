// Copyright (c) 2023 Bartłomiej Krukowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is furnished
// to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
