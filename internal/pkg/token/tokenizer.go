package token

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

func (t *Tokenizer) Tokenize(s string) (r Tokens, _ error) {
	chunks, err := t.chunker.Chunks(s)
	if err != nil {
		return nil, err
	}

	for _, ch := range chunks {
		tk, err := t.factory.Create(ch)
		if err != nil {
			return nil, err
		}
		r = append(r, tk)
	}

	return r, nil
}
