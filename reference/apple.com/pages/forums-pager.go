package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: Apple Developer Forums' pager on a middle page (`/forums/latest?page=3`),
// signed out: the nav, its row, the first page's button (a word), the previous and
// next pages' (a chevron each, hovered and focused), the page's number between.
// A direction with no page is not drawn (`../../developer-forums/`, page 1's).
func forumsPager() pagespec.Spec {
	NAV := "nav.pagination"
	BUTTON := slices.Concat(TYPE, BOX, []string{"height", "display", "cursor", "text-decoration-line"})
	return pagespec.Spec{
		URL:    "https://developer.apple.com/forums/latest?page=3",
		Wait:   3000,
		Widths: []int{1200, 390},
		Probes: []pagespec.Probe{
			{Name: "nav", Select: NAV, Props: BOX, Box: true},
			{Name: "list", Select: NAV + " ol", Props: slices.Concat(FLEX, []string{"margin-left", "margin-right"}), Box: true},
			{Name: "first", Select: NAV + " a.first-page", Props: BUTTON, Box: true, States: []string{"hover", "active", "focus-visible"}},
			{Name: "previous", Select: NAV + " a.previous-page", Props: BUTTON, Box: true, States: []string{"hover", "active", "focus-visible"}},
			{Name: "previous.icon", Select: NAV + " a.previous-page sf-symbol", Props: []string{"width", "height", "font-size", "vertical-align"}, Box: true},
			{Name: "page", Select: NAV + " .page-number", Props: slices.Concat(TYPE, BOX, []string{"display"}), Box: true},
			{Name: "next", Select: NAV + " a.next-page", Props: BUTTON, Box: true},
		},
	}
}
