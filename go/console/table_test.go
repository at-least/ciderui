package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func tableText(content string) console.ConsoleTableCell {
	return console.ConsoleTableCell{Content: g.Text(content)}
}

func jobsTable() console.ConsoleTable {
	link := func(href, label string) console.ConsoleTableCell {
		return console.ConsoleTableCell{Content: h.A(h.Class("cui-console-table-link"), h.Href(href), g.Text(label))}
	}
	return console.ConsoleTable{
		Columns: []console.ConsoleTableColumn{
			{Label: "Job", Figure: false},
			{Label: "Schedule", Figure: false},
			{Label: "Runs", Figure: true},
			{Label: "Paused until", Figure: false},
		},
		Rows: []console.ConsoleTableRow{
			{Cells: []console.ConsoleTableCell{link("#backup", "nightly-backup"), tableText("every 5 minutes"), tableText("142"), {Content: nil}}},
			{Cells: []console.ConsoleTableCell{link("#mailer", "report-mailer"), tableText("17 3 * * *"), tableText("9"), tableText("Nov 2")}},
		},
		BodyAttrs: nil,
		Class:     "",
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
		t.Errorf("no %s in %s", want, got)
	}
}

func TestTablePanics(t *testing.T) {
	example.Panics(t, jobsTable, []example.Break[console.ConsoleTable]{
		{Want: "console.ConsoleTable needs a ConsoleTableColumn", Change: func(p *console.ConsoleTable) { p.Columns = nil }},
		{Want: "console.ConsoleTableColumn needs a Label", Change: func(p *console.ConsoleTable) { p.Columns[1].Label = "" }},
		{Want: "console.ConsoleTableRow needs a cell for each column: 4 columns, 3 cells", Change: func(p *console.ConsoleTable) { p.Rows[1].Cells = p.Rows[1].Cells[:3] }},
	})
}
