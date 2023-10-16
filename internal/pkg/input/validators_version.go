package input

import (
	"errors"
	"fmt"

	"github.com/gontainer/gontainer-helpers/grouperror"
	"golang.org/x/mod/semver"
)

// DefaultVersionValidators returns validators for Input.Version.
func DefaultVersionValidators(version string) []ValidateFn {
	return []ValidateFn{
		NewVersionValidator(version).ValidateVersion,
	}
}

type VersionValidator struct {
	version string
	valid   bool
}

// NewVersionValidator expects a proper version without the leading "v".
func NewVersionValidator(version string) *VersionValidator {
	ver := version
	val := false

	if semver.IsValid("v" + version) {
		ver = "v" + version
		val = true
	}

	return &VersionValidator{
		version: ver,
		valid:   val,
	}
}

func (v *VersionValidator) ValidateVersion(i Input) (err error) {
	if i.Version == nil || !v.valid {
		return nil
	}

	defer func() {
		err = grouperror.Prefix(fmt.Sprintf("version: current: %s, given: %s: ", v.version, *i.Version), err)
	}()

	curr := semver.MajorMinor(v.version) + ".0"
	given := semver.MajorMinor(string(*i.Version)) + ".0"

	if semver.Major(v.version) == "v0" {
		if curr != given {
			return errors.New("possibly incompatible versions")
		}
		return
	}

	if semver.Major(v.version) != semver.Major(string(*i.Version)) {
		return errors.New("incompatible versions")
	}

	if semver.Compare(curr, given) < 0 {
		return errors.New("update Gontainer to use all new features")
	}

	return
}
