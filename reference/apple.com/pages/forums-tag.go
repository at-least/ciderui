package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: Apple Developer Forums' page for one tag (`/forums/tags/swiftui`), signed
// out: its header (the tag's name, its line, its links) and the count over its list.
func forumsTag() pagespec.Spec {
	return pagespec.Spec{
		URL:    "https://developer.apple.com/forums/tags/swiftui",
		Wait:   3000,
		Widths: []int{1600, 1200, 900, 735, 734, 390},
		Probes: []pagespec.Probe{
			{Name: "header", Select: "section.tag-overview", Props: slices.Concat(BOX, []string{"width"}), Box: true},
			{Name: "header.content", Select: "section.tag-overview .tag-content", Props: slices.Concat(BOX, FLEX), Box: true},
			{Name: "header.title", Select: "section.tag-overview .tag-title-and-rss h1, section.tag-overview .tag-title-and-rss h2", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "header.info", Select: "section.tag-overview p.tag-info", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "max-width"}), Box: true},
			{Name: "header.links", Select: "section.tag-overview p.tag-links", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "header.link", Select: "section.tag-overview p.tag-links a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "count.bar", Select: "main section div.header", Props: slices.Concat(BOX, FLEX), Box: true},
			{Name: "count", Select: "main section div.header span.search-results", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true},
		},
	}
}
