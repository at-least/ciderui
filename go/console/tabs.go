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
	// list's status, a chart's period): the HIG's segmented control. Of
	// links. A form's choice, collected with its fields, is the page's own
	// radios in labels (the example radios.html): the pill draws the radio's
	// checked and focus as the chosen tab's.
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
