package console

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/icon"
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
	Href     string
	Label    string
	Icon     icon.Icon
	Current  bool
	Disabled bool
	// Secondary is a count at the item's end.
	Secondary string
	Class     string
}

func (m ConsoleMenuItem) Render(w io.Writer) error {
	if m.Label == "" {
		panic("console.ConsoleMenuItem needs a Label")
	}
	if m.Disabled != (m.Href == "") {
		panic("console.ConsoleMenuItem needs an Href, unless it is disabled")
	}
	return h.A(class.Root("cui-console-menu-item", m.Class), g.If(!m.Disabled, h.Href(m.Href)), g.If(m.Current, h.Aria("current", "page")), g.If(m.Disabled, h.Aria("disabled", "true")),
		iconPart(m.Icon, "cui-console-menu-icon"),
		h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
		g.If(m.Secondary != "", h.Span(h.Class("cui-console-menu-secondary"), g.Text(m.Secondary))),
	).Render(w)
}

type ConsoleMenuGroup struct {
	Label string
	Icon  icon.Icon
	Open  bool
	Class string
	// Children: [ConsoleMenuItem]s.
	Children g.Node
}

func (m ConsoleMenuGroup) Render(w io.Writer) error {
	if m.Label == "" {
		panic("console.ConsoleMenuGroup needs a Label")
	}
	return h.Details(class.Root("cui-console-menu-group", m.Class), g.If(m.Open, h.Open()),
		h.Summary(h.Class("cui-console-menu-item"),
			iconPart(m.Icon, "cui-console-menu-icon"),
			h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
			lucide.ChevronRight.Class("cui-console-menu-chevron"),
		),
		m.Children,
	).Render(w)
}

// ConsoleChoiceMenu is a popover of choices, one of them checked: a
// [ConsolePlayer]'s subtitles or voices. The page's script hears a choice's
// click (its value), moves aria-checked to it and moves the focus among
// the choices by the arrow keys; the click closes the popover itself
// (popovertargetaction="hide").
type ConsoleChoiceMenu struct {
	// ID is the popover's id, the popovertarget of the button that opens it.
	ID      string
	Title   string
	Choices []ConsoleChoice
	Class   string
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
	if m.Title == "" {
		panic("console.ConsoleChoiceMenu needs a Title")
	}
	if len(m.Choices) == 0 {
		panic("console.ConsoleChoiceMenu needs a choice")
	}
	checked := 0
	for _, c := range m.Choices {
		if c.Label == "" || c.Value == "" {
			panic("console.ConsoleChoice needs a Label and a Value")
		}
		if c.Checked {
			checked++
		}
	}
	if checked != 1 {
		panic("console.ConsoleChoiceMenu needs one choice checked")
	}
	title := m.ID + "-title"
	return h.Div(class.Root("cui-console-popover", m.Class), h.ID(m.ID), h.Popover(),
		h.Div(h.Class("cui-console-menu"),
			h.P(h.Class("cui-console-menu-heading"), h.ID(title), g.Text(m.Title)),
			h.Div(h.Role("menu"), h.Aria("labelledby", title),
				g.Map(m.Choices, func(c ConsoleChoice) g.Node {
					return h.Button(h.Class("cui-console-menu-item"), h.Type("button"), h.Role("menuitemradio"), h.Aria("checked", strconv.FormatBool(c.Checked)), h.Value(c.Value),
						h.PopoverTarget(m.ID), h.PopoverTargetAction("hide"),
						lucide.Check.Class("cui-console-menu-icon"),
						h.Span(h.Class("cui-console-menu-label"), optionalAttr("lang", c.Lang), g.Text(c.Label)),
					)
				}),
			),
		),
	).Render(w)
}
