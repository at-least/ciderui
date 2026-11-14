package console

import (
	"fmt"
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleMenuItem is a page in a [ConsoleSidebar].
type ConsoleMenuItem struct {
	Href    string
	Label   string
	Icon    lucide.Icon
	Current bool
}

func (m ConsoleMenuItem) node() g.Node {
	need("console.ConsoleMenuItem", "an Href", m.Href)
	need("console.ConsoleMenuItem", "a Label", m.Label)
	return pageLink("cui-console-menu-item", m.Href, m.Current, g.Group{
		g.If(m.Icon != (lucide.Icon{}), m.Icon.Class("cui-console-menu-icon")),
		h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
	})
}

// @invariant: a page has one subtitles menu, its player's, so its id is the
// kit's, not the page's
const subtitlesMenuID = "subtitles"

// ConsoleChoiceMenu is a popover of options in columns side by side, each its
// title over its options, one of them checked: a [ConsolePlayer]'s voices
// and subtitles. The page's script hears an option's click (its value, and
// its column's data-column), moves aria-checked to it and moves the focus
// among the options by the arrow keys; the click closes the popover itself
// (popovertargetaction="hide").
type ConsoleChoiceMenu struct {
	Columns []ConsoleChoiceColumn
}

type ConsoleChoiceColumn struct {
	// Column is the column's data-column, what the script reads.
	Column  string
	Title   string
	Options []ConsoleChoiceOption
}

type ConsoleChoiceOption struct {
	Label string
	// Value is the <button>'s value, what the script reads.
	Value string
	// Lang is the label's language, a BCP 47 tag; "" for the page's.
	Lang    string
	Checked bool
}

func (m ConsoleChoiceMenu) Render(w io.Writer) error {
	if len(m.Columns) == 0 {
		panic("console.ConsoleChoiceMenu needs Columns")
	}
	columns := map[string]bool{}
	for _, c := range m.Columns {
		if c.Column == "" || c.Title == "" {
			panic("console.ConsoleChoiceColumn needs a Column and a Title")
		}
		if columns[c.Column] {
			panic("console.ConsoleChoiceMenu has two columns " + strconv.Quote(c.Column) + ": each needs a Column of its own")
		}
		columns[c.Column] = true
		if len(c.Options) == 0 {
			panic("console.ConsoleChoiceColumn needs Options")
		}
		checked := 0
		for _, o := range c.Options {
			if o.Label == "" || o.Value == "" {
				panic("console.ConsoleChoiceOption needs a Label and a Value")
			}
			if o.Checked {
				checked++
			}
		}
		if checked != 1 {
			panic(fmt.Sprintf("console.ConsoleChoiceColumn needs one Checked option, has %d", checked))
		}
	}
	return h.Div(h.Class("cui-console-popover"), h.ID(subtitlesMenuID), h.Popover(),
		h.Div(h.Class("cui-console-menu-columns"),
			g.Map(m.Columns, func(c ConsoleChoiceColumn) g.Node {
				titleID := subtitlesMenuID + "-" + c.Column + "-title"
				return h.Div(h.Class("cui-console-menu"), h.Data("column", c.Column),
					h.P(h.Class("cui-console-menu-title"), h.ID(titleID), g.Text(c.Title)),
					h.Div(h.Role("menu"), h.Aria("labelledby", titleID),
						g.Map(c.Options, func(o ConsoleChoiceOption) g.Node {
							return h.Button(h.Class("cui-console-menu-item"), h.Type("button"), h.Role("menuitemradio"), h.Aria("checked", strconv.FormatBool(o.Checked)), h.Value(o.Value),
								h.PopoverTarget(subtitlesMenuID), h.PopoverTargetAction("hide"),
								lucide.Check.Class("cui-console-menu-icon"),
								h.Span(h.Class("cui-console-menu-label"), attr.Optional("lang", o.Lang), g.Text(o.Label)),
							)
						}),
					),
				)
			}),
		),
	).Render(w)
}
