//go:build browser

package consolecanvas

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	asc "github.com/at-least/ciderui/reference/asc/results"
)

const (
	bar     = ".cui-console-canvas-bar"
	back    = ".cui-console-link-back"
	title   = bar + " > h1"
	actions = ".cui-console-canvas-actions"
)

func TestCases(t *testing.T) {
	scrolled := strings.Replace(h.Example("bar"), "</main>", `<div style="height: 3000px"></div></main><script>scrollTo(0, 1000)</script>`, 1)
	long := strings.Replace(h.Example("bar"), "Episode 1", "Episode 1, the night the lighthouse keeper's daughter first saw the ship", 1)
	h.Run(t, []h.Case{
		{
			Name:   "bar",
			HTML:   h.Example("bar"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: bar, Props: []string{"height", "padding-left", "padding-right", "border-bottom-width", "box-shadow", "backdrop-filter"}, ASC: h.ASCProbe(asc.Apps, "nav")},
				{Target: title, Props: []string{"font-size", "line-height", "font-weight"}, ASC: h.ASCProbe(asc.Apps, "nav.current")},
				{Target: bar, Props: []string{"background-color"}, Golden: "the console card's gray: App Store Connect's rgb(241, 241, 241) is no contract color, the card's 96.8% the nearest"},
				{Target: title, Props: []string{"box-left", "box-width", "text-align"}, Golden: "the title centered on the bar, whatever the way back and the actions beside it are"},
				{Target: back, Props: []string{"box-left"}, Golden: "the way back at the bar's start, 40px in"},
				{Target: actions, Props: []string{"box-left", "box-width", "column-gap"}, Golden: "the actions at the bar's end, 40px in, 8px apart"},
			},
		},
		{
			Name:   "bar narrow",
			HTML:   h.Example("bar"),
			Widths: []int{900, 390},
			Checks: []h.Check{
				{Target: bar, Props: []string{"padding-left", "padding-right"}, Golden: "below 1024px the console column's room, 24px from 640px, 16px below (App Store Connect was read at 1920px alone)"},
			},
		},
		{
			Name:   "bar phone",
			HTML:   h.Example("bar"),
			Widths: []int{390},
			Checks: []h.Check{
				{Target: title, Props: []string{"box-left", "box-top", "text-align"}, Golden: "below 640px iOS's large-title bar: the title on its own row under the way back and the actions, at the start, 16px in"},
				{Target: back, Props: []string{"box-left", "box-top"}, Golden: "the way back on the first row, 16px in"},
				{Target: actions, Props: []string{"box-left", "box-width", "box-top"}, Golden: "the actions on the first row, at its end, 16px in, the switch's words on one line"},
				{Target: bar, Props: []string{"padding-top", "padding-bottom", "row-gap"}, Golden: "two rows 4px apart, 8px in from the bar's edges"},
			},
		},
		{
			Name:   "bar long title",
			HTML:   long,
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: title, Props: []string{"white-space", "text-overflow", "box-height"}, Golden: "a long title wraps, never an ellipsis (decision 31), the bar growing past its 50px"},
			},
		},
		{
			Name:   "bar scrolled",
			HTML:   scrolled,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: bar, Props: []string{"position", "box-top", "z-index"}, Golden: "kept at the top as the page scrolls, as muse's was, at App Store Connect's z-index for its bar (apps:nav, 2000), where App Store Connect's own scrolls away"},
			},
		},
		{
			Name:   "bar dark",
			HTML:   `<div class="cui-dark">` + h.Example("bar") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: bar, Props: []string{"background-color", "color"}, Golden: "the dark card's gray and its label"},
			},
		},
	})
}
