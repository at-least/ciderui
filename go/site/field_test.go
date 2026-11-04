package site_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

var models = []site.Option{
	{Value: "phone-3-plus", Text: "Cider Phone 3 Plus"},
	{Value: "phone-3", Text: "Cider Phone 3"},
	{Value: "phone-2", Text: "Cider Phone 2"},
}

func TestFieldRendersTheExamples(t *testing.T) {
	example.Check(t, site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "email", Value: "", Error: "", Disabled: false, Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/field.html")
	example.Check(t, site.Field{ID: "zip", Type: "text", Name: "zip", Label: "ZIP Code", Autocomplete: "postal-code", Value: "10001", Error: "", Disabled: false, Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/filled.html")
	example.Check(t, site.Field{ID: "recipient", Type: "email", Name: "recipient", Label: "Recipient Email Address", Autocomplete: "email", Value: "", Error: "Please enter a valid email address.", Disabled: false, Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/error.html")
	example.Check(t, site.Field{ID: "sender", Type: "text", Name: "sender", Label: "Sender Name", Autocomplete: "name", Value: "", Error: "", Disabled: true, Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/disabled.html")
	example.Check(t, site.Field{ID: "amount", Type: "text", Name: "amount", Label: "Other Amount", Autocomplete: "", Value: "", Error: "", Disabled: false, Prefix: "$", Attrs: nil, Class: ""}, "site/field/examples/prefix.html")
	example.Check(t, site.Select{ID: "model", Name: "model", Label: "Choose your model", Options: models, Selected: "", Error: "", Disabled: false, Attrs: nil, Class: ""}, "site/field/examples/select.html")
	example.Check(t, site.Select{ID: "model", Name: "model", Label: "Choose your model", Options: models, Selected: "phone-3", Error: "", Disabled: false, Attrs: nil, Class: ""}, "site/field/examples/selected.html")
	example.Check(t, site.TextArea{ID: "comments", Name: "comments", Label: "Comments", Value: "", Error: "", Disabled: false, Attrs: nil, Class: ""}, "site/field/examples/textarea.html")
}

func TestFieldNoAutocompleteNoAttribute(t *testing.T) {
	got := example.Render(site.Field{ID: "code", Type: "text", Name: "code", Label: "Promo Code", Autocomplete: "", Value: "", Error: "", Disabled: false, Prefix: "", Attrs: nil, Class: "mt-4"})
	want := `<div class="cui-field mt-4"><input class="cui-field-input" id="code" type="text" name="code" placeholder=" "><label class="cui-field-label" for="code">Promo Code</label></div>`
	if got != want {
		t.Errorf("got %s", got)
	}
}

func TestFieldAPrefixAndAnErrorBothDescribeIt(t *testing.T) {
	got := example.Render(site.Field{ID: "phone", Type: "tel", Name: "phone", Label: "Phone Number", Autocomplete: "tel", Value: "", Error: "Please enter a valid phone number.", Disabled: false, Prefix: "+1", Attrs: nil, Class: ""})
	for _, want := range []string{
		`aria-invalid="true" aria-describedby="phone-prefix phone-message">`,
		`<span class="cui-field-prefix" id="phone-prefix">+1</span></div><p class="cui-field-message" id="phone-message">`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}

func TestSelectsErrorAndDisabled(t *testing.T) {
	got := example.Render(site.Select{ID: "state", Name: "state", Label: "State", Options: []site.Option{{Value: "CA", Text: "California"}}, Selected: "", Error: "Please choose a state.", Disabled: true, Attrs: nil, Class: ""})
	if !strings.Contains(got, `<select class="cui-field-input" id="state" name="state" aria-invalid="true" aria-describedby="state-message" disabled>`) || !strings.HasSuffix(got, `Please choose a state.</p>`) {
		t.Errorf("got %s", got)
	}
}

func TestTextAreasTextErrorAndDisabled(t *testing.T) {
	got := example.Render(site.TextArea{ID: "note", Name: "note", Label: "Message", Value: "Happy <birthday> & more", Error: "Please enter a message.", Disabled: true, Attrs: nil, Class: "mt-4"})
	want := `<div class="cui-field mt-4"><div class="cui-field-box"><textarea class="cui-field-textarea" id="note" name="note" placeholder=" " aria-invalid="true" aria-describedby="note-message" disabled>Happy &lt;birthday&gt; &amp; more</textarea></div>` +
		`<label class="cui-field-label" for="note">Message</label></div><p class="cui-field-message" id="note-message">` +
		example.Render(icon.Alert) + "Please enter a message.</p>"
	if got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestTextAreasFirstNewlineKept(t *testing.T) {
	got := example.Render(site.TextArea{ID: "note", Name: "note", Label: "Message", Value: "\nSecond line", Error: "", Disabled: false, Attrs: nil, Class: ""})
	if !strings.Contains(got, ">\n\nSecond line</textarea>") {
		t.Errorf("got %q", got)
	}
}

func mustPanic(t *testing.T, want string, render func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil || !strings.Contains(fmt.Sprint(r), want) {
			t.Errorf("panicked with %v, want %q", r, want)
		}
	}()
	render()
}

func TestFieldAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "needs a Label", func() {
		example.Render(site.Field{ID: "x", Type: "text", Name: "x", Label: "", Autocomplete: "", Value: "", Error: "", Disabled: false, Prefix: "", Attrs: nil, Class: ""})
	})
}

func TestSelectASelectedValueNoOptionHasPanics(t *testing.T) {
	mustPanic(t, "none of its Options", func() {
		example.Render(site.Select{ID: "m", Name: "m", Label: "Model", Options: models, Selected: "phone-1", Error: "", Disabled: false, Attrs: nil, Class: ""})
	})
}

func TestSelectAnOptionWithNoValuePanics(t *testing.T) {
	mustPanic(t, "need values", func() {
		example.Render(site.Select{ID: "m", Name: "m", Label: "Model", Options: []site.Option{{Value: "", Text: "None"}}, Selected: "", Error: "", Disabled: false, Attrs: nil, Class: ""})
	})
}

func TestFieldAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "Field needs an ID", func() {
		example.Render(site.Field{ID: "", Type: "text", Name: "x", Label: "x", Autocomplete: "", Value: "", Error: "", Disabled: false, Prefix: "", Attrs: nil, Class: ""})
	})
}

func TestSelectAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "Select needs an ID", func() {
		example.Render(site.Select{ID: "", Name: "m", Label: "Model", Options: models, Selected: "", Error: "", Disabled: false, Attrs: nil, Class: ""})
	})
}

func TestSelectAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "Select needs a Label", func() {
		example.Render(site.Select{ID: "m", Name: "m", Label: "", Options: models, Selected: "", Error: "", Disabled: false, Attrs: nil, Class: ""})
	})
}

