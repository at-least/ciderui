package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleSidebar is cui-console-sidebar (components/console/console-sidebar): App
// Store Connect's newer look's app sidebar: its wordmark (a link home), its
// list (its Children, a [ConsoleMenu]'s at the toolkit's sidebar size:
// [ConsoleMenuItem]s, headings, rules) and its foot (an
// account); a popover, so below 1024px it is a sheet the page's
// cui-console-sidebar-open button opens (popovertarget its ID) and its own
// close button closes.
type ConsoleSidebar struct {
	// ID is its id, which the page's open button names. Required: an empty
	// one panics.
	ID string
	// Label is what it is, for a screen reader (aria-label). Required.
	Label string
	// Home is where its wordmark goes.
	Home string
	// Wordmark is its wordmark's text, the product's name. Required.
	Wordmark string
	// Close is its close button's label (aria-label). Required.
	Close string
	// Class is the page's utilities on it; "" for none.
	Class string
	// Footer is what goes at its foot (an account); nil for nothing.
	Footer g.Node
	// Children is its list.
	Children g.Node
}

// Render writes the sidebar.
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
