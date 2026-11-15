//go:build browser

package consoletable

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	tt     = "component-Table"
	fitted = "the secondary label fitted to 0.76 (tokens.css, decision 21): the toolkit's 0.60 is 3.4:1 on white"
)

var desktop = []int{1600, 1200, 640}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "table",
			HTML:   h.Example("table"),
			Widths: desktop,
			Checks: []h.Check{
				{Target: "thead th", Props: []string{"box-height", "padding-top", "padding-bottom", "padding-right", "font-size", "line-height", "font-weight"}, TB: []string{tt + "-anatomy-TH-height", tt + "-anatomy-TH-paddingY", tt + "-anatomy-TH-paddingY", tt + "-anatomy-TH-paddingX", tt + "-anatomy-TH-variant-pageScale-desktop-textStyle-fontSize", tt + "-anatomy-TH-variant-pageScale-desktop-textStyle-lineHeight", tt + "-anatomy-TH-variant-pageScale-desktop-textStyle-fontWeight"}},
				{Target: "thead th", Props: []string{"color"}, Golden: fitted},
				{Target: "thead th.cui-console-table-figure", Props: []string{"font-weight", "text-align"}, Golden: "a figure's header at the header's weight, at its end"},
				{Target: "thead th", Props: []string{"border-bottom-width"}, Golden: "no rule under the header: the toolkit's tokens name none for its header row"},
				{Target: "tbody td:nth-child(2)", Props: []string{"padding-top", "padding-bottom", "padding-left", "padding-right"}, TB: []string{tt + "-anatomy-TD-variant-size-md-paddingTop", tt + "-anatomy-TD-variant-size-md-paddingBottom", tt + "-anatomy-TD-paddingX", tt + "-anatomy-TD-paddingX"}},
				{Target: ".cui-console-table", Props: []string{"color", "font-size", "line-height", "font-weight"}, TB: []string{tt + "-color", tt + "-variant-pageScale-desktop-textStyle-fontSize", tt + "-variant-pageScale-desktop-textStyle-lineHeight", tt + "-variant-pageScale-desktop-textStyle-fontWeight"}},
				{Target: "tbody tr:first-child td:last-child", Props: []string{"padding-right"}, Golden: "the last cell flush at its end, as the analytics table's first is at its start"},
				{Target: "tbody tr:first-child", Props: []string{"background-color"}, State: []string{"hover"}, ASC: h.ASCProbe("hover-row", "table.row")},
				{Target: ".cui-console-table-figure:not(th)", Props: []string{"font-size", "line-height", "font-weight", "letter-spacing", "color", "text-align", "font-variant-numeric"}, ASC: h.ASCProbe("analytics", "table.figure")},
				{Target: "tbody td:nth-child(2)", Props: []string{"vertical-align", "border-bottom-width", "border-bottom-style", "border-bottom-color"}, ASC: h.ASCProbe("analytics", "table.cell")},
				{Target: "tbody td:nth-child(2)", Props: []string{"font-size", "line-height", "font-weight", "letter-spacing", "color", "text-align"}, ASC: h.ASCProbe("analytics", "table.text")},
				{Target: ".cui-console-table-detail", Props: []string{"font-size", "line-height", "font-weight", "color", "display"}, Golden: "a line under a cell's text, 12px on 16px, the secondary label: the toolkit's display list's subtitle at the table's small size"},
				{Target: "tbody tr:first-child", Props: []string{"cursor"}, Golden: "a row that is a link, all of it"},
				{Target: ".cui-console-table", Props: []string{"box-width"}, Golden: "the full width of its place"},
				{Target: "tbody tr:first-child td:first-child", Props: []string{"hit"}, Golden: "the pointer finds the row's link anywhere in its first cell"},
			},
		},
		{
			Name:   "table phone",
			HTML:   h.Example("table"),
			Widths: []int{639, 390},
			Checks: []h.Check{
				{Target: "thead th", Props: []string{"font-size", "line-height", "font-weight"}, TB: []string{tt + "-anatomy-TH-variant-pageScale-mobile-textStyle-fontSize", tt + "-anatomy-TH-variant-pageScale-mobile-textStyle-lineHeight", tt + "-anatomy-TH-variant-pageScale-mobile-textStyle-fontWeight"}},
				{Target: ".cui-console-table", Props: []string{"font-size", "line-height", "font-weight"}, TB: []string{tt + "-variant-pageScale-mobile-textStyle-fontSize", tt + "-variant-pageScale-mobile-textStyle-lineHeight", tt + "-variant-pageScale-mobile-textStyle-fontWeight"}},
			},
		},
		{
			Name:   "table scroll",
			HTML:   h.Example("scroll"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: ".cui-console-table-scroll", Props: []string{"overflow-x"}, Golden: "a table wider than its place scrolls sideways in its box"},
				{Target: "tbody tr:first-child", Props: []string{"background-color"}, Golden: "a row whose checkbox is checked is selected: the brand's fill at 0.08, the toolkit's selected row's (--tb-component-Table-anatomy-TR-…-selected-fill), square: its 12px corners set out 12px are no table row's"},
				{Target: "tbody tr:first-child > :first-child", Props: []string{"border-left-width", "border-left-style", "border-left-color", "padding-left"}, Golden: "a selected row its leading bar in every mode, 3px of the brand label's blue, its text 8px in as a cell's: a fill at 0.08 alone is too faint to find by, and forced colors drop it (the HIG's Color: not by color alone)"},
				{Target: "tbody tr:last-child > :first-child", Props: []string{"border-left-width", "padding-left"}, Golden: "another row no bar, its first cell flush"},
				{Target: "tbody tr:first-child .cui-console-checkbox-input", Props: []string{"hit"}, Golden: "the pointer finds a linked row's checkbox, not the row's link: the link's ::after covers the row and comes later in it, and took the checkbox's click"},
				{Target: "tbody tr:first-child td:nth-child(3)", Props: []string{"hit"}, Golden: "elsewhere in the row the pointer still finds its link"},
			},
		},
		{
			Name: "table checkbox in a cell's dialog",
			HTML: `<div class="cui-console-page"><table class="cui-console-table"><thead><tr><th scope="col">Project</th><th scope="col"><span class="sr-only">Actions</span></th></tr></thead><tbody>
      <tr id="routed"><td>Shop App</td><td><form><button type="button" class="cui-console-btn" data-variant="plain">Edit…</button><dialog class="cui-console-dialog"><label class="cui-console-checkbox"><input class="cui-console-checkbox-input" type="checkbox" checked><span class="cui-console-checkbox-box" aria-hidden="true"></span>ops</label></dialog></form></td></tr>
      <tr id="own"><td><span class="cui-console-checkbox"><input class="cui-console-checkbox-input" type="checkbox" checked aria-label="Select"><span class="cui-console-checkbox-box" aria-hidden="true"></span></span></td><td>Orders API</td></tr>
    </tbody></table></div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: "#routed", Props: []string{"background-color"}, Golden: "a checkbox inside a cell's form or dialog is that control's, not the row's: no fill"},
				{Target: "#own", Props: []string{"background-color"}, Golden: "a checkbox that is the cell's own control selects its row: the brand's fill at 0.08"},
			},
		},
		{
			Name:   "table row controls",
			HTML:   `<table class="cui-console-table"><tbody><tr><td><a class="cui-console-table-link" href="#shop">Shop App</a></td><td><button type="button" class="cui-console-btn" data-variant="plain">Edit…</button></td><td><a href="#docs">Docs</a></td></tr></tbody></table>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: "button", Props: []string{"hit"}, Golden: "the pointer finds a linked row's button, not the row's link"},
				{Target: `a[href="#docs"]`, Props: []string{"hit"}, Golden: "and another link in the row"},
			},
		},
		{
			Name:   "table checkbox touch",
			HTML:   h.Example("scroll"),
			Widths: []int{390},
			Touch:  true,
			Checks: []h.Check{
				{Target: "tbody tr:first-child .cui-console-checkbox", Props: []string{"hit"}, Golden: "on a touch screen too a tap on a linked row's checkbox finds the checkbox, not the row's link"},
				{Target: "tbody tr:first-child > td:first-child", Props: []string{"padding-left", "padding-right"}, Golden: "a cell with a checkbox 14.5px in each way on a touch screen, room for its input's 44px: flush at the table's start the scroll box clipped it to 29px"},
			},
		},
		{
			Name:   "table stacked",
			HTML:   h.Example("stack"),
			Widths: []int{639, 390},
			Checks: []h.Check{
				{Target: "tbody tr:first-child", Props: []string{"display", "padding-top", "padding-bottom", "border-bottom-width"}, Golden: "below 640px a row is a block, 12px above and below, its hairline under it"},
				{Target: "thead", Props: []string{"display"}, Golden: "the header hidden: each value says its own label"},
				{Target: "tbody tr:first-child td:nth-child(2)", Props: []string{"display", "justify-content", "padding-top", "border-bottom-width"}, Golden: "a cell a line, its label at its start, its value at its end, 2px above and below"},
				{Target: "tbody tr:first-child td:nth-child(2)", Pseudo: "::before", Props: []string{"content", "color", "font-size", "font-weight"}, Golden: "its label, the cell's data-label, 13px at 500, the secondary label"},
				{Target: "tbody tr:first-child td:first-child", Pseudo: "::before", Props: []string{"content"}, Golden: "a cell with no data-label, its value alone"},
				{Target: "tbody tr:first-child td[data-empty]", Props: []string{"display"}, Golden: "a cell that says nothing (data-empty) left out of the block"},
				{Target: "tbody tr:last-child td:last-child", Props: []string{"display"}, Golden: "the same column with a value, a line"},
			},
		},
		{
			Name:   "table scroll box",
			HTML:   h.Example("scroll"),
			Widths: []int{1200, 390},
			Press:  []string{"Tab"},
			Checks: []h.Check{
				{Target: ".cui-console-table-scroll", Props: []string{"focused", "outline-style", "overflow-x"}, Golden: "the first Tab reaches the scroll box, ringed: the keyboard scrolls what overflows (axe's scrollable-region-focusable)"},
			},
		},
		{
			Name:   "table stacked desktop",
			HTML:   h.Example("stack"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: "tbody tr:first-child td:nth-child(2)", Props: []string{"display"}, Golden: "from 640px a stacking table is a table"},
				{Target: "tbody tr:first-child td[data-empty]", Props: []string{"display"}, Golden: "from 640px an empty cell keeps its column"},
			},
		},
		{
			Name:         "table forced colors",
			HTML:         h.Example("scroll"),
			Widths:       []int{1200},
			ForcedColors: true,
			Checks: []h.Check{
				{Target: "tbody tr:first-child > :first-child", Props: []string{"border-left-width", "border-left-style", "border-left-color"}, Golden: "a selected row its leading edge, 3px of the theme's Highlight on its first cell: forced colors drop the fill, and the row read as the others"},
				{Target: "tbody tr:last-child > :first-child", Props: []string{"border-left-width"}, Golden: "another row no edge"},
			},
		},
		{
			Name:   "table dark",
			HTML:   `<div class="cui-dark">` + h.Example("scroll") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: "tbody td:nth-child(2)", Props: []string{"border-bottom-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: "tbody tr:first-child > :first-child", Props: []string{"border-left-color"}, Golden: "the bar in the dark, the brand label's blue, a step lighter than the accent (decision 21)"},
				{Target: "tbody tr:first-child", Props: []string{"background-color"}, Golden: "the dark side derived: the brand blue, no step of the palette, takes the blue ramp's step nearest its lightness (tokens.css, decision 18)"},
			},
		},
	})
}