func TestSelectNoOptionsPanics(t *testing.T) {
	mustPanic(t, "Select needs Options", func() {
		example.Render(site.Select{ID: "m", Name: "m", Label: "Model", Options: nil, Selected: "", Error: "", Disabled: false, Attrs: nil, Class: ""})
	})
}

func TestTextAreaAnEmptyIDPanics(t *testing.T) {
	mustPanic(t, "TextArea needs an ID", func() {
		example.Render(site.TextArea{ID: "", Name: "note", Label: "Message", Value: "", Error: "", Disabled: false, Attrs: nil, Class: ""})
	})
}

func TestTextAreaAnEmptyLabelPanics(t *testing.T) {
	mustPanic(t, "TextArea needs a Label", func() {
		example.Render(site.TextArea{ID: "note", Name: "note", Label: "", Value: "", Error: "", Disabled: false, Attrs: nil, Class: ""})
	})
}

// htmx's attributes, inputmode, pattern and required go on the control itself
func TestEachControlTakesAttrs(t *testing.T) {
	attrs := g.Group{g.Attr("hx-post", "/check"), g.Attr("hx-trigger", "change"), h.Required()}
	for got, want := range map[string]string{
		example.Render(site.Field{ID: "code", Type: "text", Name: "code", Label: "Code", Autocomplete: "", Value: "", Error: "", Disabled: false, Prefix: "", Attrs: g.Group{attrs, h.InputMode("numeric"), h.Pattern("[0-9]{6}")}, Class: ""}): `<input class="cui-field-input" id="code" type="text" name="code" placeholder=" " hx-post="/check" hx-trigger="change" required inputmode="numeric" pattern="[0-9]{6}">`,
		example.Render(site.Select{ID: "state", Name: "state", Label: "State", Options: []site.Option{{Value: "CA", Text: "California"}}, Selected: "", Error: "", Disabled: false, Attrs: attrs, Class: ""}):                                   `<select class="cui-field-input" id="state" name="state" hx-post="/check" hx-trigger="change" required><option value="" selected>`,
		example.Render(site.TextArea{ID: "note", Name: "note", Label: "Message", Value: "Hi", Error: "", Disabled: false, Attrs: attrs, Class: ""}):                                                                                             `<textarea class="cui-field-textarea" id="note" name="note" placeholder=" " hx-post="/check" hx-trigger="change" required>Hi</textarea>`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}
