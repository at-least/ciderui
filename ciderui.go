// Package ciderui embeds the kit's stylesheets, so a site's one Go dependency
// on this module, at a commit, pins both its markup (the packages under go/)
// and its CSS. A site writes them out with WriteCSS or go/cmd/ciderui-css and
// its Tailwind entry imports the components it uses from there
// (`@import "../target/ciderui/index.css";`). The Inter font files are not
// embedded: fonts.css names them as url("fonts/…"), relative to the compiled
// stylesheet, so a site copies the kit's fonts/ directory beside its built CSS.
package ciderui

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

// root stylesheets by name, not a pattern: a pattern would embed a stray
// file in the repository's root too
//
//go:embed index.css theme.css tokens.css base.css type.css fonts.css site.css console.css docs.css components/*/*/*.css
var stylesheets embed.FS

// CSSFiles is every stylesheet's path in the kit, sorted.
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

// CSS is the text of the stylesheet at path, one of CSSFiles; any other path
// panics.
func CSS(path string) string {
	b, err := stylesheets.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// WriteCSS writes every stylesheet under dir, at its path in the kit.
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
