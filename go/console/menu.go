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
	seen := map[string]bool{}
	var columns g.Group
	for i, c := range m.Columns {
		path := need.At("console.ConsoleChoiceMenu.Columns", i)
		need.Field(path, "a Column", c.Column)
		need.Field(path, "a Title", c.Title)
		if seen[c.Column] {
			panic("console.ConsoleChoiceMenu has two columns " + strconv.Quote(c.Column) + ": each needs a Column of its own")
		}
		seen[c.Column] = true
		need.Some(path, "Options", c.Options)
		var options g.Group
		for j, o := range c.Options {
			need.Field(need.At(path+".Options", j), "a Label", o.Label)
			need.Field(need.At(path+".Options", j), "a Value", o.Value)
			options = append(options, h.Button(h.Class("cui-console-menu-item"), h.Type("button"), h.Role("menuitemradio"), h.Aria("checked", strconv.FormatBool(j == 0)), h.Value(o.Value),
				h.PopoverTarget(ConsoleSubtitlesMenuID), h.PopoverTargetAction("hide"),
				lucide.Check.Class("cui-console-menu-icon"),
				h.Span(h.Class("cui-console-menu-label"), attr.Optional("lang", o.Lang), g.Text(o.Label)),
			))
		}
		titleID := ConsoleSubtitlesMenuID + "-" + c.Column + "-title"
		columns = append(columns, h.Div(h.Class("cui-console-menu"), h.Data("column", c.Column),
			h.P(h.Class("cui-console-menu-title"), h.ID(titleID), g.Text(c.Title)),
			h.Div(h.Role("menu"), h.Aria("labelledby", titleID), options),
		))
	}
	return h.Div(h.Class("cui-console-popover"), h.ID(ConsoleSubtitlesMenuID), h.Popover(),
		h.Div(h.Class("cui-console-menu-columns"), columns),
	).Render(w)
}
