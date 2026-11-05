package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSectionHead is a section's title with its controls at the row's end.
type ConsoleSectionHead struct {
	// Title: the page's heading (<h1 class="cui-console-h2">) and any line
	// under it, in the one child that takes the row's room.
	Title g.Node
	Class string
	// Children: the controls; nil for none.
	Children g.Node
}

func (s ConsoleSectionHead) Render(w io.Writer) error {
	if s.Title == nil {
		panic("console.ConsoleSectionHead needs a Title")
	}
	return h.Header(class.Root("cui-console-section-head", s.Class),
		h.Div(s.Title),
		s.Children,
	).Render(w)
}
