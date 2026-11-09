package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestFieldRendersTheExamples(t *testing.T) {
	example.Check(t, site.Field{Type: site.FieldTypeEmail, Name: "email", Label: "Email", Autocomplete: "email", Value: "", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: nil}, "site/field/examples/field.html")
	example.Check(t, site.Field{Type: site.FieldTypeText, Name: "zip", Label: "ZIP Code", Autocomplete: "postal-code", Value: "10001", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: nil}, "site/field/examples/filled.html")
	example.Check(t, site.Field{Type: site.FieldTypeEmail, Name: "recipient", Label: "Recipient Email Address", Autocomplete: "email", Value: "", Error: "Please enter a valid email address.", Invalid: false, Help: "", Prefix: "", Attrs: nil}, "site/field/examples/error.html")
	example.Check(t, site.Field{Type: site.FieldTypeText, Name: "sender", Label: "Sender Name", Autocomplete: "name", Value: "", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: h.Disabled()}, "site/field/examples/disabled.html")
	example.Check(t, site.Field{Type: site.FieldTypeText, Name: "amount", Label: "Other Amount", Autocomplete: "", Value: "", Error: "", Invalid: false, Help: "", Prefix: "$", Attrs: nil}, "site/field/examples/prefix.html")
	example.Check(t, site.Field{Type: site.FieldTypeEmail, Name: "account", Label: "Account Email", Autocomplete: "email", Value: "", Error: "", Invalid: false, Help: "The same one you use in the app.", Prefix: "", Attrs: nil}, "site/field/examples/help.html")
}

func TestFieldItsPrefixHelpAndErrorDescribeItInThatOrder(t *testing.T) {
	got := example.Render(site.Field{Type: site.FieldTypeTel, Name: "phone", Label: "Phone", Autocomplete: "", Value: "", Error: "Please enter a number.", Invalid: false, Help: "The number you signed up with.", Prefix: "+1", Attrs: nil})
	for _, want := range []string{
		`aria-describedby="phone-prefix phone-help phone-error">`,
		`</div><p class="cui-field-help" id="phone-help">The number you signed up with.</p>`,
		`<p class="cui-field-error" id="phone-error">`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("the input is not described by its prefix, help and error in that order: no %s in %s", want, got)
		}
	}
}

func TestFieldTypeTelWritesATelInput(t *testing.T) {
	got := example.Render(site.Field{Type: site.FieldTypeTel, Name: "phone", Label: "Phone", Autocomplete: "", Value: "", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: nil})
	if !strings.Contains(got, `<input class="cui-field-input" id="phone" type="tel"`) {
		t.Errorf("a phone field does not bring up the phone keypad: no type=\"tel\" on its input in %s", got)
	}
}

func TestFieldInvalidWithNoErrorMarksItsInputAlone(t *testing.T) {
	got := example.Render(site.Field{Type: site.FieldTypeEmail, Name: "email", Label: "Email", Autocomplete: "", Value: "", Error: "", Invalid: true, Help: "", Prefix: "", Attrs: nil})
	if !strings.Contains(got, `aria-invalid="true"`) || strings.Contains(got, `cui-field-error`) {
		t.Errorf("an Invalid field with no Error does not mark its input invalid with no message under it: %s", got)
	}
}

func TestFieldPanics(t *testing.T) {
	example.Panics(t, func() site.Field {
		return site.Field{Type: site.FieldTypeText, Name: "x", Label: "x", Autocomplete: "", Value: "", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: nil}
	}, []example.Break[site.Field]{
		{Want: "site.Field needs a Name", Change: func(p *site.Field) { p.Name = "" }},
		{Want: "site.Field needs a Label", Change: func(p *site.Field) { p.Label = "" }},
	})
}
