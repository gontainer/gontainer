package input

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/grouperror"
	"github.com/gontainer/gontainer/internal/pkg/regex"
)

var (
	regexpMetaPkg                  = regex.MustCompileAz(regex.MetaPkg)
	regexpMetaContainerType        = regex.MustCompileAz(regex.MetaContainerType)
	regexpMetaContainerConstructor = regex.MustCompileAz(regex.MetaContainerConstructor)
	regexMetaImport                = regex.MustCompileAz(regex.MetaImport)
	regexMetaImportAlias           = regex.MustCompileAz(regex.MetaImportAlias)
	regexMetaFn                    = regex.MustCompileAz(regex.MetaFn)
	regexMetaGoFn                  = regex.MustCompileAz(regex.MetaGoFn)
)

// DefaultMetaValidators returns validators for Input.Meta.
func DefaultMetaValidators() []ValidateFn {
	return []ValidateFn{
		ValidateMeta,
	}
}

func ValidateMeta(i Input) error {
	vals := []func(Meta) error{
		ValidateMetaPkg,
		ValidateMetaContainerType,
		ValidateMetaContainerConstructor,
		ValidateMetaImports,
		ValidateMetaFunctions,
	}
	var errs []error
	for _, v := range vals {
		errs = append(errs, v(i.Meta))
	}
	return grouperror.Prefix("meta: ", errs...)
}

func ValidateMetaPkg(m Meta) error {
	return validateOptionalPtrField("pkg", m.Pkg, regexpMetaPkg)
}

func ValidateMetaContainerType(m Meta) error {
	if m.ContainerType != nil && !regexpMetaContainerType.MatchString(*m.ContainerType) {
		return fmt.Errorf(
			"container_type: invalid %+q",
			*m.ContainerType,
		)
	}
	return nil
}

func ValidateMetaContainerConstructor(m Meta) error {
	if m.ContainerConstructor != nil && !regexpMetaContainerConstructor.MatchString(*m.ContainerConstructor) {
		return fmt.Errorf(
			"container_constructor: invalid %+q",
			*m.ContainerConstructor,
		)
	}
	return nil
}

func ValidateMetaImports(m Meta) error {
	var errs []error
	for a, imp := range m.Imports {
		if !regexMetaImport.MatchString(imp) {
			errs = append(errs, fmt.Errorf("invalid import %+q", imp))
		}
		if !regexMetaImportAlias.MatchString(a) {
			errs = append(errs, fmt.Errorf("invalid alias %+q", a))
		}
	}
	return grouperror.Prefix("imports: ", errs...)
}

func ValidateMetaFunctions(m Meta) error {
	var errs []error
	for fn, goFn := range m.Functions {
		if !regexMetaFn.MatchString(fn) {
			errs = append(errs, fmt.Errorf("invalid function %+q", fn))
		}

		if !regexMetaGoFn.MatchString(goFn) {
			errs = append(errs, fmt.Errorf("invalid go function %+q", goFn))
		}
	}
	return grouperror.Prefix("functions: ", errs...)
}
