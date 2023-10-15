// Code generated by https://github.com/gontainer/gontainer; DO NOT EDIT.

package gontainer

// gontainer version: dev-gontainer-helpers@1.4 033aeed4f4c9521236c17b3e51b2f6f9e52b51d4-dirty (build date 2023-10-15T17:43:11Z)

import (
	ie_context "context"
	if_errors "errors"
	i1_fmt "fmt"
	i11_os "os"
	i12_strconv "strconv"

	i9_runner "github.com/gontainer/gontainer/internal/cmd/runner"
	i4_compiler "github.com/gontainer/gontainer/internal/pkg/compiler"
	i6_consts "github.com/gontainer/gontainer/internal/pkg/consts"
	i7_imports "github.com/gontainer/gontainer/internal/pkg/imports"
	i8_input "github.com/gontainer/gontainer/internal/pkg/input"
	ia_output "github.com/gontainer/gontainer/internal/pkg/output"
	i2_resolver "github.com/gontainer/gontainer/internal/pkg/resolver"
	i3_template "github.com/gontainer/gontainer/internal/pkg/template"
	i5_token "github.com/gontainer/gontainer/internal/pkg/token"

	i0_caller "github.com/gontainer/gontainer-helpers/caller"
	ib_container "github.com/gontainer/gontainer-helpers/container"
	id_copier "github.com/gontainer/gontainer-helpers/copier"
	ic_errors "github.com/gontainer/gontainer-helpers/errors"
	i10_exporter "github.com/gontainer/gontainer-helpers/exporter"
)

