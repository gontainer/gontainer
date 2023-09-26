package resolver

import (
	"fmt"
)

type ArgExpr struct {
	// Code is Go code in on of the following formats:
	// 	* dependencyProvider(%s)
	// 	* dependencyTag(%s)
	// 	* dependencyValue(%s)
	// 	* dependencyService(%s)
	Code              string
	Raw               any
	DependsOnParams   []string
	DependsOnServices []string
	DependsOnTags     []string
}

type ParamExpr struct {
	// Code is Go code in on of the following formats:
	// 	* dependencyProvider(%s)
	// 	* dependencyTag(%s)
	// 	* dependencyValue(%s)
	// 	* dependencyService(%s)
	Code            string
	Raw             any
	DependsOnParams []string
}

type resolverStrategy interface {
	ResolveArg(any) (ArgExpr, error)
	Supports(any) bool
}

type ArgResolver struct {
	strategies []resolverStrategy
}

func NewArgResolver(s ...resolverStrategy) *ArgResolver {
	return &ArgResolver{strategies: s}
}

func (a *ArgResolver) ResolveArg(i any) (e ArgExpr, _ error) {
	for _, s := range a.strategies {
		if s.Supports(i) {
			return s.ResolveArg(i)
		}
	}
	return e, fmt.Errorf("not supported %T", i)
}
