package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

var localNavLinks = []site.Link{
	{Label: "Overview", Href: "/cider-book/"},
	{Label: "Tech Specs", Href: "/cider-book/specs/"},
	{Label: "Compare", Href: "/cider-book/compare/"},
	{Label: "Switch to Cider Book", Href: "/cider-book/switch/"},
}

func TestLocalNavRendersTheExample(t *testing.T) {
	example.Check(t, site.LocalNav{Title: "Cider Book 13", Href: "/cider-book/", Links: localNavLinks, Current: "/cider-book/", Button: &site.Link{Label: "Buy", Href: "/buy/cider-book/"}, AriaLabel: "Local", MenuAriaLabel: "Menu", Class: ""}, "site/localnav/examples/localnav.html")
}

func TestLocalNavNoLinksNoMenu(t *testing.T) {
	got := example.Render(site.LocalNav{Title: "AirTag", Href: "/airtag/", Links: nil, Current: "", Button: nil, AriaLabel: "Local", MenuAriaLabel: "", Class: "cui-dark"})
	want := `<nav class="cui-localnav cui-dark" aria-label="Local"><div class="cui-localnav-content">` +
		`<a class="cui-localnav-title" href="/airtag/">AirTag</a><div class="cui-localnav-actions"></div></div></nav>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestLocalNavMarksNoLinkWithoutACurrentPage(t *testing.T) {
	got := example.Render(site.LocalNav{Title: "Cider Book 13", Href: "/cider-book/", Links: localNavLinks, Current: "", Button: nil, AriaLabel: "Local", MenuAriaLabel: "Menu", Class: ""})
	if strings.Contains(got, "aria-current") {
		t.Errorf("got %s", got)
	}
}

func TestLocalNavACurrentMatchingNoLinkPanics(t *testing.T) {
	mustPanic(t, "site.LocalNav's Current is none of its Links'", func() {
		example.Render(site.LocalNav{Title: "x", Href: "/x/", Links: localNavLinks, Current: "/elsewhere/", Button: nil, AriaLabel: "Local", MenuAriaLabel: "Menu", Class: ""})
	})
}

func TestLocalNavAnEmptyTitlePanics(t *testing.T) {
	mustPanic(t, "needs a Title", func() {
		example.Render(site.LocalNav{Title: "", Href: "/x/", Links: nil, Current: "", Button: nil, AriaLabel: "Local", MenuAriaLabel: "", Class: ""})
	})
}

func TestLocalNavAnEmptyHrefPanics(t *testing.T) {
	mustPanic(t, "LocalNav needs an Href", func() {
		example.Render(site.LocalNav{Title: "x", Href: "", Links: nil, Current: "", Button: nil, AriaLabel: "Local", MenuAriaLabel: "", Class: ""})
	})
}

func TestLocalNavAnEmptyAriaLabelPanics(t *testing.T) {
	mustPanic(t, "site.LocalNav needs an AriaLabel", func() {
		example.Render(site.LocalNav{Title: "x", Href: "/x/", Links: nil, Current: "", Button: nil, AriaLabel: "", MenuAriaLabel: "", Class: ""})
	})
}

func TestLocalNavLinksWithNoMenuAriaLabelPanics(t *testing.T) {
	mustPanic(t, "site.LocalNav's Links need a MenuAriaLabel", func() {
		example.Render(site.LocalNav{Title: "x", Href: "/x/", Links: localNavLinks, Current: "", Button: nil, AriaLabel: "Local", MenuAriaLabel: "", Class: ""})
	})
}

func TestLocalNavAnEmptyLinkPanics(t *testing.T) {
	for _, link := range []site.Link{{Label: "", Href: "/x/"}, {Label: "x", Href: ""}} {
		mustPanic(t, "site.LocalNav's Links need their Label and Href", func() {
			example.Render(site.LocalNav{Title: "x", Href: "/x/", Links: []site.Link{link}, Current: "", Button: nil, AriaLabel: "Local", MenuAriaLabel: "Menu", Class: ""})
		})
	}
}

func TestLocalNavAnEmptyButtonPanics(t *testing.T) {
	mustPanic(t, "site.LocalNav's Button needs its Label and Href", func() {
		example.Render(site.LocalNav{Title: "x", Href: "/x/", Links: nil, Current: "", Button: &site.Link{Label: "", Href: "/buy/"}, AriaLabel: "Local", MenuAriaLabel: "", Class: ""})
	})
}

func TestLocalNavWritesItsWordsInThePagesLanguage(t *testing.T) {
	got := example.Render(site.LocalNav{Title: "x", Href: "/x/", Links: localNavLinks, Current: "", Button: nil, AriaLabel: "本頁", MenuAriaLabel: "選單", Class: ""})
	for _, want := range []string{`<nav class="cui-localnav" aria-label="本頁">`, `id="cui-localnav-menu" aria-label="選單">`} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}
