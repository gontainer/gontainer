package syntax

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeImport(t *testing.T) {
	scenarios := []struct {
		input  string
		output string
	}{
		{
			input:  `"pkg"`,
			output: "pkg",
		},
		{
			input:  "config",
			output: "config",
		},
	}

	for i, s := range scenarios {
		t.Run(fmt.Sprintf("Scenario #%d", i), func(t *testing.T) {
			assert.Equal(t, s.output, SanitizeImport(s.input))
		})
	}
}

func TestCompileServiceValue(t *testing.T) {
	scenarios := []struct {
		input  string
		output string
	}{
		{
			input:  "Variable",
			output: "Variable",
		},
		{
			input:  "pkg.Variable",
			output: "i0_alias.Variable",
		},
		{
			input:  "&User{}",
			output: "&User{}",
		},
		{
			input:  `&"my/import/path".User{}`,
			output: "&i0_alias.User{}",
		},
		{
			input:  `".".GlobalConfig.DB`,
			output: "GlobalConfig.DB",
		},
	}

	a := mockAliases{alias: "i0_alias"}

	for i, s := range scenarios {
		t.Run(fmt.Sprintf("Scenario #%d", i), func(t *testing.T) {
			assert.Equal(t, s.output, CompileServiceValue(a, s.input))
		})
	}
}

type mockAliases struct {
	alias string
}

func (m mockAliases) Alias(string) string {
	return m.alias
}
