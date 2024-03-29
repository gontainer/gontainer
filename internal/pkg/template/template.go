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

package template

import (
	"bytes"
	"io/fs"
	"text/template"

	"github.com/gontainer/gontainer-helpers/v3/exporter"
	"github.com/gontainer/gontainer/internal/pkg/consts"
	"github.com/gontainer/gontainer/internal/pkg/imports"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/template/templates"
)

type aliaser interface {
	Alias(string) string
}

type importProvider interface {
	Imports() []imports.Import
}

type codeFormatter interface {
	Format(string) (string, error)
}

type Builder struct {
	aliaser         aliaser
	importsProvider importProvider
	formatter       codeFormatter
	buildInfo       string
	stub            bool
}

func NewBuilder(a aliaser, ip importProvider, cf codeFormatter, buildInfo string, stub bool) *Builder {
	return &Builder{
		aliaser:         a,
		importsProvider: ip,
		formatter:       cf,
		buildInfo:       buildInfo,
		stub:            stub,
	}
}

func (b Builder) Build(o output.Output) (string, error) {
	d := data{
		ImportCollection: b.importsProvider,
		Output:           o,
		BuildInfo:        b.buildInfo,
		Stub:             b.stub,
	}

	funcs := createDefaultFunctions(b.aliaser, o)

	var (
		body, head string
		err        error
	)

	tplBody := tpl{
		data:     d,
		funcs:    funcs,
		fsys:     templates.Body,
		name:     "body.go.tpl",
		patterns: []string{"body.go.tpl", "body-*.go.tpl"},
	}

	tplHead := tpl{
		fsys:     templates.Head,
		data:     d,
		funcs:    funcs,
		name:     "head.go.tpl",
		patterns: []string{"head.go.tpl"},
	}

	if body, err = tplBody.exec(); err != nil {
		return "", err
	}

	// we have to execute that template as the last one
	// because the previous one can add imports,
	// and we need to print all of them
	if head, err = tplHead.exec(); err != nil {
		return "", err
	}

	return b.formatter.Format(head + body)
}

func createDefaultFunctions(a aliaser, o output.Output) template.FuncMap {
	tagsServices := make(map[string]map[string]struct{}) // tagsServices[tag][serviceID] = struct{}

	for _, s := range o.Services {
		for _, t := range s.Tags {
			if _, ok := tagsServices[t.Name]; !ok {
				tagsServices[t.Name] = make(map[string]struct{})
			}
			tagsServices[t.Name][s.Name] = struct{}{}
		}
	}

	return template.FuncMap{
		"export": func(input any) (string, error) {
			return exporter.Export(input)
		},
		"importAlias": func(i string) string {
			return a.Alias(i)
		},
		"containerAlias": func() string {
			return a.Alias(consts.GontainerHelperPath + "/container")
		},
		"groupErrorAlias": func() string {
			return a.Alias(consts.GontainerHelperPath + "/grouperror")
		},
		"exporterAlias": func() string {
			return a.Alias(consts.GontainerHelperPath + "/exporter")
		},
		"callerAlias": func() string {
			return a.Alias(consts.GontainerHelperPath + "/caller")
		},
		"copierAlias": func() string {
			return a.Alias(consts.GontainerHelperPath + "/copier")
		},
		"isTagged": func(id string, tag string) bool {
			_, ok := tagsServices[tag][id]
			return ok
		},
		"isString": func(i any) bool {
			_, ok := i.(string)
			return ok
		},
	}
}

type tpl struct {
	fsys     fs.FS
	name     string
	patterns []string
	data     any
	funcs    template.FuncMap
}

func (t tpl) exec() (string, error) {
	tpl, err := template.New(t.name).Funcs(t.funcs).ParseFS(t.fsys, t.patterns...)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, t.data)
	return b.String(), err
}
