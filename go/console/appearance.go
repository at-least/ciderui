package console

import (
	"io"

	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// ConsoleAppearanceChoice is a choice of an [ConsoleAppearance]. Its zero value is none
// chosen yet.
type ConsoleAppearanceChoice struct{ value string }

var (
	// ConsoleAppearanceChoiceLight is Light, value="light".
	ConsoleAppearanceChoiceLight = ConsoleAppearanceChoice{"light"}
	// ConsoleAppearanceChoiceDark is Dark, value="dark".
	ConsoleAppearanceChoiceDark = ConsoleAppearanceChoice{"dark"}
	// ConsoleAppearanceChoiceAuto is the system's, value="auto".
	ConsoleAppearanceChoiceAuto = ConsoleAppearanceChoice{"auto"}
)

// ConsoleAppearance is cui-console-appearance (components/console/console-appearance):
// Light, Dark and Auto as one segmented control of radios (name="appearance"),
// Apple's developer documentation's footer control, for a site that offers
// its own appearance; the page's script applies the choice.
type ConsoleAppearance struct {
	// Chosen is the chosen one; the zero ConsoleAppearanceChoice for none yet (the
	// page's script chooses it from what it stored).
	Chosen ConsoleAppearanceChoice
	// Class is the page's utilities on the <fieldset>; "" for none.
	Class string
}

// Render writes the control.
func (a ConsoleAppearance) Render(w io.Writer) error {
	choices := []struct {
		choice ConsoleAppearanceChoice
		word   string
	}{
		{ConsoleAppearanceChoiceLight, "Light"},
		{ConsoleAppearanceChoiceDark, "Dark"},
		{ConsoleAppearanceChoiceAuto, "Auto"},
	}
	var labels g.Group
	for _, c := range choices {
		labels = append(labels, h.Label(h.Class("cui-console-appearance-choice"),
			h.Input(h.Type("radio"), h.Name("appearance"), h.Value(c.choice.value), g.If(a.Chosen == c.choice, h.Checked())),
			h.Span(g.Text(c.word)),
		))
	}
	return h.FieldSet(class.Root("cui-console-appearance", a.Class),
		h.Legend(g.Text("Appearance")),
		labels,
	).Render(w)
}
