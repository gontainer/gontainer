// Copyright (c) 2023 Bartłomiej Krukowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is furnished
// to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package input

import (
	"fmt"
)

const (
	ScopeShared Scope = iota + 1
	ScopeContextual
	ScopeNonShared
)

var (
	mapScopeString map[Scope]string
	mapStringScope map[string]Scope
)

func init() {
	mapScopeString = map[Scope]string{
		ScopeShared:     "shared",
		ScopeContextual: "contextual",
		ScopeNonShared:  "non_shared",
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
