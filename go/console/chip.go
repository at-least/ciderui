package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleChip is a key and its value. A filter the page applies gets
// RemoveHref; a value that opens a list filtered by it gets Href.
type ConsoleChip struct {
	Key   string
	Value string
	// Href is not allowed with RemoveHref: a link inside a link is invalid
	// HTML.
	Href       string
	RemoveHref string
	// RemoveLabel is the remove link's aria-label: the link is an icon alone.
	RemoveLabel string
	Class       string
}

func (c ConsoleChip) Render(w io.Writer) error {
	if c.Key == "" {
		panic("console.ConsoleChip needs a Key")
	}
	if c.Href != "" && c.RemoveHref != "" {
		panic("console.ConsoleChip is a link or has a remove control, not both")
	}
	if c.RemoveHref != "" && c.RemoveLabel == "" {
		panic("console.ConsoleChip needs a RemoveLabel with a RemoveHref")
	}
	parts := g.Group{h.Span(h.Class("cui-console-chip-key"), g.Text(c.Key)), g.Text(c.Value)}
	if c.Href != "" {
		return h.A(class.Root("cui-console-chip", c.Class), h.Href(c.Href), parts).Render(w)
	}
	return h.Span(class.Root("cui-console-chip", c.Class),
		parts,
		g.If(c.RemoveHref != "", h.A(h.Class("cui-console-chip-remove"), h.Href(c.RemoveHref), h.Aria("label", c.RemoveLabel), lucide.X)),
	).Render(w)
}
