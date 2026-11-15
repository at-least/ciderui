package pages

import "github.com/at-least/ciderui/internal/pagespec"

// @external: What of the home page can be selected, for the global nav's: apple.com's
// stylesheet sets user-select: none on #globalnav and on #globalnav
// .globalnav-list (globalheader.css, its text read 2026-10-22), nowhere else
// in the nav, its search's field not excepted. Read here as each element
// computes it: the nav, its list, the logo, a link and its text, the bag,
// the menu's button, a closed flyout; the page's text and its footer beside
// them; then the search opened (its button clicked once, at each of the
// nav's layouts' widths read) and its field, last. Once it has opened the
// pointer moves into it, where it hovers none (search's). At 1600px "ma" is
// typed into the field (the owner's leave to type on apple.com, 2026-10-21;
// nothing sent), its placeholder gone if the field took the keys.
func navSelect() pagespec.Spec {
	SELECT := []string{"user-select"}
	FIELD := "#globalnav-submenu-search .globalnav-searchfield-input"

	return pagespec.Spec{
		URL:    "https://www.apple.com/",
		Widths: []int{1600, 390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "nav", Select: "#globalnav", Props: SELECT},
			{Name: "list", Select: "#globalnav .globalnav-list", Props: SELECT},
			{Name: "logo", Select: "#globalnav .globalnav-link-apple", Props: SELECT},
			{Name: "link", Select: "#globalnav .globalnav-item-mac .globalnav-link", Props: SELECT},
			{Name: "link-text", Select: "#globalnav .globalnav-item-mac .globalnav-link-text", Props: SELECT},
			{Name: "bag", Select: "#globalnav .globalnav-link-bag", Props: SELECT},
			{Name: "menutrigger", Select: "#globalnav .globalnav-menutrigger-button", Props: SELECT, Widths: []int{390}},
			{Name: "flyout", Select: "#globalnav-submenu-link-mac", Props: SELECT},
			{Name: "page", Select: "main", Props: SELECT},
			{Name: "footer", Select: "#ac-globalfooter", Props: SELECT},
			{Name: "search.field", Click: "#globalnav-menubutton-link-search", Pointer: pagespec.At(200, 430), Select: FIELD, Props: SELECT, Widths: []int{1600}},
			{Name: "search.typed", Type: []string{"m", "a"}, Typed: 1500, Select: FIELD, Props: SELECT, Matches: []string{":placeholder-shown"}, Widths: []int{1600}},
			{Name: "search.field-phone", Click: "#globalnav-menubutton-link-search", Pointer: pagespec.At(200, 850), Select: FIELD, Props: SELECT, Widths: []int{390}},
		},
	}
}
