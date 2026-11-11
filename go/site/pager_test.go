package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestPagerRendersTheExamples(t *testing.T) {
	pager := site.Pager{
		AriaLabel: "Pagination",
		Page:      "Page 3",
		First:     site.PagerLink{Href: "?page=1", Label: "First"},
		Previous:  site.PagerLink{Href: "?page=2", Label: "Previous page"},
		Next:      site.PagerLink{Href: "?page=4", Label: "Next page"},
		Last:      site.PagerLink{Href: "?page=10", Label: "Last"},
	}
	example.Check(t, pager, "site/pager/examples/pager.html")
	pager.Page = "Page 1"
	pager.First = site.PagerLink{Href: "", Label: ""}
	pager.Previous = site.PagerLink{Href: "", Label: ""}
	example.Check(t, pager, "site/pager/examples/first.html")
}

func TestPagerPanics(t *testing.T) {
	example.Panics(t, func() site.Pager {
		return site.Pager{AriaLabel: "Pages", Page: "Page 1", First: site.PagerLink{Href: "", Label: ""}, Previous: site.PagerLink{Href: "", Label: ""}, Next: site.PagerLink{Href: "?page=2", Label: "Next"}, Last: site.PagerLink{Href: "", Label: ""}}
	}, []example.Break[site.Pager]{
		{Want: "site.Pager needs an AriaLabel", Change: func(p *site.Pager) { p.AriaLabel = "" }},
		{Want: "site.Pager needs a Page", Change: func(p *site.Pager) { p.Page = "" }},
		{Want: "site.Pager's links need their Label", Change: func(p *site.Pager) { p.Next.Label = "" }},
	})
}
