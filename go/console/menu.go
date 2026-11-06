package console

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/class"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleMenu is a section's pages down its left side.
type ConsoleMenu struct {
	AriaLabel string
	Class     string
	// Children: [ConsoleMenuItem]s and [ConsoleMenuGroup]s, with a
	// <p class="cui-console-menu-heading"> or an
	// <hr class="cui-console-menu-rule"> between runs of them.
	Children g.Node
}

func (m ConsoleMenu) Render(w io.Writer) error {
	if m.AriaLabel == "" {
		panic("console.ConsoleMenu needs an AriaLabel")
	}
	return h.Nav(class.Root("cui-console-menu", m.Class), h.Aria("label", m.AriaLabel), m.Children).Render(w)
}

// ConsoleMenuItem is a page in a [ConsoleMenu], a [ConsoleMenuGroup] or a
// [ConsoleSidebar].
type ConsoleMenuItem struct {
	// Href: "" for an item not available now.
	Href    string
	Label   string
	Icon    lucide.Icon
	Current bool
	// Secondary is a count at the item's end.
	Secondary string
	Class     string
}

func (m ConsoleMenuItem) Render(w io.Writer) error {
	if m.Label == "" {
		panic("console.ConsoleMenuItem needs a Label")
	}
	content := g.Group{
		iconPart(m.Icon, "cui-console-menu-icon"),
		h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
		g.If(m.Secondary != "", h.Span(h.Class("cui-console-menu-secondary"), g.Text(m.Secondary))),
	}
	if m.Href == "" {
		return h.A(class.Root("cui-console-menu-item", m.Class), h.Aria("disabled", "true"), content).Render(w)
	}
	return h.A(class.Root("cui-console-menu-item", m.Class), h.Href(m.Href), g.If(m.Current, h.Aria("current", "page")), content).Render(w)
}

type ConsoleMenuGroup struct {
	Label string
	Icon  lucide.Icon
	Open  bool
	Items []ConsoleMenuItem
	Class string
}

func (m ConsoleMenuGroup) Render(w io.Writer) error {
	if m.Label == "" {
		panic("console.ConsoleMenuGroup needs a Label")
	}
	if len(m.Items) == 0 {
		panic("console.ConsoleMenuGroup needs an Item")
	}
	return h.Details(class.Root("cui-console-menu-group", m.Class), g.If(m.Open, h.Open()),
		h.Summary(h.Class("cui-console-menu-item"),
			iconPart(m.Icon, "cui-console-menu-icon"),
			h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
			lucide.ChevronRight.Class("cui-console-menu-chevron"),
		),
		g.Map(m.Items, func(item ConsoleMenuItem) g.Node { return item }),
	).Render(w)
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
	Class   string
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
			panic("console.ConsoleChoiceMenu's columns need names of their own")
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
	return h.Div(class.Root("cui-console-popover", m.Class), h.ID(m.ID), h.Popover(),
		h.Div(h.Class("cui-console-menu-columns"),
			g.Map(m.Columns, func(c ConsoleChoiceColumn) g.Node {
				title := m.ID + "-" + c.Name + "-title"
				return h.Div(h.Class("cui-console-menu"), h.Data("column", c.Name),
					h.P(h.Class("cui-console-menu-heading"), h.ID(title), g.Text(c.Title)),
					h.Div(h.Role("menu"), h.Aria("labelledby", title),
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
