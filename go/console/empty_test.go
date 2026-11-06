package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestEmptyRendersTheExamples(t *testing.T) {
	small := console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Backed: true, Icon: lucide.ChartLine, Heading: "No Data", Caption: "Nothing was reported in this period.", Class: "", Children: nil}
	example.Check(t, small, "console/console-empty/examples/small.html")

	large := console.ConsoleEmpty{Size: console.ConsoleEmptySizeLarge, Backed: false, Icon: lucide.FolderPlus, Heading: "No Projects", Caption: "Create a project to get a DSN for your app.", Class: "",
		Children: h.P(h.Class("cui-console-empty-actions"), h.A(h.Class("cui-console-btn"), h.Data("variant", "filled"), h.Href("#new"), g.Text("New Project")))}
	example.Check(t, large, "console/console-empty/examples/large.html")
}

func TestEmptyPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleEmpty {
		return console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Backed: false, Icon: lucide.Icon{}, Heading: "x", Caption: "", Class: "", Children: nil}
	}, []example.Break[console.ConsoleEmpty]{
		{Want: "console.ConsoleEmpty needs a Heading", Change: func(p *console.ConsoleEmpty) { p.Heading = "" }},
		{Want: "console.ConsoleEmpty needs a Size", Change: func(p *console.ConsoleEmpty) { p.Size = console.ConsoleEmptySize{} }},
	})
}
