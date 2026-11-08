package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleTabs is a page's sections, each its own page: App Store Connect's
// app tabs. A form's choice is not these links but the page's own radios in
// labels, in a <div class="cui-console-tabs" data-variant="filled"> (the
// example radios.html): the pill draws the radio's checked and focus.
type ConsoleTabs struct {
	AriaLabel string
	Tabs      []ConsoleTab
	Class     string
}

type ConsoleTab struct {
	Label string
	// Href: "" for a tab not available now.
	Href    string
	Current bool
}

func (t ConsoleTabs) Render(w io.Writer) error {
	if t.AriaLabel == "" {
		panic("console.ConsoleTabs needs an AriaLabel")
	}
	currents := 0
	for _, tab := range t.Tabs {
		if tab.Current {
			currents++
		}
	}
	if currents > 1 {
		panic("console.ConsoleTabs has one current tab at most")
	}
	tabs := g.Map(t.Tabs, func(tab ConsoleTab) g.Node {
		if tab.Label == "" {
			panic("console.ConsoleTab needs a Label")
		}
		if tab.Href == "" {
			return h.A(h.Class("cui-console-tab"), h.Aria("disabled", "true"), g.Text(tab.Label))
		}
		return h.A(h.Class("cui-console-tab"), h.Href(tab.Href), g.If(tab.Current, h.Aria("current", "page")), g.Text(tab.Label))
	})
	return h.Nav(class.Root("cui-console-tabs", t.Class), h.Aria("label", t.AriaLabel), tabs).Render(w)
}
