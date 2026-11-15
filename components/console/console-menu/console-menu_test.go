//go:build browser

package consolemenu

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	menu      = "component-DisplayList-variant-pageScale"
	fitted    = "the secondary label fitted to 0.76 (tokens.css, decision 21): the toolkit's 0.60 is 3.4:1 on white"
	current   = `.cui-console-menu > [aria-current="page"]`
	sub       = `.cui-console-menu > a[href="#sources"]`
	plain     = `.cui-console-menu > a[href="#shop"]`
	open      = ".cui-console-btn[popovertarget]"
	checked   = `[role="menuitemradio"][aria-checked="true"]`
	unchecked = `[role="menuitemradio"][aria-checked="false"]`
)

var (
	desktop = []int{1600, 1200, 640}
	phone   = []int{639, 390}
	item    = []string{"box-height", "background-color", "border-top-left-radius", "padding-left", "padding-right"}
	label   = []string{"font-size", "line-height", "font-weight", "letter-spacing", "color"}
)

func page(html string) string { return `<div style="padding:24px 24px 320px">` + html + `</div>` }

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "menu long label",
			HTML:   `<div class="w-48"><nav class="cui-console-menu" aria-label="Analytics"><a class="cui-console-menu-item" href="#a"><span class="cui-console-menu-label">App Store Product Page Optimization</span></a></nav></div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-menu-label", Props: []string{"box-height"}, Golden: "a long label on two lines, whole, not cut with an ellipsis"},
			},
		},
		{
			Name:   "menu choices",
			HTML:   page(h.Example("choices")),
			Widths: []int{1200},
			Click:  []string{open},
			Checks: []h.Check{
				{Target: "#subtitles", Props: []string{"display"}, Golden: "its button opens it, a console popover (decision 27)"},
				{Target: unchecked, Props: []string{"box-height", "border-top-left-radius", "padding-left", "padding-right"}, ASC: h.ASCProbe("hover-menu", "menu.item")},
				{Target: unchecked + " .cui-console-menu-label", Props: label, ASC: h.ASCProbe("analytics", "menu.item.label")},
				{Target: unchecked, Props: []string{"background-color", "border-top-width", "text-align", "font-family"}, Golden: "a <button> drawn as the menu's <a>: no ground, no edge, its label at its start in the page's face"},
				{Target: checked, Props: []string{"background-color"}, Golden: "the chosen one unfilled, its check alone marking it, as tv.apple.com's chosen subtitles (tv-player-menu:chosen, no fill): a fill is the current page's"},
				{Target: unchecked, Props: []string{"background-color"}, State: []string{"hover"}, ASC: h.ASCProbe("hover-menu", "menu.item")},
				{Target: unchecked, Props: []string{"box-width"}, Golden: "a row the menu's width, as an <a> item is"},
				{Target: checked + " .cui-console-menu-icon", Props: []string{"box-width", "margin-left"}, TB: []string{"component-DisplayList-anatomy-Item-anatomy-Indicator-variant-inlinePosition-start-width", "component-DisplayList-anatomy-Item-anatomy-Indicator-variant-inlinePosition-start-marginInlineStart"}},
				{Target: checked + " .cui-console-menu-icon", Props: []string{"visibility", "color"}, Golden: "the check in the toolkit's indicator's place at its start, where tv.apple.com draws its own (tv-player-menu:chosen.check, 10px in); Lucide's check (decision 17); " + fitted},
				{Target: unchecked + " .cui-console-menu-icon", Props: []string{"visibility", "box-width"}, Golden: "unchecked, its check hidden in its place: every label lines up, as tv.apple.com's (its text 36px in on every row)"},
				{Target: unchecked + " .cui-console-menu-label", Props: []string{"box-left"}, Golden: "each label at the same start"},
				{Target: checked + " .cui-console-menu-label", Props: []string{"box-left"}, Golden: "each label at the same start"},
				{Target: "#subtitles .cui-console-menu-title", Props: []string{"font-size", "line-height", "font-weight", "padding-top", "padding-left"}, TB: []string{menu + "-desktop-variant-menu-anatomy-Heading-textStyle-fontSize", menu + "-desktop-variant-menu-anatomy-Heading-textStyle-lineHeight", menu + "-desktop-variant-menu-anatomy-Heading-textStyle-fontWeight", menu + "-desktop-variant-menu-anatomy-Header-paddingBlockStart", menu + "-desktop-variant-menu-anatomy-Header-paddingInline"}},
			},
		},
		{
			Name:   "menu choices columns",
			HTML:   page(h.Example("columns")),
			Widths: []int{1200},
			Click:  []string{open},
			Checks: []h.Check{
				{Target: "#subtitles .cui-console-menu-columns", Props: []string{"display", "align-items"}, Golden: "the columns side by side, the Apple TV app's Audio beside its Subtitles (the owner, 2026-11-05, for muse's player: one button, one panel), each as tall as the menu so the rule runs its height"},
				{Target: `#subtitles [data-column="voices"]`, Props: []string{"box-x", "box-width"}, Golden: "the first column at the popover's 16px side, as wide as its choices ask"},
				{Target: `#subtitles [data-column="subtitles"]`, Props: []string{"box-x", "border-left-width", "border-left-style", "border-left-color", "margin-left", "padding-left"}, Golden: "the next after a rule, the menu's own 0.5px in its border color (cui-console-menu-rule), in the middle of the popover's 8px gap (cui-console-popover)"},
				{Target: "#subtitles", Props: []string{"box-width", "max-width"}, Golden: "the popover as wide as its columns ask, past the one column's 350px"},
				{Target: `[data-column="subtitles"] ` + unchecked + " .cui-console-menu-label", Props: label, ASC: h.ASCProbe("analytics", "menu.item.label")},
				{Target: `[data-column="voices"] ` + checked + " .cui-console-menu-icon", Props: []string{"visibility"}, Golden: "each column its own check"},
			},
		},
		{
			Name:   "menu choices columns phone",
			HTML:   page(h.Example("columns")),
			Widths: []int{390},
			Click:  []string{open},
			Checks: []h.Check{
				{Target: "#subtitles", Props: []string{"box-width"}, Golden: "no wider than 85% of a phone's window, the popover's own cap"},
				{Target: `#subtitles [data-column="subtitles"]`, Props: []string{"box-right"}, Golden: "both columns inside it"},
			},
		},
		{
			Name:   "menu choices chosen",
			HTML:   page(h.Example("choices")),
			Widths: []int{1200},
			Click:  []string{open, `[role="menuitemradio"][value="ja"]`},
			Checks: []h.Check{
				{Target: "#subtitles", Props: []string{"display"}, Golden: `a choice closes its popover with no script (popovertargetaction="hide"); its check is the page's script's to move`},
			},
		},
		{
			Name:   "menu choices phone",
			HTML:   page(h.Example("choices")),
			Widths: []int{390},
			Click:  []string{open},
			Checks: []h.Check{
				{Target: unchecked, Props: []string{"padding-top", "padding-left", "border-top-left-radius", "column-gap"}, TB: []string{menu + "-mobile-variant-menu-anatomy-Item-paddingBlock", menu + "-mobile-variant-menu-anatomy-Item-paddingInline", menu + "-mobile-variant-menu-anatomy-Item-borderRadius", menu + "-mobile-variant-menu-anatomy-Item-gap"}},
			},
		},
		{
			Name:   "menu choices dark",
			HTML:   `<div class="cui-dark">` + page(h.Example("choices")) + `</div>`,
			Widths: []int{1200},
			Click:  []string{open},
			Checks: []h.Check{
				{Target: "#subtitles", Props: []string{"background-color", "color"}, Golden: "dark over a dark page: the console popover's derived dark (decision 18), a menu over the video in muse's dark console"},
				{Target: checked + " .cui-console-menu-icon", Props: []string{"color"}, Golden: "the check in the dark secondary label"},
			},
		},
		{
			Name:   "menu",
			HTML:   h.Example("menu"),
			Widths: desktop,
			Checks: []h.Check{
				{Target: current, Props: item, ASC: h.ASCProbe("analytics", "menu.current")},
				{Target: current + " .cui-console-menu-label", Props: label, ASC: h.ASCProbe("analytics", "menu.current.label")},
				{Target: sub, Props: []string{"box-height", "border-top-left-radius", "padding-right"}, ASC: h.ASCProbe("analytics", "menu.item")},
				{Target: sub, Props: []string{"background-color"}, Golden: "no ground at rest: transparent, where the toolkit writes its transparent white, oklch(1 0 0 / 0), the same nothing"},
				{Target: sub + " .cui-console-menu-label", Props: label, ASC: h.ASCProbe("analytics", "menu.item.label")},
				{Target: ".cui-console-menu-rule", Props: []string{"box-height", "background-color", "margin-top", "margin-bottom", "margin-left", "margin-right"}, ASC: h.ASCProbe("analytics", "menu.rule")},
				{Target: plain, Props: []string{"background-color"}, State: []string{"hover"}, ASC: h.ASCProbe("hover-menu", "menu.item")},
				{Target: plain + " .cui-console-menu-label", Props: []string{"color", "font-weight"}, State: []string{"hover"}, On: plain, ASC: h.ASCProbe("hover-menu", "menu.item.label")},
				{Target: ".cui-console-menu-title", Props: []string{"font-size", "line-height", "font-weight", "padding-top", "padding-bottom", "padding-left"}, TB: []string{menu + "-desktop-variant-menu-anatomy-Heading-textStyle-fontSize", menu + "-desktop-variant-menu-anatomy-Heading-textStyle-lineHeight", menu + "-desktop-variant-menu-anatomy-Heading-textStyle-fontWeight", menu + "-desktop-variant-menu-anatomy-Header-paddingBlockStart", menu + "-desktop-variant-menu-anatomy-Header-paddingBlockEnd", menu + "-desktop-variant-menu-anatomy-Header-paddingInline"}},
				{Target: current, Props: []string{"column-gap"}, TB: []string{menu + "-desktop-variant-menu-anatomy-Item-gap"}},
				{Target: current + " .cui-console-menu-icon", Props: []string{"box-width", "margin-left"}, TB: []string{"component-DisplayList-variant-variant-menu-anatomy-artStackWithSFSymbol-width", "component-DisplayList-variant-variant-menu-anatomy-ArtStack-anatomy-SFSymbol-marginInlineOffset"}},
				{Target: current + " .cui-console-menu-icon", Props: []string{"color"}, Golden: fitted},
				{Target: ".cui-console-menu-title", Props: []string{"color"}, Golden: fitted},
				{Target: plain, Props: []string{"background-color"}, State: []string{"active"}, Golden: "pressed as hovered: the toolkit's pressed fill, oklch(55.4% 0.012 286 / 0.08), is no step of its palette, so it has no dark side to derive"},
				{Target: ".cui-console-menu", Props: []string{"display"}, ASC: h.ASCProbe("analytics-page", "menu")},
				{Target: ".cui-console-menu", Props: []string{"flex-direction"}, Golden: "a column of its items"},
				{Target: sub + " .cui-console-menu-label", Props: []string{"white-space", "text-overflow", "overflow-wrap"}, Golden: "a long label wraps, whole: an ellipsis left two pages that begin alike unreadable apart"},
			},
		},
		{
			Name:   "menu phone",
			HTML:   h.Example("menu"),
			Widths: phone,
			Checks: []h.Check{
				{Target: current, Props: []string{"padding-top", "padding-left", "border-top-left-radius", "column-gap"}, TB: []string{menu + "-mobile-variant-menu-anatomy-Item-paddingBlock", menu + "-mobile-variant-menu-anatomy-Item-paddingInline", menu + "-mobile-variant-menu-anatomy-Item-borderRadius", menu + "-mobile-variant-menu-anatomy-Item-gap"}},
				{Target: current + " .cui-console-menu-label", Props: []string{"font-size", "line-height", "font-weight"}, TB: []string{menu + "-mobile-variant-menu-anatomy-Label-textStyle-fontSize", menu + "-mobile-variant-menu-anatomy-Label-textStyle-lineHeight", menu + "-mobile-variant-menu-anatomy-Label-textStyle-fontWeight"}},
				{Target: ".cui-console-menu-title", Props: []string{"font-size", "line-height", "font-weight", "padding-top", "padding-bottom", "padding-left"}, TB: []string{menu + "-mobile-variant-menu-anatomy-Heading-textStyle-fontSize", menu + "-mobile-variant-menu-anatomy-Heading-textStyle-lineHeight", menu + "-mobile-variant-menu-anatomy-Heading-textStyle-fontWeight", menu + "-mobile-variant-menu-anatomy-Header-paddingBlockStart", menu + "-mobile-variant-menu-anatomy-Header-paddingBlockEnd", menu + "-mobile-variant-menu-anatomy-Header-paddingInline"}},
				{Target: ".cui-console-menu-rule", Props: []string{"margin-left", "margin-right"}, TB: []string{menu + "-mobile-variant-menu-anatomy-HR-marginInline", menu + "-mobile-variant-menu-anatomy-HR-marginInline"}},
				{Target: current, Props: []string{"box-height"}, Golden: "the toolkit's mobile item, its padding round its line"},
				{Target: sub, Props: []string{"padding-left"}, Golden: "a group's items 48px in on a phone too: the toolkit names no mobile value, the analytics page's is read at 1920px"},
			},
		},
		{
			Name:   "menu touch",
			HTML:   h.Example("menu"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: plain, Props: []string{"box-height", "border-top-left-radius"}, Golden: "an item 44px tall on a touch screen, the HIG's 44pt (Accessibility, Mobility), a pill's corners half of it"},
			},
		},
		{
			Name:         "menu forced colors",
			HTML:         h.Example("menu"),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: current, Props: []string{"background-color", "color"}, Golden: "the current page the theme's Highlight under HighlightText: forced colors drop the fill, and it read as the others"},
				{Target: plain, Props: []string{"background-color"}, Golden: "another item no ground"},
				{Target: ".cui-console-menu-rule", Props: []string{"background-color"}, Golden: "the rule CanvasText: forced colors drew its fill in the ground's color, at its alpha"},
			},
		},
		{
			Name:   "menu dark",
			HTML:   `<div class="cui-dark">` + h.Example("menu") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: current, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: ".cui-console-menu-rule", Props: []string{"background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
