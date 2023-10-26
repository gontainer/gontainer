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
				Scope: output.SetScopeDefault,
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
