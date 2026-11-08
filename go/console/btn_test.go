package console_test

import (
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

func TestConsoleBtnLinkAndSubmit(t *testing.T) {
	link := consoleBtn("Open")
	link.Href = "/x"
	if got, want := example.Render(link), `<a class="cui-console-btn" href="/x">Open</a>`; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	submit := consoleBtn("Save")
	submit.Submit = true
	if got, want := example.Render(submit), `<button class="cui-console-btn" type="submit">Save</button>`; got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestConsoleBtnPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleBtn { return consoleBtn("Save") }, []example.Break[console.ConsoleBtn]{
		{Want: "console.ConsoleBtn has an Href and Submit: a link sends no form", Change: func(b *console.ConsoleBtn) { b.Href = "/x"; b.Submit = true }},
		{Want: "console.ConsoleBtn needs a Label, or an Icon and an AriaLabel", Change: func(b *console.ConsoleBtn) { b.Label = "" }},
		{Want: "console.ConsoleBtn needs a Label, or an Icon and an AriaLabel", Change: func(b *console.ConsoleBtn) { b.Label = ""; b.AriaLabel = "x" }},
		{Want: "console.ConsoleBtn has a Label and an AriaLabel: its words name it", Change: func(b *console.ConsoleBtn) { b.AriaLabel = "x" }},
	})
}
