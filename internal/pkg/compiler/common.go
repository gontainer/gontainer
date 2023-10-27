package compiler

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

func argExprToArg(e resolver.ArgExpr) output.Arg {
	return output.Arg{
		Code:              e.Code,
		Raw:               e.Raw,
		DependsOnParams:   e.DependsOnParams,
		DependsOnServices: e.DependsOnServices,
		DependsOnTags:     e.DependsOnTags,
	}
}

func resolveArgs(resolver argResolver, args []any) (r []output.Arg, _ error) {
	var errs []error
	if len(args) > 0 {
		r = make([]output.Arg, len(args))
	}
	for i, arg := range args {
		argExpr, err := resolver.ResolveArg(arg)
		errs = append(errs, grouperror.Prefix(fmt.Sprintf("%d: ", i), err))
		r[i] = argExprToArg(argExpr)
	}
	return r, grouperror.Prefix("args: ", errs...)
}
