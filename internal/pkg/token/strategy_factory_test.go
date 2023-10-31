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

package token_test

import (
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/token"
	"github.com/stretchr/testify/assert"
)

func TestStrategyFactory_Prepend(t *testing.T) {
	t.Run("Prepend", func(t *testing.T) {
		f := token.NewStrategyFactory()

		// func is not registered yet
		tk, err := f.Create(`%getEnv("HOST")%`)
		assert.EqualError(t, err, `not supported token: %getEnv("HOST")%`)
		assert.Empty(t, tk)

		// let's register func
		f.Prepend(token.NewFactoryFunction(mockAliaser{}, "getEnv", "pkg", "GetEnv"))
		tk, err = f.Create(`%getEnv("HOST")%`)
		assert.NoError(t, err)
		assert.Equal(t, token.Token{
			Kind:      token.KindFunc,
			Raw:       `%getEnv("HOST")%`,
			DependsOn: nil,
			Code:      `func() (r interface{}, err error) { r, err = callProvider(pkg.GetEnv, "HOST"); if err != nil { err = fmt.Errorf("%s: %w", "cannot execute %getEnv(\"HOST\")%", err) }; return }`,
		}, tk)
	})

	t.Run("RegisterFunc", func(t *testing.T) {
		f := token.NewStrategyFactory()

		// func is not registered yet
		tk, err := f.Create(`%secret("password")%`)
		assert.EqualError(t, err, `not supported token: %secret("password")%`)
		assert.Empty(t, tk)

		// let's register func using registerer
		token.NewFuncRegisterer(f, mockAliaser{}).RegisterFunc("secret", "secrets", "Secret")
		tk, err = f.Create(`%secret("password")%`)
		assert.NoError(t, err)
		assert.Equal(t, token.Token{
			Kind:      token.KindFunc,
			Raw:       `%secret("password")%`,
			DependsOn: nil,
			Code:      `func() (r interface{}, err error) { r, err = callProvider(secrets.Secret, "password"); if err != nil { err = fmt.Errorf("%s: %w", "cannot execute %secret(\"password\")%", err) }; return }`,
		}, tk)

		// let's register another func using registerer,
		// since it will prepend existing slice, the newer func has the bigger priority
		token.NewFuncRegisterer(f, mockAliaser{}).RegisterFunc("secret", "anotherSecrets", "SuperSecret")
		tk, err = f.Create(`%secret("password")%`)
		assert.NoError(t, err)
		assert.Equal(t, token.Token{
			Kind:      token.KindFunc,
			Raw:       `%secret("password")%`,
			DependsOn: nil,
			Code:      `func() (r interface{}, err error) { r, err = callProvider(anotherSecrets.SuperSecret, "password"); if err != nil { err = fmt.Errorf("%s: %w", "cannot execute %secret(\"password\")%", err) }; return }`,
		}, tk)
	})
}
