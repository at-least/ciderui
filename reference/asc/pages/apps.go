package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var apps = func() signedin.Spec {
	bar := []string{"height", "position", "top", "z-index", "background-color", "border-bottom-width", "border-bottom-style", "border-bottom-color", "box-shadow", "backdrop-filter"}
	icon := []string{"width", "height", "fill", "color"}
	banner := `div:has(> button[class*="Banner__DismissButtonWrapper"])`
	return signedin.Spec{
		Path: "/apps",
		Probes: []signedin.Probe{
			{Name: "page", Select: "html", Props: []string{"background-color", "color-scheme"}},
			{Name: "nav", Shadow: "amp-nav", Select: "header.primary-header", Props: slices.Concat(bar, []string{"padding-left", "padding-right"}), Box: true},
			{Name: "nav.content", Shadow: "amp-nav", Select: ".primary-header-content", Props: []string{"max-width", "padding-left", "padding-right", "display", "align-items", "justify-content"}, Box: true},
			{Name: "nav.logo", Shadow: "amp-nav", Select: ".logo", Props: []string{"width", "height", "color"}, Box: true},
			{Name: "nav.link", Shadow: "amp-nav", Select: ".primary-nav a", Nth: 1, Props: slices.Concat(signedin.Type, signedin.Link, []string{"padding-bottom", "border-bottom-width", "border-bottom-color"}), Box: true},
			{Name: "nav.current", Shadow: "amp-nav", Select: ".primary-nav a.active-item", Props: slices.Concat(signedin.Type, signedin.Link, []string{"padding-bottom", "border-bottom-width", "border-bottom-color"}), Box: true},
			{Name: "header", Select: `[class*="Sticky__StickyWrapper"]`, Props: []string{"position", "top", "z-index", "background-color", "padding-top", "padding-bottom"}, Box: true},
			{Name: "title", Select: `[class*="Sticky__StickyWrapper"] h2`, Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "add.icon", Select: `[class*="Sticky__StickyWrapper"] button[class*="Link-sc"] svg`, Props: icon, Box: true},
			{Name: "filter", Select: `button[class*="Link-sc"]`, Text: "All Statuses", Props: slices.Concat(signedin.Type, signedin.Link, []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "cursor"}), Box: true},
			{Name: "rule", Select: `[class*="Sticky__StickyWrapper"] hr`, Props: []string{"height", "border-top-width", "border-top-style", "border-top-color", "background-color", "margin-top", "margin-bottom"}, Box: true},
			{Name: "banner", Select: banner, Props: slices.Concat(signedin.Box, []string{"display", "align-items", "column-gap"}), Box: true},
			{Name: "banner.icon", Select: banner + " svg.glyph-box", Props: icon, Box: true},
			{Name: "banner.title", Select: banner + " p", Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "banner.text", Select: banner + " p", Nth: 1, Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "banner.link", Select: banner + " p a", Props: slices.Concat(signedin.Type, signedin.Link)},
			{Name: "banner.dismiss", Select: `button[class*="Banner__DismissButtonWrapper"]`, Props: []string{"position", "top", "right", "width", "height", "cursor"}, Box: true},
			{Name: "banner.dismiss.icon", Select: `button[class*="Banner__DismissButtonWrapper"] svg`, Props: icon, Box: true},
			{Name: "banner.info", Select: banner, Nth: 1, Props: slices.Concat(signedin.Box, []string{"display", "align-items", "column-gap"}), Box: true},
			{Name: "banner.info.icon", Select: banner + " > svg", Props: icon, Box: true},
			{Name: "banner.info.dismiss.icon", Select: `button[class*="Banner__DismissButtonWrapper"] svg`, Nth: 1, Props: icon},
			{Name: "app.icon", Select: `[role=grid] a[class*="Link-sc"]`, Props: []string{"width", "height", "border-top-left-radius", "box-shadow", "overflow"}, Box: true},
			{Name: "app.name", Select: `[role=grid] p[class*="Text__BaseText"]`, Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "app.status", Select: `[role=grid] ul[class*="List__BaseList"] a`, Props: slices.Concat(signedin.Type, signedin.Link), Box: true},
			{Name: "app.status.icon", Select: `[role=grid] ul[class*="List__BaseList"] svg`, Props: slices.Concat(icon, []string{"margin-right", "vertical-align"}), Box: true},
		},
	}
}()
