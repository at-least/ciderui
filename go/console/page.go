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
// Title, so the tab names the page its header names.
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
	return document.Document{Owner: "console.ConsolePage", Lang: p.Lang, Title: p.Content.Title, Stylesheet: ciderui.ConsoleStylesheet(), Head: p.Head, Class: p.Scheme.class, BodyClass: "cui-console-page", Children: g.Group{p.Content, p.Overlays}}.Render(w)
}
