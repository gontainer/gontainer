package output

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/errors"
)

func ValidateParamsExist(o Output) error {
	names := make(map[string]struct{}, len(o.Params))
	for _, p := range o.Params {
		names[p.Name] = struct{}{}
	}

	var errs []error

	for _, p := range o.Params {
		for _, n := range p.DependsOn {
			if _, ok := names[n]; !ok {
				errs = append(errs, fmt.Errorf("%+q: param %+q does not exist", p.Name, n))
			}
		}
	}

	return errors.PrefixedGroup("output.ValidateParamsExist: ", errs...)
}
