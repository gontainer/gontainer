package output_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

func TestValidateParamsCircularDeps(t *testing.T) {
	o := output.Output{}
	o.Params = []output.Param{
		{
			Name:      "self",
			DependsOn: []string{"self"},
		},
		{
			Name:      "firstname",
			DependsOn: []string{"name"},
		},
		{
			Name:      "name",
			DependsOn: []string{"firstname", "lastname"},
		},
	}

	expected := []string{
		`output.ValidateParamsCircularDeps: self -> self`,
		`output.ValidateParamsCircularDeps: firstname -> name -> firstname`,
	}

	err := output.ValidateParamsCircularDeps(o)
	errAssert.EqualErrorGroup(t, err, expected)
}
