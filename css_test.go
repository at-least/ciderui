package ciderui

import (
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
)

// a part's rule over another's of the same specificity relies on this order
func TestEachSetImportsItsPartsInOrder(t *testing.T) {
	for _, set := range []string{"site", "console"} {
		var imported []string
		entry, err := os.ReadFile(set + ".css")
		if err != nil {
			t.Fatal(err)
		}
		for _, line := range strings.Split(string(entry), "\n") {
			if rest, ok := strings.CutPrefix(line, `@import "./`); ok {
				imported = append(imported, strings.TrimSuffix(rest, `";`))
			}
		}
		var parts []string
		for _, p := range cssFiles() {
			if strings.HasPrefix(p, "components/"+set+"/") {
				parts = append(parts, p)
			}
		}
		if len(parts) == 0 {
			t.Fatalf("%s: no parts", set)
		}
		want := append([]string{"index.css"}, parts...)
		if !slices.Equal(imported, want) {
			t.Errorf("%s.css imports %v, want %v", set, imported, want)
		}
	}
}

// a part's stylesheet opens with its rules: what a header would say is its
// examples', its rules' and its test's (DESIGN.md, rule 14)
func TestNoPartHasAHeader(t *testing.T) {
	for _, p := range cssFiles() {
		if strings.HasPrefix(p, "components/") && strings.HasPrefix(strings.TrimSpace(css(p)), "/*") {
			t.Errorf("%s opens with a comment", p)
		}
	}
}

func TestTypeStepsAtTheMdBreakpoint(t *testing.T) {
	md := regexp.MustCompile(`--breakpoint-md: (\d+px);`).FindStringSubmatch(css("theme.css"))[1]
	steps := regexp.MustCompile(`width >= (\d+px)`).FindAllStringSubmatch(css("type.css"), -1)
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
	for _, p := range cssFiles() {
		if !strings.HasPrefix(p, "components/") {
			continue
		}
		text := css(p)
		part := strings.TrimSuffix(filepath.Base(p), ".css")
		at := strings.Index(text, "\n  .cui-"+part+" {\n")
		if at < 0 {
			continue
		}
		block, _, _ := strings.Cut(text[at+1:], "\n  }\n")
		for _, line := range strings.Split(block, "\n") {
			if line == "    margin: 0;" {
				t.Errorf("%s: .cui-%s sets %s", p, part, strings.TrimSpace(line))
			}
		}
	}
}
