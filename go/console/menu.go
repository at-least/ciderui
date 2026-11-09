package console

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleMenuItem is a page in a [ConsoleSidebar].
type ConsoleMenuItem struct {
	// Href: "" for an item not available now.
	Href    string
	Label   string
	Icon    lucide.Icon
	Current bool
}

func (m ConsoleMenuItem) Render(w io.Writer) error {
	if m.Label == "" {
		panic("console.ConsoleMenuItem needs a Label")
	}
	content := g.Group{
		iconPart(m.Icon, "cui-console-menu-icon"),
		h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
	}
	if m.Href == "" {
		return h.A(h.Class("cui-console-menu-item"), h.Aria("disabled", "true"), content).Render(w)
	}
	return h.A(h.Class("cui-console-menu-item"), h.Href(m.Href), g.If(m.Current, h.Aria("current", "page")), content).Render(w)
}

// ConsoleChoiceMenu is a popover of choices in columns side by side, each its
// title over its choices, one of them checked: a [ConsolePlayer]'s voices
// and subtitles. The page's script hears a choice's click (its value, and
// its column's data-column), moves aria-checked to it and moves the focus
// among the choices by the arrow keys; the click closes the popover itself
// (popovertargetaction="hide").
type ConsoleChoiceMenu struct {
	// ID is the popover's id, the popovertarget of the button that opens it.
	ID      string
	Columns []ConsoleChoiceColumn
}

type ConsoleChoiceColumn struct {
	// Name is the column's data-column, what the script reads.
	Name    string
	Title   string
	Choices []ConsoleChoice
}

type ConsoleChoice struct {
	Label string
	// Value is the <button>'s value, what the script reads.
	Value string
	// Lang is the label's language, a BCP 47 tag; "" for the page's.
	Lang    string
	Checked bool
}

func (m ConsoleChoiceMenu) Render(w io.Writer) error {
	if m.ID == "" {
		panic("console.ConsoleChoiceMenu needs an ID")
	}
	if len(m.Columns) == 0 {
		panic("console.ConsoleChoiceMenu needs a column")
	}
	names := map[string]bool{}
	for _, c := range m.Columns {
		if c.Name == "" || c.Title == "" {
			panic("console.ConsoleChoiceColumn needs a Name and a Title")
		}
		if names[c.Name] {
			panic("console.ConsoleChoiceMenu has two columns named " + strconv.Quote(c.Name) + ": each needs a name of its own")
		}
		names[c.Name] = true
		if len(c.Choices) == 0 {
			panic("console.ConsoleChoiceColumn needs a choice")
		}
		checked := 0
		for _, choice := range c.Choices {
			if choice.Label == "" || choice.Value == "" {
				panic("console.ConsoleChoice needs a Label and a Value")
			}
			if choice.Checked {
				checked++
			}
		}
		if checked != 1 {
			panic("console.ConsoleChoiceColumn needs one choice checked")
		}
	}
	return h.Div(h.Class("cui-console-popover"), h.ID(m.ID), h.Popover(),
		h.Div(h.Class("cui-console-menu-columns"),
			g.Map(m.Columns, func(c ConsoleChoiceColumn) g.Node {
				titleID := m.ID + "-" + c.Name + "-title"
				return h.Div(h.Class("cui-console-menu"), h.Data("column", c.Name),
					h.P(h.Class("cui-console-menu-title"), h.ID(titleID), g.Text(c.Title)),
					h.Div(h.Role("menu"), h.Aria("labelledby", titleID),
						g.Map(c.Choices, func(choice ConsoleChoice) g.Node {
							return h.Button(h.Class("cui-console-menu-item"), h.Type("button"), h.Role("menuitemradio"), h.Aria("checked", strconv.FormatBool(choice.Checked)), h.Value(choice.Value),
								h.PopoverTarget(m.ID), h.PopoverTargetAction("hide"),
								lucide.Check.Class("cui-console-menu-icon"),
								h.Span(h.Class("cui-console-menu-label"), optionalAttr("lang", choice.Lang), g.Text(choice.Label)),
							)
						}),
					),
				)
			}),
		),
	).Render(w)
}
