package ciderui

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// every component's examples write only words the stylesheets know
func TestEveryExamplePasses(t *testing.T) {
	paths, err := filepath.Glob("components/*/*/examples/*.html")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) <= 100 {
		t.Fatalf("%d examples read", len(paths))
	}
	for _, path := range paths {
		html, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if problems := CheckMarkup(string(html)); problems != nil {
			t.Errorf("%s: %q", path, problems)
		}
	}
}

func TestAMisspeltClassFails(t *testing.T) {
	want := []string{`class "cui-console-badg" on <span class="cui-console-badg mt-1">: ciderui has no such class`}
	if got := CheckMarkup(`<p>a</p><span class="cui-console-badg mt-1">b</span>`); !slices.Equal(got, want) {
		t.Errorf("got %q", got)
	}
}

func TestAValueNoRuleSelectsFails(t *testing.T) {
	want := []string{`data-color="purple" on <span class="cui-console-badge">: no rule of ciderui's selects it (cui-console-badge's: brand, error, info, success, warning)`}
	if got := CheckMarkup(`<span class="cui-console-badge" data-color="purple">b</span>`); !slices.Equal(got, want) {
		t.Errorf("got %q", got)
	}
}

// a value is held to every cui-* class on its element: cui-console-btn-icon's
// sizes are its own
func TestAValueOfAnotherClassOnTheElementPasses(t *testing.T) {
	if got := CheckMarkup(`<button class="cui-console-btn cui-console-btn-icon" data-size="sm">x</button>`); got != nil {
		t.Errorf("got %q", got)
	}
}

// a page's own data-* and classes, and what a script holds, are not ciderui's
func TestWhatIsNotCideruisPasses(t *testing.T) {
	if got := CheckMarkup(`<div class="mt-2" data-variant="whatever"><!-- <b class="cui-nope"> --><script>let s = '<a class="cui-nope">';</script></div>`); got != nil {
		t.Errorf("got %q", got)
	}
}
