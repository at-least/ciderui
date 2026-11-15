//go:build browser

package consolepopover

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	p      = "component-Popover"
	pop    = ".cui-console-popover"
	button = "[popovertarget]"
)

func page(html string) string { return `<div style="padding:24px 24px 240px">` + html + `</div>` }

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "popover",
			HTML:   page(h.Example("popover")),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: pop, Props: []string{"display"}, Golden: "closed until its button opens it"},
			},
		},
		{
			Name:   "popover open",
			HTML:   page(h.Example("popover")),
			Click:  []string{button},
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: pop, Props: []string{"background-color", "box-shadow", "border-top-left-radius", "min-width", "max-height"}, TB: []string{p + "-variant-variant-stroked-backgroundColor", p + "-variant-variant-stroked-boxShadow", p + "-variant-pageScale-desktop-borderRadius", p + "-minWidth", p + "-anatomy-Content-maxHeight"}},
				{Target: pop, Props: []string{"padding-top", "padding-left", "font-size", "line-height", "font-weight", "color", "max-width"}, Golden: "App Store Connect's analytics overview's info popover: 16px in, 13px on 18px, the label, 350px at its widest"},
				{Target: pop, Props: []string{"display", "flex-direction", "row-gap"}, Golden: "its parts stacked 8px apart: a control row's md button's (controlrow.json); the toolkit names no gap for a stack"},
				{Target: pop, Props: []string{"box-top", "box-left", "box-width"}, Golden: "open: 8px under its button, centered on it"},
				{Target: pop, Props: []string{"overflow-y"}, Golden: "scrolling past its tallest"},
			},
		},
		{
			Name:   "popover open low",
			HTML:   `<div style="padding:720px 24px 24px">` + h.Example("popover") + `</div>`,
			Click:  []string{button},
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: pop, Props: []string{"box-top"}, Golden: "above its button where there is no room below"},
			},
		},
		{
			Name:   "popover open phone",
			HTML:   page(h.Example("popover")),
			Click:  []string{button},
			Widths: []int{390},
			Checks: []h.Check{
				{Target: pop, Props: []string{"border-top-left-radius"}, TB: []string{p + "-variant-pageScale-mobile-borderRadius"}},
				{Target: pop, Props: []string{"box-left", "box-width"}, Golden: "on a phone, inside the window, 16px from its edges at the least"},
				{Target: pop, Props: []string{"max-width"}, Golden: "85% of the window at most"},
			},
		},
		{
			Name:         "popover open forced colors",
			HTML:         page(h.Example("popover")),
			Click:        []string{button},
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: pop, Props: []string{"border-top-width", "border-top-style", "border-top-color", "box-shadow"}, Golden: "a 1px CanvasText edge, as the sidebar's sheet: forced colors drop its shadow, and its ground is the page's"},
			},
		},
		{
			Name:   "popover dark",
			HTML:   `<div class="cui-dark">` + page(h.Example("popover")) + `</div>`,
			Click:  []string{button},
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: pop, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
