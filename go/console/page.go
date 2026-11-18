package console

import (
	"io"

	"github.com/at-least/ciderui"
	"github.com/at-least/ciderui/go/internal/document"
	"github.com/at-least/ciderui/go/internal/need"
	g "maragu.dev/gomponents"
)

// ConsoleScheme is the page's colors: the zero value follows the system's
// appearance (decision 35); the others hold the page to one, its <html>'s
// class.
type ConsoleScheme struct{ class string }

var (
	ConsoleSchemeSystem = ConsoleScheme{class: ""}
	ConsoleSchemeLight  = ConsoleScheme{class: "cui-light"}
	ConsoleSchemeDark   = ConsoleScheme{class: "cui-dark"}
)

// ConsolePage is a whole document, its <body> cui-console-page, linking the set's
// stylesheet: mount [ciderui.Assets] for it. Its <title> is its Content's
// Title, then, for a [ConsoleContent] with a sidebar, the sidebar's Wordmark
// ("Jobs - Acme", as apple.com writes "iPhone - Apple"), so a tab names the
// page its header names and the site it belongs to (decision 57).
type ConsolePage struct {
	// Lang: the document's language, "zh-Hant".
	Lang string
	// Head: what the page's head adds after the kit's stylesheet (a site's own
	// utilities' stylesheet, a script).
	Head    g.Node
	Scheme  ConsoleScheme
	Content ConsolePageContent
	// Overlays: what the page shows over its column, written after it in
	// <body>, outside the <main> a page's script may swap: its <dialog>s, a
	// [ConsoleViewer].
	Overlays g.Node
}

func (p ConsolePage) Render(w io.Writer) error {
	if p.Content == nil {
		panic("console.ConsolePage needs a Content")
	}
	return document.Document{Owner: "console.ConsolePage", Lang: p.Lang, Title: p.Content.pageTitle(), Stylesheet: ciderui.ConsoleStylesheet(), Head: p.Head, Class: p.Scheme.class, BodyClass: "cui-console-page", Children: g.Group{p.Content, p.Overlays}}.Render(w)
}

// ConsolePageContent is what fills a [ConsolePage]: a [ConsoleContent], the
// column most pages are, or a [ConsoleCanvas], a page that lays out the whole
// window itself. Either carries the page's one Title.
type ConsolePageContent interface {
	g.Node
	pageTitle() string
}

func (c ConsoleContent) pageTitle() string {
	need.Field("console.ConsolePage.Content", "a Title", c.Title)
	if c.Sidebar != nil && c.Sidebar.Wordmark != c.Title {
		return c.Title + " - " + c.Sidebar.Wordmark
	}
	return c.Title
}

// ConsoleCanvas is a page that fills the window with its own layout, its
// Children writing its <main> (muse's player and its picture review: a strip
// and a film edge to edge, a tool bar in place of the column's header).
// Title is the page's <title> alone: it has no header to show it.
type ConsoleCanvas struct {
	Title    string
	Children g.Node
}

func (c ConsoleCanvas) pageTitle() string {
	need.Field("console.ConsolePage.Content", "a Title", c.Title)
	return c.Title
}

func (c ConsoleCanvas) Render(w io.Writer) error {
	return g.Group{c.Children}.Render(w)
}
