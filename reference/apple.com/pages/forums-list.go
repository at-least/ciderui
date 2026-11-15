package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: Apple Developer Forums' post list (`/forums/latest`), signed out: a row's
// title, excerpt, tags and stats, the list's rule, and the count over a tag's list.
// The site draws one list for a wide window (`.desktop-view`) and another to its
// phone width (`.mobile-view`), each hidden at the other's widths.
func forumsList() pagespec.Spec {
	D := ".content-list-wrapper.desktop-view article.article-entry"
	M := ".content-list-wrapper.mobile-view article.article-entry"
	CLAMP := []string{"display", "-webkit-line-clamp", "-webkit-box-orient", "overflow", "text-overflow", "white-space"}
	return pagespec.Spec{
		URL:    "https://developer.apple.com/forums/latest",
		Wait:   3000,
		Widths: []int{1600, 1200, 900, 735, 734, 390},
		Probes: []pagespec.Probe{
			{Name: "page", Select: "body", Props: PAGE},
			{Name: "title", Select: "h1.overview-title", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "list", Select: ".content-list-wrapper.desktop-view", Props: slices.Concat(BOX, []string{"display", "width"}), Box: true},
			{Name: "row", Select: D, Props: slices.Concat(BOX, FLEX, []string{"grid-template-columns"}), Box: true, States: []string{"hover"}},
			{Name: "row.second", Select: D + " + article.article-entry", Props: BOX, Box: true},
			{Name: "row.summary", Select: D + " .article-summary", Props: slices.Concat(BOX, FLEX, []string{"flex-grow", "flex-basis", "width"}), Box: true},
			{Name: "row.title", Select: D + " a.post-title", Props: slices.Concat(TYPE, LINK), States: []string{"hover", "focus-visible"}},
			{Name: "row.title.text", Select: D + " a.post-title span.underline", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}, StateOn: D + " a.post-title"},
			{Name: "row.excerpt", Select: D + " .excerpt", Props: slices.Concat(TYPE, CLAMP, []string{"margin-top", "padding-top"}), Box: true},
			{Name: "row.chiclets", Select: D + " .chiclets", Props: slices.Concat(FLEX, []string{"margin-top", "padding-top"}), Box: true},
			{Name: "row.tag", Select: D + " a.tag", Props: slices.Concat(TYPE, BOX, []string{"text-transform"}), Box: true, States: []string{"hover"}},
			{Name: "row.stats", Select: D + " .article-stats", Props: slices.Concat(BOX, FLEX, []string{"width", "flex-shrink", "grid-template-columns"}), Box: true},
			{Name: "row.stat", Select: D + " .stat-item.replies", Props: slices.Concat(TYPE, []string{"text-align", "width", "font-variant-numeric"}), Box: true},
			{Name: "row.time", Select: D + " .stat-item.timestamp", Props: slices.Concat(TYPE, []string{"text-align", "width"}), Box: true},
			{Name: "mrow", Select: M, Props: slices.Concat(BOX, FLEX), Box: true, Widths: []int{734, 390}},
			{Name: "mrow.title", Select: M + " a.post-title", Props: TYPE, Widths: []int{734, 390}},
			{Name: "mrow.excerpt", Select: M + " .excerpt", Props: slices.Concat(TYPE, CLAMP, []string{"margin-top"}), Box: true, Widths: []int{734, 390}},
			{Name: "mrow.stats", Select: M + " dl.mobile-stats", Props: slices.Concat(BOX, FLEX), Box: true, Widths: []int{734, 390}},
			{Name: "mrow.stat", Select: M + " dd.mobile-icon.replies", Props: slices.Concat(TYPE, []string{"margin-left", "padding-left"}), Box: true, Widths: []int{734, 390}},
		},
	}
}
