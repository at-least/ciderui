//go:build browser

package consoledialog

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const m = "component-Modal"

// @external: shown by its open attribute, not showModal(): no ::backdrop
func open() string {
	return strings.Replace(h.Example("dialog"), ` aria-labelledby="ignore-title">`, ` aria-labelledby="ignore-title" open>`, 1)
}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "dialog",
			HTML:   open(),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-dialog-actions", Props: []string{"display"}, Golden: "its actions in a row"},
				{Target: ".cui-console-dialog", Props: []string{"margin-left"}, Golden: "centered: the kit's reset takes the browser's margin away"},
				{Target: ".cui-console-dialog", Props: []string{"box-width", "border-top-left-radius", "background-color", "border-top-color", "box-shadow", "color"}, TB: []string{m + "-variant-variant-modal-minWidth", m + "-variant-variant-modal-borderRadius", m + "-backgroundColor", m + "-variant-variant-modal-stroke", m + "-variant-variant-modal-boxShadow", m + "-color"}},
				{Target: ".cui-console-dialog", Props: []string{"max-width"}, TB: []string{m + "-variant-variant-modal-maxWidth"}},
				{Target: ".cui-console-dialog", Props: []string{"padding-top", "padding-left", "border-top-width"}, Golden: "ciderui's room inside, 24px (the toolkit's modal names none); a 1px edge"},
				{Target: ".cui-console-dialog", Props: []string{"font-size", "line-height"}, Golden: "the console's body, 13px on 18px"},
				{Target: ".cui-console-dialog-title", Props: []string{"font-size", "line-height", "font-weight", "margin-bottom"}, Golden: "the console's h4 at 600, 12px over what follows"},
				{Target: ".cui-console-dialog-actions", Props: []string{"justify-content", "column-gap", "margin-top"}, Golden: "its buttons at its end, 9px apart (the banner's), 24px under what it asks"},
			},
		},
		{
			Name:   "dialog modal",
			HTML:   h.Example("dialog"),
			Click:  []string{`[command="show-modal"]`},
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-dialog", Pseudo: "::backdrop", Props: []string{"background-color"}, Golden: "ciderui's scrim behind it (--cui-c-scrim): the toolkit names no backdrop"},
			},
		},
		{
			Name:   "dialog phone",
			HTML:   open(),
			Widths: []int{390},
			Checks: []h.Check{
				{Target: ".cui-console-dialog", Props: []string{"box-width"}, Golden: "never wider than the window less 22px, 11px at each side, as the sidebar's sheet"},
			},
		},
		{
			Name:   "dialog dark",
			HTML:   `<div class="cui-dark">` + open() + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-console-dialog", Props: []string{"background-color", "border-top-color", "color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
			},
		},
	})
}
