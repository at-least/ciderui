//go:build browser

package consolepage

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	asc "github.com/at-least/ciderui/reference/asc/results"
)

const stack = "tokens.css's stack, SF on Apple's platforms and Inter elsewhere, where the toolkit names -system-ui first (decision 4)"

var (
	typeProps = []string{"font-size", "line-height", "font-weight", "letter-spacing"}
	desktop   = []int{1600, 1200, 900, 640}
	styles    = [][2]string{
		{"h1-plus", "h1PlusRegular"},
		{"h1-plus-emph", "h1PlusEmph"},
		{"h1", "h1Regular"},
		{"h1-emph", "h1Emph"},
		{"h2", "h2Emph"},
		{"h3", "h3Regular"},
		{"h3-emph", "h3Emph"},
		{"h4", "h4Regular"},
		{"h4-emph", "h4Emph"},
		{"h5", "h5Regular"},
		{"h5-emph", "h5Emph"},
		{"h6", "h6Regular"},
		{"h6-emph", "h6Emph"},
		{"body", "bodyRegular"},
		{"body-emph", "bodyEmph"},
		{"small", "bodySmallRegular"},
		{"small-emph", "bodySmallEmph"},
		{"xs", "bodyXsRegular"},
		{"xs-emph", "bodyXsEmph"},
	}
)

func style(s string) []string {
	var tb []string
	for _, p := range []string{"fontSize", "lineHeight", "fontWeight", "letterSpacing"} {
		tb = append(tb, "typography-textStyle-"+s+"-"+p)
	}
	return tb
}

func typeChecks() []h.Check {
	var checks []h.Check
	for _, st := range styles {
		checks = append(checks, h.Check{Target: ".cui-console-" + st[0], Props: typeProps, TB: style(st[1])})
	}
	return append(checks,
		h.Check{Target: ".cui-console-mono", Props: []string{"font-size", "line-height", "font-weight", "letter-spacing"}, TB: style("monoBodyRegular")},
		h.Check{Target: ".cui-console-mono", Props: []string{"font-family"}, Golden: "tokens.css's code stack, SF Mono on Apple's platforms and the platform's own elsewhere, where the toolkit names SF Mono alone"},
	)
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "page",
			HTML:   h.Example("page"),
			Widths: desktop,
			Checks: []h.Check{
				{Target: ".cui-console-page", Props: []string{"color", "background-color"}, TB: []string{"color-role-neutral-label-primary", "color-role-neutral-background-primary"}},
				{Target: ".cui-console-page", Props: typeProps, TB: style("bodyRegular")},
				{Target: ".cui-console-page", Props: []string{"font-family"}, Golden: stack},
				{Target: ".cui-console-h4-emph", Props: []string{"font-size", "line-height", "font-weight", "letter-spacing", "color"}, ASC: h.ASCProbe(asc.AnalyticsPage, "section.heading")},
			},
		},
		{
			Name:   "page phone",
			HTML:   h.Example("page"),
			Widths: []int{639, 390},
			Checks: []h.Check{
				{Target: ".cui-console-page", Props: typeProps, TB: style("bodyMobileRegular")},
			},
		},
		{
			Name:   "page dark",
			HTML:   `<div class="cui-dark">` + h.Example("page") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-page", Props: []string{"color", "background-color"}, Golden: "the dark side: black and white each other's (tokens.css, decision 18)"},
			},
		},
		{
			Name:   "type",
			HTML:   h.Example("type"),
			Widths: []int{1200, 390},
			Checks: typeChecks(),
		},
		{
			Name:   "type composes",
			HTML:   `<div class="cui-console-page"><p id="variant" class="cui-console-body md:cui-console-h1">v</p><p id="plain" class="cui-console-h1 text-[20px]">p</p></div>`,
			Widths: []int{900, 390},
			Checks: []h.Check{
				{Target: "#variant", Props: []string{"font-size", "line-height"}, Golden: "utilities, not component classes, as type.css's: a variant composes with one (md:cui-console-h1 over cui-console-body from 735px)"},
				{Target: "#plain", Props: []string{"font-size"}, Golden: "a plain utility wins over one"},
			},
		},
	})
}
