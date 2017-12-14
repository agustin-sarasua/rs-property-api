package app

import (
	"errors"
	"fmt"

	m "github.com/agustin-sarasua/rs-model"
)

func validateCondition(fn func(map[string]struct{}, string) bool, m map[string]struct{}, v string, msg string, errs []error) []error {
	if ok := fn(m, v); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func validateRangeCondition(fn func(int, int, int) bool, min int, max int, val int, msg string, errs []error) []error {
	if ok := fn(min, max, val); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func validateNonEmpty(fn func() bool, msg string, errs []error) []error {
	if ok := fn(); !ok {
		errs = append(errs, errors.New(msg))
	}
	return errs
}

func validateProperty(p *m.Property) []error {
	var errs []error
	isValueInMap := func(m map[string]struct{}, val string) bool {
		_, ok := m[val]
		return ok
	}
	isValidRange := func(min int, max int, val int) bool {
		return !(val < min || val > max)
	}

	errs = validateCondition(isValueInMap, m.PropertyTypes, p.Type, "Type is incorrect", errs)
	errs = validateCondition(isValueInMap, m.Orientation, p.Orientation, "Orientation is incorrect", errs)
	errs = validateRangeCondition(isValidRange, 0, 10, p.State, fmt.Sprintf("State should be between %v and %v", 0, 10), errs)
	errs = validateNonEmpty(func() bool { return p.Address != nil }, "Address can not be empty", errs)
	return errs
}
