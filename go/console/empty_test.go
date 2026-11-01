package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestEmptyRendersTheExamples(t *testing.T) {
	empty := console.ConsoleEmpty{Size: console.ConsoleEmptySizeMedium, Backed: false, Icon: lucide.Inbox, Heading: "No Issues", Caption: "When your app reports a crash, it shows here.", Class: "",
		Children: h.P(h.Class("cui-console-empty-actions"), h.A(h.Class("cui-console-btn"), h.Href("#setup"), g.Text("Set Up Reporting")))}
	example.Check(t, empty, "console/console-empty/examples/empty.html")

	small := console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Backed: true, Icon: lucide.ChartLine, Heading: "No Data", Caption: "Nothing was reported in this period.", Class: "", Children: nil}
	example.Check(t, small, "console/console-empty/examples/small.html")

	large := console.ConsoleEmpty{Size: console.ConsoleEmptySizeLarge, Backed: false, Icon: lucide.FolderPlus, Heading: "No Projects", Caption: "Create a project to get a DSN for your app.", Class: "",
		Children: h.P(h.Class("cui-console-empty-actions"), h.A(h.Class("cui-console-btn"), h.Data("variant", "filled"), h.Href("#new"), g.Text("New Project")))}
	example.Check(t, large, "console/console-empty/examples/large.html")
}

func TestEmptyAnEmptyHeadingPanics(t *testing.T) {
	mustPanic(t, "needs a Heading", func() {
		example.Render(console.ConsoleEmpty{Size: console.ConsoleEmptySizeMedium, Backed: false, Icon: icon.Icon{}, Heading: "", Caption: "", Class: "", Children: nil})
	})
}
