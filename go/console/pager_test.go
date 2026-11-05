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

func validPager() console.ConsolePager {
	return console.ConsolePager{AriaLabel: "Pages", Summary: "", First: "", FirstLabel: "First", Previous: "", PreviousAriaLabel: "Previous page", Next: "#next", NextAriaLabel: "Next page", Class: ""}
}

func TestPagerPanics(t *testing.T) {
	example.Panics(t, validPager, []example.Break[console.ConsolePager]{
		{Want: "console.ConsolePager needs an AriaLabel", Change: func(p *console.ConsolePager) { p.AriaLabel = "" }},
		{Want: "console.ConsolePager needs a FirstLabel", Change: func(p *console.ConsolePager) { p.FirstLabel = "" }},
		{Want: "console.ConsolePager needs a PreviousAriaLabel", Change: func(p *console.ConsolePager) { p.PreviousAriaLabel = "" }},
		{Want: "console.ConsolePager needs a NextAriaLabel", Change: func(p *console.ConsolePager) { p.NextAriaLabel = "" }},
		{Want: "console.ConsolePager needs a Summary or a page to go to", Change: func(p *console.ConsolePager) { p.Next = "" }},
	})
}

// any one of them is enough to draw; the rest write nothing
func TestPagerDrawsWhatItHas(t *testing.T) {
	for _, c := range []struct {
		change func(*console.ConsolePager)
		want   string
	}{
		{func(p *console.ConsolePager) { p.Summary = "1" }, `<p class="cui-console-pager-summary">1</p></nav>`},
		{func(p *console.ConsolePager) { p.First = "#1" }, `aria-label="Pages"><a class="cui-console-btn" href="#1">First</a></nav>`},
		{func(p *console.ConsolePager) { p.Previous = "#2" }, `aria-label="Pages"><a class="cui-console-btn cui-console-btn-icon" href="#2"`},
	} {
		p := validPager()
		p.Next = ""
		c.change(&p)
		got := example.Render(p)
		if !strings.Contains(got, c.want) || strings.Contains(got, "Next page") {
			t.Errorf("no %s, or a next page, in %s", c.want, got)
		}
	}
}
