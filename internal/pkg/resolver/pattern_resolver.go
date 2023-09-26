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
		return e, err // todo decorate?
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
