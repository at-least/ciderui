package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: tv.apple.com's video player's subtitles menu, opened by its button over
// the paused player (tv-player's opening): the menu, a group and its title,
// an item hovered and the chosen one with its check, a rule between groups,
// an item with a line under its name and a chevron (its language). At 1600px
// alone: a resize closes the menu.
func tvPlayerMenu() pagespec.Spec {
	MENU := "amp-contextual-menu .contextual-menu--subtitles"
	ITEM := slices.Concat(TYPE, BOX, FLEX, []string{"height", "cursor"})
	return pagespec.Spec{
		URL:     tvPlayerURL,
		Browser: "chrome",
		Wait:    2000,
		Widths:  []int{1600},
		Probes: append(tvPlayerOpen(), []pagespec.Probe{
			{Name: "paused", Click: "button.playback-play__pause", Clicked: 1500, Pointer: pagespec.At(800, 200), Select: "button.playback-play__play", Props: []string{"opacity"}},
			{Name: "menu", Click: "button[aria-label=\"Subtitles\"]", Clicked: 1000, Pointer: pagespec.At(800, 200), Select: MENU, Props: slices.Concat(BOX, BLUR, []string{"overflow-x", "overflow-y", "width", "z-index"}), Box: true, Position: true},
			{Name: "menu.before", Select: MENU, Pseudo: "::before", Props: slices.Concat(BOX, BLUR, []string{"content", "position", "inset"})},
			{Name: "menu.overlay", Select: "amp-contextual-menu .contextual-menu__overlay", Props: slices.Concat(BOX, BLUR, []string{"position"}), Box: true},
			{Name: "menu.list", Select: MENU + " .contextual-menu__list", Props: slices.Concat(BOX, BLUR, []string{"list-style-type"}), Box: true},
			{Name: "group", Select: MENU + " .contextual-menu__group", Props: BOX, Box: true},
			{Name: "group.next", Select: MENU + " .contextual-menu__group + .contextual-menu__group", Props: BOX, Box: true},
			{Name: "group.title", Select: MENU + " .contextual-menu__group-title", Props: slices.Concat(TYPE, BOX, []string{"text-transform"}), Box: true},
			{Name: "item", Select: MENU + " .contextual-menu-item button[title=\"On\"]", Props: ITEM, Box: true, States: []string{"hover", "focus-visible", "active"}},
			{Name: "item.row", Select: MENU + " .contextual-menu-item:has(button[title=\"On\"])", Props: BOX, Box: true, States: []string{"hover"}},
			{Name: "item.text", Select: MENU + " button[title=\"On\"] .contextual-menu-item__option-text", Props: TYPE, Box: true, Position: true},
			{Name: "chosen", Select: MENU + " .contextual-menu-item button[title=\"Off\"]", Props: ITEM, Box: true},
			{Name: "chosen.check", Select: MENU + " button[title=\"Off\"] .contextual-menu-item__icon-container", Props: []string{"color", "width", "height"}, Box: true, Position: true},
			{Name: "sub", Select: MENU + " .contextual-menu-item--use-html-select label", Props: ITEM, Box: true},
			{Name: "sub.subtext", Select: MENU + " .contextual-menu-item__subtext", Props: TYPE, Box: true},
			{Name: "sub.chevron", Select: MENU + " .contextual-menu-item--use-html-select .contextual-menu-item__icon-container", Props: []string{"color", "width", "height"}, Box: true, Position: true},
		}...),
	}
}
