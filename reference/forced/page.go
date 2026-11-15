package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/at-least/ciderui/internal/build"
	"github.com/at-least/ciderui/internal/probe"
)

// @invariant: page.html, the kit's parts that draw a state or a color in a color of
// their own, on the kit's CSS compiled as a site compiles it (build.Compile): an
// option unchecked and checked (the control: its state drawn by an edge).
func writePage() error {
	names := []string{"option"}
	var imports []string
	for _, n := range names {
		imports = append(imports, fmt.Sprintf("@import \"components/site/%s/%s.css\";", n, n))
	}
	css, err := build.Compile(build.Sheet{CSS: "@import \"index.css\";\n" + strings.Join(imports, "\n"), Sources: []string{"components"}})
	if err != nil {
		return err
	}
	ex := func(n, id string) (string, error) {
		b, err := os.ReadFile(filepath.Join(build.Root, "components/site", n, "examples", id+".html"))
		return strings.TrimSpace(string(b)), err
	}
	unchecked, err := ex("option", "option")
	if err != nil {
		return err
	}
	checked, err := ex("option", "checked")
	if err != nil {
		return err
	}
	body := "<div style=\"padding: 24px\">\n<div id=\"opt\" style=\"display: flex; gap: 16px; width: 700px\">" + unchecked + checked + "</div>\n</div>"
	return os.WriteFile(filepath.Join(probe.Dir("forced"), "page.html"), []byte(`<!doctype html><html lang="en"><head><meta charset="utf-8"><style>`+css+`</style></head><body>`+body+"</body></html>\n"), 0o644)
}
