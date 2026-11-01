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

func TestStatAnEmptyValuePanics(t *testing.T) {
	mustPanic(t, "needs a Value", func() {
		example.Render(site.Stat{Large: false, Before: "", Value: "", After: "", Class: ""})
	})
}
