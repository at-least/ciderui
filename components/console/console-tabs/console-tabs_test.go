//go:build browser

package consoletabs

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	asc "github.com/at-least/ciderui/reference/asc/results"
)

const (
	tb = "component-Tabs"
	// @external: the toolkit's neutral filled tab: the picked one black under white, the
	// others the neutral label on gray; blue stays for the page's primary action
	variant = tb + "-variant-colorRole-neutral-variant"
	rest    = `.cui-console-tab[href="#issues"]`
	radio   = `.cui-console-tabs[aria-label="About"] .cui-console-tab`
	picked  = radio + ":has(> :checked)"
	plain   = radio + ":not(:has(> :checked))"
)

var sizeProps = []string{"box-height", "min-width", "column-gap", "font-size", "line-height", "font-weight"}

func tab(v, effect, active, prop string) string {
	return variant + "-" + v + "-anatomy-Tab-variant-effect-" + effect + "-isActive-" + active + "-" + prop
}

// @external: App Store Connect's version page draws its underlined tabs' labels gray, rgb(51, 51, 51)
// (asc version: tab.button.text), the current one marked by its rule alone: the toolkit's neutral tab
func underlined(effect, active, prop string) string {
	return tb + "-variant-colorRole-neutral-variant-underlined-anatomy-Tab-variant-effect-" + effect + "-isActive-" + active + "-" + prop
}

