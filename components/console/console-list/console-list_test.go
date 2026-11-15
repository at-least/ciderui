//go:build browser

package consolelist

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	d   = "component-DisplayList"
	l   = d + "-anatomy-Item-anatomy-Label-variant-pageScale"
	row = ".cui-console-list-row:nth-child(2)"
)

var typeProps = []string{"font-size", "font-weight", "line-height"}

func typeTokens(scale string) []string {
	var names []string
	for _, p := range []string{"fontSize", "fontWeight", "lineHeight"} {
		names = append(names, l+"-"+scale+"-textStyle-"+p)
	}
	return names
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "list",
			HTML:   h.Example("list"),
			Widths: []int{1200, 640},
			Checks: []h.Check{
				{Target: ".cui-console-list-label", Props: []string{"display"}, Golden: "its own line, over what follows it"},
				{Target: row, Props: []string{"padding-left", "padding-right"}, TB: []string{d + "-groupedGray-contentPaddingInlineDesktop", d + "-groupedGray-contentPaddingInlineDesktop"}},
				{Target: row, Props: []string{"min-height", "padding-top", "padding-bottom"}, TB: []string{d + "-variant-pageScale-desktop-variant-sidebar-anatomy-Item-minHeight", d + "-variant-pageScale-desktop-variant-sidebar-anatomy-Item-paddingBlock", d + "-variant-pageScale-desktop-variant-sidebar-anatomy-Item-paddingBlock"}},
				{Target: row + " .cui-console-list-label", Props: typeProps, TB: typeTokens("desktop")},
				{Target: row + " .cui-console-list-label", Props: []string{"color"}, TB: []string{d + "-anatomy-Item-variant-colorRole-neutral-effect-default-anatomy-Label-color"}},
				{Target: row + " .cui-console-list-detail", Props: []string{"color", "font-size", "line-height"}, Golden: "the secondary label at 0.76: the toolkit's 0.60 is 3.4:1 on white (tokens.css, decision 21); the body's 13px"},
				{Target: row, Pseudo: "::before", Props: []string{"left", "height", "background-color"}, Golden: "a 0.5px hairline over each row after the first, from 16px in, the secondary border (the table row's)"},
				{Target: ".cui-console-list-row:first-child", Pseudo: "::before", Props: []string{"content"}, Golden: "no hairline over the first row"},
				{Target: ".cui-console-list", Props: []string{"background-color", "border-top-left-radius"}, Golden: "the group the console card's gray and corners"},
				{Target: row, Props: []string{"background-color"}, State: []string{"hover"}, Golden: "a row is no link and does not fill"},
			},
		},
		{
			Name:   "list phone",
			HTML:   h.Example("list"),
			Widths: []int{639, 390},
			Checks: []h.Check{
				{Target: row + " .cui-console-list-label", Props: typeProps, TB: typeTokens("mobile")},
				{Target: row, Props: []string{"padding-left"}, TB: []string{d + "-groupedGray-contentPaddingInlineMobile"}},
			},
		},
		{
			Name:         "list forced colors",
			HTML:         h.Example("list"),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: row, Pseudo: "::before", Props: []string{"background-color", "height"}, Golden: "the hairline between rows CanvasText: forced colors drew its fill in the ground's color, at its alpha"},
			},
		},
		{
			Name:   "list content",
			HTML:   h.Example("content"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-list-row", Props: []string{"padding-left", "padding-top", "min-height", "background-color"}, Golden: "a row holding the page's own content after its label: a row's 16px in, 7px above, 44px at least, no fill of its own"},
			},
		},
		{
			Name:   "list dark",
			HTML:   `<div class="cui-dark">` + h.Example("list") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-list", Props: []string{"background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: row + " .cui-console-list-label", Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: row, Pseudo: "::before", Props: []string{"background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
