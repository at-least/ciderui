package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	h "maragu.dev/gomponents/html"
)

func TestAppearanceRendersTheExample(t *testing.T) {
	example.Check(t, h.Div(h.Class("cui-console-page"), console.ConsoleAppearance{Chosen: console.ConsoleAppearanceChoiceLight, Class: ""}), "console/console-appearance/examples/appearance.html")
}

func TestAppearanceNoneChosenChecksNone(t *testing.T) {
	got := example.Render(console.ConsoleAppearance{Chosen: console.ConsoleAppearanceChoice{}, Class: "hx-only"})
	if !strings.HasPrefix(got, `<fieldset class="cui-console-appearance hx-only">`) || strings.Contains(got, "checked") {
		t.Errorf("got %s", got)
	}
}
