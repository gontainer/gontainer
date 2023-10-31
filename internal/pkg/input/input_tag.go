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
	"errors"
	"fmt"
)

func (t *Tag) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var z interface{}
	if err := unmarshal(&z); err != nil {
		return err
	}

	// "tag"
	if s, ok := z.(string); ok {
		t.Name = s
		t.Priority = 0
		return nil
	}

	// {"name": "tag", "priority": 5}
	if m, ok := z.(map[string]interface{}); ok {
		if n, ok := m["name"]; ok {
			name, ok := n.(string)
			if !ok {
				return errors.New("name must be an instance of string")
			}
			t.Name = name
		} else {
			return errors.New("missing tag name")
		}

		if p, ok := m["priority"]; ok {
			priority, ok := p.(int)
			if !ok {
				return errors.New("priority must be an int")
			}
			t.Priority = priority
		} else {
			t.Priority = 0
		}

		return nil
	}

	return fmt.Errorf("unexpected type `%T`", z)
}
