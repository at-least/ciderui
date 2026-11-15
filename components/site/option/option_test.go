//go:build browser

package option

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
)

const (
	input = ".cui-option-input"
	label = ".cui-option-label"
)

var edge = []string{"border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-color", "border-top-left-radius"}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name: "zh-hant",
			HTML: `<div lang="zh-Hant">` + h.Example("option") + `</div>`,
			Checks: []h.Check{
				{Target: ".cui-option-title", Props: []string{"line-height", "letter-spacing"}, Apple: h.AppleProbe(apple.BuyTw, "option.title")},
			},
		},
		{
			Name:   "option",
			Widths: h.With1600,
			HTML:   h.Example("option"),
			Checks: []h.Check{
				{Target: label, Props: h.Props(h.Size, []string{"color"}, h.Padding, edge, []string{"background-color", "min-height", "transition-property", "transition-duration"}), Apple: h.AppleProbe(apple.Buy, "option")},
				{Target: label, State: []string{"hover"}, Props: []string{"background-color", "border-top-color", "color"}, Apple: h.AppleProbe(apple.Buy, "option", "hover")},
				{Target: label, On: input, State: []string{"focus-visible"}, Props: []string{"outline-style", "outline-width", "outline-color", "outline-offset"}, Apple: h.AppleProbe(apple.Buy, "option.focus", "focus-visible")},
				{Target: input, Props: []string{"position", "top", "left", "opacity", "margin-top", "margin-left", "display"}, Apple: h.AppleProbe(apple.Buy, "option.input")},
				{Target: ".cui-option-title", Props: h.Size, Apple: h.AppleProbe(apple.Buy, "option.title")},
				{Target: ".cui-option-title", Props: []string{"letter-spacing", "box-x"}, Golden: "at the tile's start, 16px in (its edge and padding), the system font's letter-spacing"},
				{Target: ".cui-option", Props: []string{"background-color", "border-top-left-radius", "margin-top"}, Golden: "its own light ground, the label's corners, as a field's; no margin above (apple.com's 14px, the form's spacing, buy:option.box)"},
				{Target: label, Props: []string{"display", "flex-direction", "align-items"}, Golden: "its title centered down the tile (apple.com's a flex column holding the row)"},
			},
		},
		{
			Name:   "checked",
			Widths: h.With1600,
			HTML:   h.Example("checked"),
			Checks: []h.Check{
				{Target: label, Props: h.Props(h.Size, []string{"color"}, h.Padding, edge, []string{"background-color", "min-height", "box-height"}), Apple: h.AppleProbe(apple.Buy, "option.checked")},
				{Target: ".cui-option-detail", Props: h.Props(h.Size, []string{"display", "margin-top"}), Apple: h.AppleProbe(apple.Buy, "option.detail")},
				{Target: ".cui-option-title", Props: []string{"box-width"}, Golden: "the rest of the row: nothing beside it"},
			},
		},
		{
			Name: "dark",
			HTML: `<div class="cui-dark">` + h.Example("option") + `</div>`,
			Checks: []h.Check{
				{Target: label, Props: []string{"color", "background-color", "border-top-color"}, Golden: "light in a dark block, as a field (apple.com shows options on light pages alone)"},
			},
		},
		{
			Name:   "disabled",
			HTML:   strings.Replace(h.Example("option"), `value="256"`, `value="256" disabled`, 1),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: label, Props: []string{"opacity", "border-top-color"}, Golden: "disabled, dimmed to 0.42 with its rest edge: apple.com's buy page, read 2026-10-31 (its disabled tiles 0.42, the base rule; a whole disabled step keeps 1, the page's choice)"},
				{Target: input, Props: []string{"cursor"}, Golden: "disabled, the arrow (apple.com's disabled tiles' labels default)"},
			},
		},
		{
			Name:   "disabled checked",
			HTML:   strings.Replace(h.Example("checked"), " checked", " checked disabled", 1),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: label, Props: []string{"opacity", "border-top-color", "border-top-width"}, Golden: "chosen and disabled, dimmed, its 2px edge the rest one's gray: the disabled rule after the checked one, as apple.com's are ordered"},
			},
		},
		{
			Name:   "pointer",
			HTML:   h.Example("option"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-option", Props: []string{"position"}, Golden: "what its input is placed against"},
				{Target: input, Props: []string{"cursor"}, Golden: "the pointer over a tile, apple.com's (its enabled tiles' labels and inputs pointer, read 2026-10-31); the input over the tile takes it"},
				{Target: label, Props: []string{"hit"}, Golden: "the radio over the whole tile: the pointer finds it at the tile's corners and middle, so it takes the click"},
			},
		},
	})
}
