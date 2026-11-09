package site

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Option is apple.com's buy-page choice, a radio drawn as a tile.
type Option struct {
	ID      string
	Name    string
	Value   string
	Title   string
	Detail  string
	Checked bool
	// Attrs: nodes appended to the <input> (disabled, required); nil for none.
	Attrs g.Node
}

func (o Option) Render(w io.Writer) error {
	if o.ID == "" {
		panic("site.Option needs an ID")
	}
	if o.Name == "" {
		panic("site.Option needs a Name")
	}
	if o.Title == "" {
		panic("site.Option needs a Title")
	}
	return h.Div(h.Class("cui-option"),
		h.Input(h.Class("cui-option-input"), h.Type("radio"), h.ID(o.ID), h.Name(o.Name), h.Value(o.Value), g.If(o.Checked, h.Checked()), o.Attrs),
		h.Label(h.Class("cui-option-label"), h.For(o.ID),
			h.Span(h.Class("cui-option-title"),
				g.Text(o.Title),
				g.If(o.Detail != "", h.Span(h.Class("cui-option-detail"), g.Text(o.Detail))),
			),
		),
	).Render(w)
}
