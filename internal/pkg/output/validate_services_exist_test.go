package output_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

func TestValidateServicesExist(t *testing.T) {
	d := output.Output{}
	d.Services = []output.Service{
		{
			Name: "company",
			Args: []output.Arg{{
				DependsOnServices: []string{"holding", "brand"},
			}},
		},
		{
			Name: "hr",
			Args: []output.Arg{{
				DependsOnServices: []string{"holding", "hr"},
			}},
		},
	}
	d.Decorators = []output.Decorator{
		{
			Tag:       "my-tag",
			Decorator: "",
			Args:      []output.Arg{{DependsOnServices: []string{"db"}}},
		},
	}

	expected := []string{
		`output.ValidateServicesExist: "company": service "holding" does not exist`,
		`output.ValidateServicesExist: "company": service "brand" does not exist`,
		`output.ValidateServicesExist: "hr": service "holding" does not exist`,
		`output.ValidateServicesExist: decorator(#0, "my-tag"): service "db" does not exist`,
	}

	err := output.ValidateServicesExist(d)
	errAssert.EqualErrorGroup(t, err, expected)
}
