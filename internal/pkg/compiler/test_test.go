package compiler_test

import (
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/resolver"
)

var (
	nilAliasRegisterer = aliasRegistererFunc(func(string, string) error {
		return nil
	})
	nilFuncRegisterer = funcRegistererFunc(func(string, string, string) {
	})
	simpleAliaserFunc = aliaserFunc(func(s string) string {
		return "alias_" + s
	})

	simpleArg = resolver.ArgExpr{
		Code:              "my-arg",
		Raw:               "my-raw-arg",
		DependsOnParams:   nil,
		DependsOnServices: nil,
		DependsOnTags:     nil,
	}
	simpleArgResolverFunc = argResolverFunc(func(a any) (resolver.ArgExpr, error) {
		return simpleArg, nil
	})
)

type aliasRegistererFunc func(alias string, import_ string) error

func (a aliasRegistererFunc) RegisterPrefixAlias(alias string, import_ string) error {
	return a(alias, import_)
}

type aliaserFunc func(s string) string

func (a aliaserFunc) Alias(s string) string {
	return a(s)
}

type funcRegistererFunc func(fnAlias string, goImport string, goFn string)

func (f funcRegistererFunc) RegisterFunc(fnAlias string, goImport string, goFn string) {
	f(fnAlias, goImport, goFn)
}

type paramResolveFunc func(any) (resolver.ParamExpr, error)

func (p paramResolveFunc) ResolveParam(a any) (resolver.ParamExpr, error) {
	return p(a)
}

type argResolverFunc func(any) (resolver.ArgExpr, error)

func (a argResolverFunc) ResolveArg(i any) (resolver.ArgExpr, error) {
	return a(i)
}

type inputValidatorFunc func(input input.Input) error

func (fn inputValidatorFunc) Validate(i input.Input) error {
	return fn(i)
}
