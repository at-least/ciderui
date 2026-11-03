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
	trim := func(s string) string {
		var out strings.Builder
		for _, line := range strings.Split(s, "\n") {
			out.WriteString(strings.TrimSpace(line))
		}
		return out.String()
	}
	return trim(raw[:start]) + raw[start:end] + trim(raw[end:])
}

func checkPre(t *testing.T, node g.Node, path string) {
	t.Helper()
	if got, want := example.Render(node), examplePre(path); got != want {
		t.Errorf("%s\n got: %s\nwant: %s", path, got, want)
	}
}

func TestListingRendersTheExamples(t *testing.T) {
	lines := h.Div(h.Class("cui-console-page max-w-xl"),
		console.ConsoleListing{FirstLine: 40, Lines: []console.ConsoleListingLine{{Text: "func total(_ cart: Cart) -> Int {", Highlighted: false}, {Text: "    let items = cart.items", Highlighted: false}, {Text: "    return items.reduce(0) { $0 + $1.price }", Highlighted: true}, {Text: "}", Highlighted: false}}, Class: ""},
	)
	checkPre(t, lines, "console/console-listing/examples/lines.html")

	nested := h.Div(h.Class("cui-console-page max-w-xl"),
		h.Div(h.Class("cui-console-card"),
			console.ConsoleListing{FirstLine: 41, Lines: []console.ConsoleListingLine{{Text: "    let items = cart.items", Highlighted: false}, {Text: "    return items.reduce(0) { $0 + $1.price }", Highlighted: true}}, Class: ""},
		),
	)
	checkPre(t, nested, "console/console-listing/examples/nested.html")
}

func TestListingNoLinesPanic(t *testing.T) {
	mustPanic(t, "needs a line", func() {
		example.Render(console.ConsoleListing{FirstLine: 1, Lines: nil, Class: ""})
	})
}

func TestListingHighlightsEveryMarkedLine(t *testing.T) {
	got := example.Render(console.ConsoleListing{FirstLine: 1, Lines: []console.ConsoleListingLine{{Text: "a", Highlighted: true}, {Text: "b", Highlighted: false}, {Text: "c", Highlighted: true}}, Class: ""})
	want := `<pre class="cui-console-listing"><code><span class="cui-console-listing-line" data-line="1" data-highlighted>a</span>` + "\n" + `<span class="cui-console-listing-line" data-line="2">b</span>` + "\n" + `<span class="cui-console-listing-line" data-line="3" data-highlighted>c</span></code></pre>`
	if got != want {
		t.Errorf("got %s", got)
	}
}
