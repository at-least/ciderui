package icon_test

import (
	"os"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/lucide"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ExampleIcon_Class() {
	var close icon.Icon = lucide.X
	page := h.A(h.Href("/iphone/"), g.Text("Close"), close.Class("inline"))
	_ = page.Render(os.Stdout)
	// Output: <a href="/iphone/">Close<svg class="inline" aria-hidden="true" width="1em" height="1em" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg></a>
}
