package lucide_test

import (
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
)

// lucide-static's icons/x.svg: its <svg>'s paint and its paths, at 1em where
// it has 24
func TestRendersALucideIconWithLucidesPaint(t *testing.T) {
	want := `<svg aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">` +
		`<path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>`
	if got := example.Render(lucide.X); got != want {
		t.Errorf("got %s", got)
	}
}
