//go:build browser

package consolesection

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const s = ".cui-console-section"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "section",
			HTML:   h.Example("section"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: s, Props: []string{"display", "row-gap"}, Golden: "its header, then each block, 20px apart: the room App Store Connect's section head leaves under it (head.json's section.head, margin-bottom 20px), one value for every block so no page writes its own"},
				{Target: s + " > p:last-child", Props: []string{"box-y"}, Golden: "the second block 20px under the first"},
			},
		},
	})
}
