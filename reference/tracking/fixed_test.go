package main

import (
	"math"
	"testing"
)

// TestToFixed holds toFixed to what Node's Number.prototype.toFixed gave for each
// (node -e, 2026-11-15): a half rounds away from zero, where strconv's 'f' rounds
// 0.03125 to 0.0312.
func TestToFixed(t *testing.T) {
	for _, c := range []struct {
		x      float64
		digits int
		want   string
	}{
		{0.03125, 4, "0.0313"}, {-0.03125, 4, "-0.0313"}, {1.00005, 4, "1.0001"}, {0.00005, 4, "0.0001"},
		{-0.00001, 4, "-0.0000"}, {2.5, 0, "3"}, {0.5, 0, "1"}, {math.Copysign(0, -1), 4, "0.0000"},
		{0.0651234, 4, "0.0651"}, {123.456789, 5, "123.45679"}, {1e-7, 4, "0.0000"}, {math.NaN(), 4, "NaN"},
	} {
		if got := toFixed(c.x, c.digits); got != c.want {
			t.Errorf("toFixed(%v, %d) = %s, want %s", c.x, c.digits, got, c.want)
		}
	}
}
