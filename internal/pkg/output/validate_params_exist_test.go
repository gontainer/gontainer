package output_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/errors/assert"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

func TestValidateParamsExist(t *testing.T) {
	d := output.Output{}
	d.Params = []output.Param{
		{Name: "firstname"},
		{Name: "name", DependsOn: []string{"firstname", "lastname"}},
	}

	expected := []string{
		`output.ValidateParamsExist: "name": param "lastname" does not exist`,
	}

	err := output.ValidateParamsExist(d)
	errAssert.EqualErrorGroup(t, err, expected)
}
