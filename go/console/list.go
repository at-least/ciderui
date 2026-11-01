package console

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleListItem is cui-console-list-item (components/console/console-list): a
// row of App Store Connect's newer look's grouped list, its label and what
// follows it, then what the page gives it (its controls, its lines), a link
// where it has an Href. The group, <div class="cui-console-list">, is the
// page's own.
type ConsoleListItem struct {
	// Href is where it goes; "" for a row that is no link.
	Href string
	// ID is its id; "" for none.
	ID string
	// Label is its title; "" for a row whose content is all the page's own.
	Label string
	// Secondary is what follows it, in the secondary label; "" for none.
	Secondary string
	// Class is the page's utilities on the row; "" for none.
	Class string
	// Children is what the page gives it after its label; nil for nothing.
	// A row of no label and nothing panics.
	Children g.Node
}

// Render writes the row.
func (l ConsoleListItem) Render(w io.Writer) error {
	// rendered once to check it, then written as the markup it is
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
