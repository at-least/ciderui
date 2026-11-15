//go:build browser

package consoleskip

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const skip = ".cui-console-skip"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "skip link",
			HTML:   h.Example("skip"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: skip, Props: []string{"box-width", "box-height", "clip-path"}, Golden: "unseen at rest, a 1px box clipped to nothing: still a tab stop and in the accessibility tree"},
				{Target: skip, Props: []string{"box-top", "box-left"}, State: []string{"focus"}, TB: []string{"space-pageMarginSm", "space-pageMarginSm"}},
				{Target: skip, Props: []string{"box-height", "clip-path", "position", "z-index"}, State: []string{"focus"}, Golden: "focused, the filled button whole, fixed over the page and the sidebar's sheet"},
				{Target: skip, Props: []string{"background-color", "color"}, State: []string{"focus"}, Golden: "the filled button's brand fill and white label"},
			},
		},
	})
}
