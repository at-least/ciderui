package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: an Apple Support article's text (the owner, 2026-11-13: what consumers write
// by hand, measured on Apple's sites; totality's store and website write markdown's
// elements by hand): its paragraphs, headings, lists and their items, a picture in a
// step, bold. The article with each of them: support's own (`support`) has no heading,
// numbered list or picture. Its rules are by the element before (`p + p`, `h2`), as
// markdown writes them: no class of the page's own reaches a consumer's text.
func supportProse() pagespec.Spec {
	FLOW := []string{"margin-top", "margin-bottom"}
	return pagespec.Spec{
		URL:  "https://support.apple.com/en-us/102631",
		Wait: 3000,
		Probes: []pagespec.Probe{
			{Name: "p", Select: "p.gb-paragraph + p.gb-paragraph", Props: slices.Concat(TYPE, FLOW)},
			{Name: "h2", Select: "h2.gb-header", Props: slices.Concat(TYPE, FLOW)},
			{Name: "h2.next", Select: "h2.gb-header + p.gb-paragraph", Props: FLOW},
			{Name: "h3", Select: "h3.gb-header", Props: slices.Concat(TYPE, FLOW)},
			{Name: "ul", Select: "p.gb-paragraph + ul.gb-list", Props: slices.Concat(FLOW, []string{"padding-left", "margin-left", "list-style-type", "list-style-position"})},
			{Name: "ol", Select: "ol.gb-list", Props: slices.Concat(FLOW, []string{"padding-left", "margin-left", "list-style-type", "list-style-position"})},
			{Name: "li", Select: "ul.gb-list > li", Props: slices.Concat(TYPE, FLOW, []string{"padding-left"})},
			{Name: "li.next", Select: "ul.gb-list > li + li", Props: FLOW},
			{Name: "ol.li", Select: "ol.gb-list > li", Props: slices.Concat(TYPE, FLOW)},
			{Name: "li.marker", Select: "ol.gb-list > li", Pseudo: "::marker", Props: TYPE},
			{Name: "list.next", Select: "ul.gb-list + p.gb-paragraph, ol.gb-list + p.gb-paragraph", Props: FLOW},
			{Name: "img", Select: "img.gb-image", Props: slices.Concat(FLOW, []string{"width", "max-width", "height", "display", "border-top-left-radius"}), Box: true},
			{Name: "b", Select: "p.gb-paragraph b", Props: TYPE},
			{Name: "link", Select: "p.gb-paragraph a", Props: slices.Concat(TYPE, LINK), States: []string{"hover"}},
		},
	}
}
