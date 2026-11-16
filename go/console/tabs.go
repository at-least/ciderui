package console

import (
	"fmt"
	"io"

	"github.com/at-least/ciderui/go/internal/need"
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
	Label   string
	Href    string
	Current bool
}

func (t ConsoleTabs) Render(w io.Writer) error {
	need.Field("console.ConsoleTabs", "an AriaLabel", t.AriaLabel)
	need.Some("console.ConsoleTabs", "Tabs", t.Tabs)
	currents := 0
	for _, tab := range t.Tabs {
		if tab.Current {
			currents++
		}
	}
	if currents > 1 {
		panic(fmt.Sprintf("console.ConsoleTabs has %d current tabs: one at most", currents))
	}
	var tabs g.Group
	for i, tab := range t.Tabs {
		need.Field(need.At("console.ConsoleTabs.Tabs", i), "a Label", tab.Label)
		need.Field(need.At("console.ConsoleTabs.Tabs", i), "an Href", tab.Href)
		tabs = append(tabs, pageLink("cui-console-tab", tab.Href, tab.Current, g.Text(tab.Label)))
	}
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
}

type ConsoleTabsOption struct {
	Value   string
	Label   string
	Checked bool
}

func (c ConsoleTabsChoice) Render(w io.Writer) error {
	need.Field("console.ConsoleTabsChoice", "an AriaLabel", c.AriaLabel)
	need.Field("console.ConsoleTabsChoice", "a Name", c.Name)
	if len(c.Options) < 2 {
		panic("console.ConsoleTabsChoice needs two Options at least: one is no choice")
	}
	checked := 0
	for i, o := range c.Options {
		need.Field(need.At("console.ConsoleTabsChoice.Options", i), "a Label", o.Label)
		if o.Checked {
			checked++
		}
	}
	if checked != 1 {
		panic(fmt.Sprintf("console.ConsoleTabsChoice needs one Checked option, has %d: a radio group starts with its value", checked))
	}
	return h.Div(h.Class("cui-console-tabs"), h.Data("variant", "filled"), h.Role("group"), h.Aria("label", c.AriaLabel),
		g.Map(c.Options, func(o ConsoleTabsOption) g.Node {
			return h.Label(h.Class("cui-console-tab"), h.Input(h.Type("radio"), h.Name(c.Name), h.Value(o.Value), g.If(o.Checked, h.Checked())), g.Text(o.Label))
		}),
	).Render(w)
}
