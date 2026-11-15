package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The global nav's bag, opened (its button clicked once, at the first width
// of each of the nav's layouts, as search.go opens the search): the bag's
// flyout, a link's flyout's own (flyout.go's: globalnav-flyout), its
// markup its script's, fetched (/shop/bag/flyout) and an empty bag's here:
// its message, a line with a sign-in link, the profile's heading and its
// links, each after its icon; the bar while it is open, its button, the
// curtain, each with what the click starts; the profile's links hovered and
// focused (the forced states last, as search.go's). From 834px the pointer
// moves into the flyout's room left of its content, 200px down, where it
// hovers nothing and which the flyout holds before its markup has come
// (0.9s after the click: 174px tall at 0.3s) and at every width (one left
// out of the window as it narrows has left the flyout: leaving it closes
// it); at 834px, last, it leaves for the curtain (left: leaving closes it,
// as the search). To 833px it is a sheet over the window, the menu's
// button its X; the pointer moves into the sheet below its links. The page
// is read with its locale banner above the nav (70px from 834px, 135px to
// 833px), so a box's top is the nav's top plus its own.
func bagFlyout() pagespec.Spec {
	FLY := "#globalnav-submenu-bag"
	LINKED := FLY + " .ac-gn-bagview-nav-link-orders"
	SPACE := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-right", "margin-bottom", "margin-left"}
	ITEM := slices.Concat([]string{"opacity", "transform", "transition-delay"}, MOTION)
	WIDE := []int{1600, 1200, 900, 834}
	PHONE := []int{833, 390}
	MOTIONS := "#globalnav, #globalnav-curtain"
	// @external: the flyout's room left of its content
	ROOM := pagespec.At(100, 200)

	return pagespec.Spec{
		URL:    "https://www.apple.com/",
		Widths: []int{1600, 1200, 900, 834, 833, 390},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "nav", Click: "#globalnav-menubutton-link-bag", Pointer: ROOM, Animations: MOTIONS, Select: "#globalnav", Props: slices.Concat([]string{"height", "position", "background-color"}, BLUR, []string{"z-index"}), Vars: []string{"--globalnav-background", "--globalnav-backdrop-filter"}, Box: true, Position: true, Widths: WIDE},
			{Name: "nav.phone", Click: "#globalnav-menubutton-link-bag", Pointer: pagespec.At(200, 850), Animations: MOTIONS, Select: "#globalnav", Props: slices.Concat([]string{"height", "position", "background-color"}, BLUR, []string{"z-index"}), Vars: []string{"--globalnav-background", "--globalnav-backdrop-filter"}, Box: true, Position: true, Widths: PHONE},
			{Name: "nav.content", Select: "#globalnav .globalnav-content", Props: slices.Concat([]string{"position", "z-index", "height", "background-color"}, BLUR), Box: true, Position: true},
			{Name: "root", Select: "html", Props: []string{"overflow-y"}},
			{Name: "body", Select: "body", Props: []string{"overflow-y"}},
			// @external: the bag's button while it is open, another link, the search beside them
			{Name: "link.bag", Select: "#globalnav-menubutton-link-bag", Props: slices.Concat([]string{"color", "opacity", "visibility"}, MOTION), Box: true, Position: true},
			{Name: "link.other", Select: "#globalnav .globalnav-link-mac", Props: []string{"color", "opacity", "visibility"}},
			{Name: "link.search", Select: "#globalnav .globalnav-link-search", Props: []string{"color", "opacity", "visibility"}},
			// @external: to 833px the menu's button is the sheet's X
			{Name: "menutrigger", Select: "#globalnav-menutrigger-button", Props: slices.Concat([]string{"color", "opacity", "visibility"}, MOTION), Box: true, Position: true, Widths: PHONE},
			{Name: "flyout", Select: FLY, Props: slices.Concat([]string{"position", "top", "left", "right", "height", "max-height"}, SPACE, []string{"background-color"}, BLUR, []string{"color", "overflow-y", "z-index", "visibility"}, MOTION), Vars: []string{"--r-globalnav-flyout-height", "--r-globalnav-flyout-rate"}, Box: true, Position: true},
			{Name: "flyout.scroll", Select: FLY + " .globalnav-flyout-scroll-container", Props: slices.Concat([]string{"height", "max-height", "overflow-y"}, SPACE), Box: true, Position: true},
			{Name: "flyout.content", Select: FLY + " .globalnav-flyout-content", Props: slices.Concat([]string{"max-width"}, SPACE, []string{"box-sizing"}), Box: true, Position: true},
			{Name: "bag", Select: FLY + " .ac-gn-bagview-content", Props: SPACE, Box: true, Position: true},
			{Name: "bag.header", Select: FLY + " .ac-gn-bagview-header", Props: slices.Concat(TYPE, SPACE, ITEM), Box: true, Position: true},
			{Name: "bag.subheader", Select: FLY + " .ac-gn-bagview-subheader", Props: slices.Concat(TYPE, SPACE, ITEM), Box: true, Position: true},
			{Name: "bag.nav", Select: FLY + " .ac-gn-bagview-nav-item-wrapper", Props: SPACE, Box: true, Position: true},
			{Name: "bag.nav-header", Select: FLY + " .ac-gn-bagview-nav-item-header", Props: slices.Concat(TYPE, SPACE, ITEM), Box: true, Position: true},
			{Name: "bag.nav-list", Select: FLY + " .ac-gn-bagview-nav-item-wrapper ul", Props: slices.Concat(SPACE, []string{"list-style-type"}), Box: true, Position: true},
			{Name: "bag.nav-item", Select: FLY + " .ac-gn-bagview-nav-item-orders", Props: slices.Concat(SPACE, []string{"display"}, ITEM), Box: true, Position: true},
			{Name: "bag.nav-item2", Select: FLY + " .ac-gn-bagview-nav-item-yoursaves", Props: ITEM, Box: true, Position: true},
			{Name: "bag.nav-item-last", Select: FLY + " .ac-gn-bagview-nav-item-signIn", Props: ITEM, Box: true, Position: true},
			{Name: "bag.nav-image", Select: LINKED + " .ac-gn-bagview-nav-image-container", Props: slices.Concat([]string{"position", "left", "top", "display"}, SPACE), Box: true, Position: true},
			{Name: "bag.nav-image-regular", Select: LINKED + " .ac-gn-bagview-nav-image-regular svg", Props: []string{"width", "height", "fill", "color"}, Box: true, Position: true, Widths: WIDE},
			{Name: "bag.nav-path", Select: LINKED + " .ac-gn-bagview-nav-image-regular path", Props: []string{"fill"}, Widths: WIDE},
			{Name: "bag.nav-image-compact", Select: LINKED + " .ac-gn-bagview-nav-image-compact svg", Props: []string{"width", "height", "fill", "color"}, Box: true, Position: true, Widths: PHONE},
			{Name: "bag.nav-path-compact", Select: LINKED + " .ac-gn-bagview-nav-image-compact path", Props: []string{"fill"}, Widths: PHONE},
			{Name: "curtain", Select: "#globalnav-curtain", Props: slices.Concat([]string{"position", "top", "z-index", "background-color"}, BLUR, []string{"opacity", "visibility"}, MOTION), Box: true, Position: true},
			// @invariant: the forced states last (search.go's): the sign-in link, a profile link,
			// its text and its icon
			{Name: "bag.subheader-link", Select: FLY + " .ac-gn-bagview-subheader a", Props: slices.Concat(TYPE, LINK, FOCUS), Box: true, Position: true, States: []string{"hover", "focus-visible"}},
			{Name: "bag.nav-link", Select: LINKED, Props: slices.Concat(TYPE, SPACE, []string{"display", "align-items", "position", "text-decoration-line", "background-color", "border-top-left-radius"}, FOCUS, MOTION), Box: true, Position: true, States: []string{"hover", "focus-visible"}},
			// @external: the text hovered itself: apple.com's rule is the link's :hover and the
			// text's (globalheader.css), so the pointer on the text, as a hand's
			// is, turns it #000 where the link's forced alone leaves it
			{Name: "bag.nav-text", Select: LINKED + " .ac-gn-bagview-nav-text", Props: slices.Concat(TYPE, SPACE, []string{"display"}, MOTION), Box: true, Position: true, States: []string{"hover"}},
			{Name: "bag.nav-icon", Select: LINKED + " .ac-gn-bagview-nav-image-regular svg", StateOn: LINKED, Props: slices.Concat([]string{"color", "fill"}, MOTION), States: []string{"hover"}, Widths: WIDE},
			{Name: "bag.nav-icon-compact", Select: LINKED + " .ac-gn-bagview-nav-image-compact svg", StateOn: LINKED, Props: slices.Concat([]string{"color", "fill"}, MOTION), States: []string{"hover"}, Widths: PHONE},
			// @invariant: at 834px, last: the pointer off the flyout, onto the curtain
			{Name: "left", HoverAt: pagespec.At(400, 800), Hovered: 1500, Animations: MOTIONS, Select: FLY, Props: []string{"height", "visibility"}, Box: true, Widths: []int{834}},
			{Name: "left.button", Select: "#globalnav-menubutton-link-bag", Props: []string{"color"}, Widths: []int{834}},
		},
	}
}
