package console_test

import (
	"testing"

	"github.com/at-least/ciderui/go/console"
	"github.com/at-least/ciderui/go/internal/example"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestCopyIconsRendersTheExamples(t *testing.T) {
	card := h.Div(h.Class("cui-console-page max-w-xl"),
		h.Div(h.Class("cui-console-card grid grid-cols-1 gap-3"),
			h.Div(h.Class("cui-console-copy"),
				h.Pre(h.Class("cui-console-listing"), h.TabIndex("0"), h.Code(h.ID("copy-cli"), g.Text("sentry-cli debug-files upload --org acme --project shop ./build"))),
				h.Button(h.Type("button"), h.Class("cui-console-copy-button"), h.Aria("label", "Copy"), console.ConsoleCopyIcons{}),
			),
			h.Div(h.Class("cui-console-copy"),
				h.Pre(h.Class("cui-console-listing"), h.TabIndex("0"), h.Code(h.ID("copy-dsn"), g.Text("https://4f1c@crash.example.com/2"))),
				h.Button(h.Type("button"), h.Class("cui-console-copy-button"), h.Aria("label", "Copied"), h.Data("state", "copied"), console.ConsoleCopyIcons{}),
			),
		),
	)
	example.Check(t, card, "console/console-copy/examples/card.html")

	inline := h.Div(h.Class("cui-console-page grid gap-3"),
		h.P(g.Text("Issue ID "), h.Span(h.Class("font-mono wrap-anywhere"), h.ID("issue-id"), g.Text("60674cfe203893d8499985f493f86663")), g.Text(" "), h.Button(h.Type("button"), h.Class("cui-console-copy-inline"), h.Aria("label", "Copy issue ID"), console.ConsoleCopyIcons{})),
		h.P(g.Text("Event ID "), h.Span(h.Class("font-mono wrap-anywhere"), h.ID("event-id"), g.Text("e1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6")), g.Text(" "), h.Button(h.Type("button"), h.Class("cui-console-copy-inline"), h.Aria("label", "Copied"), h.Data("state", "copied"), console.ConsoleCopyIcons{})),
		h.P(g.Text("Release "), h.Span(h.Class("font-mono wrap-anywhere"), g.Text("60674cfe203893d8499985f493f86663"))),
	)
	example.Check(t, inline, "console/console-copy/examples/inline.html")
}
