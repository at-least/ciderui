package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleTabs is a page's sections, each its own page: App Store Connect's
// app tabs. A form's choice in their look is a [ConsoleSegmentedControl].
type ConsoleTabs struct {
	AriaLabel string
	Tabs      []ConsoleTab
	// Current: the current tab's Href; "" when no tab is current.
	Current string
}

type ConsoleTab struct {
	Label string
	Href  string
}

func (t ConsoleTabs) Render(w io.Writer) error {
	need.Field("console.ConsoleTabs", "an AriaLabel", t.AriaLabel)
	need.Some("console.ConsoleTabs", "Tabs", t.Tabs)
	if t.Current != "" {
		currents := 0
		for _, tab := range t.Tabs {
			if tab.Href == t.Current {
				currents++
			}
		}
		if currents != 1 {
			panic("console.ConsoleTabs needs a Current that is one of its Tabs' Href")
		}
	}
	var tabs g.Group
	for i, tab := range t.Tabs {
		need.Field(need.At("console.ConsoleTabs.Tabs", i), "a Label", tab.Label)
		need.Field(need.At("console.ConsoleTabs.Tabs", i), "an Href", tab.Href)
		tabs = append(tabs, pageLink("cui-console-tab", tab.Href, tab.Href == t.Current, g.Text(tab.Label)))
	}
	return h.Nav(h.Class("cui-console-tabs"), h.Aria("label", t.AriaLabel), tabs).Render(w)
}

// ConsoleSegmentedControl is a form's choice of one value, drawn as filled
// tabs: radios in labels, each pill drawing its radio's checked and focus.
type ConsoleSegmentedControl struct {
	AriaLabel string
	// Name: the radios' form field.
	Name     string
	Segments []ConsoleSegment
	// Value: the checked segment's Value.
	Value string
}

type ConsoleSegment struct {
	Value string
	Label string
}

func (c ConsoleSegmentedControl) Render(w io.Writer) error {
	need.Field("console.ConsoleSegmentedControl", "an AriaLabel", c.AriaLabel)
	need.Field("console.ConsoleSegmentedControl", "a Name", c.Name)
	if len(c.Segments) < 2 {
		panic("console.ConsoleSegmentedControl needs two Segments at least: one is no choice")
	}
	checked := 0
	for i, s := range c.Segments {
		need.Field(need.At("console.ConsoleSegmentedControl.Segments", i), "a Label", s.Label)
		if s.Value == c.Value {
			checked++
		}
	}
	if checked != 1 {
		panic("console.ConsoleSegmentedControl needs a Value that is one of its Segments' Value: a radio group starts with its value")
	}
	return h.Div(h.Class("cui-console-tabs"), h.Data("variant", "filled"), h.Role("group"), h.Aria("label", c.AriaLabel),
		g.Map(c.Segments, func(s ConsoleSegment) g.Node {
			return h.Label(h.Class("cui-console-tab"), h.Input(h.Type("radio"), h.Name(c.Name), h.Value(s.Value), g.If(s.Value == c.Value, h.Checked())), g.Text(s.Label))
		}),
	).Render(w)
}
