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
	"github.com/gontainer/gontainer-helpers/v3/grouperror"
)

type factory interface {
	Create(string) (Token, error)
}

type chunker interface {
	Chunks(string) ([]string, error)
}

type Tokenizer struct {
	chunker chunker
	factory factory
}

func NewTokenizer(ch chunker, f factory) *Tokenizer {
	return &Tokenizer{
		chunker: ch,
		factory: f,
	}
}

func (t *Tokenizer) Tokenize(s string) (Tokens, error) {
	chunks, err := t.chunker.Chunks(s)
	if err != nil {
		return nil, err
	}

	errs := make([]error, len(chunks))
	tkns := make(Tokens, len(chunks))

	for i, ch := range chunks {
		tkns[i], errs[i] = t.factory.Create(ch)
	}

	return tkns, grouperror.Join(errs...)
}
