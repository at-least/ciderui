// Package document writes the page around a set's parts.
package document

import (
	"io"
	"strings"
	"testing"

	"github.com/at-least/ciderui"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Document is the whole page; the sets' page types fill it.
type Document struct {
	// Owner: the page type a panic names.
	Owner      string
	Lang       string
	Title      string
	Stylesheet string
	// Head comes after Stylesheet: a site's own sheet there wins over a part.
	Head      g.Node
	Class     string
	BodyClass string
	Children  g.Node
}

// @invariant: under go test, a page ciderui.CheckMarkup refuses panics, so every
// test of a site that renders a page checks it, with no call of the site's own
func (d Document) Render(w io.Writer) error {
	if d.Lang == "" {
		panic(d.Owner + " needs a Lang")
	}
	if d.Title == "" {
		panic(d.Owner + " needs a Title")
	}
	page := h.Doctype(h.HTML(h.Lang(d.Lang), classAttr(d.Class),
		h.Head(
			h.Meta(h.Charset("utf-8")),
			h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1")),
			h.TitleEl(g.Text(d.Title)),
			h.Link(h.Rel("stylesheet"), h.Href(d.Stylesheet)),
			d.Head,
		),
		h.Body(classAttr(d.BodyClass), d.Children),
	))
	if !testing.Testing() {
		return page.Render(w)
	}
	var html strings.Builder
	if err := page.Render(&html); err != nil {
		return err
	}
	Check(d.Owner, html.String())
	_, err := io.WriteString(w, html.String())
	return err
}

// Check panics, naming owner, on what ciderui.CheckMarkup refuses in html: the
// one wording of a refused page and a refused example
func Check(owner, html string) {
	if problems := ciderui.CheckMarkup(html); problems != nil {
		panic(owner + ": " + strings.Join(problems, "; "))
	}
}

func classAttr(value string) g.Node {
	if value == "" {
		return nil
	}
	return h.Class(value)
}
