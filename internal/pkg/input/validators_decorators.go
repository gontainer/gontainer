package input

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/errors"
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
			errors.PrefixedGroup(fmt.Sprintf("%d %+q: ", j, d.Decorator), g...),
		)
	}
	return errors.PrefixedGroup("decorators: ", errs...)
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
	return errors.PrefixedGroup("arguments: ", errs...)
}
