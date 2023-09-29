package template

import (
	"github.com/gontainer/gontainer/internal/pkg/output"
)

// data is used for syntax autocomplete in *.go.tpl files.
type data struct {
	ImportCollection importProvider
	Output           output.Output
	BuildInfo        string
	Stub             bool
}
