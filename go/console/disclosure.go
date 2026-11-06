package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleDisclosure is the HIG's disclosure control.
type ConsoleDisclosure struct {
	Open         bool
	Summary      string
	SummaryClass string
	Class        string
	Children     g.Node
}

func (d ConsoleDisclosure) Render(w io.Writer) error {
	if d.Summary == "" {
		panic("console.ConsoleDisclosure needs a Summary")
	}
	return h.Details(class.Root("cui-console-disclosure", d.Class), g.If(d.Open, h.Open()),
		h.Summary(class.Root("cui-console-disclosure-summary", d.SummaryClass), lucide.ChevronRight.Class("cui-console-disclosure-chevron"), g.Text(d.Summary)),
		d.Children,
	).Render(w)
}
