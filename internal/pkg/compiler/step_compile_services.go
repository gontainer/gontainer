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

package compiler

import (
	"errors"
	"fmt"

	"github.com/gontainer/gontainer-helpers/v3/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/maps"
	"github.com/gontainer/gontainer/internal/pkg/output"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/gontainer/gontainer/internal/pkg/regex"
	"github.com/gontainer/gontainer/internal/pkg/syntax"
)

var (
	regexServiceType        = regex.MustCompileAz(regex.ServiceType)
	regexServiceConstructor = regex.MustCompileAz(regex.ServiceConstructor)
)

type StepCompileServices struct {
	aliaser     aliaser
	argResolver argResolver
}

func NewStepCompileServices(a aliaser, ar argResolver) *StepCompileServices {
	return &StepCompileServices{aliaser: a, argResolver: ar}
}

func (s StepCompileServices) Process(i input.Input, o *output.Output) error {
	errs := make([]error, 0, len(i.Services))
	o.Services = make([]output.Service, 0, len(i.Services))
	maps.Iterate(i.Services, func(k string, v input.Service) {
		svc, err := s.processService(k, i)
		o.Services = append(o.Services, svc)
		errs = append(errs, err)
	})
	s.processScopes(i, o)
	return grouperror.Prefix("compiler.StepCompileServices: ", errs...)
}

// processScopes computes the proper scope for services. When the scope is not defined,
// but the service has a contextual dependency,
// the scope contextual will be assigned to the service instead of the default one.
func (s StepCompileServices) processScopes(i input.Input, o *output.Output) {
	for j, svc := range o.Services {
		iScope := i.Services[svc.Name].Scope
		if iScope == nil {
			o.Services[j].Scope = output.ScopeDefault
			continue
		}

		switch *iScope {
		case input.ScopeShared:
			o.Services[j].Scope = output.ScopeShared
		case input.ScopeContextual:
			o.Services[j].Scope = output.ScopeContextual
		case input.ScopeNonShared:
			o.Services[j].Scope = output.ScopeNonShared
		}
	}
}

func (s StepCompileServices) processService(name string, i input.Input) (o output.Service, _ error) {
	svc := i.Services[name]

	if ptr.Dereference(svc.Todo, defaultServiceTodo) {
		o.Name = name
		o.Todo = true
		return o, nil
	}

	var errs []error

	fields, err := s.serviceFields(svc.Fields)
	errs = append(errs, err)

	args, err := resolveArgs(s.argResolver, svc.Args)
	errs = append(errs, err)

	calls, err := s.serviceCalls(svc.Calls)
	errs = append(errs, err)

	getter, mustGetter, err := s.getter(svc, i.Meta)
	if err != nil {
		errs = append(errs, err)
	}

	return output.Service{
		Name:        name,
		Getter:      getter,
		MustGetter:  mustGetter,
		Type:        s.serviceType(svc.Type),
		Value:       s.serviceValue(svc.Value),
		Constructor: s.serviceConstructor(svc.Constructor),
		Args:        args,
		Calls:       calls,
		Fields:      fields,
		Tags:        s.serviceTags(svc.Tags),
		Todo:        false,
	}, grouperror.Prefix(fmt.Sprintf("%+q: ", name), errs...)
}

func (s StepCompileServices) getter(svc input.Service, m input.Meta) (getter string, mustGetter bool, err error) {
	getter = ptr.Dereference(svc.Getter, defaultServiceGetter)
	mustGetter = ptr.Dereference(svc.MustGetter, ptr.Dereference(m.DefaultMustGetter, defaultMetaMustGetter))
	if getter == "" && svc.MustGetter != nil && mustGetter {
		err = errors.New("cannot generate a must-getter when the getter is not specified")
	}
	if getter == "" && svc.MustGetter == nil && mustGetter {
		mustGetter = false
	}
	return
}

func (s StepCompileServices) serviceType(serviceType *string) string {
	if serviceType == nil {
		return "interface{}"
	}
	_, m := regex.Match(regexServiceType, *serviceType)
	t := m["type"]
	import_ := syntax.SanitizeImport(m["import"])
	if import_ != "" {
		t = s.aliaser.Alias(import_) + "." + t
	}
	return m["ptr"] + t
}

func (s StepCompileServices) serviceValue(serviceValue *string) string {
	if serviceValue == nil {
		return ""
	}

	return syntax.CompileServiceValue(s.aliaser, *serviceValue)
}

func (s StepCompileServices) serviceConstructor(c *string) string {
	if c == nil {
		return ""
	}
	_, m := regex.Match(regexServiceConstructor, *c)
	import_ := syntax.SanitizeImport(m["import"])
	r := ""
	if import_ != "" {
		r = s.aliaser.Alias(import_) + "."
	}
	return r + m["fn"]
}

func (s StepCompileServices) serviceCalls(calls []input.Call) (r []output.Call, _ error) {
	var errs []error
	if len(calls) > 0 {
		r = make([]output.Call, len(calls))
	}

	for i, call := range calls {
		args, err := resolveArgs(s.argResolver, call.Args)
		errs = append(errs, grouperror.Prefix(fmt.Sprintf("%d: ", i), err))
		r[i] = output.Call{
			Method:    call.Method,
			Args:      args,
			Immutable: call.Immutable,
		}
	}

	return r, grouperror.Prefix("calls: ", errs...)
}

func (s StepCompileServices) serviceFields(fields map[string]any) (r []output.Field, _ error) {
	var errs []error
	maps.Iterate(fields, func(n string, v any) {
		arg, err := s.argResolver.ResolveArg(v)
		r = append(r, output.Field{
			Name:  n,
			Value: argExprToArg(arg),
		})
		if err != nil {
			errs = append(errs, grouperror.Prefix(fmt.Sprintf("%+q: ", n), err))
		}
	})
	return r, grouperror.Prefix("fields: ", errs...)
}

func (s StepCompileServices) serviceTags(tags []input.Tag) (r []output.Tag) {
	for _, t := range tags {
		r = append(r, output.Tag{
			Name:     t.Name,
			Priority: t.Priority,
		})
	}
	return
}
