package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The global nav's search, opened (its button clicked once, at the first
// width of each of the nav's layouts: it stays open as the window narrows
// within one, and closes where the nav changes, at 833px), so its probes are
// a spec of their own: the home page's flyout, its field (the input, its
// placeholder, the search and clear buttons), its quick links (a heading, a
// list of links each after an arrow), the curtain over the page, and the
// bar while it is open. From 834px the flyout drops under the bar; to 833px it
// is a sheet over the bar and the rest of the window. Once it has opened the
// pointer moves into it, below its links (200, 430; 200, 850 on the sheet),
// where it hovers none: leaving it for the curtain closes it.
// The page is read with its locale banner above the nav (#ac-localeswitcher,
// for a reader outside the U.S.: 70px tall from 834px, 135px to 833px), so a
// box's top is the nav's top plus its own.
func search() pagespec.Spec {
	SPACE := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-right", "margin-bottom", "margin-left"}
	FLY := "#globalnav-submenu-search"
	WIDE := []int{1600, 1200, 900, 834}
	PHONE := []int{833, 390}
	NAV := slices.Concat([]string{"height", "position", "background-color"}, BLUR)
	NAV_VARS := []string{"--globalnav-background", "--globalnav-backdrop-filter"}

	return pagespec.Spec{
		URL:    "https://www.apple.com/",
		Widths: []int{1600, 1200, 900, 834, 833, 390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "nav", Click: "#globalnav-menubutton-link-search", Pointer: pagespec.At(200, 430), Animations: "#globalnav", Select: "#globalnav", Props: NAV, Vars: NAV_VARS, Box: true, Position: true, Widths: WIDE},
			{Name: "nav.phone", Click: "#globalnav-menubutton-link-search", Pointer: pagespec.At(200, 850), Animations: "#globalnav", Select: "#globalnav", Props: NAV, Vars: NAV_VARS, Box: true, Position: true, Widths: PHONE},
			// @external: the bar's content box, the sheet's ground to 833px; the page held
			{Name: "content", Select: "#globalnav .globalnav-content", Props: slices.Concat([]string{"position", "top", "height", "background-color"}, BLUR, []string{"z-index"}, MOTION), Box: true, Position: true},
			{Name: "root", Select: "html", Props: []string{"overflow-y"}},
			{Name: "body", Select: "body", Props: []string{"overflow-y"}},
			{Name: "search", Select: "#globalnav-menubutton-link-search", Props: slices.Concat([]string{"opacity", "visibility", "color"}, MOTION), Box: true, Position: true},
			{Name: "flyout", Select: FLY, Props: slices.Concat([]string{"position", "top", "left", "right", "height", "max-height", "background-color", "color", "overflow-y", "z-index"}, MOTION), Box: true, Position: true},
			{Name: "flyout.scroll", Select: FLY + " .globalnav-flyout-scroll-container", Props: slices.Concat([]string{"height", "max-height", "overflow-y"}, SPACE), Box: true, Position: true},
			{Name: "flyout.content", Select: FLY + " .globalnav-flyout-content", Props: slices.Concat([]string{"max-width"}, SPACE), Box: true, Position: true},
			{Name: "field", Select: FLY + " .globalnav-searchfield", Props: slices.Concat(SPACE, []string{"height"}), Box: true, Position: true},
			{Name: "field.wrapper", Select: FLY + " .globalnav-searchfield-wrapper", Props: slices.Concat(ANIMATION, []string{"opacity", "transform"}), Keyframes: true},
			{Name: "field.input", Select: FLY + " .globalnav-searchfield-input", Props: slices.Concat(TYPE, SPACE, []string{"height", "background-color", "border-top-width", "border-top-style", "border-top-left-radius", "outline-style", "caret-color", "text-overflow", "-webkit-appearance", "appearance"}), Box: true, Position: true},
			{Name: "field.placeholder", Select: FLY + " .globalnav-searchfield-input", Pseudo: "::placeholder", Props: []string{"color", "font-weight", "opacity"}},
			{Name: "field.submit", Select: FLY + " .globalnav-searchfield-submit", Props: slices.Concat(SPACE, []string{"width", "height", "position", "left", "top", "color", "fill", "background-color", "border-top-width", "cursor"}, MOTION), Box: true, Position: true},
			{Name: "field.submit-image", Select: FLY + " .globalnav-searchfield-submit .globalnav-image-regular svg", Props: []string{"width", "height", "fill"}, Box: true, Position: true, Widths: WIDE},
			{Name: "field.submit-compact", Select: FLY + " .globalnav-searchfield-submit .globalnav-image-compact svg", Props: []string{"width", "height", "fill"}, Box: true, Position: true, Widths: PHONE},
			{Name: "field.reset", Select: FLY + " .globalnav-searchfield-reset", Props: slices.Concat([]string{"opacity", "visibility", "width", "height"}, MOTION), Box: true, Position: true},
			{Name: "results", Select: FLY + " .globalnav-searchresults", Props: SPACE, Box: true, Position: true},
			{Name: "results.header", Select: FLY + " .globalnav-searchresults-header", Props: slices.Concat(TYPE, SPACE, ANIMATION, []string{"opacity", "transform"}), Keyframes: true, Box: true, Position: true},
			{Name: "results.list", Select: FLY + " .globalnav-searchresults-list", Props: slices.Concat(SPACE, []string{"list-style-type"}), Box: true, Position: true},
			{Name: "results.item", Select: FLY + " .globalnav-searchresults-list-item:first-child", Props: slices.Concat(SPACE, []string{"display"}), Box: true, Position: true},
			{Name: "results.item2", Select: FLY + " .globalnav-searchresults-list-item:nth-child(2)", Props: SPACE, Box: true, Position: true},
			{Name: "results.icon", Select: FLY + " .globalnav-searchresults-list-icon", Props: slices.Concat(SPACE, []string{"display", "color"}), Box: true, Position: true},
			{Name: "results.icon-compact", Select: FLY + " .globalnav-searchresults-list-icon .globalnav-image-compact svg", Props: []string{"width", "height", "fill"}, Box: true, Position: true, Widths: PHONE},
			{Name: "results.text", Select: FLY + " .globalnav-searchresults-list-text", Props: slices.Concat(TYPE, SPACE), Box: true, Position: true},
			{Name: "curtain", Select: "#globalnav-curtain", Props: slices.Concat([]string{"position", "top", "z-index", "background-color"}, BLUR, []string{"opacity", "visibility"}, MOTION), Box: true, Position: true},
			// @invariant: the forced states last: the still style that holds them cancels the
			// page's animations, and taking it off starts them again, from the rows'
			// first frame (4px up and transparent), which a probe read after them
			// reads (the link's own rest is read before its states)
			{Name: "results.link", Select: FLY + " .globalnav-searchresults-list-link", Props: slices.Concat(TYPE, SPACE, []string{"display", "align-items", "text-decoration-line", "background-color", "border-top-left-radius"}, FOCUS, MOTION), Box: true, Position: true, States: []string{"hover", "focus-visible"}},
			{Name: "field.focus", Select: FLY + " .globalnav-searchfield-input", Props: FOCUS, States: []string{"focus-visible", "focus"}},
			{Name: "results.icon-image", Select: FLY + " .globalnav-searchresults-list-icon .globalnav-image-regular svg", StateOn: FLY + " .globalnav-searchresults-list-link", Props: slices.Concat([]string{"width", "height", "fill"}, MOTION), Box: true, Position: true, Widths: WIDE, States: []string{"hover"}},
		},
	}
}
