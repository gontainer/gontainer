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
	"fmt"

	"github.com/gontainer/gontainer-helpers/v3/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/maps"
	"github.com/gontainer/gontainer/internal/pkg/regex"
	"github.com/gontainer/gontainer/internal/pkg/types"
)

var (
	regexParamName = regex.MustCompileAz(regex.ParamName)
)

// DefaultParamsValidators returns validators for Input.Params.
func DefaultParamsValidators() []ValidateFn {
	return []ValidateFn{
		ValidateParams,
	}
}

func ValidateParams(i Input) error {
	var errs []error
	for _, n := range maps.Keys(i.Params) {
		if !regexParamName.MatchString(n) {
			errs = append(errs, fmt.Errorf("%+q: invalid name", n))
		}

		v := i.Params[n]

		if !types.IsPrimitive(v) {
			errs = append(errs, newErrUnsupportedType(fmt.Sprintf("%+q", n), v))
		}
	}
	return grouperror.Prefix("parameters: ", errs...)
}
