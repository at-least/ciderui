package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleChipRemove is a filter chip's remove link, an icon alone.
type ConsoleChipRemove struct {
	Href      string
	AriaLabel string
}

// ConsoleChip is a key and its value. A filter the page applies gets Remove; a
// value that opens a list filtered by it gets Href.
type ConsoleChip struct {
	Key   string
	Value string
	// Href is not allowed with Remove: a link inside a link is invalid HTML.
	Href string
	// Remove: nil for none.
	Remove *ConsoleChipRemove
	Class  string
}

func (c ConsoleChip) Render(w io.Writer) error {
	if c.Key == "" {
		panic("console.ConsoleChip needs a Key")
	}
	if c.Href != "" && c.Remove != nil {
		panic("console.ConsoleChip is a link or has a remove control, not both")
	}
	if c.Remove != nil && (c.Remove.Href == "" || c.Remove.AriaLabel == "") {
		panic("console.ConsoleChipRemove needs an Href and an AriaLabel")
	}
	parts := g.Group{h.Span(h.Class("cui-console-chip-key"), g.Text(c.Key)), g.Text(c.Value)}
	if c.Href != "" {
		return h.A(class.Root("cui-console-chip", c.Class), h.Href(c.Href), parts).Render(w)
	}
	return h.Span(class.Root("cui-console-chip", c.Class),
		parts,
		g.Iff(c.Remove != nil, func() g.Node {
			return h.A(h.Class("cui-console-chip-remove"), h.Href(c.Remove.Href), h.Aria("label", c.Remove.AriaLabel), lucide.X)
		}),
	).Render(w)
}
