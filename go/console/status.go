package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStatus is cui-console-status (components/console/console-status): App
// Store Connect's status, a mark in its color before its word. The line
// under it (<span class="cui-console-status-detail">) is the page's own.
type ConsoleStatus struct {
	// Label is its word. Required: an empty one panics.
	Label string
	// Color is its mark's color.
	Color ConsoleColor
	// Series is a series' color for its mark, a CSS color (a chart's key,
	// --cui-console-status-color), in place of Color's; "" for none.
	Series string
	// Class is the page's utilities on it; "" for none.
	Class string
}

// Render writes the status.
func (s ConsoleStatus) Render(w io.Writer) error {
	if s.Label == "" {
		panic("console.ConsoleStatus needs a Label")
	}
	if s.Series != "" && s.Color != ConsoleColorNeutral {
		panic("console.ConsoleStatus takes a color or a series, not both")
	}
	return h.Span(class.Root("cui-console-status", s.Class), optionalAttr("data-color", s.Color.value),
		h.Span(h.Class("cui-console-status-mark"), g.If(s.Series != "", h.Style("--cui-console-status-color: "+s.Series))),
		g.Text(s.Label),
	).Render(w)
}
