package resolver_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
	"github.com/stretchr/testify/assert"
)

type resolverStrategy interface {
	ResolveArg(any) (resolver.ArgExpr, error)
	Supports(any) bool
}

type anyResolverStrategyScenario struct {
	Name       string
	Expression any
	Supports   bool
	Arg        resolver.ArgExpr
	Error      string
	Resolver   resolverStrategy
}

func assertScenarios(t *testing.T, scenarios ...anyResolverStrategyScenario) {
	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			if !s.Supports {
				assert.False(t, s.Resolver.Supports(s.Expression))
				return
			}
			assert.True(t, s.Resolver.Supports(s.Expression))
			arg, err := s.Resolver.ResolveArg(s.Expression)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, s.Arg, arg)
		})
	}
}
