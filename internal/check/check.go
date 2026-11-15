// Package check is how a test compares a value it read with the one it wants.
package check

import (
	"fmt"
	"os"
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
	dump(ok, failure)
	if !ok {
		t.Error(failure)
	}
	return ok
}

// @invariant: CIDERUI_CHECKS=<file> appends every check, passed or not, to the file,
// one line each: the Node harness wrote the same lines, so the two runs' checks diff
// line for line while the tests move to Go
func dump(ok bool, failure string) {
	path := os.Getenv("CIDERUI_CHECKS")
	if path == "" {
		return
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := fmt.Fprintf(f, "%v\t%s\n", ok, failure); err != nil {
		panic(err)
	}
}

// Number is v as JavaScript writes it in a template, for the values a page reports:
// the shortest form that reads back as v.
func Number(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) }
