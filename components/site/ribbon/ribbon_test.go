//go:build browser

package ribbon

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const link = ".cui-ribbon .cui-link"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "ribbon",
			Widths: h.With1600,
			HTML:   h.Example("ribbon"),
			Checks: []h.Check{
				{Target: ".cui-ribbon", Props: []string{"overflow", "box-height"}, Apple: h.AppleProbe("ribbon", "ribbon")},
				{Target: ".cui-ribbon-content", Props: h.Props(h.Size, []string{"color", "padding-top", "padding-bottom", "text-align"}), Apple: h.AppleProbe("ribbon", "ribbon.ground")},
				// @external: a ribbon's ground is its page's: the AirPods page's and the store's turn white, the bag's is #f5f5f7 (2026-11-13)
				{Target: ".cui-ribbon-content", Props: []string{"background-color"}, Apple: h.AppleProbe("bag", "ribbon")},
				{Target: ".cui-ribbon-content", Props: []string{"padding-left", "padding-right"}, Golden: "the message within 83.33% of the window, centered: apple.com's content's width (ribbon:ribbon.content)"},
				{Target: ".cui-ribbon-content p", Props: h.Props(h.Size, []string{"margin-top", "margin-bottom"}), Apple: h.AppleProbe("ribbon", "ribbon.text")},
				{Target: link, Props: h.Props(h.Size, []string{"color"}), Apple: h.AppleProbe("ribbon", "ribbon.link")},
				{Target: link, Props: []string{"text-decoration-line"}, Golden: "underlined where apple.com's is not: a link in a sentence told apart by more than color (axe-core's link-in-text-block)"},
				{Target: link, State: []string{"hover"}, Props: []string{"color"}, Apple: h.AppleProbe("ribbon", "ribbon.link", "hover")},
				{Target: ".cui-ribbon-content", Props: []string{"animation-name", "transform"}, Golden: "still as the page shows it: apple.com's AirPods page drops it in on the accent and turns it to its ground 1.6s on (ribbon:ribbon.drop, ribbon.ground), an entrance that only draws the eye (DESIGN.md decision 25)"},
			},
		},
	})
}
