package resolver_test

import (
	"fmt"
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

type mockAliaser struct {
}

func (a mockAliaser) Alias(i string) string {
	return "aliasPkg"
}

func valueResolverScenarios() []anyResolverStrategyScenario {
	r := resolver.NewValueResolver(mockAliaser{})

	scenariosOK := []struct {
		Expression string
		Expected   string
	}{
		{
			Expression: "!value Variable",
			Expected:   "Variable",
		},
		{
			Expression: "!value &Variable",
			Expected:   "&Variable",
		},
		{
			Expression: `!value "pkg.db".Singleton`,
			Expected:   `aliasPkg.Singleton`,
		},
		{
			Expression: `!value &"pkg.db".Singleton`,
			Expected:   `&aliasPkg.Singleton`,
		},
		{
			Expression: `!value "pkg.domain".MyType{}`,
			Expected:   `aliasPkg.MyType{}`,
		},
		{
			Expression: `!value &"pkg.domain".MyType{}`,
			Expected:   `&aliasPkg.MyType{}`,
		},
		{
			Expression: `!value "my/import".GlobalVar.Field`,
			Expected:   `aliasPkg.GlobalVar.Field`,
		},
		{
			Expression: `!value "my/import".GlobalVar.Substruct1.Substruct2.Field`,
			Expected:   `aliasPkg.GlobalVar.Substruct1.Substruct2.Field`,
		},
		{
			Expression: `!value pkg.imp.v1.Field`,
			Expected:   `aliasPkg.Field`,
		},
		{
			Expression: `!value &"my/import".GlobalVar.Field`,
			Expected:   `&aliasPkg.GlobalVar.Field`,
		},
	}

	var scenarios []anyResolverStrategyScenario

	for i, s := range scenariosOK {
		scenarios = append(scenarios, struct {
			Name       string
			Expression any
			Supports   bool
			Arg        resolver.ArgExpr
			Error      string
			Resolver   resolverStrategy
		}{
			Name:       fmt.Sprintf("OK #%d", i),
			Expression: s.Expression,
			Supports:   true,
			Arg: resolver.ArgExpr{
				Code: fmt.Sprintf(`dependencyValue(%s)`, s.Expected),
				Raw:  s.Expression,
			},
			Resolver: r,
		})
	}

	scenarios = append(
		scenarios,
		anyResolverStrategyScenario{
			Name:       "Error",
			Expression: "!value Variable()",
			Supports:   true,
			Arg:        resolver.ArgExpr{},
			Error:      "invalid value",
			Resolver:   r,
		},
		anyResolverStrategyScenario{
			Name:       "primitive",
			Expression: 5,
			Supports:   false,
			Resolver:   r,
		},
	)

	return scenarios
}

func TestValueResolver_ResolveArg(t *testing.T) {
	t.Parallel()
	assertScenarios(t, valueResolverScenarios()...)
}
