package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

func TestFieldRendersTheExamples(t *testing.T) {
	example.Check(t, site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "email", Value: "", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/field.html")
	example.Check(t, site.Field{ID: "zip", Type: "text", Name: "zip", Label: "ZIP Code", Autocomplete: "postal-code", Value: "10001", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/filled.html")
	example.Check(t, site.Field{ID: "recipient", Type: "email", Name: "recipient", Label: "Recipient Email Address", Autocomplete: "email", Value: "", Error: "Please enter a valid email address.", Invalid: false, Help: "", Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/error.html")
	example.Check(t, site.Field{ID: "sender", Type: "text", Name: "sender", Label: "Sender Name", Autocomplete: "name", Value: "", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: h.Disabled(), Class: ""}, "site/field/examples/disabled.html")
	example.Check(t, site.Field{ID: "amount", Type: "text", Name: "amount", Label: "Other Amount", Autocomplete: "", Value: "", Error: "", Invalid: false, Help: "", Prefix: "$", Attrs: nil, Class: ""}, "site/field/examples/prefix.html")
	example.Check(t, site.Field{ID: "account", Type: "email", Name: "account", Label: "Account Email", Autocomplete: "email", Value: "", Error: "", Invalid: false, Help: "The same one you use in the app.", Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/help.html")
}

func TestFieldItsPrefixAndErrorBothDescribeIt(t *testing.T) {
	got := example.Render(site.Field{ID: "phone", Type: "tel", Name: "phone", Label: "Phone", Autocomplete: "", Value: "", Error: "Please enter a number.", Invalid: false, Help: "The number you signed up with.", Prefix: "+1", Attrs: nil, Class: ""})
	for _, want := range []string{
		`aria-describedby="phone-prefix phone-help phone-message">`,
		`</div><p class="cui-field-help" id="phone-help">The number you signed up with.</p>`,
		`<p class="cui-field-message" id="phone-message">`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}

// a value too short to judge: the state with no words to show
func TestFieldIsInvalidWithNoMessage(t *testing.T) {
	got := example.Render(site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "", Value: "", Error: "", Invalid: true, Help: "", Prefix: "", Attrs: nil, Class: ""})
	if !strings.Contains(got, `aria-invalid="true"`) || strings.Contains(got, `cui-field-message`) {
		t.Errorf("want aria-invalid and no message under the field: %s", got)
	}
}

func TestFieldPanics(t *testing.T) {
	example.Panics(t, func() site.Field {
		return site.Field{ID: "x", Type: "text", Name: "x", Label: "x", Autocomplete: "", Value: "", Error: "", Invalid: false, Help: "", Prefix: "", Attrs: nil, Class: ""}
	}, []example.Break[site.Field]{
		{Want: "site.Field needs an ID", Change: func(p *site.Field) { p.ID = "" }},
		{Want: "site.Field needs a Name", Change: func(p *site.Field) { p.Name = "" }},
		{Want: "site.Field needs a Label", Change: func(p *site.Field) { p.Label = "" }},
	})
}
