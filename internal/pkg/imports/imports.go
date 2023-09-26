package imports

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gontainer/gontainer/internal/pkg/slices"
)

var (
	regexNoAlphaNum = regexp.MustCompile("[^a-zA-Z0-9]")
)

type Import struct {
	Alias string // e.g. "viper"
	Path  string // e.g. "github.com/spf13/viper"
}

type imports struct {
	counter      int64
	imports      map[string]string // map[string]string{"viper": "i0_spf13_viper", "github.com/spf13/viper": "i0_spf13_viper"}
	importsSlice []Import
	prefixes     map[string]string // map[string]string{"viper": "github.com/spf13/viper"}
}

func New() *imports {
	return &imports{
		imports:  make(map[string]string),
		prefixes: make(map[string]string),
	}
}

// RegisterPrefixAlias registers the given alias for the given path, e.g.:
//
//	i.RegisterPrefixAlias("viper", "github.com/spf13/viper")
func (i *imports) RegisterPrefixAlias(alias string, path string) error {
	if _, ok := i.prefixes[alias]; ok {
		return fmt.Errorf("prefix is already registered: %+q", alias)
	}

	i.prefixes[alias] = path
	return nil
}

// Alias generates an alias for given path and adds path to collection of all imports.
// See Imports.
func (i *imports) Alias(import_ string) string {
	import_ = i.decorateImport(import_)

	if imp, ok := i.imports[import_]; ok {
		return imp
	}

	parts := strings.Split(import_, "/")

	alias := parts[len(parts)-1]
	if len(parts) >= 2 {
		part := parts[len(parts)-2]
		alias = part + "_" + alias
	}
	alias = regexNoAlphaNum.ReplaceAllString(alias, "_")
	alias = fmt.Sprintf("i%s_%s", strconv.FormatInt(i.counter, 16), alias)

	imp := Import{
		Path:  import_,
		Alias: alias,
	}
	i.imports[import_] = alias
	i.counter++
	i.importsSlice = append(i.importsSlice, imp)

	return alias
}

// Imports returns all imports in order of using them.
func (i *imports) Imports() []Import {
	return slices.Copy(i.importsSlice)
}

func (i *imports) decorateImport(imp string) string {
	for shortcut, path := range i.prefixes {
		if strings.Index(imp, shortcut) == 0 {
			return strings.Replace(imp, shortcut, path, 1)
		}
	}

	return imp
}
