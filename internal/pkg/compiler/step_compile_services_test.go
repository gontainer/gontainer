package compiler_test

import (
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/compiler"
	"github.com/gontainer/gontainer/internal/pkg/consts"
	"github.com/gontainer/gontainer/internal/pkg/imports"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/gontainer/gontainer/internal/pkg/resolver"
	"github.com/gontainer/gontainer/internal/pkg/token"
	"github.com/stretchr/testify/assert"
)

func newRealStepCompilerServices() *compiler.StepCompileServices {
	tokenStrategyFactory := token.NewStrategyFactory(
		token.FactoryPercentMark{},
		token.FactoryReference{},
		token.FactoryUnexpectedFunction{},
		token.FactoryUnexpectedToken{},
		token.FactoryString{},
	)
	tokenizer := token.NewTokenizer(token.NewChunker(), tokenStrategyFactory)
	imports_ := imports.New()
	argResolver := resolver.NewArgResolver(
		resolver.NewNonStringPrimitiveResolver(),
		resolver.NewValueResolver(imports_),
		resolver.NewServiceResolver(),
		resolver.NewTaggedResolver(),
		resolver.NewFixedValueResolver(consts.SpecialGontainerID, consts.SpecialGontainerValue),
		resolver.NewPatternResolver(tokenizer),
	)
	return compiler.NewStepCompileServices(imports_, argResolver)
}

