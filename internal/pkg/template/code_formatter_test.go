package template_test

import (
	_ "embed"
	"testing"

	"github.com/gontainer/gontainer/internal/pkg/template"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed code_formatter_test.go
	goCode string
)

func TestCodeFormatter_Format(t *testing.T) {
	t.Parallel()
	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		f := template.NewCodeFormatter()
		_, err := f.Format("invalid go code")
		assert.ErrorContains(t, err, "CodeFormatter.Format: ")
	})

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		f := template.NewCodeFormatter()
		o, err := f.Format(goCode)
		assert.NoError(t, err)
		assert.NotEmpty(t, o)
	})
}
