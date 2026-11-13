package console

import (
	"io"
	"strings"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleList is a grouped list: its rows on one card, a hairline between.
type ConsoleList struct {
	Rows []ConsoleListRow
}

func (l ConsoleList) Render(w io.Writer) error {
	if len(l.Rows) == 0 {
		panic("console.ConsoleList needs Rows: a page with none shows a ConsoleEmpty")
	}
	return h.Div(h.Class("cui-console-list"), g.Map(l.Rows, func(r ConsoleListRow) g.Node { return r })).Render(w)
}

// ConsoleListRow is a row of a [ConsoleList].
type ConsoleListRow struct {
	Label  string
	Detail string
	// Children: the row's controls or lines, after its label.
	Children g.Node
}

func (l ConsoleListRow) Render(w io.Writer) error {
	// @trap: rendered once, then written raw: a page's node may cost, or render
	// differently, a second time
	var children strings.Builder
	if l.Children != nil {
		if err := l.Children.Render(&children); err != nil {
			return err
		}
	}
	if l.Label == "" && children.Len() == 0 {
		panic("console.ConsoleListRow needs a Label or its own content")
	}
	return h.Div(h.Class("cui-console-list-row"),
		optionalText(h.Span, "cui-console-list-label", l.Label),
		optionalText(h.Span, "cui-console-list-detail", l.Detail),
		g.Raw(children.String()),
	).Render(w)
}
