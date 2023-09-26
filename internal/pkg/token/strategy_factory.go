package token

import (
	"fmt"
)

type tokenFactoryStrategy interface {
	Supports(string) bool
	Create(string) (Token, error)
}

type StrategyFactory struct {
	strategies []tokenFactoryStrategy
}

func NewStrategyFactory(strategies ...tokenFactoryStrategy) *StrategyFactory {
	return &StrategyFactory{strategies: strategies}
}

func (f *StrategyFactory) Prepend(s tokenFactoryStrategy) {
	f.strategies = append([]tokenFactoryStrategy{s}, f.strategies...)
}

func (f *StrategyFactory) Create(i string) (t Token, _ error) {
	for _, s := range f.strategies {
		if s.Supports(i) {
			return s.Create(i)
		}
	}
	return t, fmt.Errorf("not supported token: %s", i)
}

type FuncRegisterer struct {
	prepender interface{ Prepend(s tokenFactoryStrategy) }
	aliaser   aliaser
}

func NewFuncRegisterer(
	p interface{ Prepend(s tokenFactoryStrategy) },
	a aliaser,
) *FuncRegisterer {
	return &FuncRegisterer{
		prepender: p,
		aliaser:   a,
	}
}

func (f *FuncRegisterer) RegisterFunc(fnAlias string, goImport string, goFn string) {
	f.prepender.Prepend(NewFactoryFunction(f.aliaser, fnAlias, goImport, goFn))
}
