package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleHead names what the page belongs to (an app, a project) beside its
// tabs.
type ConsoleHead struct {
	Name          string
	Href          string
	TabsAriaLabel string
	Tabs          []ConsoleTab
	Class         string
}

func (d ConsoleHead) Render(w io.Writer) error {
	if d.Name == "" {
		panic("console.ConsoleHead needs a Name")
	}
	return h.Div(class.Root("cui-console-head", d.Class),
		h.A(h.Class("cui-console-head-name"), h.Href(d.Href), g.Text(d.Name)),
		ConsoleTabs{AriaLabel: d.TabsAriaLabel, Variant: ConsoleTabsVariantUnderlined, Tabs: d.Tabs, Class: ""},
	).Render(w)
}
