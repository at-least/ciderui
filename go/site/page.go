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
	Head g.Node
	// Class: the <html>'s.
	Class string
	// Skip: the skip link's words, the link before all else the keyboard
	// reads; "" for none. It lands on the page's own <main id="main"
	// tabindex="-1"> (the example skip's).
	Skip     string
	Children g.Node
}

func (p Page) Render(w io.Writer) error {
	children := p.Children
	if p.Skip != "" {
		children = g.Group{Btn{Href: "#main", Label: p.Skip, Icon: nil, Submit: false, Variant: BtnVariantPrimary, Size: BtnSizeMedium, Block: false, Attrs: nil, Class: "cui-skip"}, children}
	}
	return document.Document{Owner: "site.Page", Lang: p.Lang, Title: p.Title, Stylesheet: ciderui.SiteStylesheet(), Head: p.Head, Class: p.Class, BodyClass: "", Children: children}.Render(w)
}
