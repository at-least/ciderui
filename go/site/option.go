package site

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Option is apple.com's buy-page choice, a radio drawn as a tile.
type Option struct {
	Name    string
	Value   string
	Title   string
	Detail  string
	Checked bool
	// Attrs: nodes appended to the <input> (disabled, required).
	Attrs g.Node
}

func (o Option) Render(w io.Writer) error {
	need("site.Option", "a Name", o.Name)
	need("site.Option", "a Value", o.Value)
	need("site.Option", "a Title", o.Title)
	id := o.Name + "-" + o.Value
	return h.Div(h.Class("cui-option"),
		h.Input(h.Class("cui-option-input"), h.Type("radio"), h.ID(id), h.Name(o.Name), h.Value(o.Value), g.If(o.Checked, h.Checked()), o.Attrs),
		h.Label(h.Class("cui-option-label"), h.For(id),
			h.Span(h.Class("cui-option-title"),
				g.Text(o.Title),
				g.If(o.Detail != "", h.Span(h.Class("cui-option-detail"), g.Text(o.Detail))),
			),
		),
	).Render(w)
}
