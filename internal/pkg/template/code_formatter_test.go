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

package template_test

import (
	_ "embed"
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/template"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed code_formatter_test.go
	goCode string
)

func TestCodeFormatter_Format(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		f := template.NewCodeFormatter()
		_, err := f.Format("invalid go code")
		assert.ErrorContains(t, err, "CodeFormatter.Format: ")
	})

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		f := template.NewCodeFormatter()
		o, err := f.Format(goCode)
		assert.NoError(t, err)
		assert.NotEmpty(t, o)
	})
}
