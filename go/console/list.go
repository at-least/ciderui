package console

import (
	"io"
	"strings"

	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleList is a grouped list: its rows on one card, a hairline between.
type ConsoleList struct {
	Rows []ConsoleListRow
}

func (l ConsoleList) Render(w io.Writer) error {
	need.Some("console.ConsoleList", "Rows: a page with none shows a ConsoleEmpty", l.Rows)
	var rows g.Group
	for i, r := range l.Rows {
		rows = append(rows, r.node(need.At("console.ConsoleList.Rows", i)))
	}
	return h.Div(h.Class("cui-console-list"), rows).Render(w)
}

// ConsoleListRow is a row of a [ConsoleList].
type ConsoleListRow struct {
	Label  string
	Detail string
	// Children: the row's controls or lines, after its label.
	Children g.Node
}

func (l ConsoleListRow) node(path string) g.Node {
	return g.NodeFunc(func(w io.Writer) error {
		// @trap: rendered once, then written raw: a page's node may cost, or
		// render differently, a second time
		var children strings.Builder
		if l.Children != nil {
			if err := l.Children.Render(&children); err != nil {
				return err
			}
		}
		if l.Label == "" && children.Len() == 0 {
			panic(path + " needs a Label or Children")
		}
		return h.Div(h.Class("cui-console-list-row"),
			optionalText(h.Span, "cui-console-list-label", l.Label),
			optionalText(h.Span, "cui-console-list-detail", l.Detail),
			g.Raw(children.String()),
		).Render(w)
	})
}
