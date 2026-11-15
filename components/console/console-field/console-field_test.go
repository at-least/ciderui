//go:build browser

package consolefield

import (
	"slices"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	tv     = "component-TextInput-variant"
	name   = "#project"
	bad    = "#url"
	off    = "#dsn"
	fitted = "fitted (tokens.css, decision 21)"
)

func size(s string) []string {
	return []string{tv + "-controlSize-" + s + "-blockSize", tv + "-controlShape-roundedRect-controlSize-" + s + "-paddingInlineStart", tv + "-controlShape-roundedRect-controlSize-" + s + "-paddingInlineEnd", tv + "-controlShape-roundedRect-controlSize-" + s + "-borderRadius", tv + "-controlSize-" + s + "-textStyle-fontSize", tv + "-controlSize-" + s + "-textStyle-lineHeight", tv + "-controlSize-" + s + "-textStyle-fontWeight"}
}

var sizeProps = []string{"box-height", "padding-left", "padding-right", "border-top-left-radius", "font-size", "line-height", "font-weight"}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "fields",
			HTML:   h.Example("fields"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: name, State: []string{"focus", "focus-visible"}, Props: []string{"outline-offset"}, Golden: "its halo starts at its edge, not the ring's gap away"},
				{Target: name, Props: sizeProps, TB: size("md")},
				{Target: name, Props: []string{"background-color", "border-top-width", "border-top-style", "color"}, TB: []string{tv + "-colorRole-neutral-variant-stroked-backgroundColor", tv + "-colorRole-neutral-focus-off-variant-stroked-border-width", tv + "-colorRole-neutral-focus-off-variant-stroked-border-style", "component-TextInput-anatomy-input-variant-effect-default-color"}},
				{Target: name, Props: []string{"border-top-color"}, Golden: "the edge coolGray-800 at 0.58, 3:1 on white: the toolkit's 0.18 is 1.4:1, " + fitted},
				{Target: name, Props: []string{"border-top-color", "outline-width", "outline-style", "outline-color"}, State: []string{"focus"}, TB: []string{tv + "-colorRole-neutral-focus-on-variant-stroked-border-color", tv + "-colorRole-neutral-focus-on-outline-width", tv + "-colorRole-neutral-focus-on-outline-style", tv + "-colorRole-neutral-focus-on-outline-color"}},
				{Target: name, Props: []string{"outline-offset"}, State: []string{"focus"}, Golden: "its halo outside its edge, none between"},
				{Target: bad, Props: []string{"background-color", "border-top-width"}, TB: []string{tv + "-colorRole-error-focus-off-variant-stroked-backgroundColor", tv + "-colorRole-error-effect-default-variant-stroked-border-width"}},
				{Target: bad, Props: []string{"border-top-color"}, Golden: "the red edge the red label, red-700: " + fitted},
				{Target: bad, Props: []string{"background-color", "outline-color"}, State: []string{"focus"}, TB: []string{tv + "-colorRole-error-focus-on-variant-stroked-backgroundColor", tv + "-colorRole-error-focus-on-outline-color"}},
				{Target: off, Props: []string{"background-color", "color"}, TB: []string{tv + "-effect-disabled-variant-stroked-backgroundColor", "component-TextInput-anatomy-input-variant-effect-disabled-color"}},
				{Target: "#search", Props: []string{"color"}, Pseudo: "::placeholder", Golden: "the placeholder the secondary label at 0.76: the toolkit's, at 0.39, is 2.1:1 on white, " + fitted},
				{Target: ".cui-console-field-label", Props: []string{"font-size", "line-height", "font-weight", "color", "margin-bottom"}, TB: []string{"component-Label-variant-controlSize-md-inline-off-textStyle-fontSize", "component-Label-variant-controlSize-md-inline-off-textStyle-lineHeight", "component-Label-variant-controlSize-md-inline-off-textStyle-fontWeight", "component-Label-color", "component-FormField-anatomy-label-variant-controlSize-md-marginBlockEnd"}},
				{Target: ".cui-console-field-error", Props: []string{"font-size", "line-height", "font-weight"}, TB: []string{"component-InputError-variant-controlSize-md-textStyle-fontSize", "component-InputError-variant-controlSize-md-textStyle-lineHeight", "component-InputError-variant-controlSize-md-textStyle-fontWeight"}},
				{Target: ".cui-console-field-error", Props: []string{"margin-top"}, TB: []string{"component-FormField-anatomy-field-variant-controlSize-md-marginBlockEnd"}},
				{Target: ".cui-console-field-error", Props: []string{"color"}, Golden: "red-700, the red label: the toolkit's red is 3.6:1 on white, " + fitted},
				{Target: "#note", Props: []string{"padding-top", "padding-bottom", "resize", "box-height"}, Golden: "a multi-line field: its height its 3 rows' and its padding, not the 30px field's, 6px above and below, taller by its handle"},
				{Target: name, Props: []string{"box-width"}, Golden: "as wide as its box, the example's 320px column"},
				{Target: ".cui-console-field", Props: []string{"display", "flex-direction"}, Golden: "its label, field and lines in a column"},
			},
		},
		{
			Name:   "fields phone",
			HTML:   h.Example("fields"),
			Widths: []int{639, 390},
			Checks: []h.Check{
				{Target: name, Props: []string{"font-size", "line-height"}, TB: []string{"typography-textStyle-bodyMobileRegular-fontSize", "typography-textStyle-bodyMobileRegular-lineHeight"}},
				{Target: name, Props: []string{"box-height"}, TB: []string{tv + "-controlSize-md-blockSize"}},
			},
		},
		{
			Name:   "fields touch",
			HTML:   h.Example("fields"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: name, Props: []string{"box-height"}, Golden: "a field 44px tall on a touch screen, the HIG's 44pt (Accessibility, Mobility)"},
			},
		},
		{
			Name:   "select",
			HTML:   h.Example("select"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: "select", Props: slices.Delete(slices.Clone(sizeProps), 2, 3), TB: slices.Delete(size("md"), 2, 3)},
				{Target: "select", Props: []string{"padding-right"}, Golden: "room for its chevron at its end"},
				{Target: ".cui-console-select-chevron", Props: []string{"box-width", "box-height", "right", "color"}, Golden: "its chevron, Lucide's chevrons-up-down, 13px, 8px in, the secondary label"},
			},
		},
		{
			Name:   "fields dark",
			HTML:   `<div class="cui-dark">` + h.Example("fields") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: name, Props: []string{"background-color", "border-top-color", "color"}, Golden: "the dark side derived; its edge fitted to 3:1 on black at 0.42 (tokens.css, decisions 18, 21)"},
				{Target: bad, Props: []string{"border-top-color", "background-color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
