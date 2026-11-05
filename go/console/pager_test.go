package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestPagerRendersTheExamples(t *testing.T) {
	pager := h.Div(h.Class("cui-console-page"),
		console.ConsolePager{AriaLabel: "Issue pages", Summary: "51–100 of 230", First: "", FirstLabel: "First", Previous: "#page-1", PreviousAriaLabel: "Previous page", Next: "#page-3", NextAriaLabel: "Next page", Class: ""},
	)
	example.Check(t, pager, "console/console-pager/examples/pager.html")
	first := h.Div(h.Class("cui-console-page"),
		console.ConsolePager{AriaLabel: "Issue pages", Summary: "1–50 of 230", First: "", FirstLabel: "First", Previous: "", PreviousAriaLabel: "Previous page", Next: "#page-2", NextAriaLabel: "Next page", Class: ""},
	)
	example.Check(t, first, "console/console-pager/examples/first.html")
}

func TestPagerALaterPageOfACursorList(t *testing.T) {
	later := h.Div(h.Class("cui-console-page"),
		console.ConsolePager{AriaLabel: "Timeline pages", Summary: "", First: "#newest", FirstLabel: "First", Previous: "", PreviousAriaLabel: "Previous page", Next: "#older", NextAriaLabel: "Next page", Class: ""},
	)
	example.Check(t, later, "console/console-pager/examples/later.html")
}

func TestPagerNothingToShowPanics(t *testing.T) {
	mustPanic(t, "needs a Summary or a page to go to", func() {
		example.Render(console.ConsolePager{AriaLabel: "Pages", Summary: "", First: "", FirstLabel: "First", Previous: "", PreviousAriaLabel: "Previous page", Next: "", NextAriaLabel: "Next page", Class: ""})
	})
}

func TestPagerAnEmptyAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ConsolePager needs an AriaLabel", func() {
		example.Render(console.ConsolePager{AriaLabel: "", Summary: "1–50 of 230", First: "", FirstLabel: "First", Previous: "", PreviousAriaLabel: "Previous page", Next: "#page-2", NextAriaLabel: "Next page", Class: ""})
	})
}

func TestPagerWithoutItsWordsPanics(t *testing.T) {
	for _, c := range []struct {
		want   string
		change func(*console.ConsolePager)
	}{
		{"ConsolePager needs a FirstLabel", func(p *console.ConsolePager) { p.FirstLabel = "" }},
		{"ConsolePager needs a PreviousAriaLabel", func(p *console.ConsolePager) { p.PreviousAriaLabel = "" }},
		{"ConsolePager needs a NextAriaLabel", func(p *console.ConsolePager) { p.NextAriaLabel = "" }},
	} {
		p := console.ConsolePager{AriaLabel: "Pages", Summary: "1–50 of 230", First: "", FirstLabel: "First", Previous: "", PreviousAriaLabel: "Previous page", Next: "#2", NextAriaLabel: "Next page", Class: ""}
		c.change(&p)
		mustPanic(t, c.want, func() { example.Render(p) })
	}
}

func TestPagerWritesTheWordsItIsGiven(t *testing.T) {
	got := example.Render(console.ConsolePager{AriaLabel: "頁", Summary: "", First: "#1", FirstLabel: "第一頁", Previous: "#2", PreviousAriaLabel: "上一頁", Next: "#4", NextAriaLabel: "下一頁", Class: ""})
	for _, want := range []string{`href="#1">第一頁</a>`, `aria-label="上一頁"`, `aria-label="下一頁"`} {
		if !strings.Contains(got, want) {
			t.Errorf("no %s in %s", want, got)
		}
	}
	for _, english := range []string{"First", "Previous", "Next"} {
		if strings.Contains(got, english) {
			t.Errorf("%s in %s", english, got)
		}
	}
}
