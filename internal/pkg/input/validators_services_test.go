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

package input_test

import (
	"testing"
	"time"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/stretchr/testify/assert"
)

func TestValidateServices(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name     string
		Services map[string]input.Service
		Errors   []string
	}{
		{
			Name: "Errors #1",
			Services: map[string]input.Service{
				"data base": {
					Constructor: ptr.New("New DB"),
					Getter:      ptr.New("GetParam"),
					Args:        []any{7, struct{}{}},
				},
				"db": {
					Constructor: ptr.New("NewDB"),
					Getter:      ptr.New("MustGetDB"),
					Calls: []input.Call{
						{Method: "ping()", Args: []any{struct{}{}}},
					},
					Fields: map[string]any{
						"Host":  "127.0.0.1",
						"Port ": complex(3306, 0),
					},
					Tags: []input.Tag{
						{Name: "sql"},
						{Name: "sql"},
						{Name: "_storage"},
					},
				},
				"ctx": {
					Constructor: ptr.New("NewContext"),
					Getter:      ptr.New("MustGetContextInContext"),
				},
			},
			Errors: []string{
				`services: "ctx": getter: prefix "Must" is not allowed`,
				`services: "ctx": getter: suffix "InContext" is not allowed`,
				`services: "data base": invalid name`,
				`services: "data base": constructor: invalid "New DB"`,
				`services: "data base": getter: "GetParam" is reserved`,
				`services: "data base": arguments: arg 1: unsupported type struct {}`,
				`services: "db": getter: prefix "Must" is not allowed`,
				`services: "db": calls: 0: method: invalid "ping()"`,
				`services: "db": calls: 0: arguments: 0: unsupported type struct {}`,
				`services: "db": fields: "Port ": invalid "Port "`,
				`services: "db": fields: "Port ": unsupported type complex128`,
				`services: "db": tags: 2: invalid "_storage"`,
				`services: "db": tags: duplicate "sql"`,
			},
		},
		{
			Name: "Errors #2",
			Services: map[string]input.Service{
				"db": {
					Args: []any{"localhost", 3306},
				},
				"logger": {},
				"server": {
					Constructor: ptr.New("pkg.NewServer"),
					Value:       ptr.New("pkg.Server{}"),
				},
				"http-client": {
					Value:  ptr.New("pkgHttp.Client{}"),
					Args:   []any{time.Second * 30},
					Getter: ptr.New("Get_http-client"),
				},
			},
			Errors: []string{
				`services: "db": missing constructor or value or type`,
				`services: "db": arguments are not empty, but constructor is missing`,
				`services: "http-client": arguments are not empty, but constructor is missing`,
				`services: "http-client": getter: invalid "Get_http-client"`,
				`services: "logger": missing constructor or value or type`,
				`services: "server": cannot define constructor and value together`,
			},
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			i := input.Input{Services: s.Services}

			t.Run("input.ValidateServices", func(t *testing.T) {
				t.Parallel()
				err := input.ValidateServices(i)
				errAssert.EqualErrorGroup(t, err, s.Errors)
			})
			t.Run("input.NewDefaultValidator", func(t *testing.T) {
				t.Parallel()
				err := input.NewDefaultValidator("dev-main").Validate(i)
				errAssert.EqualErrorGroup(t, err, s.Errors)
			})
		})
	}
}

func TestDefaultServicesValidators(t *testing.T) {
	assert.NotEmpty(t, input.DefaultServicesValidators())
}
