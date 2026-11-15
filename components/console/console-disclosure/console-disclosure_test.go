//go:build browser

package consoledisclosure

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	d      = ".cui-console-disclosure"
	closed = d + ":not([open])"
	open   = d + "[open]"
	// @invariant: a disclosure renders closed; a case reads an open one after a click opens the last
	openLast = d + ":last-child > summary"
)

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "disclosure",
			HTML:   h.Example("disclosure"),
			Click:  []string{openLast},
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: closed + " > summary", Props: []string{"list-style-type", "display", "column-gap", "color", "cursor"}, Golden: "no marker of the browser's; its chevron 4px before its label; the secondary label (the menu group's chevron's)"},
				{Target: closed + " > summary", Props: []string{"color"}, State: []string{"hover"}, Golden: "hovered, the label"},
				{Target: closed + " .cui-console-disclosure-chevron", Props: []string{"rotate", "box-width", "box-height", "box-left", "color"}, Golden: "closed, pointing inward at the label's leading edge (the HIG), 1em, Lucide's chevron-right and its secondary label as the menu group's"},
				{Target: closed + " .cui-console-disclosure-chevron", Props: []string{"color"}, State: []string{"hover"}, On: closed + " > summary", Golden: "the label hovered, the chevron still the secondary label, as the menu group's: its color its own, not the label's"},
				{Target: open + " .cui-console-disclosure-chevron", Props: []string{"rotate"}, Golden: "open, pointing down (the HIG): 90deg"},
				{Target: closed + " > summary", Props: []string{"font-size"}, Golden: "the text's size, inherited"},
				{Target: open + " .cui-console-disclosure-chevron", Props: h.Motion, Golden: "turned at once: the toolkit's controls draw no motion"},
			},
		},
		{
			Name:   "disclosure touch",
			HTML:   h.Example("disclosure"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: closed + " > summary", Props: []string{"box-height", "min-width"}, Golden: "on a touch screen the line itself 44px tall and 44px wide at least, the HIG's 44pt: a target drawn round it overlapped the next summary's and the opened content"},
			},
		},
		{
			Name:   "disclosure dark",
			HTML:   `<div class="cui-dark">` + h.Example("disclosure") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: closed + " > summary", Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
