package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: tv.apple.com's video player's playlist: the paused player's Trailers tab
// clicked (tv-player's opening), the row of trailers it shows over the
// video's foot: the tab chosen and not, the row, a trailer's frame, its
// picture, its title and line, the paddle. No trailer is marked as the one
// playing.
func tvPlayerPlaylist() pagespec.Spec {
	ITEM := "#playlist-tab-pane .shelf-grid__list-item:first-child"
	return pagespec.Spec{
		URL:     tvPlayerURL,
		Browser: "chrome",
		Wait:    2000,
		Widths:  []int{1600, 1200},
		Probes: append(tvPlayerOpen(), []pagespec.Probe{
			{Name: "tab", Select: "#playlist-tab-trigger", Props: slices.Concat(TYPE, BOX, []string{"height", "cursor"}, FOCUS), Box: true, States: []string{"hover", "focus-visible"}},
			{Name: "paused", Click: "button.playback-play__pause", Clicked: 1500, Pointer: pagespec.At(800, 200), Select: "button.playback-play__play", Props: []string{"opacity"}},
			{Name: "tab.chosen", Click: "#playlist-tab-trigger", Clicked: 1500, Pointer: pagespec.At(800, 200), Select: "#playlist-tab-trigger", Props: slices.Concat(TYPE, BOX, []string{"height", "cursor", "aria-selected"}), Box: true, Position: true, States: []string{"hover"}},
			{Name: "pane", Select: "#playlist-tab-pane", Props: slices.Concat(BOX, MOTION, []string{"opacity"}), Box: true, Position: true},
			{Name: "row", Select: "#playlist-tab-pane .shelf-grid__list", Props: slices.Concat(FLEX, BOX, []string{"grid-auto-flow", "grid-auto-columns", "grid-template-columns", "overflow-x", "scroll-snap-type", "scroll-padding-left", "padding-left", "padding-right"}), Box: true, Position: true},
			{Name: "item", Select: ITEM, Props: slices.Concat(BOX, []string{"padding-left", "padding-right", "scroll-snap-align"}), Box: true, Position: true},
			{Name: "lockup", Select: ITEM + " .lockup", Props: slices.Concat(BOX, []string{"overflow-x", "cursor", "aspect-ratio"}, FOCUS, MOTION, []string{"transform"}), Vars: []string{"--lockup-aspect-ratio", "--lockup-border-radius"}, Box: true, Position: true, States: []string{"hover", "focus-visible"}},
			{Name: "lockup.after", Select: ITEM + " .lockup", Pseudo: "::after", Props: slices.Concat(BOX, []string{"content", "position", "inset", "opacity"})},
			{Name: "lockup.grid", Select: ITEM + " .lockup .grid", Props: slices.Concat(BOX, []string{"overflow-x"}), Box: true, Position: true},
			{Name: "artwork", Select: ITEM + " .artwork", Props: slices.Concat(BOX, []string{"overflow-x"}), Box: true},
			{Name: "title", Select: ITEM + " .metadata .title", Props: slices.Concat(TYPE, []string{"padding-left", "padding-right"}), Box: true, Position: true},
			{Name: "line", Select: ITEM + " .progress-and-attribution", Props: TYPE, Box: true, Position: true},
			{Name: "gradient", Select: ITEM + " .legibility-gradient", Props: []string{"background-image", "opacity"}, Box: true, Position: true},
			{Name: "paddle", Select: "#playlist-tab-pane .shelf-grid-nav__arrow--right", Props: slices.Concat(BOX, []string{"color", "opacity", "width", "height"}), Box: true, Position: true, States: []string{"hover"}},
		}...),
	}
}
