package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The store's sign-in, Apple Account's widget in its iframe (idmsa.apple.com):
// the title, a text field and its floating label at rest (the page focuses
// the field on load: blurred first) and focused, the checkbox (its parts,
// checked and focused on a copy the probe adds), the continue button, a link.
func signin() pagespec.Spec {
	F := "idmsa.apple.com"
	CHECKED := "<div class=\"form-checkbox\"><input type=\"checkbox\" id=\"cui-checked\" class=\"form-checkbox-input\" checked><label for=\"cui-checked\" class=\"form-label\"><span class=\"form-checkbox-indicator\" aria-hidden=\"true\"></span>Remember me</label></div>"
	return pagespec.Spec{
		URL:  "https://secure.store.apple.com/shop/account/home",
		Wait: 5000,
		Probes: []pagespec.Probe{
			{Name: "title", Frame: F, Select: ".si-container-title", Props: TYPE},
			{Name: "field", Frame: F, Select: ".form-textbox", Props: slices.Concat(BOX, []string{"width", "height", "position"}), Box: true},
			{Name: "field.input", Frame: F, Blur: true, Select: "#account_name_text_field", Props: slices.Concat(TYPE, BOX, []string{"height"}, FOCUS, MOTION), Box: true, States: []string{"focus", "focus-visible", "hover"}},
			{Name: "field.label", Frame: F, Blur: true, Select: "#apple_id_field_label", StateOn: "#account_name_text_field", Props: slices.Concat(TYPE, []string{"position", "top", "left", "transform", "transform-origin"}, MOTION), States: []string{"focus"}},
			{Name: "checkbox", Frame: F, Select: ".form-checkbox-indicator", Props: slices.Concat(BOX, []string{"width", "height", "position", "top", "left", "display"}, MOTION), Box: true},
			{Name: "checkbox.label", Frame: F, Select: "#remember-me-label", Props: slices.Concat(TYPE, []string{"padding-top", "padding-left", "position", "display"})},
			{Name: "checkbox.box", Frame: F, Select: ".form-checkbox", Props: []string{"position", "margin-bottom"}, Box: true},
			{Name: "checkbox.input", Frame: F, Select: "#remember-me", Props: []string{"position", "top", "left", "width", "height", "opacity", "margin-top", "margin-left"}},
			// @external: checked, and focused from the keyboard: apple.com's classes on a copy the
			// probe adds beside the widget's own (which is not checked until clicked)
			{Name: "checkbox.checked", Frame: F, Inject: CHECKED, Into: ":has(> .form-checkbox)", Select: ".form-checkbox-indicator", StateOn: "#cui-checked", Props: slices.Concat(BOX, []string{"color"}, FOCUS), States: []string{"focus-visible", "hover"}},
			{Name: "checkbox.check", Frame: F, Inject: CHECKED, Into: ":has(> .form-checkbox)", Select: ".form-checkbox-indicator", Pseudo: "::after", Props: []string{"content", "font-family", "font-size", "line-height", "position", "top", "left", "width", "height", "color"}},
			{Name: "checkbox.unchecked-focus", Frame: F, Select: ".form-checkbox-indicator", StateOn: "#remember-me", Props: FOCUS, States: []string{"focus-visible"}},
			{Name: "continue", Frame: F, Blur: true, Select: "#sign-in", Props: slices.Concat(TYPE, BOX, []string{"width", "height", "opacity"}, FOCUS), Box: true, States: []string{"hover", "focus-visible"}},
			{Name: "link", Frame: F, Select: "#iforgot-link", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
		},
	}
}
