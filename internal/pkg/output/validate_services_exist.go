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

package output

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/v3/grouperror"
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
