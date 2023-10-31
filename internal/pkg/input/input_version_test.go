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
			Error: `version must follow the semver scheme, and it must not be prefixed by "v", see https://semver.org/`,
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
