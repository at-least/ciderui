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
			console.ConsoleHelp{ID: "project", Text: "Shown in alerts and on every page."},
		}},
		console.ConsoleFormField{ID: "search", Label: g.Text("Search"), Children: h.Input(h.Class("cui-console-field"), h.ID("search"), h.Placeholder("is:unresolved level:error"))},
		console.ConsoleFormField{ID: "url", Label: g.Text("Webhook URL"), Children: g.Group{
			h.Input(h.Class("cui-console-field"), h.ID("url"), h.Value("htp:/example"), h.Aria("invalid", "true"), h.Aria("describedby", console.ConsoleErrorID("url"))),
			console.ConsoleError{ID: "url", Text: "Enter a URL that starts with https://.", Class: ""},
		}},
		console.ConsoleFormField{ID: "dsn", Label: g.Text("DSN"), Children: h.Input(h.Class("cui-console-field"), h.ID("dsn"), h.Value("https://key@crashcart.test/1"), h.Disabled())},
		console.ConsoleFormField{ID: "note", Label: g.Text("Note"), Children: h.Textarea(h.Class("cui-console-field"), h.ID("note"), h.Rows("3"), h.Placeholder("What did you decide?"))},
	)
	example.Check(t, fields, "console/console-field/examples/fields.html")
}

func TestErrorNoTextIsTheHiddenSlotAScriptFills(t *testing.T) {
	if got, want := example.Render(console.ConsoleError{ID: "", Text: "", Class: "mt-2"}), `<p class="cui-console-error mt-2" role="alert" hidden></p>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
	if got, want := example.Render(console.ConsoleError{ID: "url", Text: "", Class: ""}), `<p class="cui-console-error" id="url-error" role="alert" hidden></p>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestErrorATextWithNoControlPanics(t *testing.T) {
	mustPanic(t, "console.ConsoleError with a Text needs its control's ID", func() {
		example.Render(console.ConsoleError{ID: "", Text: "Enter a name.", Class: ""})
	})
}

func TestFormFieldAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "ConsoleFormField needs an ID", func() {
		example.Render(console.ConsoleFormField{ID: "", Label: g.Text("Name"), Children: h.Input(h.Class("cui-console-field"))})
	})
}

func TestHelpAnEmptyTextPanics(t *testing.T) {
	mustPanic(t, "ConsoleHelp needs a Text", func() {
		example.Render(console.ConsoleHelp{ID: "name", Text: ""})
	})
}
