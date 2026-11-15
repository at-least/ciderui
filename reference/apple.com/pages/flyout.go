package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The global nav's flyouts, a desktop's menus under its links: Mac's,
// opened by the pointer (`hover`: its link reached in steps, apple.com's
// script opening it 120ms on) at each width, then the pointer into the
// flyout's room at the window's right edge, where it hovers nothing: the bar
// while it is open, its links, the flyout, its content, its groups (the
// first elevated, its links large), their headings, lists, rows and links,
// the curtain over the page, the button beside the link that the keyboard
// opens it by and its chevron; then the pointer along the bar to iPad's link
// (the change from one flyout to another) and off the nav to the page (the
// flyout closing), each recording what it starts. From 834px; to 833px the
// links are the menu's (flyout-phone). The page is read with its locale
// banner above the nav (70px), so a box's top is the nav's top plus its own.
func flyout() pagespec.Spec {
	FLY := "#globalnav-submenu-link-mac"
	ELEVATED := FLY + " .globalnav-submenu-group-elevated"
	GROUP := FLY + " .globalnav-submenu-group:nth-child(2)"
	LAST := FLY + " .globalnav-submenu-group:nth-child(3)"
	TRIGGER := "#globalnav-menubutton-link-mac"
	SPACE := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-right", "margin-bottom", "margin-left"}
	ITEM := slices.Concat([]string{"opacity", "transform", "transition-delay"}, MOTION)
	SIZE := []string{"flex-grow", "flex-shrink", "flex-basis", "max-width", "width"}
	MOTIONS := "#globalnav, #globalnav-curtain"
	// @external: the flyout's room at the window's right edge, below its links' first rows
	ROOM := pagespec.Point(func(w int) [2]float64 { return [2]float64{float64(w) - 10, 400} })

	return pagespec.Spec{
		URL:    "https://www.apple.com/",
		Widths: []int{1600, 1200, 900, 834},
		Wait:   2500,
		Probes: []pagespec.Probe{
			{Name: "nav", Hover: "#globalnav .globalnav-link-mac", Pointer: ROOM, Animations: MOTIONS, Select: "#globalnav", Props: slices.Concat([]string{"height", "position", "background-color"}, BLUR, []string{"z-index"}), Vars: []string{"--globalnav-background", "--globalnav-backdrop-filter"}, Box: true, Position: true},
			{Name: "nav.content", Select: "#globalnav .globalnav-content", Props: slices.Concat([]string{"position", "z-index", "background-color"}, BLUR), Box: true, Position: true},
			// @external: the open flyout's link, another, the search beside them
			{Name: "link.open", Select: "#globalnav .globalnav-link-mac", Props: slices.Concat([]string{"color", "opacity"}, MOTION)},
			{Name: "link.other", Select: "#globalnav .globalnav-link-ipad", Props: []string{"color", "opacity"}},
			{Name: "link.search", Select: "#globalnav .globalnav-link-search", Props: []string{"color", "opacity"}},
			{Name: "flyout", Select: FLY, Props: slices.Concat([]string{"position", "top", "left", "right", "height", "max-height"}, SPACE, []string{"background-color"}, BLUR, []string{"color", "overflow-y", "z-index", "visibility"}, MOTION), Vars: []string{"--r-globalnav-flyout-height", "--r-globalnav-flyout-rate"}, Box: true, Position: true},
			{Name: "flyout.scroll", Select: FLY + " .globalnav-flyout-scroll-container", Props: slices.Concat([]string{"height", "max-height", "overflow-y"}, SPACE), Box: true, Position: true},
			{Name: "flyout.content", Select: FLY + " .globalnav-flyout-content", Props: slices.Concat([]string{"max-width"}, SPACE, FLEX, []string{"box-sizing"}), Box: true, Position: true},
			{Name: "group.elevated", Select: ELEVATED, Props: slices.Concat(SIZE, SPACE, []string{"display"}), Box: true, Position: true},
			{Name: "group", Select: GROUP, Props: slices.Concat(SIZE, SPACE, []string{"display"}), Box: true, Position: true},
			{Name: "group.last", Select: LAST, Props: slices.Concat(SIZE, SPACE), Box: true, Position: true},
			{Name: "header.elevated", Select: ELEVATED + " .globalnav-submenu-header", Props: slices.Concat(TYPE, SPACE, []string{"display"}, ITEM), Box: true, Position: true},
			{Name: "header", Select: GROUP + " .globalnav-submenu-header", Props: slices.Concat(TYPE, SPACE, ITEM), Box: true, Position: true},
			{Name: "header.last", Select: LAST + " .globalnav-submenu-header", Props: ITEM},
			{Name: "list", Select: GROUP + " .globalnav-submenu-list", Props: slices.Concat(SPACE, []string{"list-style-type", "display"}), Box: true, Position: true},
			{Name: "item.elevated", Select: ELEVATED + " .globalnav-submenu-list-item-elevated", Props: slices.Concat(TYPE, SPACE, []string{"display"}, ITEM), Box: true, Position: true},
			{Name: "item.elevated2", Select: ELEVATED + " .globalnav-submenu-list-item-elevated:nth-child(2)", Props: slices.Concat(SPACE, ITEM), Box: true, Position: true},
			{Name: "item.after", Select: ELEVATED + " .globalnav-submenu-list-item", Props: slices.Concat(TYPE, SPACE, ITEM), Box: true, Position: true},
			{Name: "item.after2", Select: ELEVATED + " .globalnav-submenu-list-item ~ .globalnav-submenu-list-item", Props: slices.Concat(SPACE, ITEM), Box: true, Position: true},
			{Name: "item", Select: GROUP + " .globalnav-submenu-list-item", Props: slices.Concat(TYPE, SPACE, ITEM), Box: true, Position: true},
			{Name: "item2", Select: GROUP + " .globalnav-submenu-list-item:nth-child(2)", Props: slices.Concat(SPACE, ITEM), Box: true, Position: true},
			{Name: "item.last", Select: LAST + " .globalnav-submenu-list-item", Props: ITEM},
			{Name: "link.elevated", Select: ELEVATED + " .globalnav-submenu-list-item-elevated .globalnav-submenu-link", Props: slices.Concat(TYPE, SPACE, []string{"display", "text-decoration-line", "background-color", "border-top-left-radius"}, FOCUS, MOTION), Box: true, Position: true, States: []string{"hover", "focus-visible"}},
			{Name: "link.after", Select: ELEVATED + " .globalnav-submenu-list-item .globalnav-submenu-link", Props: slices.Concat(TYPE, SPACE, []string{"display"}, MOTION), Box: true, Position: true, States: []string{"hover"}},
			{Name: "link", Select: GROUP + " .globalnav-submenu-link", Props: slices.Concat(TYPE, SPACE, []string{"display", "text-decoration-line", "background-color"}, FOCUS, MOTION), Box: true, Position: true, States: []string{"hover", "focus-visible"}},
			{Name: "curtain", Select: "#globalnav-curtain", Props: slices.Concat([]string{"position", "top", "z-index", "background-color"}, BLUR, []string{"opacity", "visibility", "pointer-events", "transition-delay"}, MOTION), Box: true, Position: true},
			// @external: the button beside the link the keyboard opens the flyout by: its box, its
			// chevron shown focused
			{Name: "trigger", Select: TRIGGER, Props: slices.Concat([]string{"position", "top", "left", "right", "width", "height"}, BOX, []string{"color", "opacity", "cursor"}, FLEX, FOCUS, MOTION), Box: true, Position: true, States: []string{"focus-visible", "hover"}},
			{Name: "trigger.icon", Select: TRIGGER + " .globalnav-image-regular", StateOn: TRIGGER, Props: slices.Concat([]string{"display", "width", "height"}, SPACE, []string{"opacity", "color"}, MOTION), Box: true, Position: true, States: []string{"focus-visible", "hover"}},
			{Name: "trigger.compact", Select: TRIGGER + " .globalnav-image-compact", Props: []string{"display"}},
			{Name: "trigger.group", Select: "#globalnav .globalnav-item-mac .globalnav-submenu-trigger-group", Props: slices.Concat([]string{"position", "display", "width", "height"}, SPACE, []string{"list-style-type"}), Box: true, Position: true},
			{Name: "trigger.item", Select: "#globalnav .globalnav-item-mac .globalnav-submenu-trigger-item", Props: []string{"display", "position"}, Box: true, Position: true},
			// @external: the change: the pointer along the bar to iPad's link
			{Name: "change", Hover: "#globalnav .globalnav-link-ipad", Pointer: ROOM, Animations: MOTIONS, Select: "#globalnav-submenu-link-ipad", Props: slices.Concat([]string{"height", "background-color", "visibility"}, MOTION), Vars: []string{"--r-globalnav-flyout-height", "--r-globalnav-flyout-rate"}, Box: true},
			{Name: "change.previous", Select: FLY, Props: []string{"height", "visibility"}, Box: true},
			{Name: "change.link", Select: "#globalnav .globalnav-link-mac", Props: []string{"color"}},
			// @external: and off the nav, down to the page: it closes
			{Name: "close", HoverAt: func(w int) [2]float64 { return [2]float64{float64(w) / 2, 890} }, Hovered: 1200, Animations: MOTIONS, Select: "#globalnav-submenu-link-ipad", Props: []string{"height", "visibility", "background-color"}, Box: true},
			{Name: "close.curtain", Select: "#globalnav-curtain", Props: []string{"opacity", "visibility"}},
			{Name: "close.nav", Select: "#globalnav", Props: slices.Concat([]string{"background-color"}, BLUR)},
		},
	}
}
