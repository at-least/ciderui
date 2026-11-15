//go:build browser

package consolecode

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	textStyle = "typography-textStyle-monoBodyRegular"
	code      = "p:first-child .cui-console-code"
)

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "code",
			HTML:   h.Example("code"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: code, Props: []string{"font-size", "line-height", "font-weight", "letter-spacing"}, TB: []string{textStyle + "-fontSize", textStyle + "-lineHeight", textStyle + "-fontWeight", textStyle + "-letterSpacing"}},
				{Target: code, Props: []string{"font-family"}, Golden: `ciderui's mono stack: the toolkit's "SF Mono" names a font only Apple's platforms have`},
				{Target: code, Props: []string{"color", "background-color"}, Golden: "the label's color on the gray badge's fill (the neutral fill at 0.12): the text style has no ground"},
				{Target: code, Props: []string{"padding-left", "padding-right", "border-top-left-radius"}, Golden: "4px in, 4px corners: the toolkit's smallest rounded rectangle, 8px, rounds an 18px line nearly to a capsule"},
			},
		},
		{
			Name:   "code dark",
			HTML:   `<div class="cui-dark">` + h.Example("code") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: code, Props: []string{"color", "background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
