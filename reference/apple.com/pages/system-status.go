package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

var systemStatusMARK = []string{"width", "height", "margin-right", "position", "top", "left", "vertical-align", "display", "background-size", "background-repeat"}

func systemStatus() pagespec.Spec {
	WIDE := []int{1600, 1200}
	NARROW := []int{900, 390}
	return pagespec.Spec{
		URL:  "https://developer.apple.com/system-status/",
		Wait: 3000,
		Probes: []pagespec.Probe{
			{Name: "page", Select: "body", Props: PAGE},
			{Name: "title", Select: "h1.section-headline", Props: slices.Concat(TYPE, []string{"margin-bottom"}), Box: true},
			{Name: "legend", Select: "section.section-legend", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Widths: WIDE},
			{Name: "legend.item", Select: ".section-legend li", Props: []string{"display", "margin-right", "padding-left", "padding-right"}, Box: true, Widths: WIDE},
			{Name: "legend.mark", Select: ".section-legend li > div", Props: systemStatusMARK, Image: []string{"background-image"}, Box: true, Widths: WIDE},
			{Name: "legend.word", Select: ".section-legend li > span:not(.a11y)", Props: TYPE, Widths: WIDE},
			{Name: "table", Select: ".section-lights table", Props: []string{"border-collapse", "border-spacing", "margin-top", "margin-bottom"}, Box: true, Widths: WIDE},
			{Name: "row", Select: ".section-lights table tr", Props: []string{"background-color", "border-bottom-width", "border-bottom-style", "border-bottom-color"}, Box: true, Widths: WIDE},
			{Name: "cell", Select: ".section-lights td", Props: slices.Concat(TYPE, BOX, []string{"vertical-align", "text-align"}), Box: true, Widths: WIDE},
			{Name: "cell.mark", Select: ".section-lights td .light-available", Props: systemStatusMARK, Image: []string{"background-image"}, Box: true, Widths: WIDE},
			{Name: "cell.name", Select: ".section-lights td .light-name > span:not(.a11y)", Props: TYPE, Widths: WIDE},
			{Name: "cell.link", Select: ".section-lights td a.light-url-icon", Props: []string{"width", "height", "display", "margin-left", "background-size"}, Image: []string{"background-image"}, Box: true, Widths: WIDE},
			{Name: "links", Select: ".section-links p.link-copy", Props: slices.Concat(TYPE, []string{"margin-top"})},
			{Name: "links.link", Select: ".section-links p.link-copy a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
			{Name: "date", Select: ".section-date p.date-copy", Props: slices.Concat(TYPE, []string{"margin-top"})},
			{Name: "all", Select: ".lights-small .light-container", Props: slices.Concat(BOX, []string{"display", "width"}), Box: true, Widths: NARROW},
			{Name: "all.mark", Select: ".lights-small .light-available", Props: systemStatusMARK, Image: []string{"background-image"}, Box: true, Widths: NARROW},
			{Name: "all.words", Select: ".lights-small h2", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true, Widths: NARROW},
			{Name: "all.button", Select: ".lights-small h2 button", Props: slices.Concat(TYPE, BOX, []string{"cursor", "text-align"}), Box: true, Widths: NARROW},
			{Name: "all.toggle", Select: ".lights-small .light-toggle-icon", Props: systemStatusMARK, Image: []string{"background-image"}, Box: true, Widths: NARROW},
		},
	}
}
