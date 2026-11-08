package console

import (
	"io"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/internal/document"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
)

// ConsolePage is a whole document, its <body> cui-console-page, linking the set's
// stylesheet: mount [ciderui.Assets] for it.
type ConsolePage struct {
	// Lang: the document's language, "zh-Hant".
	Lang  string
	Title string
	// Head: what the page's head adds after the kit's stylesheet (a site's own
	// utilities' stylesheet, a script); nil for none.
	Head g.Node
	// Class: the <html>'s.
	Class string
	// SkipLabel: the skip link's words, the link before all else the keyboard reads;
	// "" for none. It lands on the page's own <main id="main" tabindex="-1">
	// (the example console-skip's).
	SkipLabel string
	Children  g.Node
}

func (p ConsolePage) Render(w io.Writer) error {
	children := p.Children
	if p.SkipLabel != "" {
		children = g.Group{
			ConsoleBtn{Href: "#main", Label: p.SkipLabel, AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: ConsoleBtnVariantFilled, Size: ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: "cui-console-skip"},
			children,
		}
	}
	return document.Document{Owner: "console.ConsolePage", Lang: p.Lang, Title: p.Title, Stylesheet: ciderui.ConsoleStylesheet(), Head: p.Head, Class: p.Class, BodyClass: "cui-console-page", Children: children}.Render(w)
}
