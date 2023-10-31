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

package input_test

import (
	"fmt"
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
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
			Name:     "non_shared",
			Input:    "non_shared",
			Expected: input.ScopeNonShared,
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
