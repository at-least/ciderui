//go:build browser

package consolemessage

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const m = ".cui-console-message"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "message",
			HTML:   h.Example("message"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: m + ":not([data-state])", Props: []string{"color", "font-size", "line-height"}, Golden: "what the request did, in the text's color at the body's size"},
				{Target: m + `[data-state="error"]`, Props: []string{"color"}, Golden: "a request that failed in the error color, its words saying what to do next"},
				{Target: m + ":empty", Props: []string{"display", "position"}, Golden: "empty, it takes no room (no margin or gap of the page's is left behind) and stays in the accessibility tree, so the line its script fills is announced"},
			},
		},
	})
}
