package console

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleMetric is a metric's cui-console-card (components/console/console-card):
// App Store Connect's newer look's analytics card, its title, its value, a
// change up or down and a caption; with an Href, all of it a link to what
// it counts (a.cui-console-card). A card holding anything else (a list, a
// chart) is the page's own <div class="cui-console-card">.
type ConsoleMetric struct {
	// Title is its title. Required: an empty one panics.
	Title string
	// Value is its value, text or markup (a figure in a color, a badge
	// beside it). Required: one that renders empty panics.
	Value g.Node
	// Change is its change beside the value; "" for none.
	Change string
	// ChangeColor is the change's color: the neutral (no change), success
	// (up, for the better) or error (for the worse); any other panics.
	ChangeColor ConsoleColor
	// Caption is a line under it; "" for none.
	Caption string
	// Href is where it goes, the whole card a link; "" for a card that is
	// no link.
	Href string
	// Class is the page's utilities on the card; "" for none.
	Class string
}

// Render writes the card.
func (m ConsoleMetric) Render(w io.Writer) error {
	if m.Title == "" {
		panic("console.ConsoleMetric needs a Title")
	}
	// rendered once to check it, then written as the markup it is
	var value strings.Builder
	if m.Value != nil {
		if err := m.Value.Render(&value); err != nil {
			return err
		}
	}
	if value.Len() == 0 {
		panic("console.ConsoleMetric needs a Value")
	}
	if m.ChangeColor != ConsoleColorNeutral && m.ChangeColor != ConsoleColorSuccess && m.ChangeColor != ConsoleColorError {
		panic("console.ConsoleMetric's change is neutral, success or error")
	}
	content := g.Group{
		h.P(h.Class("cui-console-card-title"), g.Text(m.Title)),
		h.P(h.Class("cui-console-card-figure"),
			h.Span(h.Class("cui-console-card-value"), g.Raw(value.String())),
			g.If(m.Change != "", h.Span(h.Class("cui-console-card-change"), optionalAttr("data-color", m.ChangeColor.value), g.Text(m.Change))),
		),
		g.If(m.Caption != "", h.P(h.Class("cui-console-card-caption"), g.Text(m.Caption))),
	}
	if m.Href != "" {
		return h.A(class.Root("cui-console-card", m.Class), h.Href(m.Href), content).Render(w)
	}
	return h.Div(class.Root("cui-console-card", m.Class), content).Render(w)
}
