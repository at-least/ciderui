package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: The retail page's field in a dark color scheme (prefers-color-scheme:
// dark): apple.com's forms keep their light ground and text; a field in
// error and a disabled one, as copies the probe adds.
func retailDark() pagespec.Spec {
	return pagespec.Spec{
		URL:         "https://www.apple.com/retail/",
		ColorScheme: "dark",
		Probes:      pick(retailProbes(), "field.input", "field.label", "field.error", "field.error-label", "field.error-message", "field.error-icon", "field.disabled", "field.disabled-label", "textarea", "textarea.label", "textarea.error", "dropdown.error", "dropdown.disabled"),
	}
}
