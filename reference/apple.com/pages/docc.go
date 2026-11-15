package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

func doccProbes() []pagespec.Probe {
	ASIDE := slices.Concat(BOX, TYPE, []string{"border-left-width", "border-left-style", "border-left-color", "border-right-color"})
	return []pagespec.Probe{
		{Name: "page", Select: "body", Props: PAGE},
		{Name: "eyebrow", Select: ".eyebrow", Props: TYPE},
		{Name: "title", Select: "h1.title", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
		{Name: "abstract", Select: ".abstract", Props: TYPE},
		{Name: "heading", Select: "h2:has(+ p)", Props: slices.Concat(TYPE, []string{"margin-top", "margin-bottom"})},
		{Name: "text", Select: "h2 + p", Props: slices.Concat(TYPE, []string{"margin-top"})},
		{Name: "link", Select: "p a.inline-link", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
		{Name: "code", Select: "p code.inline-code", Props: slices.Concat(TYPE, []string{"background-color", "padding-left", "padding-right", "border-top-left-radius"})},
		{Name: "listing", Select: ".code-listing", Props: slices.Concat(BOX, []string{"display", "overflow-x"}), Box: true},
		{Name: "listing.scroller", Select: ".code-listing .container-general", Props: []string{"display", "overflow-x"}, Box: true},
		{Name: "listing.pre", Select: ".code-listing pre", Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-bottom", "white-space"}, Box: true},
		{Name: "listing.code", Select: ".code-listing pre code", Props: slices.Concat(TYPE, []string{"white-space", "tab-size"})},
		{Name: "listing.line", Select: ".code-listing .code-line-container", Props: []string{"display", "min-height", "padding-left", "padding-right", "background-color"}, Box: true},
		{Name: "copy", Select: ".code-listing button.copy-button", Props: slices.Concat(BOX, []string{"width", "height", "position", "top", "right", "color", "opacity", "cursor"}, FOCUS, MOTION), Box: true, States: []string{"hover", "focus-visible"}},
		{Name: "copy.icon", Select: ".code-listing button.copy-button svg", Props: []string{"width", "height", "color", "fill"}, Box: true},
		{Name: "bar", Select: "nav.documentation-nav .nav__wrapper", Props: []string{"height", "background-color", "border-bottom-width", "border-bottom-style", "border-bottom-color", "backdrop-filter", "-webkit-backdrop-filter"}, Box: true},
		{Name: "bar.ground", Select: "nav.documentation-nav .nav__background", Props: []string{"background-color", "border-bottom-width", "border-bottom-style", "border-bottom-color", "box-shadow", "opacity", "backdrop-filter", "-webkit-backdrop-filter"}, Box: true},
		{Name: "bar.title", Select: "nav.documentation-nav .nav-title", Props: TYPE, Box: true},
		{Name: "bar.toggle", Select: "nav.documentation-nav button.sidenav-toggle", Widths: []int{390}, Props: []string{"width", "height", "color", "background-color"}, Box: true},
		{Name: "navigator", Select: ".navigator", Props: []string{"display", "width", "border-right-width", "border-right-style", "border-right-color", "background-color"}, Box: true},
		{Name: "navigator.item", Select: ".navigator-card-item:not(.active)", Props: slices.Concat(TYPE, []string{"height", "padding-left", "padding-right", "background-color", "border-top-left-radius"}), Box: true, States: []string{"hover"}},
		{Name: "navigator.link", Select: ".navigator-card-item:not(.active) a.leaf-link", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
		{Name: "navigator.current", Select: ".navigator-card-item.active", Props: slices.Concat(TYPE, []string{"height", "padding-left", "background-color", "border-top-left-radius"}), Box: true},
		{Name: "navigator.current.link", Select: ".navigator-card-item.active a.leaf-link", Props: slices.Concat(TYPE, LINK)},
		{Name: "note", Select: "aside.note", Props: ASIDE, Box: true},
		{Name: "note.label", Select: "aside.note p.label", Props: slices.Concat(TYPE, []string{"margin-bottom"})},
		{Name: "note.text", Select: "aside.note p:not(.label)", Props: TYPE},
	}
}

func docc() pagespec.Spec {
	return pagespec.Spec{
		URL:    "https://developer.apple.com/documentation/xcode/examining-the-fields-in-a-crash-report",
		Wait:   4000,
		Probes: doccProbes(),
	}
}
