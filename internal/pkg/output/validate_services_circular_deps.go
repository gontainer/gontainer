package output

import (
	"github.com/gontainer/gontainer-helpers/v2/container/graph"
	"github.com/gontainer/gontainer-helpers/v2/grouperror"
)

func ValidateServicesCircularDeps(o Output) error {
	return grouperror.Prefix(
		"output.ValidateServicesCircularDeps: ",
		graph.CircularDepsToError(o.BuildDependencyGraph().CircularDeps()),
	)
}
