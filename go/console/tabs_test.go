package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
)

func TestTabsRendersTheExamples(t *testing.T) {
	underlined := []console.ConsoleTab{
		{Label: "Overview", Href: "#overview", Current: true},
		{Label: "Issues", Href: "#issues", Current: false},
		{Label: "Releases", Href: "#releases", Current: false},
		{Label: "Jobs", Href: "", Current: false},
	}
	example.Check(t, console.ConsoleTabs{AriaLabel: "Project", Variant: console.ConsoleTabsVariantUnderlined, Tabs: underlined, Class: ""}, "console/console-tabs/examples/tabs.html")
	filled := []console.ConsoleTab{
		{Label: "24h", Href: "#24h", Current: false},
		{Label: "7d", Href: "#7d", Current: true},
		{Label: "30d", Href: "#30d", Current: false},
		{Label: "90d", Href: "#90d", Current: false},
	}
	example.Check(t, console.ConsoleTabs{AriaLabel: "Period", Variant: console.ConsoleTabsVariantFilled, Tabs: filled, Class: ""}, "console/console-tabs/examples/filled.html")
}

func TestTabsTwoCurrentTabsPanic(t *testing.T) {
	tabs := []console.ConsoleTab{{Label: "A", Href: "#a", Current: true}, {Label: "B", Href: "#b", Current: true}}
	mustPanic(t, "one current tab at most", func() {
		example.Render(console.ConsoleTabs{AriaLabel: "X", Variant: console.ConsoleTabsVariantUnderlined, Tabs: tabs, Class: ""})
	})
}

// ConsoleTabsVariant has no variant that writes no attribute
func TestTabsTheZeroVariantPanics(t *testing.T) {
	mustPanic(t, "needs a Variant", func() {
		example.Render(console.ConsoleTabs{AriaLabel: "X", Variant: console.ConsoleTabsVariant{}, Tabs: nil, Class: ""})
	})
}

func TestTabsAnEmptyAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleTabs needs an AriaLabel", func() {
		example.Render(console.ConsoleTabs{AriaLabel: "", Variant: console.ConsoleTabsVariantUnderlined, Tabs: nil, Class: ""})
	})
}

func TestTabsATabWithAnEmptyLabelPanics(t *testing.T) {
	tabs := []console.ConsoleTab{{Label: "", Href: "#a", Current: false}}
	mustPanic(t, "ConsoleTab needs a Label", func() {
		example.Render(console.ConsoleTabs{AriaLabel: "X", Variant: console.ConsoleTabsVariantUnderlined, Tabs: tabs, Class: ""})
	})
}

func pressedButtons() []console.ConsoleTabButton {
	return []console.ConsoleTabButton{
		{Label: "None", Value: "none", Pressed: false},
		{Label: "Keep", Value: "keep", Pressed: true},
		{Label: "Redraw", Value: "redraw", Pressed: false},
	}
}

func TestTabButtonsRendersTheExample(t *testing.T) {
	example.Check(t, console.ConsoleTabButtons{AriaLabel: "Mark", Name: "mark", Buttons: pressedButtons(), Class: ""}, "console/console-tabs/examples/pressed.html")
}

func TestTabButtonsTwoPressedPanic(t *testing.T) {
	buttons := pressedButtons()
	buttons[0].Pressed = true
	mustPanic(t, "one pressed button at most", func() {
		example.Render(console.ConsoleTabButtons{AriaLabel: "Mark", Name: "mark", Buttons: buttons, Class: ""})
	})
}

func TestTabButtonsAnEmptyAriaLabelPanics(t *testing.T) {
	mustPanic(t, "ConsoleTabButtons needs an AriaLabel", func() {
		example.Render(console.ConsoleTabButtons{AriaLabel: "", Name: "mark", Buttons: nil, Class: ""})
	})
}

func TestTabButtonsAnEmptyNamePanics(t *testing.T) {
	mustPanic(t, "ConsoleTabButtons needs a Name", func() {
		example.Render(console.ConsoleTabButtons{AriaLabel: "Mark", Name: "", Buttons: nil, Class: ""})
	})
}

func TestTabButtonsAButtonWithAnEmptyLabelPanics(t *testing.T) {
	buttons := []console.ConsoleTabButton{{Label: "", Value: "keep", Pressed: false}}
	mustPanic(t, "ConsoleTabButton needs a Label", func() {
		example.Render(console.ConsoleTabButtons{AriaLabel: "Mark", Name: "mark", Buttons: buttons, Class: ""})
	})
}

func TestTabButtonsAButtonWithAnEmptyValuePanics(t *testing.T) {
	buttons := []console.ConsoleTabButton{{Label: "Keep", Value: "", Pressed: false}}
	mustPanic(t, "ConsoleTabButton needs a Value", func() {
		example.Render(console.ConsoleTabButtons{AriaLabel: "Mark", Name: "mark", Buttons: buttons, Class: ""})
	})
}
