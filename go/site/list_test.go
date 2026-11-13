package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
)

func TestListRendersTheExamples(t *testing.T) {
	example.Check(t, site.List{Rows: []site.ListRow{
		{Title: "Software keyboard fails to appear in a playground app", Href: "#keyboard", Detail: "The on-screen keyboard does not appear when editing a text field in an app made and run in Swift Playgrounds. It happens in a new, empty app too, with no code of its own.", Aside: g.Text("3 replies · 8m")},
		{Title: "Enrollment still pending after ten days", Href: "#enrollment", Detail: "Payment went through and the ID was sent the same day.", Aside: g.Text("12 replies · 2h")},
		{Title: "Widgets stop refreshing after an update", Href: "#widgets", Detail: "", Aside: nil},
	}}, "site/list/examples/list.html")
}

func TestListPanics(t *testing.T) {
	example.Panics(t, func() site.List {
		return site.List{Rows: []site.ListRow{{Title: "x", Href: "#x", Detail: "", Aside: nil}}}
	}, []example.Break[site.List]{
		{Want: "site.List needs Rows: a page with none says so in its own words", Change: func(l *site.List) { l.Rows = nil }},
		{Want: "site.ListRow needs a Title and an Href", Change: func(l *site.List) { l.Rows[0].Title = "" }},
		{Want: "site.ListRow needs a Title and an Href", Change: func(l *site.List) { l.Rows[0].Href = "" }},
	})
}
