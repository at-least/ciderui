//go:build browser

package link

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
)

var linkProps = []string{"color", "text-decoration-line", "text-decoration-color"}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name: "inline",
			HTML: h.Example("inline"),
			Checks: []h.Check{
				{Target: ".cui-link", Props: h.Props([]string{"font-size", "line-height", "font-weight"}, linkProps), Apple: h.AppleProbe(apple.Iphone, "link.inline")},
				{Target: ".cui-link", State: []string{"hover"}, Props: linkProps, Apple: h.AppleProbe(apple.Iphone, "link.inline", "hover")},
				{Target: ".cui-link", State: []string{"active"}, Props: linkProps, Apple: h.AppleProbe(apple.Iphone, "link.inline", "active")},
			},
		},
		{
			Name: "inline-dark",
			HTML: `<section class="cui-dark">` + h.Example("inline") + `</section>`,
			Checks: []h.Check{
				{Target: ".cui-link", Props: linkProps, Apple: h.AppleProbe(apple.Product, "link.inline")},
				{Target: ".cui-link", State: []string{"hover"}, Props: linkProps, Apple: h.AppleProbe(apple.Product, "link.inline", "hover")},
				{Target: ".cui-link", State: []string{"active"}, Props: linkProps, Apple: h.AppleProbe(apple.Product, "link.inline", "active")},
			},
		},
		{
			Name:   "inline link no touch target",
			HTML:   h.Example("inline"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: ".cui-link", Pseudo: "::after", Props: []string{"content"}, Golden: "a link in running text no ::after on a touch screen (text, exempt from the HIG's control sizes)"},
			},
		},
	})
}
