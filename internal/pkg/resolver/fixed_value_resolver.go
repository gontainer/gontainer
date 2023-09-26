package resolver

import (
	"fmt"

	"github.com/gontainer/gontainer/internal/pkg/consts"
)

type FixedValueResolver struct {
	id    string
	value string
}

func NewFixedValueResolver(id string, value string) *FixedValueResolver {
	return &FixedValueResolver{id: id, value: value}
}

func (f FixedValueResolver) ResolveArg(a any) (ArgExpr, error) {
	return ArgExpr{
		Code:              fmt.Sprintf(consts.TplDependencyValue, f.value),
		Raw:               a,
		DependsOnParams:   nil,
		DependsOnServices: nil,
		DependsOnTags:     nil,
	}, nil
}

func (f FixedValueResolver) Supports(a any) bool {
	s, ok := a.(string)
	return ok && f.id == s
}
