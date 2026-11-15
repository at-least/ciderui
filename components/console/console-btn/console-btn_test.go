//go:build browser

package consolebtn

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	asc "github.com/at-least/ciderui/reference/asc/results"
)

const (
	b         = "component-Button-variant"
	gray      = ".flex > :nth-child(1)"
	filled    = ".flex > :nth-child(2)"
	plain     = ".flex > :nth-child(3)"
	red       = ".flex > :nth-child(4)"
	redFilled = ".flex > :nth-child(5)"
	redPlain  = ".flex > :nth-child(6)"
	fitted    = "fitted to 4.5:1 (tokens.css, decision 21)"
)

var (
	desktop   = []int{1600, 1200, 390}
	sizeProps = []string{"box-height", "padding-left", "padding-right", "border-top-left-radius", "column-gap", "min-width"}
)

func at(size string) []string {
	return []string{b + "-controlSize-" + size + "-blockSize", b + "-controlSize-" + size + "-paddingInlineStart", b + "-controlSize-" + size + "-paddingInlineEnd", b + "-controlShape-roundedRect-controlSize-" + size + "-borderRadius", b + "-controlSize-" + size + "-gap", b + "-controlSize-" + size + "-minWidth"}
}

func iconSizes() []h.Check {
	var checks []h.Check
	for _, size := range []string{"sm", "lg"} {
		checks = append(checks, h.Check{Target: `[data-size="` + size + `"]`, Props: []string{"font-size"}, TB: []string{"component-Button-anatomy-artStart-withoutLabel-controlSize-" + size + "-fontSize"}})
	}
	return checks
}

