package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The bag's accordion, opened (clicked once, at the first width: what it
// opens stays so, so these probes are a spec of their own): the outer item
// (the bag's questions, a large title) and its chevron, then, opened, its
// compact items (a question, its chevron), the first of those opened (its
// answer, the chevron's color).
func bagAccordion() pagespec.Spec {
	return pagespec.Spec{
		URL:  "https://www.apple.com/shop/bag",
		Wait: 2500,
		Probes: []pagespec.Probe{
			{Name: "accordion.title", Select: ".rc-accordion-item > h2 .rc-accordion-title", Props: TYPE},
			{Name: "accordion.icon", Select: ".rc-accordion-item > h2 .rc-accordion-chevrondown", Props: slices.Concat([]string{"color", "margin-left"}, MOTION), Box: true},
			{Name: "accordion.svg", Select: ".rc-accordion-item > h2 .accordion-icon-svg", Props: []string{"width", "height"}, Box: true},
			{Name: "accordion.line", Select: ".rc-accordion-item > h2 .accordion-icon-svg polyline", Props: []string{"stroke-width", "stroke-linecap", "stroke-linejoin"}},
			{Name: "accordion.open", Click: ".rc-accordion-item > h2 .rc-accordion-button", Select: ".rc-accordion-item > h2 .rc-accordion-chevrondown", Props: []string{"color"}},
			{Name: "accordion.panel", Select: ".rc-accordion-item > .r-height-transition-enter-done", Props: slices.Concat(MOTION, []string{"overflow"}), Box: true},
			{Name: "compact.item", Select: ".rc-accordion-compact > .rc-accordion-item", Props: BOX, Box: true},
			{Name: "compact.button", Select: ".rc-accordion-compact > .rc-accordion-item .rc-accordion-button", Props: []string{"padding-top", "padding-bottom", "display", "text-align", "cursor"}, Box: true},
			{Name: "compact.title", Select: ".rc-accordion-compact .rc-accordion-title", Props: TYPE},
			{Name: "compact.icon", Select: ".rc-accordion-compact .rc-accordion-chevrondown", Props: slices.Concat([]string{"color"}, MOTION), Box: true, States: []string{"hover"}},
			// @external: its button hovered: the chevron's color (the button's hover, not the icon's)
			{Name: "compact.icon-hover", Select: ".rc-accordion-compact .rc-accordion-chevrondown", StateOn: ".rc-accordion-compact > .rc-accordion-item .rc-accordion-button", Props: []string{"color"}, States: []string{"hover"}},
			{Name: "compact.svg", Select: ".rc-accordion-compact .accordion-icon-svg", Props: []string{"width", "height"}, Box: true},
			{Name: "compact.line", Select: ".rc-accordion-compact .accordion-icon-svg polyline", Props: []string{"stroke-width", "stroke-linecap", "stroke-linejoin"}},
			{Name: "compact.open", Click: ".rc-accordion-compact > .rc-accordion-item:first-child .rc-accordion-button", Select: ".rc-accordion-compact > .rc-accordion-item:first-child .rc-accordion-chevrondown", Props: []string{"color"}},
			{Name: "compact.panel", Select: ".rc-accordion-compact > .rc-accordion-item:first-child > div", Props: slices.Concat([]string{"margin-top"}, MOTION), Box: true},
			{Name: "compact.answer", Select: ".rc-accordion-compact > .rc-accordion-item:first-child .rc-accordion-content", Props: slices.Concat(TYPE, []string{"padding-top", "padding-right", "padding-bottom", "padding-left"}), Box: true},
		},
	}
}
