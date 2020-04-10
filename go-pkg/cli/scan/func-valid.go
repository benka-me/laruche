package scan

import (
	"errors"
	"regexp"
)

type StringValid func(s string) error

func IsAlphaNum(s string) error {
	var rgx = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if rgx.MatchString(s) {
		return nil
	} else {
		return errors.New("alphanumeric and dash only")
	}
}

func IsAlphaDash(s string) error {
	var rgx = regexp.MustCompile(`^[a-zA-Z-]+$`)
	if rgx.MatchString(s) {
		return nil
	} else {
		return errors.New("alphanumeric and dash only")
	}
}

func IsAlphanumDash(s string) error {
	var rgx = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	if rgx.MatchString(s) {
		return nil
	} else {
		return errors.New("alphanumeric and dash only")
	}
}

func PortValidator(Public bool) StringValid {
	var portValidator StringValid

	if Public {
		portValidator = func(s string) error {
			// TODO: verify availability of public port
			return nil
		}
	} else {
		portValidator = func(s string) error {
			// TODO: verify availability of port from author existing cells...
			return nil
		}
	}
	return portValidator
}
