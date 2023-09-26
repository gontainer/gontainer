package cmd

import (
	"io"

	"github.com/gontainer/gontainer/internal/cmd/runner"
	"github.com/gontainer/gontainer/internal/pkg/compiler"
	"github.com/gontainer/gontainer/internal/pkg/consts"
	"github.com/gontainer/gontainer/internal/pkg/imports"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/resolver"
	"github.com/gontainer/gontainer/internal/pkg/template"
	"github.com/gontainer/gontainer/internal/pkg/token"
)

func buildRunner(
	w io.Writer,
	version string,
	buildInfo string,
	paramsExistActive bool,
	servicesExistActive bool,
	inputPatterns []string,
	outputFile string,
) *runner.Runner {
	printer := runner.NewPrinter(w)
	tokenStrategyFactory := token.NewStrategyFactory(
		token.FactoryPercentMark{},
		token.FactoryReference{},
		token.FactoryUnexpectedFunction{},
		token.FactoryUnexpectedToken{},
		token.FactoryString{},
	)
	tokenizer := token.NewTokenizer(token.NewChunker(), tokenStrategyFactory)
	imports_ := imports.New()
	fnRegisterer := token.NewFuncRegisterer(tokenStrategyFactory, imports_)
	paramResolver := resolver.NewParamResolver(resolver.NewArgResolver(
		resolver.NewNonStringPrimitiveResolver(),
		resolver.NewPatternResolver(tokenizer),
	))
	argResolver := resolver.NewArgResolver(
		resolver.NewNonStringPrimitiveResolver(),
		resolver.NewValueResolver(imports_),
		resolver.NewServiceResolver(),
		resolver.NewTaggedResolver(),
		resolver.NewFixedValueResolver(consts.SpecialGontainerID, consts.SpecialGontainerValue),
		resolver.NewPatternResolver(tokenizer),
	)
	compiler_ := compiler.New(
		compiler.NewStepValidateInput(input.NewDefaultValidator(version)),
		compiler.NewStepCompileMeta(imports_, fnRegisterer),
		compiler.NewStepCompileParams(paramResolver),
		compiler.NewStepCompileServices(imports_, argResolver),
		compiler.NewStepCompileDecorators(imports_, argResolver),
	)
	templateBuilder := template.NewBuilder(
		imports_,
		imports_,
		template.NewCodeFormatter(),
		template.WithBuildInfo(buildInfo),
	)
	v := func(s runner.Step) *runner.StepVerboseSwitchable {
		return runner.NewStepVerboseSwitchable(s, printer, printer)
	}
	stepServicesScopes := v(runner.NewStepOutputValidationRule(output.ValidateServicesScopes, "Scope"))
	stepParamsCircular := v(runner.NewStepOutputValidationRule(output.ValidateParamsCircularDeps, "Circular deps in params"))
	stepServicesCircular := v(runner.NewStepOutputValidationRule(output.ValidateServicesCircularDeps, "Circular deps in services"))
	stepParamsExist := v(runner.NewStepOutputValidationRule(output.ValidateParamsExist, "Missing parameters"))
	stepServicesExist := v(runner.NewStepOutputValidationRule(output.ValidateServicesExist, "Missing services"))
	stepParamsExist.Active(paramsExistActive)
	stepServicesExist.Active(servicesExistActive)

	return runner.NewRunner(
		v(runner.StepDefaultInput{}),
		v(runner.NewStepReadConfig(printer, inputPatterns)),
		v(runner.NewStepCompile(compiler_)),
		v(runner.NewStepAmalgamated(
			"Validate output",
			stepServicesScopes,
			stepParamsCircular,
			stepServicesCircular,
			stepParamsExist,
			stepServicesExist,
		)),
		v(runner.NewStepCodeGenerator(printer, templateBuilder, outputFile)),
	)
}
