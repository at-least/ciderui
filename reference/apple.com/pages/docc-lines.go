package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

func doccLinesProbes() []pagespec.Probe {
	LINE := []string{"display", "width", "box-sizing", "padding-left", "padding-right", "background-color", "border-left-width", "border-left-style", "border-left-color"}
	NUMBER := slices.Concat(TYPE, []string{"display", "min-width", "text-align", "user-select", "padding-top", "padding-right", "padding-bottom", "padding-left", "border-left-width", "border-left-style", "border-left-color"})
	return []pagespec.Probe{
		{Name: "listing", Select: ".code-listing:has(.code-number)", Props: slices.Concat(BOX, []string{"display", "overflow-x"}), Box: true},
		{Name: "listing.pre", Select: ".code-listing:has(.code-number) pre", Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "white-space"}, Box: true},
		{Name: "listing.code", Select: ".code-listing:has(.code-number) pre code", Props: TYPE},
		{Name: "line", Select: ".code-listing:has(.code-number) .code-line-container:not(.highlighted)", Props: LINE, Box: true},
		{Name: "line.number", Select: ".code-listing:has(.code-number) .code-line-container:not(.highlighted) .code-number", Props: NUMBER, Box: true},
		{Name: "highlighted", Select: ".code-listing:has(.code-number) .code-line-container.highlighted", Props: LINE, Box: true},
		{Name: "highlighted.number", Select: ".code-listing:has(.code-number) .code-line-container.highlighted .code-number", Props: NUMBER, Box: true},
	}
}

func doccLines() pagespec.Spec {
	return pagespec.Spec{
		URL:    "https://developer.apple.com/tutorials/swiftui/creating-and-combining-views",
		Wait:   6000,
		Probes: doccLinesProbes(),
	}
}
