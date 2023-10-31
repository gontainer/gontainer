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

package compiler_test

import (
	"fmt"
	"testing"

	errAssert "github.com/gontainer/gontainer-helpers/v2/grouperror/assert"
	"github.com/gontainer/gontainer/internal/pkg/compiler"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/stretchr/testify/assert"
)

func TestStepCompileMeta_Process(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		Name     string
		Input    input.Meta
		Output   output.Meta
		Compiler *compiler.StepCompileMeta
		Errors   []string
	}{
		{
			Name:  "Default values",
			Input: input.Meta{},
			Output: output.Meta{
				Pkg:                  "main",
				ContainerType:        "Gontainer",
				ContainerConstructor: "NewGontainer",
			},
			Compiler: compiler.NewStepCompileMeta(nilAliasRegisterer, nilFuncRegisterer),
		},
		{
			Name: "Errors",
			Input: input.Meta{
				Imports: map[string]string{
					"http":   "net/http",
					"assert": "github.com/stretchr/testify/assert",
				},
				Functions: map[string]string{
					"env": "os.GetEnv",
				},
			},
			Output: output.Meta{
				Pkg:                  "main",
				ContainerType:        "Gontainer",
				ContainerConstructor: "NewGontainer",
			},
			Errors: []string{
				`compiler.StepCompileMeta: imports: could not register alias "assert" for import "github.com/stretchr/testify/assert"`,
				`compiler.StepCompileMeta: imports: could not register alias "http" for import "net/http"`,
			},
			Compiler: compiler.NewStepCompileMeta(
				aliasRegistererFunc(func(alias string, import_ string) error {
					return fmt.Errorf("could not register alias %+q for import %+q", alias, import_)
				}),
				nilFuncRegisterer,
			),
		},
	}

	for _, tmp := range scenarios {
		s := tmp
		t.Run(s.Name, func(t *testing.T) {
			t.Parallel()
			o := output.Output{}
			err := s.Compiler.Process(input.Input{Meta: s.Input}, &o)
			errAssert.EqualErrorGroup(t, err, s.Errors)
			assert.Equal(t, s.Output, o.Meta)
		})
	}

	t.Run("Register imports & funcs", func(t *testing.T) {
		t.Parallel()
		o := output.Output{}
		i := input.Input{}
		i.Meta.Functions = map[string]string{
			"yamlFunc1": "goFunc1",
			"yamlFunc2": "pkg.goFunc2",
		}
		i.Meta.Imports = map[string]string{
			"alias1": "path1",
			"alias2": "pkg.path2",
		}

		aliases := make(map[string]string)
		funcs := make(map[string][2]string) // alias => [import, fn]

		cmplr := compiler.NewStepCompileMeta(
			aliasRegistererFunc(func(alias string, import_ string) error {
				aliases[alias] = import_
				return nil
			}),
			funcRegistererFunc(func(fnAlias string, goImport string, goFn string) {
				funcs[fnAlias] = [2]string{goImport, goFn}
			}),
		)
		assert.NoError(t, cmplr.Process(i, &o))

		assert.Equal(
			t,
			map[string]string{
				"alias1": "path1",
				"alias2": "pkg.path2",
			},
			aliases,
		)
		assert.Equal(
			t,
			map[string][2]string{
				"yamlFunc1": {"", "goFunc1"},
				"yamlFunc2": {"pkg", "goFunc2"},
			},
			funcs,
		)
	})
}
