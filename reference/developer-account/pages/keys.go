package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var keys = func() signedin.Spec {
	pad := []string{"padding-top", "padding-right", "padding-bottom", "padding-left"}
	rule := []string{"border-bottom-width", "border-bottom-style", "border-bottom-color"}
	row := "[role=row]:not(.th)"
	return signedin.Spec{
		Path: "/account/resources/authkeys/list",
		Probes: []signedin.Probe{
			{Name: "page", Select: ".main-wrapper", Props: signedin.Type},
			{Name: "nav", Select: "nav.app-nav", Props: []string{"width", "margin-right"}, Box: true},
			{Name: "nav.item", Select: "nav.app-nav li:not(.active)", Props: []string{"height", "margin-bottom", "background-color", "border-top-left-radius"}, Box: true},
			{Name: "nav.word", Select: "nav.app-nav li:not(.active) > span", Props: slices.Concat(signedin.Type, pad), Box: true},
			{Name: "nav.current", Select: "nav.app-nav li.active", Props: []string{"background-color", "border-top-left-radius"}, Box: true},
			{Name: "nav.current.word", Select: "nav.app-nav li.active > span", Props: slices.Concat(signedin.Type, pad)},
			{Name: "header", Select: ".list-header", Props: slices.Concat(pad, rule, []string{"height", "margin-bottom"}), Box: true},
			{Name: "title", Select: ".list-header h2", Props: slices.Concat(signedin.Type, []string{"margin-top", "margin-bottom"}), Box: true},
			{Name: "add", Select: ".list-header h2 ~ span a svg", Props: []string{"width", "height", "color", "fill", "margin-left"}, Box: true},
			{Name: "filter", Select: `.list-header button[class*="Link-sc"]`, Props: slices.Concat(signedin.Type, []string{"cursor"}), Box: true},
			{Name: "filter.chevron", Select: `.list-header button[class*="Link-sc"] svg`, Props: []string{"width", "height", "color", "fill", "margin-left"}, Box: true},
			{Name: "column", Select: "[role=columnheader]:not(.sortable)", Props: slices.Concat(signedin.Type, pad, rule, []string{"background-color", "text-align"}), Box: true},
			{Name: "column.sortable", Select: "[role=columnheader].sortable", Props: slices.Concat(signedin.Type, []string{"cursor"}), Box: true},
			{Name: "row", Select: row, Props: []string{"height", "background-color", "cursor"}, Box: true},
			{Name: "cell", Select: row + " > [role=cell]", Props: slices.Concat(signedin.Type, pad, rule, []string{"vertical-align"}), Box: true},
			{Name: "cell.glyph", Select: row + " [role=cell] img", Props: []string{"width", "height", "vertical-align"}, Box: true},
		},
	}
}()
