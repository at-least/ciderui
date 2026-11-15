//go:build browser

package list

import (
	"strings"
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
)

const (
	row    = ".cui-list-row"
	second = row + ":nth-child(2)"
	title  = row + ":first-child .cui-list-title"
	text   = row + ":first-child .cui-list-detail"
	aside  = row + ":first-child .cui-list-aside"
)

var (
	size  = []string{"font-size", "font-weight"}
	clamp = []string{"-webkit-line-clamp", "-webkit-box-orient", "overflow-x", "overflow-y"}
)

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "list",
			HTML:   h.Example("list"),
			Widths: []int{1600, 1200, 900, 735},
			Checks: []h.Check{
				{Target: second, Props: []string{"padding-top", "padding-bottom", "border-top-width", "border-top-style"}, Apple: h.AppleProbe(apple.ForumsList, "row.second")},
				{Target: row, Props: []string{"column-gap"}, Apple: h.AppleProbe(apple.ForumsList, "row")},
				{Target: second, Props: []string{"padding-left", "padding-right", "border-top-color"}, Golden: "flush with the page's column, the kit's divider: the forums draw their rows inset on a white card (20px by 40px, a #d3d3d3 rule), a site's page is white already"},
				{Target: ".cui-list-row:first-child", Props: []string{"border-top-width"}, Golden: "no rule over the first row (forums-list:row has one, the card's top edge under it)"},
				{Target: title, Pseudo: "::after", Props: []string{"position", "top", "bottom", "left", "right"}, Golden: "the whole row is the title's target, as a list row is on iOS: the forums' title alone is, 21px tall"},
				{Target: aside, Props: []string{"position"}, Golden: "over the title's target, so a link in the aside is its own"},
				{Target: row, Props: []string{"grid-template-columns"}, Golden: "from 735px the row's text, then its aside as wide as it needs (forums-list:row, its stats a fixed 274px of four columns)"},
				{Target: title, Props: h.Props(size), Apple: h.AppleProbe(apple.ForumsList, "row.title")},
				{Target: title, Props: []string{"color", "line-height"}, Golden: "the kit's text, where the forums' title is black; the ladder's 21px (cui-body-tight), where theirs computes 21.0012px"},
				{Target: title, State: []string{"hover"}, Props: []string{"text-decoration-line"}, Apple: h.AppleProbe(apple.ForumsList, "row.title", "hover")},
				{Target: text, Props: h.Props(size, []string{"color", "margin-top"}, clamp), Apple: &h.Apple{Page: apple.ForumsList, Probe: "row.excerpt", As: map[string]string{"overflow-x": "overflow", "overflow-y": "overflow"}}},
				{Target: text, Props: []string{"line-height"}, Golden: "the ladder's 18px (cui-body-reduced-tight), where the forums' excerpt computes 18.0008px"},
				{Target: aside, Props: []string{"color", "line-height"}, Apple: h.AppleProbe(apple.ForumsList, "row.stat")},
				{Target: aside, Props: []string{"font-size", "text-align", "margin-top", "font-variant-numeric"}, Golden: "the forums' stats at 14px, the ladder's (they compute 14.875px), at the row's end, its figures tabular"},
			},
		},
		{
			Name:   "list phone",
			HTML:   h.Example("list"),
			Widths: []int{734, 390},
			Checks: []h.Check{
				{Target: second, Props: []string{"padding-top", "padding-bottom"}, Apple: h.AppleProbe(apple.ForumsList, "mrow")},
				{Target: row, Props: []string{"grid-template-columns"}, Golden: "one column: the aside under the text (forums-list:mrow)"},
				{Target: title, Props: h.Props(size), Apple: h.AppleProbe(apple.ForumsList, "mrow.title")},
				{Target: text, Props: []string{"margin-top", "-webkit-line-clamp"}, Apple: h.AppleProbe(apple.ForumsList, "mrow.excerpt")},
				{Target: aside, Props: []string{"margin-top"}, Apple: h.AppleProbe(apple.ForumsList, "mrow.stats")},
				{Target: aside, Props: []string{"font-size", "color", "text-align"}, Golden: "the wide row's aside, under the text: the forums' phone row draws its stats as icons at 16px, the kit draws no icon the page did not choose"},
			},
		},
		{
			Name:   "list zh",
			HTML:   strings.Replace(h.Example("list"), "<ul class", `<ul lang="zh-Hant" class`, 1),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: title, Props: []string{"line-height"}, Golden: "Traditional Chinese a step taller, as cui-body-tight is"},
				{Target: text, Props: []string{"line-height"}, Golden: "Traditional Chinese a pixel taller, as cui-body-reduced-tight is"},
				{Target: aside, Props: []string{"line-height"}, Golden: "Traditional Chinese a pixel taller, as cui-body-reduced is"},
			},
		},
	})
}
