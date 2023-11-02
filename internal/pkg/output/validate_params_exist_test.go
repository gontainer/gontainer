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
