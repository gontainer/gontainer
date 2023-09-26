package resolver

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/gontainer/gontainer/internal/pkg/consts"
	"github.com/gontainer/gontainer/internal/pkg/regex"
)

var (
	taggedPrefixRegex = regexp.MustCompile(`\A(` + regex.PrefixTagged + `)`)
	taggedRegex       = regex.MustCompileAz(regex.ArgTagged)
)

type TaggedResolver struct {
	patternGetByTag string
}

func NewTaggedResolver() *TaggedResolver {
	return &TaggedResolver{
		patternGetByTag: consts.TplDependencyTag,
	}
}

func (t TaggedResolver) ResolveArg(i any) (ArgExpr, error) {
	s := i.(string)
	ok, m := regex.Match(taggedRegex, s)

	if !ok {
		return ArgExpr{}, errors.New("invalid tag")
	}

	return ArgExpr{
		Code:              fmt.Sprintf(t.patternGetByTag, m["tag"]),
		Raw:               s,
		DependsOnParams:   nil,
		DependsOnServices: nil,
		DependsOnTags:     []string{m["tag"]},
	}, nil
}

func (t TaggedResolver) Supports(p any) bool {
	s, ok := p.(string)
	if !ok {
		return false
	}
	return taggedPrefixRegex.MatchString(s)
}
