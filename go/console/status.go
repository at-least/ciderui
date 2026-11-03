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
	// Series: a chart legend's series color, as CSS.
	Series string
	Class  string
}

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
