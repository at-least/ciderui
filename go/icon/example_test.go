package icon_test

import (
	"os"

	"github.com/at-least/ciderui/go/icon"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ExampleIcon_Class() {
	page := h.A(h.Href("/iphone/"), g.Text("Learn more"), icon.More.Class("cui-link-icon"))
	_ = page.Render(os.Stdout)
	// Output: <a href="/iphone/">Learn more<svg class="cui-link-icon" aria-hidden="true" width="0.442em" height="0.54em" viewBox="0 0 442 540"><path d="M107.7 47.7 332.3 272.5 107.7 497.3" fill="none" stroke="currentColor" stroke-width="67.9" stroke-linecap="round" stroke-linejoin="round"/></svg></a>
}
