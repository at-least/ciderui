package console

import (
	"io"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/internal/document"
	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
)

// ConsoleScheme is the page's colors: the zero value follows the system's
// appearance (decision 35); the others hold the page to one, its <html>'s
// class.
type ConsoleScheme struct{ class string }

var (
	ConsoleSchemeSystem = ConsoleScheme{class: ""}
	ConsoleSchemeLight  = ConsoleScheme{class: "cui-light"}
	ConsoleSchemeDark   = ConsoleScheme{class: "cui-dark"}
)

// ConsolePage is a whole document, its <body> cui-console-page, linking the set's
// stylesheet: mount [ciderui.Assets] for it. Its <title> is its Content's
// Title, then its sidebar's Wordmark when it has one ("Jobs - Acme", as
// apple.com writes "iPhone - Apple"), so a tab names the page its header
// names and the site it belongs to (decision 57).
type ConsolePage struct {
	// Lang: the document's language, "zh-Hant".
	Lang string
	// Head: what the page's head adds after the kit's stylesheet (a site's own
	// utilities' stylesheet, a script).
	Head    g.Node
	Scheme  ConsoleScheme
	Content ConsoleContent
	// Overlays: what the page shows over its column, written after it in
	// <body>, outside the <main> a page's script may swap: its <dialog>s, a
	// [ConsoleViewer].
	Overlays g.Node
}

func (p ConsolePage) Render(w io.Writer) error {
	need.Field("console.ConsolePage.Content", "a Title", p.Content.Title)
	title := p.Content.Title
	if p.Content.Sidebar != nil && p.Content.Sidebar.Wordmark != title {
		title += " - " + p.Content.Sidebar.Wordmark
	}
	return document.Document{Owner: "console.ConsolePage", Lang: p.Lang, Title: title, Stylesheet: ciderui.ConsoleStylesheet(), Head: p.Head, Class: p.Scheme.class, BodyClass: "cui-console-page", Children: g.Group{p.Content, p.Overlays}}.Render(w)
}
