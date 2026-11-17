//go:build browser

package consolecode

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const code = "p:first-child .cui-console-code"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "code",
			HTML:   h.Example("code"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: code, Props: []string{"font-family"}, Golden: `ciderui's mono stack: the toolkit's "SF Mono" names a font only Apple's platforms have`},
				{Target: code, Props: []string{"font-size", "line-height", "color", "letter-spacing"}, Golden: "mono at the text's own size, line and color, its tracking none, as cui-console-prose's code (DocC's inline code)"},
				{Target: code, Props: []string{"padding-left", "border-top-left-radius", "background-color"}, Golden: "no chip: no padding, corners or fill, as cui-console-prose's code (DocC's inline code)"},
			},
		},
		{
			Name:   "code dark",
			HTML:   `<div class="cui-dark">` + h.Example("code") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: code, Props: []string{"color"}, Golden: "the text's own color, dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
