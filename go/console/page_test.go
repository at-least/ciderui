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
	return console.ConsolePage{Lang: "en", Title: "Jobs", Head: nil, Class: "", Skip: "", Children: h.P(g.Text("x"))}
}

func TestConsolePageWritesTheDocument(t *testing.T) {
	want := `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Jobs</title><link rel="stylesheet" href="` + ciderui.ConsoleStylesheet() + `"></head><body class="cui-console-page"><p>x</p></body></html>`
	if got := example.Render(consolePage()); got != want {
		t.Errorf("got %s\nwant %s", got, want)
	}
}

func TestConsolePageWritesItsClasses(t *testing.T) {
	p := consolePage()
	p.Class = "cui-light"
	got := example.Render(p)
	for _, want := range []string{`<html lang="en" class="cui-light">`, `<body class="cui-console-page">`} {
		if !strings.Contains(got, want) {
			t.Errorf("no %s in %s", want, got)
		}
	}
}

func TestConsolePageWritesItsHeadAfterTheKitsStylesheet(t *testing.T) {
	p := consolePage()
	p.Head = h.Link(h.Rel("stylesheet"), h.Href("/own.css"))
	got := example.Render(p)
	want := `<link rel="stylesheet" href="` + ciderui.ConsoleStylesheet() + `"><link rel="stylesheet" href="/own.css">`
	if !strings.Contains(got, want) {
		t.Errorf("no %s in %s", want, got)
	}
}

func TestConsolePageWritesItsSkipLinkFirst(t *testing.T) {
	p := consolePage()
	p.Skip = "Skip to Content"
	p.Children = h.Main(h.ID("main"), h.TabIndex("-1"))
	got := example.Render(p)
	want := `<body class="cui-console-page"><a class="cui-console-btn cui-console-skip" data-variant="filled" href="#main">Skip to Content</a><main id="main" tabindex="-1">`
	if !strings.Contains(got, want) {
		t.Errorf("no %s in %s", want, got)
	}
}

func TestConsolePagePanics(t *testing.T) {
	example.Panics(t, consolePage, []example.Break[console.ConsolePage]{
		{Want: "console.ConsolePage needs a Lang", Change: func(p *console.ConsolePage) { p.Lang = "" }},
		{Want: "console.ConsolePage needs a Title", Change: func(p *console.ConsolePage) { p.Title = "" }},
	})
}
