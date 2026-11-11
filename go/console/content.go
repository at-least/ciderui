package console

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleContent is a console page's column: its [ConsoleSidebar], if any,
// beside the page's <main id="main" tabindex="-1">, which spaces what it
// holds (rule 6): a [ConsoleSectionHeader] of Level Page, then
// [ConsoleSection]s.
type ConsoleContent struct {
	// Sidebar: nil for a page with none.
	Sidebar *ConsoleSidebar
	// Children: a [ConsoleSidebarOpen] first where Sidebar is set.
	Children g.Node
}

func (c ConsoleContent) Render(w io.Writer) error {
	var sidebar g.Node
	if c.Sidebar != nil {
		sidebar = *c.Sidebar
	}
	return h.Div(h.Class("cui-console-content"),
		sidebar,
		h.Main(h.ID("main"), h.TabIndex("-1"), h.Class("cui-console-content-main"), c.Children),
	).Render(w)
}
