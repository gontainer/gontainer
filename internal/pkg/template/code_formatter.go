package template

import (
	"go/format"
	"regexp"

	"github.com/gontainer/gontainer-helpers/v2/grouperror"
	"golang.org/x/tools/imports"
)

func init() {
	imports.LocalPrefix = "github.com/gontainer/gontainer-helpers/v2"
}

type CodeFormatter struct {
}

func NewCodeFormatter() *CodeFormatter {
	return &CodeFormatter{}
}

var (
	reEmptyNewLines = regexp.MustCompile(`\n\n+\t`)
)

func (CodeFormatter) Format(c string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = grouperror.Prefix("CodeFormatter.Format: ", err)
		}
	}()

	var r []byte
	r, err = format.Source([]byte(c))
	if err != nil {
		return "", err
	}
	r = reEmptyNewLines.ReplaceAll(r, []byte("\n\t"))

	// remove unused imports
	// required for generating stubs
	r, err = imports.Process("", r, nil)

	return string(r), err
}
