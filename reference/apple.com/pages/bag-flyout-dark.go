package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The bag's flyout on a dark page (iPhone 18 Pro's: its bar clear, the
// flyout dark), opened as bag-flyout.go opens it (the pointer into its
// room, 100, 200), at 1600 and 390px: the
// bar, the flyout's ground, its message, line, link, heading and profile
// links (hovered and focused last), the curtain.
func bagFlyoutDark() pagespec.Spec {
	FLY := "#globalnav-submenu-bag"
	LINKED := FLY + " .ac-gn-bagview-nav-link-orders"

	return pagespec.Spec{
		URL:    "https://www.apple.com/iphone-18-pro/",
		Widths: []int{1600, 390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "nav", Click: "#globalnav-menubutton-link-bag", Pointer: pagespec.At(100, 200), Select: "#globalnav", Props: slices.Concat([]string{"background-color"}, BLUR), Widths: []int{1600}},
			{Name: "nav.phone", Click: "#globalnav-menubutton-link-bag", Pointer: pagespec.At(200, 850), Select: "#globalnav", Props: slices.Concat([]string{"background-color"}, BLUR), Widths: []int{390}},
			{Name: "nav.content", Select: "#globalnav .globalnav-content", Props: slices.Concat([]string{"background-color"}, BLUR)},
			{Name: "link.bag", Select: "#globalnav-menubutton-link-bag", Props: []string{"color", "opacity"}},
			{Name: "menutrigger", Select: "#globalnav-menutrigger-button", Props: []string{"color", "opacity"}, Widths: []int{390}},
			{Name: "flyout", Select: FLY, Props: slices.Concat([]string{"background-color"}, BLUR, []string{"color"})},
			{Name: "bag.header", Select: FLY + " .ac-gn-bagview-header", Props: TYPE},
			{Name: "bag.subheader", Select: FLY + " .ac-gn-bagview-subheader", Props: TYPE},
			{Name: "bag.nav-header", Select: FLY + " .ac-gn-bagview-nav-item-header", Props: TYPE},
			{Name: "curtain", Select: "#globalnav-curtain", Props: slices.Concat([]string{"background-color"}, BLUR, []string{"opacity", "visibility"})},
			{Name: "bag.subheader-link", Select: FLY + " .ac-gn-bagview-subheader a", Props: slices.Concat(TYPE, LINK, FOCUS), States: []string{"hover", "focus-visible"}},
			{Name: "bag.nav-link", Select: LINKED, Props: slices.Concat(TYPE, []string{"background-color"}, FOCUS), States: []string{"hover", "focus-visible"}},
			{Name: "bag.nav-text", Select: LINKED + " .ac-gn-bagview-nav-text", Props: []string{"color"}, States: []string{"hover"}},
			{Name: "bag.nav-icon", Select: LINKED + " .ac-gn-bagview-nav-image-regular svg", StateOn: LINKED, Props: []string{"color", "fill"}, States: []string{"hover"}, Widths: []int{1600}},
			{Name: "bag.nav-icon-compact", Select: LINKED + " .ac-gn-bagview-nav-image-compact svg", StateOn: LINKED, Props: []string{"color", "fill"}, States: []string{"hover"}, Widths: []int{390}},
		},
	}
}
