package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleDisclosure is the HIG's disclosure control.
type ConsoleDisclosure struct {
	Summary  string
	Children g.Node
}

func (d ConsoleDisclosure) Render(w io.Writer) error {
	if d.Summary == "" {
		panic("console.ConsoleDisclosure needs a Summary")
	}
	return h.Details(h.Class("cui-console-disclosure"),
		h.Summary(h.Class("cui-console-disclosure-summary"), lucide.ChevronRight.Class("cui-console-disclosure-chevron"), g.Text(d.Summary)),
		d.Children,
	).Render(w)
}
