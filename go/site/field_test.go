package site_test

import (
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/example"
	"github.com/at-least/ciderui/go/internal/glyph"
	"github.com/at-least/ciderui/go/site"
	h "maragu.dev/gomponents/html"
)

var models = []site.Option{
	{Value: "phone-3-plus", Label: "Cider Phone 3 Plus"},
	{Value: "phone-3", Label: "Cider Phone 3"},
	{Value: "phone-2", Label: "Cider Phone 2"},
}

func TestFieldRendersTheExamples(t *testing.T) {
	example.Check(t, site.Field{ID: "email", Type: "email", Name: "email", Label: "Email", Autocomplete: "email", Value: "", Help: "", Error: "", Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/field.html")
	example.Check(t, site.Field{ID: "zip", Type: "text", Name: "zip", Label: "ZIP Code", Autocomplete: "postal-code", Value: "10001", Help: "", Error: "", Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/filled.html")
	example.Check(t, site.Field{ID: "recipient", Type: "email", Name: "recipient", Label: "Recipient Email Address", Autocomplete: "email", Value: "", Help: "", Error: "Please enter a valid email address.", Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/error.html")
	example.Check(t, site.Field{ID: "sender", Type: "text", Name: "sender", Label: "Sender Name", Autocomplete: "name", Value: "", Help: "", Error: "", Prefix: "", Attrs: h.Disabled(), Class: ""}, "site/field/examples/disabled.html")
	example.Check(t, site.Field{ID: "amount", Type: "text", Name: "amount", Label: "Other Amount", Autocomplete: "", Value: "", Help: "", Error: "", Prefix: "$", Attrs: nil, Class: ""}, "site/field/examples/prefix.html")
	example.Check(t, site.Select{ID: "model", Name: "model", Label: "Choose your model", Placeholder: "Select", Options: models, Selected: "", Error: "", Attrs: nil, Class: ""}, "site/field/examples/select.html")
	example.Check(t, site.Select{ID: "model", Name: "model", Label: "Choose your model", Placeholder: "Select", Options: models, Selected: "phone-3", Error: "", Attrs: nil, Class: ""}, "site/field/examples/selected.html")
	example.Check(t, site.TextArea{ID: "comments", Name: "comments", Label: "Comments", Value: "", Error: "", Attrs: nil, Class: ""}, "site/field/examples/textarea.html")
}

func TestSelectsErrorAndDisabled(t *testing.T) {
	got := example.Render(site.Select{ID: "state", Name: "state", Label: "State", Placeholder: "Select", Options: []site.Option{{Value: "CA", Label: "California"}}, Selected: "", Error: "Please choose a state.", Attrs: h.Disabled(), Class: ""})
	if !strings.Contains(got, `<select class="cui-field-input" id="state" name="state" aria-invalid="true" aria-describedby="state-message" disabled>`) || !strings.HasSuffix(got, `Please choose a state.</p>`) {
		t.Errorf("got %s", got)
	}
}

func TestTextAreasTextErrorAndDisabled(t *testing.T) {
	got := example.Render(site.TextArea{ID: "note", Name: "note", Label: "Message", Value: "Happy <birthday> & more", Error: "Please enter a message.", Attrs: h.Disabled(), Class: "mt-4"})
	want := `<div class="cui-field mt-4"><div class="cui-field-box"><textarea class="cui-field-textarea" id="note" name="note" placeholder=" " aria-invalid="true" aria-describedby="note-message" disabled>Happy &lt;birthday&gt; &amp; more</textarea></div>` +
		`<label class="cui-field-label" for="note">Message</label></div><p class="cui-field-message" id="note-message">` +
		example.Render(glyph.Alert) + "Please enter a message.</p>"
	if got != want {
		t.Errorf("got  %s\nwant %s", got, want)
	}
}

func TestTextAreasFirstNewlineKept(t *testing.T) {
	got := example.Render(site.TextArea{ID: "note", Name: "note", Label: "Message", Value: "\nSecond line", Error: "", Attrs: nil, Class: ""})
	if !strings.Contains(got, ">\n\nSecond line</textarea>") {
		t.Errorf("got %q", got)
	}
}

func TestSelectWritesItsPlaceholderInThePagesLanguage(t *testing.T) {
	got := example.Render(site.Select{ID: "x", Name: "x", Label: "x", Placeholder: "請選擇", Options: models, Selected: "", Error: "", Attrs: nil, Class: ""})
	if !strings.Contains(got, `<option value="" selected>請選擇</option>`) {
		t.Errorf("got %s", got)
	}
}

func TestFieldRendersTheHelpExample(t *testing.T) {
	example.Check(t, site.Field{ID: "account", Type: "email", Name: "account", Label: "Account Email", Autocomplete: "email", Value: "", Help: "The same one you use in the app.", Error: "", Prefix: "", Attrs: nil, Class: ""}, "site/field/examples/help.html")
}

func TestFieldItsPrefixHelpAndErrorAllDescribeIt(t *testing.T) {
	got := example.Render(site.Field{ID: "phone", Type: "tel", Name: "phone", Label: "Phone", Autocomplete: "", Value: "", Help: "With its area code.", Error: "Please enter a valid number.", Prefix: "+1", Attrs: nil, Class: ""})
	for _, want := range []string{
		`aria-describedby="phone-prefix phone-help phone-message">`,
		`</div><p class="cui-field-help" id="phone-help">With its area code.</p><p class="cui-field-message" id="phone-message">`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("%s has no %s", got, want)
		}
	}
}

func TestFieldPanics(t *testing.T) {
	example.Panics(t, func() site.Field {
		return site.Field{ID: "x", Type: "text", Name: "x", Label: "x", Autocomplete: "", Value: "", Help: "", Error: "", Prefix: "", Attrs: nil, Class: ""}
	}, []example.Break[site.Field]{
		{Want: "site.Field needs an ID", Change: func(p *site.Field) { p.ID = "" }},
		{Want: "site.Field needs a Name", Change: func(p *site.Field) { p.Name = "" }},
		{Want: "site.Field needs a Label", Change: func(p *site.Field) { p.Label = "" }},
	})
}

func TestSelectPanics(t *testing.T) {
	example.Panics(t, func() site.Select {
		return site.Select{ID: "x", Name: "x", Label: "x", Placeholder: "x", Options: []site.Option{{Value: "a", Label: "A"}}, Selected: "", Error: "", Attrs: nil, Class: ""}
	}, []example.Break[site.Select]{
		{Want: "site.Select needs an ID", Change: func(p *site.Select) { p.ID = "" }},
		{Want: "site.Select needs a Name", Change: func(p *site.Select) { p.Name = "" }},
		{Want: "site.Select needs a Label", Change: func(p *site.Select) { p.Label = "" }},
		{Want: "site.Select needs a Placeholder", Change: func(p *site.Select) { p.Placeholder = "" }},
		{Want: "site.Select needs Options", Change: func(p *site.Select) { p.Options = nil }},
		{Want: "site.Select's Options need values: the empty one is its empty choice's", Change: func(p *site.Select) { p.Options[0].Value = "" }},
		{Want: "site.Select's Selected value is none of its Options'", Change: func(p *site.Select) { p.Selected = "b" }},
	})
}

func TestTextAreaPanics(t *testing.T) {
	example.Panics(t, func() site.TextArea {
		return site.TextArea{ID: "x", Name: "x", Label: "x", Value: "", Error: "", Attrs: nil, Class: ""}
	}, []example.Break[site.TextArea]{
		{Want: "site.TextArea needs an ID", Change: func(p *site.TextArea) { p.ID = "" }},
		{Want: "site.TextArea needs a Name", Change: func(p *site.TextArea) { p.Name = "" }},
		{Want: "site.TextArea needs a Label", Change: func(p *site.TextArea) { p.Label = "" }},
	})
}
