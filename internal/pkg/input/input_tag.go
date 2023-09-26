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
