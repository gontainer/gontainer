package output_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

func TestValidateParamsExist(t *testing.T) {
	d := output.Output{}
	d.Params = []output.Param{
		{Name: "firstname"},
		{Name: "name", DependsOn: []string{"firstname", "lastname"}},
	}
	d.Services = []output.Service{
		{
			Name: "db",
			Fields: []output.Field{
				{Name: "Host", Value: output.Arg{DependsOnParams: []string{"host"}}},
				{Name: "Port", Value: output.Arg{DependsOnParams: []string{"port"}}},
			},
		},
	}

	expected := []string{
		`output.ValidateParamsExist: "%name%": param "lastname" does not exist`,
		`output.ValidateParamsExist: "@db": param "host" does not exist`,
		`output.ValidateParamsExist: "@db": param "port" does not exist`,
	}

	err := output.ValidateParamsExist(d)
	errAssert.EqualErrorGroup(t, err, expected)
}
