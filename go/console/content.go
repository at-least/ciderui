package console

import (
	"io"

	"github.com/at-least/ciderui/go/lucide"
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
	Sidebar *ConsoleSidebar
	Back    *ConsoleContentBack
	Title   string
	// Detail: the line under the title, what the page is for.
	Detail string
	// Controls: at the header's end.
	Controls g.Node
	Children g.Node
}

func (c ConsoleContent) Render(w io.Writer) error {
	need("console.ConsoleContent", "a Title", c.Title)
	var sidebar, open, back g.Node
	if c.Sidebar != nil {
		sidebar, open = *c.Sidebar, sidebarOpen(c.Sidebar.OpenAriaLabel)
	}
	if c.Back != nil {
		need("console.ConsoleContentBack", "an Href", c.Back.Href)
		need("console.ConsoleContentBack", "a Label", c.Back.Label)
		back = h.A(h.Class("cui-console-link cui-console-header-back"), h.Href(c.Back.Href), lucide.ChevronLeft, g.Text(c.Back.Label))
	}
	return h.Div(h.Class("cui-console-content"),
		sidebar,
		h.Main(h.ID("main"), h.TabIndex("-1"), h.Class("cui-console-content-main"),
			open,
			header(g.Group{
				back,
				h.H1(h.Class("cui-console-h2"), g.Text(c.Title)),
			}, c.Detail, c.Controls),
			c.Children,
		),
	).Render(w)
}
