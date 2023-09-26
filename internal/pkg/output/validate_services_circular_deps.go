package output

import (
	"github.com/gontainer/gontainer-helpers/container/graph"
	"github.com/gontainer/gontainer-helpers/errors"
)

func ValidateServicesCircularDeps(o Output) error {
	return errors.PrefixedGroup(
		"output.ValidateServicesCircularDeps: ",
		graph.CircularDepsToError(o.BuildDependencyGraph().CircularDeps()),
	)
}
