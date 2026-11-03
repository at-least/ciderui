package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestSectionRendersTheSectionExample(t *testing.T) {
	example.Check(t, site.Section{Ground: site.GroundPage, Rhythm: site.RhythmRegular, Dark: false, Class: "", Children: h.P(g.Text("Content"))}, "site/section/examples/section.html")
}

func TestSectionHeaderRendersTheHeaderExample(t *testing.T) {
	links := []site.Link{{Text: "Watch the video", Href: "#film"}, {Text: "Take the tour", Href: "#event"}}
	example.Check(t, site.SectionHeader{Headline: "See it in action.", Links: links, Class: ""}, "site/section/examples/header.html")
}

func TestSectionHeaderWithoutLinksHasNoList(t *testing.T) {
	got := example.Render(site.SectionHeader{Headline: "Values", Links: nil, Class: ""})
	want := `<header class="cui-section-header"><h2 class="cui-section-headline">Values</h2></header>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestSectionTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.Section{Ground: site.GroundPage, Rhythm: site.RhythmRegular, Dark: false, Class: "pt-0", Children: g.Text("x")})
	if !strings.HasPrefix(got, `<section class="cui-section pt-0">`) {
		t.Errorf("got %s", got)
	}
}

func TestSectionHeaderAnEmptyHeadlinePanics(t *testing.T) {
	mustPanic(t, "needs a Headline", func() {
		example.Render(site.SectionHeader{Headline: "", Links: nil, Class: ""})
	})
}

func TestSectionHeaderTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.SectionHeader{Headline: "x", Links: nil, Class: "mb-0"})
	if !strings.HasPrefix(got, `<header class="cui-section-header mb-0">`) {
		t.Errorf("got %s", got)
	}
}

func TestSectionDarkThenThePagesUtilities(t *testing.T) {
	got := example.Render(site.Section{Ground: site.GroundPage, Rhythm: site.RhythmRegular, Dark: true, Class: "pt-0", Children: g.Text("x")})
	if !strings.HasPrefix(got, `<section class="cui-section cui-dark pt-0">`) {
		t.Errorf("got %s", got)
	}
}
