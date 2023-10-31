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
)

var (
	servicePrefixRegex = regexp.MustCompile(`\A(` + regex.PrefixService + `)`)
	serviceRegex       = regex.MustCompileAz(regex.ArgService)
)

type ServiceResolver struct {
	patternGetService string
}

func NewServiceResolver() *ServiceResolver {
	return &ServiceResolver{
		patternGetService: consts.TplDependencyService,
	}
}

func (s ServiceResolver) ResolveArg(i any) (ArgExpr, error) {
	st := i.(string)
	ok, m := regex.Match(serviceRegex, st)

	if !ok {
		return ArgExpr{}, errors.New("invalid service")
	}

	return ArgExpr{
		Code:              fmt.Sprintf(s.patternGetService, m["service"]),
		Raw:               i,
		DependsOnParams:   nil,
		DependsOnServices: []string{m["service"]},
		DependsOnTags:     nil,
	}, nil
}

func (s ServiceResolver) Supports(i any) bool {
	st, ok := i.(string)
	if !ok {
		return false
	}
	return servicePrefixRegex.MatchString(st)
}
