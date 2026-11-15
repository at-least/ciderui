//go:build browser

package field

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	input  = ".cui-field-input"
	label  = ".cui-field-label"
	errorM = ".cui-field-error"
	mark   = errorM + " svg"
	prefix = ".cui-field-prefix"
	long   = "A label far longer than the field it names on a phone, and longer still"
)

var (
	edge = []string{"border-top-width", "border-top-style", "border-top-color", "border-bottom-width", "border-bottom-color", "border-top-left-radius"}
	font = []string{"font-size", "line-height"}
)

func dark(html string) string { return `<div class="cui-dark">` + html + `</div>` }

// @external: the stylesheet's :-webkit-autofill rules copied onto a class: no script can make
// a browser autofill. Its reason is the copy's, so a rule dropped fails here
const autofilled = `<script>
  const copy = (rules) => { for (const rule of [...rules]) {
    if (rule.cssRules) copy(rule.cssRules);
    if (rule.selectorText?.includes(':-webkit-autofill')) (rule.parentRule ?? rule.parentStyleSheet).insertRule(rule.cssText.replaceAll(':-webkit-autofill', '.autofilled'));
  } };
  for (const sheet of document.styleSheets) copy(sheet.cssRules);
  document.querySelector('.cui-field-input').classList.add('autofilled');
</script>`

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name: "field",
			HTML: h.Example("field"),
			Checks: []h.Check{
				{Target: input, Props: []string{"display"}, Golden: "no line box under it"},
				{Target: ".cui-field", Props: []string{"position"}, Golden: "what its label and prefix are placed against, wherever it sits on a page"},
				{Target: input, Props: h.Props(h.Size, []string{"color"}, h.Padding, edge, []string{"background-color", "height", "box-sizing", "box-height", "transition-property", "transition-duration", "outline-style"}), Apple: h.AppleProbe("retail", "field.input")},
				{Target: input, State: []string{"hover"}, Props: []string{"color", "background-color", "border-top-color"}, Apple: h.AppleProbe("retail", "field.input", "hover")},
				{Target: input, State: []string{"focus"}, Props: h.Props(h.Padding, edge), Apple: h.AppleProbe("retail", "field.input", "focus")},
				// @trap: against the rest probe: focused, the 2px edge takes the padding's pixel, the box
				// unchanged
				{Target: input, State: []string{"focus"}, Props: []string{"box-height"}, Apple: h.AppleProbe("retail", "field.input")},
				{Target: input, State: []string{"focus", "focus-visible"}, Props: []string{"padding-left", "border-top-width", "border-top-color", "outline-style"}, Apple: h.AppleProbe("retail", "field.input", "focus+focus-visible")},
				{Target: label, Props: h.Props(h.Size, []string{"color", "position", "top", "left", "white-space", "transition-property", "transition-duration", "transition-timing-function"}), Apple: h.AppleProbe("retail", "field.label")},
				{Target: label, On: input, State: []string{"focus"}, Props: h.Props(h.Size, []string{"color", "top", "left"}), Apple: h.AppleProbe("retail", "field.label-floated")},
				{Target: ".cui-field", Props: []string{"background-color", "border-top-left-radius"}, Golden: "its own light ground under the translucent input, its corners the input's: it reads in a dark block too, where apple.com shows no field"},
				{Target: input, Props: []string{"margin-bottom"}, Golden: "no margin of its own: apple.com's 14px below each field is the form's spacing, the page's (retail:field.input)"},
			},
		},
		{
			Name:   "help",
			HTML:   h.Example("help"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: ".cui-field-help", Props: h.Props(h.Size, []string{"color", "letter-spacing", "margin-top"}), Golden: "the error message's size and place, in the secondary text: the store's form hints (retail: not one measured alone)"},
			},
		},
		{
			Name:   "long label",
			HTML:   strings.Replace(h.Example("field"), ">Email</label>", ">"+long+"</label>", 1),
			Widths: []int{390},
			Checks: []h.Check{
				{Target: label, Props: []string{"box-width", "overflow-x", "text-overflow", "white-space"}, Golden: "a long label stops 17px short of the field's end, with an ellipsis, where it ran past the field (apple.com's labels are short: not measured)"},
			},
		},
		{
			Name:   "autofilled",
			HTML:   h.Example("field") + autofilled,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: label, Props: []string{"top", "font-size", "line-height"}, Golden: "an autofilled field's label floats, as over any value, so it does not lie over the filled text (a browser's real fill not read: no script can make one)"},
				{Target: input, Props: []string{"box-shadow"}, Golden: "its ground the page's, an inset shadow over the browser's own autofill color, which is !important"},
			},
		},
		{
			Name: "filled",
			HTML: h.Example("filled"),
			Checks: []h.Check{
				{Target: input, Props: h.Props(h.Padding, edge, []string{"color", "background-color"}), Apple: h.AppleProbe("retail", "field.input-filled")},
				{Target: label, Props: h.Props(h.Size, []string{"color", "top", "left"}), Apple: h.AppleProbe("retail", "field.label-floated")},
			},
		},
		{
			Name: "dark",
			HTML: dark(h.Example("field")),
			Checks: []h.Check{
				{Target: input, Props: []string{"color", "background-color", "border-top-color"}, Apple: h.AppleProbe("retail-dark", "field.input")},
				{Target: label, Props: []string{"color"}, Apple: h.AppleProbe("retail-dark", "field.label")},
			},
		},
		{
			Name: "error",
			HTML: h.Example("error"),
			Checks: []h.Check{
				{Target: mark, Props: []string{"flex-shrink"}, Golden: "a long message does not squeeze its icon"},
				{Target: input, Props: h.Props(h.Size, []string{"color", "height", "border-top-width", "border-top-color", "border-bottom-color", "background-color"}), Apple: h.AppleProbe("giftcard", "error")},
				{Target: input, State: []string{"focus"}, Props: h.Props(h.Padding, []string{"border-top-width", "border-top-color", "background-color"}), Apple: h.AppleProbe("giftcard", "error", "focus")},
				{Target: label, Props: h.Props(h.Size, []string{"color", "top", "left"}), Apple: h.AppleProbe("giftcard", "error.label")},
				{Target: label, On: input, State: []string{"focus"}, Props: []string{"color"}, Golden: "its label red focused too: apple.com's keeps the class that makes it so (giftcard:error.label), not read focused"},
				{Target: errorM, Props: h.Props(h.Size, []string{"color"}), Apple: h.AppleProbe("giftcard", "error.message")},
				{Target: errorM, Props: []string{"margin-top", "display"}, Apple: h.AppleProbe("giftcard", "error.wrapper")},
				{Target: errorM, Props: []string{"margin-bottom"}, Golden: "no margin below it: apple.com's 12px is the form's spacing, the page's (giftcard:error.wrapper)"},
				{Target: mark, Props: []string{"margin-right"}, Apple: &h.Apple{Page: "giftcard", Probe: "error.icon", As: map[string]string{"margin-right": "padding-right"}}},
				{Target: mark, Props: []string{"box-width", "box-height", "margin-top", "color"}, Golden: "apple.com's mark (U+F312) drawn as ciderui's ALERT: its advance by its ink's height at 12px, its ink 2.75px down the row as apple.com's (reference/icons/form.mjs), in the message's red"},
			},
		},
		{
			Name: "error-dark",
			HTML: dark(h.Example("error")),
			Checks: []h.Check{
				{Target: input, Props: []string{"color", "background-color", "border-top-color", "border-top-width"}, Apple: h.AppleProbe("retail-dark", "field.error")},
				{Target: label, Props: []string{"color"}, Apple: h.AppleProbe("retail-dark", "field.error-label")},
				{Target: errorM, Props: []string{"color"}, Golden: "on the dark block, the fitted dark red (tokens.css: apple.com's #e30000 is below 4.5:1 on black; its forms show no message on a dark page)"},
			},
		},
		{
			Name: "disabled",
			HTML: h.Example("disabled"),
			Checks: []h.Check{
				{Target: input, Props: []string{"color", "background-color", "border-top-color", "cursor"}, Apple: h.AppleProbe("retail", "field.disabled")},
				{Target: label, Props: []string{"color"}, Apple: h.AppleProbe("retail", "field.disabled-label")},
			},
		},
		{
			// @external: the gift card's page read at 1600px alone (before its choices)
			Name:   "disabled-page",
			HTML:   h.Example("disabled"),
			Widths: []int{1600},
			Checks: []h.Check{
				{Target: input, Props: h.Props(h.Size, []string{"color", "background-color", "border-top-color", "border-top-width", "cursor", "height"}), Apple: h.AppleProbe("giftcard", "field.disabled")},
				{Target: label, Props: h.Props(h.Size, []string{"color"}), Apple: h.AppleProbe("giftcard", "label.disabled")},
			},
		},
		{
			Name: "disabled-dark",
			HTML: dark(h.Example("disabled")),
			Checks: []h.Check{
				{Target: input, Props: []string{"color", "background-color", "border-top-color", "cursor"}, Apple: h.AppleProbe("retail-dark", "field.disabled")},
				{Target: label, Props: []string{"color"}, Apple: h.AppleProbe("retail-dark", "field.disabled-label")},
			},
		},
		{
			// @external: read at 1600px alone (the gift card's page enables it at its first width)
			Name:   "prefix",
			HTML:   h.Example("prefix"),
			Widths: []int{1600},
			Checks: []h.Check{
				{Target: prefix, Props: []string{"pointer-events"}, Golden: "a click on it reaches the input under it"},
				{Target: input, Props: h.Props(h.Padding, edge, []string{"height", "background-color"}), Apple: h.AppleProbe("giftcard", "prefix.input")},
				{Target: input, State: []string{"focus"}, Props: []string{"padding-left", "border-top-width", "border-top-color"}, Apple: h.AppleProbe("giftcard", "prefix.input", "focus")},
				{Target: label, Props: h.Props(h.Size, []string{"color", "top", "left", "padding-left"}), Apple: h.AppleProbe("giftcard", "prefix.label")},
				{Target: prefix, Props: h.Props(font, []string{"color", "padding-top", "padding-left", "padding-right", "background-color", "border-right-width", "border-right-style", "border-right-color", "border-top-left-radius", "border-bottom-left-radius", "border-top-right-radius", "position", "top", "left", "height"}), Apple: h.AppleProbe("giftcard", "prefix.currency")},
				{Target: prefix, Props: []string{"box-width", "letter-spacing"}, Golden: "its text's width and 16px either side, as apple.com's (43.34px in SF: giftcard:prefix.currency); the fitted letter-spacing"},
			},
		},
	})
}
