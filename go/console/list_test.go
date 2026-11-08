package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestListItemRendersTheExample(t *testing.T) {
	list := h.Div(h.Class("cui-console-list max-w-md"),
		console.ConsoleListItem{Href: "#shop", ID: "", Label: "Shop App", Detail: "8 open issues · crash-free 99.2%", Class: "", Children: nil},
		console.ConsoleListItem{Href: "#api", ID: "", Label: "Orders API", Detail: "2 open issues", Class: "", Children: nil},
		console.ConsoleListItem{Href: "", ID: "", Label: "nightly-backup", Detail: "Last run 4 hours ago", Class: "", Children: nil},
	)
	example.Check(t, list, "console/console-list/examples/list.html")
}

func TestListItemRendersTheContentExample(t *testing.T) {
	list := h.Div(h.Class("cui-console-list max-w-md"),
		console.ConsoleListItem{Href: "", ID: "nightly-backup", Label: "nightly-backup", Detail: "Last run 4 hours ago", Class: "",
			Children: h.P(h.Class("mt-2"), h.Button(h.Class("cui-console-btn"), h.Type("button"), g.Text("Pause")))},
	)
	example.Check(t, list, "console/console-list/examples/content.html")
}

func TestListItemARowOfItsOwnContentNeedsNoLabel(t *testing.T) {
	got := example.Render(console.ConsoleListItem{Href: "", ID: "", Label: "", Detail: "", Class: "py-3", Children: h.P(g.Text("Shop App"))})
	if want := `<div class="cui-console-list-item py-3"><p>Shop App</p></div>`; got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestListItemPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleListItem {
		return console.ConsoleListItem{Href: "", ID: "", Label: "x", Detail: "", Class: "", Children: nil}
	}, []example.Break[console.ConsoleListItem]{
		{Want: "console.ConsoleListItem needs a Label or its own content", Change: func(p *console.ConsoleListItem) { p.Label = "" }},
	})
}

func TestListItemALinkTakesItsID(t *testing.T) {
	got := example.Render(console.ConsoleListItem{Href: "#a", ID: "row", Label: "x", Detail: "", Class: "", Children: nil})
	if !strings.HasPrefix(got, `<a class="cui-console-list-item" id="row" href="#a">`) {
		t.Errorf("the link is not the row with its id and href: %s", got)
	}
}
