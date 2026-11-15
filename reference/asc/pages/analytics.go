package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var analytics = func() signedin.Spec {
	icon := []string{"width", "height", "fill", "color"}
	item := []string{"height", "background-color", "border-top-left-radius", "padding-left", "padding-right", "column-gap", "color"}
	label := `span[class*="rt-Text--trim--cap-alphabetic"]`
	step := []string{"height", "background-color", "color", "border-top-left-radius", "border-top-right-radius", "border-bottom-left-radius", "border-bottom-right-radius", "padding-left", "padding-right", "margin-left", "border-left-width", "border-left-style", "border-left-color"}
	stop := []string{"stop-color", "stop-opacity"}
	return signedin.Spec{
		Path: "/apps/<app>/analytics/overview",
		Probes: []signedin.Probe{
			{Name: "menu.current", Select: `nav[class*="rt-DisplayList"] a[class*="rt-Item"]`, Props: item, Box: true},
			{Name: "menu.current.label", Select: `nav[class*="rt-DisplayList"] a[class*="rt-Item"] > div > div`, Props: signedin.Type, Box: true},
			{Name: "menu.item", Select: `nav[class*="rt-DisplayList"] a[class*="rt-Item"]`, Nth: 1, Props: item, Box: true},
			{Name: "menu.item.label", Select: `nav[class*="rt-DisplayList"] a[class*="rt-Item"] > div > div`, Nth: 1, Props: signedin.Type, Box: true},
			{Name: "menu.icon", Select: `nav[class*="rt-DisplayList"] a[class*="rt-Item"] svg`, Props: icon, Box: true},
			{Name: "menu.rule", Select: `nav[class*="rt-DisplayList"] hr`, Props: []string{"height", "background-color", "border-top-style", "margin-top", "margin-bottom", "margin-left", "margin-right"}, Box: true},
			{Name: "date.previous", Select: "button:has(+ button > " + label + ")", Props: step, Box: true},
			{Name: "date.period", Select: "button:has(> " + label + ")", Props: step, Box: true},
			{Name: "date.period.label", Select: "button > " + label, Props: signedin.Type, Box: true},
			{Name: "date.next", Select: "button:has(> " + label + ") + button", Props: step, Box: true},
			{Name: "date.chevron", Select: "button:has(+ button > " + label + ") svg", Props: icon, Box: true},
			{Name: "card.title", Select: `p[class*="rt-Text--textStyle--$bodyEmph"]`, Props: signedin.Type, Box: true},
			{Name: "card.info", Select: `button[class*="rt-InfoSystemButton"]`, Props: []string{"width", "height", "color", "background-color", "border-top-left-radius"}, Box: true},
			{Name: "card.info.icon", Select: `button[class*="rt-InfoSystemButton"] svg`, Props: icon, Box: true},
			{Name: "card.value", Select: `h2[class*="rt-Text--textStyle--$h1PlusEmph"]`, Props: slices.Concat(signedin.Type, []string{"font-variant-numeric"}), Box: true},
			{Name: "card.change", Select: `span[class*="rt-Text--textStyle--$h4Medium"]`, Props: signedin.Type, Box: true},
			{Name: "spark.area", Select: "svg g.group.path path", Props: []string{"fill", "fill-opacity", "opacity", "stroke"}, Box: true},
			{Name: "spark.line", Select: "svg g.group.path path", Nth: 1, Props: []string{"fill", "stroke", "stroke-width", "stroke-linecap", "stroke-linejoin", "opacity"}, Box: true},
			{Name: "spark.top", Select: `linearGradient[id^="linear_"] stop`, Unseen: true, Props: stop},
			{Name: "spark.bottom", Select: `linearGradient[id^="linear_"] stop`, Nth: 1, Unseen: true, Props: stop},
			{Name: "chart.grid", Select: "svg g > line", Props: []string{"stroke", "stroke-width", "stroke-dasharray", "opacity"}, Box: true},
			{Name: "chart.y", Select: "svg g.axis.y text", Props: slices.Concat(signedin.Type, []string{"fill", "text-anchor"}), Box: true},
			{Name: "chart.x", Select: "svg g.axis.x text", Props: slices.Concat(signedin.Type, []string{"fill", "text-anchor"}), Box: true},
			{Name: "chart.bar", Select: `svg path[fill*="lg_"]`, Props: []string{"fill", "stroke", "opacity"}, Box: true},
			{Name: "chart.bar.top", Select: `linearGradient[id^="lg_"] stop`, Unseen: true, Props: stop},
			{Name: "chart.bar.bottom", Select: `linearGradient[id^="lg_"] stop`, Nth: 1, Unseen: true, Props: stop},
			{Name: "table.row", Select: `table[class*="rt-Table"] tr`, Props: []string{"height", "background-color", "border-bottom-width", "border-bottom-style", "border-bottom-color"}, Box: true},
			{Name: "table.cell", Select: `table[class*="rt-Table"] td`, Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "vertical-align", "text-align", "border-bottom-width", "border-bottom-style", "border-bottom-color"}, Box: true},
			{Name: "table.text", Select: `table[class*="rt-Table"] td p`, Props: slices.Concat(signedin.Type, []string{"text-align"}), Box: true},
			{Name: "table.figure", Select: `table[class*="rt-Table"] td p`, Nth: 1, Props: slices.Concat(signedin.Type, []string{"text-align", "font-variant-numeric"}), Box: true},
		},
	}
}()
