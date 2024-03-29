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

	errAssert "github.com/gontainer/gontainer-helpers/v3/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenizer_Tokenize(t *testing.T) {
	t.Parallel()

	tokenStrategies := token.NewStrategyFactory(
		token.NewFactoryFunction(mockAliaser{}, "env", "helper", "GetEnv"),
		token.FactoryPercentMark{},
		token.FactoryReference{},
		token.FactoryUnexpectedFunction{},
		token.FactoryString{},
	)
	tokenizer := token.NewTokenizer(token.NewChunker(), tokenStrategies)

	scenarios := []struct {
		Name      string
		Tokenizer *token.Tokenizer
		Input     string
		Tokens    token.Tokens
		Error     string
		Code      string
		CodeError string
	}{
		{
			Name:      "Reference",
			Tokenizer: tokenizer,
			Input:     "%host%",
			Tokens: token.Tokens{{
				Kind:      token.KindReference,
				Raw:       "%host%",
				DependsOn: []string{"host"},
				Code:      `func() (interface{}, error) { return getParam("host") }`,
			}},
			Error:     "",
			Code:      `dependencyProvider(func() (interface{}, error) { return getParam("host") })`,
			CodeError: "",
		},
		{
			Name:      "%%",
			Tokenizer: tokenizer,
			Input:     "%%",
			Tokens: token.Tokens{{
				Kind:      token.KindString,
				Raw:       "%%",
				DependsOn: nil,
				Code:      `func() (r interface{}, err error) { return "%", nil }`,
			}},
			Error:     "",
			Code:      `dependencyProvider(func() (r interface{}, err error) { return "%", nil })`,
			CodeError: "",
		},
		{
			Name:      "host:port",
			Tokenizer: tokenizer,
			Input:     "%host%:%port%",
			Tokens: token.Tokens{
				{
					Kind:      token.KindReference,
					Raw:       "%host%",
					DependsOn: []string{"host"},
					Code:      `func() (interface{}, error) { return getParam("host") }`,
				},
				{
					Kind:      token.KindString,
					Raw:       ":",
					DependsOn: nil,
					Code:      `func() (r interface{}, err error) { return ":", nil }`,
				},
				{
					Kind:      token.KindReference,
					Raw:       "%port%",
					DependsOn: []string{"port"},
					Code:      `func() (interface{}, error) { return getParam("port") }`,
				},
			},
			Error: "",
			Code: `dependencyProvider(func () (string, error) { return concatenateChunks(` +
				`func() (interface{}, error) { return getParam("host") }, ` +
				`func() (r interface{}, err error) { return ":", nil }, ` +
				`func() (interface{}, error) { return getParam("port") }) })`,
			CodeError: "",
		},
		{
			Name:      "Chunker error",
			Tokenizer: tokenizer,
			Input:     "%",
			Tokens:    nil,
			Error:     `not closed token: "%"`,
			Code:      "",
			CodeError: "",
		},
		{
			Name:      "Empty factories",
			Tokenizer: token.NewTokenizer(token.NewChunker(), token.NewStrategyFactory()),
			Input:     "input",
			Tokens:    nil,
			Error:     "not supported token: input",
			Code:      "",
			CodeError: "",
		},
		{
			Name:      "Empty input",
			Tokenizer: tokenizer,
			Input:     "",
			Tokens: token.Tokens{{
				Kind:      token.KindString,
				Raw:       "",
				DependsOn: nil,
				Code:      `func() (r interface{}, err error) { return "", nil }`,
			}},
			Error:     "",
			Code:      `dependencyProvider(func() (r interface{}, err error) { return "", nil })`,
			CodeError: "",
		},
		{
			Name:      "Func",
			Tokenizer: tokenizer,
			Input:     `%env("HOST")%`,
			Tokens: token.Tokens{{
				Kind:      token.KindFunc,
				Raw:       `%env("HOST")%`,
				DependsOn: nil,
				Code:      `func() (r interface{}, err error) { r, err = callProvider(helper.GetEnv, "HOST"); if err != nil { err = fmt.Errorf("%s: %w", "cannot execute %env(\"HOST\")%", err) }; return }`,
			}},
			Error:     "",
			Code:      `dependencyProvider(func() (r interface{}, err error) { r, err = callProvider(helper.GetEnv, "HOST"); if err != nil { err = fmt.Errorf("%s: %w", "cannot execute %env(\"HOST\")%", err) }; return })`,
			CodeError: "",
		},
		{
			Name:      "Unexpected func",
			Tokenizer: tokenizer,
			Input:     `%secret("password")%`,
			Tokens:    nil,
			Error:     `unexpected function: "secret": "%secret(\"password\")%"`,
			Code:      "",
			CodeError: "",
		},
		{
			Name:      "Unexpected token",
			Tokenizer: token.NewTokenizer(token.NewChunker(), token.NewStrategyFactory(token.FactoryUnexpectedToken{})),
			Input:     `%unexpectedToken%`,
			Tokens:    nil,
			Error:     `unexpected token: "%unexpectedToken%"`,
			Code:      "",
			CodeError: "",
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()

			tkns, err := s.Tokenizer.Tokenize(s.Input)
			if s.Error != "" {
				assert.EqualError(t, err, s.Error)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, s.Tokens, tkns)

			code, err := tkns.GoCode()
			if s.CodeError != "" {
				assert.EqualError(t, err, s.CodeError)
				assert.Empty(t, code)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, s.Code, code)
		})
	}

	t.Run("Error group", func(t *testing.T) {
		_, err := tokenizer.Tokenize("%getHost()%:%getPort(80)%")
		expected := []string{
			`unexpected function: "getHost": "%getHost()%"`,
			`unexpected function: "getPort": "%getPort(80)%"`,
		}
		errAssert.EqualErrorGroup(t, err, expected)
	})
}
