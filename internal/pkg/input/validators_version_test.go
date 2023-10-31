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

package input_test

import (
	"fmt"
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/stretchr/testify/assert"
)

func TestVersionValidator_ValidateVersion(t *testing.T) {
	t.Parallel()

	scenarios := []struct {
		libraryVer string
		givenVer   *input.Version
		error      string
	}{
		{
			libraryVer: "0.1.0",
			givenVer:   ptr.New(input.Version("v0.1.1")),
			error:      "",
		},
		{
			libraryVer: "1.0.0-alpha.1",
			givenVer:   ptr.New(input.Version("v1.0.0")),
			error:      "",
		},
		{
			libraryVer: "dev-master",
			givenVer:   ptr.New(input.Version("v999.999.999")),
			error:      "",
		},
		{
			libraryVer: "0.3.0",
			givenVer:   ptr.New(input.Version("v0.2.0")),
			error:      "version: current: v0.3.0, given: v0.2.0: possibly incompatible versions",
		},
		{
			libraryVer: "1.3.0",
			givenVer:   ptr.New(input.Version("v1.2.0")),
			error:      "",
		},
		{
			libraryVer: "1.2.0",
			givenVer:   ptr.New(input.Version("v1.3.0")),
			error:      "version: current: v1.2.0, given: v1.3.0: update Gontainer to use all new features",
		},
		{
			libraryVer: "1.0.0",
			givenVer:   ptr.New(input.Version("v2.0.0")),
			error:      "version: current: v1.0.0, given: v2.0.0: incompatible versions",
		},
		{
			libraryVer: "1.0.0",
			givenVer:   ptr.New(input.Version("v0.999.0")),
			error:      "version: current: v1.0.0, given: v0.999.0: incompatible versions",
		},
	}

	for i, tmp := range scenarios {
		s := tmp
		t.Run(fmt.Sprintf("scenario #%d", i), func(t *testing.T) {
			t.Parallel()
			i := input.Input{
				Version: s.givenVer,
			}

			t.Run("input.NewVersionValidator", func(t *testing.T) {
				t.Parallel()
				err := input.NewVersionValidator(s.libraryVer).ValidateVersion(i)
				if s.error == "" {
					assert.NoError(t, err)
					return
				}
				assert.EqualError(t, err, s.error)
			})
			t.Run("input.NewDefaultValidator", func(t *testing.T) {
				t.Parallel()
				err := input.NewDefaultValidator(s.libraryVer).Validate(i)
				if s.error == "" {
					assert.NoError(t, err)
					return
				}
				assert.EqualError(t, err, s.error)
			})
		})
	}
}
