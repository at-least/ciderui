package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestFootnoteRendersTheExample(t *testing.T) {
	example.Check(t, h.P(h.Class("cui-body"), g.Text("A full day of battery life"), site.Footnote{Href: "#footnote-1", Mark: "1", AriaLabel: "Footnote 1", Class: ""}), "site/footnote/examples/footnote.html")
}

func TestFootnotePanics(t *testing.T) {
	example.Panics(t, func() site.Footnote {
		return site.Footnote{Href: "#f", Mark: "1", AriaLabel: "x", Class: ""}
	}, []example.Break[site.Footnote]{
		{Want: "site.Footnote needs a Mark", Change: func(p *site.Footnote) { p.Mark = "" }},
		{Want: "site.Footnote needs an AriaLabel", Change: func(p *site.Footnote) { p.AriaLabel = "" }},
	})
}
