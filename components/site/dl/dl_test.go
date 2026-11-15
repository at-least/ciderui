//go:build browser

package dl

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const d = ".cui-dl"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "dl",
			HTML:   h.Example("dl"),
			Widths: []int{1200, 735},
			Checks: []h.Check{
				{Target: d, Props: []string{"display", "grid-template-columns", "row-gap", "column-gap"}, Golden: "the pairs in two columns from 735px, 12px down and 24px across: totality's order details (pages.go's dlRow)"},
				{Target: d + " > div:first-child > dt", Props: []string{"color", "font-size", "line-height"}, Golden: "the term in the secondary text over its value, the body's reduced size"},
				{Target: d + " > div:first-child > dd", Props: []string{"color", "margin-left"}, Golden: "the value in the text's color under its term, no indent"},
			},
		},
		{
			Name:   "dl phone",
			HTML:   h.Example("dl"),
			Widths: []int{734, 390},
			Checks: []h.Check{
				{Target: d, Props: []string{"grid-template-columns"}, Golden: "one column below 735px"},
			},
		},
		{
			Name:   "dl zh",
			HTML:   strings.Replace(h.Example("dl"), "<dl class", `<dl lang="zh-Hant" class`, 1),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: d + " > div:first-child > dt", Props: []string{"line-height"}, Golden: "Traditional Chinese a pixel taller, as cui-body-reduced is"},
			},
		},
	})
}
