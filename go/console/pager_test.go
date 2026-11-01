package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestPagerRendersTheExamples(t *testing.T) {
	pager := h.Div(h.Class("cui-console-page"),
		console.ConsolePager{Label: "Issue pages", Summary: "51–100 of 230", First: "", Previous: "#page-1", Next: "#page-3", Class: ""},
	)
	example.Check(t, pager, "console/console-pager/examples/pager.html")
	first := h.Div(h.Class("cui-console-page"),
		console.ConsolePager{Label: "Issue pages", Summary: "1–50 of 230", First: "", Previous: "", Next: "#page-2", Class: ""},
	)
	example.Check(t, first, "console/console-pager/examples/first.html")
}

func TestPagerALaterPageOfACursorList(t *testing.T) {
	later := h.Div(h.Class("cui-console-page"),
		console.ConsolePager{Label: "Timeline pages", Summary: "", First: "#newest", Previous: "", Next: "#older", Class: ""},
	)
	example.Check(t, later, "console/console-pager/examples/later.html")
}

func TestPagerNothingToShowPanics(t *testing.T) {
	mustPanic(t, "needs a Summary or a page to go to", func() {
		example.Render(console.ConsolePager{Label: "Pages", Summary: "", First: "", Previous: "", Next: "", Class: ""})
	})
}
