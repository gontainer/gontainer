package types_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/types"
	"github.com/stretchr/testify/assert"
)

type MyInt int

func TestIsPrimitive(t *testing.T) {
	scenarios := []struct {
		Name     string
		Value    any
		Expected bool
	}{
		{
			Name:     "MyInt(5)",
			Value:    MyInt(5),
			Expected: true,
		},
		{
			Name:     "nil",
			Value:    nil,
			Expected: true,
		},
		{
			Name:     "(*error)(nil)",
			Value:    (*error)(nil),
			Expected: false,
		},
		{
			Name:     "(error)(nil)",
			Value:    (error)(nil),
			Expected: true,
		},
		{
			Name:     "(*int)(nil)",
			Value:    (*int)(nil),
			Expected: false,
		},
		{
			Name:     "(*interface{})(nil)",
			Value:    (*interface{})(nil),
			Expected: false,
		},
		{
			Name:     "(interface{})(nil)",
			Value:    (interface{})(nil),
			Expected: true,
		},
		{
			Name:     "(interface{ Foo() })(nil)",
			Value:    (interface{ Foo() })(nil),
			Expected: true,
		},
		{
			Name:     "struct{}{}",
			Value:    struct{}{},
			Expected: false,
		},
		{
			Name:     "int",
			Value:    5,
			Expected: true,
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, s.Expected, types.IsPrimitive(s.Value))
		})
	}
}