// ············································································
// ···································PARAMS···································
// ············································································
// #### buildInfo
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = i0_caller.CallProvider(paramTodo); if err != nil { err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································
// #### inputPatterns
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = i0_caller.CallProvider(paramTodo); if err != nil { err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································
// #### outputFile
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = i0_caller.CallProvider(paramTodo); if err != nil { err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································
// #### stub
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = i0_caller.CallProvider(paramTodo); if err != nil { err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································
// #### version
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = i0_caller.CallProvider(paramTodo); if err != nil { err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································

// ············································································
// ··································SERVICES··································
// ············································································
// #### argResolver
// var service interface{}
// service = i2_resolver.NewArgResolver(eval("@nonStringPrimitiveResolver"), eval("@valueResolver"), eval("@serviceResolver"), eval("@taggedResolver"), eval("@gontainerValueResolver"), eval("@patternResolver"))
// ············································································
// #### codeFormatter
// var service interface{}
// service = i3_template.NewCodeFormatter()
// ············································································
// #### compiler
// var service interface{}
// service = i4_compiler.New(eval("@stepValidateInput"), eval("@stepCompileMeta"), eval("@stepCompileParams"), eval("@stepCompileServices"), eval("@stepCompileDecorators"))
// ············································································
// #### fnRegisterer
// var service interface{}
// service = i5_token.NewFuncRegisterer(eval("@tokenStrategyFactory"), eval("@imports"))
// ············································································
// #### gontainerValueResolver
// var service interface{}
// service = i2_resolver.NewFixedValueResolver(eval("!value consts.SpecialGontainerID"), eval("!value consts.SpecialGontainerValue"))
// ············································································
// #### imports
// var service interface{}
// service = i7_imports.New()
// ············································································
// #### inputValidator
// var service interface{}
// service = i8_input.NewDefaultValidator(eval("%version%"))
// ············································································
// #### nonStringPrimitiveResolver
// var service interface{}
// service = i2_resolver.NewNonStringPrimitiveResolver()
// ············································································
// #### paramResolver
// var service interface{}
// service = i2_resolver.NewParamResolver(eval("@primitiveArgResolver"))
// ············································································
// #### patternResolver
// var service interface{}
// service = i2_resolver.NewPatternResolver(eval("@tokenizer"))
// ············································································
// #### primitiveArgResolver
// var service interface{}
// service = i2_resolver.NewArgResolver(eval("@nonStringPrimitiveResolver"), eval("@patternResolver"))
// ············································································
// #### printer
// var service interface{}
// service = i9_runner.NewPrinter(eval("@writer"))
// ············································································
// #### runner
// var service *i9_runner.Runner
// service = i9_runner.NewRunner(eval("@stepDefaultInput"), eval("@stepReadConfig"), eval("@stepCompile"), eval("@stepValidateOutput"), eval("@stepCodeGenerator"))
// ············································································
// #### serviceResolver
// var service interface{}
// service = i2_resolver.NewServiceResolver()
// ············································································
// #### stepCodeGenerator
// var service interface{}
// service = i9_runner.NewStepCodeGenerator(eval("@printer"), eval("@templateBuilder"), eval("%outputFile%"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepCodeGenerator", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepCompile
// var service interface{}
// service = i9_runner.NewStepCompile(eval("@compiler"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepCompile", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepCompileDecorators
// var service interface{}
// service = i4_compiler.NewStepCompileDecorators(eval("@imports"), eval("@argResolver"))
// ············································································
// #### stepCompileMeta
// var service interface{}
// service = i4_compiler.NewStepCompileMeta(eval("@imports"), eval("@fnRegisterer"))
// ············································································
// #### stepCompileParams
// var service interface{}
// service = i4_compiler.NewStepCompileParams(eval("@paramResolver"))
// ············································································
// #### stepCompileServices
// var service interface{}
// service = i4_compiler.NewStepCompileServices(eval("@imports"), eval("@argResolver"))
// ············································································
// #### stepDefaultInput
// service := i9_runner.StepDefaultInput{}
// service = i9_runner.DecorateStepVerboseSwitchable("stepDefaultInput", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputParamsCircular
// var service interface{}
// service = i9_runner.NewStepOutputValidationRule(eval("!value output.ValidateParamsCircularDeps"), eval("Circular deps in params"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepOutputParamsCircular", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputParamsExist
// var service *i9_runner.StepVerboseSwitchable
// service = i9_runner.NewStepOutputValidationRule(eval("!value output.ValidateParamsExist"), eval("Missing parameters"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepOutputParamsExist", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputServicesCircular
// var service interface{}
// service = i9_runner.NewStepOutputValidationRule(eval("!value output.ValidateServicesCircularDeps"), eval("Circular deps in services"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepOutputServicesCircular", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputServicesExist
// var service *i9_runner.StepVerboseSwitchable
// service = i9_runner.NewStepOutputValidationRule(eval("!value output.ValidateServicesExist"), eval("Missing services"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepOutputServicesExist", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputServicesScopes
// var service interface{}
// service = i9_runner.NewStepOutputValidationRule(eval("!value output.ValidateServicesScopes"), eval("Scope"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepOutputServicesScopes", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepReadConfig
// var service interface{}
// service = i9_runner.NewStepReadConfig(eval("@printer"), eval("%inputPatterns%"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepReadConfig", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepValidateInput
// var service interface{}
// service = i4_compiler.NewStepValidateInput(eval("@inputValidator"))
// ············································································
// #### stepValidateOutput
// var service interface{}
// service = i9_runner.NewStepAmalgamated(eval("Validate output"), eval("@stepOutputServicesScopes"), eval("@stepOutputParamsCircular"), eval("@stepOutputServicesCircular"), eval("@stepOutputParamsExist"), eval("@stepOutputServicesExist"))
// service = i9_runner.DecorateStepVerboseSwitchable("stepValidateOutput", service, eval("@printer"), eval("@printer"))
// ············································································
// #### taggedResolver
// var service interface{}
// service = i2_resolver.NewTaggedResolver()
// ············································································
// #### templateBuilder
// var service interface{}
// service = i3_template.NewBuilder(eval("@imports"), eval("@imports"), eval("@codeFormatter"), eval("%buildInfo%"), eval("%stub%"))
// ············································································
// #### tokenChunker
// var service interface{}
// service = i5_token.NewChunker()
// ············································································
// #### tokenStrategyFactory
// var service interface{}
// service = i5_token.NewStrategyFactory(eval("!value token.FactoryPercentMark{}"), eval("!value token.FactoryReference{}"), eval("!value token.FactoryUnexpectedFunction{}"), eval("!value token.FactoryUnexpectedToken{}"), eval("!value token.FactoryString{}"))
// ············································································
// #### tokenizer
// var service interface{}
// service = i5_token.NewTokenizer(eval("@tokenChunker"), eval("@tokenStrategyFactory"))
// ············································································
// #### valueResolver
// var service interface{}
// service = i2_resolver.NewValueResolver(eval("@imports"))
// ············································································
// #### writer
// panic("todo")
// ············································································

type gontainer struct {
	*ib_container.SuperContainer
}

func (c *gontainer) GetRunner() (result *i9_runner.Runner, err error) {
	var s interface{}
	s, err = c.Get("runner")
	if err != nil {
		return nil, ic_errors.PrefixedGroup(
			i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetRunner"),
			err,
		)
	}
	err = ic_errors.PrefixedGroup(
		i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetRunner"),
		id_copier.ConvertAndCopy(s, &result),
	)
	return
}

func (c *gontainer) GetRunnerContext(ctx ie_context.Context) (result *i9_runner.Runner, err error) {
	var s interface{}
	s, err = c.GetWithContext(ctx, "runner")
	if err != nil {
		return nil, ic_errors.PrefixedGroup(
			i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetRunner"),
			err,
		)
	}
	err = ic_errors.PrefixedGroup(
		i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetRunner"),
		id_copier.ConvertAndCopy(s, &result),
	)
	return
}

func (c *gontainer) MustGetRunner() *i9_runner.Runner {
	r, err := c.GetRunner()
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) MustGetRunnerContext(ctx ie_context.Context) *i9_runner.Runner {
	r, err := c.GetRunnerContext(ctx)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) GetStepValidateParamsExist() (result *i9_runner.StepVerboseSwitchable, err error) {
	var s interface{}
	s, err = c.Get("stepOutputParamsExist")
	if err != nil {
		return nil, ic_errors.PrefixedGroup(
			i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateParamsExist"),
			err,
		)
	}
	err = ic_errors.PrefixedGroup(
		i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateParamsExist"),
		id_copier.ConvertAndCopy(s, &result),
	)
	return
}

func (c *gontainer) GetStepValidateParamsExistContext(ctx ie_context.Context) (result *i9_runner.StepVerboseSwitchable, err error) {
	var s interface{}
	s, err = c.GetWithContext(ctx, "stepOutputParamsExist")
	if err != nil {
		return nil, ic_errors.PrefixedGroup(
			i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateParamsExist"),
			err,
		)
	}
	err = ic_errors.PrefixedGroup(
		i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateParamsExist"),
		id_copier.ConvertAndCopy(s, &result),
	)
	return
}

func (c *gontainer) MustGetStepValidateParamsExist() *i9_runner.StepVerboseSwitchable {
	r, err := c.GetStepValidateParamsExist()
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) MustGetStepValidateParamsExistContext(ctx ie_context.Context) *i9_runner.StepVerboseSwitchable {
	r, err := c.GetStepValidateParamsExistContext(ctx)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) GetStepValidateServicesExist() (result *i9_runner.StepVerboseSwitchable, err error) {
	var s interface{}
	s, err = c.Get("stepOutputServicesExist")
	if err != nil {
		return nil, ic_errors.PrefixedGroup(
			i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateServicesExist"),
			err,
		)
	}
	err = ic_errors.PrefixedGroup(
		i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateServicesExist"),
		id_copier.ConvertAndCopy(s, &result),
	)
	return
}

func (c *gontainer) GetStepValidateServicesExistContext(ctx ie_context.Context) (result *i9_runner.StepVerboseSwitchable, err error) {
	var s interface{}
	s, err = c.GetWithContext(ctx, "stepOutputServicesExist")
	if err != nil {
		return nil, ic_errors.PrefixedGroup(
			i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateServicesExist"),
			err,
		)
	}
	err = ic_errors.PrefixedGroup(
		i1_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateServicesExist"),
		id_copier.ConvertAndCopy(s, &result),
	)
	return
}

func (c *gontainer) MustGetStepValidateServicesExist() *i9_runner.StepVerboseSwitchable {
	r, err := c.GetStepValidateServicesExist()
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) MustGetStepValidateServicesExistContext(ctx ie_context.Context) *i9_runner.StepVerboseSwitchable {
	r, err := c.GetStepValidateServicesExistContext(ctx)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func New() (rootGontainer interface {
	// service container
	Get(serviceID string) (interface{}, error)
	GetWithContext(ctx ie_context.Context, serviceID string) (interface{}, error)
	CircularDeps() error
	OverrideService(serviceID string, s ib_container.Service)
	AddDecorator(tag string, decorator interface{}, deps ...ib_container.Dependency)
	IsTaggedBy(serviceID string, tag string) bool
	GetTaggedBy(tag string) ([]interface{}, error)
	GetTaggedByWithContext(ctx ie_context.Context, tag string) ([]interface{}, error)
	// param container
	GetParam(paramID string) (interface{}, error)
	OverrideParam(paramID string, d ib_container.Dependency)
	// getters
	GetRunner() (*i9_runner.Runner, error)
	GetRunnerContext(ctx ie_context.Context) (*i9_runner.Runner, error)
	MustGetRunner() *i9_runner.Runner
	GetStepValidateParamsExist() (*i9_runner.StepVerboseSwitchable, error)
	GetStepValidateParamsExistContext(ctx ie_context.Context) (*i9_runner.StepVerboseSwitchable, error)
	MustGetStepValidateParamsExist() *i9_runner.StepVerboseSwitchable
	GetStepValidateServicesExist() (*i9_runner.StepVerboseSwitchable, error)
	GetStepValidateServicesExistContext(ctx ie_context.Context) (*i9_runner.StepVerboseSwitchable, error)
	MustGetStepValidateServicesExist() *i9_runner.StepVerboseSwitchable
}) {
	sc := &gontainer{
		SuperContainer: ib_container.NewSuperContainer(),
	}
	rootGontainer = sc
	//
	//
	// #####################################
	// ############## Helpers ##############
	//
	//
	dependencyService := ib_container.NewDependencyService
	_ = dependencyService
	dependencyValue := ib_container.NewDependencyValue
	_ = dependencyValue
	dependencyTag := ib_container.NewDependencyTag
	_ = dependencyTag
	dependencyProvider := ib_container.NewDependencyProvider
	_ = dependencyProvider
	newService := ib_container.NewService
	_ = newService
	concatenateChunks := sc._concatenateChunks
	_ = concatenateChunks
	paramTodo := sc._paramTodo
	_ = paramTodo
	getEnv := sc._getEnv
	_ = getEnv
	getEnvInt := sc._getEnvInt
	_ = getEnvInt
	getParam := sc.GetParam
	_ = getParam
	//
	//
	// ####################################
	// ############## Params ##############
	//
	//
	// "buildInfo": "%todo()%"
	sc.OverrideParam("buildInfo", dependencyProvider(func() (r interface{}, err error) {
		r, err = i0_caller.CallProvider(paramTodo)
		if err != nil {
			err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	// "inputPatterns": "%todo()%"
	sc.OverrideParam("inputPatterns", dependencyProvider(func() (r interface{}, err error) {
		r, err = i0_caller.CallProvider(paramTodo)
		if err != nil {
			err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	// "outputFile": "%todo()%"
	sc.OverrideParam("outputFile", dependencyProvider(func() (r interface{}, err error) {
		r, err = i0_caller.CallProvider(paramTodo)
		if err != nil {
			err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	// "stub": "%todo()%"
	sc.OverrideParam("stub", dependencyProvider(func() (r interface{}, err error) {
		r, err = i0_caller.CallProvider(paramTodo)
		if err != nil {
			err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	// "version": "%todo()%"
	sc.OverrideParam("version", dependencyProvider(func() (r interface{}, err error) {
		r, err = i0_caller.CallProvider(paramTodo)
		if err != nil {
			err = i1_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	//
	//
	// ######################################
	// ############## Services ##############
	//
	//
	// "argResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewArgResolver,
			// "@nonStringPrimitiveResolver"
			dependencyService("nonStringPrimitiveResolver"),
			// "@valueResolver"
			dependencyService("valueResolver"),
			// "@serviceResolver"
			dependencyService("serviceResolver"),
			// "@taggedResolver"
			dependencyService("taggedResolver"),
			// "@gontainerValueResolver"
			dependencyService("gontainerValueResolver"),
			// "@patternResolver"
			dependencyService("patternResolver"),
		)
		s.ScopeDefault()
		sc.OverrideService("argResolver", s)
	}
	// "codeFormatter"
	{
		s := newService()
		s.SetConstructor(
			i3_template.NewCodeFormatter,
		)
		s.ScopeDefault()
		sc.OverrideService("codeFormatter", s)
	}
	// "compiler"
	{
		s := newService()
		s.SetConstructor(
			i4_compiler.New,
			// "@stepValidateInput"
			dependencyService("stepValidateInput"),
			// "@stepCompileMeta"
			dependencyService("stepCompileMeta"),
			// "@stepCompileParams"
			dependencyService("stepCompileParams"),
			// "@stepCompileServices"
			dependencyService("stepCompileServices"),
			// "@stepCompileDecorators"
			dependencyService("stepCompileDecorators"),
		)
		s.ScopeDefault()
		sc.OverrideService("compiler", s)
	}
	// "fnRegisterer"
	{
		s := newService()
		s.SetConstructor(
			i5_token.NewFuncRegisterer,
			// "@tokenStrategyFactory"
			dependencyService("tokenStrategyFactory"),
			// "@imports"
			dependencyService("imports"),
		)
		s.ScopeDefault()
		sc.OverrideService("fnRegisterer", s)
	}
	// "gontainerValueResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewFixedValueResolver,
			// "!value consts.SpecialGontainerID"
			dependencyValue(i6_consts.SpecialGontainerID),
			// "!value consts.SpecialGontainerValue"
			dependencyValue(i6_consts.SpecialGontainerValue),
		)
		s.ScopeDefault()
		sc.OverrideService("gontainerValueResolver", s)
	}
	// "imports"
	{
		s := newService()
		s.SetConstructor(
			i7_imports.New,
		)
		s.ScopeDefault()
		sc.OverrideService("imports", s)
	}
	// "inputValidator"
	{
		s := newService()
		s.SetConstructor(
			i8_input.NewDefaultValidator,
			// "%version%"
			dependencyProvider(func() (interface{}, error) { return getParam("version") }),
		)
		s.ScopeDefault()
		sc.OverrideService("inputValidator", s)
	}
	// "nonStringPrimitiveResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewNonStringPrimitiveResolver,
		)
		s.ScopeDefault()
		sc.OverrideService("nonStringPrimitiveResolver", s)
	}
	// "paramResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewParamResolver,
			// "@primitiveArgResolver"
			dependencyService("primitiveArgResolver"),
		)
		s.ScopeDefault()
		sc.OverrideService("paramResolver", s)
	}
	// "patternResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewPatternResolver,
			// "@tokenizer"
			dependencyService("tokenizer"),
		)
		s.ScopeDefault()
		sc.OverrideService("patternResolver", s)
	}
	// "primitiveArgResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewArgResolver,
			// "@nonStringPrimitiveResolver"
			dependencyService("nonStringPrimitiveResolver"),
			// "@patternResolver"
			dependencyService("patternResolver"),
		)
		s.ScopeDefault()
		sc.OverrideService("primitiveArgResolver", s)
	}
	// "printer"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewPrinter,
			// "@writer"
			dependencyService("writer"),
		)
		s.ScopeDefault()
		sc.OverrideService("printer", s)
	}
	// "runner"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewRunner,
			// "@stepDefaultInput"
			dependencyService("stepDefaultInput"),
			// "@stepReadConfig"
			dependencyService("stepReadConfig"),
			// "@stepCompile"
			dependencyService("stepCompile"),
			// "@stepValidateOutput"
			dependencyService("stepValidateOutput"),
			// "@stepCodeGenerator"
			dependencyService("stepCodeGenerator"),
		)
		s.ScopeDefault()
		sc.OverrideService("runner", s)
	}
	// "serviceResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewServiceResolver,
		)
		s.ScopeDefault()
		sc.OverrideService("serviceResolver", s)
	}
	// "stepCodeGenerator"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepCodeGenerator,
			// "@printer"
			dependencyService("printer"),
			// "@templateBuilder"
			dependencyService("templateBuilder"),
			// "%outputFile%"
			dependencyProvider(func() (interface{}, error) { return getParam("outputFile") }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepCodeGenerator", s)
	}
	// "stepCompile"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepCompile,
			// "@compiler"
			dependencyService("compiler"),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepCompile", s)
	}
	// "stepCompileDecorators"
	{
		s := newService()
		s.SetConstructor(
			i4_compiler.NewStepCompileDecorators,
			// "@imports"
			dependencyService("imports"),
			// "@argResolver"
			dependencyService("argResolver"),
		)
		s.ScopeDefault()
		sc.OverrideService("stepCompileDecorators", s)
	}
	// "stepCompileMeta"
	{
		s := newService()
		s.SetConstructor(
			i4_compiler.NewStepCompileMeta,
			// "@imports"
			dependencyService("imports"),
			// "@fnRegisterer"
			dependencyService("fnRegisterer"),
		)
		s.ScopeDefault()
		sc.OverrideService("stepCompileMeta", s)
	}
	// "stepCompileParams"
	{
		s := newService()
		s.SetConstructor(
			i4_compiler.NewStepCompileParams,
			// "@paramResolver"
			dependencyService("paramResolver"),
		)
		s.ScopeDefault()
		sc.OverrideService("stepCompileParams", s)
	}
	// "stepCompileServices"
	{
		s := newService()
		s.SetConstructor(
			i4_compiler.NewStepCompileServices,
			// "@imports"
			dependencyService("imports"),
			// "@argResolver"
			dependencyService("argResolver"),
		)
		s.ScopeDefault()
		sc.OverrideService("stepCompileServices", s)
	}
	// "stepDefaultInput"
	{
		s := newService()
		s.SetConstructor(func() interface{} { return i9_runner.StepDefaultInput{} })
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepDefaultInput", s)
	}
	// "stepOutputParamsCircular"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepOutputValidationRule,
			// "!value output.ValidateParamsCircularDeps"
			dependencyValue(ia_output.ValidateParamsCircularDeps),
			// "Circular deps in params"
			dependencyProvider(func() (r interface{}, err error) { return "Circular deps in params", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepOutputParamsCircular", s)
	}
	// "stepOutputParamsExist"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepOutputValidationRule,
			// "!value output.ValidateParamsExist"
			dependencyValue(ia_output.ValidateParamsExist),
			// "Missing parameters"
			dependencyProvider(func() (r interface{}, err error) { return "Missing parameters", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepOutputParamsExist", s)
	}
	// "stepOutputServicesCircular"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepOutputValidationRule,
			// "!value output.ValidateServicesCircularDeps"
			dependencyValue(ia_output.ValidateServicesCircularDeps),
			// "Circular deps in services"
			dependencyProvider(func() (r interface{}, err error) { return "Circular deps in services", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepOutputServicesCircular", s)
	}
	// "stepOutputServicesExist"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepOutputValidationRule,
			// "!value output.ValidateServicesExist"
			dependencyValue(ia_output.ValidateServicesExist),
			// "Missing services"
			dependencyProvider(func() (r interface{}, err error) { return "Missing services", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepOutputServicesExist", s)
	}
	// "stepOutputServicesScopes"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepOutputValidationRule,
			// "!value output.ValidateServicesScopes"
			dependencyValue(ia_output.ValidateServicesScopes),
			// "Scope"
			dependencyProvider(func() (r interface{}, err error) { return "Scope", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepOutputServicesScopes", s)
	}
	// "stepReadConfig"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepReadConfig,
			// "@printer"
			dependencyService("printer"),
			// "%inputPatterns%"
			dependencyProvider(func() (interface{}, error) { return getParam("inputPatterns") }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepReadConfig", s)
	}
	// "stepValidateInput"
	{
		s := newService()
		s.SetConstructor(
			i4_compiler.NewStepValidateInput,
			// "@inputValidator"
			dependencyService("inputValidator"),
		)
		s.ScopeDefault()
		sc.OverrideService("stepValidateInput", s)
	}
	// "stepValidateOutput"
	{
		s := newService()
		s.SetConstructor(
			i9_runner.NewStepAmalgamated,
			// "Validate output"
			dependencyProvider(func() (r interface{}, err error) { return "Validate output", nil }),
			// "@stepOutputServicesScopes"
			dependencyService("stepOutputServicesScopes"),
			// "@stepOutputParamsCircular"
			dependencyService("stepOutputParamsCircular"),
			// "@stepOutputServicesCircular"
			dependencyService("stepOutputServicesCircular"),
			// "@stepOutputParamsExist"
			dependencyService("stepOutputParamsExist"),
			// "@stepOutputServicesExist"
			dependencyService("stepOutputServicesExist"),
		)
		s.Tag("step-runner-verbose", int(0))
		s.ScopeDefault()
		sc.OverrideService("stepValidateOutput", s)
	}
	// "taggedResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewTaggedResolver,
		)
		s.ScopeDefault()
		sc.OverrideService("taggedResolver", s)
	}
	// "templateBuilder"
	{
		s := newService()
		s.SetConstructor(
			i3_template.NewBuilder,
			// "@imports"
			dependencyService("imports"),
			// "@imports"
			dependencyService("imports"),
			// "@codeFormatter"
			dependencyService("codeFormatter"),
			// "%buildInfo%"
			dependencyProvider(func() (interface{}, error) { return getParam("buildInfo") }),
			// "%stub%"
			dependencyProvider(func() (interface{}, error) { return getParam("stub") }),
		)
		s.ScopeDefault()
		sc.OverrideService("templateBuilder", s)
	}
	// "tokenChunker"
	{
		s := newService()
		s.SetConstructor(
			i5_token.NewChunker,
		)
		s.ScopeDefault()
		sc.OverrideService("tokenChunker", s)
	}
	// "tokenStrategyFactory"
	{
		s := newService()
		s.SetConstructor(
			i5_token.NewStrategyFactory,
			// "!value token.FactoryPercentMark{}"
			dependencyValue(i5_token.FactoryPercentMark{}),
			// "!value token.FactoryReference{}"
			dependencyValue(i5_token.FactoryReference{}),
			// "!value token.FactoryUnexpectedFunction{}"
			dependencyValue(i5_token.FactoryUnexpectedFunction{}),
			// "!value token.FactoryUnexpectedToken{}"
			dependencyValue(i5_token.FactoryUnexpectedToken{}),
			// "!value token.FactoryString{}"
			dependencyValue(i5_token.FactoryString{}),
		)
		s.ScopeDefault()
		sc.OverrideService("tokenStrategyFactory", s)
	}
	// "tokenizer"
	{
		s := newService()
		s.SetConstructor(
			i5_token.NewTokenizer,
			// "@tokenChunker"
			dependencyService("tokenChunker"),
			// "@tokenStrategyFactory"
			dependencyService("tokenStrategyFactory"),
		)
		s.ScopeDefault()
		sc.OverrideService("tokenizer", s)
	}
	// "valueResolver"
	{
		s := newService()
		s.SetConstructor(
			i2_resolver.NewValueResolver,
			// "@imports"
			dependencyService("imports"),
		)
		s.ScopeDefault()
		sc.OverrideService("valueResolver", s)
	}
	// "writer"
	{
		s := newService()
		s.SetConstructor(func() (interface{}, error) { return nil, if_errors.New("service todo") })
		sc.OverrideService("writer", s)
	}
	//
	//
	// ########################################
	// ############## Decorators ##############
	//
	//
	sc.AddDecorator(
		"step-runner-verbose",
		i9_runner.DecorateStepVerboseSwitchable,
		dependencyService("printer"),
		dependencyService("printer"),
	)
	return
}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *gontainer) _concatenateChunks(first func() (interface{}, error), chunks ...func() (interface{}, error)) (string, error) {
	r := ""
	for _, p := range append([]func() (interface{}, error){first}, chunks...) {
		chunk, err := p()
		if err != nil {
			return "", err
		}
		s, err := i10_exporter.CastToString(chunk)
		if err != nil {
			return "", err
		}
		r += s
	}
	return r, nil
}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *gontainer) _paramTodo(params ...string) (interface{}, error) {
	if len(params) > 0 {
		return nil, if_errors.New(params[0])
	}
	return nil, if_errors.New("parameter todo")
}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *gontainer) _getEnv(key string, def ...string) (string, error) {
	val, ok := i11_os.LookupEnv(key)
	if !ok {
		if len(def) > 0 {
			return def[0], nil
		}
		return "", i1_fmt.Errorf("environment variable %+q does not exist", key)
	}
	return val, nil
}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *gontainer) _getEnvInt(key string, def ...int) (int, error) {
	val, ok := i11_os.LookupEnv(key)
	if !ok {
		if len(def) > 0 {
			return def[0], nil
		}
		return 0, i1_fmt.Errorf("environment variable %+q does not exist", key)
	}
	res, err := i12_strconv.Atoi(val)
	if err != nil {
		return 0, i1_fmt.Errorf("cannot cast env(%+q) to int: %w", key, err)
	}
	return res, nil
}
