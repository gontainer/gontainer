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

package resolver

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/gontainer/gontainer/internal/pkg/consts"
	"github.com/gontainer/gontainer/internal/pkg/regex"
	"github.com/gontainer/gontainer/internal/pkg/syntax"
)

var (
	valuePrefixRegex = regexp.MustCompile(`\A(` + regex.PrefixValue + `)`)
	valueRegex       = regex.MustCompileAz(regex.ArgValue)
)

type aliaser interface {
	// Alias returns an alias for given import, e.g. "github.com/spf13/viper" => "i0_viper".
	Alias(string) string
}

type ValueResolver struct {
	aliaser aliaser
}

func NewValueResolver(a aliaser) *ValueResolver {
	return &ValueResolver{aliaser: a}
}

func (v ValueResolver) ResolveArg(p any) (ArgExpr, error) {
	s := p.(string)
	ok, m := regex.Match(valueRegex, s)

	if !ok {
		return ArgExpr{}, errors.New("invalid value")
	}

	return ArgExpr{
		Code:              fmt.Sprintf(consts.TplDependencyValue, syntax.CompileServiceValue(v.aliaser, m["argval"])),
		Raw:               s,
		DependsOnParams:   nil,
		DependsOnServices: nil,
		DependsOnTags:     nil,
	}, nil
}

func (v ValueResolver) Supports(p any) bool {
	s, ok := p.(string)
	if !ok {
		return false
	}
	return valuePrefixRegex.MatchString(s)
}
