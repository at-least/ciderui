package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var testflight = func() signedin.Spec {
	icon := []string{"width", "height", "fill", "color"}
	cell := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "vertical-align", "text-align", "background-color", "border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-style", "border-bottom-color"}
	row := []string{"height", "background-color", "border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-style", "border-bottom-color"}
	return signedin.Spec{
		Path: "/teams/<team>/apps/<app>/testflight/ios",
		Probes: []signedin.Probe{
			{Name: "tabs", Select: "nav.sub-nav-container", Props: []string{"height", "border-bottom-width", "border-bottom-style", "border-bottom-color"}, Box: true},
			{Name: "tabs.link", Select: "nav.sub-nav-container a", Props: slices.Concat(signedin.Type, signedin.Link, []string{"padding-left", "padding-right", "border-bottom-width", "border-bottom-style", "border-bottom-color"}), Box: true},
			{Name: "tabs.current", Select: "nav.sub-nav-container a", Nth: 2, Props: slices.Concat(signedin.Type, signedin.Link, []string{"padding-left", "padding-right", "border-bottom-width", "border-bottom-style", "border-bottom-color"}), Box: true},
			{Name: "sidebar", Select: "[role=navigation]", Props: []string{"width", "padding-right"}, Box: true},
			{Name: "sidebar.heading", Select: "[role=navigation] h2 span", Props: signedin.Type, Box: true},
			{Name: "sidebar.link", Select: "[role=navigation] a:not([aria-current])", Props: slices.Concat(signedin.Type, signedin.Link, []string{"padding-left", "padding-right", "height"}), Box: true},
			{Name: "sidebar.link.row", Select: "[role=navigation] li:has(> a:not([aria-current]))", Props: []string{"background-color", "border-top-left-radius", "height", "margin-bottom"}, Box: true},
			{Name: "sidebar.current", Select: "[role=navigation] a[aria-current]", Props: slices.Concat(signedin.Type, signedin.Link, []string{"padding-left", "padding-right", "height"}), Box: true},
			{Name: "sidebar.current.row", Select: "[role=navigation] li:has(> a[aria-current])", Props: []string{"background-color", "border-top-left-radius", "height", "margin-bottom"}, Box: true},
			{Name: "sidebar.rule", Select: "[role=navigation] hr", Props: []string{"height", "background-color", "border-top-style", "margin-top", "margin-bottom"}, Box: true},
			{Name: "title", Select: "h2:not([role=navigation] h2)", Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "group", Select: "button[type=button][aria-expanded=true]", Props: slices.Concat(signedin.Box, []string{"height", "display", "align-items", "column-gap"}), Box: true},
			{Name: "group.chevron", Select: "button[type=button][aria-expanded=true] svg", Props: slices.Concat(icon, []string{"transform"}), Box: true},
			{Name: "group.title", Select: "button[type=button][aria-expanded=true] h3", Props: slices.Concat(signedin.Type, []string{"margin-left"}), Box: true},
			{Name: "table", Select: "table[role=grid]", Props: slices.Concat([]string{"border-collapse", "background-color", "margin-top", "margin-left"}, signedin.Type), Box: true},
			{Name: "head", Select: "thead tr", Props: row, Box: true},
			{Name: "head.cell", Select: "thead th", Props: slices.Concat(cell, []string{"height"}), Box: true},
			{Name: "head.label", Select: `thead th span[class*="Text__BaseText"]`, Props: slices.Concat(signedin.Type, []string{"text-transform"}), Box: true},
			{Name: "row", Select: "tbody tr", Props: row, Box: true},
			{Name: "cell", Select: "tbody td[role=gridcell]", Props: slices.Concat(cell, signedin.Type), Box: true},
			{Name: "cell.first", Select: "tbody td[role=rowheader]", Props: slices.Concat(cell, signedin.Type), Box: true},
			{Name: "build.link", Select: `tbody td[role=rowheader] span[class*="Link-sc"]`, Props: slices.Concat(signedin.Type, signedin.Link)},
			{Name: "status.icon", Select: `tbody td[role=gridcell] svg[class*="SvgIcon"]`, Props: slices.Concat(icon, []string{"margin-right", "vertical-align"}), Box: true},
			{Name: "status.word", Select: `tbody td[role=gridcell] svg[class*="SvgIcon"] + span`, Props: slices.Concat(signedin.Type, []string{"margin-left"}), Box: true},
			{Name: "status.line", Select: `tbody td[role=gridcell] span[class*="Text__BaseText"]`, Nth: 1, Props: slices.Concat(signedin.Type, []string{"display", "margin-top"}), Box: true},
			{Name: "count", Select: "tbody tr td:nth-child(5) p", Props: slices.Concat(signedin.Type, []string{"text-align"}), Box: true},
			{Name: "info.icon", Select: `tbody button[class*="Link-sc"] svg`, Props: icon, Box: true},
		},
	}
}()
