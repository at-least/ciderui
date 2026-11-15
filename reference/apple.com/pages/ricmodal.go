package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The marketing pages' modal (ric-modal), the iPhone page's "Read more about
// Apple Trade In", opened (its button clicked once, at the first width; it
// stays open as the window narrows): the layer over the window and its
// scrim (its ::after: a dimming and a blur), the page held, the overlay
// centered in it, its corners (its content clipped round), its content's
// room, its close button (a circle in a 44px target) at rest, hovered,
// pressed and focused; every animation the click starts. It steps
// at the marketing pages' widths, 1069 and 735px. Read only.
func ricmodal() pagespec.Spec {
	M := "#trade-in-modal"
	SPACE := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-right", "margin-bottom", "margin-left"}

	return pagespec.Spec{
		URL:    "https://www.apple.com/iphone/",
		Widths: []int{1600, 1200, 1069, 1068, 900, 735, 734, 640, 390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "modal", Click: "button[data-modal-open=\"trade-in-modal\"]", Animations: M, Select: M, Props: []string{"position", "top", "left", "right", "bottom", "overflow-y", "z-index", "opacity", "display", "background-color"}, Box: true},
			{Name: "scrim", Select: M, Pseudo: "::after", Props: slices.Concat([]string{"content", "position", "top", "left", "right", "bottom", "background-color", "background-image"}, BLUR, []string{"opacity", "z-index"}, MOTION)},
			{Name: "root", Select: "html", Props: []string{"overflow-y"}},
			{Name: "container", Select: M + " .ric-modal-overlay-container", Props: slices.Concat(SPACE, []string{"display", "align-items", "justify-content", "min-height", "box-sizing"}), Box: true},
			{Name: "overlay", Select: M + " .ric-modal-overlay", Props: slices.Concat(SPACE, []string{"width", "max-width", "border-top-left-radius", "border-bottom-right-radius", "overflow-x", "overflow-y", "position", "z-index", "opacity", "transform", "background-color"}, MOTION), Box: true, Position: true},
			{Name: "content", Select: M + " .ric-modal-content-wrapper", Props: slices.Concat(SPACE, []string{"background-color", "color", "border-top-left-radius"}), Box: true, Position: true},
			{Name: "viewport", Select: M + " .ric-modal-viewport-content", Props: SPACE, Box: true, Position: true},
			{Name: "close.wrapper", Select: M + " .ric-modal-close-button-wrapper", Props: []string{"position", "top", "right", "left", "width", "height", "margin-left", "margin-right", "z-index", "display", "justify-content", "align-items", "order"}, Box: true, Position: true},
			{Name: "close", Select: M + " .ric-modal-close-button", Props: slices.Concat([]string{"width", "height"}, SPACE, []string{"display", "align-items", "justify-content", "background-color", "color", "cursor", "border-top-width"}, MOTION), Box: true, Position: true},
			// @external: its circle (an element between the button and its svg) and the svg
			{Name: "close.icon", Select: M + " .ric-modal-close-icon", Props: slices.Concat([]string{"width", "height", "display", "align-items", "position", "background-color", "color", "border-top-left-radius"}, MOTION), Box: true, Position: true},
			{Name: "close.svg", Select: M + " .ric-modal-close-icon svg", Props: []string{"width", "height", "fill", "position"}, Box: true, Position: true},
			// @invariant: hovered, pressed and focused last: the still style that holds a state
			// cancels the page's animations, and taking it off starts them again
			{Name: "close.icon-states", Select: M + " .ric-modal-close-icon", StateOn: M + " .ric-modal-close-button", Props: []string{"background-color", "color", "outline-style", "outline-color", "outline-width", "outline-offset"}, States: []string{"hover", "active", "focus-visible"}},
		},
	}
}
