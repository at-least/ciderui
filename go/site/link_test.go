package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestMoreLinkRendersTheExample(t *testing.T) {
	example.Check(t, site.MoreLink{Href: "#learn", Label: "Learn more", Class: ""}, "site/link/examples/more.html")
}

func TestMoreLinkTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.MoreLink{Href: "/x", Label: "More", Class: "mt-4"})
	if !strings.HasPrefix(got, `<a class="cui-link mt-4" data-variant="more"`) {
		t.Errorf("got %s", got)
	}
}

func TestMoreLinkEscapesItsText(t *testing.T) {
	got := example.Render(site.MoreLink{Href: "/a?b&c", Label: "<b>", Class: ""})
	for _, want := range []string{`href="/a?b&amp;c"`, `<span class="cui-link-text">&lt;b&gt;</span>`} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}

func TestMoreLinkAnEmptyTextPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(site.MoreLink{Href: "/x", Label: "", Class: ""})
	})
}

func TestMoreLinkAnEmptyHrefPanics(t *testing.T) {
	mustPanic(t, "MoreLink needs an Href", func() {
		example.Render(site.MoreLink{Href: "", Label: "x", Class: ""})
	})
}
