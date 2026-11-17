package console

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/attr"
	"github.com/at-least/ciderui/go/internal/need"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleMenuItem is a page in a [ConsoleSidebar].
type ConsoleMenuItem struct {
	Href  string
	Label string
	Icon  lucide.Icon
}

func (m ConsoleMenuItem) node(path string, current bool) g.Node {
	need.Field(path, "an Href", m.Href)
	need.Field(path, "a Label", m.Label)
	return pageLink("cui-console-menu-item", m.Href, current, g.Group{
		g.If(m.Icon != (lucide.Icon{}), m.Icon.Class("cui-console-menu-icon")),
		h.Span(h.Class("cui-console-menu-label"), g.Text(m.Label)),
	})
}

// ConsoleSubtitlesMenuID is the [ConsoleChoiceMenu]'s id, for a page's
// popovertarget and script.
//
// @invariant: a page has one subtitles menu, its player's, so its id is the
// kit's, not the page's
const ConsoleSubtitlesMenuID = "cui-console-subtitles"

// ConsoleChoiceMenu is a popover of options in columns side by side, each its
// title over its options, written at rest: each column's first option
// checked. The page's script checks the stored choice, hears an option's
// click (its value, and its column's data-column), moves aria-checked to it
// and moves the focus among the options by the arrow keys; the click closes
// the popover itself (popovertargetaction="hide"). A [ConsolePlayer]'s voices
// and subtitles.
type ConsoleChoiceMenu struct {
	// AriaLabel names the button that opens the menu, the player's subtitles
	// button.
	AriaLabel string
	Columns   []ConsoleChoiceColumn
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
	Lang string
}

func (m ConsoleChoiceMenu) Render(w io.Writer) error {
	need.Some("console.ConsoleChoiceMenu", "Columns", m.Columns)
	columns := map[string]bool{}
	for i, c := range m.Columns {
		column := need.At("console.ConsoleChoiceMenu.Columns", i)
		need.Field(column, "a Column", c.Column)
		need.Field(column, "a Title", c.Title)
		if columns[c.Column] {
			panic("console.ConsoleChoiceMenu has two columns " + strconv.Quote(c.Column) + ": each needs a Column of its own")
		}
		columns[c.Column] = true
		need.Some(column, "Options", c.Options)
		for j, o := range c.Options {
			need.Field(need.At(column+".Options", j), "a Label", o.Label)
			need.Field(need.At(column+".Options", j), "a Value", o.Value)
		}
	}
	return h.Div(h.Class("cui-console-popover"), h.ID(ConsoleSubtitlesMenuID), h.Popover(),
		h.Div(h.Class("cui-console-menu-columns"),
			g.Map(m.Columns, func(c ConsoleChoiceColumn) g.Node {
				titleID := ConsoleSubtitlesMenuID + "-" + c.Column + "-title"
				var options g.Group
				for j, o := range c.Options {
					options = append(options, h.Button(h.Class("cui-console-menu-item"), h.Type("button"), h.Role("menuitemradio"), h.Aria("checked", strconv.FormatBool(j == 0)), h.Value(o.Value),
						h.PopoverTarget(ConsoleSubtitlesMenuID), h.PopoverTargetAction("hide"),
						lucide.Check.Class("cui-console-menu-icon"),
						h.Span(h.Class("cui-console-menu-label"), attr.Optional("lang", o.Lang), g.Text(o.Label)),
					))
				}
				return h.Div(h.Class("cui-console-menu"), h.Data("column", c.Column),
					h.P(h.Class("cui-console-menu-title"), h.ID(titleID), g.Text(c.Title)),
					h.Div(h.Role("menu"), h.Aria("labelledby", titleID), options),
				)
			}),
		),
	).Render(w)
}
