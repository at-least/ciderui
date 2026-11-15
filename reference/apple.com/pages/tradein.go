package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The store's dropdown (its form-dropdown), on iPhone 18 Pro's buy page,
// where its trade-in asks for the model to trade: the buy page's choices
// picked once, at the first width (its size, color and capacity by their
// labels, which are over their inputs; Add a trade-in by its input, which is
// over its label), and the dropdown it shows, empty (its label the placeholder),
// hovered and focused; as a copy beside it with a model chosen (its label
// small at its top: the classes apple.com's script sets). Read only: nothing
// chosen in the page's own dropdown, nothing sent.
func tradein() pagespec.Spec {
	SECTION := ".rf-tradeupinline-form-section"
	DD := SECTION + " .form-dropdown:not([data-cui-injected])"
	CHOSEN := "<div class=\"form-dropdown\"><label class=\"form-dropdown-label\" for=\"cui-dropdown\">Choose your model</label><select id=\"cui-dropdown\" name=\"cui-dropdown\" class=\"form-dropdown-select\"><option>iPhone 17 Pro</option></select><span class=\"form-dropdown-chevron\" aria-hidden=\"true\"></span></div>"
	EDGE := []string{"border-top-width", "border-top-style", "border-top-color", "border-top-left-radius"}

	return pagespec.Spec{
		URL:  "https://www.apple.com/shop/buy-iphone/iphone-18-pro",
		Wait: 2500,
		Probes: []pagespec.Probe{
			{Name: "choose.size", Click: "input[name=dimensionScreensize] + label", Select: "input[name=dimensionScreensize]", Props: []string{"opacity"}},
			{Name: "choose.color", Click: "input[name=dimensionColor] + label", Select: "input[name=dimensionColor]", Props: []string{"opacity"}},
			{Name: "choose.capacity", Click: "input[name=dimensionCapacity] + label", Select: "input[name=dimensionCapacity]", Props: []string{"opacity"}},
			{Name: "choose.tradein", Click: "input[name=tradeupinline][value=tradeIn]", Select: "input[name=tradeupinline][value=tradeIn]", Props: []string{"opacity"}},
			{Name: "dropdown", Select: DD, Props: []string{"position", "display", "height", "margin-bottom", "color-scheme"}, Box: true, Position: true},
			{Name: "dropdown.select", Select: DD + " .form-dropdown-select", Props: slices.Concat(TYPE, BOX, EDGE, []string{"height", "box-sizing", "background-color", "appearance", "-webkit-appearance", "cursor", "text-overflow"}, FOCUS, MOTION), Box: true, Position: true, States: []string{"hover", "focus", "focus-visible"}},
			{Name: "dropdown.label", Select: DD + " .form-dropdown-label", Props: slices.Concat(TYPE, []string{"position", "top", "left", "right", "pointer-events", "white-space", "overflow-x", "text-overflow"}, MOTION), Box: true, Position: true},
			{Name: "dropdown.chevron", Select: DD + " .form-dropdown-chevron", Props: []string{"position", "top", "right", "width", "height", "background-size", "background-position-x", "background-position-y", "background-repeat", "pointer-events", "opacity"}, Box: true, Position: true},
			{Name: "dropdown.chosen", Inject: CHOSEN, Into: SECTION, Select: ".form-dropdown-select", Props: slices.Concat(TYPE, BOX, []string{"height"})},
			{Name: "dropdown.chosen-label", Inject: CHOSEN, Into: SECTION, Select: ".form-dropdown-label", Props: slices.Concat(TYPE, []string{"top", "left"})},
		},
	}
}
