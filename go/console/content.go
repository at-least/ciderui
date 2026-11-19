package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/need"
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

func (b ConsoleContentBack) link(owner string) g.Node {
	need.Field(owner, "an Href", b.Href)
	need.Field(owner, "a Label", b.Label)
	return h.A(h.Class("cui-console-link cui-console-link-back"), h.Href(b.Href), lucide.ChevronLeft, g.Text(b.Label))
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
	// Actions: the buttons, links or tabs at the header's end.
	Actions  g.Node
	Children g.Node
}

func (c ConsoleContent) Render(w io.Writer) error {
	need.Field("console.ConsoleContent", "a Title", c.Title)
	var sidebar, open, back g.Node
	if c.Sidebar != nil {
		sidebar, open = *c.Sidebar, sidebarOpen(c.Sidebar.OpenAriaLabel)
	}
	if c.Back != nil {
		back = c.Back.link("console.ConsoleContent.Back")
	}
	return h.Div(h.Class("cui-console-content"),
		sidebar,
		h.Main(h.ID("main"), h.TabIndex("-1"), h.Class("cui-console-content-main"),
			open,
			header(g.Group{
				back,
				h.H1(h.Class("cui-console-h2"), g.Text(c.Title)),
			}, c.Detail, c.Actions),
			c.Children,
		),
	).Render(w)
}
