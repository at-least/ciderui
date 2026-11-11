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
		return h.Input(attrs, g.Group(rest))
	}
}

func TestSelectRendersTheExample(t *testing.T) {
	field := console.ConsoleField{ID: "level", Label: "Level", Error: "", Control: func(attrs g.Node) g.Node {
		return console.ConsoleSelect{Children: h.Select(attrs, h.Name("level"),
			h.Option(g.Text("Any Level")),
			h.Option(g.Text("Fatal")),
			h.Option(g.Text("Error")),
		)}
	}}
	example.Check(t, field, "console/console-field/examples/select.html")
}

func TestFieldRendersTheFieldsExample(t *testing.T) {
	fields := h.Div(h.Class("grid max-w-80 gap-4"),
		console.ConsoleField{ID: "project", Label: "Project Name", Error: "", Control: input(h.Value("Shop"))},
		console.ConsoleField{ID: "search", Label: "Search", Error: "", Control: input(h.Placeholder("is:unresolved level:error"))},
		console.ConsoleField{ID: "url", Label: "Webhook URL", Error: "Enter a URL that starts with https://.", Control: input(h.Value("htp:/example"))},
		console.ConsoleField{ID: "dsn", Label: "DSN", Error: "", Control: input(h.Value("https://key@crashcart.test/1"), h.Disabled())},
		console.ConsoleField{ID: "note", Label: "Note", Error: "", Control: func(attrs g.Node) g.Node {
			return h.Textarea(attrs, h.Rows("3"), h.Placeholder("What did you decide?"))
		}},
	)
	example.Check(t, fields, "console/console-field/examples/fields.html")
}

func TestFieldPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleField {
		return console.ConsoleField{ID: "x", Label: "x", Error: "", Control: input()}
	}, []example.Break[console.ConsoleField]{
		{Want: "console.ConsoleField needs an ID", Change: func(p *console.ConsoleField) { p.ID = "" }},
		{Want: "console.ConsoleField needs a Label", Change: func(p *console.ConsoleField) { p.Label = "" }},
		{Want: "console.ConsoleField needs a Control", Change: func(p *console.ConsoleField) { p.Control = nil }},
		{Want: `example.Render: class on <input class="cui-console-field-input">: written twice, and the browser keeps the first`, Change: func(p *console.ConsoleField) { p.Control = input(h.Class("cui-console-field-input")) }},
	})
}

func search() console.ConsoleSearch {
	return console.ConsoleSearch{AriaLabel: "Search Accounts", Control: input(h.Name("q"), h.Placeholder("Name, email or phone"))}
}

func TestSearchRendersTheExample(t *testing.T) {
	example.Check(t, h.Form(h.Class("max-w-80"), h.Role("search"), h.Action("#accounts"), search()), "console/console-field/examples/search.html")
}

func TestSearchPanics(t *testing.T) {
	example.Panics(t, search, []example.Break[console.ConsoleSearch]{
		{Want: "console.ConsoleSearch needs an AriaLabel", Change: func(p *console.ConsoleSearch) { p.AriaLabel = "" }},
		{Want: "console.ConsoleSearch needs a Control", Change: func(p *console.ConsoleSearch) { p.Control = nil }},
	})
}
