package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The global nav's flyouts on a phone (to 833px): the menu opened by its
// button, its Mac row hovered (a pointer at a phone's width: its chevron
// slides in), then clicked: the row's flyout over the menu, a back button at
// the bar's start and the menu's X at its end, the menu's rows sliding out
// and the flyout's in (`animations`), its groups one under another, their
// headings, rows and links; then its back button clicked: the menu's rows
// back. At 390px; the page read with its locale banner above the nav (135px),
// which the opened menu covers.
func flyoutPhone() pagespec.Spec {
	FLY := "#globalnav-submenu-link-mac"
	ELEVATED := FLY + " .globalnav-submenu-group-elevated"
	GROUP := FLY + " .globalnav-submenu-group:nth-child(2)"
	ROW := "#globalnav .globalnav-link-mac"
	BACK := "#globalnav .globalnav-menuback"
	SPACE := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-right", "margin-bottom", "margin-left"}
	SHOWN := []string{"opacity", "visibility", "transform"}

	return pagespec.Spec{
		URL:    "https://www.apple.com/",
		Widths: []int{390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "menu", Click: "#globalnav .globalnav-menutrigger-button", Select: "#globalnav .globalnav-content", Props: []string{"background-color", "height"}, Box: true},
			// @external: a row hovered, and one not: its chevron at its end
			{Name: "row.chevron", Hover: ROW, Hovered: 900, Animations: "#globalnav .globalnav-item-mac", Select: ROW + " .globalnav-link-chevron", Props: slices.Concat([]string{"display", "width", "height"}, SPACE, []string{"color"}, SHOWN, ANIMATION, MOTION), Keyframes: true, Box: true, Position: true},
			{Name: "row.chevron-image", Select: ROW + " .globalnav-link-chevron svg", Props: []string{"width", "height", "transform"}, Box: true, Position: true},
			{Name: "row.text", Select: ROW + " .globalnav-link-text", Props: slices.Concat([]string{"color"}, SHOWN), Box: true, Position: true},
			{Name: "row.other-chevron", Select: "#globalnav .globalnav-link-ipad .globalnav-link-chevron", Props: slices.Concat(SHOWN, ANIMATION), Keyframes: true},
			// @external: the row clicked: its flyout over the menu
			{Name: "flyout", Click: ROW, Animations: "#globalnav", Select: FLY, Props: slices.Concat([]string{"position", "top", "left", "height", "background-color", "overflow-y", "z-index"}, SHOWN, ANIMATION), Keyframes: true, Box: true, Position: true},
			{Name: "flyout.scroll", Select: FLY + " .globalnav-flyout-scroll-container", Props: slices.Concat([]string{"height", "overflow-y"}, SPACE), Box: true, Position: true},
			{Name: "flyout.content", Select: FLY + " .globalnav-flyout-content", Props: slices.Concat(SPACE, []string{"display", "flex-direction"}), Box: true, Position: true},
			{Name: "content", Select: "#globalnav .globalnav-content", Props: []string{"background-color", "overflow-y", "height"}, Box: true},
			{Name: "group.elevated", Select: ELEVATED, Props: slices.Concat(SPACE, []string{"max-width"}), Box: true, Position: true},
			{Name: "group", Select: GROUP, Props: slices.Concat(SPACE, []string{"max-width"}), Box: true, Position: true},
			{Name: "header.elevated", Select: ELEVATED + " .globalnav-submenu-header", Props: []string{"display"}},
			{Name: "header", Select: GROUP + " .globalnav-submenu-header", Props: slices.Concat(TYPE, SPACE, SHOWN), Box: true, Position: true},
			{Name: "item.elevated", Select: ELEVATED + " .globalnav-submenu-list-item-elevated", Props: slices.Concat(TYPE, SPACE, SHOWN, []string{"transition-delay"}), Box: true, Position: true},
			{Name: "item.elevated2", Select: ELEVATED + " .globalnav-submenu-list-item-elevated:nth-child(2)", Props: SPACE, Box: true, Position: true},
			{Name: "item.after", Select: ELEVATED + " .globalnav-submenu-list-item", Props: slices.Concat(TYPE, SPACE), Box: true, Position: true},
			{Name: "item", Select: GROUP + " .globalnav-submenu-list-item", Props: slices.Concat(TYPE, SPACE), Box: true, Position: true},
			{Name: "item2", Select: GROUP + " .globalnav-submenu-list-item:nth-child(2)", Props: SPACE, Box: true, Position: true},
			{Name: "link.elevated", Select: ELEVATED + " .globalnav-submenu-list-item-elevated .globalnav-submenu-link", Props: slices.Concat(TYPE, SPACE, []string{"display", "max-width"}, FOCUS, MOTION), Box: true, Position: true, States: []string{"hover", "focus-visible"}},
			{Name: "link", Select: GROUP + " .globalnav-submenu-link", Props: slices.Concat(TYPE, SPACE, []string{"display", "max-width"}, MOTION), Box: true, Position: true, States: []string{"hover"}},
			// @external: the bar over it: the back button at its start, the menu's X at its end
			{Name: "back", Select: BACK, Props: slices.Concat([]string{"position", "top", "left", "right", "width", "height", "z-index"}, SHOWN, []string{"transition-delay"}, MOTION), Box: true, Position: true},
			{Name: "back.button", Select: BACK + "-button", Props: slices.Concat([]string{"width", "height"}, SPACE, []string{"color", "opacity", "background-color", "cursor"}, FOCUS, MOTION), Box: true, Position: true, States: []string{"hover", "focus-visible"}},
			{Name: "back.image", Select: BACK + "-button svg", Props: []string{"width", "height", "transform"}, Box: true, Position: true},
			{Name: "close", Select: "#globalnav .globalnav-menutrigger-button", Props: []string{"color", "opacity"}, Box: true, Position: true},
			{Name: "row.away", Select: ROW + " .globalnav-link-text", Props: SHOWN, Box: true, Position: true},
			// @external: its back button clicked: the menu's rows back
			{Name: "back.clicked", Click: BACK + "-button", Animations: "#globalnav", Select: FLY, Props: slices.Concat(SHOWN, ANIMATION), Box: true, Position: true},
			{Name: "back.row", Select: ROW + " .globalnav-link-text", Props: SHOWN, Box: true, Position: true},
			{Name: "back.gone", Select: BACK, Props: SHOWN},
			// @external: the slides' keyframes, as the page's stylesheet defines them (an element
			// the probe adds, naming them)
			{Name: "keyframes", Inject: "<div style=\"animation-name: globalnav-flyout-slide-forward-previous, globalnav-flyout-slide-forward-next, globalnav-flyout-slide-back-previous, globalnav-flyout-slide-back-next\"></div>", Into: "#globalnav", Props: []string{"animation-name"}, Keyframes: true},
		},
	}
}
