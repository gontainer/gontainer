package output

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/maps"
)

func ValidateServicesScopes(o Output) error {
	var errs []error

	services := make(map[string]Service)
	for _, s := range o.Services {
		services[s.Name] = s
	}

	graph := o.BuildDependencyGraph()

	maps.Iterate(services, func(_ string, s Service) {
		if s.Scope != ScopeShared {
			return
		}

		deps := graph.Deps(s.Name)
		for _, d := range deps {
			if !d.IsService() {
				continue
			}
			s2 := services[d.Resource]
			if s2.Scope == ScopeContextual {
				errs = append(
					errs,
					fmt.Errorf("%+q: service is shared, but dependant %+q is contextual", s.Name, s2.Name),
				)
			}
		}
	})

	return grouperror.Prefix("output.ValidateServicesScopes: ", errs...)
}
