//go:build browser

package consolestatus

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	status = ".cui-console-status"
	mark   = ".cui-console-status-mark"
)

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "statuses",
			HTML:   h.Example("status"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: mark, Props: []string{"height", "margin-right"}, ASC: h.ASCProbe("testflight", "status.icon")},
				{Target: status, Props: []string{"font-size", "line-height"}, ASC: h.ASCProbe("testflight", "status.word")},
				{Target: status, Props: []string{"display", "align-items"}, Golden: "the mark and the word on one line, the mark centered on it"},
				{Target: status, Props: []string{"color"}, Golden: "the word in the console's label: App Store Connect's older look wrote rgb(51, 51, 51)"},
				{Target: mark, Props: []string{"box-width"}, Golden: "as wide as App Store Connect's 12px (an asc check reads no width: its text was SF)"},
				{Target: mark, Props: []string{"background-color", "border-top-left-radius"}, Golden: "the mark in the console's warning label, the one color a status takes: it says look here (tokens.css, fitted to 4.5:1, decision 21), a circle (App Store Connect's glyph's shape is unread)"},
			},
		},
		{
			Name:   "status detail",
			HTML:   h.Example("detail"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-status-detail", Props: []string{"font-size", "line-height", "margin-top"}, ASC: h.ASCProbe("testflight", "status.line")},
				{Target: ".cui-console-status-detail", Props: []string{"color", "padding-left"}, Golden: "the secondary label, 18px in: under the word, past the mark and its 6px"},
			},
		},
		{
			Name:         "statuses forced colors",
			HTML:         h.Example("status"),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: mark, Props: []string{"background-color"}, Golden: "CanvasText: forced colors would drop the fill and the mark with it"},
			},
		},
		{
			Name:   "statuses dark",
			HTML:   `<div class="cui-dark">` + h.Example("status") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{{Target: mark, Props: []string{"background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"}},
		},
	})
}