func TestStepCompileServices_Process(t *testing.T) {
	t.Parallel()

	defStep := compiler.NewStepCompileServices(simpleAliaserFunc, simpleArgResolverFunc)

	scenarios := []struct {
		Name     string
		Input    input.Input
		Expected output.Output
		Step     *compiler.StepCompileServices
		Errors   []string
	}{
		{
			Name: "Must getter #1 (error)",
			Input: input.Input{
				Services: map[string]input.Service{
					"db": {
						MustGetter: ptr.New(true),
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name:       "db",
						MustGetter: true,
						Type:       "interface{}",
					},
				},
			},
			Step: defStep,
			Errors: []string{
				`compiler.StepCompileServices: "db": cannot generate a must-getter when the getter is not specified`,
			},
		},
		{
			Name: "Must getter #2 (ok)",
			Input: input.Input{
				Meta: input.Meta{
					DefaultMustGetter: ptr.New(true),
				},
				Services: map[string]input.Service{
					"db": {
						MustGetter: nil,
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name:       "db",
						MustGetter: false,
						Type:       "interface{}",
					},
				},
			},
			Step: defStep,
		},
		{
			Name: "Must getter #3 (ok)",
			Input: input.Input{
				Meta: input.Meta{
					DefaultMustGetter: ptr.New(true),
				},
				Services: map[string]input.Service{
					"db": {
						Getter:     ptr.New("GetDB"),
						MustGetter: nil,
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name:       "db",
						Getter:     "GetDB",
						MustGetter: true,
						Type:       "interface{}",
					},
				},
			},
			Step: defStep,
		},
		{
			Name: "Scopes",
			Input: input.Input{
				Services: map[string]input.Service{
					"serviceA": {
						Scope: ptr.New(input.ScopeShared),
					},
					"serviceB": {
						Scope: ptr.New(input.ScopeContextual),
					},
					"serviceC": {
						Scope: ptr.New(input.ScopeNonShared),
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name:  "serviceA",
						Type:  "interface{}",
						Scope: output.ScopeShared,
					},
					{
						Name:  "serviceB",
						Type:  "interface{}",
						Scope: output.ScopeContextual,
					},
					{
						Name:  "serviceC",
						Type:  "interface{}",
						Scope: output.ScopeNonShared,
					},
				},
			},
			Step:   defStep,
			Errors: nil,
		},
		{
			Name: "Contextual scope",
			Input: input.Input{
				Services: map[string]input.Service{
					"transaction": {
						Scope: ptr.New(input.ScopeContextual),
					},
					"repository": {
						Fields: map[string]any{
							"Transaction": "@transaction",
						},
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name:  "repository",
						Type:  "interface{}",
						Scope: output.ScopeDefault,
						Fields: []output.Field{
							{
								Name: "Transaction",
								Value: output.Arg{
									Code:              `dependencyService("transaction")`,
									Raw:               "@transaction",
									DependsOnParams:   nil,
									DependsOnServices: []string{"transaction"},
									DependsOnTags:     nil,
								},
							},
						},
					},
					{
						Name:  "transaction",
						Type:  "interface{}",
						Scope: output.ScopeContextual,
					},
				},
			},
			Step:   newRealStepCompilerServices(),
			Errors: nil,
		},
		{
			Name: "TODO",
			Input: input.Input{
				Services: map[string]input.Service{
					"serviceTodo": {
						Todo: ptr.New(true),
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name: "serviceTodo",
						Todo: true,
					},
				},
			},
			Step:   defStep,
			Errors: nil,
		},
		{
			Name: "Compile type and value",
			Input: input.Input{
				Services: map[string]input.Service{
					"db": {
						Type:  ptr.New("*sql.DB"),
						Value: ptr.New("&sql.DB"),
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name:  "db",
						Type:  "*alias_sql.DB",
						Value: "&alias_sql.DB",
					},
				},
			},
			Step:   defStep,
			Errors: nil,
		},
		{
			Name: "Compile constructor",
			Input: input.Input{
				Services: map[string]input.Service{
					"db": {
						Constructor: ptr.New("pkg.NewDB"),
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name:        "db",
						Constructor: "alias_pkg.NewDB",
						Type:        "interface{}",
					},
				},
			},
			Step:   defStep,
			Errors: nil,
		},
		{
			Name: "Cannot compile field",
			Input: input.Input{
				Services: map[string]input.Service{
					"db": {
						Fields: map[string]any{
							"Port": "%port%%",
						},
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name: "db",
						Type: "interface{}",
						Fields: []output.Field{
							{
								Name:  "Port",
								Value: output.Arg{},
							},
						},
					},
				},
			},
			Step: newRealStepCompilerServices(),
			Errors: []string{
				`compiler.StepCompileServices: "db": fields: "Port": not closed token: "%"`,
			},
		},
		{
			Name: "Calls",
			Input: input.Input{
				Services: map[string]input.Service{
					"db": {
						Calls: []input.Call{
							{
								Method:    "SetHost",
								Args:      []any{"localhost"},
								Immutable: false,
							},
							{
								Method:    "WithPort",
								Args:      []any{"%"},
								Immutable: true,
							},
						},
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name: "db",
						Type: "interface{}",
						Calls: []output.Call{
							{
								Method: "SetHost",
								Args: []output.Arg{
									{
										Code:              `dependencyProvider(func() (r interface{}, err error) { return "localhost", nil })`,
										Raw:               "localhost",
										DependsOnParams:   nil,
										DependsOnServices: nil,
										DependsOnTags:     nil,
									},
								},
								Immutable: false,
							},
							{
								Method: "WithPort",
								Args: []output.Arg{
									{},
								},
								Immutable: true,
							},
						},
					},
				},
			},
			Step: newRealStepCompilerServices(),
			Errors: []string{
				`compiler.StepCompileServices: "db": calls: 1: args: 0: not closed token: "%"`,
			},
		},
		{
			Name: "Tags",
			Input: input.Input{
				Services: map[string]input.Service{
					"endpointHome": {
						Tags: []input.Tag{
							{
								Name:     "http-endpoint",
								Priority: 50,
							},
							{
								Name:     "http",
								Priority: 0,
							},
						},
					},
				},
			},
			Expected: output.Output{
				Services: []output.Service{
					{
						Name: "endpointHome",
						Type: "interface{}",
						Tags: []output.Tag{
							{
								Name:     "http-endpoint",
								Priority: 50,
							},
							{
								Name:     "http",
								Priority: 0,
							},
						},
					},
				},
			},
			Step:   newRealStepCompilerServices(),
			Errors: nil,
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()

			o := output.Output{}
			err := s.Step.Process(s.Input, &o)
			errAssert.EqualErrorGroup(t, err, s.Errors)
			assert.Equal(t, s.Expected, o)
		})
	}
}
