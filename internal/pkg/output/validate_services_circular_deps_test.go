package output_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

func TestValidateServicesCircularDeps(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		o := output.Output{}
		o.Services = []output.Service{
			{
				Name: "holding",
				Args: []output.Arg{{
					DependsOnServices: []string{"company", "holding"},
				}},
				Tags: []output.Tag{
					{Name: "organization"},
				},
			},
			{
				Name: "company",
				Args: []output.Arg{{
					DependsOnServices: []string{"department"},
				}},
				Tags: []output.Tag{
					{Name: "organization"},
				},
			},
			{
				Name: "department",
				Args: []output.Arg{{
					DependsOnServices: []string{"team"},
				}},
			},
			{
				Name: "team",
				Args: []output.Arg{{
					DependsOnServices: []string{"hr"},
				}},
				Fields: []output.Field{{
					Name:  "Department",
					Value: output.Arg{DependsOnServices: []string{"department"}},
				}},
			},
			{
				Name: "hr",
				Args: []output.Arg{{
					DependsOnServices: []string{"chro"},
				}},
				Calls: []output.Call{{
					Method: "Organizations",
					Args:   []output.Arg{{DependsOnTags: []string{"organization"}}},
				}},
				Fields: []output.Field{{
					Name:  "Team",
					Value: output.Arg{DependsOnServices: []string{"hr"}},
				}},
			},
			{
				Name: "chro",
				Args: []output.Arg{{
					DependsOnServices: []string{"hr"},
				}},
			},
		}

		expected := []string{
			`output.ValidateServicesCircularDeps: @company -> @department -> @team -> @hr -> !tagged organization -> @holding -> @company`,
			`output.ValidateServicesCircularDeps: @company -> @department -> @team -> @hr -> !tagged organization -> @company`,
			`output.ValidateServicesCircularDeps: @holding -> @holding`,
			`output.ValidateServicesCircularDeps: @department -> @team -> @department`,
			`output.ValidateServicesCircularDeps: @chro -> @hr -> @chro`,
			`output.ValidateServicesCircularDeps: @hr -> @hr`,
		}

		err := output.ValidateServicesCircularDeps(o)
		errAssert.EqualErrorGroup(t, err, expected)
	})

	t.Run("Decorators", func(t *testing.T) { // !tagged storage -> db
		o := output.Output{
			Meta:   output.Meta{},
			Params: nil,
			Services: []output.Service{
				{
					Name: "db",
					Tags: []output.Tag{{Name: "sql.DB"}},
				},
				{
					Name: "serviceA",
					Tags: []output.Tag{{Name: "tagB"}},
				},
			},
			Decorators: []output.Decorator{
				{
					Tag:  "sql.DB",
					Raw:  "traceableDB",
					Args: []output.Arg{{DependsOnServices: []string{"db"}}},
				},
				{
					Tag:  "tagB",
					Raw:  "decoratorC",
					Args: []output.Arg{{DependsOnTags: []string{"tagB"}}},
				},
			},
		}
		expected := []string{
			`output.ValidateServicesCircularDeps: @db -> decorate(!tagged sql.DB) -> decorator(#0) -> @db`,
			`output.ValidateServicesCircularDeps: @serviceA -> decorate(!tagged tagB) -> decorator(#1) -> !tagged tagB -> @serviceA`,
		}

		err := output.ValidateServicesCircularDeps(o)
		errAssert.EqualErrorGroup(t, err, expected)
	})
}
