//go:build browser

package consolesidebar

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	s       = "component-Sidebar"
	d       = "component-DisplayList-variant"
	ds      = d + "-pageScale-desktop-variant-sidebar"
	ms      = d + "-pageScale-mobile-variant-sidebar"
	sidebar = ".cui-console-sidebar"
	item    = `.cui-console-sidebar .cui-console-menu-item[href="#projects"]`
	current = `.cui-console-sidebar .cui-console-menu-item[aria-current="page"]`
	fitted  = "the secondary label fitted to 0.76 (tokens.css, decision 21)"
)

func page(html string) string { return `<div class="cui-console-page">` + html + `</div>` }

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "sidebar",
			HTML:   page(h.Example("sidebar")),
			Widths: []int{1600, 1200, 1024},
			Checks: []h.Check{
				{Target: sidebar, Props: []string{"bottom"}, Golden: "in the page from 1024px, not the sheet's full height"},
				{Target: sidebar, Props: []string{"box-width"}, TB: []string{"size-sidebar-width"}},
				{Target: sidebar, Props: []string{"color"}, TB: []string{"color-role-neutral-label-primary"}},
				{Target: sidebar, Props: []string{"flex-direction"}, Golden: "a column: its header, list and footer one under another"},
				{Target: sidebar, Props: []string{"display", "position", "background-color", "box-shadow"}, Golden: "from 1024px it stands in the page: the toolkit's plain sidebar, no ground, no shadow"},
				{Target: ".cui-console-sidebar-header", Props: []string{"padding-top", "padding-bottom", "padding-left", "padding-right", "column-gap", "min-height"}, TB: []string{s + "-anatomy-Header-paddingBlockStart", s + "-anatomy-Header-paddingBlockEnd", s + "-anatomy-Header-paddingInline", s + "-anatomy-Header-paddingInline", s + "-anatomy-Header-anatomy-Nav-gap", s + "-anatomy-Header-anatomy-Nav-minHeight"}},
				{Target: ".cui-console-sidebar-wordmark", Props: []string{"margin-left"}, TB: []string{s + "-anatomy-Header-anatomy-Nav-anatomy-Wordmark-marginInlineStart"}},
				{Target: ".cui-console-sidebar-wordmark", Props: []string{"font-size", "line-height", "font-weight"}, Golden: "the wordmark the product's name, 17px on 22px at 600: the toolkit's is a picture"},
				{Target: ".cui-console-sidebar-close", Props: []string{"display"}, Golden: "from 1024px no close button: it stands in the page"},
				{Target: ".cui-console-sidebar-body", Props: []string{"padding-left", "padding-right"}, TB: []string{s + "-anatomy-Body-paddingInline", s + "-anatomy-Body-paddingInline"}},
				{Target: item, Props: []string{"min-height", "padding-top", "padding-left", "border-top-left-radius", "column-gap", "font-size", "line-height", "font-weight"}, TB: []string{ds + "-anatomy-Item-minHeight", ds + "-anatomy-Item-paddingBlock", ds + "-anatomy-Item-paddingInline", ds + "-anatomy-Item-borderRadius", ds + "-anatomy-Item-gap", ds + "-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-fontSize", ds + "-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-lineHeight", ds + "-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-fontWeight"}},
				{Target: item, Props: []string{"color"}, TB: []string{strings.Replace(d, "-variant", "", 1) + "-anatomy-Item-variant-colorRole-neutral-effect-default-anatomy-Label-color"}},
				{Target: item, Props: []string{"background-color"}, State: []string{"hover"}, TB: []string{d + "-variant-sidebar-anatomy-Item-variant-effect-hover-focus-off-selected-none-backgroundColor"}},
				{Target: item, Props: []string{"background-color"}, State: []string{"active"}, Golden: "pressed as hovered: the toolkit's pressed fill, oklch(42.6% 0.008 286 / 0.03), is no step of its palette"},
				{Target: current, Props: []string{"background-color", "font-weight"}, TB: []string{d + "-variant-sidebar-anatomy-Item-variant-effect-default-focus-off-selected-fill-backgroundColor", ds + "-anatomy-Item-variant-selected-fill-anatomy-Label-textStyle-fontWeight"}},
				{Target: ".cui-console-sidebar .cui-console-menu-title", Props: []string{"font-size", "line-height", "font-weight", "padding-top", "padding-bottom", "padding-left", "margin-top"}, TB: []string{ds + "-anatomy-Heading-textStyle-fontSize", ds + "-anatomy-Heading-textStyle-lineHeight", ds + "-anatomy-Heading-textStyle-fontWeight", ds + "-anatomy-Header-paddingBlockStart", ds + "-anatomy-Header-paddingBlockEnd", ds + "-anatomy-Header-paddingInline", d + "-variant-sidebar-anatomy-groupGap-marginBlockEnd"}},
				{Target: ".cui-console-sidebar .cui-console-menu-title", Props: []string{"color"}, Golden: fitted},
				{Target: item + " .cui-console-menu-label", Props: []string{"white-space", "text-overflow", "overflow-wrap"}, Golden: "a long label wraps, whole, as a menu's"},
				{Target: ".cui-console-sidebar .cui-console-menu-rule", Props: []string{"margin-top", "margin-left", "background-color", "box-height"}, TB: []string{d + "-variant-sidebar-anatomy-HR-marginBlock", ds + "-anatomy-HR-marginInline", "color-role-neutral-border-secondary", "border-width-fine"}},
				{Target: ".cui-console-sidebar .cui-console-menu-icon", Props: []string{"box-width", "box-height"}, Golden: "its icon Lucide's, 20px in 24px (the toolkit's art's least width): the toolkit draws SF Symbols at its text's size"},
				{Target: ".cui-console-sidebar-footer", Props: []string{"padding-top", "padding-bottom"}, TB: []string{s + "-anatomy-UniversalIDFooter-paddingBlockStart", s + "-anatomy-UniversalIDFooter-paddingBlockEnd"}},
			},
		},
		{
			Name:   "sidebar sheet tall",
			HTML:   page(h.Example("open")),
			Widths: []int{390},
			Height: 1400,
			Click:  []string{".cui-console-sidebar-open"},
			Checks: []h.Check{
				{Target: ".cui-console-sidebar-footer", Props: []string{"box-top", "box-height"}, Golden: "the footer at the sheet's foot, 11px above the window's: the body takes the rest of the column"},
			},
		},
		{
			Name:   "sidebar sheet",
			HTML:   page(h.Example("open")),
			Widths: []int{390},
			Click:  []string{".cui-console-sidebar-open"},
			Checks: []h.Check{
				{Target: ".cui-console-menu-title", Props: []string{"padding-left", "font-size", "line-height"}, Golden: "in the sidebar, a heading lined up with its rows, in their size"},
				{Target: sidebar, Props: []string{"box-left", "box-top", "border-top-left-radius", "background-color", "box-shadow"}, TB: []string{s + "-inset", s + "-inset", s + "-borderRadius", s + "-variant-variant-opaque-backgroundColor", s + "-variant-variant-opaque-boxShadow"}},
				{Target: sidebar, Props: []string{"position", "display"}, Golden: "below 1024px, opened, a sheet over the page: a popover, fixed"},
				{Target: ".cui-console-sidebar-close", Props: []string{"display"}, Golden: "its close button shown in the sheet"},
				{Target: ".cui-console-sidebar-close", Props: []string{"box-right"}, Golden: "at the header's end, 6px in"},
				{Target: item, Props: []string{"min-height", "padding-top", "padding-left", "border-top-left-radius", "font-size", "line-height"}, TB: []string{ms + "-anatomy-Item-minHeight", ms + "-anatomy-Item-paddingBlock", ms + "-anatomy-Item-paddingInline", ms + "-anatomy-Item-borderRadius", ms + "-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-fontSize", ms + "-anatomy-Item-variant-selected-none-anatomy-Label-textStyle-lineHeight"}},
			},
		},
		{
			Name:   "sidebar opened focus",
			HTML:   strings.Replace(h.Example("sidebar"), `<nav class="cui-console-sidebar"`, `<button id="open" popovertarget="cui-console-sidebar">Open</button><nav class="cui-console-sidebar"`, 1),
			Click:  []string{"#open"},
			Widths: []int{390},
			Checks: []h.Check{
				{Target: ".cui-console-sidebar-close", Props: []string{"focused"}, Golden: "opened, the focus on its close button, as the global menu's (HTML's autofocus in a popover): a screen reader is in the sheet, and Escape or Return closes it"},
			},
		},
		{
			Name:   "sidebar load focus",
			HTML:   h.Example("sidebar"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-sidebar-close", Props: []string{"focused"}, Golden: "from 1024px, the page loaded, the focus left alone: its close button stands in the page, not shown"},
			},
		},
		{
			Name:   "sidebar closed",
			HTML:   page(h.Example("open")),
			Widths: []int{1023, 390},
			Checks: []h.Check{
				{Target: sidebar, Props: []string{"display"}, Golden: "below 1024px, closed, not shown"},
				{Target: ".cui-console-sidebar-open", Props: []string{"display"}, Golden: "below 1024px, the page's button that opens it, shown"},
			},
		},
		{
			Name:   "sidebar open button wide",
			HTML:   page(h.Example("open")),
			Widths: []int{1024},
			Checks: []h.Check{
				{Target: ".cui-console-sidebar-open", Props: []string{"display"}, Golden: "from 1024px no button opens it: it stands in the page"},
			},
		},
		{
			Name:         "sidebar forced colors",
			HTML:         page(h.Example("sidebar")),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: current, Props: []string{"background-color", "color"}, Golden: "the current page the theme's Highlight under HighlightText: forced colors drop the fill, and it read as the others"},
				{Target: item, Props: []string{"background-color"}, Golden: "another item no ground"},
				{Target: ".cui-console-sidebar .cui-console-menu-rule", Props: []string{"background-color"}, Golden: "the rule CanvasText: forced colors drew its fill in the ground's color, at its alpha"},
			},
		},
		{
			Name:         "sidebar sheet forced colors",
			HTML:         page(h.Example("open")),
			Widths:       []int{390},
			Click:        []string{".cui-console-sidebar-open"},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: sidebar, Props: []string{"border-top-width", "border-top-style", "border-top-color", "box-shadow"}, Golden: "the sheet a 1px CanvasText edge, as the popover's and the toast's: forced colors drop its shadow, and its ground is the page's"},
			},
		},
		{
			Name:   "sidebar dark",
			HTML:   `<div class="cui-dark">` + page(h.Example("sidebar")) + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: current, Props: []string{"background-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
		{
			Name:   "wordmark touch",
			HTML:   page(h.Example("sidebar")),
			Touch:  true,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-sidebar-wordmark", Props: []string{"box-height"}, Golden: "the wordmark unchanged by touch, its 22px line"},
				{Target: ".cui-console-sidebar-wordmark", Pseudo: "::after", Props: []string{"height"}, Golden: "its target a transparent ::after 44px tall at least on a touch screen, the HIG's 44pt (22px is under its 28pt minimum), as the console's link draws one"},
			},
		},
	})
}
