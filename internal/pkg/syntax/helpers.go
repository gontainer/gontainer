package syntax

import (
	"strings"

	"github.com/gontainer/gontainer/internal/pkg/regex"
)

var (
	regexServiceValue = regex.MustCompileAz(regex.ServiceValue)
)

type aliaser interface {
	Alias(string) string
}

// SanitizeImport removes surrounding quotation marks.
// If `result == "."`, it returns an empty string.
func SanitizeImport(i string) string {
	r := strings.Trim(i, `"`)
	if r == "." {
		r = ""
	}
	return r
}

// CompileServiceValue expects correct expr, validation must be done earlier.
func CompileServiceValue(a aliaser, expr string) string {
	_, m := regex.Match(regexServiceValue, expr)

	if m["v1"] != "" {
		parts := make([]string, 0)
		import_ := SanitizeImport(m["import"])
		if import_ != "" {
			parts = append(parts, a.Alias(import_))
		}
		return m["ptr"] + strings.Join(append(parts, m["value"]), ".")
	}

	parts := make([]string, 0)
	import_ := SanitizeImport(m["import2"])
	if import_ != "" {
		parts = append(parts, a.Alias(import_))
	}
	return m["ptr2"] + strings.Join(append(parts, m["struct2"]), ".") + "{}"
}
