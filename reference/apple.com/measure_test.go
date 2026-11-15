package main

import "testing"

// TestStrip holds strip to the Node probe's: each want is what its strip gave
// (2026-11-15).
func TestStrip(t *testing.T) {
	for head, want := range map[string]string{
		`<nav id="g" aria-label="Global">x</nav>`:                `<nav id="g" aria-label="…"></nav>`,
		`<a title="Hi there" alt="A">b</a><img placeholder="cut`: `<a title="…" alt="…"></a><img placeholder="…`,
		`<p class="x">text<span aria-label="">y`:                 `<p class="x"><span aria-label="">`,
	} {
		if got := strip(head); got != want {
			t.Errorf("strip(%q) = %q, want %q", head, got, want)
		}
	}
}
