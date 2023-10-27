package output

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
)

func ValidateParamsExist(o Output) error {
	existing := make(map[string]struct{}, len(o.Params))
	for _, p := range o.Params {
		existing[p.Name] = struct{}{}
	}

	var errs []error
	errs = append(errs, validateParamsExistsInParams(existing, o.Params)...)
	errs = append(errs, validateParamsExistsInServices(existing, o.Services)...)

	return grouperror.Prefix("output.ValidateParamsExist: ", errs...)
}

func validateParamsExistsInParams(existing map[string]struct{}, params []Param) []error {
	var errs []error

	for _, p := range params {
		for _, n := range p.DependsOn {
			if _, ok := existing[n]; !ok {
				errs = append(errs, fmt.Errorf("%+q: param %+q does not exist", "%"+p.Name+"%", n))
			}
		}
	}

	return errs
}

func validateParamsExistsInServices(existing map[string]struct{}, services []Service) []error {
	var errs []error

	for _, s := range services {
		for _, a := range s.AllArgs() {
			for _, n := range a.DependsOnParams {
				if _, ok := existing[n]; !ok {
					errs = append(errs, fmt.Errorf("%+q: param %+q does not exist", "@"+s.Name, n))
				}
			}
		}
	}

	return errs
}
