package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var hoverButton = signedin.Spec{
	Path: "/apps/<app>/analytics/acquisition/sources",
	Probes: []signedin.Probe{
		{Name: "button", Select: `button[class*="rt-Button--variant--gray"]:hover`, Props: []string{"background-color", "color", "cursor", "opacity"}, Box: true},
	},
}

var hoverDate = signedin.Spec{
	Path: "/apps/<app>/analytics/overview",
	Probes: []signedin.Probe{
		{Name: "date.period", Select: `button:hover:has(> span[class*="rt-Text--trim--cap-alphabetic"])`, Props: []string{"background-color", "color", "opacity", "cursor", "filter"}, Box: true},
	},
}

var hoverLink = signedin.Spec{
	Path: "/apps/<app>/analytics/overview",
	Probes: []signedin.Probe{
		{Name: "section.link", Select: `a[class*="rt-Link--variant--inline"]:hover`, Props: slices.Concat([]string{"color"}, signedin.Link, []string{"opacity", "cursor"}), Box: true},
	},
}

// @external: The analytics menu's item under the pointer (the page marks it
// rt-Item--effect--hover): read with the pointer resting on a menu item that
// is not the current one.
var hoverMenu = func() signedin.Spec {
	item := []string{"height", "background-color", "border-top-left-radius", "padding-left", "padding-right", "color", "cursor"}
	return signedin.Spec{
		Path: "/apps/<app>/analytics/overview",
		Probes: []signedin.Probe{
			{Name: "menu.item", Select: `nav[class*="rt-DisplayList"] a[class*="rt-Item--effect--hover"]`, Props: item, Box: true},
			{Name: "menu.item.label", Select: `nav[class*="rt-DisplayList"] a[class*="rt-Item--effect--hover"] > div > div`, Props: []string{"color", "font-weight"}, Box: true},
		},
	}
}()

var hoverRow = signedin.Spec{
	Path: "/apps/<app>/analytics/overview",
	Probes: []signedin.Probe{
		{Name: "table.row", Select: `table[class*="rt-Table"] tr:hover`, Props: []string{"background-color", "border-bottom-color", "cursor"}, Box: true},
		{Name: "table.cell", Select: `table[class*="rt-Table"] tr:hover td`, Props: []string{"background-color", "color", "border-bottom-color"}, Box: true},
	},
}
