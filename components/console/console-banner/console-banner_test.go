//go:build browser

package consolebanner

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	b      = "component-Banner"
	d      = b + "-anatomy"
	banner = ".cui-console-banner"
	link   = "a link in its sentence in the brand's blue fitted to the tints (--cui-c-console-brand-label-on-tint, held to 4.5:1 on each by test/tokens/), not the console link's own"
)

var (
	colors = []string{"info", "success", "error"}
	body   = []string{"typography-textStyle-bodyRegular-fontSize", "typography-textStyle-bodyRegular-lineHeight", "typography-textStyle-bodyRegular-fontWeight"}
)

func colorChecks() []h.Check {
	var checks []h.Check
	for _, c := range colors {
		checks = append(checks, h.Check{Target: banner + `[data-color="` + c + `"]`, Props: []string{"background-color"}, TB: []string{b + "-variant-colorRole-" + c + "-variant-tinted-backgroundColor"}})
	}
	for _, c := range colors {
		checks = append(checks, h.Check{Target: banner + `[data-color="` + c + `"] .cui-console-banner-icon`, Props: []string{"color"}, Golden: "its icon in the fitted " + c + " label (tokens.css, decision 21)"})
	}
	return append(checks,
		h.Check{Target: banner + `[data-color="info"] .cui-console-link`, Props: []string{"color"}, Golden: link},
		h.Check{Target: banner + `[data-color="info"] .cui-console-link`, Props: []string{"text-decoration-line"}, Golden: "a link in its sentence underlined, not told apart by color alone (WCAG 1.4.1)"},
	)
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "banner",
			HTML:   h.Example("banner"),
			Widths: []int{1600, 1200, 640},
			Checks: []h.Check{
				{Target: banner, Props: []string{"border-top-left-radius", "padding-top", "padding-bottom", "padding-left", "padding-right", "min-height"}, TB: []string{b + "-borderRadius", b + "-paddingBlock", b + "-paddingBlock", b + "-paddingInlineStart", b + "-variant-pageScale-desktop-paddingInlineEnd", b + "-variant-pageScale-desktop-minHeight"}},
				{Target: banner, Props: []string{"background-color"}, TB: []string{b + "-variant-colorRole-warning-variant-tinted-backgroundColor"}},
				{Target: ".cui-console-banner-title", Props: []string{"font-size", "line-height", "font-weight", "color"}, TB: []string{d + "-title-variant-pageScale-desktop-textStyle-fontSize", d + "-title-variant-pageScale-desktop-textStyle-lineHeight", d + "-title-variant-pageScale-desktop-textStyle-fontWeight", d + "-title-color"}},
				{Target: ".cui-console-banner-body", Props: []string{"margin-top", "color"}, TB: []string{d + "-body-variant-pageScale-desktop-marginBlockStart", d + "-body-color"}},
				{Target: ".cui-console-banner-body", Props: h.Size, TB: body},
				{Target: ".cui-console-banner-icon", Props: []string{"box-width"}, TB: []string{d + "-art-variant-layout-inline-pageScale-desktop-inlineSize"}},
				{Target: ".cui-console-banner-icon", Props: []string{"box-height", "color"}, Golden: "its icon 20px tall, the toolkit's art size at desktop, in the fitted label of its color (tokens.css, decision 21)"},
				{Target: ".cui-console-banner-actions", Props: []string{"column-gap"}, TB: []string{d + "-actions-variant-pageScale-desktop-gap"}},
				{Target: ".cui-console-banner-actions", Props: []string{"display", "margin-top"}, Golden: "its actions a row of their own, 8px down, as its body"},
			},
		},
		{
			Name:   "banner phone",
			HTML:   h.Example("banner"),
			Widths: []int{639, 390},
			Checks: []h.Check{
				{Target: banner, Props: []string{"padding-right", "min-height"}, TB: []string{b + "-variant-pageScale-mobile-paddingInlineEnd", b + "-variant-pageScale-mobile-minHeight"}},
				{Target: ".cui-console-banner-title", Props: []string{"font-size", "line-height", "font-weight"}, TB: []string{d + "-title-variant-pageScale-mobile-textStyle-fontSize", d + "-title-variant-pageScale-mobile-textStyle-lineHeight", d + "-title-variant-pageScale-mobile-textStyle-fontWeight"}},
				{Target: ".cui-console-banner-body", Props: []string{"margin-top"}, TB: []string{d + "-body-variant-pageScale-mobile-marginBlockStart"}},
				{Target: ".cui-console-banner-body", Props: h.Size, TB: body},
				{Target: ".cui-console-banner-icon", Props: []string{"box-height"}, Golden: "its icon 20px tall below 640px too: the toolkit's 30px is SF Symbols' size, where a Lucide icon at 30px outweighs its 16px title"},
			},
		},
		{
			Name:   "banner colors",
			HTML:   h.Example("colors"),
			Widths: []int{1200, 390},
			Checks: colorChecks(),
		},
		{
			Name:   "banner dark",
			HTML:   `<div class="cui-dark">` + h.Example("banner") + h.Example("colors") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: `[data-color="warning"]`, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: `[data-color="error"]`, Props: []string{"background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: `[data-color="warning"] .cui-console-link`, Props: []string{"color"}, Golden: link},
			},
		},
	})
}
