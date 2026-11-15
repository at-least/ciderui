package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var crashes = func() signedin.Spec {
	cell := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "vertical-align", "text-align", "background-color", "border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-style", "border-bottom-color"}
	return signedin.Spec{
		Path: "/teams/<team>/apps/<app>/testflight/crashes",
		Probes: []signedin.Probe{
			{Name: "head.cell", Select: "thead th", Props: slices.Concat(cell, []string{"height"}), Box: true},
			{Name: "head.label", Select: `thead th span[class*="Text__BaseText"]`, Props: slices.Concat(signedin.Type, []string{"text-transform"}), Box: true},
			{Name: "row", Select: "tbody tr", Props: []string{"height", "background-color", "border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-style", "border-bottom-color"}, Box: true},
			{Name: "empty.cell", Select: "tbody td", Props: slices.Concat(cell, []string{"height"}), Box: true},
			{Name: "empty", Select: "tbody td > div", Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "text-align", "background-color"}, Box: true},
			{Name: "empty.text", Select: "tbody td p", Props: slices.Concat(signedin.Type, []string{"text-align"}), Box: true},
		},
	}
}()
