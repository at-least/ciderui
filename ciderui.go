// Package ciderui embeds the kit's stylesheets and fonts, so a site's one Go
// dependency on this module, at a commit, pins both its markup (the packages
// under go/) and its CSS: [Assets] serves them. A site with utilities of its own
// builds a second stylesheet from what go/cmd/ciderui-css writes.
package ciderui

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// @trap: by name, not a pattern: a pattern would embed a stray file in the
// repository's root too.
//
//go:embed theme.css base.css type.css utilities.css components/*/*/*.css fonts
var kit embed.FS

func cssFiles() []string {
	var paths []string
	err := fs.WalkDir(kit, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".css") {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
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

// WriteCSS writes under dir what a site's own utilities' stylesheet imports:
// utilities.css and the two it names.
func WriteCSS(dir string) error {
	// @trap: wrapped alone: MkdirAll's error names the component that failed,
	// not the directory asked for, and a file's write error names itself
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("write %s: %w", dir, err)
	}
	for _, path := range []string{"utilities.css", "theme.css", "type.css"} {
		if err := os.WriteFile(filepath.Join(dir, path), []byte(css(path)), 0o644); err != nil {
			return err
		}
	}
	return nil
}
