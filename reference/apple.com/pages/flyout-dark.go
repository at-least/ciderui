package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: The global nav's flyouts on a dark page, iPhone 18 Pro's (its bar clear,
// product:nav): flyout.go's probes, at 1600px. The Newsroom's dark nav
// opens none (`globalnav-submenus-disabled` on its nav, 2026-10-21).
func flyoutDark() pagespec.Spec {
	s := flyout()
	s.URL = "https://www.apple.com/iphone-18-pro/"
	s.Widths = []int{1600}
	return s
}
