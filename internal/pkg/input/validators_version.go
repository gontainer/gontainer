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

import (
	"errors"
	"fmt"

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"golang.org/x/mod/semver"
)

// DefaultVersionValidators returns validators for Input.Version.
func DefaultVersionValidators(version string) []ValidateFn {
	return []ValidateFn{
		NewVersionValidator(version).ValidateVersion,
	}
}

type VersionValidator struct {
	version string
	valid   bool
}

// NewVersionValidator expects a proper version without the leading "v".
func NewVersionValidator(version string) *VersionValidator {
	ver := version
	val := false

	if semver.IsValid("v" + version) {
		ver = "v" + version
		val = true
	}

	return &VersionValidator{
		version: ver,
		valid:   val,
	}
}

func (v *VersionValidator) ValidateVersion(i Input) (err error) {
	if i.Version == nil || !v.valid {
		return nil
	}

	defer func() {
		err = grouperror.Prefix(fmt.Sprintf("version: current: %s, given: %s: ", v.version, *i.Version), err)
	}()

	curr := semver.MajorMinor(v.version) + ".0"
	given := semver.MajorMinor(string(*i.Version)) + ".0"

	if semver.Major(v.version) == "v0" {
		if curr != given {
			return errors.New("possibly incompatible versions")
		}
		return
	}

	if semver.Major(v.version) != semver.Major(string(*i.Version)) {
		return errors.New("incompatible versions")
	}

	if semver.Compare(curr, given) < 0 {
		return errors.New("update Gontainer to use all new features")
	}

	return
}
