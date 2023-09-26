package compiler

import (
	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

type aliasRegisterer interface {
	RegisterPrefixAlias(alias string, import_ string) error
}

type funcRegisterer interface {
	RegisterFunc(fnAlias string, goImport string, goFn string)
}

type paramResolver interface {
	ResolveParam(any) (resolver.ParamExpr, error)
}

type argResolver interface {
	ResolveArg(any) (resolver.ArgExpr, error)
}

type aliaser interface {
	// Alias returns an alias for given import, e.g. "github.com/spf13/viper" => "i0_viper".
	Alias(string) string
}
