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

package input

type Input struct {
	Version    *Version           `yaml:"version"`
	Meta       Meta               `yaml:"meta"`
	Params     map[string]any     `yaml:"parameters"`
	Services   map[string]Service `yaml:"services"`
	Decorators []Decorator        `yaml:"decorators"`
}

type Version string

type Meta struct {
	Pkg                  *string           `yaml:"pkg"`                   // default "main"
	ContainerType        *string           `yaml:"container_type"`        // default "Gontainer"
	ContainerConstructor *string           `yaml:"container_constructor"` // default "NewGontainer"
	DefaultMustGetter    *bool             `yaml:"default_must_getter"`   // default false
	Imports              map[string]string `yaml:"imports"`               // {"alias": "my/long/path", ...}
	Functions            map[string]string `yaml:"functions"`             // {"env": "os.Getenv", ...}
}

type Scope uint

type Service struct {
	Getter      *string        `yaml:"getter"`      // e.g. GetDB
	MustGetter  *bool          `yaml:"must_getter"` // if true adds MustGet func, it requires non-empty Getter
	Type        *string        `yaml:"type"`        // *?my/import/path.Type
	Value       *string        `yaml:"value"`       // my/import/path.Variable
	Constructor *string        `yaml:"constructor"` // NewDB
	Args        []any          `yaml:"arguments"`   // ["%host%", "%port%", "@logger"]
	Calls       []Call         `yaml:"calls"`       // [["SetLogger", ["@logger"]], ...]
	Fields      map[string]any `yaml:"fields"`      // Field: "%value%"
	Tags        []Tag          `yaml:"tags"`        // ["service_decorator", ...]
	Scope       *Scope         `yaml:"scope"`       // scope for the service, if nil, default value will be applied
	Todo        *bool          `yaml:"todo"`        // if true skips validation and returns error whenever users asks container for a service
}

type Tag struct {
	Name     string
	Priority int
}

type Call struct {
	Method    string
	Args      []any
	Immutable bool
}

type Decorator struct {
	Tag       string `yaml:"tag"`       // http-client
	Decorator string `yaml:"decorator"` // myImport/pkgCopy.MakeTracedHttpClient
	Args      []any  `yaml:"arguments"` // ["@logger"]
}
