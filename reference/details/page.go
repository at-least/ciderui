package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/probe"
)

// @invariant: page.html, the accordion's example on the kit's CSS, compiled as a site
// compiles it (build.Compile), and ring.html, for the focus ring in an answer: 24px
// in from the window's edge, the example as an answer of itself, and the example
// with a link beginning its first answer.
func writePages() error {
	css, err := build.Compile(build.Sheet{CSS: "@import \"index.css\";\n@import \"components/site/accordion/accordion.css\";", Sources: []string{"components"}})
	if err != nil {
		return err
	}
	example, err := os.ReadFile(filepath.Join(build.Root, "components/site/accordion/examples/accordion.html"))
	if err != nil {
		return err
	}
	page := func(body string) []byte {
		return []byte(`<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width"><style>` + css + `</style></head><body>` + body + "</body></html>\n")
	}
	dir := probe.Dir("details")
	html := strings.TrimSpace(string(example))
	if err := os.WriteFile(filepath.Join(dir, "page.html"), page(html), 0o644); err != nil {
		return err
	}
	const first = "<p>Add your address at checkout to see a delivery date.</p>"
	// @invariant: the example's items render closed; the page opens the first, as a page wanting one open writes it
	opened := strings.Replace(html, `<details class="cui-accordion-item">`, `<details class="cui-accordion-item" open>`, 1)
	nested := strings.Replace(opened, first, html, 1)
	linked := strings.Replace(opened, first, `<p><a href="#store">Visit a store</a> to collect your order.</p>`, 1)
	if opened == html || nested == opened || linked == opened {
		return errors.New("the examples changed: ring.html is not what it says")
	}
	return os.WriteFile(filepath.Join(dir, "ring.html"), page(`<main style="padding: 24px"><div id="nested">`+nested+`</div><div id="linked">`+linked+`</div></main>`), 0o644)
}
