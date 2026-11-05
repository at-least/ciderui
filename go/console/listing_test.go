package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// the <pre> kept as written, each line outside it trimmed: example.Read
// would drop the code's indent and the newlines between its lines
func examplePre(path string) string {
	raw := readRaw(path)
	start := strings.Index(raw, "<pre")
	end := strings.Index(raw, "</pre>") + len("</pre>")
	if start < 0 || end < len("</pre>") {
		panic(path + " has no <pre>…</pre>")
	}
	return example.JoinTrimmed(raw[:start]) + raw[start:end] + example.JoinTrimmed(raw[end:])
}

func checkPre(t *testing.T, node g.Node, path string) {
	t.Helper()
	if got, want := example.Render(node), examplePre(path); got != want {
		t.Errorf("%s\n got: %s\nwant: %s", path, got, want)
	}
}

func TestListingRendersTheExamples(t *testing.T) {
	lines := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleListing{FirstLine: 40, Lines: []console.ConsoleListingLine{{Label: "func total(_ cart: Cart) -> Int {", Highlighted: false}, {Label: "    let items = cart.items", Highlighted: false}, {Label: "    return items.reduce(0) { $0 + $1.price }", Highlighted: true}, {Label: "}", Highlighted: false}}, Class: ""},
	)
	checkPre(t, lines, "console/console-listing/examples/lines.html")

	nested := h.Div(h.Class("cui-console-page max-w-xl"),
		h.Div(h.Class("cui-console-card"),
			console.ConsoleListing{FirstLine: 41, Lines: []console.ConsoleListingLine{{Label: "    let items = cart.items", Highlighted: false}, {Label: "    return items.reduce(0) { $0 + $1.price }", Highlighted: true}}, Class: ""},
		),
	)
	checkPre(t, nested, "console/console-listing/examples/nested.html")
}

func TestListingPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleListing {
		return console.ConsoleListing{FirstLine: 1, Lines: []console.ConsoleListingLine{{Label: "a", Highlighted: false}}, Class: ""}
	}, []example.Break[console.ConsoleListing]{
		{Want: "console.ConsoleListing needs a line", Change: func(p *console.ConsoleListing) { p.Lines = nil }},
	})
}
