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
