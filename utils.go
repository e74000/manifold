package manifold

import (
	"cmp"
	"fmt"
)

// mustOk returns a function that enforces a successful operation (`ok` is true).
// If `ok` is false, it panics with the given `okContext`.
func mustOk[T any](okContext string) func(a T, ok bool) T {
	return func(a T, ok bool) T {
		if !ok {
			panic(fmt.Errorf("mustOk: %s", okContext))
		}

		return a
	}
}

// checkOneOf checks if the value is in the allowed set of values.
func checkOneOf[T comparable](value T, allowed ...T) error {
	for _, a := range allowed {
		if value == a {
			return nil
		}
	}
	return fmt.Errorf("invalid value: %v, allowed values are: %v", value, allowed)
}

// checkInRange checks if a value is within the allowed range [min, max] and returns a custom error if not.
func checkInRange[T cmp.Ordered](value, min, max T) error {
	if value < min || value > max {
		return fmt.Errorf("invalid value: %v, must be within range [%v, %v]", value, min, max)
	}
	return nil
}
