package ciderui

import (
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
)

func TestTheKitIsEmbedded(t *testing.T) {
	paths := CSSFiles()
	for _, want := range []string{"index.css", "tokens.css", "type.css", "fonts.css", "components/site/button/button.css"} {
		if !slices.Contains(paths, want) {
			t.Errorf("%s is not embedded", want)
		}
	}
	if !strings.Contains(CSS("tokens.css"), "--cui-c-accent:") {
		t.Error("tokens.css has no --cui-c-accent")
	}
}

// a part's rule over another's of the same specificity relies on this order
func TestEachSetImportsItsPartsInOrder(t *testing.T) {
	for _, set := range []string{"site", "console", "docs"} {
		var imported []string
		for _, line := range strings.Split(CSS(set+".css"), "\n") {
			if rest, ok := strings.CutPrefix(line, `@import "./`); ok {
				imported = append(imported, strings.TrimSuffix(rest, `";`))
			}
		}
		var parts []string
		for _, p := range CSSFiles() {
			if strings.HasPrefix(p, "components/"+set+"/") {
				parts = append(parts, p)
			}
		}
		if len(parts) == 0 {
			t.Fatalf("%s: no parts", set)
		}
		if !slices.Equal(imported, parts) {
			t.Errorf("%s.css imports %v, the parts are %v", set, imported, parts)
		}
	}
}

// a part's stylesheet opens with its rules: what a header would say is its
// examples', its rules' and its test's (DESIGN.md, rule 14)
func TestNoPartHasAHeader(t *testing.T) {
	for _, p := range CSSFiles() {
		if strings.HasPrefix(p, "components/") && strings.HasPrefix(strings.TrimSpace(CSS(p)), "/*") {
			t.Errorf("%s opens with a comment", p)
		}
	}
}

// a page's md: layout and its type change at one width
func TestTypeStepsAtTheMdBreakpoint(t *testing.T) {
	md := regexp.MustCompile(`--breakpoint-md: (\d+px);`).FindStringSubmatch(CSS("theme.css"))[1]
	steps := regexp.MustCompile(`width >= (\d+px)`).FindAllStringSubmatch(CSS("type.css"), -1)
	if len(steps) == 0 {
		t.Fatal("type.css has no step")
	}
	for _, step := range steps {
		if step[1] != md {
			t.Errorf("type.css steps at %s, md: at %s", step[1], md)
		}
	}
}

// a part's margin: 0, the browser's margin undone, is in :where(), with no
// specificity: a container spacing its children (a page's content) wins, where
// an equal selector would leave it to the order the stylesheets come in. A
// margin of the part's own design keeps its class.
func TestNoPartResetsItsMarginOverAContainer(t *testing.T) {
	for _, p := range CSSFiles() {
		if !strings.HasPrefix(p, "components/") {
			continue
		}
		css := CSS(p)
		part := strings.TrimSuffix(filepath.Base(p), ".css")
		at := strings.Index(css, "\n  .cui-"+part+" {\n")
		if at < 0 {
			continue
		}
		block, _, _ := strings.Cut(css[at+1:], "\n  }\n")
		for _, line := range strings.Split(block, "\n") {
			if line == "    margin: 0;" {
				t.Errorf("%s: .cui-%s sets %s", p, part, strings.TrimSpace(line))
			}
		}
	}
}

func TestWritesTheKitsLayout(t *testing.T) {
	dir := t.TempDir()
	if err := WriteCSS(dir); err != nil {
		t.Fatal(err)
	}
	button, err := os.ReadFile(filepath.Join(dir, "components/site/button/button.css"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(button), ".cui-btn") {
		t.Error("button.css has no .cui-btn")
	}
}
