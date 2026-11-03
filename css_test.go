package ciderui

import (
	"os"
	"path/filepath"
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

// a set's entry imports each of its parts once, in the order of their names,
// and no other set's: a part's rule over another's of the same specificity
// relies on that order
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
