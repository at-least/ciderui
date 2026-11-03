package console

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleMetric is a metric's cui-console-card. A card holding anything else
// (a list, a chart) is the page's own <div class="cui-console-card">.
type ConsoleMetric struct {
	Title string
	Value g.Node
	// Change is "+0.4%" or "No change"; "" for none.
	Change string
	// ChangeColor: success for the better, error for the worse, whichever
	// direction the number moved.
	ChangeColor ConsoleColor
	Caption     string
	Href        string
	Class       string
}

func (m ConsoleMetric) Render(w io.Writer) error {
	if m.Title == "" {
		panic("console.ConsoleMetric needs a Title")
	}
	// rendered once, then written raw: a page's node may cost, or render
	// differently, a second time
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
