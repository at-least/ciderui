package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The gift card's buy page, for the store's form parts a bag is not needed
// to reach (its checkout's are not: a headless browser's Add to Bag is refused,
// 2026-10-20): its Select button disabled before its choices, its delivery fields
// disabled before the design and the amount; its sections; the choices that enable the
// fields (picked once, at the first width: Email, the first design, Other Amount, then
// $50); the amount's field with its currency before it, while Other Amount is chosen; a
// plain field enabled; an empty required field the focus has left (clicked, then the
// focus taken off: nothing typed), its error and its message. Read only: nothing typed,
// nothing sent. Email is clicked on its label: from 2026-11-13 the
// delivery's label covers its radio, which takes no click (the design's and the
// amount's radios still lie over their labels).
func giftcard() pagespec.Spec {
	FIELDS := ".rf-giftcard-delivery-section-fields"
	NAME := FIELDS + " .form-textbox:first-child"
	EMAIL := FIELDS + " .form-textbox:nth-child(2)"
	OTHER := ".rf-choose-amount-dimensions > div:nth-child(4) .form-selector-input"
	PREFIX := ".rf-custom-input"
	AMOUNT := ".rf-choose-amount-dimensions > div:nth-child(2) .form-selector-input"
	EDGE := []string{"border-top-width", "border-top-style", "border-top-color", "border-right-width", "border-right-style", "border-right-color", "border-top-left-radius", "border-bottom-left-radius", "border-top-right-radius"}

	return pagespec.Spec{
		URL:  "https://www.apple.com/shop/buy-giftcard/giftcard",
		Wait: 2500,
		Probes: []pagespec.Probe{
			// @external: before a choice: its button disabled (the first width alone, as the
			// fields below: the choices after them enable them)
			{Name: "button.disabled", Select: "button.button-block[disabled]", Props: []string{"background-color", "color", "opacity", "cursor", "border-top-color"}, Widths: []int{1600}},
			// @external: a section, its rule and its title
			{Name: "section", Select: ".rf-giftcard-selection-group", Props: slices.Concat(BOX, []string{"border-top-width", "border-top-style", "border-top-color"}), Box: true},
			{Name: "section.title", Select: ".rf-choose-amount-section-title", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			// @external: the choices that enable the fields, each clicked once
			{Name: "choose.type", Click: "input[name=dimensionGiftCardType][value=electronic] + label", Select: "input[name=dimensionGiftCardType][value=electronic]", Props: []string{"opacity"}},
			// @external: its delivery fields shown, disabled until a design and an amount are
			// chosen
			{Name: "field.disabled", Select: NAME + " .form-textbox-input", Props: slices.Concat(TYPE, []string{"background-color", "border-top-color", "border-top-width", "opacity", "cursor", "height"}), Widths: []int{1600}},
			{Name: "label.disabled", Select: NAME + " .form-textbox-label", Props: slices.Concat(TYPE, []string{"opacity"}), Widths: []int{1600}},
			{Name: "choose.design", Click: "input[name=dimensionGiftCardTheme]", Select: "input[name=dimensionGiftCardTheme]", Props: []string{"opacity"}},
			{Name: "choose.other", Click: OTHER, Select: OTHER, Props: []string{"opacity"}},
			// @external: the amount's field, its currency before it (enabled while Other
			// Amount is chosen: at the first width alone, $50 chosen after it)
			{Name: "prefix", Select: PREFIX, Props: []string{"position", "display", "height"}, Box: true, Position: true, Widths: []int{1600}},
			{Name: "prefix.currency", Select: PREFIX + " .form-textbox-countrycode", Props: slices.Concat(TYPE, BOX, EDGE, []string{"position", "top", "left", "width", "height", "display", "align-items", "justify-content", "background-color", "z-index"}), Box: true, Position: true, Widths: []int{1600}},
			{Name: "prefix.input", Select: PREFIX + " .form-textbox-input", Props: slices.Concat(TYPE, BOX, EDGE, []string{"height", "background-color"}), Box: true, Position: true, States: []string{"focus"}, Widths: []int{1600}},
			{Name: "prefix.label", Select: PREFIX + " .form-textbox-label", Props: slices.Concat(TYPE, []string{"position", "top", "left", "padding-left"}), Box: true, Position: true, Widths: []int{1600}},
			// @external: an amount, which enables the delivery fields
			{Name: "choose.amount", Click: AMOUNT, Select: AMOUNT, Props: []string{"opacity"}},
			// @external: a plain field, enabled (the first width alone: after it, the error
			// below has the page's script mark the delivery fields invalid, and it
			// read so at 1200, 900 and 390px, 2026-10-20)
			{Name: "field", Select: NAME + " .form-textbox-input", Props: slices.Concat(TYPE, BOX, []string{"height", "background-color", "border-top-color", "border-top-width"}), Box: true, Widths: []int{1600}},
			// @external: an empty required field the focus left: clicked, then the focus taken off
			{Name: "error.focus", Click: EMAIL + " .form-textbox-input", Select: EMAIL + " .form-textbox-input", Props: []string{"opacity"}},
			{Name: "error", Blur: true, Select: EMAIL + " .form-textbox-input", Props: slices.Concat(TYPE, BOX, []string{"height", "background-color", "border-top-color", "border-top-width"}, FOCUS, MOTION), Box: true, Position: true, States: []string{"focus"}},
			{Name: "error.label", Select: EMAIL + " .form-textbox-label", Props: slices.Concat(TYPE, []string{"top", "left"}), Box: true, Position: true},
			{Name: "error.wrapper", Select: EMAIL + " .form-message-wrapper", Props: slices.Concat(BOX, []string{"display", "align-items"}), Box: true, Position: true},
			{Name: "error.message", Select: EMAIL + " .form-message", Props: slices.Concat(TYPE, BOX, []string{"display", "align-items"}), Box: true, Position: true},
			// @external: its icon, the wrapper's ::before: an SF Pro Icons glyph (U+F312)
			{Name: "error.icon", Select: EMAIL + " .form-message-wrapper", Pseudo: "::before", Props: []string{"content", "font-family", "font-size", "line-height", "width", "height", "padding-right", "margin-top", "color", "display", "position", "top"}},
		},
	}
}
