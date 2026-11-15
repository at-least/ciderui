package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: apple.com's Newsroom, a light page under a dark global nav: the bar's dark
// material (the home page's is light), its links and icons, its phone menu
// opened.
func newsroom() pagespec.Spec {
	return pagespec.Spec{
		URL: "https://www.apple.com/newsroom/",
		Probes: []pagespec.Probe{
			{Name: "nav", Select: "#globalnav", Props: slices.Concat([]string{"height", "position", "background-color"}, BLUR), Vars: []string{"--globalnav-background", "--globalnav-backdrop-filter"}, Box: true},
			{Name: "nav.content", Select: "#globalnav .globalnav-content", Props: slices.Concat([]string{"background-color"}, BLUR)},
			{Name: "nav.link", Select: "#globalnav .globalnav-item:not(.globalnav-item-apple) .globalnav-link", Props: slices.Concat(TYPE, MOTION), States: []string{"hover"}},
			{Name: "nav.search", Select: "#globalnav .globalnav-link-search", Props: []string{"color"}, States: []string{"hover"}},
			{Name: "nav.bag", Select: "#globalnav .globalnav-link-bag", Props: []string{"color"}, States: []string{"hover"}},
			{Name: "nav.menutrigger", Select: "#globalnav .globalnav-menutrigger-button", Props: []string{"color", "opacity"}, Widths: []int{900, 390}, States: []string{"hover"}},
			// @invariant: the phone's menu, opened: last, as the page stays open after the click
			{Name: "menu.open", Click: "#globalnav-menutrigger-button", Select: "#globalnav", Props: slices.Concat([]string{"background-color"}, BLUR), Widths: []int{390}},
			{Name: "menu.content", Select: "#globalnav .globalnav-content", Props: slices.Concat([]string{"background-color"}, BLUR), Widths: []int{390}},
			{Name: "menu.link", Select: "#globalnav .globalnav-menu-list .globalnav-link", Props: TYPE, Widths: []int{390}, States: []string{"hover"}},
			{Name: "menu.trigger", Select: "#globalnav .globalnav-menutrigger-button", Props: []string{"color", "opacity"}, Widths: []int{390}},
		},
	}
}
