package pages

import (
	"testing"

	"github.com/at-least/ciderui/reference/apple.com/results"
)

// TestSpecs holds the registry to the records: every spec builds (touched panics on a
// name its base lacks) and has its record in results/, and every record its spec.
func TestSpecs(t *testing.T) {
	var names []string
	for name, build := range Specs {
		names = append(names, name)
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%s: %v", name, r)
				}
			}()
			if s := build(); s.URL == "" || len(s.Probes) == 0 {
				t.Errorf("%s: no url or no probes", name)
			}
		}()
		if _, ok := results.Pages[name]; !ok {
			t.Errorf("%s: no record in results/", name)
		}
	}
	// @invariant: a record no spec reads again is a probe nobody can rerun
	for name := range results.Pages {
		if _, ok := Specs[name]; !ok {
			t.Errorf("results/%s.go: no spec in Specs", name)
		}
	}
}
