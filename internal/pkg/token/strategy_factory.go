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
	"fmt"
)

type tokenFactoryStrategy interface {
	Supports(string) bool
	Create(string) (Token, error)
}

type StrategyFactory struct {
	strategies []tokenFactoryStrategy
}

func NewStrategyFactory(strategies ...tokenFactoryStrategy) *StrategyFactory {
	return &StrategyFactory{strategies: strategies}
}

func (f *StrategyFactory) Prepend(s tokenFactoryStrategy) {
	f.strategies = append([]tokenFactoryStrategy{s}, f.strategies...)
}

func (f *StrategyFactory) Create(i string) (t Token, _ error) {
	for _, s := range f.strategies {
		if s.Supports(i) {
			return s.Create(i)
		}
	}
	return t, fmt.Errorf("not supported token: %s", i)
}

type FuncRegisterer struct {
	prepender interface{ Prepend(s tokenFactoryStrategy) }
	aliaser   aliaser
}

func NewFuncRegisterer(
	p interface{ Prepend(s tokenFactoryStrategy) },
	a aliaser,
) *FuncRegisterer {
	return &FuncRegisterer{
		prepender: p,
		aliaser:   a,
	}
}

func (f *FuncRegisterer) RegisterFunc(fnAlias string, goImport string, goFn string) {
	f.prepender.Prepend(NewFactoryFunction(f.aliaser, fnAlias, goImport, goFn))
}
