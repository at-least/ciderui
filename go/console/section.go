package console

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSection is one of a page's sections: its header, then its blocks,
// which it spaces (rule 6).
type ConsoleSection struct {
	Title string
	// Detail: the line under the title, what the section is for.
	Detail string
	// Controls: at the header's end.
	Controls g.Node
	Children g.Node
}

func (s ConsoleSection) Render(w io.Writer) error {
	need("console.ConsoleSection", "a Title", s.Title)
	return h.Section(h.Class("cui-console-section"),
		sectionHeader(s.Title, s.Detail, s.Controls),
		s.Children,
	).Render(w)
}
