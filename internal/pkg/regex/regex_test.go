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

package regex

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		regex  string
		input  string
		match  bool
		result map[string]string
	}{
		{
			regex:  "^(?P<firstname>[A-Z][a-z]+) (?P<lastname>[A-Z][a-z]+)$",
			input:  "Jane Doe",
			match:  true,
			result: map[string]string{"firstname": "Jane", "lastname": "Doe"},
		},
		{
			regex:  "^(?P<fullname>(?P<firstname>[A-Z][a-z]+) (?P<lastname>[A-Z][a-z]+))$",
			input:  "Jane Doe",
			match:  true,
			result: map[string]string{"firstname": "Jane", "lastname": "Doe", "fullname": "Jane Doe"},
		},
		{
			regex:  "^(?P<firstname>[A-Z][a-z]+) (?P<lastname>[A-Z][a-z]+)$",
			input:  "Jane Doe-Doe",
			match:  false,
			result: nil,
		},
	}

	for id, tmp := range scenarios {
		s := tmp
		t.Run(fmt.Sprintf("scenario #%d", id), func(t *testing.T) {
			t.Parallel()
			m, r := Match(regexp.MustCompile(s.regex), s.input)
			assert.Equal(t, s.match, m)
			assert.Equal(t, s.result, r)
		})
	}
}

func TestMustCompileWrapped(t *testing.T) {
	t.Run("Given scenario", func(t *testing.T) {
		r := MustCompileAz(".*")
		assert.Equal(t, `\A(.*)\z`, r.String())
	})
	t.Run("Given error", func(t *testing.T) {
		defer func() {
			assert.NotNil(t, recover())
		}()
		MustCompileAz(`[A-z`)
	})
}
