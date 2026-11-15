package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: A dark product page modal (ric-modal theme-dark), MacBook Pro's "Explore
// the M5 family of chips", opened (its button clicked once, at the first
// width): its scrim (on the layer itself here, its ::after empty), its
// panel's ground and text, a heading's color, its close button's
// circle at rest, hovered, pressed and focused. Its panel is the larger kind
// (l2-modal, 1260px), so its colors alone are read. Read only.
func ricmodalDark() pagespec.Spec {
	M := "#performance-modal"

	return pagespec.Spec{
		URL:    "https://www.apple.com/macbook-pro/",
		Widths: []int{1600, 390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "modal", Click: "[data-modal-open=\"performance-modal\"]", Select: M, Props: slices.Concat([]string{"position", "color-scheme", "background-color", "background-image"}, BLUR)},
			{Name: "scrim", Select: M, Pseudo: "::after", Props: slices.Concat([]string{"background-color", "background-image"}, BLUR, []string{"opacity"})},
			{Name: "content", Select: M + " .ric-modal-content-wrapper", Props: []string{"background-color", "color"}},
			{Name: "heading", Select: M + " .ric-modal-content-wrapper h2", Props: []string{"color"}},
			{Name: "close.icon", Select: M + " .ric-modal-close-icon", Props: []string{"background-color", "color", "border-top-left-radius"}},
			{Name: "close.icon-states", Select: M + " .ric-modal-close-icon", StateOn: M + " .ric-modal-close-button", Props: []string{"background-color", "color", "outline-style", "outline-color", "outline-width", "outline-offset"}, States: []string{"hover", "active", "focus-visible"}},
		},
	}
}
