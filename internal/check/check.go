// Package check is how a test compares a value it read with the one it wants.
package check

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

var length = regexp.MustCompile(`^(-?[\d.]+)px$`)

// Tolerance is how far apart two lengths Equal holds the same may be.
const Tolerance = 1.0 / 32

// Same is whether got is want, or both are lengths within tolerance px of each other.
func Same(got, want string, tolerance float64) bool {
	if got == want {
		return true
	}
	a, b := length.FindStringSubmatch(got), length.FindStringSubmatch(want)
	if a == nil || b == nil {
		return false
	}
	x, errX := strconv.ParseFloat(a[1], 64)
	y, errY := strconv.ParseFloat(b[1], 64)
	if errX != nil || errY != nil {
		return false
	}
	d := x - y
	return d <= tolerance && d >= -tolerance
}

// Equal fails t where got is not want (Same within Tolerance), naming the check.
func Equal(t testing.TB, name, got, want string) bool {
	t.Helper()
	return Within(t, name, got, want, Tolerance)
}

// Within is Equal with lengths held within tolerance px.
func Within(t testing.TB, name, got, want string, tolerance float64) bool {
	t.Helper()
	return Hold(t, Same(got, want, tolerance), fmt.Sprintf("%s: %s (want %s)", name, got, want))
}

// Hold fails t with failure where ok is false.
func Hold(t testing.TB, ok bool, failure string) bool {
	t.Helper()
	if !ok {
		t.Error(failure)
	}
	return ok
}

// Number is v as JavaScript writes it in a template, for the values a page reports:
// the shortest form that reads back as v.
func Number(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }
