package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSidebar is a popover: below 1024px it is a sheet that the page's own
// <button class="cui-console-sidebar-open" popovertarget="ID"> opens.
type ConsoleSidebar struct {
	ID string
	// Label is the <nav>'s aria-label.
	Label string
	// Home is the wordmark's href.
	Home       string
	Wordmark   string
	CloseLabel string
	Class      string
	// Footer: an account.
	Footer g.Node
	// Children: [ConsoleMenuItem]s, headings and rules, as in a [ConsoleMenu]
	// (decision 27).
	Children g.Node
}

func (s ConsoleSidebar) Render(w io.Writer) error {
	if s.ID == "" {
		panic("console.ConsoleSidebar needs an ID")
	}
	if s.Label == "" {
		panic("console.ConsoleSidebar needs a Label")
	}
	if s.Wordmark == "" {
		panic("console.ConsoleSidebar needs a Wordmark")
	}
	if s.CloseLabel == "" {
		panic("console.ConsoleSidebar needs a CloseLabel")
	}
	return h.Nav(class.Root("cui-console-sidebar", s.Class), h.ID(s.ID), h.Popover(), h.Aria("label", s.Label),
		h.Div(h.Class("cui-console-sidebar-header"),
			h.A(h.Class("cui-console-sidebar-wordmark"), h.Href(s.Home), g.Text(s.Wordmark)),
			h.Button(h.Class("cui-console-btn cui-console-btn-icon cui-console-sidebar-close"), h.Data("variant", "plain"), h.Type("button"), h.PopoverTarget(s.ID), h.PopoverTargetAction("hide"), h.Aria("label", s.CloseLabel), lucide.X),
		),
		h.Div(h.Class("cui-console-sidebar-body cui-console-menu"), h.Data("variant", "sidebar"), s.Children),
		h.Div(h.Class("cui-console-sidebar-footer"), s.Footer),
	).Render(w)
}
