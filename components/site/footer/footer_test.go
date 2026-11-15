//go:build browser

package footer

import (
	"testing"

	h "github.com/at-least/ciderui/internal/harness"
)

const section = ".cui-footer-section:first-child"

var (
	wide  = []int{1600, 1200, 900}
	phone = []int{390}
)

func TestCases(t *testing.T) {
	h.Run(t, []h.Case{
		{
			Name: "footer",
			HTML: h.Example("footer"),
			Checks: []h.Check{
				{Target: ".cui-footer", Props: h.Props(h.Size, []string{"color", "background-color"}), Apple: h.AppleProbe("home", "footer")},
				{Target: ".cui-footer-content", Props: []string{"max-width", "padding-left", "padding-right"}, Apple: h.AppleProbe("home", "footer.content")},
				{Target: ".cui-footer-legal", Props: []string{"padding-top", "padding-bottom"}, Apple: h.AppleProbe("home", "footer.footer")},
				{Target: ".cui-footer-copyright", Props: h.Props(h.Size, []string{"color", "margin-top"}), Apple: h.AppleProbe("home", "footer.copyright")},
				{Target: ".cui-footer-copyright a", Props: []string{"color", "text-decoration-line"}, Apple: h.AppleProbe("home", "footer.legal-link")},
				{Target: ".cui-footer", Props: []string{"box-top"}, Golden: "the body's last child at the window's foot on a page shorter than the window, not under the last line of a short page"},
				{Target: ".cui-footer-copyright a", State: []string{"hover"}, Props: []string{"text-decoration-line"}, Golden: "a link in the copyright underlined under the pointer, as apple.com's footer links"},
			},
		},
		{
			Name:   "touch target",
			HTML:   h.Example("footer"),
			Touch:  true,
			Widths: []int{390},
			Checks: []h.Check{
				{Target: ".cui-footer-copyright a", Pseudo: "::after", Props: []string{"height"}, Golden: "the company link's target 44pt on a touch screen, the HIG's default, past its 16px line"},
			},
		},
		{
			Name:   "dark",
			HTML:   `<div class="cui-dark">` + h.Example("directory") + `</div>`,
			Widths: wide,
			Checks: []h.Check{
				{Target: ".cui-footer", Props: []string{"color", "background-color"}, Apple: h.AppleProbe("product", "footer")},
				{Target: section + " .cui-footer-title", Props: []string{"color"}, Apple: h.AppleProbe("product", "footer.title")},
				{Target: section + " .cui-footer-list a", Props: []string{"color"}, Apple: h.AppleProbe("product", "footer.link")},
			},
		},
		{
			Name: "directory",
			HTML: h.Example("directory"),
			Checks: []h.Check{
				{Target: ".cui-footer-directory", Props: []string{"display", "flex-direction", "flex-wrap", "padding-top"}, Apple: h.AppleProbe("home", "footer.directory")},
			},
		},
		{
			// @trap: apple.com's column holds sections; here a section is the column, read against its column
			Name:   "directory-wide",
			HTML:   h.Example("directory"),
			Widths: wide,
			Checks: []h.Check{
				{Target: section, Props: []string{"flex-grow", "flex-shrink", "flex-basis"}, Apple: h.AppleProbe("home", "footer.column")},
				{Target: section + " .cui-footer-title", Props: h.Props(h.Size, []string{"color"}), Apple: h.AppleProbe("home", "footer.title")},
				{Target: section + " .cui-footer-title", Props: []string{"margin-right", "margin-bottom"}, Apple: h.AppleProbe("home", "footer.title-text")},
				{Target: section + " .cui-footer-list", Props: []string{"margin-right", "padding-top", "padding-bottom"}, Apple: h.AppleProbe("home", "footer.list")},
				{Target: section + " .cui-footer-list li:first-child", Props: []string{"margin-bottom"}, Apple: h.AppleProbe("home", "footer.item")},
				{Target: section + " .cui-footer-list li:last-child", Props: []string{"margin-bottom"}, Apple: h.AppleProbe("home", "footer.item-last")},
				{Target: section + " .cui-footer-list a", Props: h.Props(h.Size, []string{"color", "text-decoration-line"}), Apple: h.AppleProbe("home", "footer.link")},
				{Target: section + " .cui-footer-list a", State: []string{"hover"}, Props: []string{"color", "text-decoration-line"}, Apple: h.AppleProbe("home", "footer.link", "hover")},
				{Target: section + " .cui-footer-toggle", Props: []string{"display"}, Golden: "its toggle gone from 834px: every section shown"},
			},
		},
		{
			Name:   "directory-phone",
			HTML:   h.Example("directory"),
			Widths: phone,
			Checks: []h.Check{
				{Target: section, Props: []string{"padding-left", "padding-right", "padding-bottom", "margin-left", "position", "box-height"}, Apple: h.AppleProbe("home", "footer.section-phone")},
				{Target: section, Pseudo: "::after", Props: []string{"content", "position", "bottom", "left", "right", "height", "background-color"}, Apple: h.AppleProbe("home", "footer.section-rule")},
				{Target: section + " .cui-footer-title label", Props: h.Props(h.Size, []string{"color", "padding-top", "padding-bottom", "cursor", "box-height"}), Apple: h.AppleProbe("home", "footer.title-button")},
				{Target: section + " .cui-footer-title svg", Props: []string{"color", "margin-top", "box-height"}, Apple: h.AppleProbe("home", "footer.title-icon")},
				{Target: section + " .cui-footer-title svg", Props: []string{"transition-property", "transition-duration", "transition-timing-function"}, Golden: "its turn the kit's one chevron turn, 0.32s cubic-bezier(.4,0,.6,1), the accordion's, where apple.com's reads 0.3s ease here"},
				{Target: section + " .cui-footer-list", Props: []string{"display"}, Golden: "closed, its links out of sight and out of the focus order (apple.com's hidden: home:footer.list-closed)"},
			},
		},
		{
			Name:   "directory-open",
			HTML:   h.Example("directory"),
			Widths: phone,
			Click:  []string{section + " .cui-footer-title label"},
			Checks: []h.Check{
				{Target: section + " .cui-footer-list", Props: []string{"padding-top", "padding-bottom"}, Apple: h.AppleProbe("home-footer", "list.open")},
				{Target: section + " .cui-footer-list li:first-child", Props: []string{"box-height"}, Apple: h.AppleProbe("home-footer", "item.open")},
				{Target: section + " .cui-footer-list a", Props: h.Props(h.Size, []string{"color", "display", "padding-top", "padding-right", "padding-bottom", "padding-left"}), Apple: h.AppleProbe("home-footer", "link.open")},
				{Target: section + " .cui-footer-title svg", Props: []string{"color"}, Apple: h.AppleProbe("home-footer", "icon.open")},
				{Target: section + " .cui-footer-title svg", Props: []string{"transform"}, Golden: "opened, its chevron turned up: apple.com's line moved to the same points"},
			},
		},
		{
			Name:   "directory step at 834px",
			HTML:   h.Example("directory"),
			Widths: []int{834, 833},
			Checks: []h.Check{
				{Target: ".cui-footer-content", Props: []string{"padding-left", "padding-right"}, Golden: "22px in from 834px, 16px to 833px: apple.com's footer steps there, with the global nav (reference/apple.com/MEASURED.md, not read at these widths)"},
				{Target: ".cui-footer-directory", Props: []string{"flex-direction"}, Golden: "a row of columns from 834px, a column of sections to 833px, as apple.com's"},
				{Target: section + " .cui-footer-toggle", Props: []string{"display"}, Golden: "its toggle from 833px down alone"},
			},
		},
		{
			Name:       "directory open at once",
			HTML:       h.Example("directory"),
			Widths:     phone,
			Click:      []string{section + " .cui-footer-title label"},
			FirstFrame: true,
			Checks: []h.Check{
				{Target: section + " .cui-footer-list", Props: []string{"display", "box-height"}, Golden: "opened at once, its whole list at the click: apple.com's moves in 0.3s"},
			},
		},
		{
			Name:   "directory toggle focus",
			HTML:   h.Example("directory"),
			Widths: phone,
			Press:  []string{"Tab"},
			Checks: []h.Check{
				{Target: section + " .cui-footer-toggle", Props: []string{"focused"}, Golden: "the toggle, out of sight, the page's first stop"},
				{Target: section + " .cui-footer-title label", Props: h.Outline, Golden: "its focus ring on its title"},
			},
		},
		{
			// @trap: from 834px: below it a closed section's links have no box
			Name:   "directory touch targets",
			HTML:   h.Example("directory"),
			Touch:  true,
			Widths: []int{1200},
			Checks: []h.Check{
				{Target: section + " .cui-footer-title label", Props: []string{"box-height"}, Golden: "a section's title 44px tall on a touch screen"},
				{Target: section + " .cui-footer-list li:first-child a", Props: []string{"box-height"}, Golden: "a link 44px tall on a touch screen, the HIG's 44pt (Accessibility, Mobility), a tablet's columns too"},
				{Target: section + " .cui-footer-list li:first-child", Props: []string{"margin-bottom"}, Golden: "the list's gaps inside the links' 44px"},
			},
		},
	})
}
