package compiler

import (
	"github.com/gontainer/gontainer-helpers/grouperror"
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
