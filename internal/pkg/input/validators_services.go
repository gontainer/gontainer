package input

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gontainer/gontainer-helpers/container"
	"github.com/gontainer/gontainer-helpers/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/maps"
	"github.com/gontainer/gontainer/internal/pkg/ptr"
	"github.com/gontainer/gontainer/internal/pkg/regex"
	"github.com/gontainer/gontainer/internal/pkg/types"
)

var (
	regexServiceName        = regex.MustCompileAz(regex.ServiceName)
	regexServiceGetter      = regex.MustCompileAz(regex.ServiceGetter)
	regexServiceType        = regex.MustCompileAz(regex.ServiceType)
	regexServiceValue       = regex.MustCompileAz(regex.ServiceValue)
	regexServiceConstructor = regex.MustCompileAz(regex.ServiceConstructor)
	regexServiceCallName    = regex.MustCompileAz(regex.ServiceCallName)
	regexServiceFieldName   = regex.MustCompileAz(regex.ServiceFieldName)
	regexServiceTag         = regex.MustCompileAz(regex.ServiceTag)
)

type ValidateService func(Service) error

// DefaultServicesValidators returns validators for Input.Services.
func DefaultServicesValidators() []ValidateFn {
	return []ValidateFn{
		ValidateServices,
	}
}

func ValidateServices(i Input) error {
	validators := []ValidateService{
		ValidateConstructorType,
		ValidateServiceConstructor,
		ValidateServiceGetter,
		ValidateServiceType,
		ValidateServiceValue,
		ValidateServiceArgs,
		ValidateServiceCalls,
		ValidateServiceFields,
		ValidateServiceTags,
	}

	var errs []error

	for _, n := range maps.Keys(i.Services) {
		var sErrs []error
		s := i.Services[n]

		sErrs = append(sErrs, ValidateServiceName(n))
		if !ptr.Dereference(s.Todo, DefaultServiceTodo) {
			for _, v := range validators {
				sErrs = append(sErrs, v(s))
			}
		}
		errs = append(errs, grouperror.Prefix(fmt.Sprintf("%+q: ", n), sErrs...))
	}
	return grouperror.Prefix("services: ", errs...)
}

func ValidateServiceName(n string) error {
	if !regexServiceName.MatchString(n) {
		return errors.New("invalid name")
	}
	return nil
}

// ValidateConstructorType validates whether the given service has a proper constructor.
// To initialize a service, it must have either a valid constructor or a value.
func ValidateConstructorType(s Service) error {
	var errs []error

	if s.Constructor == nil && s.Value == nil && s.Type == nil {
		errs = append(errs, errors.New("missing constructor or value or type"))
	}

	if s.Constructor != nil && s.Value != nil {
		errs = append(errs, errors.New("cannot define constructor and value together"))
	}

	if len(s.Args) > 0 && s.Constructor == nil {
		errs = append(errs, errors.New("arguments are not empty, but constructor is missing"))
	}

	return grouperror.Join(errs...)
}

var reservedGetters map[string]bool

func init() {
	reservedGetters = map[string]bool{}
	r := reflect.TypeOf(container.NewSuperContainer())
	for i := 0; i < r.NumMethod(); i++ {
		reservedGetters[r.Method(i).Name] = true
	}
}

func ValidateServiceGetter(s Service) error {
	if s.Getter == nil {
		return nil
	}
	if reservedGetters[*s.Getter] {
		return fmt.Errorf("getter: %+q is reserved", *s.Getter)
	}

	var errs []error
	if strings.HasPrefix(*s.Getter, "Must") {
		errs = append(errs, errors.New(`getter: prefix "Must" is not allowed`))
	}
	if strings.HasSuffix(*s.Getter, "InContext") {
		errs = append(errs, errors.New(`getter: suffix "InContext" is not allowed`))
	}
	errs = append(errs, validateRegexField("getter", *s.Getter, regexServiceGetter))
	return grouperror.Join(errs...)
}

func ValidateServiceType(s Service) error {
	return validateOptionalPtrField("type", s.Type, regexServiceType)
}

func ValidateServiceValue(s Service) error {
	return validateOptionalPtrField("value", s.Value, regexServiceValue)
}

func ValidateServiceConstructor(s Service) error {
	return validateOptionalPtrField("constructor", s.Constructor, regexServiceConstructor)
}

func ValidateServiceArgs(s Service) error {
	var errs []error
	for i, a := range s.Args {
		if !types.IsPrimitive(a) {
			errs = append(errs, newErrUnsupportedType(fmt.Sprintf("arg %d", i), a))
		}
	}
	return grouperror.Prefix("arguments: ", errs...)
}

func ValidateServiceCalls(s Service) error {
	var errs []error
	for j, c := range s.Calls {
		var cErrs []error
		cErrs = append(
			cErrs,
			validateRegexField("method", c.Method, regexServiceCallName),
		)
		for i, a := range c.Args {
			var aErrs []error
			if !types.IsPrimitive(a) {
				aErrs = append(aErrs, newErrUnsupportedType(fmt.Sprintf("%d", i), a))
			}
			cErrs = append(cErrs, grouperror.Prefix("arguments: ", aErrs...))
		}
		errs = append(errs, grouperror.Prefix(fmt.Sprintf("%d: ", j), cErrs...))
	}
	return grouperror.Prefix("calls: ", errs...)
}

func ValidateServiceFields(s Service) error {
	var errs []error
	for _, n := range maps.Keys(s.Fields) {
		v := s.Fields[n]
		errs = append(errs, validateRegexField(fmt.Sprintf("%+q", n), n, regexServiceFieldName))
		if !types.IsPrimitive(v) {
			errs = append(errs, newErrUnsupportedType(fmt.Sprintf("%+q", n), v))
		}
	}
	return grouperror.Prefix("fields: ", errs...)
}

func ValidateServiceTags(s Service) error {
	var errs []error
	counters := make(map[string]int)
	for i, t := range s.Tags {
		errs = append(errs, validateRegexField(fmt.Sprintf("%d", i), t.Name, regexServiceTag))
		counters[t.Name]++
	}
	for _, n := range maps.Keys(counters) {
		if counters[n] > 1 {
			errs = append(errs, fmt.Errorf("duplicate %+q", n))
		}
	}
	return grouperror.Prefix("tags: ", errs...)
}
