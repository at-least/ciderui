package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func jobsTable() console.ConsoleTable {
	link := func(href, label string) g.Node {
		return h.A(h.Class("cui-console-table-link"), h.Href(href), g.Text(label))
	}
	return console.ConsoleTable{
		Columns: []console.ConsoleTableColumn{
			{Label: "Job", Figure: false},
			{Label: "Schedule", Figure: false},
			{Label: "Runs", Figure: true},
			{Label: "Paused until", Figure: false},
		},
		Rows: []console.ConsoleTableRow{
			{Cells: []g.Node{link("#backup", "nightly-backup"), g.Text("every 5 minutes"), g.Text("142"), nil}},
			{Cells: []g.Node{link("#mailer", "report-mailer"), g.Text("17 3 * * *"), g.Text("9"), g.Text("Nov 2")}},
		},
		BodyAttrs: nil,
	}
}

func TestTableRendersTheExample(t *testing.T) {
	example.Check(t, jobsTable(), "console/console-table/examples/stack.html")
}

func TestTableBodyAttrsGoOnItsBody(t *testing.T) {
	jobs := jobsTable()
	jobs.Rows = nil
	jobs.BodyAttrs = h.ID("rows")
	got := example.Render(jobs)
	if want := `<tbody id="rows"></tbody>`; !strings.Contains(got, want) {
		t.Errorf("the table's BodyAttrs are not on its <tbody>: no %s in %s", want, got)
	}
}

func TestTablePanics(t *testing.T) {
	example.Panics(t, jobsTable, []example.Break[console.ConsoleTable]{
		{Want: "console.ConsoleTable needs Columns", Change: func(p *console.ConsoleTable) { p.Columns = nil }},
		{Want: "console.ConsoleTable.Columns[1] needs a Label", Change: func(p *console.ConsoleTable) { p.Columns[1].Label = "" }},
		{Want: "console.ConsoleTable.Rows[1] needs a cell for each column: 4 columns, 3 cells", Change: func(p *console.ConsoleTable) { p.Rows[1].Cells = p.Rows[1].Cells[:3] }},
		{Want: "console.ConsoleTable.Rows[1] needs a cell for each column: 4 columns, 5 cells", Change: func(p *console.ConsoleTable) { p.Rows[1].Cells = append(p.Rows[1].Cells, g.Text("x")) }},
	})
}

func TestTableRowIsTheRowTheTableWrites(t *testing.T) {
	jobs := jobsTable()
	row := example.Render(jobs.Row(jobs.Rows[1]))
	if table := example.Render(jobs); !strings.Contains(table, row) {
		t.Errorf("Row wrote %s, which the table does not hold: %s", row, table)
	}
	jobs.Rows[1].Cells = jobs.Rows[1].Cells[:3]
	defer func() {
		if got, want := recover(), "console.ConsoleTableRow needs a cell for each column: 4 columns, 3 cells"; got != want {
			t.Errorf("Row of 3 cells panicked with %v, want %q", got, want)
		}
	}()
	jobs.Row(jobs.Rows[1])
}

func TestTableCellRendersTheExamplesCells(t *testing.T) {
	table := readRaw("console/console-table/examples/table.html")
	for _, cell := range []console.ConsoleTableCell{
		{Href: "#v1", Label: "2025.10.16", Detail: "iOS"},
		{Href: "", Label: "2.12", Detail: "iOS"},
	} {
		if got := example.Render(cell); !strings.Contains(table, "<td>"+got+"</td>") {
			t.Errorf("the cell wrote %s, which table.html holds in no <td>", got)
		}
	}
	if got, want := example.Render(console.ConsoleTableCell{Href: "#v1", Label: "2025.10.16", Detail: ""}), `<a class="cui-console-table-link" href="#v1">2025.10.16</a>`; got != want {
		t.Errorf("a cell with no Detail wrote %s, not %s", got, want)
	}
}

func TestTableCellPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleTableCell {
		return console.ConsoleTableCell{Href: "#v1", Label: "x", Detail: ""}
	}, []example.Break[console.ConsoleTableCell]{
		{Want: "console.ConsoleTableCell needs a Label", Change: func(p *console.ConsoleTableCell) { p.Label = "" }},
	})
}
