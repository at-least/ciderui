//go:build browser

// @trap: Every example on every surface, not listed cases: six times a site built on
// ciderui (crashcart) found a failing pair the cases had not listed.
package surfaces

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/browser"
	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/check"
	"github.com/at-least/ciderui/internal/pinned"
	"github.com/mxschmitt/playwright-go"
)

var rules = []string{"color-contrast", "link-in-text-block", "button-name", "link-name", "aria-valid-attr-value"}

type surface struct {
	name, open, close string
}

var surfaces = map[string][]surface{
	"site": {
		{"page", `<div class="cui-light">`, `</div>`},
		{"dark", `<div class="cui-dark">`, `</div>`},
	},
	"console": {
		{"page", `<div class="cui-console-page">`, `</div>`},
		{"dark", `<div class="cui-console-page cui-dark">`, `</div>`},
		{"card", `<div class="cui-console-page"><div class="cui-console-card">`, `</div></div>`},
		{"card dark", `<div class="cui-console-page cui-dark"><div class="cui-console-card">`, `</div></div>`},
	},
}

func TestSurfaces(t *testing.T) {
	axeDir, err := pinned.Axe()
	if err != nil {
		t.Fatal(err)
	}
	axe, err := os.ReadFile(filepath.Join(axeDir, "axe.min.js"))
	if err != nil {
		t.Fatal(err)
	}
	parts, err := build.Parts()
	if err != nil {
		t.Fatal(err)
	}
	b := browser.Launch(t)
	for _, part := range parts {
		css, err := build.Shipped(part.Set)
		if err != nil {
			t.Fatal(err)
		}
		files, err := build.ExampleFiles(part)
		if err != nil {
			t.Fatal(err)
		}
		for _, file := range files {
			example, err := os.ReadFile(filepath.Join(build.Root, part.Dir, "examples", file))
			if err != nil {
				t.Fatal(err)
			}
			for _, s := range surfaces[part.Set] {
				found := audit(t, b, css, s.open+string(example)+s.close, string(axe))
				check.Equal(t, fmt.Sprintf("%s/%s on the %s", part.Name, file, s.name), strings.Join(found, " | "), "")
			}
		}
	}
}

func audit(t *testing.T, b *browser.Browser, css, html, axe string) []string {
	t.Helper()
	page, err := b.Open(css, html, browser.Options{})
	if err != nil {
		t.Fatal(err)
	}
	defer page.Close()
	if err := page.Settle(); err != nil {
		t.Fatal(err)
	}
	if _, err := page.Page.AddScriptTag(playwright.PageAddScriptTagOptions{Content: playwright.String(axe)}); err != nil {
		t.Fatal(err)
	}
	v, err := page.Page.Evaluate(`async (rules) => {
  const r = await axe.run(document, { runOnly: { type: 'rule', values: rules }, resultTypes: ['violations'] });
  return r.violations.flatMap((v) => v.nodes.map((n) => v.id + ' ' + n.target.join(' ') + ': ' + n.failureSummary.replace(/\s+/g, ' ')));
}`, rules)
	if err != nil {
		t.Fatal(err)
	}
	var found []string
	if err := browser.Decode(v, &found); err != nil {
		t.Fatal(err)
	}
	return found
}
