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

package regex

const (
	BaseImport = `[A-Za-z](\/?[A-Z-a-z0-9._-])*`
	Import     = `((` + BaseImport + `)|("` + BaseImport + `")|"\.")`
	GoToken    = `[A-Za-z][A-Za-z0-9_]*`
	YamlToken  = `[A-Za-z]((\.|-|_)?[A-Za-z0-9])*`
	GoFunc     = `((?P<import>` + Import + `)\.)?(?P<fn>` + GoToken + `)`

	MetaPkg                  = GoToken
	MetaContainerType        = GoToken
	MetaContainerConstructor = GoToken
	MetaImport               = Import
	MetaImportAlias          = YamlToken
	MetaFn                   = GoToken
	MetaGoFn                 = GoFunc

	ParamName = YamlToken

	ServiceName   = YamlToken
	ServiceGetter = GoToken
	ServiceType   = `(?P<ptr>\*)?` + `((?P<import>` + Import + `)\.)?(?P<type>` + GoToken + `)`
	// ServiceValue supports
	// Value
	// &Value
	// my/import.Value
	// "my/import".GlobalVar.Field
	// &"my/import".GlobalVar.Field
	// MyStruct{}
	// &MyStruct{}
	// my/import.MyStruct{}
	// &my/import.MyStruct{}
	ServiceValue = `((?P<v1>(?P<ptr>\&)?((?P<import>` + Import + `)\.)?(?P<value>` + GoToken + `(\.` + GoToken + `)*` + `))` +
		`|(?P<v2>(?P<ptr2>\&)?((?P<import2>` + Import + `)\.)?(?P<struct2>` + GoToken + `)\{\}))`
	ServiceConstructor = GoFunc
	ServiceCallName    = GoToken
	ServiceFieldName   = GoToken
	ServiceTag         = YamlToken

	PrefixService = `@`
	ArgService    = PrefixService + "(?P<service>" + YamlToken + ")"
	PrefixTagged  = `!tagged\s+`
	ArgTagged     = PrefixTagged + "(?P<tag>" + YamlToken + ")"
	PrefixValue   = `!value\s+`
	ArgValue      = PrefixValue + "((?P<argval>" + ServiceValue + "))"

	DecoratorTag    = `(\*|(` + ServiceTag + `))`
	DecoratorMethod = GoFunc
)
