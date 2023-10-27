package resolver

import (
	"fmt"
	"strings"

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
)

type argResolver interface {
	ResolveArg(any) (ArgExpr, error)
}

// ParamResolver resolves input parameters. In opposition to arguments,
// parameters may depend o on other parameters only.
type ParamResolver struct {
	resolver argResolver
}

func NewParamResolver(resolver argResolver) *ParamResolver {
	return &ParamResolver{resolver: resolver}
}

func (p ParamResolver) ResolveParam(i any) (ParamExpr, error) {
	a, err := p.resolver.ResolveArg(i)
	// params cannot depend on services or tags
	// if they do, most likely a wrong ArgResolver has been injected
	// to the constructor NewParamResolver
	var errs []error
	if len(a.DependsOnServices) > 0 {
		errs = append(
			errs,
			fmt.Errorf("param cannot depend on any service: %s", strings.Join(a.DependsOnServices, ", ")),
		)
	}
	if len(a.DependsOnTags) > 0 {
		errs = append(
			errs,
			fmt.Errorf("param cannot depend on any tag: %s", strings.Join(a.DependsOnTags, ", ")),
		)
	}
	if len(errs) > 0 {
		return ParamExpr{}, grouperror.Join(errs...)
	}
	return ParamExpr{
		Code:            a.Code,
		Raw:             a.Raw,
		DependsOnParams: a.DependsOnParams,
	}, err
}
