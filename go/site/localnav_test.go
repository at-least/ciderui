package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

var localNavLinks = []site.Link{
	{Text: "Overview", Href: "/cider-book/"},
	{Text: "Tech Specs", Href: "/cider-book/specs/"},
	{Text: "Compare", Href: "/cider-book/compare/"},
	{Text: "Switch to Cider Book", Href: "/cider-book/switch/"},
}

func TestLocalNavRendersTheExample(t *testing.T) {
	example.Check(t, site.LocalNav{Title: "Cider Book 13", Href: "/cider-book/", Links: localNavLinks, Current: "/cider-book/", Button: &site.Link{Text: "Buy", Href: "/buy/cider-book/"}, Dark: false, Class: ""}, "site/localnav/examples/localnav.html")
}

func TestLocalNavNoLinksNoMenu(t *testing.T) {
	got := example.Render(site.LocalNav{Title: "AirTag", Href: "/airtag/", Links: nil, Current: "", Button: nil, Dark: true, Class: ""})
	want := `<nav class="cui-localnav cui-dark" aria-label="Local"><div class="cui-localnav-content">` +
		`<a class="cui-localnav-title" href="/airtag/">AirTag</a><div class="cui-localnav-actions"></div></div></nav>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestLocalNavMarksNoLinkWithoutACurrentPage(t *testing.T) {
	got := example.Render(site.LocalNav{Title: "Cider Book 13", Href: "/cider-book/", Links: localNavLinks, Current: "", Button: nil, Dark: false, Class: ""})
	if strings.Contains(got, "aria-current") {
		t.Errorf("got %s", got)
	}
}

func TestLocalNavAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(site.LocalNav{Title: "", Href: "/x/", Links: nil, Current: "", Button: nil, Dark: false, Class: ""})
	})
}

func TestLocalNavDarkThenThePagesUtilities(t *testing.T) {
	got := example.Render(site.LocalNav{Title: "x", Href: "/x/", Links: nil, Current: "", Button: nil, Dark: true, Class: "mt-0"})
	if !strings.HasPrefix(got, `<nav class="cui-localnav cui-dark mt-0" aria-label="Local">`) {
		t.Errorf("got %s", got)
	}
}

func TestLocalNavAnEmptyHrefPanics(t *testing.T) {
	mustPanic(t, "LocalNav needs an Href", func() {
		example.Render(site.LocalNav{Title: "x", Href: "", Links: nil, Current: "", Button: nil, Dark: false, Class: ""})
	})
}
