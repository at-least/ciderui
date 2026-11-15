// Package reference holds what was measured on Apple's pages (reference/apple.com/
// results/, reference/<signed-in site>/results/, reference/tracking/table/), the
// values the tests hold ciderui to.
//
// @invariant: A probe, width or property the record lacks is an error: a check never
// compares with a value nobody measured.
package reference

import (
	"strconv"
)

// Widths are the windows a case is read at unless it says otherwise.
//
// @trap: not apple.com's fourth, 1600: no part's rule starts above 1200, so only what
// the window sizes, up to its max-width, reads otherwise there, in the cases that say
// harness.With1600
var Widths = []int{1200, 900, 390}

func px(v float64) string { return strconv.FormatFloat(v, 'f', -1, 64) + "px" }
