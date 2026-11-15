package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: The home page on a touch screen 1024px wide (a tablet's: the bar's links in
// a row), Mac's link in the bar tapped (`tap`, the DevTools protocol's touch
// events): apple.com's script, having seen a touch on the bar, opens its
// flyout in place of following the link. Read only.
func tapFlyout() pagespec.Spec {
	return pagespec.Spec{
		URL:    "https://www.apple.com/",
		Wait:   2500,
		Touch:  true,
		Widths: []int{1024},
		Probes: []pagespec.Probe{
			{Name: "link", Tap: "#globalnav .globalnav-link-mac", Select: "#globalnav .globalnav-link-mac", Props: []string{"color"}, Matches: []string{":hover", ":focus", ":focus-visible"}},
			{Name: "flyout", Select: "#globalnav-submenu-link-mac", Props: []string{"height", "visibility", "background-color"}, Box: true},
			{Name: "trigger", Select: "#globalnav-menubutton-link-mac", Props: []string{"display"}, Vars: []string{}, Box: true},
			{Name: "curtain", Select: "#globalnav-curtain", Props: []string{"opacity", "visibility"}},
			{Name: "page", Select: "link[rel=\"canonical\"]", Props: []string{}},
		},
	}
}
