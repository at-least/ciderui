package ciderui

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

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

func TestAWordInACommentIsNoClass(t *testing.T) {
	classes := classesIn(`/* cui-ghost, named here alone */ .cui-real { color: red; } @utility cui-type { color: red; }`)
	if classes["cui-ghost"] {
		t.Error("cui-ghost is only in a comment")
	}
	if !classes["cui-real"] || !classes["cui-type"] {
		t.Errorf("got %v", classes)
	}
}

func TestAWordNoSelectorNamesFails(t *testing.T) {
	want := []string{`class "cui-globalnav-row-in" on <li class="cui-globalnav-row-in">: ciderui has no such class`}
	if got := CheckMarkup(`<li class="cui-globalnav-row-in">a</li>`); !slices.Equal(got, want) {
		t.Errorf("a keyframes name: got %q", got)
	}
}

func TestAClassWithNoRulePasses(t *testing.T) {
	for _, class := range classesWithNoRule {
		if got := CheckMarkup(`<div class="` + class + `">a</div>`); got != nil {
			t.Errorf("%s: %q", class, got)
		}
	}
}

// a class that gained a rule leaves the list
func TestClassesWithNoRuleHaveNoRule(t *testing.T) {
	for _, path := range cssFiles() {
		classes := classesIn(css(path))
		for _, class := range classesWithNoRule {
			if classes[class] {
				t.Errorf("%s has a rule for %s", path, class)
			}
		}
	}
}

func TestAValueNoRuleSelectsFails(t *testing.T) {
	want := []string{`data-color="purple" on <span class="cui-console-badge">: no rule of ciderui's selects it (cui-console-badge's: error, info, success, warning)`}
	if got := CheckMarkup(`<span class="cui-console-badge" data-color="purple">b</span>`); !slices.Equal(got, want) {
		t.Errorf("got %q", got)
	}
}

func TestAShapeNoRuleSelectsFails(t *testing.T) {
	want := []string{`data-shape="circle" on <button class="cui-console-btn">: no rule of ciderui's selects it (cui-console-btn's: capsule)`}
	if got := CheckMarkup(`<button class="cui-console-btn" data-shape="circle">x</button>`); !slices.Equal(got, want) {
		t.Errorf("got %q", got)
	}
}

func TestATypoInAnySelectedAttributeFails(t *testing.T) {
	want := []string{`data-orientation="vertcal" on <dl class="cui-console-dl">: no rule of ciderui's selects it (cui-console-dl's: vertical)`}
	if got := CheckMarkup(`<dl class="cui-console-dl" data-orientation="vertcal"></dl>`); !slices.Equal(got, want) {
		t.Errorf("got %q", got)
	}
	if got := CheckMarkup(`<dl class="cui-console-dl" data-orientation="vertical"></dl>`); got != nil {
		t.Errorf("the value itself: got %q", got)
	}
}

// [data-sticky] selects any value, on its own class alone
func TestAnAttributeSelectedWithNoValue(t *testing.T) {
	for _, html := range []string{
		`<table class="cui-console-table" data-sticky></table>`,
		`<table class="cui-console-table" data-sticky="true"></table>`,
		`<tr class="cui-console-listing-line" data-highlighted></tr>`,
	} {
		if got := CheckMarkup(html); got != nil {
			t.Errorf("%s: got %q", html, got)
		}
	}
	want := []string{`data-sticky="" on <div class="cui-console-card">: no rule of ciderui's selects it (cui-console-card's: none)`}
	if got := CheckMarkup(`<div class="cui-console-card" data-sticky></div>`); !slices.Equal(got, want) {
		t.Errorf("got %q", got)
	}
}

// .cui-posts:not([data-size]) is its default: the bare attribute leaves it, and
// a misspelt value would as well, with no rule of the size's own
func TestAnAttributeOnlyANotSelects(t *testing.T) {
	if got := CheckMarkup(`<ul class="cui-console-tabs" data-overflow-end></ul>`); got != nil {
		t.Errorf("got %q", got)
	}
	want := []string{`data-overflow-end="yes" on <ul class="cui-console-tabs">: no rule of ciderui's selects it (cui-console-tabs's: "")`}
	if got := CheckMarkup(`<ul class="cui-console-tabs" data-overflow-end="yes"></ul>`); !slices.Equal(got, want) {
		t.Errorf("got %q", got)
	}
	want = []string{`data-size="tiny" on <div class="cui-story">: no rule of ciderui's selects it (cui-story's: "", lg, md)`}
	if got := CheckMarkup(`<div class="cui-story" data-size="tiny"></div>`); !slices.Equal(got, want) {
		t.Errorf("got %q", got)
	}
}

// :is(.cui-console-copy-button, .cui-console-copy-inline)[data-state="copied"]
func TestAValueSelectedForEachClassOfAnIs(t *testing.T) {
	for _, class := range []string{"cui-console-copy-button", "cui-console-copy-inline"} {
		if got := CheckMarkup(`<button class="` + class + `" data-state="copied">x</button>`); got != nil {
			t.Errorf("%s: got %q", class, got)
		}
	}
}

// .cui-console-player-end > [data-control="replay"] names no class of the element
func TestAnAttributeOnlyARuleForAnotherElementSelectsPasses(t *testing.T) {
	if got := CheckMarkup(`<p class="cui-console-player-caption" data-control="caption">x</p>`); got != nil {
		t.Errorf("got %q", got)
	}
}

// :where(.cui-btn, .cui-link[data-variant="more"], .cui-console-link, …) gives
// cui-console-link no "more"
func TestAValueIsNotAnotherAlternativesClasss(t *testing.T) {
	want := []string{`data-variant="more" on <a class="cui-console-link">: no rule of ciderui's selects it (cui-console-link's: label, text)`}
	if got := CheckMarkup(`<a class="cui-console-link" data-variant="more">x</a>`); !slices.Equal(got, want) {
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

func TestWhatIsNotCideruisPasses(t *testing.T) {
	if got := CheckMarkup(`<div class="mt-2" data-variant="whatever"><!-- <b class="cui-nope"> --><script>let s = '<a class="cui-nope">';</script></div>`); got != nil {
		t.Errorf("got %q", got)
	}
	if got := CheckMarkup(`<a class="cui-console-btn" data-testid="save">x</a>`); got != nil {
		t.Errorf("a data-* no rule selects anywhere, the page's own: got %q", got)
	}
}
