package output

import (
	"github.com/gontainer/gontainer-helpers/container/graph"
	"github.com/gontainer/gontainer-helpers/grouperror"
)

func ValidateServicesCircularDeps(o Output) error {
	return grouperror.Prefix(
		"output.ValidateServicesCircularDeps: ",
		graph.CircularDepsToError(o.BuildDependencyGraph().CircularDeps()),
	)
}
