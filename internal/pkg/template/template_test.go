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

package template_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gontainer/gontainer/internal/pkg/imports"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/template"
	"github.com/stretchr/testify/assert"
)

type aliaser struct {
}

func (a aliaser) Alias(string) string {
	return "pkg"
}

type importProvider struct {
	imports []imports.Import
}

func (ip importProvider) Imports() []imports.Import {
	return ip.imports
}

func TestBuilder_Build(t *testing.T) {
	ver := fmt.Sprintf("v99.99.99 %s", time.Now())
	b := template.NewBuilder(
		aliaser{},
		importProvider{
			imports: []imports.Import{
				{
					Alias: "myPkg",
					Path:  "pkg",
				},
			},
		},
		template.NewCodeFormatter(),
		ver,
		false,
	)
	code, err := b.Build(output.Output{
		Meta: output.Meta{
			Pkg:                  "main",
			ContainerType:        "Gontainer",
			ContainerConstructor: "NewGontainer",
		},
		Services: []output.Service{
			{
				Name:        "db",
				Getter:      "DB",
				MustGetter:  false,
				Type:        "*pkg.DB",
				Value:       "",
				Constructor: "pkg.NewDB",
				Args: []output.Arg{
					{Raw: "localhost", Code: `func () (interface{}, error) { return "localhost", nil}()`},
				},
				Calls: []output.Call{
					{Method: `Ping`},
				},
				Fields: []output.Field{
					{Name: "Logger", Value: output.Arg{Code: `container.Logger()`, Raw: `@logger`}},
				},
				Tags:  []output.Tag{{Name: "Traceable"}},
				Scope: output.ScopeDefault,
				Todo:  false,
			},
		},
		Decorators: []output.Decorator{
			{
				Tag:       "http-handler",
				Decorator: "http.TraceableHandler",
			},
		},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, code)
	assert.Contains(t, code, ver)
}
