package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func consoleContent() console.ConsoleContent {
	return console.ConsoleContent{Sidebar: nil, Back: nil, Title: "Jobs", Detail: "", Actions: nil, Children: nil}
}

func consolePage() console.ConsolePage {
	return console.ConsolePage{Lang: "en", Head: nil, Scheme: console.ConsoleSchemeSystem, Content: consoleContent(), Overlays: h.P(g.Text("x"))}
}

func TestConsolePageWritesTheDocumentItsHeadAfterTheStylesheet(t *testing.T) {
	p := consolePage()
	p.Scheme = console.ConsoleSchemeLight
	p.Head = h.Link(h.Rel("stylesheet"), h.Href("/own.css"))
	want := `<!doctype html><html lang="en" class="cui-light"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Jobs</title><link rel="stylesheet" href="` + ciderui.ConsoleStylesheet() + `"><link rel="stylesheet" href="/own.css"></head><body class="cui-console-page">` + example.Render(consoleContent()) + `<p>x</p></body></html>`
	if got := example.Render(p); got != want {
		t.Errorf("the page is not the document, its <title> its content's Title, its <html> taking the scheme's class, its head after the console stylesheet, its <body> cui-console-page holding its content then its overlays:\n got: %s\nwant: %s", got, want)
	}
}

// A tab names the page and, on a page with a sidebar, the site its wordmark
// names; a page titled as its site says it once (decision 57).
func TestConsolePageTitleAddsTheSidebarsWordmark(t *testing.T) {
	sidebar := console.ConsoleSidebar{AriaLabel: "Main", Home: "/", Wordmark: "CrashCart", SkipLabel: "Skip", OpenAriaLabel: "Menu", CloseAriaLabel: "Close", Footer: nil, Current: "/jobs", Sections: []console.ConsoleMenuSection{
		{Title: "", Items: []console.ConsoleMenuItem{{Href: "/jobs", Label: "Jobs", Icon: lucide.Icon{}}}},
	}}
	for _, c := range []struct {
		sidebar *console.ConsoleSidebar
		title   string
		want    string
	}{
		{nil, "Jobs", "<title>Jobs</title>"},
		{&sidebar, "Jobs", "<title>Jobs - CrashCart</title>"},
		{&sidebar, "CrashCart", "<title>CrashCart</title>"},
	} {
		content := consoleContent()
		content.Sidebar, content.Title = c.sidebar, c.title
		p := consolePage()
		p.Content = content
		if got := example.Render(p); !strings.Contains(got, c.want) {
			t.Errorf("the page does not write %s: %s", c.want, got)
		}
	}
}

func TestConsolePagePanics(t *testing.T) {
	example.Panics(t, consolePage, []example.Break[console.ConsolePage]{
		{Want: "console.ConsolePage needs a Lang", Change: func(p *console.ConsolePage) { p.Lang = "" }},
		{Want: "console.ConsolePage needs a Content", Change: func(p *console.ConsolePage) { p.Content = nil }},
		{Want: "console.ConsolePage.Content needs a Title", Change: func(p *console.ConsolePage) {
			content := consoleContent()
			content.Title = ""
			p.Content = content
		}},
		{Want: "console.ConsolePage.Content needs a Title", Change: func(p *console.ConsolePage) {
			p.Content = console.ConsoleCanvas{Back: nil, Title: "", Actions: nil, Children: nil}
		}},
	})
}

// A canvas page is its own layout: the body holds its bar, then its Children as
// written, no column around them, and the tab its Title alone.
func TestConsolePageOfACanvasWritesItsBarThenItsChildren(t *testing.T) {
	p := consolePage()
	p.Content = console.ConsoleCanvas{Back: nil, Title: "Player", Actions: nil, Children: h.Main(g.Text("film"))}
	want := `<title>Player</title>`
	got := example.Render(p)
	if !strings.Contains(got, want) || !strings.Contains(got, `<body class="cui-console-page"><header class="cui-console-canvas-bar"><h1 class="cui-console-h6-emph">Player</h1></header><main>film</main><p>x</p></body>`) {
		t.Errorf("the canvas page is not its title, then its bar, its children and its overlays in the body: %s", got)
	}
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

func canvasBar() console.ConsoleCanvas {
	return console.ConsoleCanvas{
		Back:  &console.ConsoleContentBack{Href: "#story", Label: "Shore Story"},
		Title: "Episode 1",
		Actions: g.Group{
			console.ConsoleSwitch{Label: "Only what needs you", Control: input(h.Name("todo"))},
			console.ConsoleBtn{Href: "#assets", Label: "Assets", Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: ""},
		},
		Children: nil,
	}
}

func TestCanvasRendersTheExample(t *testing.T) {
	example.Check(t, canvasBar(), "console/console-canvas/examples/bar.html")
}

func TestCanvasWithoutBackOrActionsWritesTheTitleAlone(t *testing.T) {
	c := canvasBar()
	c.Back, c.Actions = nil, nil
	if got, want := example.Render(c), `<header class="cui-console-canvas-bar"><h1 class="cui-console-h6-emph">Episode 1</h1></header>`; got != want {
		t.Errorf("a canvas with no Back or Actions wrote %s, not %s", got, want)
	}
}

func TestCanvasPanics(t *testing.T) {
	example.Panics(t, canvasBar, []example.Break[console.ConsoleCanvas]{
		{Want: "console.ConsoleCanvas needs a Title", Change: func(c *console.ConsoleCanvas) { c.Title = "" }},
		{Want: "console.ConsoleCanvas.Back needs an Href", Change: func(c *console.ConsoleCanvas) { c.Back.Href = "" }},
		{Want: "console.ConsoleCanvas.Back needs a Label", Change: func(c *console.ConsoleCanvas) { c.Back.Label = "" }},
	})
}
