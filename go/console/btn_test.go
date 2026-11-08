package console_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func consoleBtn(label string) console.ConsoleBtn {
	return console.ConsoleBtn{Href: "", Label: label, AriaLabel: "", Icon: lucide.Icon{}, Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: ""}
}

func row(children ...g.Node) g.Node {
	return h.Div(h.Class("flex flex-wrap items-center gap-3"), g.Group(children))
}

func TestConsoleBtnRendersTheExamples(t *testing.T) {
	withVariant := func(label string, v console.ConsoleBtnVariant, destructive bool) console.ConsoleBtn {
		b := consoleBtn(label)
		b.Variant = v
		b.Destructive = destructive
		return b
	}
	gray, filled, plain := console.ConsoleBtnVariantGray, console.ConsoleBtnVariantFilled, console.ConsoleBtnVariantPlain
	sized := func(label string, size console.ConsoleBtnSize) console.ConsoleBtn {
		b := consoleBtn(label)
		b.Size = size
		return b
	}
	disabled := func(b console.ConsoleBtn) console.ConsoleBtn {
		b.Attrs = h.Disabled()
		return b
	}
	copyIcon := func(size console.ConsoleBtnSize, capsule bool) console.ConsoleBtn {
		b := sized("", size)
		b.AriaLabel = "Copy"
		b.Icon = lucide.Copy
		if capsule {
			b.Attrs = h.Data("shape", "capsule")
		}
		return b
	}
	labelled := consoleBtn("Copy")
	labelled.Icon = lucide.Copy
	compare := consoleBtn("Compare")
	compare.Attrs = h.Data("shape", "capsule")
	medium := console.ConsoleBtnSizeMedium
	for _, c := range []struct {
		path string
		node g.Node
	}{
		{"variants", row(withVariant("Compare", gray, false), withVariant("Resolve", filled, false), withVariant("See All", plain, false), withVariant("Delete", gray, true), withVariant("Delete Project", filled, true), withVariant("Remove", plain, true))},
		{"sizes", row(sized("Small", console.ConsoleBtnSizeSmall), sized("Medium", medium), sized("Large", console.ConsoleBtnSizeLarge))},
		{"disabled", row(disabled(withVariant("Compare", gray, false)), disabled(withVariant("Resolve", filled, false)), disabled(withVariant("See All", plain, false)), disabled(withVariant("Delete", gray, true)))},
		{"icon", row(copyIcon(medium, false), labelled)},
		{"capsule", row(copyIcon(console.ConsoleBtnSizeSmall, true), copyIcon(medium, true), copyIcon(console.ConsoleBtnSizeLarge, true), compare)},
	} {
		example.Check(t, c.node, "console/console-btn/examples/"+c.path+".html")
	}
}

// an icon-only link is named by its AriaLabel and carries the page's attributes,
// as a <button> does
func TestConsoleBtnAnIconLinkTakesItsAriaLabelAndAttrs(t *testing.T) {
	link := consoleBtn("")
	link.Href, link.AriaLabel, link.Icon, link.Attrs = "/export.csv", "Download", lucide.Download, g.Attr("download")
	want := `<a class="cui-console-btn cui-console-btn-icon" href="/export.csv" aria-label="Download" download><svg`
	if got := example.Render(link); !strings.HasPrefix(got, want) {
		t.Errorf("the link is not named Download with its download attribute:\n got: %s\nwant: %s…", got, want)
	}
}

func TestConsoleBtnPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleBtn { return consoleBtn("Save") }, []example.Break[console.ConsoleBtn]{
		{Want: "console.ConsoleBtn has an Href and Submit: a link sends no form", Change: func(b *console.ConsoleBtn) { b.Href = "/x"; b.Submit = true }},
		{Want: "console.ConsoleBtn needs a Label, or an Icon and an AriaLabel", Change: func(b *console.ConsoleBtn) { b.Label = "" }},
		{Want: "console.ConsoleBtn needs a Label, or an Icon and an AriaLabel", Change: func(b *console.ConsoleBtn) { b.Label = ""; b.AriaLabel = "x" }},
		{Want: "console.ConsoleBtn needs a Label, or an Icon and an AriaLabel", Change: func(b *console.ConsoleBtn) { b.Label = ""; b.Icon = lucide.Copy }},
		{Want: "console.ConsoleBtn has a Label and an AriaLabel: its words name it", Change: func(b *console.ConsoleBtn) { b.AriaLabel = "x" }},
	})
}
