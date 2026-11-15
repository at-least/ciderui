package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

func retailProbes() []pagespec.Probe {
	FIELD := func(id, labelClass, inputClass string) string {
		return `<div class="form-textbox"><input type="text" class="form-textbox-input` + inputClass + `" id="` + id + `" required value=""><label class="form-textbox-label` + labelClass + `" for="` + id + `">Email</label></div>`
	}
	// @external: a field in error, as the gift card's page marks an empty required one the
	// focus has left (giftcard:error*): its classes, its message and its icon
	ERROR := "<div class=\"form-textbox is-error\"><input type=\"text\" class=\"form-textbox-input\" id=\"cui-error\" aria-invalid=\"true\" value=\"\"><label class=\"form-textbox-label\" for=\"cui-error\">Email</label><div><div class=\"is-error\"><div class=\"form-message-wrapper\" id=\"cui-error-message\"><span class=\"form-message\">Please enter a valid email.</span></div></div></div></div>"
	DISABLED := "<div class=\"form-textbox\"><input type=\"text\" class=\"form-textbox-input\" id=\"cui-disabled\" disabled value=\"\"><label class=\"form-textbox-label\" for=\"cui-disabled\">Email</label></div>"
	// @external: the multi-line field, the feedback form's markup (feedback:textarea*: a
	// bordered box, its textarea in it, its label after it), with the store's
	// stylesheet here: at rest, focused and in error (the classes apple.com's
	// script sets: the box's focused, the label's eyebrow, the field's is-error)
	AREA := func(cls, box, label string) string {
		return `<div class="form-textbox` + cls + `"><div class="form-textarea` + box + `"><textarea id="cui-area" aria-labelledby="cui-area-label"></textarea></div><span class="form-textbox-label` + label + `" id="cui-area-label">Comments</span></div>`
	}
	// @external: the store's dropdown (tradein:dropdown*), in error and disabled: the
	// trade-in's markup, nothing chosen, with the class apple.com's script sets
	// (is-error, as the feedback form's empty required dropdowns get it)
	DROPDOWN := func(cls, attrs string) string {
		return `<div class="form-dropdown` + cls + `"><label class="form-dropdown-label form-dropdown-label-placeholder" for="cui-dd">Choose your model</label><select id="cui-dd" class="form-dropdown-select form-dropdown-selectnone"` + attrs + `><option value="">Select</option></select><span class="form-dropdown-chevron" aria-hidden="true"></span></div>`
	}
	return []pagespec.Probe{
		{Name: "field.search", Select: "#store-search-input", Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "height"}, Box: true},
		{Name: "field", Inject: FIELD("cui-field", "", ""), Into: "main", Props: []string{"position", "height"}, Box: true},
		{Name: "field.input", Inject: FIELD("cui-field", "", ""), Into: "main", Select: "input", Props: slices.Concat(TYPE, BOX, []string{"height", "box-sizing"}, FOCUS, MOTION), Box: true, States: []string{"hover", "focus", "focus-visible", "focus+focus-visible"}},
		{Name: "field.label", Inject: FIELD("cui-field", "", ""), Into: "main", Select: "label", Props: slices.Concat(TYPE, []string{"position", "top", "left", "transform", "pointer-events", "white-space"}, MOTION)},
		{Name: "field.label-floated", Inject: FIELD("cui-field", " form-textbox-eyebrow-label", " form-textbox-entered"), Into: "main", Select: "label", Props: slices.Concat(TYPE, []string{"position", "top", "left"})},
		{Name: "field.error", Inject: ERROR, Into: "main", Select: "input", Props: []string{"color", "background-color", "border-top-color", "border-top-width"}},
		{Name: "field.error-label", Inject: ERROR, Into: "main", Select: "label", Props: []string{"color"}},
		{Name: "field.error-message", Inject: ERROR, Into: "main", Select: ".form-message", Props: TYPE},
		{Name: "field.error-icon", Inject: ERROR, Into: "main", Select: ".form-message-wrapper", Pseudo: "::before", Props: []string{"content", "font-family", "font-size", "line-height", "color", "padding-right", "display"}},
		{Name: "field.disabled", Inject: DISABLED, Into: "main", Select: "input", Props: []string{"color", "background-color", "border-top-color", "opacity", "cursor"}},
		{Name: "field.disabled-label", Inject: DISABLED, Into: "main", Select: "label", Props: []string{"color", "opacity"}},
		{Name: "textarea", Inject: AREA("", "", ""), Into: "main", Select: ".form-textarea", Props: slices.Concat(TYPE, BOX, []string{"height", "box-sizing", "display", "border-top-width", "border-top-style", "border-top-color", "border-top-left-radius", "outline-style"}, MOTION), Box: true},
		{Name: "textarea.input", Inject: AREA("", "", ""), Into: "main", Select: "textarea", Props: slices.Concat(TYPE, BOX, []string{"height", "min-height", "box-sizing", "resize", "outline-style", "border-top-width", "border-top-style"}), Box: true},
		{Name: "textarea.label", Inject: AREA("", "", ""), Into: "main", Select: ".form-textbox-label", Props: slices.Concat(TYPE, []string{"position", "top", "left", "pointer-events", "white-space"}, MOTION)},
		{Name: "textarea.focused", Inject: AREA("", " focused", " form-textbox-eyebrow-label"), Into: "main", Select: ".form-textarea", Props: slices.Concat(BOX, []string{"height", "border-top-width", "border-top-color", "outline-style"}), Box: true},
		{Name: "textarea.label-floated", Inject: AREA("", " focused", " form-textbox-eyebrow-label"), Into: "main", Select: ".form-textbox-label", Props: slices.Concat(TYPE, []string{"top", "left"})},
		{Name: "textarea.error", Inject: AREA(" is-error", "", ""), Into: "main", Select: ".form-textarea", Props: []string{"color", "background-color", "border-top-color", "border-top-width"}},
		{Name: "textarea.error-label", Inject: AREA(" is-error", "", ""), Into: "main", Select: ".form-textbox-label", Props: []string{"color"}},
		{Name: "textarea.disabled", Inject: AREA("", " disabled", ""), Into: "main", Select: ".form-textarea", Props: []string{"color", "background-color", "border-top-color"}},
		{Name: "textarea.disabled-input", Inject: AREA("", " disabled", ""), Into: "main", Select: "textarea", Props: []string{"color", "-webkit-text-fill-color"}},
		{Name: "dropdown.error", Inject: DROPDOWN(" is-error", ""), Into: "main", Select: "select", Props: []string{"background-color", "border-top-color", "border-top-width", "cursor"}, States: []string{"focus"}},
		{Name: "dropdown.error-label", Inject: DROPDOWN(" is-error", ""), Into: "main", Select: "label", Props: []string{"color"}},
		{Name: "dropdown.disabled", Inject: DROPDOWN("", " disabled"), Into: "main", Select: "select", Props: []string{"color", "background-color", "border-top-color", "cursor", "opacity"}},
		{Name: "dropdown.disabled-label", Inject: DROPDOWN("", " disabled"), Into: "main", Select: "label", Props: []string{"color"}},
		{Name: "dropdown.disabled-chevron", Inject: DROPDOWN("", " disabled"), Into: "main", Select: ".form-dropdown-chevron", Props: []string{"opacity"}},
		{Name: "field.input-filled", Inject: FIELD("cui-field", " form-textbox-eyebrow-label", " form-textbox-entered"), Into: "main", Select: "input", Props: slices.Concat(TYPE, BOX)},
	}
}

// @external: The retail page, for apple.com's own text field (its form-textbox, the
// store's forms' field): the page's store search (with its icons) and a
// plain field the probe adds with apple.com's classes, at rest, hovered,
// focused; its label floated (the class apple.com's script sets on it as the
// field is focused or filled, on a copy the probe adds) and the field
// filled (form-textbox-entered).
func retail() pagespec.Spec {
	return pagespec.Spec{
		URL:    "https://www.apple.com/retail/",
		Probes: retailProbes(),
	}
}
