//go:build browser

package consolecard

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const card = ".cui-console-card"

var typeProps = []string{"font-size", "line-height", "font-weight", "letter-spacing"}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "card",
			HTML:   h.Example("list"),
			Widths: []int{1600, 1200},
			Checks: []h.Check{
				{Target: card, Props: []string{"padding-top", "border-top-left-radius", "background-color", "overflow", "box-shadow", "border-top-width"}, ASC: h.ASCProbe("analytics-page", "card")},
				{Target: ".cui-console-card-title", Props: h.Props(typeProps, []string{"color"}), ASC: h.ASCProbe("analytics", "card.title")},
			},
		},
		{
			Name:   "card list",
			HTML:   h.Example("list"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-card-header", Props: []string{"display", "justify-content", "align-items"}, Golden: "a header: its title at its start, its link at its end, on the baseline"},
			},
		},
		{
			Name:   "card dark",
			HTML:   `<div class="cui-dark">` + h.Example("list") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: card, Props: []string{"background-color"}, Golden: "the dark side derived: the toolkit's gray coolGray-50's dark step (tokens.css, decision 18)"},
			},
		},
	})
}
