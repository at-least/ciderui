package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: apple.com's Newsroom, its index: the latest stories, a grid of white
// rounded tiles on the alternate ground (one across, two, three), each a
// link with its picture, its category, its headline and its date. The
// event banner above (#takeover) and the modules a script draws below are
// not read. The page's script gives the tiles their classes by the window
// (from 735px to 1068px every tile is tile-2up, two across), so a tile is
// picked by its place in the grid: the first (one across from 1069px and
// below 735px), the third (two across), the fourth (three across).
func newsroomIndex() pagespec.Spec {
	MEDIA := []string{"aspect-ratio", "overflow", "border-top-left-radius", "border-bottom-left-radius"}

	return pagespec.Spec{
		URL:    "https://www.apple.com/newsroom/",
		Widths: []int{1600, 1200, 1069, 1068, 900, 735, 734, 390},
		Probes: []pagespec.Probe{
			{Name: "feed", Select: ".everydayfeed", Props: slices.Concat(BOX, []string{"padding-top", "padding-bottom"})},
			{Name: "feed.content", Select: ".everydayfeed .section-content", Props: []string{"width", "max-width", "margin-left", "margin-right", "padding-left", "padding-right"}, Box: true},
			{Name: "feed.head", Select: ".everydayfeed .section-head", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "tiles", Select: ".everydayfeed .section-tiles", Props: slices.Concat(FLEX, []string{"grid-template-columns", "list-style-type", "padding-left", "margin-top", "margin-bottom"}), Box: true},
			{Name: "tiles.item", Select: ".everydayfeed .section-tiles > li", Props: []string{"grid-column-start", "grid-column-end", "display", "margin-top", "margin-bottom", "flex-basis", "width"}, Box: true},
			// @external: the first tile: one across, its picture beside its text
			{Name: "tile.first", Select: ".everydayfeed li:nth-child(1) > a", Props: slices.Concat(BOX, TEXT, []string{"display", "grid-template-columns", "column-gap", "transform", "transition-property", "transition-duration", "overflow"}), Box: true, States: []string{"hover", "focus-visible"}},
			{Name: "tile.first-media", Select: ".everydayfeed li:nth-child(1) > a .tile__media", Props: MEDIA, Box: true},
			{Name: "tile.first-description", Select: ".everydayfeed li:nth-child(1) > a .tile__description", Props: slices.Concat(BOX, []string{"padding-left", "padding-right", "display", "flex-direction", "justify-content"}), Box: true},
			{Name: "tile.first-eyebrow", Select: ".everydayfeed li:nth-child(1) > a :is(.tile__hero--eyebrow, .tile__category)", Props: slices.Concat(TYPE, TEXT, []string{"text-transform"})},
			{Name: "tile.first-headline", Select: ".everydayfeed li:nth-child(1) > a .tile__headline", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "tile.first-date", Select: ".everydayfeed li:nth-child(1) > a .tile__timestamp", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			// @external: the third: two across, its picture over its text
			{Name: "tile.third", Select: ".everydayfeed li:nth-child(3) > a", Props: slices.Concat(BOX, []string{"display", "flex-direction", "transform", "overflow"}), Box: true, States: []string{"hover"}},
			{Name: "tile.third-media", Select: ".everydayfeed li:nth-child(3) > a .tile__media", Props: MEDIA, Box: true},
			{Name: "tile.third-description", Select: ".everydayfeed li:nth-child(3) > a .tile__description", Props: slices.Concat(BOX, []string{"padding-left", "padding-right", "display", "flex-direction", "justify-content", "flex-grow"}), Box: true},
			{Name: "tile.third-category", Select: ".everydayfeed li:nth-child(3) > a .tile__category", Props: slices.Concat(TYPE, TEXT, []string{"text-transform"})},
			{Name: "tile.third-headline", Select: ".everydayfeed li:nth-child(3) > a .tile__headline", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "tile.third-date", Select: ".everydayfeed li:nth-child(3) > a .tile__timestamp", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			// @external: the fourth: three across
			{Name: "tile.fourth", Select: ".everydayfeed li:nth-child(4) > a", Props: slices.Concat(BOX, []string{"display", "flex-direction", "transform", "overflow"}), Box: true, States: []string{"hover"}},
			{Name: "tile.fourth-media", Select: ".everydayfeed li:nth-child(4) > a .tile__media", Props: MEDIA, Box: true},
			{Name: "tile.fourth-description", Select: ".everydayfeed li:nth-child(4) > a .tile__description", Props: slices.Concat(BOX, []string{"padding-left", "padding-right", "display", "flex-direction", "justify-content", "flex-grow"}), Box: true},
			{Name: "tile.fourth-category", Select: ".everydayfeed li:nth-child(4) > a .tile__category", Props: slices.Concat(TYPE, TEXT, []string{"text-transform"})},
			{Name: "tile.fourth-headline", Select: ".everydayfeed li:nth-child(4) > a .tile__headline", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "tile.fourth-date", Select: ".everydayfeed li:nth-child(4) > a .tile__timestamp", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
			{Name: "tile.link", Select: ".everydayfeed li:nth-child(4) > a", Props: LINK},
		},
	}
}
