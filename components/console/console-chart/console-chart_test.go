//go:build browser

package consolechart

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	asc "github.com/at-least/ciderui/reference/asc/results"
)

const c = ".cui-console-chart"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "chart",
			HTML:   h.Example("chart"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: c + "-grid", Props: []string{"stroke-width"}, ASC: h.ASCProbe(asc.Analytics, "chart.grid")},
				{Target: c + "-scale", Props: []string{"font-size", "font-weight"}, ASC: h.ASCProbe(asc.Analytics, "chart.y")},
				{Target: c + "-axis", Props: []string{"font-size", "font-weight"}, ASC: h.ASCProbe(asc.Analytics, "chart.x")},
				{Target: c + "-grid", Props: []string{"stroke"}, Golden: "the tertiary hairline, where App Store Connect's is rgb(214, 214, 214)"},
				{Target: c + "-grid line:nth-child(2)", Props: []string{"box-y"}, Golden: "the kit's gridlines where its scale labels them, the plot's top and middle: the middle one under the half scale"},
				{Target: c + "-scale", Props: []string{"color", "background-color", "right", "top"}, Golden: "the fitted secondary label (App Store Connect's rgb(136, 136, 136) is 3.5:1 on white, decision 21) on the ground, at the plot's trailing top (the HIG keeps the leading edge clean)"},
				{Target: c + `-scale[data-at="half"]`, Props: []string{"box-top"}, Golden: "the half scale on the middle gridline"},
				{Target: c + "-axis", Props: []string{"color", "margin-top", "justify-content"}, Golden: "the first and last labels at the plot's ends, 6px under it, the fitted secondary label"},
				{Target: c + "-mark", Props: []string{"stroke", "stroke-width"}, Golden: "a 1px line of the ground round each mark: stacked marks stand apart (the HIG's Color)"},
				{Target: c + "-mark", Props: []string{"fill"}, Golden: "the class owns its color: the toolkit's theme blue (tokens.css, decision 18), never a fill of the page's — a consumer once filled its bars with the site set's blue"},
				{Target: c + "-svg", Props: []string{"box-height", "display"}, Golden: "its height the page's, the svg's 140"},
				{Target: c + "-svg", Props: []string{"box-width"}, Golden: "as wide as its plot"},
				{Target: c + "-svg", Props: []string{"position"}, Golden: "positioned, after the grid: its marks paint over the gridlines"},
				{Target: c + "-grid", Props: []string{"overflow-y"}, Golden: "the top gridline drawn whole: at y 0, half its stroke is above the plot"},
			},
		},
		{
			Name:   "chart hover",
			HTML:   h.Example("chart"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: c + "-bar:last-of-type " + c + "-mark", Props: []string{"opacity"}, State: []string{"hover"}, On: c + "-bar:last-of-type", Golden: "the only cue naming the bar under the pointer, at once: no date pill to read (ASC's chart hover), no motion (the HIG's feedback)"},
			},
		},
		{
			Name:   "chart in a card",
			HTML:   `<div class="cui-console-page"><div class="cui-console-card">` + h.Example("chart") + `</div></div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: c + "-scale", Props: []string{"border-top-left-radius"}, Golden: "the scale's rounded corners"},
				{Target: c + "-svg", Props: []string{"overflow-y"}, Golden: "a mark on the edge of the plot is not clipped"},
				{Target: c + "-mark", Props: []string{"stroke"}, Golden: "in a card, the line round a mark is the card's ground"},
				{Target: c + "-scale", Props: []string{"background-color"}, Golden: "in a card, the scale on the card's ground"},
			},
		},
		{
			Name:   "chart empty",
			HTML:   h.Example("empty"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: c + "-empty", Props: []string{"font-size", "color", "box-top", "box-height", "align-items", "justify-content"}, Golden: "centered on the plot, 13px, the secondary label"},
			},
		},
		{
			Name:   "chart dark",
			HTML:   `<div class="cui-dark">` + h.Example("chart") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: c + "-grid", Props: []string{"stroke"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: c + "-scale", Props: []string{"color", "background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: c + "-mark", Props: []string{"fill"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
