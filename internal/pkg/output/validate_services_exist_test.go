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

	errAssert "github.com/gontainer/gontainer-helpers/v3/grouperror/assert"
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
