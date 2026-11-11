package console

import (
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func pageHeader(back ConsoleContentBack, title, detail string, controls g.Node) g.Node {
	return header(g.Group{
		g.If(back.Href != "", h.A(h.Class("cui-console-link cui-console-section-header-back"), h.Href(back.Href), lucide.ChevronLeft, g.Text(back.Label))),
		h.H1(h.Class("cui-console-h2"), g.Text(title)),
	}, detail, controls)
}

func sectionHeader(title, detail string, controls g.Node) g.Node {
	return header(h.H2(h.Class("cui-console-h4-emph"), g.Text(title)), detail, controls)
}

func header(heading g.Node, detail string, controls g.Node) g.Node {
	return h.Header(h.Class("cui-console-section-header"),
		h.Div(heading, optionalText(h.P, "cui-console-section-header-detail", detail)),
		controls,
	)
}
