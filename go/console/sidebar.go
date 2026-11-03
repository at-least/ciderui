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
	Home     string
	Wordmark string
	// Close is the close button's aria-label.
	Close string
	Class string
	// Footer: an account.
	Footer g.Node
	// Children: [ConsoleMenuItem]s, headings and rules, as in a [ConsoleMenu]
	// (decision 27).
	Children g.Node
}

func (s ConsoleSidebar) Render(w io.Writer) error {
	for _, field := range []struct{ name, value string }{
		{"ID", s.ID},
		{"Label", s.Label},
		{"Wordmark", s.Wordmark},
		{"Close", s.Close},
	} {
		if field.value == "" {
			panic("console.ConsoleSidebar needs a " + field.name)
		}
	}
	return h.Nav(class.Root("cui-console-sidebar", s.Class), h.ID(s.ID), h.Popover(), h.Aria("label", s.Label),
		h.Div(h.Class("cui-console-sidebar-header"),
			h.A(h.Class("cui-console-sidebar-wordmark"), h.Href(s.Home), g.Text(s.Wordmark)),
			h.Button(h.Class("cui-console-btn cui-console-btn-icon cui-console-sidebar-close"), h.Data("variant", "plain"), h.Type("button"), h.PopoverTarget(s.ID), h.PopoverTargetAction("hide"), h.Aria("label", s.Close), lucide.X),
		),
		h.Div(h.Class("cui-console-sidebar-body cui-console-menu"), h.Data("variant", "sidebar"), s.Children),
		h.Div(h.Class("cui-console-sidebar-footer"), s.Footer),
	).Render(w)
}
