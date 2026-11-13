package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func consolePage() console.ConsolePage {
	return console.ConsolePage{Lang: "en", Title: "Jobs", Head: nil, Scheme: console.ConsoleSchemeSystem, Children: h.P(g.Text("x"))}
}

func TestConsolePageWritesTheDocumentItsHeadAfterTheStylesheet(t *testing.T) {
	p := consolePage()
	p.Scheme = console.ConsoleSchemeLight
	p.Head = h.Link(h.Rel("stylesheet"), h.Href("/own.css"))
	want := `<!doctype html><html lang="en" class="cui-light"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Jobs</title><link rel="stylesheet" href="` + ciderui.ConsoleStylesheet() + `"><link rel="stylesheet" href="/own.css"></head><body class="cui-console-page"><p>x</p></body></html>`
	if got := example.Render(p); got != want {
		t.Errorf("the page is not the document, its <html> taking the scheme's class, its head after the console stylesheet, its <body> cui-console-page:\n got: %s\nwant: %s", got, want)
	}
}

func TestConsolePagePanics(t *testing.T) {
	example.Panics(t, consolePage, []example.Break[console.ConsolePage]{
		{Want: "console.ConsolePage needs a Lang", Change: func(p *console.ConsolePage) { p.Lang = "" }},
		{Want: "console.ConsolePage needs a Title", Change: func(p *console.ConsolePage) { p.Title = "" }},
	})
}

func TestConsolePageWritesItsScheme(t *testing.T) {
	for _, c := range []struct {
		scheme console.ConsoleScheme
		html   string
	}{
		{console.ConsoleSchemeSystem, `<html lang="en">`},
		{console.ConsoleSchemeLight, `<html lang="en" class="cui-light">`},
		{console.ConsoleSchemeDark, `<html lang="en" class="cui-dark">`},
	} {
		p := consolePage()
		p.Scheme = c.scheme
		if got := example.Render(p); !strings.Contains(got, c.html) {
			t.Errorf("the page does not write %s: %s", c.html, got)
		}
	}
}
