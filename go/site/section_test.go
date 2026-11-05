package site_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSectionRendersTheSectionExample(t *testing.T) {
	example.Check(t, site.Section{Alt: false, Large: false, Class: "", Children: h.P(g.Text("Content"))}, "site/section/examples/section.html")
}

func TestSectionHeaderRendersTheHeaderExample(t *testing.T) {
	links := []site.Link{{Label: "Watch the video", Href: "#film"}, {Label: "Take the tour", Href: "#event"}}
	example.Check(t, site.SectionHeader{Headline: "See it in action.", Links: links, Class: ""}, "site/section/examples/header.html")
}

func TestSectionHeaderWithoutLinksHasNoList(t *testing.T) {
	got := example.Render(site.SectionHeader{Headline: "Values", Links: nil, Class: ""})
	want := `<header class="cui-section-header"><h2 class="cui-section-headline">Values</h2></header>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestSectionHeaderPanics(t *testing.T) {
	example.Panics(t, func() site.SectionHeader {
		return site.SectionHeader{Headline: "x", Links: nil, Class: ""}
	}, []example.Break[site.SectionHeader]{
		{Want: "site.SectionHeader needs a Headline", Change: func(p *site.SectionHeader) { p.Headline = "" }},
	})
}
