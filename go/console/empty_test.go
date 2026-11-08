package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
)

func TestEmptyRendersTheExamples(t *testing.T) {
	small := console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Icon: lucide.ChartLine, Title: "No Data", Body: "Nothing was reported in this period.", Class: ""}
	example.Check(t, small, "console/console-empty/examples/small.html")

	large := console.ConsoleEmpty{Size: console.ConsoleEmptySizeLarge, Icon: lucide.FolderPlus, Title: "No Projects", Body: "Create a project to get a DSN for your app.", Class: ""}
	example.Check(t, large, "console/console-empty/examples/large.html")
}

func TestEmptyPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleEmpty {
		return console.ConsoleEmpty{Size: console.ConsoleEmptySizeSmall, Icon: lucide.Icon{}, Title: "x", Body: "", Class: ""}
	}, []example.Break[console.ConsoleEmpty]{
		{Want: "console.ConsoleEmpty needs a Title", Change: func(p *console.ConsoleEmpty) { p.Title = "" }},
		{Want: "console.ConsoleEmpty needs a Size: ConsoleEmptySizeSmall (a card) or ConsoleEmptySizeLarge (a whole page)", Change: func(p *console.ConsoleEmpty) { p.Size = console.ConsoleEmptySize{} }},
	})
}
