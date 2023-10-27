package output

import (
	"github.com/gontainer/gontainer-helpers/v2/container/graph"
	"github.com/gontainer/gontainer-helpers/v2/grouperror"
)

func ValidateCircularDeps(o Output) error {
	return grouperror.Prefix(
		"output.ValidateCircularDeps: ",
		graph.CircularDepsToError(o.BuildDependencyGraph().CircularDeps()), //nolint:staticcheck
	)
}
