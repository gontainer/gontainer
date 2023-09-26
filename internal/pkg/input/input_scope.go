package input

import (
	"fmt"
)

const (
	ScopeShared Scope = iota + 1
	ScopeContextual
	ScopePrivate
)

var (
	mapScopeString map[Scope]string
	mapStringScope map[string]Scope
)

func init() {
	mapScopeString = map[Scope]string{
		ScopeShared:     "shared",
		ScopeContextual: "contextual",
		ScopePrivate:    "private",
	}
	mapStringScope = make(map[string]Scope)
	for sc, str := range mapScopeString {
		mapStringScope[str] = sc
	}
}

func (s Scope) String() string {
	if v, ok := mapScopeString[s]; ok {
		return v
	}

	return fmt.Sprintf("invalid (%d)", s)
}

func (s *Scope) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var z string
	if err := unmarshal(&z); err != nil {
		return err
	}

	if val, ok := mapStringScope[z]; ok {
		*s = val
		return nil
	}

	return fmt.Errorf("invalid value for %T: %+q", *s, z)
}
