package token_test

import (
	"fmt"
	"regexp"

	"github.com/gontainer/gontainer/internal/pkg/regex"
)

type mockAliaser struct {
}

var (
	aliasRegex = regexp.MustCompile(`(?P<last>[a-zA-Z][a-zA-Z0-9]*)\z`)
)

func (mockAliaser) Alias(s string) string {
	ok, m := regex.Match(aliasRegex, s)
	if !ok {
		panic(fmt.Sprintf("token_test: mockAliaser.Alias: invalid import `%s`", s))
	}
	return m["last"]
}
