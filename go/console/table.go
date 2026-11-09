package console

import (
	"fmt"
	"io"

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
	// BodyAttrs: the <tbody>'s attributes (an id a swap targets); nil for none.
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

func (t ConsoleTable) Render(w io.Writer) error {
	if len(t.Columns) == 0 {
		panic("console.ConsoleTable needs a ConsoleTableColumn")
	}
	return h.Table(h.Class("cui-console-table"), g.Attr("data-stack"),
		h.THead(h.Tr(g.Map(t.Columns, func(c ConsoleTableColumn) g.Node {
			if c.Label == "" {
				panic("console.ConsoleTableColumn needs a Label")
			}
			return h.Th(h.Scope("col"), g.If(c.Figure, h.Class("cui-console-table-figure")), g.Text(c.Label))
		}))),
		h.TBody(t.BodyAttrs, g.Map(t.Rows, t.Row)),
	).Render(w)
}

// Row is one row as this table writes it, for a response that swaps one row.
func (t ConsoleTable) Row(r ConsoleTableRow) g.Node {
	if len(r.Cells) != len(t.Columns) {
		panic(fmt.Sprintf("console.ConsoleTableRow needs a cell for each column: %d columns, %d cells", len(t.Columns), len(r.Cells)))
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
