package pages

import (
	"slices"

	"github.com/at-least/ciderui/internal/signedin"
)

var list = func() signedin.Spec {
	card := `[class*="ListItemWrapper"] a[href^="/news/"] [class^="Card-sc-"]`
	sides := []string{"border-left-width", "border-left-style", "border-left-color", "border-right-width", "border-right-style", "border-right-color"}
	return signedin.Spec{
		Path: "/",
		Probes: []signedin.Probe{
			{Name: "page", Select: "body", Props: signedin.Page},
			{Name: "page.ground", Select: `div:has(> [class*="NavbarStyles__NavbarWrapper"])`, Props: []string{"background-color"}, Box: true},
			{Name: "bar", Select: "nav.localnav", Props: []string{"height", "position", "top", "z-index"}, Box: true},
			{Name: "bar.ground", Select: "nav.localnav .localnav-background", Props: []string{"background-color", "backdrop-filter", "-webkit-backdrop-filter", "border-bottom-width", "border-bottom-color"}, Box: true},
			{Name: "bar.layer", Select: "nav.localnav .localnav-background", Pseudo: "::after", Props: []string{"content", "position", "top", "bottom", "height", "background-color", "backdrop-filter", "-webkit-backdrop-filter"}},
			{Name: "bar.title", Select: `nav.localnav [class*="NavbarStyles__BrandText"]`, Props: signedin.Type, Box: true},
			{Name: "bar.compose", Select: `[class*="NavbarStyles__ComposeIcon"]`, Props: slices.Concat(signedin.Type, []string{"width", "height"}), Box: true},
			{Name: "bar.help", Select: `a[class*="NavbarStyles__HelpIconWrapper"]`, Props: slices.Concat(signedin.Type, []string{"width", "height"}), Box: true},
			{Name: "bar.menu", Select: `[class*="NavbarStyles__UserMenu"]`, Props: slices.Concat(signedin.Type, []string{"cursor"}), Box: true},
			{Name: "pane", Select: `aside[class^="LeftPane-sc-"]`, Props: slices.Concat([]string{"width", "background-color"}, sides), Box: true},
			{Name: "search", Select: `[class*="ItemSearchField__SearchWrapper"]`, Props: slices.Concat(signedin.Box, []string{"height"}), Box: true},
			{Name: "search.field", Select: `[class*="ItemSearchField__SearchWrapper"] input`, Props: slices.Concat(signedin.Type, signedin.Box, []string{"height"}), Box: true},
			{Name: "search.placeholder", Select: `[class*="ItemSearchField__SearchWrapper"] input`, Pseudo: "::placeholder", Props: signedin.Type},
			{Name: "sort", Select: `[class*="ListContainerStyles__ItemSortChooser"]`, Props: slices.Concat(signedin.Box, []string{"height"}), Box: true},
			{Name: "sort.label", Select: `[class*="ListContainerStyles__ItemSortChooser"] .itemSorter`, Props: slices.Concat(signedin.Type, []string{"cursor"}), Box: true},
			{Name: "sort.chevron", Select: `[class*="ListContainerStyles__ItemSortChooser"] .icon-chevrondown`, Props: slices.Concat(signedin.Type, []string{"width", "height"}), Box: true},
			{Name: "filter", Select: `[class*="ListContainerStyles__ItemFilter"] button.filterIcon`, Props: slices.Concat(signedin.Type, signedin.Box, []string{"width", "height", "cursor"}), Box: true},
			{Name: "filter.icon", Select: `[class*="ListContainerStyles__ItemFilter"] button.filterIcon svg`, Props: []string{"width", "height", "color", "fill"}, Box: true},
			{Name: "item", Select: card, Props: slices.Concat(signedin.Box, sides, []string{"height", "display"}), Box: true},
			{Name: "item.mark", Select: card + " svg.unread", Props: []string{"width", "height", "margin-top", "margin-left"}, Box: true},
			{Name: "item.mark.fill", Select: card + " svg.unread path", Props: []string{"fill"}},
			{Name: "item.title", Select: card + ` [class^="Title-sc-"]`, Props: slices.Concat(signedin.Type, []string{"white-space", "text-overflow", "overflow"}), Box: true},
			{Name: "item.kind", Select: card + ` [class^="Type-sc-"]`, Props: signedin.Type, Box: true},
			{Name: "item.line", Select: card + ` [class^="Subtext-sc-"]`, Props: signedin.Type, Box: true},
			{Name: "item.date", Select: card + ` [class^="Stamp-sc-"]`, Props: slices.Concat(signedin.Type, []string{"text-align"}), Box: true},
			{Name: "empty", Select: `article[class^="RightPane-sc-"] [class*="routes__Center"]`, Props: slices.Concat(signedin.Type, []string{"display", "align-items", "justify-content", "text-align"}), Box: true},
		},
	}
}()
