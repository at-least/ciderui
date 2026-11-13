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
		{Column: "subtitles", Title: "Subtitles", Options: []console.ConsoleChoiceOption{
			{Label: "Off", Value: "off", Lang: "", Checked: false},
			{Label: "English", Value: "en", Lang: "en", Checked: true},
			{Label: "日本語", Value: "ja", Lang: "ja", Checked: false},
		}},
	}}
}

// @external: two columns, the voices beside the subtitles, as the Apple TV app's
func columnsMenu() console.ConsoleChoiceMenu {
	m := choiceMenu()
	m.Columns = append([]console.ConsoleChoiceColumn{{Column: "voices", Title: "Voices", Options: []console.ConsoleChoiceOption{
		{Label: "English", Value: "en", Lang: "en", Checked: true},
		{Label: "日本語", Value: "ja", Lang: "ja", Checked: false},
	}}}, m.Columns...)
	return m
}

func TestAPageWithNoHrefIsNotAvailableEvenIfCurrent(t *testing.T) {
	for _, c := range []struct {
		part string
		node g.Node
		want string
	}{
		{"ConsoleMenuItem", console.ConsoleMenuItem{Href: "", Label: "x", Icon: lucide.Icon{}, Current: true}, `<a class="cui-console-menu-item" aria-disabled="true"><span class="cui-console-menu-label">x</span></a>`},
		{"ConsoleTabs", console.ConsoleTabs{AriaLabel: "Project", Tabs: []console.ConsoleTab{{Label: "Jobs", Href: "", Current: true}}}, `<nav class="cui-console-tabs" aria-label="Project"><a class="cui-console-tab" aria-disabled="true">Jobs</a></nav>`},
	} {
		if got := example.Render(c.node); got != c.want {
			t.Errorf("%s: a page with no Href is a link or marked current, not unavailable:\n got: %s\nwant: %s", c.part, got, c.want)
		}
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
		{Want: "console.ConsoleChoiceColumn needs a Column and a Title", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Column = "" }},
		{Want: "console.ConsoleChoiceColumn needs a Column and a Title", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Title = "" }},
		{Want: `console.ConsoleChoiceMenu has two columns "voices": each needs a Column of its own`, Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Column = "voices" }},
		{Want: "console.ConsoleChoiceColumn needs an Option", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options = nil }},
		{Want: "console.ConsoleChoiceOption needs a Label and a Value", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options[0].Label = "" }},
		{Want: "console.ConsoleChoiceOption needs a Label and a Value", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options[0].Value = "" }},
		{Want: "console.ConsoleChoiceColumn needs one Checked option, has 2", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options[0].Checked = true }},
		{Want: "console.ConsoleChoiceColumn needs one Checked option, has 0", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options[1].Checked = false }},
	})
}
