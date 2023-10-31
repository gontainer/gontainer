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
	"github.com/gontainer/gontainer/internal/pkg/token"
)

type tokenizer interface {
	Tokenize(pattern string) (token.Tokens, error)
}

type PatternResolver struct {
	tokenizer tokenizer
}

func NewPatternResolver(t tokenizer) *PatternResolver {
	return &PatternResolver{tokenizer: t}
}

func (p PatternResolver) ResolveArg(i any) (e ArgExpr, _ error) {
	v := i.(string)
	tkns, err := p.tokenizer.Tokenize(v)
	if err != nil {
		return e, err
	}

	var dependsOn []string

	c, err := tkns.GoCode()
	if err != nil {
		return e, err
	}

	for _, t := range tkns {
		dependsOn = append(dependsOn, t.DependsOn...)
	}

	return ArgExpr{
		Code:              c,
		Raw:               v,
		DependsOnParams:   dependsOn,
		DependsOnServices: nil,
		DependsOnTags:     nil,
	}, nil
}

func (p PatternResolver) Supports(i any) bool {
	_, ok := i.(string)
	return ok
}
