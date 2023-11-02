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

package input

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/v3/grouperror"
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
