package input

import (
	"fmt"
	"regexp"
)

func validateOptionalPtrField(field string, val *string, r *regexp.Regexp) error {
	if val == nil {
		return nil
	}
	return validateRegexField(field, *val, r)
}

func validateRegexField(field string, v string, r *regexp.Regexp) error {
	if !r.MatchString(v) {
		return fmt.Errorf("%s: invalid %+q", field, v)
	}
	return nil
}

func newErrUnsupportedType(name string, val any) error {
	return fmt.Errorf("%s: unsupported type %T", name, val)
}
