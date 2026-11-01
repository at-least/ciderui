package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// the example at path with its listing's lines kept as written: trimming
// each line, as example.Read does, would drop the code's indent and the
// newlines between its lines; each line outside the <pre> is trimmed as
// example.Read trims it
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
		console.ConsoleListing{FirstLine: 40, Lines: []string{"func total(_ cart: Cart) -> Int {", "    let items = cart.items", "    return items.reduce(0) { $0 + $1.price }", "}"}, Highlighted: 42, Class: ""},
	)
	checkPre(t, lines, "console/console-listing/examples/lines.html")

	nested := h.Div(h.Class("cui-console-page max-w-xl"),
		h.Div(h.Class("cui-console-card"),
			console.ConsoleListing{FirstLine: 41, Lines: []string{"    let items = cart.items", "    return items.reduce(0) { $0 + $1.price }"}, Highlighted: 42, Class: ""},
		),
	)
	checkPre(t, nested, "console/console-listing/examples/nested.html")
}

func TestListingNoLinesPanic(t *testing.T) {
	mustPanic(t, "needs a line", func() {
		example.Render(console.ConsoleListing{FirstLine: 1, Lines: nil, Highlighted: 0, Class: ""})
	})
}

// not in the Rust crate, whose highlighted line is an Option: here 0 is
// none, so a first line 0 would make line 0 unhighlightable unseen
func TestListingAFirstLineOf0Panics(t *testing.T) {
	mustPanic(t, "FirstLine is at least 1", func() {
		example.Render(console.ConsoleListing{FirstLine: 0, Lines: []string{"a"}, Highlighted: 0, Class: ""})
	})
}
