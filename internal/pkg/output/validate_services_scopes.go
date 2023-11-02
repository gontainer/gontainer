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
	"github.com/gontainer/gontainer/internal/pkg/maps"
)

func ValidateServicesScopes(o Output) error {
	var errs []error

	services := make(map[string]Service)
	for _, s := range o.Services {
		services[s.Name] = s
	}

	graph := o.BuildDependencyGraph()

	maps.Iterate(services, func(_ string, s Service) {
		if s.Scope != ScopeShared {
			return
		}

		deps := graph.Deps(s.Name)
		for _, d := range deps {
			if !d.IsService() {
				continue
			}
			s2 := services[d.Resource]
			if s2.Scope == ScopeContextual {
				errs = append(
					errs,
					fmt.Errorf("%+q: service is shared, but dependant %+q is contextual", s.Name, s2.Name),
				)
			}
		}
	})

	return grouperror.Prefix("output.ValidateServicesScopes: ", errs...)
}
