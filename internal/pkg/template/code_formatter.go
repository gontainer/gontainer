package template

import (
	"go/format"
	"regexp"

	"github.com/gontainer/gontainer-helpers/errors"
)

type CodeFormatter struct {
}

func NewCodeFormatter() *CodeFormatter {
	return &CodeFormatter{}
}

var (
	reEmptyNewLines = regexp.MustCompile(`\n\n+\t`)
)

func (CodeFormatter) Format(c string) (string, error) {
	var r []byte
	r, err := format.Source([]byte(c))
	err = errors.PrefixedGroup("CodeFormatter.Format: ", err)
	r = reEmptyNewLines.ReplaceAll(r, []byte("\n\t"))
	return string(r), err
}
