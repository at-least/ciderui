package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleTabsVariant is what the tabs are for; it sets their markup as well
// as their look.
type ConsoleTabsVariant struct{ value string }

var (
	// ConsoleTabsVariantUnderlined is a page's sections, each its own page:
	// App Store Connect's app tabs.
	ConsoleTabsVariantUnderlined = ConsoleTabsVariant{"underlined"}
	// ConsoleTabsVariantFilled is a choice of a few views of one thing (a
	// list's status, a chart's period): the HIG's segmented control.
	ConsoleTabsVariantFilled = ConsoleTabsVariant{"filled"}
)

type ConsoleTabs struct {
	AriaLabel string
	Variant   ConsoleTabsVariant
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
	var current string
	switch t.Variant {
	case ConsoleTabsVariantUnderlined:
		current = "page"
	case ConsoleTabsVariantFilled:
		current = "true"
	default:
		panic("console.ConsoleTabs needs a Variant")
	}
	tabs := g.Map(t.Tabs, func(tab ConsoleTab) g.Node {
		if tab.Label == "" {
			panic("console.ConsoleTab needs a Label")
		}
		if tab.Href == "" {
			return h.A(h.Class("cui-console-tab"), h.Aria("disabled", "true"), g.Text(tab.Label))
		}
		return h.A(h.Class("cui-console-tab"), h.Href(tab.Href), g.If(tab.Current, h.Aria("current", current)), g.Text(tab.Label))
	})
	if t.Variant == ConsoleTabsVariantUnderlined {
		return h.Nav(class.Root("cui-console-tabs", t.Class), h.Aria("label", t.AriaLabel), tabs).Render(w)
	}
	return h.Div(class.Root("cui-console-tabs", t.Class), h.Data("variant", "filled"), h.Role("group"), h.Aria("label", t.AriaLabel), tabs).Render(w)
}

// ConsoleTabButtons is filled [ConsoleTabs] of submit buttons, the pressed
// one chosen: a choice the page's own form posts as Name=Value.
type ConsoleTabButtons struct {
	AriaLabel string
	// Name is the form field every button posts.
	Name    string
	Buttons []ConsoleTabButton
	Class   string
}

type ConsoleTabButton struct {
	Label   string
	Value   string
	Pressed bool
}

func (t ConsoleTabButtons) Render(w io.Writer) error {
	if t.AriaLabel == "" {
		panic("console.ConsoleTabButtons needs an AriaLabel")
	}
	if t.Name == "" {
		panic("console.ConsoleTabButtons needs a Name")
	}
	pressed := 0
	for _, b := range t.Buttons {
		if b.Pressed {
			pressed++
		}
	}
	if pressed > 1 {
		panic("console.ConsoleTabButtons has one pressed button at most")
	}
	buttons := g.Map(t.Buttons, func(b ConsoleTabButton) g.Node {
		if b.Label == "" {
			panic("console.ConsoleTabButton needs a Label")
		}
		if b.Value == "" {
			panic("console.ConsoleTabButton needs a Value")
		}
		// @external: aria-pressed="false" too: it is what makes a button a toggle to a
		// screen reader
		ariaPressed := "false"
		if b.Pressed {
			ariaPressed = "true"
		}
		return h.Button(h.Class("cui-console-tab"), h.Type("submit"), h.Name(t.Name), h.Value(b.Value), h.Aria("pressed", ariaPressed), g.Text(b.Label))
	})
	return h.Div(class.Root("cui-console-tabs", t.Class), h.Data("variant", "filled"), h.Role("group"), h.Aria("label", t.AriaLabel), buttons).Render(w)
}
