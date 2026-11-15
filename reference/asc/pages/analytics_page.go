package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var analyticsPage = func() signedin.Spec {
	item := []string{"height", "background-color", "border-top-left-radius", "padding-left", "padding-right", "column-gap", "color", "opacity", "cursor"}
	label := "> div > div"
	return signedin.Spec{
		Path: "/apps/<app>/analytics/overview",
		Probes: []signedin.Probe{
			{Name: "card", Select: `a:has(p[class*="rt-Text--textStyle--$bodyEmph"])`, Props: slices.Concat(signedin.Box, []string{"overflow"}), Box: true},
			{Name: "section.heading", Select: `p[class*="rt-Text--textStyle--$h4Emph"]`, Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "section.link", Select: `a[class*="rt-Link--variant--inline"]`, Props: slices.Concat(signedin.Type, signedin.Link, []string{"cursor"}), Box: true},
			{Name: "menu", Select: `nav[class*="rt-DisplayList--variant--menu"]`, Props: slices.Concat(signedin.Box, []string{"row-gap", "display"}), Box: true},
			{Name: "group.button", Select: `button[class*="rt-Item"]`, Props: item, Box: true},
			{Name: "group.button.label", Select: `button[class*="rt-Item"] ` + label, Props: signedin.Type, Box: true},
			{Name: "group.chevron", Select: `button[class*="rt-Item"] svg`, Props: []string{"width", "height", "color", "fill", "transform"}, Box: true},
			{Name: "group.item", Select: `[class*="rt-DisclosureGroup--state--open"] > a[class*="rt-Item--effect--default"]`, Props: item, Box: true},
			{Name: "group.item.label", Select: `[class*="rt-DisclosureGroup--state--open"] > a[class*="rt-Item--effect--default"] ` + label, Props: slices.Concat(signedin.Type, []string{"margin-left", "padding-left"}), Box: true},
			{Name: "item.disabled", Select: `a[class*="rt-Item--effect--disabled"]`, Props: item, Box: true},
			{Name: "item.disabled.label", Select: `a[class*="rt-Item--effect--disabled"] ` + label, Props: signedin.Type, Box: true},
		},
	}
}()
