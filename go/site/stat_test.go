package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestStatRendersTheExamples(t *testing.T) {
	example.Check(t, site.Stat{Large: false, Before: "Up to", Value: "2 days", After: "on a single charge", Class: ""}, "site/stat/examples/stat.html")
	example.Check(t, site.Stat{Large: true, Before: "Charges up to", Value: "3x faster", After: "than before", Class: ""}, "site/stat/examples/large.html")
}

func TestStatPanics(t *testing.T) {
	example.Panics(t, func() site.Stat {
		return site.Stat{Large: false, Before: "", Value: "x", After: "", Class: ""}
	}, []example.Break[site.Stat]{
		{Want: "site.Stat needs a Value", Change: func(p *site.Stat) { p.Value = "" }},
	})
}

func TestStatOfAValueAlone(t *testing.T) {
	got := example.Render(site.Stat{Large: false, Before: "", Value: "x", After: "", Class: ""})
	if want := `<p class="cui-stat"><span class="cui-stat-value">x</span></p>`; got != want {
		t.Errorf("got %s", got)
	}
}
