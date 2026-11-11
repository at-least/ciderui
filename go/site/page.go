package site

import (
	"io"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/internal/document"
	g "maragu.dev/gomponents"
)

// Page is a whole document, linking the set's stylesheet: mount
// [ciderui.Assets] for it.
type Page struct {
	// Lang: the document's language, "zh-Hant".
	Lang  string
	Title string
	// Head: what the page's head adds after the kit's stylesheet (a site's own
	// utilities' stylesheet, a script); nil for none.
	Head     g.Node
	Children g.Node
}

func (p Page) Render(w io.Writer) error {
	return document.Document{Owner: "site.Page", Lang: p.Lang, Title: p.Title, Stylesheet: ciderui.SiteStylesheet(), Head: p.Head, Class: "", BodyClass: "", Children: p.Children}.Render(w)
}
