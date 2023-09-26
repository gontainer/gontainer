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
