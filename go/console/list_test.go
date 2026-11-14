package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestListItemRendersTheExamples(t *testing.T) {
	list := h.Div(h.Class("max-w-md"), console.ConsoleList{Rows: []console.ConsoleListRow{
		{Label: "Shop App", Detail: "8 open issues · crash-free 99.2%", Children: nil},
		{Label: "Orders API", Detail: "2 open issues", Children: nil},
		{Label: "nightly-backup", Detail: "Last run 4 hours ago", Children: nil},
	}})
	example.Check(t, list, "console/console-list/examples/list.html")
	content := h.Div(h.Class("max-w-md"), console.ConsoleList{Rows: []console.ConsoleListRow{
		{Label: "nightly-backup", Detail: "Last run 4 hours ago",
			Children: h.P(h.Class("mt-2"), h.Button(h.Class("cui-console-btn"), h.Type("button"), g.Text("Pause")))},
	}})
	example.Check(t, content, "console/console-list/examples/content.html")
}

func TestListItemARowOfItsOwnContentNeedsNoLabel(t *testing.T) {
	got := example.Render(console.ConsoleList{Rows: []console.ConsoleListRow{{Label: "", Detail: "", Children: h.P(g.Text("Shop App"))}}})
	if want := `<div class="cui-console-list"><div class="cui-console-list-row"><p>Shop App</p></div></div>`; got != want {
		t.Errorf("a row of its own content wrote more than that content:\n got: %s\nwant: %s", got, want)
	}
}

func TestListPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleList {
		return console.ConsoleList{Rows: []console.ConsoleListRow{{Label: "x", Detail: "", Children: nil}}}
	}, []example.Break[console.ConsoleList]{
		{Want: "console.ConsoleList needs Rows: a page with none shows a ConsoleEmpty", Change: func(p *console.ConsoleList) { p.Rows = nil }},
		{Want: "console.ConsoleListRow needs a Label or Children", Change: func(p *console.ConsoleList) { p.Rows[0].Label = "" }},
	})
}
