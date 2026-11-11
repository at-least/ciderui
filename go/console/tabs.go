package console

import (
	"fmt"
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleTabs is a page's sections, each its own page: App Store Connect's
// app tabs. A form's choice in their look is a [ConsoleTabsChoice].
type ConsoleTabs struct {
	AriaLabel string
	Tabs      []ConsoleTab
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
		panic(fmt.Sprintf("console.ConsoleTabs has %d current tabs: one at most", currents))
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
	return h.Nav(h.Class("cui-console-tabs"), h.Aria("label", t.AriaLabel), tabs).Render(w)
}

// ConsoleTabsChoice is a form's choice of one value, drawn as filled tabs (a
// segmented control): radios in labels, each pill drawing its radio's checked
// and focus.
type ConsoleTabsChoice struct {
	AriaLabel string
	// Name: the radios' form field.
	Name    string
	Options []ConsoleTabsOption
	// Attrs: the group's other attributes, for the page's script (data-*);
	// nil for none.
	Attrs g.Node
}

type ConsoleTabsOption struct {
	Value   string
	Label   string
	Checked bool
}

func (c ConsoleTabsChoice) Render(w io.Writer) error {
	if c.AriaLabel == "" {
		panic("console.ConsoleTabsChoice needs an AriaLabel")
	}
	if c.Name == "" {
		panic("console.ConsoleTabsChoice needs a Name")
	}
	if len(c.Options) < 2 {
		panic("console.ConsoleTabsChoice needs two Options at least: one is no choice")
	}
	checked := 0
	for _, o := range c.Options {
		if o.Label == "" {
			panic("console.ConsoleTabsOption needs a Label")
		}
		if o.Checked {
			checked++
		}
	}
	if checked != 1 {
		panic(fmt.Sprintf("console.ConsoleTabsChoice needs one Checked option, has %d: a radio group starts with its value", checked))
	}
	return h.Div(h.Class("cui-console-tabs"), h.Data("variant", "filled"), h.Role("group"), h.Aria("label", c.AriaLabel), c.Attrs,
		g.Map(c.Options, func(o ConsoleTabsOption) g.Node {
			return h.Label(h.Class("cui-console-tab"), h.Input(h.Type("radio"), h.Name(c.Name), h.Value(o.Value), g.If(o.Checked, h.Checked())), g.Text(o.Label))
		}),
	).Render(w)
}
