package token

import (
	"github.com/gontainer/gontainer-helpers/grouperror"
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
