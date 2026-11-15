package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var button = signedin.Spec{
	Path: "/apps/<app>/analytics/acquisition/sources",
	Probes: []signedin.Probe{
		{Name: "button", Select: `button[class*="rt-Button--variant--gray"]`, Props: slices.Concat(signedin.Type, signedin.Box, signedin.Text, []string{"height", "column-gap", "border-top-right-radius", "border-bottom-left-radius", "outline-style", "transition-property", "transition-duration"}), Box: true},
	},
}
