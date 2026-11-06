// Package document writes the page around a set's parts.
package document

import (
	"io"

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

func (d Document) Render(w io.Writer) error {
	if d.Lang == "" {
		panic(d.Owner + " needs a Lang")
	}
	if d.Title == "" {
		panic(d.Owner + " needs a Title")
	}
	return h.Doctype(h.HTML(h.Lang(d.Lang), classAttr(d.Class),
		h.Head(
			h.Meta(h.Charset("utf-8")),
			h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1")),
			h.TitleEl(g.Text(d.Title)),
			h.Link(h.Rel("stylesheet"), h.Href(d.Stylesheet)),
			d.Head,
		),
		h.Body(classAttr(d.BodyClass), d.Children),
	)).Render(w)
}

func classAttr(value string) g.Node {
	if value == "" {
		return nil
	}
	return h.Class(value)
}
