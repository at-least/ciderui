package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The store's overlay, opened (clicked once, at the first width: it stays
// open as the window narrows, so its probes are a spec of their own): iPhone
// 18 Pro's buy page, its first "Show more" (Which model is right for you?),
// a modal dialog over the page: the window-wide layer that dims the page and
// scrolls the dialog (rc-overlay), the dialog (role=dialog, aria-modal), its
// white panel and its content's room, its close button (a 44px target, a
// 36px circle, a 20px X), hovered and focused; the page under it held. The
// store steps it at 736/737px, where its own media queries are (the panel
// 816px at most above, 640px below), not at 735; and the panel is the
// window's width below 816 and 640px.
func overlay() pagespec.Spec {
	SPACE := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-right", "margin-bottom", "margin-left"}

	return pagespec.Spec{
		URL:    "https://www.apple.com/shop/buy-iphone/iphone-18-pro",
		Widths: []int{1600, 1200, 900, 816, 737, 736, 735, 734, 640, 390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			// @external: the layer: fixed over the window, the page dimmed, fading in; it is
			// what scrolls when the dialog is taller than the window
			{Name: "overlay", Click: "a.rc-decisionsection-button", Animations: ".rc-overlay", Select: ".rc-overlay", Props: slices.Concat([]string{"position", "top", "left", "z-index", "display", "align-items", "justify-content", "overflow-x", "overflow-y", "overscroll-behavior", "background-color", "opacity"}, BLUR, MOTION), Box: true},
			// @external: the page under it, held still
			{Name: "root", Select: "html", Props: []string{"overflow-x", "overflow-y"}},
			{Name: "dialog", Select: ".rc-overlay [role=\"dialog\"][aria-modal=\"true\"]", Props: []string{"position", "width", "max-width", "margin-left", "margin-right", "outline-style"}, Box: true, Position: true},
			{Name: "panel", Select: ".rc-overlay-popup-outer", Props: slices.Concat(BOX, []string{"margin-right", "margin-left", "border-top-right-radius", "border-bottom-left-radius", "border-bottom-right-radius", "overflow-x", "overflow-y", "display", "flex-direction", "color"}), Box: true, Position: true},
			{Name: "content", Select: ".rc-overlay-popup-content", Props: SPACE, Box: true, Position: true},
			// @external: the close button: its target, its circle, its X
			{Name: "close", Select: ".rc-overlay-close", Props: slices.Concat([]string{"width", "height"}, SPACE, []string{"display", "align-items", "justify-content", "background-color", "border-top-width", "cursor", "z-index", "position"}), Box: true, Position: true, States: []string{"focus-visible"}},
			{Name: "close.focus", Select: ".rc-overlay-close", Props: FOCUS, States: []string{"focus-visible", "focus+focus-visible"}},
			{Name: "close.circle", Select: ".rc-overlay-close .rc-overlay-closesvg", Props: slices.Concat([]string{"width", "height", "background-color", "color", "border-top-left-radius"}, MOTION), Box: true, States: []string{"hover"}},
			{Name: "close.circle-hover", Select: ".rc-overlay-close .rc-overlay-closesvg", StateOn: ".rc-overlay-close", Props: []string{"background-color", "color"}, States: []string{"hover", "active"}},
			{Name: "close.icon", Select: ".rc-overlay-close svg", Props: []string{"width", "height", "fill", "top", "right", "position"}, Box: true},
		},
	}
}
