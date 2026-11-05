package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleStatus is a colored mark before a word. A line under it,
// <span class="cui-console-status-detail">, is the page's own.
type ConsoleStatus struct {
	Label string
	Color ConsoleColor
	Class string
}

func (s ConsoleStatus) Render(w io.Writer) error {
	if s.Label == "" {
		panic("console.ConsoleStatus needs a Label")
	}
	return statusPart(s.Label, s.Color, "", s.Class).Render(w)
}

// seriesColor: a chart legend's series, as CSS; "" for color's.
func statusPart(label string, color ConsoleColor, seriesColor, rootClass string) g.Node {
	return h.Span(class.Root("cui-console-status", rootClass), optionalAttr("data-color", color.value),
		h.Span(h.Class("cui-console-status-mark"), g.If(seriesColor != "", h.Style("--cui-console-status-color: "+seriesColor))),
		g.Text(label),
	)
}
