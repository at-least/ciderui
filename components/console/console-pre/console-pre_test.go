//go:build browser

package consolepre

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
)

const l = ".cui-console-pre"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "pre",
			HTML:   h.Example("pre"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: l, Props: []string{"border-top-left-radius"}, Apple: h.AppleProbe(apple.Docc, "listing")},
				{Target: l, Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left"}, Apple: h.AppleProbe(apple.Docc, "listing.pre")},
				{Target: l, Props: []string{"white-space"}, Apple: h.AppleProbe(apple.Docc, "listing.pre")},
				{Target: l, Props: []string{"font-size", "line-height", "font-weight"}, Golden: "the console's 13px mono body (monoBodyRegular), its line in DocC's proportion, 25px on 15px"},
				{Target: l, Props: []string{"font-family"}, Golden: "the console's mono face (tokens.css's code stack)"},
				{Target: l, Props: []string{"background-color", "color", "overflow-x"}, Golden: "the console card's gray, the nearest the contract holds to DocC's rgb(245, 245, 247); the label; a long line scrolls sideways inside it, as DocC's scroller does"},
			},
		},
		{
			Name:   "pre nested",
			HTML:   h.Example("nested"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: l, Props: []string{"background-color"}, Golden: "in a console card, whose gray is its own, the page's white"},
			},
		},
		{
			Name:   "pre dark",
			HTML:   `<div class="cui-dark">` + h.Example("pre") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: l, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
