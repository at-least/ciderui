package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleChip is cui-console-chip (components/console/console-chip): App Store
// Connect's newer look's chip, a key and its value on the gray fill. A
// filter the page applies has a control that removes it (RemoveHref, a link
// with a Lucide x); a value that opens a list filtered by it is a link as a
// whole (Href).
type ConsoleChip struct {
	// Key is its key, as the page words it. Required: an empty one panics.
	Key string
	// Value is its value.
	Value string
	// Href, not "": the chip is a link to this URL. Not with RemoveHref: a
	// control inside a link is no valid HTML, so both panic.
	Href string
	// RemoveHref, not "": the chip has a remove control, a link to this
	// URL.
	RemoveHref string
	// RemoveLabel is the remove control's aria-label. Required with
	// RemoveHref: the control is an icon alone.
	RemoveLabel string
	// Class is the page's utilities on the chip; "" for none.
	Class string
}

// Render writes the chip.
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
