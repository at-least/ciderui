package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: The iPhone page on a touch screen, its product gallery's next paddle tapped
// at 390px (`tap`, the DevTools protocol's touch events: a touch screen's
// tap, not a forced :hover): what the tap leaves hovered, against
// touch-iphone:paddle (forced: rgba(210, 210, 215, 0.64) at rest,
// rgba(223, 223, 227, 0.698) hovered). Read only.
func tapIphone() pagespec.Spec {
	// @external: the one touch-iphone:paddle reads, the first enabled
	PADDLE := "main .section-select .paddlenav-arrow-next"

	return pagespec.Spec{
		URL:    iphone().URL,
		Wait:   2500,
		Touch:  true,
		Widths: []int{390},
		Probes: []pagespec.Probe{
			// @external: the gallery's current item before the tap (its markup names it)
			{Name: "gallery.current-before", Select: "main .section-select .gallery-item.current", Props: []string{"display"}},
			{Name: "paddle", Tap: PADDLE, Select: PADDLE, Props: []string{"background-color", "color", "border-top-color", "opacity"}, Matches: []string{":hover", ":active", ":focus", ":focus-visible"}},
			// @external: the gallery the tap moved: its current item the next
			{Name: "gallery.current", Select: "main .section-select .gallery-item.current", Props: []string{"display"}},
		},
	}
}