func capsule() []h.Check {
	var checks []h.Check
	for _, size := range []string{"sm", "md", "lg"} {
		target := `.cui-console-btn-icon[data-size="` + size + `"]`
		if size == "md" {
			target = ".cui-console-btn-icon:not([data-size])"
		}
		checks = append(checks, h.Check{Target: target, Props: []string{"box-height", "box-width", "border-top-left-radius", "border-bottom-right-radius"}, TB: []string{b + "-controlSize-" + size + "-blockSize", b + "-controlSize-" + size + "-minWidth", b + "-controlShape-capsule-controlSize-" + size + "-borderRadius", b + "-controlShape-capsule-controlSize-" + size + "-borderRadius"}})
	}
	return append(checks, h.Check{Target: ".cui-console-btn:not(.cui-console-btn-icon)", Props: []string{"box-height", "border-top-left-radius", "padding-left", "padding-right"}, TB: []string{b + "-controlSize-md-blockSize", b + "-controlShape-capsule-controlSize-md-borderRadius", b + "-controlSize-md-paddingInlineStart", b + "-controlSize-md-paddingInlineEnd"}})
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "variants",
			HTML:   h.Example("variants"),
			Widths: desktop,
			Checks: []h.Check{
				{Target: gray, Props: []string{"box-height", "font-size", "line-height", "font-weight", "letter-spacing", "color", "background-color", "padding-left", "padding-right", "border-top-left-radius", "border-top-width", "min-width", "column-gap", "text-align", "text-decoration-line", "cursor", "transition-duration", "box-shadow"}, ASC: h.ASCProbe(asc.Button, "button")},
				{Target: gray, Props: []string{"background-color", "color", "cursor"}, State: []string{"hover"}, ASC: h.ASCProbe(asc.HoverButton, "button")},
				{Target: gray, Props: []string{"background-color", "color"}, State: []string{"active"}, TB: []string{b + "-colorRole-neutral-effect-press-variant-gray-backgroundColor", b + "-colorRole-neutral-effect-press-variant-gray-color"}},
				{Target: filled, Props: []string{"background-color", "color"}, TB: []string{b + "-colorRole-brand-effect-default-variant-filled-backgroundColor", b + "-colorRole-brand-effect-default-variant-filled-color"}},
				{Target: filled, Props: []string{"background-color"}, State: []string{"hover"}, Golden: "the filled blue hovered a step darker, blue-600: white on the toolkit's lighter blue is 3.9:1, " + fitted},
				{Target: filled, Props: []string{"background-color", "color"}, State: []string{"active"}, TB: []string{b + "-colorRole-brand-effect-press-variant-filled-backgroundColor", b + "-colorRole-brand-effect-press-variant-filled-color"}},
				{Target: plain, Props: []string{"background-color"}, Golden: "no fill: transparent, where the toolkit writes its transparent white, the same nothing"},
				{Target: plain, Props: []string{"color", "padding-left", "padding-right", "min-width"}, TB: []string{b + "-colorRole-brand-effect-default-variant-plain-color", b + "-variant-plain-paddingInlineStart", b + "-variant-plain-paddingInlineEnd", "component-Button-anatomy-label-variant-variant-plain-minWidth"}},
				{Target: plain, Props: []string{"color"}, State: []string{"hover"}, Golden: "the brand's label hovered a step darker, blue-600: the toolkit's lighter blue is 3.5:1 on white, " + fitted},
				{Target: plain, Props: []string{"color"}, State: []string{"active"}, TB: []string{b + "-colorRole-brand-effect-press-variant-plain-color"}},
				{Target: red, Props: []string{"background-color"}, TB: []string{b + "-colorRole-error-effect-default-variant-gray-backgroundColor"}},
				{Target: red, Props: []string{"color"}, Golden: "the red label red-700: the toolkit's is 3.6:1 on white, " + fitted},
				{Target: red, Props: []string{"color"}, State: []string{"hover"}, Golden: "red-700 hovered too: the toolkit's lighter red is 2.9:1 on white, " + fitted},
				{Target: red, Props: []string{"background-color"}, State: []string{"hover"}, TB: []string{b + "-colorRole-error-effect-hover-variant-gray-backgroundColor"}},
				{Target: red, Props: []string{"background-color"}, State: []string{"active"}, TB: []string{b + "-colorRole-error-effect-press-variant-gray-backgroundColor"}},
				{Target: red, Props: []string{"color"}, State: []string{"active"}, Golden: "red-700 pressed too: the toolkit's is lighter, " + fitted},
				{Target: redFilled, Props: []string{"background-color", "color"}, Golden: "the red fill red-700: white on the toolkit's red is 3.3:1, " + fitted},
				{Target: redFilled, Props: []string{"background-color"}, State: []string{"hover"}, Golden: "hovered red-800, a step darker: " + fitted},
				{Target: redFilled, Props: []string{"color"}, State: []string{"hover"}, TB: []string{b + "-colorRole-error-effect-hover-variant-filled-color"}},
				{Target: redFilled, Props: []string{"background-color"}, State: []string{"active"}, Golden: "pressed red-800 too: the toolkit's is lighter, " + fitted},
				{Target: redFilled, Props: []string{"color"}, State: []string{"active"}, TB: []string{b + "-colorRole-error-effect-press-variant-filled-color"}},
				{Target: redPlain, Props: []string{"color", "background-color"}, Golden: "the red label red-700, no fill: " + fitted},
				{Target: redPlain, Props: []string{"color", "background-color"}, State: []string{"hover"}, Golden: "hovered, red-700, no fill: " + fitted},
				{Target: redPlain, Props: []string{"color", "background-color"}, State: []string{"active"}, Golden: "pressed, red-700, no fill: " + fitted},
			},
		},
		{
			Name:   "alone",
			HTML:   `<p><button class="cui-console-btn" type="button">Compare</button></p>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-btn", Props: []string{"display", "vertical-align", "white-space"}, ASC: h.ASCProbe(asc.Button, "button")},
			},
		},
		{
			Name:   "in a row",
			HTML:   `<div style="display: flex; width: 300px"><input style="flex: 1"><button class="cui-console-btn" type="button">Generate Again</button></div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-btn", Props: []string{"flex-shrink", "box-width", "box-height", "white-space"}, Golden: "a control keeps its size beside a field that fills the row: as wide as its label, which stays on one line in its 30px (squashed, the label wrapped out of it; muse wrote shrink-0 on seven buttons); its white-space App Store Connect's normal"},
			},
		},
		{
			Name:   "disabled",
			HTML:   h.Example("disabled"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: gray, Props: []string{"background-color", "color"}, TB: []string{b + "-colorRole-neutral-effect-disabled-variant-gray-backgroundColor", b + "-colorRole-neutral-effect-disabled-variant-gray-color"}},
				{Target: filled, Props: []string{"background-color", "color"}, TB: []string{b + "-colorRole-brand-effect-disabled-variant-filled-backgroundColor", b + "-colorRole-brand-effect-disabled-variant-filled-color"}},
				{Target: plain, Props: []string{"color"}, TB: []string{b + "-colorRole-brand-effect-disabled-variant-plain-color"}},
				{Target: red, Props: []string{"color"}, TB: []string{b + "-colorRole-error-effect-disabled-variant-gray-color"}},
				{Target: gray, Props: []string{"cursor"}, Golden: "no pointer over a disabled button"},
				{Target: gray, Props: []string{"background-color", "color"}, State: []string{"hover"}, Golden: "hovered, a disabled button stays as it is"},
			},
		},
		{
			Name:   "aria-disabled",
			HTML:   `<a class="cui-console-btn" aria-disabled="true" href="#all">See All</a>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-btn", Props: []string{"background-color", "color"}, TB: []string{b + "-colorRole-neutral-effect-disabled-variant-gray-backgroundColor", b + "-colorRole-neutral-effect-disabled-variant-gray-color"}},
				{Target: ".cui-console-btn", Props: []string{"cursor"}, Golden: "a link marked disabled: no pointer, as a disabled button"},
			},
		},
		{
			Name:   "sizes",
			HTML:   h.Example("sizes"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: `[data-size="sm"]`, Props: sizeProps, TB: at("sm")},
				{Target: ".flex > :nth-child(2)", Props: sizeProps, TB: at("md")},
				{Target: `[data-size="lg"]`, Props: sizeProps, TB: at("lg")},
				{Target: `[data-size="sm"]`, Props: []string{"font-size", "line-height", "font-weight"}, TB: []string{"component-Tabs-variant-controlSize-sm-textStyle-fontSize", "component-Tabs-variant-controlSize-sm-textStyle-lineHeight", "component-Tabs-variant-controlSize-sm-textStyle-fontWeight"}},
				{Target: `[data-size="lg"]`, Props: []string{"font-size", "line-height", "font-weight"}, TB: []string{"component-Tabs-variant-controlSize-lg-textStyle-fontSize", "component-Tabs-variant-controlSize-lg-textStyle-lineHeight", "component-Tabs-variant-controlSize-lg-textStyle-fontWeight"}},
			},
		},
		{
			Name:   "icon",
			HTML:   h.Example("icon"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-btn-icon", Props: []string{"box-height", "box-width", "font-size"}, TB: []string{b + "-controlSize-md-blockSize", b + "-controlSize-md-minWidth", "component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize"}},
				{Target: ".cui-console-btn-icon svg", Props: []string{"box-height", "box-width"}, TB: []string{"component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize", "component-Button-anatomy-artStart-withoutLabel-controlSize-md-fontSize"}},
				{Target: ".cui-console-btn:not(.cui-console-btn-icon) svg", Props: []string{"box-height", "box-width"}, Golden: "an icon beside a label at the label's size, 1em"},
			},
		},
		{
			Name:   "icon sizes",
			HTML:   `<div class="flex gap-3"><button class="cui-console-btn cui-console-btn-icon" data-size="sm" type="button" aria-label="Copy"></button><button class="cui-console-btn cui-console-btn-icon" data-size="lg" type="button" aria-label="Copy"></button></div>`,
			Widths: []int{1200},
			Checks: iconSizes(),
		},
		{
			Name:   "capsule",
			HTML:   h.Example("capsule"),
			Widths: []int{1200},
			Checks: capsule(),
		},
		{
			Name:   "capsule touch",
			HTML:   h.Example("capsule"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: `.cui-console-btn-icon[data-size="sm"]`, Props: []string{"box-width", "box-height", "border-top-left-radius"}, Golden: "a capsule icon button 44px square on a touch screen, its corners half that: a circle still, as the toolkit's capsule is at each size"},
			},
		},
		{
			Name:   "touch target",
			HTML:   h.Example("icon"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: ".cui-console-btn-icon", Props: []string{"box-width", "box-height"}, Golden: "an icon button 44px square on a touch screen, the HIG's 44pt (Accessibility, Mobility): a target drawn round a 30px button covered its neighbour's, 8px away"},
				{Target: ".cui-console-btn:not(.cui-console-btn-icon)", Props: []string{"box-height"}, Golden: "a labelled button 44px tall on a touch screen, the HIG's 44pt (Accessibility, Mobility)"},
			},
		},
		{
			Name:   "touch sizes",
			HTML:   h.Example("sizes"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: `[data-size="sm"]`, Props: []string{"box-height"}, Golden: "the small button 44px tall on a touch screen too, the HIG's 44pt (Accessibility, Mobility)"},
			},
		},
		{
			Name:   "dark",
			HTML:   `<div class="cui-dark">` + h.Example("variants") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: gray, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: filled, Props: []string{"background-color", "color"}, Golden: "the dark blue fill dark blue-400: white on the derived blue-600 is 2.7:1, " + fitted},
				{Target: plain, Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: red, Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: redFilled, Props: []string{"background-color"}, Golden: "the dark red fill dark red-300: " + fitted},
			},
		},
		{
			Name:          "busy",
			HTML:          h.Example("busy"),
			Widths:        []int{1200, 390},
			ReducedMotion: true,
			Checks: []h.Check{
				{Target: ".cui-console-btn", Props: []string{"cursor"}, Golden: "aria-busy, set by the page's script while the request runs: the pointer says so"},
				{Target: ".cui-console-btn", Pseudo: "::before", Props: []string{"width", "height", "border-top-width", "border-top-color", "border-right-color", "border-top-left-radius", "animation-name", "animation-duration", "animation-iteration-count"}, Golden: "a spinner before the label, at the label's size, a ring open on one side so its turn shows, the label kept so the button keeps its width and its name; under Reduce Motion it keeps turning, as iOS's activity indicator does (base.css)"},
			},
		},
	})
}
