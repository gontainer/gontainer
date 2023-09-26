package input

import (
	"fmt"
)

func (c *Call) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var z []interface{}
	if err := unmarshal(&z); err != nil {
		return err
	}

	if len(z) == 0 || len(z) > 3 {
		return fmt.Errorf("the object Call must contain 1 - 3 values, %d given", len(z))
	}

	if s, ok := z[0].(string); !ok {
		return fmt.Errorf("first element of the object Call must be a string, `%T` given", z[0])
	} else {
		c.Method = s
	}

	if len(z) >= 2 {
		if args, ok := z[1].([]any); !ok {
			return fmt.Errorf("second element of the object Call must be an array, `%T` given", z[1])
		} else {
			c.Args = args
		}
	}

	if len(z) >= 3 {
		if i, ok := z[2].(bool); !ok {
			return fmt.Errorf("third element of the object Call must be a bool, `%T` given", z[2])
		} else {
			c.Immutable = i
		}
	}

	return nil
}
