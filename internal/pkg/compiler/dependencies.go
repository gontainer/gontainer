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

package compiler

import (
	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

type aliasRegisterer interface {
	RegisterPrefixAlias(alias string, import_ string) error
}

type funcRegisterer interface {
	RegisterFunc(fnAlias string, goImport string, goFn string)
}

type paramResolver interface {
	ResolveParam(any) (resolver.ParamExpr, error)
}

type argResolver interface {
	ResolveArg(any) (resolver.ArgExpr, error)
}

type aliaser interface {
	// Alias returns an alias for given import, e.g. "github.com/spf13/viper" => "i0_viper".
	Alias(string) string
}
