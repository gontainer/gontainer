package input

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/errors"
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
	return errors.PrefixedGroup("parameters: ", errs...)
}
