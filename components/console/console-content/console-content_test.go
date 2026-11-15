//go:build browser

package consolecontent

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const (
	c    = ".cui-console-content"
	l    = "component-Content-legacy"
	room = "the sidebar's room dropped in the column (the toolkit's sidebar stands at the window's edge)"
)

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "content wide",
			HTML:   h.Example("content"),
			Widths: []int{1920, 1440, 1024},
			Checks: []h.Check{
				{Target: c, Props: []string{"max-width", "padding-left", "padding-right"}, TB: []string{l + "-maxWidth", l + "-xl-paddingInline", l + "-xl-paddingInline"}},
				{Target: c, Props: []string{"column-gap"}, TB: []string{"space-spacer-30"}},
				{Target: c, Props: []string{"box-left", "grid-template-columns"}, Golden: "centered at its widest; the sidebar's 225px (App Store Connect's menu: analytics menu.item) beside the page"},
				{Target: c + " .cui-console-menu-item[aria-current]", Props: []string{"box-left", "box-width"}, Golden: "the sidebar's items flush with the column, as the analytics menu's are"},
				{Target: c + " main", Props: []string{"box-left"}, Golden: "the page 30px past the sidebar's column"},
				{Target: c + " .cui-console-sidebar-header", Props: []string{"padding-left", "padding-right"}, Golden: room},
				{Target: c + "-main", Props: []string{"padding-top", "padding-bottom"}, Golden: "the page 32px under the window's top and 64px over its foot: crashcart's and totality's own (pt-8, pb-16), App Store Connect's unread"},
				{Target: c + " .cui-console-sidebar-body", Props: []string{"padding-left", "padding-right"}, Golden: room},
			},
		},
		{
			Name:   "content tablet",
			HTML:   h.Example("content"),
			Widths: []int{1023, 640},
			Checks: []h.Check{
				{Target: c, Props: []string{"padding-left", "padding-right"}, TB: []string{l + "-md-paddingInline", l + "-md-paddingInline"}},
				{Target: c, Props: []string{"display"}, Golden: "one column below 1024px: the sidebar is a sheet there"},
				{Target: c + " .cui-console-sidebar", Props: []string{"display"}, Golden: "the closed sheet"},
			},
		},
		{
			Name:   "content phone",
			HTML:   h.Example("content"),
			Widths: []int{639, 390},
			Checks: []h.Check{
				{Target: c, Props: []string{"padding-left", "padding-right"}, TB: []string{l + "-sm-paddingInline", l + "-sm-paddingInline"}},
			},
		},
		{
			Name:   "content rhythm",
			HTML:   h.Example("rhythm"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: c + "-main", Props: []string{"display", "row-gap"}, Golden: "the page's head and its sections 40px apart, the room crashcart and totality each wrote (gap-10); a grid, so a part the page hides leaves no room"},
				{Target: c + "-main > .cui-console-section:last-child", Props: []string{"box-y"}, Golden: "the second section 40px under the first"},
			},
		},
		{
			Name:   "content open",
			HTML:   h.Example("content"),
			Widths: []int{390},
			Checks: []h.Check{
				{Target: c + "-main > .cui-console-sidebar-open", Props: []string{"justify-self", "box-width"}, Golden: "the sheet's button its own size, not the column's"},
				{Target: c + "-main > header", Props: []string{"box-y"}, Golden: "40px under the button"},
			},
		},
	})
}
