//go:build browser

package zola

import (
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/browser"
)

// @invariant: zola/test/site/config.toml's base_url: the pages link their
// stylesheets there; served at another origin a page was drawn unstyled, a code
// line running it 645px wide at 320px, so each page's two sheets are counted
const siteOrigin = "https://check.test/"

type layout struct {
	Sheets int `json:"sheets"`
	Scroll int `json:"scroll"`
	First  int `json:"first"`
	Footer int `json:"footer"`
}

// TestNoPageScrollsSideways holds every page of the test site at a phone's, a
// tablet's and two desktops' widths: it scrolls no wider than the window (a code
// block or a table scrolls inside its own box), and its content's left edge is the
// footer's text's.
func TestNoPageScrollsSideways(t *testing.T) {
	out := buildSite(t, nil)
	var paths []string
	for path := range pages(t, out) {
		paths = append(paths, strings.TrimSuffix(filepath.ToSlash(path), "index.html"))
	}
	slices.Sort(paths)
	b := browser.Launch(t)
	for _, width := range []int{320, 390, 834, 1600} {
		p, err := b.OpenSite(siteOrigin, out, browser.Options{Width: width})
		if err != nil {
			t.Fatal(err)
		}
		for _, path := range paths {
			if _, err := p.Goto(siteOrigin + path); err != nil {
				t.Fatal(err)
			}
			v, err := p.Evaluate(`() => {
				const left = (el) => el && Math.round(el.getBoundingClientRect().left);
				const read = (sheet) => { try { return sheet.cssRules.length > 0; } catch { return false; } };
				return {
					sheets: [...document.styleSheets].filter(read).length,
					scroll: document.documentElement.scrollWidth,
					first: left(document.querySelector("main > *")),
					footer: left(document.querySelector(".cui-footer-copyright")),
				};
			}`)
			if err != nil {
				t.Fatal(err)
			}
			var got layout
			if err := browser.Decode(v, &got); err != nil {
				t.Fatal(err)
			}
			if got.Sheets != 2 {
				t.Fatalf("/%s at %dpx: %d stylesheets read, want dist/site.css and the theme's", path, width, got.Sheets)
			}
			if got.Scroll > width {
				t.Errorf("/%s at %dpx scrolls %dpx wide", path, width, got.Scroll)
			}
			if got.First != got.Footer {
				t.Errorf("/%s at %dpx: its content starts at %dpx, the footer's text at %dpx", path, width, got.First, got.Footer)
			}
		}
		if err := p.Context().Close(); err != nil {
			t.Fatal(err)
		}
	}
}
