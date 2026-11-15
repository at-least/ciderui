//go:build browser

package btn

import (
	"slices"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	flexGolden = "inline-flex, not apple.com's inline-block: an icon half an em before its label, on its middle (apple.com's pill holds text alone)"
	ownGolden  = "the kit's stack, not apple.com's"
)

var (
	look     = h.Props(h.Size, h.Padding, h.Border, h.Radius, []string{"min-width", "box-height"}, slices.DeleteFunc(slices.Clone(h.Text), func(p string) bool { return p == "display" }))
	flexProp = []string{"display", "align-items", "justify-content", "column-gap"}
	// @external: what the home page's probes read (no text-align, wrapping or opacity; its
	// primary button's cursor alone); not its display, which is its place's
	lookHome = h.Props(h.Size, h.Padding, h.Border, h.Radius, []string{"min-width", "box-height"})
	ownProp  = []string{"font-family"}
)

func flex(target string) h.Check {
	return h.Check{Target: target, Props: flexProp, Golden: flexGolden}
}

func own(target string) h.Check {
	return h.Check{Target: target, Props: ownProp, Golden: ownGolden}
}

func states(target, page, probe string, props []string) []h.Check {
	return []h.Check{
		{Target: target, State: []string{"hover"}, Props: props, Apple: h.AppleProbe(page, probe, "hover")},
		{Target: target, State: []string{"active"}, Props: props, Apple: h.AppleProbe(page, probe, "active")},
	}
}

func TestCases(t *testing.T) {
	var disabled []h.Check
	for _, v := range [][2]string{{"", "button.disabled"}, {"secondary", "button.secondary.disabled"}} {
		target := ".cui-btn:not([data-variant])"
		if v[0] != "" {
			target = `.cui-btn[data-variant="` + v[0] + `"]`
		}
		disabled = append(disabled,
			// @external: a <button>: the class over the preflight's reset of the element
			h.Check{Target: target, Props: h.Props(look, h.Colors), Apple: h.AppleProbe("store", v[1])},
			h.Check{Target: target, State: []string{"hover"}, Props: h.Colors, Golden: "a disabled button keeps its ground under the pointer, where apple.com's changes it"},
		)
	}
	cases := []h.Case{
		{
			Name: "default",
			HTML: h.Example("default"),
			Checks: slices.Concat(
				[]h.Check{{Target: ".cui-btn", Props: h.Props(look, h.Colors), Apple: h.AppleProbe("store", "button")}},
				states(".cui-btn", "store", "button", h.Colors),
				[]h.Check{own(".cui-btn"), flex(".cui-btn")},
			),
		},
		{
			Name: "sizes",
			HTML: h.Example("sizes"),
			Checks: []h.Check{
				{Target: ".cui-btn:not([data-size])", Props: look, Apple: h.AppleProbe("store", "button")},
				{Target: `.cui-btn[data-size="lg"]`, Props: look, Apple: h.AppleProbe("store", "button.elevated")},
				own(`.cui-btn[data-size="lg"]`),
				flex(`.cui-btn[data-size="lg"]`),
			},
		},
		{
			Name: "variants",
			HTML: h.Example("variants"),
			Checks: slices.Concat(
				[]h.Check{{Target: `.cui-btn[data-variant="secondary"]`, Props: h.Props(look, h.Colors), Apple: h.AppleProbe("store", "button.secondary")}},
				states(`.cui-btn[data-variant="secondary"]`, "store", "button.secondary", h.Colors),
				[]h.Check{{Target: `.cui-btn[data-variant="secondary-neutral"]`, Props: h.Props(look, h.Colors), Apple: h.AppleProbe("store", "button.secondary-neutral")}},
				states(`.cui-btn[data-variant="secondary-neutral"]`, "store", "button.secondary-neutral", h.Colors),
			),
		},
		{
			Name:   "dark secondary-neutral",
			HTML:   `<section class="cui-dark"><a class="cui-btn" data-variant="secondary-neutral" href="#buy">Buy</a></section>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-btn", Props: h.Colors, Golden: "the neutral's dark side, near-white on no ground: apple.com's in a dark block not measured"},
				{Target: ".cui-btn", State: []string{"hover"}, Props: h.Colors, Golden: "filled with it hovered"},
			},
		},
		{
			Name:   "dark",
			HTML:   h.Example("dark"),
			Widths: []int{1600, 1200, 900},
			Checks: slices.Concat(
				[]h.Check{{Target: ".cui-btn:not([data-variant])", Props: h.Props(h.Size, h.Padding, h.Radius, []string{"min-width"}, h.Colors), Apple: h.AppleProbe("home", "tile.dark-button")}},
				states(".cui-btn:not([data-variant])", "home", "tile.dark-button", h.Colors),
				[]h.Check{{Target: `.cui-btn[data-variant="secondary"]`, Props: h.Props(h.Size, h.Padding, h.Radius, []string{"min-width"}, h.Colors), Apple: h.AppleProbe("home", "tile.dark-button-secondary")}},
				states(`.cui-btn[data-variant="secondary"]`, "home", "tile.dark-button-secondary", h.Colors),
			),
		},
		{
			Name:   "disabled",
			HTML:   h.Example("disabled"),
			Checks: disabled,
		},
	}
	for _, block := range [][2]string{{"block", "a.cui-btn-block"}, {"block-button", "button.cui-btn-block"}} {
		name, target := block[0], block[1]
		cases = append(cases, h.Case{
			Name: name,
			HTML: h.Example(name),
			Checks: []h.Check{
				{Target: target, Props: h.Props(h.Size, h.Padding, h.Radius, []string{"box-height"}, h.Colors), Apple: h.AppleProbe("store", "button.block")},
				{Target: target, Props: []string{"display", "box-width"}, Golden: "flex, not apple.com's block, so an icon keeps its place; the row's whole width, a <button> as a link"},
			},
		})
	}
	cases = append(cases,
		h.Case{
			Name:   "icon",
			HTML:   h.Example("icon"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-btn > svg", Props: []string{"box-width", "box-height"}, Golden: "the icon 1em square, the label's 14px"},
				{Target: ".cui-btn", Props: []string{"box-height"}, Golden: "the icon adds no height: the 36px pill"},
			},
		},
		h.Case{
			Name:   "touch target",
			HTML:   h.Example("sizes"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: ".cui-btn:not([data-size])", Pseudo: "::after", Props: []string{"height"}, Golden: "the 36px pill's target 44px tall"},
			},
		},
		h.Case{
			Name:   "no touch target with a mouse",
			HTML:   h.Example("sizes"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-btn:not([data-size])", Pseudo: "::after", Props: []string{"content"}, Golden: "no invisible overlay where no touch screen is: the pointer finds the label, not a sheet over it"},
			},
		},
	)
	h.Run(t, cases)
}
