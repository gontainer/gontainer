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

package token

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gontainer/gontainer/internal/pkg/consts"
)

type Kind uint

const (
	KindString    Kind = iota // raw string
	KindReference             // e.g. %name%
	KindFunc                  // e.g. %env("APP_HOST")%
)

type Token struct {
	Kind      Kind
	Raw       string
	DependsOn []string // list of dependencies for KindReference
	Code      string   // GO code in the following format of func: func() (interface{}, error) { ... }
}

type Tokens []Token

// GoCode returns Go code that is such func `func() (interface{}, error)`,surrounded by `dependencyProvider(%s)`.
// Possibly `dependencyValue(%s)` in the future.
//
//	dependencyProvider(func() (interface{}, error) { return "hello world", nil })
func (tkns Tokens) GoCode() (string, error) {
	if len(tkns) == 0 {
		return "", errors.New("unexpected error: len(tokens) == 0")
	}

	if len(tkns) == 1 { // single token may return a non-string value, e.g. %envInt("PORT")%
		return fmt.Sprintf(consts.TplDependencyProvider, tkns[0].Code), nil
	}

	parts := make([]string, 0)
	for _, t := range tkns {
		parts = append(parts, t.Code)
	}

	return fmt.Sprintf(consts.TplDependencyConcatenateChunks, strings.Join(parts, ", ")), nil
}
