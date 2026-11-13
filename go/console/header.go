package console

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func header(heading g.Node, detail string, controls g.Node) g.Node {
	return h.Header(h.Class("cui-console-header"),
		h.Div(heading, optionalText(h.P, "cui-console-header-detail", detail)),
		controls,
	)
}
