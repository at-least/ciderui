package console

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleListItem is a row of a grouped list; the group,
// <div class="cui-console-list">, is the page's own.
type ConsoleListItem struct {
	Href      string
	ID        string
	Label     string
	Secondary string
	Class     string
	// Children: the row's controls or lines, after its label.
	Children g.Node
}

func (l ConsoleListItem) Render(w io.Writer) error {
	// @trap: rendered once, then written raw: a page's node may cost, or render
	// differently, a second time
	var children strings.Builder
	if l.Children != nil {
		if err := l.Children.Render(&children); err != nil {
			return err
		}
	}
	if l.Label == "" && children.Len() == 0 {
		panic("console.ConsoleListItem needs a Label or its own content")
	}
	content := g.Group{
		g.If(l.Label != "", h.Span(h.Class("cui-console-list-label"), g.Text(l.Label))),
		g.If(l.Secondary != "", h.Span(h.Class("cui-console-list-secondary"), g.Text(l.Secondary))),
		g.Raw(children.String()),
	}
	if l.Href != "" {
		return h.A(class.Root("cui-console-list-item", l.Class), g.If(l.ID != "", h.ID(l.ID)), h.Href(l.Href), content).Render(w)
	}
	return h.Div(class.Root("cui-console-list-item", l.Class), g.If(l.ID != "", h.ID(l.ID)), content).Render(w)
}
