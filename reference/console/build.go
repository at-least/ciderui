package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/at-least/ciderui/internal/build"
)

// @invariant: writes the pages run serves into out: ciderui compiled as
// internal/harness compiles it (every component's sheet), its fonts beside it, and
// a page per console part from its example, each loaded in a frame of a set width
// by index.html (run reads the frames).
func writeSite(out string) error {
	if err := os.MkdirAll(out, 0o755); err != nil {
		return err
	}
	parts, err := build.Parts()
	if err != nil {
		return err
	}
	var sheets []string
	for _, p := range parts {
		sheets = append(sheets, fmt.Sprintf("@import \"%s/%s.css\";", p.Dir, p.Name))
	}
	css, err := build.Compile(build.Sheet{CSS: "@import \"fonts.css\";\n@import \"index.css\";\n" + strings.Join(sheets, "\n"), Sources: []string{"components"}})
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(out, "style.css"), []byte(css), 0o644); err != nil {
		return err
	}
	if err := copyDir(filepath.Join(build.Root, "fonts"), filepath.Join(out, "fonts")); err != nil {
		return err
	}
	ex := func(component, name string) (string, error) {
		for _, p := range parts {
			if p.Name == component {
				b, err := os.ReadFile(filepath.Join(build.Root, p.Dir, "examples", name+".html"))
				return string(b), err
			}
		}
		return "", fmt.Errorf("no part %s", component)
	}
	page := func(body string) []byte {
		return []byte(`<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link rel="stylesheet" href="style.css"></head><body>` + body + "</body></html>\n")
	}
	for _, p := range []struct{ file, component, example string }{{"tabs.html", "console-tabs", "many"}, {"content.html", "console-content", "content"}} {
		body, err := ex(p.component, p.example)
		if err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(out, p.file), page(body), 0o644); err != nil {
			return err
		}
	}
	frames := []struct {
		name          string
		width, height int
	}{
		{"tabs", 320, 60},
		{"content", 900, 200},
		{"content", 1200, 200},
		{"content", 1920, 200},
	}
	var iframes strings.Builder
	for _, f := range frames {
		fmt.Fprintf(&iframes, `<iframe id="%s-%d" src="%s.html" width="%d" height="%d"></iframe>`, f.name, f.width, f.name, f.width, f.height)
	}
	return os.WriteFile(filepath.Join(out, "index.html"), []byte(`<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><style>body { margin: 0; background: white; } iframe { display: block; border: 0; }</style></head><body>`+iframes.String()+"</body></html>\n"), 0o644)
}

// @invariant: src's tree into dst, a file there already overwritten, as Node's cpSync
// copies it
func copyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, b, 0o644)
	})
}
