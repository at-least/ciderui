package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

// @external: The app's icon is a div drawn by its background image, inside a box of its size.
var head = func() signedin.Spec {
	room := slices.Concat(signedin.Box, []string{"height", "margin-left", "margin-right"})
	return signedin.Spec{
		Path: "/apps/<app>/analytics/overview",
		Probes: []signedin.Probe{
			{Name: "head", Select: "nav.sub-nav-container", Up: 2, Props: slices.Concat(room, signedin.Flex), Box: true},
			{Name: "head.wrap", Select: "nav.sub-nav-container", Up: 3, Props: slices.Concat(room, signedin.Flex), Box: true},
			{Name: "head.wrap.outer", Select: "nav.sub-nav-container", Up: 4, Props: slices.Concat(room, signedin.Flex), Box: true},
			{Name: "head.name", Select: "div.dd-container", Props: slices.Concat(room, signedin.Flex), Box: true},
			{Name: "head.name.button", Select: "div.dd-container [role=button]", Props: slices.Concat(room, signedin.Flex, []string{"cursor"}), Box: true},
			{Name: "head.name.text", Select: "div.dd-container [role=button] span", Props: slices.Concat(signedin.Type, []string{"margin-left", "margin-right", "margin-top", "margin-bottom"}), Box: true},
			{Name: "head.name.icon", Select: `div.dd-container [role=button] div[class*="Box-sc"]`, Nth: 1, Props: []string{"width", "height", "border-top-left-radius", "margin-left", "margin-right"}, Box: true},
			{Name: "head.name.chevron", Select: "div.dd-container svg", Props: []string{"width", "height", "color", "fill", "margin-left", "margin-right"}, Box: true},
			{Name: "head.tabs", Select: "nav.sub-nav-container", Props: slices.Concat(room, signedin.Flex), Box: true},
			{Name: "head.tab", Select: "ul.sub-nav-links > li", Props: room, Box: true},
			{Name: "head.tab.link", Select: "ul.sub-nav-links a", Props: slices.Concat(signedin.Type, room), Box: true},
			{Name: "head.tab.current", Select: "ul.sub-nav-links a", Nth: 1, Props: slices.Concat(signedin.Type, room), Box: true},
			{Name: "content", Select: `article[class*="rt-Content"]`, Props: room, Box: true},
			{Name: "section.head", Select: `div:has(> p[class*="rt-Text--textStyle--$h4Emph"])`, Props: slices.Concat(room, signedin.Flex), Box: true},
			{Name: "section.title", Select: `p[class*="rt-Text--textStyle--$h4Emph"]`, Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
		},
	}
}()
