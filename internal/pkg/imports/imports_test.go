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

package imports_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/imports"
	"github.com/stretchr/testify/assert"
)

const (
	viperPkg = "github.com/spf13/viper"
)

func Test_imports_RegisterPrefixAlias(t *testing.T) {
	t.Run("Given scenario", func(t *testing.T) {
		i := imports.New()
		shortcut := "viper"
		expectedViper := "i0_viper"
		expectedRemote := "i1_remote"

		if !assert.NoError(t, i.RegisterPrefixAlias(shortcut, viperPkg)) {
			return
		}

		assert.Empty(t, i.Imports()) // no import has been used yet

		// aliases for "viper" and "github.com/spf13/viper" should be equal
		assert.Equal(t, expectedViper, i.Alias(shortcut))
		assert.Equal(t, expectedViper, i.Alias(viperPkg))

		// aliases for "viper/remote" and "github.com/spf13/viper/remote" should be equal
		assert.Equal(t, expectedRemote, i.Alias(shortcut+"/remote"))
		assert.Equal(t, expectedRemote, i.Alias(viperPkg+"/remote"))

		assert.Equal(
			t,
			[]imports.Import{
				{Alias: expectedViper, Path: "github.com/spf13/viper"},
				{Alias: expectedRemote, Path: "github.com/spf13/viper/remote"},
			},
			i.Imports(),
		)
	})

	t.Run("Given error", func(t *testing.T) {
		i := imports.New()
		if !assert.NoError(t, i.RegisterPrefixAlias("viper", viperPkg)) {
			return
		}
		assert.EqualError(
			t,
			i.RegisterPrefixAlias("viper", viperPkg),
			`prefix is already registered: "viper"`,
		)
	})
}
