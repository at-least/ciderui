//go:build browser

package consoledl

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const d = "component-DL"

var typeProps = []string{"font-size", "line-height", "font-weight"}

func style(part, scale, variant string) []string {
	v := ""
	if variant != "" {
		v = "-variant-" + variant
	}
	var names []string
	for _, p := range []string{"fontSize", "lineHeight", "fontWeight"} {
		names = append(names, d+"-anatomy-"+part+"-variant-pageScale-"+scale+v+"-textStyle-"+p)
	}
	return names
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "dl",
			HTML:   h.Example("dl"),
			Widths: []int{1200, 640},
			Checks: []h.Check{
				{Target: ".cui-console-dl", Props: []string{"column-gap", "row-gap"}, TB: []string{d + "-variant-orientation-horizontal-pageScale-desktop-columnGap", d + "-variant-orientation-horizontal-pageScale-desktop-rowGap"}},
				{Target: "dt:first-of-type", Props: typeProps, TB: style("DT", "desktop", "muted")},
				{Target: "dd:first-of-type", Props: h.Props(typeProps, []string{"color"}), TB: h.Props(style("DD", "desktop", ""), []string{d + "-anatomy-DD-color"})},
				{Target: "dt:first-of-type", Props: []string{"color"}, Golden: "the muted label, the secondary label at 0.76: the toolkit's 0.60 is 3.4:1 on white (tokens.css, decision 21)"},
				{Target: ".cui-console-dl", Props: []string{"grid-template-columns"}, Golden: "a label column as wide as its longest label, the values in the rest"},
				{Target: "dd:last-of-type", Props: []string{"overflow-wrap"}, Golden: "a value breaks anywhere it must: an id, a URL"},
			},
		},
		{
			Name:   "dl phone",
			HTML:   h.Example("dl"),
			Widths: []int{639, 390},
			Checks: []h.Check{
				{Target: ".cui-console-dl", Props: []string{"column-gap", "row-gap"}, TB: []string{d + "-variant-orientation-horizontal-pageScale-mobile-columnGap", d + "-variant-orientation-horizontal-pageScale-mobile-rowGap"}},
				{Target: "dt:first-of-type", Props: typeProps, TB: style("DT", "mobile", "muted")},
				{Target: "dd:first-of-type", Props: typeProps, TB: style("DD", "mobile", "")},
			},
		},
		{
			Name:   "dl dark",
			HTML:   `<div class="cui-dark">` + h.Example("dl") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: "dt:first-of-type", Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: "dd:first-of-type", Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
