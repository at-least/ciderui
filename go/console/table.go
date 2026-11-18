package console

import (
	"fmt"
	"io"

	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleTable is a table whose rows stack on a phone, each cell after the
// first named by its column. A table that scrolls wraps in its own scroll box,
// the keyboard's to reach (the example scroll.html); none is written here,
// that every table not overflowing would carry a tab stop of its own.
type ConsoleTable struct {
	Columns []ConsoleTableColumn
	Rows    []ConsoleTableRow
	// BodyAttrs: the <tbody>'s attributes (an id a swap targets).
	BodyAttrs g.Node
}

type ConsoleTableColumn struct {
	Label string
	// Figure: its numbers at the cell's end.
	Figure bool
}

// ConsoleTableRow is drawn selected while a cell's own checkbox, a
// cui-console-checkbox that is the cell's child, is checked.
type ConsoleTableRow struct {
	// Cells: one for each of the table's Columns; nil for a cell that says
	// nothing, the table writing "—" and leaving it out of its stacked row on
	// a phone.
	Cells []g.Node
}

// ConsoleTableCell is a cell's words: its Label, a link where Href is set
// (the whole row its target, as a row has one link), and a quieter Detail
// line under it ("" for none).
type ConsoleTableCell struct {
	Href   string
	Label  string
	Detail string
}

func (c ConsoleTableCell) Render(w io.Writer) error {
	need.Field("console.ConsoleTableCell", "a Label", c.Label)
	label := g.Text(c.Label)
	if c.Href != "" {
		label = h.A(h.Class("cui-console-table-link"), h.Href(c.Href), label)
	}
	return g.Group{label, optionalText(h.Span, "cui-console-table-detail", c.Detail)}.Render(w)
}

func (t ConsoleTable) Render(w io.Writer) error {
	need.Some("console.ConsoleTable", "Columns", t.Columns)
	var headers, rows g.Group
	for i, c := range t.Columns {
		need.Field(need.At("console.ConsoleTable.Columns", i), "a Label", c.Label)
		headers = append(headers, h.Th(h.Scope("col"), g.If(c.Figure, h.Class("cui-console-table-figure")), g.Text(c.Label)))
	}
	for i, r := range t.Rows {
		rows = append(rows, t.row(need.At("console.ConsoleTable.Rows", i), r))
	}
	return h.Table(h.Class("cui-console-table"), g.Attr("data-stack"),
		h.THead(h.Tr(headers)),
		h.TBody(t.BodyAttrs, rows),
	).Render(w)
}

// Row is one row as this table writes it, for a response that swaps one row.
func (t ConsoleTable) Row(r ConsoleTableRow) g.Node {
	return t.row("console.ConsoleTableRow", r)
}

func (t ConsoleTable) row(path string, r ConsoleTableRow) g.Node {
	if len(r.Cells) != len(t.Columns) {
		panic(fmt.Sprintf("%s needs a cell for each column: %d columns, %d cells", path, len(t.Columns), len(r.Cells)))
	}
	var cells g.Group
	for i, content := range r.Cells {
		column := t.Columns[i]
		if content == nil {
			content = g.Group{g.Attr("data-empty"), g.Text("—")}
		}
		cells = append(cells, h.Td(g.If(column.Figure, h.Class("cui-console-table-figure")), g.If(i > 0, h.Data("label", column.Label)), content))
	}
	return h.Tr(cells)
}
