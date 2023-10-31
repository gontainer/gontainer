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

package imports

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	regexNoAlphaNum = regexp.MustCompile("[^a-zA-Z0-9]")
)

type Import struct {
	Alias string // e.g. "viper"
	Path  string // e.g. "github.com/spf13/viper"
}

type imports struct {
	counter  int64
	imports  map[string]string // map[string]string{"viper": "i0_spf13_viper", "github.com/spf13/viper": "i0_spf13_viper"}
	prefixes map[string]string // map[string]string{"viper": "github.com/spf13/viper"}
}

func New() *imports {
	return &imports{
		imports:  make(map[string]string),
		prefixes: make(map[string]string),
	}
}

// RegisterPrefixAlias registers the given alias for the given path, e.g.:
//
//	i.RegisterPrefixAlias("viper", "github.com/spf13/viper")
func (i *imports) RegisterPrefixAlias(alias string, path string) error {
	if _, ok := i.prefixes[alias]; ok {
		return fmt.Errorf("prefix is already registered: %+q", alias)
	}

	i.prefixes[alias] = path
	return nil
}

// Alias generates an alias for given path and adds path to collection of all imports.
// See Imports.
func (i *imports) Alias(import_ string) string {
	import_ = i.decorateImport(import_)

	if imp, ok := i.imports[import_]; ok {
		return imp
	}

	parts := strings.Split(import_, "/")

	alias := parts[len(parts)-1]
	alias = regexNoAlphaNum.ReplaceAllString(alias, "_")
	alias = fmt.Sprintf("i%s_%s", strconv.FormatInt(i.counter, 16), alias)

	i.imports[import_] = alias
	i.counter++

	return alias
}

// Imports returns all imports in order of using them.
func (i *imports) Imports() []Import {
	imps := make([]Import, 0, len(i.imports))
	for imp, alias := range i.imports {
		imps = append(imps, Import{
			Alias: alias,
			Path:  imp,
		})
	}

	sort.SliceStable(imps, func(i, j int) bool {
		return imps[i].Path < imps[j].Path
	})
	return imps
}

func (i *imports) decorateImport(imp string) string {
	for shortcut, path := range i.prefixes {
		if strings.Index(imp, shortcut) == 0 {
			return strings.Replace(imp, shortcut, path, 1)
		}
	}

	return imp
}
