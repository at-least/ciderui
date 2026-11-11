package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleDisclosure is the HIG's disclosure control.
type ConsoleDisclosure struct {
	Summary string
	// Attrs: the root's other attributes, for the page's script (data-*, an
	// htmx attribute).
	Attrs    g.Node
	Children g.Node
}

func (d ConsoleDisclosure) Render(w io.Writer) error {
	need("console.ConsoleDisclosure", "a Summary", d.Summary)
	return h.Details(h.Class("cui-console-disclosure"), d.Attrs,
		h.Summary(h.Class("cui-console-disclosure-summary"), lucide.ChevronRight.Class("cui-console-disclosure-chevron"), g.Text(d.Summary)),
		d.Children,
	).Render(w)
}
