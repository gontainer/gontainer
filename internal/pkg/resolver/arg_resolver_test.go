package resolver_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/resolver"
	"github.com/stretchr/testify/assert"
)

func TestArgResolver_ResolveArg(t *testing.T) {
	t.Run("Not supported", func(t *testing.T) {
		arg, err := resolver.NewArgResolver().ResolveArg("arg")
		assert.EqualError(t, err, "not supported string")
		assert.Empty(t, arg)
	})
	t.Run("OK", func(t *testing.T) {
		arg, err := resolver.NewArgResolver(resolver.NewNonStringPrimitiveResolver()).ResolveArg(5)
		assert.NoError(t, err)
		assert.Equal(
			t,
			resolver.ArgExpr{
				Code:              `dependencyValue(int(5))`,
				Raw:               5,
				DependsOnParams:   nil,
				DependsOnServices: nil,
				DependsOnTags:     nil,
			},
			arg,
		)
	})
}
