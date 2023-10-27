// Code generated by https://github.com/gontainer/gontainer; DO NOT EDIT.

package gontainer

// gontainer version: dev-gontainer-helpers@v2.0.0-alpha 2374796403abc52bf74bc1e44e94435b5f5b76a7-dirty (build date 2023-10-27T17:45:03Z)

import (
	ib_context "context"
	if_errors "errors"
	i0_fmt "fmt"
	i11_os "os"
	ic_reflect "reflect"
	i12_strconv "strconv"

	i8_runner "github.com/gontainer/gontainer/internal/cmd/runner"
	i3_compiler "github.com/gontainer/gontainer/internal/pkg/compiler"
	i5_consts "github.com/gontainer/gontainer/internal/pkg/consts"
	i6_imports "github.com/gontainer/gontainer/internal/pkg/imports"
	i7_input "github.com/gontainer/gontainer/internal/pkg/input"
	i9_output "github.com/gontainer/gontainer/internal/pkg/output"
	i1_resolver "github.com/gontainer/gontainer/internal/pkg/resolver"
	i2_template "github.com/gontainer/gontainer/internal/pkg/template"
	i4_token "github.com/gontainer/gontainer/internal/pkg/token"

	i13_caller "github.com/gontainer/gontainer-helpers/v2/caller"
	ia_container "github.com/gontainer/gontainer-helpers/v2/container"
	ie_copier "github.com/gontainer/gontainer-helpers/v2/copier"
	i10_exporter "github.com/gontainer/gontainer-helpers/v2/exporter"
	id_grouperror "github.com/gontainer/gontainer-helpers/v2/grouperror"
)

