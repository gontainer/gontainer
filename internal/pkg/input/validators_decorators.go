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

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/regex"
	"github.com/gontainer/gontainer/internal/pkg/types"
)

var (
	regexDecoratorsTag   = regex.MustCompileAz(regex.DecoratorTag)
	regexDecoratorMethod = regex.MustCompileAz(regex.DecoratorMethod)
)

// DefaultDecoratorsValidators returns validators for Input.Decorators.
func DefaultDecoratorsValidators() []ValidateFn {
	return []ValidateFn{
		ValidateDecorators,
	}
}

func ValidateDecorators(i Input) error {
	var errs []error
	for j, d := range i.Decorators {
		g := []error{
			ValidateDecoratorTag(d),
			ValidateDecoratorMethod(d),
			ValidateDecoratorArgs(d),
		}
		errs = append(
			errs,
			grouperror.Prefix(fmt.Sprintf("%d %+q: ", j, d.Decorator), g...),
		)
	}
	return grouperror.Prefix("decorators: ", errs...)
}

func ValidateDecoratorTag(d Decorator) error {
	return validateRegexField("tag", d.Tag, regexDecoratorsTag)
}

func ValidateDecoratorMethod(d Decorator) error {
	return validateRegexField("method", d.Decorator, regexDecoratorMethod)
}

func ValidateDecoratorArgs(d Decorator) error {
	var errs []error
	for i, a := range d.Args {
		if !types.IsPrimitive(a) {
			errs = append(errs, newErrUnsupportedType(fmt.Sprintf("%d", i), a))
		}
	}
	return grouperror.Prefix("arguments: ", errs...)
}
