package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: apple.com's search page for a query nothing matches: its empty state (a
// title and a line under it) and the product tiles it offers instead, a picture
// over its name.
func searchNull() pagespec.Spec {
	return pagespec.Spec{
		URL:  "https://www.apple.com/us/search/xqzvwqk?src=globalnav",
		Wait: 3000,
		Probes: []pagespec.Probe{
			{Name: "null", Select: ".rf-serp-nullsearch", Props: slices.Concat(BOX, []string{"width"}), Box: true},
			{Name: "null.heading", Select: ".rf-serp-nullsearch-heading", Props: slices.Concat(BOX, FLEX, []string{"text-align"}), Box: true},
			{Name: "null.title", Select: ".rf-serp-nullsearch-title", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "text-align"}), Box: true},
			{Name: "null.desc", Select: ".rf-serp-nullsearch-desc", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom", "text-align"}), Box: true},
			{Name: "null.results", Select: ".rf-serp-nullsearch-results", Props: BOX, Box: true},
			{Name: "section.label", Select: ".rf-serp-product-sectionlabel", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "tile", Select: ".rf-serp-product-tile", Props: slices.Concat(BOX, FLEX, []string{"width", "text-align", "overflow"}), Box: true, States: []string{"hover"}},
			{Name: "tile.img", Select: ".rf-serp-product-tile img", Props: []string{"width", "height", "object-fit", "border-top-left-radius", "margin-bottom", "background-color"}, Box: true},
			{Name: "tile.title", Select: ".rf-serp-product-title", Props: slices.Concat(TYPE, []string{"margin-top", "text-align", "-webkit-line-clamp", "overflow"}), Box: true},
			{Name: "tile.link", Select: ".rf-serp-product-tile a, a:has(> .rf-serp-product-tile)", Props: slices.Concat(TYPE, []string{"text-decoration-line"}), States: []string{"hover"}},
		},
	}
}
