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

func choiceMenu() console.ConsoleChoiceMenu {
	return console.ConsoleChoiceMenu{ID: "subtitles", Title: "Subtitles", Class: "", Choices: []console.ConsoleChoice{
		{Label: "Off", Value: "off", Lang: "", Checked: false},
		{Label: "English", Value: "en", Lang: "en", Checked: true},
		{Label: "日本語", Value: "ja", Lang: "ja", Checked: false},
	}}
}

func TestChoiceMenuRendersTheExample(t *testing.T) {
	page := h.Div(h.Class("flex justify-center"),
		h.Button(h.Type("button"), h.Class("cui-console-btn"), h.PopoverTarget("subtitles"), g.Text("Subtitles")),
		choiceMenu(),
	)
	example.Check(t, page, "console/console-menu/examples/choices.html")
}

func TestChoiceMenuPanics(t *testing.T) {
	for _, c := range []struct {
		want   string
		change func(*console.ConsoleChoiceMenu)
	}{
		{"ConsoleChoiceMenu needs an ID", func(m *console.ConsoleChoiceMenu) { m.ID = "" }},
		{"ConsoleChoiceMenu needs a Title", func(m *console.ConsoleChoiceMenu) { m.Title = "" }},
		{"ConsoleChoiceMenu needs a choice", func(m *console.ConsoleChoiceMenu) { m.Choices = nil }},
		{"ConsoleChoice needs a Label and a Value", func(m *console.ConsoleChoiceMenu) { m.Choices[0].Label = "" }},
		{"ConsoleChoice needs a Label and a Value", func(m *console.ConsoleChoiceMenu) { m.Choices[0].Value = "" }},
		{"needs one choice checked", func(m *console.ConsoleChoiceMenu) { m.Choices[0].Checked = true }},
		{"needs one choice checked", func(m *console.ConsoleChoiceMenu) { m.Choices[1].Checked = false }},
	} {
		m := choiceMenu()
		c.change(&m)
		mustPanic(t, c.want, func() { example.Render(m) })
	}
}
