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
	return console.ConsoleBtn{Href: "", Label: label, Submit: false, Variant: console.ConsoleBtnVariantGray, Size: console.ConsoleBtnSizeMedium, Destructive: false, Attrs: nil, Class: ""}
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
	// @invariant: a console button is words alone; a page's icon button is
	// these classes on its own <button>, written here as a page writes them
	copyIcon := func(size string) g.Node {
		return h.Button(h.Class("cui-console-btn cui-console-btn-icon"), g.If(size != "", h.Data("size", size)), h.Data("shape", "capsule"), h.Type("button"), h.Aria("label", "Copy"), lucide.Copy)
	}
	compare := consoleBtn("Compare")
	compare.Attrs = h.Data("shape", "capsule")
	busy := withVariant("Save", filled, false)
	busy.Submit, busy.Attrs = true, h.Aria("busy", "true")
	medium := console.ConsoleBtnSizeMedium
	for _, c := range []struct {
		path string
		node g.Node
	}{
		{"variants", row(withVariant("Compare", gray, false), withVariant("Resolve", filled, false), withVariant("See All", plain, false), withVariant("Delete", gray, true), withVariant("Delete Project", filled, true), withVariant("Remove", plain, true))},
		{"sizes", row(sized("Small", console.ConsoleBtnSizeSmall), sized("Medium", medium), sized("Large", console.ConsoleBtnSizeLarge))},
		{"disabled", row(disabled(withVariant("Compare", gray, false)), disabled(withVariant("Resolve", filled, false)), disabled(withVariant("See All", plain, false)), disabled(withVariant("Delete", gray, true)))},
		{"icon", copyIcon("")},
		{"busy", busy},
		{"capsule", row(copyIcon("sm"), copyIcon(""), copyIcon("lg"), compare)},
	} {
		example.Check(t, c.node, "console/console-btn/examples/"+c.path+".html")
	}
}

func TestConsoleBtnPanics(t *testing.T) {
	example.Panics(t, func() console.ConsoleBtn { return consoleBtn("Save") }, []example.Break[console.ConsoleBtn]{
		{Want: "console.ConsoleBtn has an Href and Submit: a link sends no form", Change: func(b *console.ConsoleBtn) { b.Href = "/x"; b.Submit = true }},
		{Want: "console.ConsoleBtn needs a Label", Change: func(b *console.ConsoleBtn) { b.Label = "" }},
	})
}
