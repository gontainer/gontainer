package token

import (
	"fmt"

	"github.com/gontainer/gontainer-helpers/v2/exporter"
	"github.com/gontainer/gontainer/internal/pkg/consts"
	"github.com/gontainer/gontainer/internal/pkg/regex"
)

var (
	regexTokenRef = regex.MustCompileAz(regex.ParamName)
	regexSimpleFn = regex.MustCompileAz(`(?P<fn>` + regex.MetaFn + `)\((?P<params>.*)\)`)
)

// FactoryPercentMark handles %%.
type FactoryPercentMark struct{}

func (f FactoryPercentMark) Supports(expr string) bool {
	return expr == "%%"
}

func (f FactoryPercentMark) Create(string) (Token, error) {
	return Token{
		Kind:      KindString,
		Raw:       "%%",
		DependsOn: nil,
		Code:      fmt.Sprintf(consts.TplTokenProvider, `return "%", nil`),
	}, nil
}

// FactoryReference handles %my.param%.
type FactoryReference struct{}

func (FactoryReference) Supports(s string) bool {
	expr, ok := toExpr(s)

	return ok && regexTokenRef.MatchString(expr)
}

func (FactoryReference) Create(s string) (Token, error) {
	ref, _ := toExpr(s)

	return Token{
		Kind:      KindReference,
		Raw:       s,
		DependsOn: []string{ref},
		Code:      fmt.Sprintf(consts.TplTokenGetParam, ref),
	}, nil
}

type FactoryString struct{}

func (FactoryString) Supports(string) bool {
	return true
}

func (FactoryString) Create(expr string) (Token, error) {
	return Token{
		Kind:      KindString,
		Raw:       expr,
		DependsOn: nil,
		Code:      fmt.Sprintf(consts.TplTokenProvider, fmt.Sprintf("return %s, nil", exporter.MustExport(expr))),
	}, nil
}

// FactoryFunction handles %env(ENV_VAR)%.
type FactoryFunction struct {
	aliaser  aliaser
	fn       string
	goImport string
	goFn     string
}

func NewFactoryFunction(
	a aliaser,
	fn string,
	goImport string,
	goFn string,
) *FactoryFunction {
	return &FactoryFunction{
		aliaser:  a,
		fn:       fn,
		goImport: goImport,
		goFn:     goFn,
	}
}

func (f *FactoryFunction) Supports(expr string) bool {
	e, ok := toExpr(expr)
	if !ok {
		return false
	}

	ok, m := regex.Match(regexSimpleFn, e)
	return ok && m["fn"] == f.fn
}

func (f *FactoryFunction) Create(expr string) (Token, error) {
	e, _ := toExpr(expr)
	_, m := regex.Match(regexSimpleFn, e)
	goFn := f.goFn
	if f.goImport != "" {
		goFn = fmt.Sprintf("%s.%s", f.aliaser.Alias(f.goImport), goFn)
	}

	callFn := fmt.Sprintf(
		"%s.CallProvider(%s",
		f.aliaser.Alias(consts.GontainerHelperPath+"/caller"),
		goFn,
	)
	if m["params"] != "" {
		callFn += fmt.Sprintf(", []interface{}{%s}, true", m["params"])
	} else {
		callFn += ", nil, false"
	}
	callFn += ")"

	body := fmt.Sprintf(
		`r, err = %s; if err != nil { err = %s.Errorf("%%s: %%w", %s, err) }; return`,
		callFn,
		f.aliaser.Alias("fmt"),
		exporter.MustExport(fmt.Sprintf("cannot execute %s", expr)),
	)

	return Token{
		Kind: KindFunc,
		Raw:  expr,
		Code: fmt.Sprintf(consts.TplTokenProvider, body),
	}, nil
}

type FactoryUnexpectedFunction struct {
}

func (FactoryUnexpectedFunction) Supports(expr string) bool {
	e, ok := toExpr(expr)
	if !ok {
		return false
	}

	return regexSimpleFn.MatchString(e)
}

func (FactoryUnexpectedFunction) Create(expr string) (t Token, _ error) {
	e, _ := toExpr(expr)
	_, m := regex.Match(regexSimpleFn, e)

	return t, fmt.Errorf("unexpected function: %+q: %+q", m["fn"], expr)
}

type FactoryUnexpectedToken struct {
}

func (FactoryUnexpectedToken) Supports(expr string) bool {
	_, ok := toExpr(expr)
	return ok
}

func (f FactoryUnexpectedToken) Create(expr string) (t Token, _ error) {
	return t, fmt.Errorf("unexpected token: %+q", expr)
}
