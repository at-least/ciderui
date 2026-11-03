package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestMenuRendersTheExample(t *testing.T) {
	menu := console.ConsoleMenu{AriaLabel: "Analytics", Class: "", Children: g.Group{
		console.ConsoleMenuItem{Href: "#overview", Label: "Overview", Icon: lucide.LayoutGrid, Current: true, Disabled: false, Secondary: "", Class: ""},
		console.ConsoleMenuGroup{Label: "Acquisition", Icon: lucide.UserPlus, Open: true, Class: "", Children: g.Group{
			console.ConsoleMenuItem{Href: "#sources", Label: "Sources", Icon: icon.Icon{}, Current: false, Disabled: false, Secondary: "", Class: ""},
			console.ConsoleMenuItem{Href: "", Label: "Product Pages", Icon: icon.Icon{}, Current: false, Disabled: true, Secondary: "", Class: ""},
		}},
		console.ConsoleMenuGroup{Label: "Monetization", Icon: lucide.ChartLine, Open: false, Class: "", Children: console.ConsoleMenuItem{Href: "#sales", Label: "Sales", Icon: icon.Icon{}, Current: false, Disabled: false, Secondary: "", Class: ""}},
		h.Hr(h.Class("cui-console-menu-rule")),
		h.P(h.Class("cui-console-menu-heading"), g.Text("Projects")),
		console.ConsoleMenuItem{Href: "#shop", Label: "Shop", Icon: icon.Icon{}, Current: false, Disabled: false, Secondary: "8", Class: ""},
		console.ConsoleMenuItem{Href: "#demo", Label: "Demo Shop", Icon: icon.Icon{}, Current: false, Disabled: false, Secondary: "12", Class: ""},
	}}
	example.Check(t, menu, "console/console-menu/examples/menu.html")
}

func TestMenuItemWithNoHrefPanics(t *testing.T) {
	mustPanic(t, "needs an Href, unless it is disabled", func() {
		example.Render(console.ConsoleMenuItem{Href: "", Label: "x", Icon: icon.Icon{}, Current: false, Disabled: false, Secondary: "", Class: ""})
	})
}

func TestMenuItemDisabledWithAnHrefPanics(t *testing.T) {
	mustPanic(t, "needs an Href, unless it is disabled", func() {
		example.Render(console.ConsoleMenuItem{Href: "#x", Label: "x", Icon: icon.Icon{}, Current: false, Disabled: true, Secondary: "", Class: ""})
	})
}

func TestMenuWithNoAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleMenu needs an AriaLabel", func() {
		example.Render(console.ConsoleMenu{AriaLabel: "", Class: "", Children: g.Text("x")})
	})
}

func TestMenuItemWithNoLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleMenuItem needs a Label", func() {
		example.Render(console.ConsoleMenuItem{Href: "#x", Label: "", Icon: icon.Icon{}, Current: false, Disabled: false, Secondary: "", Class: ""})
	})
}

func TestMenuGroupWithNoLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleMenuGroup needs a Label", func() {
		example.Render(console.ConsoleMenuGroup{Label: "", Icon: icon.Icon{}, Open: false, Class: "", Children: g.Text("x")})
	})
}
