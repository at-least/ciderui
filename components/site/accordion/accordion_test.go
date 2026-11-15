//go:build browser

package accordion

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	item   = ".cui-accordion-item"
	open   = ".cui-accordion-item[open]"
	closed = ".cui-accordion-item:not([open])"
	q      = ".cui-accordion-question"
	// @invariant: an item renders closed; a case reads an open one after a click opens the first
	openFirst = item + ":nth-child(1) > " + q
)

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:  "zh-hant",
			HTML:  `<div lang="zh-Hant">` + h.Example("accordion") + `</div>`,
			Click: []string{openFirst},
			Checks: []h.Check{
				{Target: open + " .cui-accordion-answer", Props: []string{"line-height", "letter-spacing"}, Apple: h.AppleProbe("buy-tw", "accordion.content")},
			},
		},
		{
			Name:  "accordion",
			HTML:  h.Example("accordion"),
			Click: []string{openFirst},
			Checks: []h.Check{
				{Target: item, Props: []string{"border-top-width", "border-top-style", "border-top-color", "padding-top", "margin-top"}, Apple: h.AppleProbe("bag", "accordion.item")},
				{Target: closed + " > " + q, Props: []string{"padding-top", "padding-bottom", "display", "cursor", "box-height"}, Apple: h.AppleProbe("bag-accordion", "compact.button")},
				{Target: closed + " > " + q, Props: h.Props(h.Size, []string{"color"}), Apple: h.AppleProbe("bag-accordion", "compact.title")},
				{Target: closed + " > " + q + " svg", Props: []string{"color"}, Apple: h.AppleProbe("bag-accordion", "compact.icon")},
				{Target: closed + " > " + q + " svg", Props: []string{"box-height"}, Apple: h.AppleProbe("bag-accordion", "compact.svg")},
				{Target: closed + " > " + q + " svg", Props: []string{"transition-property", "transition-duration", "transition-timing-function"}, Golden: "its turn in 0.32s, its color at once, where apple.com's eases in 0.1s linear (bag-accordion:compact.icon): motion that adds nothing to a pointer's color change (DESIGN.md decision 25)"},
				{Target: closed + " > " + q + " svg", On: closed + " > " + q, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe("bag-accordion", "compact.icon-hover", "hover")},
				{Target: open + " > " + q + " svg", Props: []string{"color"}, Apple: h.AppleProbe("bag-accordion", "compact.open")},
				{Target: open + " .cui-accordion-answer", Props: h.Props(h.Size, []string{"color", "padding-right", "padding-bottom"}), Apple: h.AppleProbe("bag-accordion", "compact.answer")},
				{Target: open + " .cui-accordion-answer", Props: []string{"padding-top"}, Apple: &h.Apple{Page: "bag-accordion", Probe: "compact.panel", As: map[string]string{"padding-top": "margin-top"}}},
				{Target: open + " > " + q + " svg", Props: []string{"transform"}, Golden: "open, turned up: apple.com's line moved to the same points, through a flat one, in 320ms"},
				{Target: closed + " > " + q + " svg", Props: []string{"box-width", "margin-right"}, Golden: "apple.com's 17px chevron (14px) in the middle of its 32px column at the end (bag-accordion:compact.svg, .icon)"},
				{Target: item, Pseudo: "::details-content", Props: h.Motion, Golden: "its answer opens at once, where apple.com's height slides in 0.3s ease-in-out (bag-accordion:compact.panel): the chevron's turn says it opened, and the words are there to read (DESIGN.md decision 25)"},
				{Target: q, Props: []string{"list-style-type"}, Golden: "no marker: the chevron is the question's"},
			},
		},
		{
			Name:   "click",
			HTML:   h.Example("accordion"),
			Widths: []int{1200},
			Click:  []string{".cui-accordion-item:nth-child(2) > .cui-accordion-question"},
			Checks: []h.Check{
				{Target: ".cui-accordion-item:nth-child(2) > .cui-accordion-question svg", Props: []string{"transform"}, Golden: "clicked, its question's chevron turned up: <details> opens with no script"},
				{Target: ".cui-accordion-item:nth-child(2) .cui-accordion-answer", Props: []string{"box-height"}, Golden: "its answer shown, at its height (13.6px above, 28px below its line)"},
			},
		},
	})
}
