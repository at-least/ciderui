package console

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleContentBack is the way back to the page a page is under, over its
// title, named by that page (iOS's back button).
type ConsoleContentBack struct {
	Href  string
	Label string
}

// ConsoleContent is a console page's column: its [ConsoleSidebar], if any,
// beside the page's <main id="main" tabindex="-1">, which spaces what it
// holds (rule 6): the sidebar's open button, the page's header, then
// Children, its [ConsoleSection]s.
type ConsoleContent struct {
	// Sidebar: nil for a page with none.
	Sidebar *ConsoleSidebar
	// Back: the zero value for none.
	Back  ConsoleContentBack
	Title string
	// Detail: the line under the title, what the page is for; "" for none.
	Detail string
	// Controls: at the header's end; nil for none.
	Controls g.Node
	Children g.Node
}

func (c ConsoleContent) Render(w io.Writer) error {
	if c.Title == "" {
		panic("console.ConsoleContent needs a Title")
	}
	if (c.Back.Href == "") != (c.Back.Label == "") {
		panic("console.ConsoleContent needs a Back with both an Href and a Label, or neither")
	}
	var sidebar, open g.Node
	if c.Sidebar != nil {
		sidebar, open = *c.Sidebar, sidebarOpen(c.Sidebar.OpenAriaLabel)
	}
	return h.Div(h.Class("cui-console-content"),
		sidebar,
		h.Main(h.ID("main"), h.TabIndex("-1"), h.Class("cui-console-content-main"),
			open,
			pageHeader(c.Back, c.Title, c.Detail, c.Controls),
			c.Children,
		),
	).Render(w)
}
