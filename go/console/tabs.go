package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleTabsVariant is what a row of [ConsoleTabs] is for, which sets its look and its
// markup. It has no default: its zero value panics in Render.
type ConsoleTabsVariant struct{ value string }

var (
	// ConsoleTabsVariantUnderlined is a page's sections, each its own page: App
	// Store Connect's app tabs. A <nav>; the current tab
	// aria-current="page".
	ConsoleTabsVariantUnderlined = ConsoleTabsVariant{"underlined"}
	// ConsoleTabsVariantFilled is a choice of a few views of one thing (a list's
	// status, a chart's period): the HIG's segmented control, Apple
	// Developer Forums' Topics and All Posts. A group; the chosen tab
	// aria-current="true".
	ConsoleTabsVariantFilled = ConsoleTabsVariant{"filled"}
)

// ConsoleTabs is cui-console-tabs (components/console/console-tabs): a row of
// tabs, each a link to a page.
type ConsoleTabs struct {
	// Label is what the row is called (its aria-label). Required: an empty
	// one panics.
	Label string
	// Variant is underlined or filled. Required: the zero ConsoleTabsVariant
	// panics.
	Variant ConsoleTabsVariant
	// Tabs are the tabs, in order.
	Tabs []ConsoleTab
	// Class is the page's utilities on the row; "" for none.
	Class string
}

// ConsoleTab is one tab of a [ConsoleTabs].
type ConsoleTab struct {
	// Label is its text. Required: an empty one panics.
	Label string
	// Href is the page it opens; "" for a tab not available now
	// (aria-disabled).
	Href string
	// Current is whether it is the page or the choice shown.
	Current bool
}

// Render writes the row.
func (t ConsoleTabs) Render(w io.Writer) error {
	if t.Label == "" {
		panic("console.ConsoleTabs needs a Label")
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
		return h.Nav(class.Root("cui-console-tabs", t.Class), h.Aria("label", t.Label), tabs).Render(w)
	}
	return h.Div(class.Root("cui-console-tabs", t.Class), h.Data("variant", "filled"), h.Role("group"), h.Aria("label", t.Label), tabs).Render(w)
}
