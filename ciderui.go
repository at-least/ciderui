// Package ciderui embeds the kit's stylesheets and fonts, so a site's one Go
// dependency on this module, at a commit, pins both its markup (the packages
// under go/) and its CSS. A site writes them out with go/cmd/ciderui-css and its
// Tailwind entry imports its set from there
// (`@import "../target/ciderui/site.css";`). The Tailwind CLI leaves fonts.css's
// url("fonts/…") as it is, relative to the built stylesheet: a site whose entry
// imports fonts.css serves the written fonts/ beside that stylesheet.
package ciderui

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// root stylesheets by name, not a pattern: a pattern would embed a stray
// file in the repository's root too
//
//go:embed index.css theme.css tokens.css base.css type.css fonts.css site.css console.css components/*/*/*.css fonts
var kit embed.FS

// every embedded file's path, sorted
func kitFiles() []string {
	var paths []string
	err := fs.WalkDir(kit, ".", func(path string, d fs.DirEntry, err error) error {
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

func cssFiles() []string {
	var paths []string
	for _, path := range kitFiles() {
		if strings.HasSuffix(path, ".css") {
			paths = append(paths, path)
		}
	}
	return paths
}

func css(path string) string {
	b, err := kit.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// WriteCSS writes every stylesheet and font under dir, at its path in the kit.
func WriteCSS(dir string) error {
	for _, path := range kitFiles() {
		out := filepath.Join(dir, filepath.FromSlash(path))
		if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
			return err
		}
		b, err := kit.ReadFile(path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(out, b, 0o644); err != nil {
			return err
		}
	}
	return nil
}
