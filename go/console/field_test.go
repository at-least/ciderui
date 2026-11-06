package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func input(rest ...g.Node) func(g.Node) g.Node {
	return func(attrs g.Node) g.Node {
		return h.Input(h.Class("cui-console-field"), attrs, g.Group(rest))
	}
}

func TestSelectRendersTheExample(t *testing.T) {
	field := console.ConsoleFormField{ID: "level", Label: "Level", Help: "", Error: "", Control: func(attrs g.Node) g.Node {
		return console.ConsoleSelect{Class: "", Children: h.Select(h.Class("cui-console-field"), attrs, h.Name("level"),
			h.Option(g.Text("Any Level")),
			h.Option(g.Text("Fatal")),
			h.Option(g.Text("Error")),
		)}
	}}
	example.Check(t, field, "console/console-field/examples/select.html")
}

func TestFormFieldRendersTheFieldsExample(t *testing.T) {
	fields := h.Div(h.Class("grid max-w-80 gap-4"),
		console.ConsoleFormField{ID: "project", Label: "Project Name", Help: "Shown in alerts and on every page.", Error: "", Control: input(h.Value("Shop"))},
		console.ConsoleFormField{ID: "search", Label: "Search", Help: "", Error: "", Control: input(h.Placeholder("is:unresolved level:error"))},
		console.ConsoleFormField{ID: "url", Label: "Webhook URL", Help: "", Error: "Enter a URL that starts with https://.", Control: input(h.Value("htp:/example"))},
		console.ConsoleFormField{ID: "dsn", Label: "DSN", Help: "", Error: "", Control: input(h.Value("https://key@crashcart.test/1"), h.Disabled())},
		console.ConsoleFormField{ID: "note", Label: "Note", Help: "", Error: "", Control: func(attrs g.Node) g.Node {
			return h.Textarea(h.Class("cui-console-field"), attrs, h.Rows("3"), h.Placeholder("What did you decide?"))
		}},
	)
	example.Check(t, fields, "console/console-field/examples/fields.html")
}

func TestFormFieldItsHelpAndErrorBothDescribeIt(t *testing.T) {
	got := example.Render(console.ConsoleFormField{ID: "x", Label: "x", Help: "h", Error: "e", Control: input()})
	want := `<div class="cui-console-form-field"><label class="cui-console-label" for="x">x</label><input class="cui-console-field" id="x" aria-invalid="true" aria-describedby="x-help x-error"><p class="cui-console-help" id="x-help">h</p><p class="cui-console-error" id="x-error">e</p></div>`
	if got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestFormFieldPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleFormField {
		return console.ConsoleFormField{ID: "x", Label: "x", Help: "", Error: "", Control: input()}
	}, []example.Break[console.ConsoleFormField]{
		{Want: "console.ConsoleFormField needs an ID", Change: func(p *console.ConsoleFormField) { p.ID = "" }},
		{Want: "console.ConsoleFormField needs a Label", Change: func(p *console.ConsoleFormField) { p.Label = "" }},
		{Want: "console.ConsoleFormField needs a Control", Change: func(p *console.ConsoleFormField) { p.Control = nil }},
	})
}
