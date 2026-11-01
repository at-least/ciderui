package site_test

import (
	"os"

	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func ExampleField() {
	page := h.Form(site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "email", Value: "", Error: "", Disabled: false, Prefix: "", Class: ""})
	_ = page.Render(os.Stdout)
	// Output: <form><div class="cui-field"><input class="cui-field-input" id="email" type="email" name="email" placeholder=" " autocomplete="email"><label class="cui-field-label" for="email">Email</label></div></form>
}
