package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestMenuRendersTheExample(t *testing.T) {
	menu := console.ConsoleMenu{AriaLabel: "Analytics", Class: "", Children: g.Group{
		console.ConsoleMenuItem{Href: "#overview", Label: "Overview", Icon: lucide.LayoutGrid, Current: true, Secondary: "", Class: ""},
		console.ConsoleMenuGroup{Label: "Acquisition", Icon: lucide.UserPlus, Open: true, Class: "", Items: []console.ConsoleMenuItem{
			{Href: "#sources", Label: "Sources", Icon: lucide.Icon{}, Current: false, Secondary: "", Class: ""},
			{Href: "", Label: "Product Pages", Icon: lucide.Icon{}, Current: false, Secondary: "", Class: ""},
		}},
		console.ConsoleMenuGroup{Label: "Monetization", Icon: lucide.ChartLine, Open: false, Class: "", Items: []console.ConsoleMenuItem{{Href: "#sales", Label: "Sales", Icon: lucide.Icon{}, Current: false, Secondary: "", Class: ""}}},
		h.Hr(h.Class("cui-console-menu-rule")),
		h.P(h.Class("cui-console-menu-heading"), g.Text("Projects")),
		console.ConsoleMenuItem{Href: "#shop", Label: "Shop", Icon: lucide.Icon{}, Current: false, Secondary: "8", Class: ""},
		console.ConsoleMenuItem{Href: "#demo", Label: "Demo Shop", Icon: lucide.Icon{}, Current: false, Secondary: "12", Class: ""},
	}}
	example.Check(t, menu, "console/console-menu/examples/menu.html")
}

func choiceMenu() console.ConsoleChoiceMenu {
	return console.ConsoleChoiceMenu{ID: "subtitles", Class: "", Columns: []console.ConsoleChoiceColumn{
		{Name: "subtitles", Title: "Subtitles", Choices: []console.ConsoleChoice{
			{Label: "Off", Value: "off", Lang: "", Checked: false},
			{Label: "English", Value: "en", Lang: "en", Checked: true},
			{Label: "日本語", Value: "ja", Lang: "ja", Checked: false},
		}},
	}}
}

// two columns, the voices beside the subtitles, as the Apple TV app's
func columnsMenu() console.ConsoleChoiceMenu {
	m := choiceMenu()
	m.Columns = append([]console.ConsoleChoiceColumn{{Name: "voices", Title: "Voices", Choices: []console.ConsoleChoice{
		{Label: "English", Value: "en", Lang: "en", Checked: true},
		{Label: "日本語", Value: "ja", Lang: "ja", Checked: false},
	}}}, m.Columns...)
	return m
}

func TestChoiceMenuRendersTheExample(t *testing.T) {
	page := h.Div(h.Class("flex justify-center"),
		h.Button(h.Type("button"), h.Class("cui-console-btn"), h.PopoverTarget("subtitles"), g.Text("Subtitles")),
		choiceMenu(),
	)
	example.Check(t, page, "console/console-menu/examples/choices.html")
	page = h.Div(h.Class("flex justify-center"),
		h.Button(h.Type("button"), h.Class("cui-console-btn"), h.PopoverTarget("subtitles"), g.Text("Audio and Subtitles")),
		columnsMenu(),
	)
	example.Check(t, page, "console/console-menu/examples/columns.html")
}

// not the link a current item is: no aria-current on what cannot be followed
func TestMenuItemWithNoHrefIsNotAvailable(t *testing.T) {
	got := example.Render(console.ConsoleMenuItem{Href: "", Label: "x", Icon: lucide.Icon{}, Current: true, Secondary: "", Class: ""})
	if want := `<a class="cui-console-menu-item" aria-disabled="true"><span class="cui-console-menu-label">x</span></a>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestMenuPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleMenu {
		return console.ConsoleMenu{AriaLabel: "x", Class: "", Children: g.Text("x")}
	}, []example.Break[console.ConsoleMenu]{
		{Want: "console.ConsoleMenu needs an AriaLabel", Change: func(p *console.ConsoleMenu) { p.AriaLabel = "" }},
	})
}

func validMenuItem() console.ConsoleMenuItem {
	return console.ConsoleMenuItem{Href: "#x", Label: "x", Icon: lucide.Icon{}, Current: false, Secondary: "", Class: ""}
}

func TestMenuItemPanics(t *testing.T) {
	example.Panics(t, validMenuItem, []example.Break[console.ConsoleMenuItem]{
		{Want: "console.ConsoleMenuItem needs a Label", Change: func(p *console.ConsoleMenuItem) { p.Label = "" }},
	})
}

func TestMenuGroupPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleMenuGroup {
		return console.ConsoleMenuGroup{Label: "x", Icon: lucide.Icon{}, Open: false, Class: "", Items: []console.ConsoleMenuItem{validMenuItem()}}
	}, []example.Break[console.ConsoleMenuGroup]{
		{Want: "console.ConsoleMenuGroup needs a Label", Change: func(p *console.ConsoleMenuGroup) { p.Label = "" }},
		{Want: "console.ConsoleMenuGroup needs an Item", Change: func(p *console.ConsoleMenuGroup) { p.Items = nil }},
	})
}

func TestChoiceMenuPanics(t *testing.T) {
	example.Panics(t, columnsMenu, []example.Break[console.ConsoleChoiceMenu]{
		{Want: "console.ConsoleChoiceMenu needs an ID", Change: func(p *console.ConsoleChoiceMenu) { p.ID = "" }},
		{Want: "console.ConsoleChoiceMenu needs a column", Change: func(p *console.ConsoleChoiceMenu) { p.Columns = nil }},
		{Want: "console.ConsoleChoiceColumn needs a Name and a Title", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Name = "" }},
		{Want: "console.ConsoleChoiceColumn needs a Name and a Title", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Title = "" }},
		{Want: "console.ConsoleChoiceMenu's columns need names of their own", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Name = "voices" }},
		{Want: "console.ConsoleChoiceColumn needs a choice", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices = nil }},
		{Want: "console.ConsoleChoice needs a Label and a Value", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices[0].Label = "" }},
		{Want: "console.ConsoleChoice needs a Label and a Value", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices[0].Value = "" }},
		{Want: "console.ConsoleChoiceColumn needs one choice checked", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices[0].Checked = true }},
		{Want: "console.ConsoleChoiceColumn needs one choice checked", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices[1].Checked = false }},
	})
}
