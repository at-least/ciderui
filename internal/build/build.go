// Package build is the repository's layout and its Tailwind builds: the parts, the
// stylesheets the tests and dist/ compile, and a site's own utilities' sheet.
package build

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/at-least/ciderui/internal/pinned"
)

// Root is the repository's top directory: the nearest above the working directory
// that holds this module's go.mod.
//
// @invariant: found from the working directory, not from where this file was compiled: a
// binary runs on another machine (a probe on air) or in a copy of the repository
// (test/probes), and reads and writes the tree it runs in
var Root = func() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for {
		mod, err := os.ReadFile(filepath.Join(dir, "go.mod"))
		if err == nil && strings.HasPrefix(string(mod), "module github.com/at-least/ciderui\n") {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			panic("build: no github.com/at-least/ciderui go.mod above " + dir)
		}
		dir = parent
	}
}()

// Sets are the component sets, each a directory of components/ and a dist/ sheet.
var Sets = []string{"site", "console"}

// Part is one component: components/<Set>/<Name>, Dir that path from Root.
type Part struct {
	Set  string
	Name string
	Dir  string
}

// Parts are every component, sorted by name across the sets.
//
// @invariant: the order the tests' stylesheets and dist's import them in, which a part's
// rule over another's of the same specificity relies on
func Parts() ([]Part, error) {
	var parts []Part
	for _, set := range Sets {
		entries, err := os.ReadDir(filepath.Join(Root, "components", set))
		if err != nil {
			return nil, fmt.Errorf("build: parts: %w", err)
		}
		for _, e := range entries {
			parts = append(parts, Part{Set: set, Name: e.Name(), Dir: "components/" + set + "/" + e.Name()})
		}
	}
	slices.SortStableFunc(parts, func(a, b Part) int { return strings.Compare(a.Name, b.Name) })
	return parts, nil
}

// @external: entry, a file of its own: `@import` resolves from where it sits. Minified,
// as a site ships it.
func tailwind(entry string) (string, error) {
	bin, err := pinned.Tailwind()
	if err != nil {
		return "", err
	}
	dir, err := os.MkdirTemp("", "ciderui-out-")
	if err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	defer os.RemoveAll(dir)
	out := filepath.Join(dir, "out.css")
	if msg, err := exec.Command(bin, "-i", entry, "-o", out, "--minify").CombinedOutput(); err != nil {
		return "", fmt.Errorf("build: tailwindcss -i %s: %w\n%s", entry, err, msg)
	}
	css, err := os.ReadFile(out)
	if err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	return string(css), nil
}

// Sheet is what Compile builds: CSS, the entry after Tailwind's import, its @import
// paths relative to Root; Classes, class strings for Tailwind to find; Sources,
// paths relative to Root that Tailwind scans, as a site's @source does.
type Sheet struct {
	CSS     string
	Classes string
	Sources []string
}

var importPath = regexp.MustCompile(`@import "([^"]+)"`)

// Compile builds s with Tailwind's base, as dist/ and the tests' pages are.
func Compile(s Sheet) (string, error) {
	dir, err := os.MkdirTemp("", "ciderui-test-")
	if err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	defer os.RemoveAll(dir)
	classes, err := json.Marshal(s.Classes)
	if err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	var entry strings.Builder
	// @external: the standalone binary's own Tailwind: dist/ byte-identical to the npm
	// CLI's build from node_modules/tailwindcss/index.css
	entry.WriteString(`@import "tailwindcss" source(none);` + "\n")
	fmt.Fprintf(&entry, "@source inline(%s);\n", classes)
	for _, src := range s.Sources {
		fmt.Fprintf(&entry, "@source %q;\n", filepath.Join(Root, src))
	}
	entry.WriteString(importPath.ReplaceAllStringFunc(s.CSS, func(m string) string {
		return fmt.Sprintf("@import %q", filepath.Join(Root, importPath.FindStringSubmatch(m)[1]))
	}))
	entry.WriteString("\n")
	path := filepath.Join(dir, "entry.css")
	if err := os.WriteFile(path, []byte(entry.String()), 0o644); err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	return tailwind(path)
}

// OwnSheet is a site's own utilities' stylesheet: the checkout's utilities.css,
// built with Tailwind for classes.
func OwnSheet(classes string) (string, error) {
	dir, err := os.MkdirTemp("", "ciderui-own-")
	if err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "entry.css")
	if err := os.WriteFile(path, pinned.UtilitiesEntry(filepath.Join(Root, "utilities.css"), []string{fmt.Sprintf("inline(%q)", classes)}), 0o644); err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	return tailwind(path)
}

// Shipped is dist/<set>.css, what a site's page links.
func Shipped(set string) (string, error) {
	b, err := os.ReadFile(filepath.Join(Root, "dist", set+".css"))
	if err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	return string(b), nil
}

var utilityName = regexp.MustCompile(`(?m)^@utility (cui-[\w-]+)`)

// Dist is dist/<set>.css as the sources build it: Tailwind's base, Inter, the set's
// parts, every class of the type and every utility an example writes, so a page
// copying an example needs nothing more. Any other utility is the site's own
// stylesheet's (utilities.css).
func Dist(set string) (string, error) {
	all, err := Parts()
	if err != nil {
		return "", err
	}
	sheets := []string{"type.css"}
	imports := []string{"fonts.css", "index.css"}
	var sources []string
	for _, p := range all {
		if p.Set != set {
			continue
		}
		sheets = append(sheets, p.Dir+"/"+p.Name+".css")
		imports = append(imports, p.Dir+"/"+p.Name+".css")
		sources = append(sources, p.Dir+"/examples")
	}
	imports = append(imports, set+".css")
	var classes []string
	for _, sheet := range sheets {
		b, err := os.ReadFile(filepath.Join(Root, sheet))
		if err != nil {
			return "", fmt.Errorf("build: %w", err)
		}
		for _, m := range utilityName.FindAllSubmatch(b, -1) {
			classes = append(classes, string(m[1]))
		}
	}
	var css []string
	for _, sheet := range imports {
		css = append(css, fmt.Sprintf("@import %q;", sheet))
	}
	return Compile(Sheet{CSS: strings.Join(css, "\n"), Classes: strings.Join(classes, " "), Sources: sources})
}

// ExampleFiles are the part's examples' file names, examples/<id>.html, sorted.
func ExampleFiles(p Part) ([]string, error) {
	entries, err := os.ReadDir(filepath.Join(Root, p.Dir, "examples"))
	if err != nil {
		return nil, fmt.Errorf("build: %w", err)
	}
	var files []string
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".html") {
			files = append(files, e.Name())
		}
	}
	slices.Sort(files)
	return files, nil
}

// ExampleHTML is the part's example file as a page holds it: a console part's
// inside the console's page.
func ExampleHTML(p Part, file string) (string, error) {
	b, err := os.ReadFile(filepath.Join(Root, p.Dir, "examples", file))
	if err != nil {
		return "", fmt.Errorf("build: %w", err)
	}
	if p.Set == "console" {
		return `<div class="cui-console-page">` + string(b) + `</div>`, nil
	}
	return string(b), nil
}
