package site

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type Stat struct {
	// Large: apple.com's AirPods Pro page.
	Large  bool
	Before string
	Value  string
	After  string
	Class  string
}

func (s Stat) Render(w io.Writer) error {
	if s.Value == "" {
		panic("site.Stat needs a Value")
	}
	return h.P(class.Root("cui-stat", s.Class), g.If(s.Large, h.Data("size", "lg")),
		g.If(s.Before != "", h.Span(h.Class("cui-stat-caption"), g.Text(s.Before))),
		h.Span(h.Class("cui-stat-value"), g.Text(s.Value)),
		g.If(s.After != "", h.Span(h.Class("cui-stat-caption"), g.Text(s.After))),
	).Render(w)
}
