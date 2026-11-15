//go:build browser

package consolecheckbox

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	c        = "component-Checkbox-variant"
	off      = "label:nth-child(1)"
	on       = "label:nth-child(2)"
	disabled = "label:nth-child(3)"
	// @invariant: a box with no words beside it, as a table's cell writes it by hand
	// (console-table's scroll.html): no Go type writes it, so no example of this part holds it
	unlabeled = `<div class="grid gap-3"><span class="cui-console-checkbox"><input class="cui-console-checkbox-input" type="checkbox" aria-label="Select all" name="all"><span class="cui-console-checkbox-box" aria-hidden="true"><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 6 9 17l-5-5"/></svg><svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14"/></svg></span></span></div>`
)

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "checkbox",
			HTML:   h.Example("checkbox"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: off + " .cui-console-checkbox-box", Props: []string{"box-width", "box-height", "border-top-left-radius", "border-top-width", "border-top-style"}, TB: []string{c + "-controlSize-md-size", c + "-controlSize-md-size", c + "-controlSize-md-borderRadius", c + "-controlSize-md-effect-default-state-off-border-width", c + "-controlSize-md-effect-default-state-off-border-style"}},
				{Target: off, Props: []string{"column-gap", "font-size", "line-height", "font-weight", "color"}, TB: []string{"component-Label-anatomy-checkbox-variant-controlSize-md-inline-on-gap", "component-Label-variant-controlSize-md-inline-on-textStyle-fontSize", "component-Label-variant-controlSize-md-inline-on-textStyle-lineHeight", "component-Label-variant-controlSize-md-inline-on-textStyle-fontWeight", "component-Label-color"}},
				{Target: off + " .cui-console-checkbox-box", Props: []string{"border-top-color", "background-color"}, Golden: "its edge the field's, coolGray-800 at 0.58 (3:1; the toolkit's 0.18 is 1.4:1: tokens.css, decision 21), on white"},
				{Target: on + " .cui-console-checkbox-box", Props: []string{"background-color", "border-top-color"}, Golden: "checked, the brand's flat blue, its edge gone: the toolkit's is a gradient around it under a shadow"},
				{Target: on + " .cui-console-checkbox-box svg:first-child", Props: []string{"visibility", "color", "box-width"}, Golden: "checked, Lucide's check, white, 11px"},
				{Target: off + " .cui-console-checkbox-box svg:first-child", Props: []string{"visibility"}, Golden: "unchecked, no check"},
				{Target: disabled + " .cui-console-checkbox-box", Props: []string{"background-color"}, TB: []string{c + "-effect-disabled-backgroundColor"}},
				{Target: disabled + " .cui-console-checkbox-box", Props: []string{"border-top-color"}, TB: []string{c + "-controlSize-md-effect-disabled-state-off-border-color"}},
				{Target: disabled, Props: []string{"color"}, Golden: "a disabled checkbox's label the disabled label"},
				{Target: off + " .cui-console-checkbox-input", Props: []string{"opacity", "box-width", "box-height", "box-left", "box-top"}, Golden: "the input transparent, the box's size, over it"},
				{Target: off + " .cui-console-checkbox-box", Props: []string{"box-left", "box-top"}, Golden: "the box where its input is"},
			},
		},
		{
			Name:   "checkbox focus",
			HTML:   h.Example("checkbox"),
			Widths: []int{1200},
			Press:  []string{"Tab"},
			Checks: []h.Check{
				{Target: off + " .cui-console-checkbox-box", Props: []string{"outline-style", "outline-width", "outline-color", "outline-offset"}, Golden: "focused from the keyboard, base.css's ring on the box (rule 9; the toolkit's is 3px of the blue at 0.15)"},
			},
		},
		{
			Name:   "checkbox indeterminate",
			HTML:   unlabeled + `<script>document.querySelector('[name="all"]').indeterminate = true</script>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: `[name="all"] + .cui-console-checkbox-box`, Props: []string{"background-color"}, Golden: "indeterminate, the brand's blue"},
				{Target: `[name="all"] + .cui-console-checkbox-box svg:last-child`, Props: []string{"visibility"}, Golden: "indeterminate, Lucide's minus"},
				{Target: `[name="all"] + .cui-console-checkbox-box svg:first-child`, Props: []string{"visibility"}, Golden: "indeterminate, no check"},
				{Target: `[name="all"] + .cui-console-checkbox-box svg:last-child`, Props: []string{"color"}, Golden: "indeterminate, its minus white"},
			},
		},
		{
			Name:   "checkbox touch",
			HTML:   h.Example("checkbox") + unlabeled,
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: off + " .cui-console-checkbox-input", Props: []string{"box-width", "box-height", "box-left", "box-top"}, Golden: "the transparent input 44px square on a touch screen, the HIG's 44pt (Accessibility, Mobility), centered on the box"},
				{Target: off + " .cui-console-checkbox-box", Props: []string{"box-width", "box-height", "box-left", "box-top"}, Golden: "the box unchanged by touch, 15px"},
				{Target: "span.cui-console-checkbox", Props: []string{"box-height"}, Golden: "a table row's <span> not made 44px tall: its row is taller than the input already"},
				{Target: off, Props: []string{"box-height"}, Golden: "a checkbox with its label 44px tall on a touch screen, as the radio, so its input's 44px fits its row: stacked 12px apart, the next one's reached 14px into it"},
			},
		},
		{
			Name:   "checkbox dark",
			HTML:   `<div class="cui-dark">` + h.Example("checkbox") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: off + " .cui-console-checkbox-box", Props: []string{"border-top-color", "background-color"}, Golden: "the dark side: its edge at 0.42 (tokens.css, decisions 18, 21)"},
				{Target: on + " .cui-console-checkbox-box", Props: []string{"background-color"}, Golden: "the dark blue fill, fitted (tokens.css, decision 21)"},
			},
		},
	})
}
