//go:build browser

package prose

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
	apple "github.com/at-least/ciderui/reference/apple.com/results"
)

const p = ".cui-prose"

var flow = []string{"margin-top", "margin-bottom"}

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "prose",
			HTML:   h.Example("prose"),
			Widths: h.With1600,
			Checks: []h.Check{
				{Target: p + " > p + p", Props: h.Props(h.Size, []string{"color"}, flow), Apple: h.AppleProbe(apple.SupportProse, "p")},
				{Target: p + " > h2", Props: h.Props(h.Size, []string{"color"}, flow), Apple: h.AppleProbe(apple.SupportProse, "h2")},
				{Target: p + " > h2 + p", Props: flow, Apple: h.AppleProbe(apple.SupportProse, "h2.next")},
				{Target: p + " > h3", Props: h.Props(h.Size, []string{"color"}, flow), Apple: h.AppleProbe(apple.SupportProse, "h3")},
				{Target: p + " > ul", Props: h.Props(flow, []string{"padding-left", "margin-left", "list-style-type", "list-style-position"}), Apple: h.AppleProbe(apple.SupportProse, "ul")},
				{Target: p + " > ol", Props: h.Props(flow, []string{"padding-left", "margin-left", "list-style-type", "list-style-position"}), Apple: h.AppleProbe(apple.SupportProse, "ol")},
				{Target: p + " > ul > li", Props: h.Props(h.Size, flow, []string{"padding-left"}), Apple: h.AppleProbe(apple.SupportProse, "li")},
				{Target: p + " > ul > li + li", Props: flow, Apple: h.AppleProbe(apple.SupportProse, "li.next")},
				{Target: p + " > ol + h3", Props: []string{"margin-top"}, Apple: h.AppleProbe(apple.SupportProse, "h3")},
				{Target: p + " > ul + p", Props: flow, Apple: h.AppleProbe(apple.SupportProse, "list.next")},
				{Target: p + " b", Props: h.Size, Apple: h.AppleProbe(apple.SupportProse, "b")},
				{Target: p + " a", Props: h.Props(h.Size, []string{"color", "text-decoration-line"}), Apple: h.AppleProbe(apple.SupportProse, "link")},
				{Target: p + " > p:first-child", Props: []string{"margin-top"}, Golden: "the first block flush with the container: the space above the text is the page's"},
				{Target: p + " > table", Props: []string{"margin-top", "overflow-x", "border-collapse"}, Golden: "a table 14px under the text before it, as Apple Support's wrapper (support:table.wrapper); it scrolls itself, markdown writing no wrapper"},
				{Target: p + " th", Props: h.Props(h.Size, []string{"padding-top", "padding-left", "text-align", "vertical-align", "border-bottom-width", "border-bottom-style"}), Apple: h.AppleProbe(apple.Support, "table.head")},
				{Target: p + " td", Props: h.Props(h.Size, []string{"padding-top", "padding-left", "text-align", "vertical-align", "border-bottom-width", "border-bottom-style"}), Apple: h.AppleProbe(apple.Support, "table.cell")},
				{Target: p + " td", Props: []string{"border-bottom-color"}, Golden: "the kit's divider, one rule color: Apple Support's #ddd is a gray of its own"},
			},
		},
		{
			Name:   "prose code, quote and notes",
			HTML:   h.Example("code"),
			Widths: h.With1600,
			Checks: []h.Check{
				{Target: p + " p code", Props: h.Size, Apple: h.AppleProbe(apple.Docc, "code")},
				{Target: p + " p code", Props: []string{"background-color", "padding-left"}, Apple: h.AppleProbe(apple.Docc, "code")},
				{Target: p + " > pre", Props: []string{"margin-top", "border-top-left-radius", "background-color"}, Apple: h.AppleProbe(apple.Docc, "listing")},
				{Target: p + " > pre", Props: []string{"padding-top", "padding-right", "padding-bottom", "padding-left"}, Apple: h.AppleProbe(apple.Docc, "listing.pre")},
				{Target: p + " > pre code", Props: h.Size, Apple: h.AppleProbe(apple.Docc, "listing.code")},
				{Target: p + " > pre", Props: []string{"overflow-x", "tab-size"}, Golden: "a long line scrolls inside the box, tabs four wide: DocC's listing (docc:listing.scroller)"},
				{Target: p + " > blockquote", Props: []string{"margin-top", "border-top-left-radius", "background-color", "padding-top", "padding-left"}, Apple: h.AppleProbe(apple.Docc, "note")},
				{Target: p + " > blockquote", Props: []string{"border-top-width"}, Golden: "no edge: DocC's note edges its box in #696969, a gray of its own; the alternate ground alone sets it apart"},
				{Target: p + " > blockquote > p + p", Props: []string{"margin-top"}, Golden: "a quote's paragraphs 14px apart, as the text's"},
				{Target: p + " > hr", Props: []string{"margin-top", "border-top-width", "border-top-style", "border-top-color"}, Golden: "a rule in the kit's divider, 27.2px under the text as a listing is: no Apple article measured draws one"},
				{Target: p + " > .footnotes li", Props: h.Props(h.Size, []string{"color"}), Apple: h.AppleProbe(apple.NewsroomArticle, "footnotes.li")},
				{Target: p + " > .footnotes ol", Props: []string{"padding-left", "margin-left", "list-style-type"}, Apple: h.AppleProbe(apple.NewsroomArticle, "footnotes.ol")},
				{Target: p + " > p + .footnotes", Props: []string{"margin-top"}, Apple: h.AppleProbe(apple.NewsroomArticle, "footnotes")},
				{Target: p + " > .footnotes li + li", Props: []string{"margin-top"}, Apple: h.AppleProbe(apple.NewsroomArticle, "footnotes.li")},
			},
		},
		{
			Name:   "prose with a part",
			HTML:   `<div class="cui-prose"><p>Ready?</p><p><a class="cui-btn" href="#">Buy</a></p></div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: ".cui-prose .cui-btn", Props: []string{"color", "text-decoration-line"}, Golden: "a button in the text keeps its own white and no underline: the text's link rule leaves a part's link alone"},
			},
		},
		{
			Name:   "prose picture",
			HTML:   `<div class="cui-prose" style="width: 300px"><p>Step one.</p><img src="data:image/svg+xml,%3Csvg xmlns=%22http://www.w3.org/2000/svg%22 width=%22640%22 height=%22320%22/%3E" alt=""><p>Step two.</p></div>`,
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: ".cui-prose > img", Props: []string{"margin-top", "margin-bottom", "display"}, Apple: h.AppleProbe(apple.SupportProse, "img")},
				{Target: ".cui-prose > img", Props: []string{"box-width", "box-height"}, Golden: "a picture no wider than the text, its height kept: Apple Support's are 320px screenshots (min(100%, 320px)), a consumer's any size"},
			},
		},
	})
}
