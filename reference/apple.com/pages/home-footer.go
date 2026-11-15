package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/pagespec"
)

func homeFooter() pagespec.Spec {
	SECTION := "#ac-globalfooter .ac-gf-directory-column-section"
	return pagespec.Spec{
		URL:    "https://www.apple.com/",
		Widths: []int{390},
		Probes: []pagespec.Probe{
			{Name: "section.open", Click: SECTION + " .ac-gf-directory-column-section-title-button", Select: SECTION, Props: []string{"padding-bottom"}, Box: true},
			{Name: "list.open", Select: SECTION + " .ac-gf-directory-column-section-list", Props: []string{"visibility", "padding-top", "padding-bottom", "transition-property", "transition-duration", "transition-timing-function"}, Box: true},
			{Name: "item.open", Select: SECTION + " .ac-gf-directory-column-section-item", Props: []string{"margin-bottom", "padding-top"}, Box: true},
			{Name: "link.open", Select: SECTION + " .ac-gf-directory-column-section-link", Props: slices.Concat(TYPE, LINK, []string{"display", "padding-top", "padding-right", "padding-bottom", "padding-left"}), Box: true},
			{Name: "icon.open", Select: SECTION + " .ac-gf-directory-column-section-title-icon", Props: []string{"color", "transform"}},
		},
	}
}
