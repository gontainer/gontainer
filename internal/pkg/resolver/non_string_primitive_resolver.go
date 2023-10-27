package resolver

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/v2/exporter"
	"github.com/gontainer/gontainer/internal/pkg/consts"
	"github.com/gontainer/gontainer/internal/pkg/types"
)

type NonStringPrimitiveResolver struct {
}

func NewNonStringPrimitiveResolver() *NonStringPrimitiveResolver {
	return &NonStringPrimitiveResolver{}
}

func (NonStringPrimitiveResolver) ResolveArg(i any) (e ArgExpr, _ error) {
	// Method NonStringPrimitiveResolver{}.Supports checks whether the underlying type of `i` is primitive.
	// exporter.MustExport never panics for primitive types, so there is no reason to handle an error.
	return ArgExpr{
		Code:              fmt.Sprintf(consts.TplDependencyValue, exporter.MustExport(i)),
		Raw:               i,
		DependsOnParams:   nil,
		DependsOnServices: nil,
		DependsOnTags:     nil,
	}, nil
}

func (NonStringPrimitiveResolver) Supports(i any) bool {
	_, ok := i.(string)
	return !ok && types.IsPrimitive(i)
}
