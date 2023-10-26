package input

import (
	"errors"

	"golang.org/x/mod/semver"
)

func (v *Version) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val any
	if err := unmarshal(&val); err != nil {
		return err
	}

	vs, ok := val.(string)
	if !ok {
		return errors.New("it must be a string")
	}

	if !semver.IsValid("v" + vs) {
		return errors.New(`version must follow the semver scheme, and it must not be Prefix by "v", see https://semver.org/`)
	}

	*v = Version(vs)
	return nil
}
