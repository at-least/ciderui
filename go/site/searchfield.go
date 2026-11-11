package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// SearchField is apple.com's search page's field: a magnifier before the
// text, its Label shown while it is empty, and a clear link after the text
// once it holds some. Its input's id is its Name; the page writes the <form>.
type SearchField struct {
	Name  string
	Label string
	Value string
	// ClearHref: where the clear link goes, the search with nothing in it. A
	// link, not a reset button: a reset puts back the Value the page wrote.
	ClearHref      string
	ClearAriaLabel string
	// Attrs: nodes appended to the <input> (autofocus, a page script's data-*).
	Attrs g.Node
}

func (f SearchField) Render(w io.Writer) error {
	if f.Name == "" {
		panic("site.SearchField needs a Name")
	}
	if f.Label == "" {
		panic("site.SearchField needs a Label")
	}
	if f.ClearHref == "" || f.ClearAriaLabel == "" {
		panic("site.SearchField needs a ClearHref and a ClearAriaLabel")
	}
	return h.Div(h.Class("cui-field"),
		h.Input(h.Class("cui-field-input"), h.ID(f.Name), h.Type("search"), h.Name(f.Name), h.Placeholder(f.Label), h.Aria("label", f.Label), h.AutoComplete("off"),
			attr.Optional("value", f.Value),
			f.Attrs,
		),
		glyph.Magnifier.Class("cui-field-icon"),
		h.A(h.Class("cui-field-clear"), h.Href(f.ClearHref), h.Aria("label", f.ClearAriaLabel), glyph.Reset),
	).Render(w)
}
