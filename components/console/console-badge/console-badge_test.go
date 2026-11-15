//go:build browser

package consolebadge

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	b      = "component-Badge-variant"
	fitted = "its label fitted to 4.5:1 on the page, its tint and the gray (tokens.css, decision 21)"
)

var (
	colorRoles = []string{"neutral", "info", "success", "warning", "error"}
	md         = []string{b + "-size-md-height", b + "-size-md-paddingInline", b + "-size-md-paddingInline", b + "-controlShape-capsule-size-md-borderRadius", b + "-size-md-textStyle-fontSize", b + "-size-md-textStyle-lineHeight", b + "-size-md-textStyle-fontWeight"}
	box        = []string{"box-height", "padding-left", "padding-right", "border-top-left-radius", "font-size", "line-height", "font-weight"}
)

func at(c string) string {
	if c == "neutral" {
		return ".cui-console-badge:not([data-color])"
	}
	return `.cui-console-badge[data-color="` + c + `"]`
}

func TestCases(t *testing.T) {
	badges := []h.Check{{Target: at("neutral"), Props: box, TB: md}}
	for _, c := range colorRoles {
		badges = append(badges, h.Check{Target: at(c), Props: []string{"background-color"}, TB: []string{b + "-colorRole-" + c + "-variant-tinted-backgroundColor"}})
	}
	badges = append(badges,
		h.Check{Target: at("neutral"), Props: []string{"color"}, TB: []string{b + "-colorRole-neutral-variant-tinted-color"}},
		h.Check{Target: at("error"), Props: []string{"color"}, Golden: "red-700: " + fitted},
	)
	for _, c := range []string{"info", "success", "warning"} {
		badges = append(badges, h.Check{Target: at(c), Props: []string{"color"}, Golden: c + ": " + fitted})
	}
	var dark []h.Check
	for _, c := range colorRoles {
		dark = append(dark, h.Check{Target: at(c), Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"})
	}
	h.Run(t, []h.Case{
		{
			Name:   "badges",
			HTML:   h.Example("badges"),
			Widths: []int{1200, 390},
			Checks: badges,
		},
		{
			Name:   "badges dark",
			HTML:   `<div class="cui-dark">` + h.Example("badges") + `</div>`,
			Widths: []int{1200},
			Checks: dark,
		},
	})
}
