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

package syntax

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeImport(t *testing.T) {
	scenarios := []struct {
		input  string
		output string
	}{
		{
			input:  `"pkg"`,
			output: "pkg",
		},
		{
			input:  "config",
			output: "config",
		},
	}

	for i, s := range scenarios {
		t.Run(fmt.Sprintf("Scenario #%d", i), func(t *testing.T) {
			assert.Equal(t, s.output, SanitizeImport(s.input))
		})
	}
}

func TestCompileServiceValue(t *testing.T) {
	scenarios := []struct {
		input  string
		output string
	}{
		{
			input:  "Variable",
			output: "Variable",
		},
		{
			input:  "pkg.Variable",
			output: "i0_alias.Variable",
		},
		{
			input:  "&User{}",
			output: "&User{}",
		},
		{
			input:  `&"my/import/path".User{}`,
			output: "&i0_alias.User{}",
		},
		{
			input:  `".".GlobalConfig.DB`,
			output: "GlobalConfig.DB",
		},
	}

	a := mockAliases{alias: "i0_alias"}

	for i, s := range scenarios {
		t.Run(fmt.Sprintf("Scenario #%d", i), func(t *testing.T) {
			assert.Equal(t, s.output, CompileServiceValue(a, s.input))
		})
	}
}

type mockAliases struct {
	alias string
}

func (m mockAliases) Alias(string) string {
	return m.alias
}
