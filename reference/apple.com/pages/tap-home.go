package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: The home page on a touch screen, its menu's button tapped at 390px (`tap`,
// the DevTools protocol's touch events: a touch screen's tap, not a forced
// :hover): what the tap leaves hovered, against touch-home:nav.menutrigger
// (forced: 0.8 at rest, 1 hovered) and home:menu.trigger (the menu opened by
// a click, the pointer gone: 0.8). Read only.
func tapHome() pagespec.Spec {
	TRIGGER := "#globalnav .globalnav-menutrigger-button"

	return pagespec.Spec{
		URL:    home().URL,
		Wait:   2500,
		Touch:  true,
		Widths: []int{390},
		Probes: []pagespec.Probe{
			{Name: "menutrigger", Tap: TRIGGER, Select: TRIGGER, Props: []string{"color", "opacity"}, Matches: []string{":hover", ":active", ":focus", ":focus-visible"}},
			// @external: the menu the tap opened: the page held
			{Name: "root", Select: "html", Props: []string{"overflow-y"}},
		},
	}
}
