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
	return console.ConsolePage{Lang: "en", Title: "Jobs", Head: nil, Class: "", Children: h.P(g.Text("x"))}
}

func TestConsolePageWritesTheDocumentItsHeadAfterTheStylesheet(t *testing.T) {
	p := consolePage()
	p.Class = "cui-light"
	p.Head = h.Link(h.Rel("stylesheet"), h.Href("/own.css"))
	want := `<!doctype html><html lang="en" class="cui-light"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Jobs</title><link rel="stylesheet" href="` + ciderui.ConsoleStylesheet() + `"><link rel="stylesheet" href="/own.css"></head><body class="cui-console-page"><p>x</p></body></html>`
	if got := example.Render(p); got != want {
		t.Errorf("the page is not the document, its <html> taking Class, its head after the console stylesheet, its <body> cui-console-page:\n got: %s\nwant: %s", got, want)
	}
}

func TestConsolePageWithASidebarWritesItsSkipLinkFirst(t *testing.T) {
	p := consolePage()
	sidebar := console.ConsoleSidebar{AriaLabel: "Main", Home: "/", Wordmark: "Jobs", SkipLabel: "Skip to Content", OpenAriaLabel: "Menu", CloseAriaLabel: "Close", Footer: nil, Children: nil}
	p.Children = console.ConsoleContent{Sidebar: &sidebar, Back: console.ConsoleContentBack{Href: "", Label: ""}, Title: "Jobs", Detail: "", Controls: nil, Children: nil}
	got := example.Render(p)
	want := `<body class="cui-console-page"><div class="cui-console-content"><a class="cui-console-btn cui-console-skip" data-variant="filled" href="#main">Skip to Content</a><nav class="cui-console-sidebar"`
	if !strings.Contains(got, want) {
		t.Errorf("the sidebar's skip link to #main is not the page's first element: no %s in %s", want, got)
	}
}

func TestConsolePagePanics(t *testing.T) {
	example.Panics(t, consolePage, []example.Break[console.ConsolePage]{
		{Want: "console.ConsolePage needs a Lang", Change: func(p *console.ConsolePage) { p.Lang = "" }},
		{Want: "console.ConsolePage needs a Title", Change: func(p *console.ConsolePage) { p.Title = "" }},
	})
}
