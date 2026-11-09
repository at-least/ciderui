package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestMenuRendersTheExamples(t *testing.T) {
	item := func(href, label string, icon lucide.Icon, current bool) console.ConsoleMenuItem {
		return console.ConsoleMenuItem{Href: href, Label: label, Icon: icon, Current: current}
	}
	menu := h.Nav(h.Class("cui-console-menu"), h.Aria("label", "Analytics"),
		item("#overview", "Overview", lucide.LayoutGrid, true),
		item("#sources", "Sources", lucide.Icon{}, false),
		item("", "Product Pages", lucide.Icon{}, false),
		h.Hr(h.Class("cui-console-menu-rule")),
		h.P(h.Class("cui-console-menu-title"), g.Text("Projects")),
		item("#shop", "Shop", lucide.Icon{}, false),
		item("#demo", "Demo Shop", lucide.Icon{}, false),
	)
	example.Check(t, menu, "console/console-menu/examples/menu.html")
	choices := h.Div(h.Class("flex justify-center"),
		h.Button(h.Type("button"), h.Class("cui-console-btn"), h.PopoverTarget("subtitles"), g.Text("Subtitles")),
		choiceMenu(),
	)
	example.Check(t, choices, "console/console-menu/examples/choices.html")
	columns := h.Div(h.Class("flex justify-center"),
		h.Button(h.Type("button"), h.Class("cui-console-btn"), h.PopoverTarget("subtitles"), g.Text("Audio and Subtitles")),
		columnsMenu(),
	)
	example.Check(t, columns, "console/console-menu/examples/columns.html")
}

func choiceMenu() console.ConsoleChoiceMenu {
	return console.ConsoleChoiceMenu{ID: "subtitles", Columns: []console.ConsoleChoiceColumn{
		{Name: "subtitles", Title: "Subtitles", Choices: []console.ConsoleChoice{
			{Label: "Off", Value: "off", Lang: "", Checked: false},
			{Label: "English", Value: "en", Lang: "en", Checked: true},
			{Label: "日本語", Value: "ja", Lang: "ja", Checked: false},
		}},
	}}
}

// @external: two columns, the voices beside the subtitles, as the Apple TV app's
func columnsMenu() console.ConsoleChoiceMenu {
	m := choiceMenu()
	m.Columns = append([]console.ConsoleChoiceColumn{{Name: "voices", Title: "Voices", Choices: []console.ConsoleChoice{
		{Label: "English", Value: "en", Lang: "en", Checked: true},
		{Label: "日本語", Value: "ja", Lang: "ja", Checked: false},
	}}}, m.Columns...)
	return m
}

// not the link a current item is: no aria-current on what cannot be followed
func TestMenuItemWithNoHrefIsNotAvailable(t *testing.T) {
	got := example.Render(console.ConsoleMenuItem{Href: "", Label: "x", Icon: lucide.Icon{}, Current: true})
	if want := `<a class="cui-console-menu-item" aria-disabled="true"><span class="cui-console-menu-label">x</span></a>`; got != want {
		t.Errorf("an item with no Href is a link or marked current, not unavailable:\n got: %s\nwant: %s", got, want)
	}
}

func validMenuItem() console.ConsoleMenuItem {
	return console.ConsoleMenuItem{Href: "#x", Label: "x", Icon: lucide.Icon{}, Current: false}
}

func TestMenuItemPanics(t *testing.T) {
	example.Panics(t, validMenuItem, []example.Break[console.ConsoleMenuItem]{
		{Want: "console.ConsoleMenuItem needs a Label", Change: func(p *console.ConsoleMenuItem) { p.Label = "" }},
	})
}

func TestChoiceMenuPanics(t *testing.T) {
	example.Panics(t, columnsMenu, []example.Break[console.ConsoleChoiceMenu]{
		{Want: "console.ConsoleChoiceMenu needs an ID", Change: func(p *console.ConsoleChoiceMenu) { p.ID = "" }},
		{Want: "console.ConsoleChoiceMenu needs a column", Change: func(p *console.ConsoleChoiceMenu) { p.Columns = nil }},
		{Want: "console.ConsoleChoiceColumn needs a Name and a Title", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Name = "" }},
		{Want: "console.ConsoleChoiceColumn needs a Name and a Title", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Title = "" }},
		{Want: `console.ConsoleChoiceMenu has two columns named "voices": each needs a name of its own`, Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Name = "voices" }},
		{Want: "console.ConsoleChoiceColumn needs a choice", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices = nil }},
		{Want: "console.ConsoleChoice needs a Label and a Value", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices[0].Label = "" }},
		{Want: "console.ConsoleChoice needs a Label and a Value", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices[0].Value = "" }},
		{Want: "console.ConsoleChoiceColumn needs one choice checked", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices[0].Checked = true }},
		{Want: "console.ConsoleChoiceColumn needs one choice checked", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Choices[1].Checked = false }},
	})
}
