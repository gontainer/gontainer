// Copyright (c) 2023 Bartłomiej Krukowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is furnished
// to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
			`output.ValidateCircularDeps: @company -> @department -> @team -> @hr -> !tagged organization -> @holding -> @company`,
			`output.ValidateCircularDeps: @company -> @department -> @team -> @hr -> !tagged organization -> @company`,
			`output.ValidateCircularDeps: @holding -> @holding`,
			`output.ValidateCircularDeps: @department -> @team -> @department`,
			`output.ValidateCircularDeps: @chro -> @hr -> @chro`,
			`output.ValidateCircularDeps: @hr -> @hr`,
			`output.ValidateCircularDeps: %self% -> %self%`,
			`output.ValidateCircularDeps: %firstname% -> %name% -> %firstname%`,
		}

		err := output.ValidateCircularDeps(o)
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
			`output.ValidateCircularDeps: @db -> decorate(!tagged sql.DB) -> decorator(#0) -> @db`,
			`output.ValidateCircularDeps: @serviceA -> decorate(!tagged tagB) -> decorator(#1) -> !tagged tagB -> @serviceA`,
		}

		err := output.ValidateCircularDeps(o)
		errAssert.EqualErrorGroup(t, err, expected)
	})
}
