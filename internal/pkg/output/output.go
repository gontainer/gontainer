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

package output

type Param struct {
	Name      string
	Code      string
	Raw       any
	DependsOn []string
}

type Arg struct {
	Code              string
	Raw               any
	DependsOnParams   []string
	DependsOnServices []string
	DependsOnTags     []string
}

type Call struct {
	Method    string
	Args      []Arg
	Immutable bool
}

type Field struct {
	Name  string
	Value Arg
}

type Tag struct {
	Name     string
	Priority int
}

type Scope uint

const (
	ScopeDefault Scope = iota
	ScopeShared
	ScopeContextual
	ScopeNonShared
)

func (s Scope) IsDefault() bool {
	return s == ScopeDefault
}

func (s Scope) IsShared() bool {
	return s == ScopeShared
}

func (s Scope) IsContextual() bool {
	return s == ScopeContextual
}

func (s Scope) IsNonShared() bool {
	return s == ScopeNonShared
}

type Service struct {
	Name        string
	Getter      string
	MustGetter  bool
	Type        string
	Value       string
	Constructor string
	Args        []Arg
	Calls       []Call
	Fields      []Field
	Tags        []Tag
	Scope       Scope
	Todo        bool
}

type Decorator struct {
	Tag       string
	Decorator string
	Args      []Arg
	Raw       string
}

type Meta struct {
	Pkg                  string
	ContainerType        string
	ContainerConstructor string
}

type Output struct {
	Meta       Meta
	Params     []Param
	Services   []Service
	Decorators []Decorator
}

// AllArgs returns arguments passed to constructor, calls and fields to fetch information about all dependencies.
// It does not include arguments passed to related decorators.
func (s Service) AllArgs() []Arg {
	var res []Arg
	res = append(res, s.Args...)
	for _, c := range s.Calls {
		res = append(res, c.Args...)
	}
	for _, f := range s.Fields {
		res = append(res, f.Value)
	}
	return res
}
