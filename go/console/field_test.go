package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSelectRendersTheExample(t *testing.T) {
	selectField := h.Div(h.Class("cui-console-form-field"),
		h.Label(h.Class("cui-console-label"), h.For("level"), g.Text("Level")),
		console.ConsoleSelect{Class: "", Children: h.Select(h.Class("cui-console-field"), h.ID("level"), h.Name("level"),
			h.Option(g.Text("Any Level")),
			h.Option(g.Text("Fatal")),
			h.Option(g.Text("Error")),
		)},
	)
	example.Check(t, selectField, "console/console-field/examples/select.html")
}

func TestFormFieldRendersTheFieldsExample(t *testing.T) {
	fields := h.Div(h.Class("grid max-w-80 gap-4"),
		console.ConsoleFormField{ID: "project", Label: g.Text("Project Name"), Children: g.Group{
			h.Input(h.Class("cui-console-field"), h.ID("project"), h.Value("Shop"), h.Aria("describedby", console.ConsoleHelpID("project"))),
			console.ConsoleHelp{ControlID: "project", Label: "Shown in alerts and on every page."},
		}},
		console.ConsoleFormField{ID: "search", Label: g.Text("Search"), Children: h.Input(h.Class("cui-console-field"), h.ID("search"), h.Placeholder("is:unresolved level:error"))},
		console.ConsoleFormField{ID: "url", Label: g.Text("Webhook URL"), Children: g.Group{
			h.Input(h.Class("cui-console-field"), h.ID("url"), h.Value("htp:/example"), h.Aria("invalid", "true"), h.Aria("describedby", console.ConsoleErrorID("url"))),
			console.ConsoleError{ControlID: "url", Label: "Enter a URL that starts with https://.", Class: ""},
		}},
		console.ConsoleFormField{ID: "dsn", Label: g.Text("DSN"), Children: h.Input(h.Class("cui-console-field"), h.ID("dsn"), h.Value("https://key@crashcart.test/1"), h.Disabled())},
		console.ConsoleFormField{ID: "note", Label: g.Text("Note"), Children: h.Textarea(h.Class("cui-console-field"), h.ID("note"), h.Rows("3"), h.Placeholder("What did you decide?"))},
	)
	example.Check(t, fields, "console/console-field/examples/fields.html")
}

func TestErrorNoLabelIsTheHiddenSlotAScriptFills(t *testing.T) {
	if got, want := example.Render(console.ConsoleError{ControlID: "", Label: "", Class: "mt-2"}), `<p class="cui-console-error mt-2" role="alert" hidden></p>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
	if got, want := example.Render(console.ConsoleError{ControlID: "url", Label: "", Class: ""}), `<p class="cui-console-error" id="url-error" role="alert" hidden></p>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestFormFieldPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleFormField {
		return console.ConsoleFormField{ID: "x", Label: g.Text("x"), Children: h.Input(h.Class("cui-console-field"))}
	}, []example.Break[console.ConsoleFormField]{
		{Want: "console.ConsoleFormField needs an ID", Change: func(p *console.ConsoleFormField) { p.ID = "" }},
		{Want: "console.ConsoleFormField needs a Label", Change: func(p *console.ConsoleFormField) { p.Label = nil }},
	})
}

func TestHelpPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleHelp {
		return console.ConsoleHelp{ControlID: "x", Label: "x"}
	}, []example.Break[console.ConsoleHelp]{
		{Want: "console.ConsoleHelp needs a Label", Change: func(p *console.ConsoleHelp) { p.Label = "" }},
	})
}

func TestErrorPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleError {
		return console.ConsoleError{ControlID: "x", Label: "x", Class: ""}
	}, []example.Break[console.ConsoleError]{
		{Want: "console.ConsoleError with a Label needs a ControlID", Change: func(p *console.ConsoleError) { p.ControlID = "" }},
	})
}
