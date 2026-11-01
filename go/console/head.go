package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleHead is cui-console-head (components/console/console-head): what a page
// belongs to (an app, a project) by name, a link to it, with its underlined
// tabs beside the name, App Store Connect's app head.
type ConsoleHead struct {
	// Name is what the page belongs to. Required: an empty one panics.
	Name string
	// Href is its page.
	Href string
	// Label is what its tabs are called (their aria-label).
	Label string
	// Tabs are its tabs, in order.
	Tabs []ConsoleTab
	// Class is the page's utilities on the row; "" for none.
	Class string
}

// Render writes the head.
func (d ConsoleHead) Render(w io.Writer) error {
	if d.Name == "" {
		panic("console.ConsoleHead needs a Name")
	}
	return h.Div(class.Root("cui-console-head", d.Class),
		h.A(h.Class("cui-console-head-name"), h.Href(d.Href), g.Text(d.Name)),
		ConsoleTabs{Label: d.Label, Variant: ConsoleTabsVariantUnderlined, Tabs: d.Tabs, Class: ""},
	).Render(w)
}
