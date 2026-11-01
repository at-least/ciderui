package icon_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
)

func TestRendersAnSVGInEm(t *testing.T) {
	want := `<svg aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540">` +
		`<path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg>`
	if got := example.Render(icon.More); got != want {
		t.Errorf("got %s", got)
	}
}

// lucide-static's icons/x.svg: its <svg>'s paint and its paths, at 1em where
// it has 24
func TestRendersALucideIconWithLucidesPaint(t *testing.T) {
	want := `<svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">` +
		`<path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>`
	if got := example.Render(lucide.X); got != want {
		t.Errorf("got %s", got)
	}
}

func TestEscapesItsClass(t *testing.T) {
	if got := example.Render(icon.More.Class(`a"b<c&d`)); !strings.HasPrefix(got, `<svg class="a&quot;b&lt;c&amp;d"`) {
		t.Errorf("got %s", got)
	}
}
