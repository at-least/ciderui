//go:build browser

package consoleprose

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const p = ".cui-console-prose"

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name:   "console prose",
			HTML:   h.Example("prose"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: p + " > p + p", Props: h.Props(h.Size, []string{"color", "margin-top"}), Golden: "the console's body (cui-console-body, 16px on a phone as the page's), blocks cui-prose's 14px on 25px apart in the console's line"},
				{Target: p + " > p:first-child", Props: []string{"margin-top"}, Golden: "the first block flush with the container: the room above is the section's"},
				{Target: p + " > h2", Props: h.Props(h.Size, []string{"margin-top"}), Golden: "one rung under the section's title (cui-console-h5-emph), 20px under the text as the section's blocks"},
				{Target: p + " > h3", Props: h.Props(h.Size, []string{"margin-top"}), Golden: "the body's size at 600"},
				{Target: p + " > ul", Props: []string{"margin-left", "padding-left", "list-style-type", "list-style-position"}, Golden: "a list 20px in, its marks outside, as cui-prose's"},
				{Target: p + " > ol", Props: []string{"list-style-type"}, Golden: "numbered"},
				{Target: p + " > ul > li + li", Props: []string{"margin-top"}, Golden: "items 0.4em apart, as cui-prose's"},
				{Target: p + " b", Props: []string{"font-weight"}, Golden: "bold at 600, the console's emphasis"},
				{Target: p + " a", Props: []string{"color", "text-decoration-line"}, Golden: "the console link's blue, underlined in running text (WCAG 1.4.1), as the banner's links"},
				{Target: p + " p code", Props: []string{"font-family"}, Golden: "the console's mono face at the text's size"},
				{Target: p + " > table", Props: []string{"display", "overflow-x"}, Golden: "markdown writes no wrapper: the table scrolls itself"},
				{Target: p + " th", Props: h.Props(h.Size, []string{"color", "padding-top", "padding-left", "text-align"}), Golden: "cui-console-table's column label, flush at the table's start"},
				{Target: p + " td:last-child", Props: []string{"padding-top", "padding-left", "padding-right", "border-bottom-width", "border-bottom-color", "vertical-align"}, Golden: "cui-console-table's cell over its divider, flush at the table's end, its lines at the top"},
			},
		},
		{
			Name:   "console prose code, quote and rule",
			HTML:   h.Example("code"),
			Widths: []int{1200, 390},
			Checks: []h.Check{
				{Target: p + " > pre", Props: h.Props(h.Size, []string{"border-top-left-radius", "padding-top", "padding-left", "background-color", "overflow-x"}), Golden: "cui-console-pre's listing: a long line scrolls inside the box"},
				{Target: p + " > pre code", Props: []string{"font-size"}, Golden: "the listing's 13px, not the code's own"},
				{Target: p + " > blockquote", Props: []string{"border-top-left-radius", "padding-top", "padding-left", "background-color", "border-top-width"}, Golden: "the listing's box with no edge, the text's size"},
				{Target: p + " > blockquote > p + p", Props: []string{"margin-top"}, Golden: "a quote's paragraphs as far apart as the text's"},
				{Target: p + " > hr", Props: []string{"border-top-width", "border-top-style", "border-top-color"}, Golden: "the console's divider"},
			},
		},
		{
			Name:   "console prose long word",
			HTML:   `<div class="cui-console-page" style="width: 200px"><div class="cui-console-prose"><ul><li>claude: at com.example.store.backup.BackupScheduler$RetryingUploader.uploadSnapshotToColdStorage(BackupScheduler.kt:412)</li></ul></div></div>`,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: p + " li", Props: []string{"overflow-wrap", "box-width"}, Golden: "a word with no space breaks at the column's edge, so the page never scrolls sideways"},
			},
		},
		{
			Name:   "console prose in a card",
			HTML:   `<div class="cui-console-page"><div class="cui-console-card"><div class="cui-console-prose"><pre><code>x</code></pre></div></div></div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: p + " > pre", Props: []string{"background-color"}, Golden: "in a card, whose gray is its own, the page's white, as cui-console-pre"},
			},
		},
		{
			Name:   "console prose with a part",
			HTML:   `<div class="cui-console-page"><div class="cui-console-prose"><p><a class="cui-console-link" href="#">Open</a></p></div></div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: p + " .cui-console-link", Props: []string{"text-decoration-line"}, Golden: "a part's link keeps its own look: the text's link rule leaves it alone"},
			},
		},
		{
			Name:   "console prose dark",
			HTML:   `<div class="cui-dark">` + h.Example("prose") + `</div>`,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: p + " > p + p", Props: []string{"color"}, Golden: "the dark side derived (tokens.css, decision 18)"},
				{Target: p + " td:last-child", Props: []string{"border-bottom-color"}, Golden: "the dark divider"},
			},
		},
	})
}
