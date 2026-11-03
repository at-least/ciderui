package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

var badgeColors = []struct {
	color console.ConsoleColor
	word  string
}{
	{console.ConsoleColorNeutral, "Unhandled"},
	{console.ConsoleColorBrand, "New"},
	{console.ConsoleColorInfo, "Info"},
	{console.ConsoleColorSuccess, "Resolved"},
	{console.ConsoleColorWarning, "Warning"},
	{console.ConsoleColorError, "Fatal"},
}

func badgeRow(count int, variant console.ConsoleBadgeVariant, small bool) g.Node {
	var badges g.Group
	for _, c := range badgeColors[:count] {
		badges = append(badges, console.ConsoleBadge{Label: c.word, Color: c.color, Variant: variant, Small: small, Class: ""})
	}
	return h.P(h.Class("flex flex-wrap gap-2"), badges)
}

func TestBadgeRendersTheExamples(t *testing.T) {
	example.Check(t, badgeRow(6, console.ConsoleBadgeVariantTinted, false), "console/console-badge/examples/badges.html")
	example.Check(t, badgeRow(6, console.ConsoleBadgeVariantGray, false), "console/console-badge/examples/gray.html")
	example.Check(t, badgeRow(6, console.ConsoleBadgeVariantPlain, false), "console/console-badge/examples/plain.html")
	example.Check(t, badgeRow(3, console.ConsoleBadgeVariantTinted, true), "console/console-badge/examples/small.html")
}

func TestBadgeAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(console.ConsoleBadge{Label: "", Color: console.ConsoleColorNeutral, Variant: console.ConsoleBadgeVariantTinted, Small: false, Class: ""})
	})
}
