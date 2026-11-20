package zola

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/at-least/ciderui/internal/build"
)

const (
	nav   = `[{"text": "Store", "href": "/store"}, {"text": "Laptops", "href": "/laptops"}, {"text": "Tablets", "href": "/tablets"}, {"text": "Support", "href": "/support"}]`
	words = `aria_label="Global" skip_label="Skip to Content" open_aria_label="Menu" close_aria_label="Close"`
)

var markupCases = []struct{ example, tera string }{
	{"globalnav/globalnav", `{{ <cui_globalnav wordmark="Cider" home="/" links={` + nav + `} ` + words + ` /> }}`},
	{"footer/footer", `{{ <cui_footer copyright="Example Inc." href="/company/" /> }}`},
	{"footer/directory", `{{ <cui_footer copyright="Example Inc." href="/company/" directory_aria_label="Directory" directory={[{"title": "Products", "links": [{"text": "Store", "href": "/store/"}, {"text": "Laptops", "href": "/laptops/"}, {"text": "Phones", "href": "/phones/"}]}, {"title": "Account", "links": [{"text": "Your Account", "href": "/account/"}, {"text": "Order History", "href": "/orders/"}]}]} /> }}`},
	{"list/list", `{% <cui_list> %}` + strings.Join([]string{
		`{{ <cui_list_row title="Software keyboard fails to appear in a playground app" href="#keyboard" detail="The on-screen keyboard does not appear when editing a text field in an app made and run in Swift Playgrounds. It happens in a new, empty app too, with no code of its own." aside="3 replies · 8m" /> }}`,
		`{{ <cui_list_row title="Enrollment still pending after ten days" href="#enrollment" detail="Payment went through and the ID was sent the same day." aside="12 replies · 2h" /> }}`,
		`{{ <cui_list_row title="Widgets stop refreshing after an update" href="#widgets" /> }}`,
	}, "\n") + `{% </cui_list> %}`},
}

// @invariant: an example read as the kit's Go tests read it (go/internal/example's
// Check): each line trimmed, the lines joined
func trimmed(html string) string {
	var b strings.Builder
	for line := range strings.SplitSeq(html, "\n") {
		b.WriteString(strings.TrimSpace(line))
	}
	return b.String()
}

// @invariant: every case on one page, each between its own begin and end comment:
// one build for all, and a case's markup cut out exactly
func TestComponentsWriteTheExamples(t *testing.T) {
	var page strings.Builder
	for _, c := range markupCases {
		fmt.Fprintf(&page, "<!--begin %s-->\n%s\n<!--end %s-->\n", c.example, c.tera, c.example)
	}
	out := buildSite(t, map[string]string{
		"content/parity/_index.md": "+++\ntitle = \"Parity\"\ntemplate = \"parity.html\"\n+++\n",
		"templates/parity.html":    page.String(),
	})
	built, err := os.ReadFile(filepath.Join(out, "parity/index.html"))
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range markupCases {
		_, after, ok := strings.Cut(string(built), "<!--begin "+c.example+"-->")
		if !ok {
			t.Fatalf("%s: no begin comment in the built page", c.example)
		}
		markup, _, ok := strings.Cut(after, "<!--end "+c.example+"-->")
		if !ok {
			t.Fatalf("%s: no end comment in the built page", c.example)
		}
		part, file, _ := strings.Cut(c.example, "/")
		example, err := os.ReadFile(filepath.Join(build.Root, "components/site", part, "examples", file+".html"))
		if err != nil {
			t.Fatal(err)
		}
		if got, want := trimmed(markup), trimmed(string(example)); got != want {
			at := 0
			for at < min(len(got), len(want)) && got[at] == want[at] {
				at++
			}
			t.Errorf("%s: first differs at byte %d:\n got: %q\nwant: %q", c.example, at, got[max(0, at-40):min(len(got), at+60)], want[max(0, at-40):min(len(want), at+60)])
		}
	}
}
