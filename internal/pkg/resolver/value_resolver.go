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
