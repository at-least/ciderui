package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestFootnoteRendersTheExample(t *testing.T) {
	example.Check(t, h.P(h.Class("cui-body"), g.Text("A full day of battery life"), site.Footnote{Href: "#footnote-1", Mark: "1", Class: ""}), "site/footnote/examples/footnote.html")
}

func TestFootnoteAnEmptyMarkPanics(t *testing.T) {
	mustPanic(t, "needs a Mark", func() {
		example.Render(site.Footnote{Href: "#f", Mark: "", Class: ""})
	})
}

func TestFootnoteTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.Footnote{Href: "#f", Mark: "1", Class: "ml-1"})
	if !strings.HasPrefix(got, `<sup class="cui-footnote ml-1">`) {
		t.Errorf("got %s", got)
	}
}
