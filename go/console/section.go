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
	// Detail: the line under the title, what the section is for; "" for none.
	Detail string
	// Controls: at the header's end; nil for none.
	Controls g.Node
	Children g.Node
}

func (s ConsoleSection) Render(w io.Writer) error {
	return h.Section(h.Class("cui-console-section"),
		ConsoleSectionHeader{Level: ConsoleSectionHeaderLevelSection, Back: ConsoleSectionHeaderBack{Href: "", Label: ""}, Title: s.Title, Detail: s.Detail, Children: s.Controls},
		s.Children,
	).Render(w)
}
