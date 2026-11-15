//go:build browser

package consolelink

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	asc "github.com/at-least/ciderui/reference/asc/results"
)

const link = ".cui-console-link"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "link",
			HTML:   h.Example("link"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: link, Props: []string{"color", "text-decoration-line", "cursor"}, ASC: h.ASCProbe(asc.AnalyticsPage, "section.link")},
				{Target: link, Props: []string{"color", "text-decoration-line", "cursor"}, State: []string{"hover"}, ASC: h.ASCProbe(asc.HoverLink, "section.link")},
				{Target: link, Pseudo: "::after", Props: []string{"content"}, Golden: "with a mouse, no target drawn: none"},
			},
		},
		{
			Name:   "link touch",
			HTML:   h.Example("link"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: link, Pseudo: "::after", Props: []string{"top", "left", "transform"}, Golden: "the 44px target centered on the link, wider than it"},
				{Target: link, Pseudo: "::after", Props: []string{"width", "height"}, Golden: "a touch screen's target 44px square at least, the HIG's 44pt (crashcart's HIG pass: its See All and Open issue links, 13px tall, were the only controls without one)"},
			},
		},
		{
			Name:   "link dark",
			HTML:   `<div class="cui-dark">` + h.Example("link") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: link, Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
