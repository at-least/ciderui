package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

func localnavProbes() []pagespec.Probe {
	S := 1500
	return []pagespec.Probe{
		{Name: "localnav", Select: "#ac-localnav", Scroll: S, Props: []string{"position", "top", "z-index", "min-width"}, Box: true},
		{Name: "localnav.top", Select: "#ac-localnav .ac-ln-background", Props: slices.Concat([]string{"background-color"}, BLUR)},
		{Name: "localnav.background", Select: "#ac-localnav .ac-ln-background", Scroll: S, Props: slices.Concat([]string{"background-color"}, BLUR, MOTION)},
		{Name: "localnav.rule", Select: "#ac-localnav .ac-ln-background", Pseudo: "::after", Scroll: S, Props: []string{"background-color", "height", "position", "bottom", "z-index"}},
		{Name: "localnav.content", Select: "#ac-localnav .ac-ln-content", Scroll: S, Props: []string{"max-width", "padding-left", "padding-right", "margin-left", "margin-right", "height"}, Box: true},
		{Name: "localnav.title", Select: "#ac-localnav .ac-ln-title a", Scroll: S, Props: slices.Concat(TYPE, []string{"opacity", "text-decoration-line", "cursor"}, MOTION), Box: true, Position: true, States: []string{"hover"}},
		{Name: "localnav.item", Select: "#ac-localnav .ac-ln-menu-item:nth-child(2)", Scroll: S, Props: []string{"margin-left", "margin-right", "display"}},
		{Name: "localnav.link", Select: "#ac-localnav a.ac-ln-menu-link", Scroll: S, Props: slices.Concat(TYPE, []string{"opacity", "text-decoration-line", "display", "cursor"}, MOTION), Box: true, Position: true, States: []string{"hover"}},
		{Name: "localnav.current", Select: "#ac-localnav .ac-ln-menu-link.current", Scroll: S, Props: slices.Concat(TYPE, []string{"opacity", "cursor", "display"}), Box: true, Position: true},
		{Name: "localnav.current-mark", Select: "#ac-localnav .ac-ln-menu-link.current", Pseudo: "::after", Scroll: S, Props: []string{"background-color", "opacity", "width", "height", "top", "left", "bottom", "border-top-left-radius", "position"}},
		{Name: "localnav.action", Select: "#ac-localnav .ac-ln-action-button", Scroll: S, Props: []string{"margin-left", "margin-top"}},
		{Name: "localnav.button", Select: "#ac-localnav .ac-ln-button", Scroll: S, Props: slices.Concat(TYPE, BOX, []string{"margin-top", "cursor"}), Box: true, Position: true, States: []string{"hover", "active"}},
		{Name: "localnav.chevron", Select: "#ac-localnav .ac-ln-menucta", Scroll: S, Props: []string{"width", "height", "margin-right", "display", "cursor"}, Box: true, Position: true, Widths: []int{833, 390}},
		{Name: "localnav.chevron-image", Select: "#ac-localnav .ac-ln-menucta-chevron", Scroll: S, Props: slices.Concat([]string{"color", "opacity"}, MOTION, FLEX), Box: true, States: []string{"hover"}, Widths: []int{833, 390}},
		{Name: "localnav.chevron-svg", Select: "#ac-localnav .ac-ln-menucta-chevron svg", Scroll: S, Props: []string{"width", "height", "transform"}, Box: true, Position: true, Widths: []int{833, 390}},
		{Name: "localnav.tray", Select: "#ac-localnav .ac-ln-menu-tray", Scroll: S, Props: slices.Concat([]string{"visibility", "max-height", "overflow-y"}, MOTION), Box: true, Widths: []int{833, 390}},
		// @invariant: its menu opened by the chevron: last, as the page stays so
		{Name: "tray.open", Click: "#ac-localnav .ac-ln-menucta", Animations: "#ac-localnav", Select: "#ac-localnav .ac-ln-menu-tray", Scroll: S, Props: slices.Concat([]string{"visibility", "max-height"}, MOTION), Box: true, Position: true, Widths: []int{390}},
		{Name: "tray.background", Select: "#ac-localnav .ac-ln-background", Scroll: S, Props: slices.Concat([]string{"background-color"}, BLUR, MOTION), Box: true, Widths: []int{390}},
		{Name: "tray.rule", Select: "#ac-localnav .ac-ln-background", Pseudo: "::after", Scroll: S, Props: []string{"background-color", "height", "top"}, Widths: []int{390}},
		{Name: "tray.items", Select: "#ac-localnav .ac-ln-menu-items", Scroll: S, Props: slices.Concat(BOX, []string{"padding-left", "padding-right"}), Box: true, Position: true, Widths: []int{390}},
		{Name: "tray.link", Select: "#ac-localnav a.ac-ln-menu-link", Scroll: S, Props: slices.Concat(TYPE, []string{"opacity", "display", "align-items", "transform"}, MOTION), Box: true, Position: true, Widths: []int{390}, States: []string{"hover"}},
		{Name: "tray.current", Select: "#ac-localnav .ac-ln-menu-link.current", Scroll: S, Props: slices.Concat(TYPE, []string{"opacity"}), Box: true, Position: true, Widths: []int{390}},
		{Name: "tray.current-mark", Select: "#ac-localnav .ac-ln-menu-link.current", Pseudo: "::after", Scroll: S, Props: []string{"background-color", "opacity", "width", "height", "top", "left", "border-top-left-radius"}, Widths: []int{390}},
		{Name: "tray.chevron-svg", Select: "#ac-localnav .ac-ln-menucta-chevron svg", Scroll: S, Props: []string{"transform"}, Widths: []int{390}},
	}
}

// @external: apple.com's local nav, the one most product pages carry under the global
// nav (MacBook Air's, light): the bar, sticky, its material and its rule,
// the product's name, its links (the current page's marked), its Buy
// button, and to 833px its menu, opened by its chevron. Everything is read
// with the page scrolled (the bar stuck to the top), but `*.top`, at the top;
// the menu is the bar to 833px, from 834px (the global nav's step).
func localnav() pagespec.Spec {
	return pagespec.Spec{URL: "https://www.apple.com/macbook-air/", Widths: []int{1600, 1200, 900, 834, 833, 390}, Probes: localnavProbes()}
}
