package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleCheckbox is a checkbox and its words. Control writes attrs on the
// page's own <input>, which keeps its name, value and state.
type ConsoleCheckbox struct {
	// Label: the words beside the box. AriaLabel: the name of a box with no
	// words beside it, as a table row's. One of the two, not both.
	Label     string
	AriaLabel string
	Control   func(attrs g.Node) g.Node
}

func (c ConsoleCheckbox) Render(w io.Writer) error {
	if (c.Label == "") == (c.AriaLabel == "") {
		panic("console.ConsoleCheckbox needs a Label or an AriaLabel, not both")
	}
	if c.Control == nil {
		panic("console.ConsoleCheckbox needs a Control")
	}
	attrs := g.Group{h.Class("cui-console-checkbox-input"), h.Type("checkbox")}
	root := h.Label
	if c.AriaLabel != "" {
		attrs = append(attrs, h.Aria("label", c.AriaLabel))
		root = h.Span
	}
	return root(h.Class("cui-console-checkbox"),
		c.Control(attrs),
		h.Span(h.Class("cui-console-checkbox-box"), h.Aria("hidden", "true"), lucide.Check, lucide.Minus),
		g.Text(c.Label),
	).Render(w)
}
