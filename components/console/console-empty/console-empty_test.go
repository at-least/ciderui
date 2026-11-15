//go:build browser

package consoleempty

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const e = "component-EmptyState"

func textStyle(s, part string) []string {
	var tb []string
	for _, p := range []string{"fontSize", "lineHeight", "fontWeight"} {
		tb = append(tb, e+"-variant-size-"+s+"-anatomy-"+part+"-textStyle-"+p)
	}
	return tb
}

func sized(s, m string) []h.Check {
	return []h.Check{
		{Target: ".cui-console-empty", Props: []string{"padding-top", "padding-bottom"}, TB: []string{e + "-variant-size-" + s + "-paddingBlock", e + "-variant-size-" + s + "-paddingBlock"}},
		{Target: ".cui-console-empty-title", Props: []string{"font-size", "line-height", "font-weight", "max-width", "margin-bottom", "color"}, TB: h.Props(textStyle(s, "Heading"), []string{e + "-maxInlineSize" + m, e + "-variant-size-" + s + "-anatomy-Heading-marginBlockEnd", e + "-anatomy-Heading-color"})},
		{Target: ".cui-console-empty-body", Props: []string{"font-size", "line-height", "font-weight", "max-width", "margin-bottom"}, TB: h.Props(textStyle(s, "Caption"), []string{e + "-maxInlineSize" + m, e + "-variant-size-" + s + "-anatomy-Caption-marginBlockEnd"})},
		{Target: ".cui-console-empty-body", Props: []string{"color"}, Golden: "the secondary label fitted to 0.76 (tokens.css, decision 21)"},
	}
}

var icon = []h.Check{
	{Target: ".cui-console-empty-icon", Props: []string{"box-height", "margin-bottom"}, TB: []string{e + "-variant-size-lg-anatomy-Art-fontSize", e + "-variant-size-lg-anatomy-Art-marginBlockEnd"}},
	{Target: ".cui-console-empty-icon", Props: []string{"color"}, Golden: "the secondary label fitted to 0.76 (tokens.css, decision 21)"},
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "empty small",
			HTML:   h.Example("small"),
			Widths: []int{1200},
			Checks: append(sized("sm", "Sm"),
				h.Check{Target: ".cui-console-empty", Props: []string{"background-color", "border-top-left-radius"}, TB: []string{e + "-variant-backed-on-backgroundColor", e + "-borderRadius"}},
			),
		},
		{
			Name:   "empty large",
			HTML:   h.Example("large"),
			Widths: []int{1200, 390},
			Checks: append(append(sized("lg", "Lg"), icon...),
				h.Check{Target: ".cui-console-empty", Props: []string{"background-color"}, Golden: "a whole page's: no fill, on the page's ground"},
				h.Check{Target: ".cui-console-empty", Props: []string{"flex-direction", "text-align", "align-items"}, Golden: "a column, centered"},
			),
		},
		{
			Name:   "empty dark",
			HTML:   `<div class="cui-dark">` + h.Example("small") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-empty", Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
