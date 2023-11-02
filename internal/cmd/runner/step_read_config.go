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

package runner

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gontainer/gontainer-helpers/v3/exporter"
	"github.com/gontainer/gontainer-helpers/v3/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"gopkg.in/yaml.v3"
)

type StepReadConfig struct {
	printer  printer
	patterns []string
}

func NewStepReadConfig(printer printer, patterns []string) *StepReadConfig {
	return &StepReadConfig{printer: printer, patterns: patterns}
}

func (s *StepReadConfig) Name() string {
	return "Read config"
}

func (s *StepReadConfig) Run(i *input.Input, _ *output.Output) (err error) {
	defer func() {
		err = grouperror.Prefix("runner.StepReadConfig: ", err)
	}()

	if len(s.patterns) == 0 {
		return errors.New("missing file patterns")
	}

	processed := make(map[string][]string)

	var errs []error
	found := false
	s.printer.Println("Patterns")
	for j, p := range s.patterns {
		s.printer.Println(fmt.Sprintf("%d. %s", j+1, p))
		files, err := s.findFiles(p)
		errs = append(errs, err)
		if len(files) == 0 {
			s.printer.Println("   No files")
		}
		for _, f := range files {
			var pErrs []error

			func() {
				defer func() {
					errs = append(errs, grouperror.Prefix(fmt.Sprintf("`%s`: ", f), pErrs...))
				}()

				noErrors := false

				defer func() {
					mark := checkMark
					if !noErrors {
						mark = xMark
					}
					s.printer.Println(fmt.Sprintf("   • %s %s", f, mark))
				}()

				buff, err := os.ReadFile(f)
				if err != nil {
					pErrs = append(pErrs, grouperror.Prefix("could not read the file: ", err))
					return
				}

				tmp := input.Input{}
				if err := yaml.Unmarshal(buff, &tmp); err != nil {
					pErrs = append(pErrs, grouperror.Prefix("parsing yaml: ", err))
					return
				}
				*i = input.Merge(*i, tmp)
				found = true
				noErrors = true
				processed[f] = append(processed[f], p)
			}()
		}
	}
	if !found {
		errs = append(errs, errors.New("could not process any files"))
	}

	for f, p := range processed {
		if len(p) > 1 {
			tmpPatterns := fmt.Sprintf("%#v", p)
			tmpPatterns = strings.TrimPrefix(tmpPatterns, "[]string")
			errs = append(errs, fmt.Errorf("file %+q matches more than one pattern: %s", f, tmpPatterns))
		}
	}

	err = grouperror.Join(errs...)
	return
}

// findFiles returns list of files found by given pattern.
// Names are returned in the lexical order.
func (s *StepReadConfig) findFiles(pattern string) ([]string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, grouperror.Prefix(
			fmt.Sprintf("pattern: %s: ", exporter.MustExport(pattern)),
			err,
		)
	}
	for i, m := range matches {
		matches[i] = filepath.Clean(m)
	}
	sort.Strings(matches)
	return matches, nil
}
