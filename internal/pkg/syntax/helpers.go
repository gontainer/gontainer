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
	"strings"

	"github.com/gontainer/gontainer/internal/pkg/regex"
)

var (
	regexServiceValue = regex.MustCompileAz(regex.ServiceValue)
)

type aliaser interface {
	Alias(string) string
}

// SanitizeImport removes surrounding quotation marks.
// If `result == "."`, it returns an empty string.
func SanitizeImport(i string) string {
	r := strings.Trim(i, `"`)
	if r == "." {
		r = ""
	}
	return r
}

// CompileServiceValue expects correct expr, validation must be done earlier.
func CompileServiceValue(a aliaser, expr string) string {
	_, m := regex.Match(regexServiceValue, expr)

	if m["v1"] != "" {
		parts := make([]string, 0)
		import_ := SanitizeImport(m["import"])
		if import_ != "" {
			parts = append(parts, a.Alias(import_))
		}
		return m["ptr"] + strings.Join(append(parts, m["value"]), ".")
	}

	parts := make([]string, 0)
	import_ := SanitizeImport(m["import2"])
	if import_ != "" {
		parts = append(parts, a.Alias(import_))
	}
	return m["ptr2"] + strings.Join(append(parts, m["struct2"]), ".") + "{}"
}
