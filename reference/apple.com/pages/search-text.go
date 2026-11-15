package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The global nav's search with text in it, opened as search.go opens it: the
// field given a value and its two buttons what apple.com's script gives them
// while it holds text (its _updateSearchFieldButtons: tabindex 0, no
// disabled, no aria-hidden), set by the probe (`set`), so nothing is typed
// and no suggestion is asked for. The input's text; the clear button shown
// (its place, size, fill and motion, its images), the search button's fill;
// both hovered and focused, last. Read only.
func searchText() pagespec.Spec {
	SPACE := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-right", "margin-bottom", "margin-left"}
	FLY := "#globalnav-submenu-search"
	WIDE := []int{1600, 1200, 900, 834}
	PHONE := []int{833, 390}
	INPUT := FLY + " .globalnav-searchfield-input"
	RESET := FLY + " .globalnav-searchfield-reset"
	SUBMIT := FLY + " .globalnav-searchfield-submit"
	FILLED := []pagespec.Set{pagespec.Attr(INPUT, "value", "iPhone")}
	for _, b := range []string{RESET, SUBMIT} {
		FILLED = append(FILLED, pagespec.Attr(b, "tabindex", "0"), pagespec.NoAttr(b, "disabled"), pagespec.NoAttr(b, "aria-hidden"))
	}
	BUTTON := slices.Concat(SPACE, []string{"width", "height", "position", "display", "align-items", "justify-content", "fill", "color", "background-color", "border-top-width", "cursor", "opacity", "visibility", "z-index", "order"}, MOTION)

	s := search()
	s.Probes = append(pick(s.Probes, "nav", "nav.phone"), []pagespec.Probe{
		{Name: "input", Set: FILLED, Select: INPUT, Props: slices.Concat(TYPE, SPACE, []string{"width", "height", "caret-color"}), Box: true, Position: true},
		{Name: "reset", Select: RESET, Props: BUTTON, Box: true, Position: true},
		{Name: "reset.image", Select: RESET + " .globalnav-image-regular svg", Props: []string{"width", "height", "fill", "display"}, Box: true, Position: true, Widths: WIDE},
		{Name: "reset.image-compact", Select: RESET + " .globalnav-image-compact svg", Props: []string{"width", "height", "fill", "display"}, Box: true, Position: true, Widths: PHONE},
		{Name: "submit", Select: SUBMIT, Props: BUTTON, Box: true, Position: true},
		// @invariant: hovered and focused last: the still style that holds a state cancels
		// the page's animations
		{Name: "reset.states", Select: RESET, Props: slices.Concat([]string{"fill"}, FOCUS), States: []string{"hover", "focus-visible"}},
		{Name: "submit.states", Select: SUBMIT, Props: slices.Concat([]string{"fill"}, FOCUS), States: []string{"hover", "focus-visible"}},
	}...)
	return s
}
