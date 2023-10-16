package output_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

func TestValidateServicesScopes(t *testing.T) {
	o := output.Output{}
	o.Services = []output.Service{
		{
			Name:  "transaction",
			Scope: output.ScopeContextual,
		},
		{
			Name:  "userRepository",
			Scope: output.ScopeShared,
			Args:  []output.Arg{{DependsOnServices: []string{"transaction"}}},
		},
		{
			Name:  "userService",
			Scope: output.ScopeShared,
			Args:  []output.Arg{{DependsOnServices: []string{"userRepository"}, DependsOnTags: []string{"some-tag"}}},
		},
	}

	expected := []string{
		`output.ValidateServicesScopes: "userRepository": service is shared, but dependant "transaction" is contextual`,
		`output.ValidateServicesScopes: "userService": service is shared, but dependant "transaction" is contextual`,
	}

	err := output.ValidateServicesScopes(o)
	errAssert.EqualErrorGroup(t, err, expected)
}
