package output

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/grouperror"
)

func ValidateServicesExist(o Output) error {
	existing := make(map[string]struct{}, len(o.Services))
	for _, s := range o.Services {
		existing[s.Name] = struct{}{}
	}

	var errs []error
	errs = append(errs, validateServicesExistsInServices(existing, o.Services)...)
	errs = append(errs, validateServicesExistsInDecorators(existing, o.Decorators)...)

	return grouperror.Prefix("output.ValidateServicesExist: ", errs...)
}

func validateServicesExistsInServices(existing map[string]struct{}, services []Service) []error {
	var errs []error

	for _, s := range services {
		for _, a := range s.AllArgs() {
			for _, n := range a.DependsOnServices {
				if _, ok := existing[n]; !ok {
					errs = append(errs, fmt.Errorf("%+q: service %+q does not exist", s.Name, n))
				}
			}
		}
	}

	return errs
}

func validateServicesExistsInDecorators(existing map[string]struct{}, decorators []Decorator) []error {
	var errs []error

	for dID, d := range decorators {
		for _, a := range d.Args {
			for _, n := range a.DependsOnServices {
				if _, ok := existing[n]; !ok {
					errs = append(errs, fmt.Errorf("decorator(#%d, %+q): service %+q does not exist", dID, d.Tag, n))
				}
			}
		}
	}

	return errs
}
