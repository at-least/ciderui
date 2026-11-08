package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestListItemRendersTheExample(t *testing.T) {
	list := h.Div(h.Class("cui-console-list max-w-md"),
		console.ConsoleListItem{Href: "#shop", Label: "Shop App", Detail: "8 open issues · crash-free 99.2%", Class: "", Children: nil},
		console.ConsoleListItem{Href: "#api", Label: "Orders API", Detail: "2 open issues", Class: "", Children: nil},
		console.ConsoleListItem{Href: "", Label: "nightly-backup", Detail: "Last run 4 hours ago", Class: "", Children: nil},
	)
	example.Check(t, list, "console/console-list/examples/list.html")
}

func TestListItemRendersTheContentExample(t *testing.T) {
	list := h.Div(h.Class("cui-console-list max-w-md"),
		console.ConsoleListItem{Href: "", Label: "nightly-backup", Detail: "Last run 4 hours ago", Class: "",
			Children: h.P(h.Class("mt-2"), h.Button(h.Class("cui-console-btn"), h.Type("button"), g.Text("Pause")))},
	)
	example.Check(t, list, "console/console-list/examples/content.html")
}

func TestListItemARowOfItsOwnContentNeedsNoLabel(t *testing.T) {
	got := example.Render(console.ConsoleListItem{Href: "", Label: "", Detail: "", Class: "py-3", Children: h.P(g.Text("Shop App"))})
	if want := `<div class="cui-console-list-item py-3"><p>Shop App</p></div>`; got != want {
		t.Errorf("a row of its own content wrote more than that content:\n got: %s\nwant: %s", got, want)
	}
}

func TestListItemPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleListItem {
		return console.ConsoleListItem{Href: "", Label: "x", Detail: "", Class: "", Children: nil}
	}, []example.Break[console.ConsoleListItem]{
		{Want: "console.ConsoleListItem needs a Label or its own content", Change: func(p *console.ConsoleListItem) { p.Label = "" }},
	})
}
