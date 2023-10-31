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

package template

import (
	"go/format"
	"regexp"

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"golang.org/x/tools/imports"
)

func init() {
	imports.LocalPrefix = "github.com/gontainer/gontainer-helpers/v2"
}

type CodeFormatter struct {
}

func NewCodeFormatter() *CodeFormatter {
	return &CodeFormatter{}
}

var (
	reEmptyNewLines = regexp.MustCompile(`\n\n+\t`)
)

func (CodeFormatter) Format(c string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = grouperror.Prefix("CodeFormatter.Format: ", err)
		}
	}()

	var r []byte
	r, err = format.Source([]byte(c))
	if err != nil {
		return "", err
	}
	r = reEmptyNewLines.ReplaceAll(r, []byte("\n\t"))

	// remove unused imports
	// required for generating stubs
	r, err = imports.Process("", r, nil)

	return string(r), err
}
