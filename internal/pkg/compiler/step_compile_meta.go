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

package compiler

import (
	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/maps"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/gontainer/gontainer/internal/pkg/regex"
	"github.com/gontainer/gontainer/internal/pkg/syntax"
)

var (
	regexMetaGoFn = regex.MustCompileAz(regex.MetaGoFn)
)

type StepCompileMeta struct {
	aliasRegisterer aliasRegisterer
	funcRegisterer  funcRegisterer
}

func NewStepCompileMeta(a aliasRegisterer, fn funcRegisterer) *StepCompileMeta {
	return &StepCompileMeta{
		aliasRegisterer: a,
		funcRegisterer:  fn,
	}
}

func (s *StepCompileMeta) Process(i input.Input, d *output.Output) error {
	var errs []error
	d.Meta.Pkg = ptr.Dereference(i.Meta.Pkg, defaultMetaPkg)
	d.Meta.ContainerType = ptr.Dereference(i.Meta.ContainerType, defaultMetaContainerType)
	d.Meta.ContainerConstructor = ptr.Dereference(i.Meta.ContainerConstructor, defaultMetaContainerConstructor)
	errs = append(errs, s.handleImports(i.Meta.Imports))
	s.handleFunctions(i.Meta.Functions)
	return grouperror.Prefix("compiler.StepCompileMeta: ", errs...)
}

func (s *StepCompileMeta) handleImports(imports map[string]string) error {
	var errs []error
	maps.Iterate(imports, func(alias string, import_ string) {
		errs = append(errs, s.aliasRegisterer.RegisterPrefixAlias(alias, import_))
	})
	return grouperror.Prefix("imports: ", errs...)
}

func (s *StepCompileMeta) handleFunctions(fns map[string]string) {
	maps.Iterate(fns, func(alias string, goFn string) {
		_, m := regex.Match(regexMetaGoFn, goFn)
		s.funcRegisterer.RegisterFunc(
			alias,
			syntax.SanitizeImport(m["import"]),
			m["fn"],
		)
	})
}
