package output

import (
	"fmt"
	"strings"

	"github.com/gontainer/gontainer-helpers/graph"
	"github.com/gontainer/gontainer-helpers/grouperror"
)

func ValidateParamsCircularDeps(o Output) error {
	depsGraph := graph.New()
	for _, p := range o.Params {
		for _, dep := range p.DependsOn {
			depsGraph.AddDep(p.Name, dep)
		}
	}

	var errs []error

	for _, cycle := range depsGraph.CircularDeps() {
		errs = append(errs, fmt.Errorf("%s", strings.Join(cycle, " -> ")))
	}

	return grouperror.Prefix("output.ValidateParamsCircularDeps: ", errs...)
}
