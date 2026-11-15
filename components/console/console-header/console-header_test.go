//go:build browser

package consoleheader

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	asc "github.com/at-least/ciderui/reference/asc/results"
)

const s = ".cui-console-header"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "header",
			HTML:   h.Example("page"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: s, Props: []string{"display", "align-items", "row-gap", "column-gap"}, ASC: h.ASCProbe(asc.Head, "section.head")},
				{Target: s, Props: []string{"margin-bottom"}, Golden: "none of its own: the room under it is its container's (rule 6), App Store Connect's 20px a section's (console-section)"},
				{Target: s + " > :first-child", Props: []string{"flex-grow", "min-width"}, Golden: "the title takes the room, the controls at the row's end (App Store Connect's title 941px of 1105px)"},
				{Target: s + " h1", Props: []string{"font-size", "line-height"}, Golden: "the title's type its own class's (cui-console-h2), none from the header: App Store Connect's section heading is its h4, 18px on 22px"},
				{Target: s + " > .cui-console-tabs", Props: []string{"box-y"}, Golden: "at 1200px on the title's row; at 390px under it (the row wraps; App Store Connect's does not)"},
				{Target: s + "-detail", Props: []string{"color", "margin-top", "max-width"}, Golden: "a page's line under its title: the secondary label, 8px under the title, a readable measure (the shape crashcart and totality wrote with utilities)"},
			},
		},
		{
			Name:   "header section",
			HTML:   h.Example("section"),
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: s + " h2", Props: []string{"font-size", "line-height", "font-weight"}, Golden: "a section's title its own class's (cui-console-h4-emph): App Store Connect's section heading, 18px on 22px"},
				{Target: s + "-detail", Props: []string{"color", "margin-top"}, Golden: "a section's line under its title: the secondary label, 4px under its smaller title"},
			},
		},
		{
			Name:   "header back",
			HTML:   h.Example("back"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: s + "-back", Props: []string{"display", "color", "font-size", "margin-bottom", "box-width"}, Golden: "the way back over the title, named by the page it goes to with a chevron before it (iOS's back button), a console link at the body's size, 8px over the title, as wide as its words"},
				{Target: s + "-back > svg", Props: []string{"box-width"}, Golden: "the chevron at the text's size"},
			},
		},
	})
}
