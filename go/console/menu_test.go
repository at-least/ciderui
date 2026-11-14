package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestChoiceMenuRendersTheExamples(t *testing.T) {
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
	return console.ConsoleChoiceMenu{Columns: []console.ConsoleChoiceColumn{
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

func TestChoiceMenuPanics(t *testing.T) {
	example.Panics(t, columnsMenu, []example.Break[console.ConsoleChoiceMenu]{
		{Want: "console.ConsoleChoiceMenu needs Columns", Change: func(p *console.ConsoleChoiceMenu) { p.Columns = nil }},
		{Want: "console.ConsoleChoiceColumn needs a Column and a Title", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Column = "" }},
		{Want: "console.ConsoleChoiceColumn needs a Column and a Title", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Title = "" }},
		{Want: `console.ConsoleChoiceMenu has two columns "voices": each needs a Column of its own`, Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Column = "voices" }},
		{Want: "console.ConsoleChoiceColumn needs Options", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options = nil }},
		{Want: "console.ConsoleChoiceOption needs a Label and a Value", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options[0].Label = "" }},
		{Want: "console.ConsoleChoiceOption needs a Label and a Value", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options[0].Value = "" }},
		{Want: "console.ConsoleChoiceColumn needs one Checked option, has 2", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options[0].Checked = true }},
		{Want: "console.ConsoleChoiceColumn needs one Checked option, has 0", Change: func(p *console.ConsoleChoiceMenu) { p.Columns[1].Options[1].Checked = false }},
	})
}
