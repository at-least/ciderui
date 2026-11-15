package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: apple.com's search page with results (`airpods`): its field (a magnifier
// at its start, a reset at its end), the count over the results, and a result's
// tile, a picture over its name and price, in a grid.
func searchSerp() pagespec.Spec {
	return pagespec.Spec{
		URL:  "https://www.apple.com/us/search/airpods?src=globalnav",
		Wait: 3000,
		Probes: []pagespec.Probe{
			{Name: "search", Select: ".rf-serp-search", Props: BOX, Box: true},
			{Name: "box", Select: ".rf-serp-searchbox-wrapper", Props: slices.Concat(BOX, []string{"width"}), Box: true},
			{Name: "field", Select: ".rf-serp-searchbox .form-textbox", Props: slices.Concat(BOX, []string{"width", "position"}), Box: true},
			{Name: "input", Select: ".rf-serp-searchbox input.form-textbox-input", Props: slices.Concat(TYPE, BOX, []string{"height", "border-left-width", "border-left-color", "padding-left", "padding-right", "border-top-right-radius"}), Box: true, States: []string{"hover", "focus"}},
			{Name: "label", Select: ".rf-serp-searchbox label.form-textbox-label", Props: slices.Concat(TYPE, []string{"top", "left", "position"}), Box: true},
			{Name: "icon.left", Select: ".rf-serp-searchbox .form-textbox-lefticon", Props: []string{"position", "left", "top", "width", "height", "color"}, Box: true},
			{Name: "icon.left.button", Select: ".rf-serp-searchbox .form-textbox-lefticon button", Props: []string{"width", "height", "color", "font-size", "line-height"}, Box: true},
			{Name: "icon.right", Select: ".rf-serp-searchbox .form-textbox-righticon", Props: []string{"position", "right", "top", "width", "height", "visibility", "opacity"}, Box: true},
			{Name: "count", Select: ".rf-serp-resultcount", Props: slices.Concat(TYPE, BOX), Box: true},
			{Name: "grid", Select: ".rf-accessoriesgrid", Props: slices.Concat(FLEX, []string{"grid-template-columns", "width"}), Box: true},
			{Name: "tile", Select: ".rf-producttile", Props: slices.Concat(BOX, []string{"width", "flex-basis"}), Box: true},
			{Name: "tile.gallery", Select: ".rf-producttile-gallery", Props: slices.Concat(BOX, []string{"width", "height", "overflow"}), Box: true},
			{Name: "tile.img", Select: ".rf-producttile-gallery img", Props: []string{"width", "height", "object-fit", "border-top-left-radius"}, Box: true},
			{Name: "tile.info", Select: ".rf-producttile-info", Props: slices.Concat(BOX, []string{"text-align"}), Box: true},
			{Name: "tile.name", Select: ".rf-producttile-name", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "tile.name.link", Select: ".rf-producttile-name a", Props: slices.Concat(TYPE, LINK), States: []string{"hover", "focus-visible"}},
			{Name: "tile.price", Select: ".rf-producttile-priceinfo", Props: slices.Concat(TYPE, []string{"margin-top"}), Box: true},
		},
	}
}
