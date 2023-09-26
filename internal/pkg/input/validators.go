package input

import (
	"github.com/gontainer/gontainer-helpers/errors"
)

type ValidateFn func(Input) error

type Validator struct {
	validators []ValidateFn
}

func NewValidator(validators ...ValidateFn) *Validator {
	return &Validator{validators: validators}
}

func (v Validator) Validate(m Input) error {
	var errs []error
	for _, v := range v.validators {
		if err := v(m); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Group(errs...)
}

// NewDefaultValidator expects a version that follows semver schema without the leading "v".
func NewDefaultValidator(version string) *Validator {
	var sv []ValidateFn
	sv = append(sv, DefaultVersionValidators(version)...)
	sv = append(sv, DefaultMetaValidators()...)
	sv = append(sv, DefaultParamsValidators()...)
	sv = append(sv, DefaultServicesValidators()...)
	sv = append(sv, DefaultDecoratorsValidators()...)
	return NewValidator(sv...)
}
