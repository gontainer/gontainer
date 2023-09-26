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
