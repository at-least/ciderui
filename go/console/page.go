package console

import (
	"io"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/internal/document"
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
// stylesheet: mount [ciderui.Assets] for it.
type ConsolePage struct {
	// Lang: the document's language, "zh-Hant".
	Lang  string
	Title string
	// Head: what the page's head adds after the kit's stylesheet (a site's own
	// utilities' stylesheet, a script).
	Head     g.Node
	Scheme   ConsoleScheme
	Children g.Node
}

func (p ConsolePage) Render(w io.Writer) error {
	return document.Document{Owner: "console.ConsolePage", Lang: p.Lang, Title: p.Title, Stylesheet: ciderui.ConsoleStylesheet(), Head: p.Head, Class: p.Scheme.class, BodyClass: "cui-console-page", Children: p.Children}.Render(w)
}
