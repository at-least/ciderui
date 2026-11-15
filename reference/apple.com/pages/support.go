package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: support-dark reads the same in the dark, which Apple Support follows.
func support() pagespec.Spec {
	CELL := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "text-align", "vertical-align", "background-color", "border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-style", "border-bottom-color"}
	return pagespec.Spec{
		URL:  "https://support.apple.com/en-us/100100",
		Wait: 3000,
		Probes: []pagespec.Probe{
			{Name: "page", Select: "body", Props: PAGE},
			{Name: "title", Select: "h1.gb-header", Props: slices.Concat(TYPE, []string{"margin-bottom"})},
			{Name: "subheader", Select: "p.gb-subheader", Props: TYPE},
			{Name: "text", Select: "p.gb-paragraph", Props: slices.Concat(TYPE, []string{"margin-top"})},
			{Name: "link", Select: "a.gb-anchor", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "list", Select: "ul.gb-list", Props: []string{"padding-left", "margin-left", "list-style-type"}},
			{Name: "note", Select: ".gb-note", Props: TYPE},
			{Name: "table.wrapper", Select: ".table-wrapper", Props: []string{"overflow-x", "margin-top", "margin-bottom"}, Box: true},
			{Name: "table", Select: "table.gb-table", Props: []string{"border-collapse", "background-color"}, Box: true},
			{Name: "table.head", Select: "table.gb-table th", Props: slices.Concat(TYPE, CELL), Box: true},
			{Name: "table.cell", Select: "table.gb-table td", Props: slices.Concat(TYPE, CELL), Box: true},
			{Name: "table.link", Select: "table.gb-table td a.gb-anchor", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
		},
	}
}
