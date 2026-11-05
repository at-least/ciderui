package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
)

func TestSkipLinkRendersTheExample(t *testing.T) {
	example.Check(t, site.SkipLink{Href: "#main", Label: "Skip to Content", Class: ""}, "site/skip/examples/skip.html")
}

func TestSkipLinkAnEmptyHrefPanics(t *testing.T) {
	mustPanic(t, "site.SkipLink needs an Href", func() {
		example.Render(site.SkipLink{Href: "", Label: "x", Class: ""})
	})
}

func TestSkipLinkAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "site.SkipLink needs a Label", func() {
		example.Render(site.SkipLink{Href: "#main", Label: "", Class: ""})
	})
}

func TestSkipLinkTakesThePagesUtilities(t *testing.T) {
	got := example.Render(site.SkipLink{Href: "#main", Label: "跳到主要內容", Class: "z-50"})
	if !strings.HasPrefix(got, `<a class="cui-skip cui-btn z-50" href="#main">跳到主要內容</a>`) {
		t.Errorf("got %s", got)
	}
}