// ············································································
// ···································PARAMS···································
// ············································································
// #### buildInfo
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = callProvider(paramTodo); if err != nil { err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································
// #### inputPatterns
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = callProvider(paramTodo); if err != nil { err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································
// #### outputFile
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = callProvider(paramTodo); if err != nil { err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································
// #### stub
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = callProvider(paramTodo); if err != nil { err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································
// #### version
// Raw: "%todo()%"
// GO:  dependencyProvider(func() (r interface{}, err error) { r, err = callProvider(paramTodo); if err != nil { err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err) }; return })
// ············································································

// ············································································
// ··································SERVICES··································
// ············································································
// #### argResolver
// var service interface{}
// service = i1_resolver.NewArgResolver(eval("@nonStringPrimitiveResolver"), eval("@valueResolver"), eval("@serviceResolver"), eval("@taggedResolver"), eval("@gontainerValueResolver"), eval("@patternResolver"))
// ············································································
// #### codeFormatter
// var service interface{}
// service = i2_template.NewCodeFormatter()
// ············································································
// #### compiler
// var service interface{}
// service = i3_compiler.New(eval("@stepValidateInput"), eval("@stepCompileMeta"), eval("@stepCompileParams"), eval("@stepCompileServices"), eval("@stepCompileDecorators"))
// ············································································
// #### fnRegisterer
// var service interface{}
// service = i4_token.NewFuncRegisterer(eval("@tokenStrategyFactory"), eval("@imports"))
// ············································································
// #### gontainerValueResolver
// var service interface{}
// service = i1_resolver.NewFixedValueResolver(eval("!value consts.SpecialGontainerID"), eval("!value consts.SpecialGontainerValue"))
// ············································································
// #### imports
// var service interface{}
// service = i6_imports.New()
// ············································································
// #### inputValidator
// var service interface{}
// service = i7_input.NewDefaultValidator(eval("%version%"))
// ············································································
// #### nonStringPrimitiveResolver
// var service interface{}
// service = i1_resolver.NewNonStringPrimitiveResolver()
// ············································································
// #### paramResolver
// var service interface{}
// service = i1_resolver.NewParamResolver(eval("@primitiveArgResolver"))
// ············································································
// #### patternResolver
// var service interface{}
// service = i1_resolver.NewPatternResolver(eval("@tokenizer"))
// ············································································
// #### primitiveArgResolver
// var service interface{}
// service = i1_resolver.NewArgResolver(eval("@nonStringPrimitiveResolver"), eval("@patternResolver"))
// ············································································
// #### printer
// var service interface{}
// service = i8_runner.NewPrinter(eval("@writer"))
// ············································································
// #### runner
// var service *i8_runner.Runner
// service = i8_runner.NewRunner(eval("@stepDefaultInput"), eval("@stepReadConfig"), eval("@stepCompile"), eval("@stepValidateOutput"), eval("@stepCodeGenerator"))
// ············································································
// #### serviceResolver
// var service interface{}
// service = i1_resolver.NewServiceResolver()
// ············································································
// #### stepCodeGenerator
// var service interface{}
// service = i8_runner.NewStepCodeGenerator(eval("@printer"), eval("@templateBuilder"), eval("%outputFile%"))
// service = i8_runner.DecorateStepVerboseSwitchable("stepCodeGenerator", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepCompile
// var service interface{}
// service = i8_runner.NewStepCompile(eval("@compiler"))
// service = i8_runner.DecorateStepVerboseSwitchable("stepCompile", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepCompileDecorators
// var service interface{}
// service = i3_compiler.NewStepCompileDecorators(eval("@imports"), eval("@argResolver"))
// ············································································
// #### stepCompileMeta
// var service interface{}
// service = i3_compiler.NewStepCompileMeta(eval("@imports"), eval("@fnRegisterer"))
// ············································································
// #### stepCompileParams
// var service interface{}
// service = i3_compiler.NewStepCompileParams(eval("@paramResolver"))
// ············································································
// #### stepCompileServices
// var service interface{}
// service = i3_compiler.NewStepCompileServices(eval("@imports"), eval("@argResolver"))
// ············································································
// #### stepDefaultInput
// service := i8_runner.StepDefaultInput{}
// service = i8_runner.DecorateStepVerboseSwitchable("stepDefaultInput", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputCircularDeps
// var service interface{}
// service = i8_runner.NewStepOutputValidationRule(eval("!value output.ValidateCircularDeps"), eval("Circular dependencies"))
// service = i8_runner.DecorateStepVerboseSwitchable("stepOutputCircularDeps", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputParamsExist
// var service *i8_runner.StepVerboseSwitchable
// service = i8_runner.NewStepOutputValidationRule(eval("!value output.ValidateParamsExist"), eval("Missing parameters"))
// service = i8_runner.DecorateStepVerboseSwitchable("stepOutputParamsExist", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputServicesExist
// var service *i8_runner.StepVerboseSwitchable
// service = i8_runner.NewStepOutputValidationRule(eval("!value output.ValidateServicesExist"), eval("Missing services"))
// service = i8_runner.DecorateStepVerboseSwitchable("stepOutputServicesExist", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepOutputServicesScopes
// var service interface{}
// service = i8_runner.NewStepOutputValidationRule(eval("!value output.ValidateServicesScopes"), eval("Scope"))
// service = i8_runner.DecorateStepVerboseSwitchable("stepOutputServicesScopes", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepReadConfig
// var service interface{}
// service = i8_runner.NewStepReadConfig(eval("@printer"), eval("%inputPatterns%"))
// service = i8_runner.DecorateStepVerboseSwitchable("stepReadConfig", service, eval("@printer"), eval("@printer"))
// ············································································
// #### stepValidateInput
// var service interface{}
// service = i3_compiler.NewStepValidateInput(eval("@inputValidator"))
// ············································································
// #### stepValidateOutput
// var service interface{}
// service = i8_runner.NewStepAmalgamated(eval("Validate output"), eval("@stepOutputServicesScopes"), eval("@stepOutputCircularDeps"), eval("@stepOutputParamsExist"), eval("@stepOutputServicesExist"))
// service = i8_runner.DecorateStepVerboseSwitchable("stepValidateOutput", service, eval("@printer"), eval("@printer"))
// ············································································
// #### taggedResolver
// var service interface{}
// service = i1_resolver.NewTaggedResolver()
// ············································································
// #### templateBuilder
// var service interface{}
// service = i2_template.NewBuilder(eval("@imports"), eval("@imports"), eval("@codeFormatter"), eval("%buildInfo%"), eval("%stub%"))
// ············································································
// #### tokenChunker
// var service interface{}
// service = i4_token.NewChunker()
// ············································································
// #### tokenStrategyFactory
// var service interface{}
// service = i4_token.NewStrategyFactory(eval("!value token.FactoryPercentMark{}"), eval("!value token.FactoryReference{}"), eval("!value token.FactoryUnexpectedFunction{}"), eval("!value token.FactoryUnexpectedToken{}"), eval("!value token.FactoryString{}"))
// ············································································
// #### tokenizer
// var service interface{}
// service = i4_token.NewTokenizer(eval("@tokenChunker"), eval("@tokenStrategyFactory"))
// ············································································
// #### valueResolver
// var service interface{}
// service = i1_resolver.NewValueResolver(eval("@imports"))
// ············································································
// #### writer
// panic("todo")
// ············································································

type gontainer struct {
	*ia_container.Container
}

func init() {
	interface_ := (*interface {
		// service container
		Get(serviceID string) (interface{}, error)
		GetInContext(ctx ib_context.Context, serviceID string) (interface{}, error)
		CircularDeps() error
		OverrideService(serviceID string, s ia_container.Service)
		AddDecorator(tag string, decorator interface{}, deps ...ia_container.Dependency)
		IsTaggedBy(serviceID string, tag string) bool
		GetTaggedBy(tag string) ([]interface{}, error)
		GetTaggedByInContext(ctx ib_context.Context, tag string) ([]interface{}, error)
		// param container
		GetParam(paramID string) (interface{}, error)
		OverrideParam(paramID string, d ia_container.Dependency)
		// misc
		HotSwap(func(ia_container.MutableContainer))
		Root() *ia_container.Container
		// getters
		GetRunner() (*i8_runner.Runner, error)
		GetRunnerInContext(ctx ib_context.Context) (*i8_runner.Runner, error)
		MustGetRunner() *i8_runner.Runner
		MustGetRunnerInContext(ctx ib_context.Context) *i8_runner.Runner
		GetStepValidateParamsExist() (*i8_runner.StepVerboseSwitchable, error)
		GetStepValidateParamsExistInContext(ctx ib_context.Context) (*i8_runner.StepVerboseSwitchable, error)
		MustGetStepValidateParamsExist() *i8_runner.StepVerboseSwitchable
		MustGetStepValidateParamsExistInContext(ctx ib_context.Context) *i8_runner.StepVerboseSwitchable
		GetStepValidateServicesExist() (*i8_runner.StepVerboseSwitchable, error)
		GetStepValidateServicesExistInContext(ctx ib_context.Context) (*i8_runner.StepVerboseSwitchable, error)
		MustGetStepValidateServicesExist() *i8_runner.StepVerboseSwitchable
		MustGetStepValidateServicesExistInContext(ctx ib_context.Context) *i8_runner.StepVerboseSwitchable
	})(nil)
	var nilContainer *gontainer
	interfaceType := ic_reflect.TypeOf(interface_).Elem()
	implements := ic_reflect.TypeOf(nilContainer).Implements(interfaceType)
	if !implements {
		panic("generated container does not implement expected interface")
	}
}

func (c *gontainer) GetRunner() (result *i8_runner.Runner, err error) {
	var s interface{}
	s, err = c.Get("runner")
	if err != nil {
		return nil, id_grouperror.Prefix(
			i0_fmt.Sprintf("%s.%s(): ", "gontainer", "GetRunner"),
			err,
		)
	}
	err = id_grouperror.Prefix(
		i0_fmt.Sprintf("%s.%s(): ", "gontainer", "GetRunner"),
		ie_copier.Copy(s, &result, true),
	)
	return
}

func (c *gontainer) GetRunnerInContext(ctx ib_context.Context) (result *i8_runner.Runner, err error) {
	var s interface{}
	s, err = c.GetInContext(ctx, "runner")
	if err != nil {
		return nil, id_grouperror.Prefix(
			i0_fmt.Sprintf("%s.%sInContext(): ", "gontainer", "GetRunner"),
			err,
		)
	}
	err = id_grouperror.Prefix(
		i0_fmt.Sprintf("%s.%sInContext(): ", "gontainer", "GetRunner"),
		ie_copier.Copy(s, &result, true),
	)
	return
}

func (c *gontainer) MustGetRunner() *i8_runner.Runner {
	r, err := c.GetRunner()
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) MustGetRunnerInContext(ctx ib_context.Context) *i8_runner.Runner {
	r, err := c.GetRunnerInContext(ctx)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) GetStepValidateParamsExist() (result *i8_runner.StepVerboseSwitchable, err error) {
	var s interface{}
	s, err = c.Get("stepOutputParamsExist")
	if err != nil {
		return nil, id_grouperror.Prefix(
			i0_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateParamsExist"),
			err,
		)
	}
	err = id_grouperror.Prefix(
		i0_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateParamsExist"),
		ie_copier.Copy(s, &result, true),
	)
	return
}

func (c *gontainer) GetStepValidateParamsExistInContext(ctx ib_context.Context) (result *i8_runner.StepVerboseSwitchable, err error) {
	var s interface{}
	s, err = c.GetInContext(ctx, "stepOutputParamsExist")
	if err != nil {
		return nil, id_grouperror.Prefix(
			i0_fmt.Sprintf("%s.%sInContext(): ", "gontainer", "GetStepValidateParamsExist"),
			err,
		)
	}
	err = id_grouperror.Prefix(
		i0_fmt.Sprintf("%s.%sInContext(): ", "gontainer", "GetStepValidateParamsExist"),
		ie_copier.Copy(s, &result, true),
	)
	return
}

func (c *gontainer) MustGetStepValidateParamsExist() *i8_runner.StepVerboseSwitchable {
	r, err := c.GetStepValidateParamsExist()
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) MustGetStepValidateParamsExistInContext(ctx ib_context.Context) *i8_runner.StepVerboseSwitchable {
	r, err := c.GetStepValidateParamsExistInContext(ctx)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) GetStepValidateServicesExist() (result *i8_runner.StepVerboseSwitchable, err error) {
	var s interface{}
	s, err = c.Get("stepOutputServicesExist")
	if err != nil {
		return nil, id_grouperror.Prefix(
			i0_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateServicesExist"),
			err,
		)
	}
	err = id_grouperror.Prefix(
		i0_fmt.Sprintf("%s.%s(): ", "gontainer", "GetStepValidateServicesExist"),
		ie_copier.Copy(s, &result, true),
	)
	return
}

func (c *gontainer) GetStepValidateServicesExistInContext(ctx ib_context.Context) (result *i8_runner.StepVerboseSwitchable, err error) {
	var s interface{}
	s, err = c.GetInContext(ctx, "stepOutputServicesExist")
	if err != nil {
		return nil, id_grouperror.Prefix(
			i0_fmt.Sprintf("%s.%sInContext(): ", "gontainer", "GetStepValidateServicesExist"),
			err,
		)
	}
	err = id_grouperror.Prefix(
		i0_fmt.Sprintf("%s.%sInContext(): ", "gontainer", "GetStepValidateServicesExist"),
		ie_copier.Copy(s, &result, true),
	)
	return
}

func (c *gontainer) MustGetStepValidateServicesExist() *i8_runner.StepVerboseSwitchable {
	r, err := c.GetStepValidateServicesExist()
	if err != nil {
		panic(err.Error())
	}
	return r
}

func (c *gontainer) MustGetStepValidateServicesExistInContext(ctx ib_context.Context) *i8_runner.StepVerboseSwitchable {
	r, err := c.GetStepValidateServicesExistInContext(ctx)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func New() (rootGontainer *gontainer) {
	c := &gontainer{
		Container: ia_container.New(),
	}
	rootGontainer = c
	//
	//
	// #####################################
	// ############## Helpers ##############
	//
	//
	dependencyService := ia_container.NewDependencyService
	_ = dependencyService
	dependencyValue := ia_container.NewDependencyValue
	_ = dependencyValue
	dependencyTag := ia_container.NewDependencyTag
	_ = dependencyTag
	dependencyProvider := ia_container.NewDependencyProvider
	_ = dependencyProvider
	newService := ia_container.NewService
	_ = newService
	concatenateChunks := c._concatenateChunks
	_ = concatenateChunks
	paramTodo := c._paramTodo
	_ = paramTodo
	getEnv := c._getEnv
	_ = getEnv
	getEnvInt := c._getEnvInt
	_ = getEnvInt
	getParam := c.GetParam
	_ = getParam
	callProvider := c._callProvider
	_ = callProvider
	//
	//
	// ####################################
	// ############## Params ##############
	//
	//
	// "buildInfo": "%todo()%"
	c.OverrideParam("buildInfo", dependencyProvider(func() (r interface{}, err error) {
		r, err = callProvider(paramTodo)
		if err != nil {
			err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	// "inputPatterns": "%todo()%"
	c.OverrideParam("inputPatterns", dependencyProvider(func() (r interface{}, err error) {
		r, err = callProvider(paramTodo)
		if err != nil {
			err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	// "outputFile": "%todo()%"
	c.OverrideParam("outputFile", dependencyProvider(func() (r interface{}, err error) {
		r, err = callProvider(paramTodo)
		if err != nil {
			err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	// "stub": "%todo()%"
	c.OverrideParam("stub", dependencyProvider(func() (r interface{}, err error) {
		r, err = callProvider(paramTodo)
		if err != nil {
			err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
		}
		return
	}))
	// "version": "%todo()%"
	c.OverrideParam("version", dependencyProvider(func() (r interface{}, err error) {
		r, err = callProvider(paramTodo)
		if err != nil {
			err = i0_fmt.Errorf("%s: %w", "cannot execute %todo()%", err)
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
			i1_resolver.NewArgResolver,
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
		s.SetScopeDefault()
		c.OverrideService("argResolver", s)
	}
	// "codeFormatter"
	{
		s := newService()
		s.SetConstructor(
			i2_template.NewCodeFormatter,
		)
		s.SetScopeDefault()
		c.OverrideService("codeFormatter", s)
	}
	// "compiler"
	{
		s := newService()
		s.SetConstructor(
			i3_compiler.New,
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
		s.SetScopeDefault()
		c.OverrideService("compiler", s)
	}
	// "fnRegisterer"
	{
		s := newService()
		s.SetConstructor(
			i4_token.NewFuncRegisterer,
			// "@tokenStrategyFactory"
			dependencyService("tokenStrategyFactory"),
			// "@imports"
			dependencyService("imports"),
		)
		s.SetScopeDefault()
		c.OverrideService("fnRegisterer", s)
	}
	// "gontainerValueResolver"
	{
		s := newService()
		s.SetConstructor(
			i1_resolver.NewFixedValueResolver,
			// "!value consts.SpecialGontainerID"
			dependencyValue(i5_consts.SpecialGontainerID),
			// "!value consts.SpecialGontainerValue"
			dependencyValue(i5_consts.SpecialGontainerValue),
		)
		s.SetScopeDefault()
		c.OverrideService("gontainerValueResolver", s)
	}
	// "imports"
	{
		s := newService()
		s.SetConstructor(
			i6_imports.New,
		)
		s.SetScopeDefault()
		c.OverrideService("imports", s)
	}
	// "inputValidator"
	{
		s := newService()
		s.SetConstructor(
			i7_input.NewDefaultValidator,
			// "%version%"
			dependencyProvider(func() (interface{}, error) { return getParam("version") }),
		)
		s.SetScopeDefault()
		c.OverrideService("inputValidator", s)
	}
	// "nonStringPrimitiveResolver"
	{
		s := newService()
		s.SetConstructor(
			i1_resolver.NewNonStringPrimitiveResolver,
		)
		s.SetScopeDefault()
		c.OverrideService("nonStringPrimitiveResolver", s)
	}
	// "paramResolver"
	{
		s := newService()
		s.SetConstructor(
			i1_resolver.NewParamResolver,
			// "@primitiveArgResolver"
			dependencyService("primitiveArgResolver"),
		)
		s.SetScopeDefault()
		c.OverrideService("paramResolver", s)
	}
	// "patternResolver"
	{
		s := newService()
		s.SetConstructor(
			i1_resolver.NewPatternResolver,
			// "@tokenizer"
			dependencyService("tokenizer"),
		)
		s.SetScopeDefault()
		c.OverrideService("patternResolver", s)
	}
	// "primitiveArgResolver"
	{
		s := newService()
		s.SetConstructor(
			i1_resolver.NewArgResolver,
			// "@nonStringPrimitiveResolver"
			dependencyService("nonStringPrimitiveResolver"),
			// "@patternResolver"
			dependencyService("patternResolver"),
		)
		s.SetScopeDefault()
		c.OverrideService("primitiveArgResolver", s)
	}
	// "printer"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewPrinter,
			// "@writer"
			dependencyService("writer"),
		)
		s.SetScopeDefault()
		c.OverrideService("printer", s)
	}
	// "runner"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewRunner,
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
		s.SetScopeDefault()
		c.OverrideService("runner", s)
	}
	// "serviceResolver"
	{
		s := newService()
		s.SetConstructor(
			i1_resolver.NewServiceResolver,
		)
		s.SetScopeDefault()
		c.OverrideService("serviceResolver", s)
	}
	// "stepCodeGenerator"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewStepCodeGenerator,
			// "@printer"
			dependencyService("printer"),
			// "@templateBuilder"
			dependencyService("templateBuilder"),
			// "%outputFile%"
			dependencyProvider(func() (interface{}, error) { return getParam("outputFile") }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepCodeGenerator", s)
	}
	// "stepCompile"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewStepCompile,
			// "@compiler"
			dependencyService("compiler"),
		)
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepCompile", s)
	}
	// "stepCompileDecorators"
	{
		s := newService()
		s.SetConstructor(
			i3_compiler.NewStepCompileDecorators,
			// "@imports"
			dependencyService("imports"),
			// "@argResolver"
			dependencyService("argResolver"),
		)
		s.SetScopeDefault()
		c.OverrideService("stepCompileDecorators", s)
	}
	// "stepCompileMeta"
	{
		s := newService()
		s.SetConstructor(
			i3_compiler.NewStepCompileMeta,
			// "@imports"
			dependencyService("imports"),
			// "@fnRegisterer"
			dependencyService("fnRegisterer"),
		)
		s.SetScopeDefault()
		c.OverrideService("stepCompileMeta", s)
	}
	// "stepCompileParams"
	{
		s := newService()
		s.SetConstructor(
			i3_compiler.NewStepCompileParams,
			// "@paramResolver"
			dependencyService("paramResolver"),
		)
		s.SetScopeDefault()
		c.OverrideService("stepCompileParams", s)
	}
	// "stepCompileServices"
	{
		s := newService()
		s.SetConstructor(
			i3_compiler.NewStepCompileServices,
			// "@imports"
			dependencyService("imports"),
			// "@argResolver"
			dependencyService("argResolver"),
		)
		s.SetScopeDefault()
		c.OverrideService("stepCompileServices", s)
	}
	// "stepDefaultInput"
	{
		s := newService()
		s.SetConstructor(func() interface{} { return i8_runner.StepDefaultInput{} })
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepDefaultInput", s)
	}
	// "stepOutputCircularDeps"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewStepOutputValidationRule,
			// "!value output.ValidateCircularDeps"
			dependencyValue(i9_output.ValidateCircularDeps),
			// "Circular dependencies"
			dependencyProvider(func() (r interface{}, err error) { return "Circular dependencies", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepOutputCircularDeps", s)
	}
	// "stepOutputParamsExist"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewStepOutputValidationRule,
			// "!value output.ValidateParamsExist"
			dependencyValue(i9_output.ValidateParamsExist),
			// "Missing parameters"
			dependencyProvider(func() (r interface{}, err error) { return "Missing parameters", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepOutputParamsExist", s)
	}
	// "stepOutputServicesExist"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewStepOutputValidationRule,
			// "!value output.ValidateServicesExist"
			dependencyValue(i9_output.ValidateServicesExist),
			// "Missing services"
			dependencyProvider(func() (r interface{}, err error) { return "Missing services", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepOutputServicesExist", s)
	}
	// "stepOutputServicesScopes"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewStepOutputValidationRule,
			// "!value output.ValidateServicesScopes"
			dependencyValue(i9_output.ValidateServicesScopes),
			// "Scope"
			dependencyProvider(func() (r interface{}, err error) { return "Scope", nil }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepOutputServicesScopes", s)
	}
	// "stepReadConfig"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewStepReadConfig,
			// "@printer"
			dependencyService("printer"),
			// "%inputPatterns%"
			dependencyProvider(func() (interface{}, error) { return getParam("inputPatterns") }),
		)
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepReadConfig", s)
	}
	// "stepValidateInput"
	{
		s := newService()
		s.SetConstructor(
			i3_compiler.NewStepValidateInput,
			// "@inputValidator"
			dependencyService("inputValidator"),
		)
		s.SetScopeDefault()
		c.OverrideService("stepValidateInput", s)
	}
	// "stepValidateOutput"
	{
		s := newService()
		s.SetConstructor(
			i8_runner.NewStepAmalgamated,
			// "Validate output"
			dependencyProvider(func() (r interface{}, err error) { return "Validate output", nil }),
			// "@stepOutputServicesScopes"
			dependencyService("stepOutputServicesScopes"),
			// "@stepOutputCircularDeps"
			dependencyService("stepOutputCircularDeps"),
			// "@stepOutputParamsExist"
			dependencyService("stepOutputParamsExist"),
			// "@stepOutputServicesExist"
			dependencyService("stepOutputServicesExist"),
		)
		s.Tag("step-runner-verbose", int(0))
		s.SetScopeDefault()
		c.OverrideService("stepValidateOutput", s)
	}
	// "taggedResolver"
	{
		s := newService()
		s.SetConstructor(
			i1_resolver.NewTaggedResolver,
		)
		s.SetScopeDefault()
		c.OverrideService("taggedResolver", s)
	}
	// "templateBuilder"
	{
		s := newService()
		s.SetConstructor(
			i2_template.NewBuilder,
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
		s.SetScopeDefault()
		c.OverrideService("templateBuilder", s)
	}
	// "tokenChunker"
	{
		s := newService()
		s.SetConstructor(
			i4_token.NewChunker,
		)
		s.SetScopeDefault()
		c.OverrideService("tokenChunker", s)
	}
	// "tokenStrategyFactory"
	{
		s := newService()
		s.SetConstructor(
			i4_token.NewStrategyFactory,
			// "!value token.FactoryPercentMark{}"
			dependencyValue(i4_token.FactoryPercentMark{}),
			// "!value token.FactoryReference{}"
			dependencyValue(i4_token.FactoryReference{}),
			// "!value token.FactoryUnexpectedFunction{}"
			dependencyValue(i4_token.FactoryUnexpectedFunction{}),
			// "!value token.FactoryUnexpectedToken{}"
			dependencyValue(i4_token.FactoryUnexpectedToken{}),
			// "!value token.FactoryString{}"
			dependencyValue(i4_token.FactoryString{}),
		)
		s.SetScopeDefault()
		c.OverrideService("tokenStrategyFactory", s)
	}
	// "tokenizer"
	{
		s := newService()
		s.SetConstructor(
			i4_token.NewTokenizer,
			// "@tokenChunker"
			dependencyService("tokenChunker"),
			// "@tokenStrategyFactory"
			dependencyService("tokenStrategyFactory"),
		)
		s.SetScopeDefault()
		c.OverrideService("tokenizer", s)
	}
	// "valueResolver"
	{
		s := newService()
		s.SetConstructor(
			i1_resolver.NewValueResolver,
			// "@imports"
			dependencyService("imports"),
		)
		s.SetScopeDefault()
		c.OverrideService("valueResolver", s)
	}
	// "writer"
	{
		s := newService()
		s.SetConstructor(func() (interface{}, error) { return nil, if_errors.New("service todo") })
		c.OverrideService("writer", s)
	}
	//
	//
	// ########################################
	// ############## Decorators ##############
	//
	//
	c.AddDecorator(
		"step-runner-verbose",
		i8_runner.DecorateStepVerboseSwitchable,
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
		return "", i0_fmt.Errorf("environment variable %+q does not exist", key)
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
		return 0, i0_fmt.Errorf("environment variable %+q does not exist", key)
	}
	res, err := i12_strconv.Atoi(val)
	if err != nil {
		return 0, i0_fmt.Errorf("cannot cast env(%+q) to int: %w", key, err)
	}
	return res, nil
}

// Deprecated: do not use it, only for internal purposes, that method can be changed at any time
func (c *gontainer) _callProvider(provider interface{}, args ...interface{}) (interface{}, error) {
	return i13_caller.CallProvider(provider, args, true)
}