func size(s string) []string {
	return []string{tb + "-variant-controlSize-" + s + "-anatomy-Tab-height", tb + "-variant-controlSize-" + s + "-anatomy-Tab-minWidth", tb + "-variant-controlSize-" + s + "-anatomy-Tab-gap", tb + "-variant-controlSize-" + s + "-textStyle-fontSize", tb + "-variant-controlSize-" + s + "-textStyle-lineHeight", tb + "-variant-controlSize-" + s + "-textStyle-fontWeight"}
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "tabs",
			HTML:   h.Example("tabs"),
			Widths: []int{1600, 1200, 390},
			Checks: []h.Check{
				{Target: rest, Props: sizeProps, TB: size("md")},
				{Target: rest, Props: []string{"color", "padding-left", "padding-right"}, TB: []string{underlined("default", "false", "color"), tb + "-underlinedPaddingInline", tb + "-underlinedPaddingInline"}},
				{Target: rest, Props: []string{"border-bottom-width", "border-bottom-style", "border-bottom-color"}, ASC: h.ASCProbe(asc.Version, "tab.button")},
				{Target: rest, Props: []string{"color"}, State: []string{"active"}, TB: []string{underlined("press", "false", "color")}},
				{Target: rest, Props: []string{"color"}, State: []string{"hover"}, Golden: "hovered, the gray button's hovered label (tokens.css's neutral label hovered), where the toolkit's tab hovers to oklch(23.8% 0 0 / 0.85): one hovered gray for a label"},
				{Target: ".cui-console-tabs", Props: []string{"column-gap", "border-bottom-width", "border-bottom-style", "border-bottom-color"}, TB: []string{tb + "-variant-controlSize-md-variant-underlined-gap", tb + "-underlinedFullWidthKeyline-width", tb + "-underlinedFullWidthKeyline-style", tb + "-underlinedFullWidthKeyline-color"}},
				{Target: `[aria-current="page"]`, Props: []string{"color", "border-bottom-width"}, TB: []string{underlined("default", "true", "color"), tb + "-underlinedActiveBorderBottomWidth"}},
				{Target: `[aria-current="page"]`, Props: []string{"border-bottom-width", "border-bottom-style"}, ASC: h.ASCProbe(asc.Version, "tab.button.current")},
				{Target: `[aria-current="page"]`, Props: []string{"border-bottom-color"}, Golden: "the current tab's rule in its label's color, as App Store Connect's (asc version: tab.button.current, its label's rgb(51, 51, 51))"},
				{Target: `[aria-current="page"]`, Props: []string{"color"}, State: []string{"hover"}, TB: []string{underlined("hover", "true", "color")}},
				{Target: `[aria-current="page"]`, Props: []string{"color"}, State: []string{"active"}, TB: []string{underlined("press", "true", "color")}},
			},
		},
		{
			Name:   "tabs many on a phone",
			HTML:   h.Example("many"),
			Widths: []int{375},
			Checks: []h.Check{
				{Target: ".cui-console-tabs", Props: []string{"flex-wrap", "overflow-x", "box-height"}, Golden: "a phone: one row (no second hairline), scrolling sideways"},
				{Target: ".cui-console-tabs", Props: []string{"scrollbar-width"}, Golden: "no scrollbar: the tab clipped at the row's edge says more is there"},
				{Target: `.cui-console-tab[href="#settings"]`, Props: []string{"outline-offset"}, State: []string{"focus-visible"}, Golden: "its focus ring inside the tab, where the scroll box would clip it"},
				{Target: ".cui-console-tabs", Props: []string{"scroll-padding-left", "scroll-padding-right"}, Golden: "a focused tab is brought in whole: a browser's focus scrolls only a tab it judges out of view, and left one 3px in view where it was (reference/console/)"},
			},
		},
		{
			Name:   "tabs many wide",
			HTML:   h.Example("many"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-tabs", Props: []string{"flex-wrap", "overflow-x"}, Golden: "wide, the bar as the toolkit draws it"},
			},
		},
		{
			Name:   "tabs not current",
			HTML:   `<div class="cui-console-tabs" aria-label="Report"><a class="cui-console-tab" href="#crashes" aria-current="false">Crashes</a></div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: `[aria-current="false"]`, Props: []string{"border-bottom-color"}, Golden: `aria-current="false" not the current one: no rule`},
			},
		},
		{
			Name:   "tabs of radios",
			HTML:   h.Example("radios"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: plain, Props: sizeProps, TB: size("md")},
				{Target: plain, Props: []string{"background-color", "padding-left", "padding-right", "border-top-left-radius"}, TB: []string{tab("filledGray", "default", "false", "backgroundColor"), tb + "-defaultMdPaddingInline", tb + "-defaultMdPaddingInline", tb + "-variant-controlShape-capsule-controlSize-md-borderRadius"}},
				{Target: plain, Props: []string{"color"}, TB: []string{tab("filledGray", "default", "false", "color")}},
				{Target: plain, Props: []string{"background-color"}, State: []string{"hover"}, TB: []string{tab("filledGray", "hover", "false", "backgroundColor")}},
				{Target: plain, Props: []string{"color"}, State: []string{"hover"}, TB: []string{tab("filledGray", "hover", "false", "color")}},
				{Target: plain, Props: []string{"background-color"}, State: []string{"active"}, TB: []string{tab("filledGray", "press", "false", "backgroundColor")}},
				{Target: plain, Props: []string{"color"}, State: []string{"active"}, TB: []string{tab("filledGray", "press", "false", "color")}},
				{Target: ".cui-console-tabs", Props: []string{"column-gap"}, TB: []string{tb + "-variant-controlSize-md-variant-filledGray-gap"}},
				{Target: ".cui-console-tabs", Props: []string{"border-bottom-width"}, Golden: "no hairline under pills"},
				{Target: plain, Props: []string{"border-bottom-width"}, Golden: "no underline under a pill: an underlined tab's transparent 2px rule lifted its label 1px off its fill's middle"},
				{Target: picked, Props: []string{"background-color", "color"}, TB: []string{tab("filledGray", "default", "true", "backgroundColor"), tab("filledGray", "default", "true", "color")}},
				{Target: picked, Props: []string{"background-color", "color"}, State: []string{"hover"}, TB: []string{tab("filledGray", "hover", "true", "backgroundColor"), tab("filledGray", "hover", "true", "color")}},
				{Target: picked, Props: []string{"background-color", "color"}, State: []string{"active"}, TB: []string{tab("filledGray", "press", "true", "backgroundColor"), tab("filledGray", "press", "true", "color")}},
				{Target: radio + " > input", Props: []string{"position", "opacity", "pointer-events"}, Golden: "the page's own radio out of the way: the pill is the control's look, its words the radio's label"},
				{Target: picked, Props: []string{"outline-style", "outline-width"}, On: radio + " > input", State: []string{"focus-visible"}, Golden: "the keyboard's ring on the pill, not around the hidden radio (the one ring, moved)"},
			},
		},
		{
			Name:         "tabs of radios forced colors",
			HTML:         h.Example("radios"),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: picked, Props: []string{"background-color", "color"}, Golden: "the chosen one the theme's Highlight under HighlightText: forced colors drop the fill, and it read as the others"},
				{Target: plain, Props: []string{"background-color", "color"}, Golden: "the others' gray fill forced to the ground, at its alpha, under the forced text: not Highlight"},
			},
		},
		{
			Name:   "tabs touch",
			HTML:   h.Example("radios"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: ".cui-console-tab", Props: []string{"justify-content"}, Golden: "its label centered in the 44px tab"},
				{Target: plain, Props: []string{"box-height", "min-width", "border-top-left-radius"}, Golden: "a tab 44px each way at least on a touch screen, the HIG's 44pt (Accessibility, Mobility), a pill's corners half its height: a target drawn round it was clipped by a phone's sideways-scrolling row"},
			},
		},
		{
			Name:   "tabs underlined touch",
			HTML:   h.Example("many"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: rest, Props: []string{"box-height", "min-width"}, Golden: "an underlined tab 44px each way at least on a touch screen, in its scrolling row"},
			},
		},
		{
			Name:         "tabs forced colors",
			HTML:         h.Example("tabs"),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: `.cui-console-tab[aria-current="page"]`, Props: []string{"border-bottom-color", "border-bottom-width"}, Golden: "the current tab's rule, in the forced text color"},
				{Target: rest, Props: []string{"border-bottom-color", "border-bottom-width"}, Golden: "the others' rule in the ground's color, its 2px kept: forced colors drew even the transparent rule in the text's color, every tab ruled as the current one"},
			},
		},
		{
			Name:   "tabs dark",
			HTML:   `<div class="cui-dark">` + h.Example("tabs") + h.Example("radios") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: rest, Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: plain, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: picked, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18): the chosen pill white under black"},
			},
		},
	})
}
