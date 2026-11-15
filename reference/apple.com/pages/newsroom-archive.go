package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: apple.com's Newsroom, its archive: each month a heading over ruled rows,
// each row a link with its picture, its category, its headline and its
// date. The filter bar above (a script's) and the button that loads more
// are not read.
func newsroomArchive() pagespec.Spec {
	return pagespec.Spec{
		URL:    "https://www.apple.com/newsroom/archive/",
		Widths: []int{1600, 1200, 1069, 1068, 900, 735, 734, 390},
		Probes: []pagespec.Probe{
			{Name: "results", Select: ".results__content", Props: []string{"width", "max-width", "margin-left", "margin-right", "padding-left", "padding-right", "padding-top", "padding-bottom"}, Box: true},
			{Name: "month", Select: ".result", Props: []string{"margin-top", "margin-bottom", "padding-top", "padding-bottom"}},
			{Name: "month.title", Select: ".result__title", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "row", Select: ".result__item", Props: slices.Concat(BOX, FLEX, LINK, []string{"border-bottom-width", "border-bottom-style", "border-bottom-color", "cursor"}), Box: true, States: []string{"hover", "focus-visible"}},
			{Name: "row.last", Select: ".result .result__item:last-child", Props: []string{"border-bottom-width", "border-bottom-style", "border-bottom-color", "padding-bottom"}},
			{Name: "row.img", Select: ".result__item picture", Props: []string{"width", "height", "aspect-ratio", "border-top-left-radius", "overflow", "margin-right", "flex-shrink"}, Box: true},
			{Name: "row.copy", Select: ".result__item .item__content", Props: []string{"padding-left", "margin-left", "flex-grow", "flex-basis"}, Box: true},
			{Name: "row.category", Select: ".result__item .item__category", Props: slices.Concat(TYPE, TEXT, []string{"text-transform"})},
			{Name: "row.headline", Select: ".result__item .item__headline", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), States: []string{"hover"}},
			{Name: "row.date", Select: ".result__item .item__date", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
		},
	}
}
