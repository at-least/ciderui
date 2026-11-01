// Package ciderui holds the kit's stylesheets, embedded, so a site's Go
// dependency on this module, at a commit, is the one pin for both its markup
// (the packages under go/: site, console, docs, icon, lucide) and its CSS:
// the site writes the stylesheets out with WriteCSS (or `go run
// github.com/at-least/ciderui/go/cmd/ciderui-css <dir>`) and its Tailwind
// entry imports the components it uses from there
// (`@import "../target/ciderui/index.css";`). The Inter font files are not
// embedded: fonts.css refers to them as url("fonts/…"), relative to the
// compiled stylesheet, so a site copies the kit's fonts/ directory beside its
// built CSS once.
//
// CheckMarkup holds a rendered page's ciderui words to these stylesheets, for
// a site's tests.
package ciderui

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

// the root stylesheets by name: a pattern would take a stray file in the
// repository's root too
//
//go:embed index.css theme.css tokens.css base.css type.css fonts.css site.css console.css docs.css components/*/*/*.css
var stylesheets embed.FS

// CSSFiles is every stylesheet of the kit, by its path in the kit, in the
// order of their paths.
func CSSFiles() []string {
	var paths []string
	err := fs.WalkDir(stylesheets, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return paths
}

// CSS is the text of the kit's stylesheet at path (one of CSSFiles); a path
// not there panics.
func CSS(path string) string {
	b, err := stylesheets.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// WriteCSS writes every stylesheet of the kit under dir, at its path in the
// kit, creating the directories.
func WriteCSS(dir string) error {
	for _, path := range CSSFiles() {
		out := filepath.Join(dir, filepath.FromSlash(path))
		if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(out, []byte(CSS(path)), 0o644); err != nil {
			return err
		}
	}
	return nil
}
