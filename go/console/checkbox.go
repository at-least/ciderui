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
	Label   string
	Control func(attrs g.Node) g.Node
}

func (c ConsoleCheckbox) Render(w io.Writer) error {
	need("console.ConsoleCheckbox", "a Label", c.Label)
	if c.Control == nil {
		panic("console.ConsoleCheckbox needs a Control")
	}
	return h.Label(h.Class("cui-console-checkbox"),
		c.Control(g.Group{h.Class("cui-console-checkbox-input"), h.Type("checkbox")}),
		h.Span(h.Class("cui-console-checkbox-box"), h.Aria("hidden", "true"), lucide.Check, lucide.Minus),
		g.Text(c.Label),
	).Render(w)
}
