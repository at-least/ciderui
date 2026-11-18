//go:build browser

package consoleswitch

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	t        = "component-Toggle-anatomy-"
	md       = "variant-controlSize-md-"
	off      = "label:nth-child(1)"
	on       = "label:nth-child(2)"
	disabled = "label:nth-child(3)"
)

func TestCases(tt *testing.T) {
	h.Run(tt, []h.Case{
		{
			Name:   "switch",
			HTML:   h.Example("switch"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: off + " .cui-console-switch-track", Props: []string{"box-width", "box-height", "border-top-left-radius"}, TB: []string{t + "background-" + md + "width", t + "background-" + md + "height", t + "background-" + md + "borderRadius"}},
				{Target: off + " .cui-console-switch-track", Pseudo: "::before", Props: []string{"width", "height", "box-shadow"}, TB: []string{t + "knob-" + md + "width", t + "knob-" + md + "height", t + "knob-boxShadow"}},
				{Target: off, Props: []string{"column-gap", "font-size", "line-height", "font-weight", "color"}, TB: []string{"component-Label-anatomy-toggle-variant-controlSize-md-inline-on-gap", "component-Label-variant-controlSize-md-inline-on-textStyle-fontSize", "component-Label-variant-controlSize-md-inline-on-textStyle-lineHeight", "component-Label-variant-controlSize-md-inline-on-textStyle-fontWeight", "component-Label-color"}},
				{Target: off + " .cui-console-switch-track", Props: []string{"background-color"}, TB: []string{t + "background-variant-effect-default-state-off-backgroundColor"}},
				{Target: on + " .cui-console-switch-track", Props: []string{"background-color"}, TB: []string{t + "background-variant-effect-default-state-on-backgroundColor"}},
				{Target: off + " .cui-console-switch-track", Props: []string{"border-top-width", "border-top-color"}, Golden: "off, its edge the field's, coolGray-800 at 0.58 (3:1, a control's boundary: decision 21), where the toolkit's is a 1px shadow at 0.16 (1.3:1)"},
				{Target: on + " .cui-console-switch-track", Props: []string{"border-top-color"}, Golden: "on, its edge gone into the blue"},
				{Target: off + " .cui-console-switch-track", Pseudo: "::before", Props: []string{"left", "top", "translate", "background-color"}, Golden: "off, the knob white at the start of the track's 1px edge, 1px from its outer edge as the toolkit's state-off-left"},
				{Target: on + " .cui-console-switch-track", Pseudo: "::before", Props: []string{"translate"}, Golden: "on, the knob moved 12px, 13px from the track's outer edge as the toolkit's state-on-left"},
				{Target: disabled + " .cui-console-switch-track", Props: []string{"opacity"}, TB: []string{"component-Toggle-disabledOpacity"}},
				{Target: disabled, Props: []string{"color"}, Golden: "a disabled switch's label the disabled label, as the checkbox's"},
				{Target: off + " .cui-console-switch-input", Props: []string{"opacity", "box-width", "box-height", "box-left", "box-top"}, Golden: "the input transparent, the track's size, over it"},
			},
		},
		{
			Name:   "switch focus",
			HTML:   h.Example("switch"),
			Widths: []int{1200},
			Press:  []string{"Tab"},
			Checks: []h.Check{
				{Target: off + " .cui-console-switch-track", Props: []string{"outline-style", "outline-width", "outline-color", "outline-offset"}, Golden: "focused from the keyboard, base.css's ring on the track (rule 9; the toolkit's is 3px of the blue at 0.15)"},
			},
		},
		{
			Name:   "switch touch",
			HTML:   h.Example("switch"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: off + " .cui-console-switch-input", Props: []string{"box-width", "box-height", "box-left", "box-top"}, Golden: "the transparent input 44px square on a touch screen, the HIG's 44pt (Accessibility, Mobility), centered on the track"},
				{Target: off + " .cui-console-switch-track", Props: []string{"box-width", "box-height"}, Golden: "the track unchanged by touch, 30 by 18px"},
				{Target: off, Props: []string{"box-height"}, Golden: "a switch with its label 44px tall on a touch screen, as the checkbox, so its input's 44px fits its row"},
			},
		},
		{
			Name:         "switch forced colors",
			HTML:         h.Example("switch"),
			ForcedColors: true,
			Widths:       []int{1200},
			Checks: []h.Check{
				{Target: off + " .cui-console-switch-track", Props: []string{"border-top-style"}, Golden: "off, in forced colors its edge still drawn: a border, which forced colors keep where they drop a shadow"},
				{Target: off + " .cui-console-switch-track", Pseudo: "::before", Props: []string{"background-color"}, Golden: "off, the knob the text's color on the page's"},
				{Target: on + " .cui-console-switch-track", Props: []string{"background-color"}, Golden: "on, the track the system's highlight"},
				{Target: on + " .cui-console-switch-track", Pseudo: "::before", Props: []string{"background-color"}, Golden: "on, the knob the highlight's text color"},
			},
		},
		{
			Name:   "switch dark",
			HTML:   `<div class="cui-dark">` + h.Example("switch") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: off + " .cui-console-switch-track", Props: []string{"border-top-color", "background-color"}, Golden: "the dark side: its edge at 0.42, its fill the derived gray (tokens.css, decisions 18, 21)"},
				{Target: on + " .cui-console-switch-track", Props: []string{"background-color"}, Golden: "the dark blue fill, fitted (tokens.css, decision 21)"},
				{Target: off + " .cui-console-switch-track", Pseudo: "::before", Props: []string{"background-color"}, Golden: "the knob white in dark too, as iOS's"},
			},
		},
	})
}
