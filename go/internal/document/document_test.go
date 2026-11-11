package document_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/at-least/ciderui/go/internal/document"
	g "maragu.dev/gomponents"
)

func page(title string, children g.Node) document.Document {
	return document.Document{Owner: "Page", Lang: "en", Title: title, Stylesheet: "/ciderui.css", Head: nil, Class: "", BodyClass: "", Children: children}
}

func TestDocumentWritesItsTitleAsText(t *testing.T) {
	var b strings.Builder
	if err := page("Tom & Jerry <3", nil).Render(&b); err != nil {
		t.Fatal(err)
	}
	if want := "<title>Tom &amp; Jerry &lt;3</title>"; !strings.Contains(b.String(), want) {
		t.Errorf("the page's title is not its text escaped, %s: %s", want, b.String())
	}
}

type failing struct{}

func (failing) Render(io.Writer) error { return errors.New("disk full") }

func TestDocumentReturnsTheErrorOfAPartThatCannotWrite(t *testing.T) {
	if err := page("Saved", failing{}).Render(io.Discard); err == nil || err.Error() != "disk full" {
		t.Errorf("a page whose part cannot write returned %v, want the part's error", err)
	}
}
