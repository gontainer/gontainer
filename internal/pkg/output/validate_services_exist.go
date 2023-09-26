package output

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/errors"
)

func ValidateServicesExist(o Output) error {
	names := make(map[string]struct{}, len(o.Services))
	for _, s := range o.Services {
		names[s.Name] = struct{}{}
	}

	var errs []error

	for _, s := range o.Services {
		for _, a := range s.AllArgs() {
			for _, n := range a.DependsOnServices {
				if _, ok := names[n]; !ok {
					errs = append(errs, fmt.Errorf("%+q: service %+q does not exist", s.Name, n))
				}
			}
		}
	}

	for dID, d := range o.Decorators {
		for _, a := range d.Args {
			for _, n := range a.DependsOnServices {
				if _, ok := names[n]; !ok {
					errs = append(errs, fmt.Errorf("decorator(#%d, %+q): service %+q does not exist", dID, d.Tag, n))
				}
			}
		}
	}

	return errors.PrefixedGroup("output.ValidateServicesExist: ", errs...)
}
