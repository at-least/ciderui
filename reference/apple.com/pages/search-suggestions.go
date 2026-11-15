package pages

import (
	"fmt"
	"slices"
	"strings"

	"github.com/at-least/ciderui/internal/pagespec"
)

// @external: The global nav's search showing its suggestions, opened as search.go opens
// it, with nothing typed and no suggestion asked for: the field given a value
// and its buttons as search-text.go gives them, the search marked as
// apple.com's script marks it with suggestions (globalnav-search-with-results,
// added with no event: `set`'s classList), the quick links hidden, and in
// their place a copy the probe adds of what its script writes then (its
// _createAMLSearchTemplate, read in globalheader.umd.js on 2026-10-21): two
// sections, Suggested Links and Suggested Searches, each a heading over rows,
// each row a link with its icon (empty SVGs of apple.com's images' sizes,
// 9 by 16 and 13 by 25) and its text, the part the service marks in a
// span (-text-highlight). The copy is outside the results' current box, so
// none of its script's motion runs on it. Read only.
func searchSuggestions() pagespec.Spec {
	SPACE := []string{"padding-top", "padding-right", "padding-bottom", "padding-left", "margin-top", "margin-right", "margin-bottom", "margin-left"}
	FLY := "#globalnav-submenu-search"
	WIDE := []int{1600, 1200, 900, 834}
	PHONE := []int{833, 390}
	INPUT := FLY + " .globalnav-searchfield-input"
	SET := []pagespec.Set{pagespec.Attr(INPUT, "value", "iph")}
	for _, b := range []string{FLY + " .globalnav-searchfield-reset", FLY + " .globalnav-searchfield-submit"} {
		SET = append(SET, pagespec.Attr(b, "tabindex", "0"), pagespec.NoAttr(b, "disabled"), pagespec.NoAttr(b, "aria-hidden"))
	}
	SET = append(SET,
		pagespec.Attr("#globalnav .globalnav-search", "classList", "globalnav-search-with-results"),
		pagespec.Attr(FLY+" .globalnav-searchresults-current", "hidden", ""),
	)
	images := func(w, h, W, H int) string {
		return fmt.Sprintf(`<span class="globalnav-image-regular globalnav-link-image"><svg width="%d" height="%d" viewBox="0 0 %d %d"></svg></span><span class="globalnav-image-compact globalnav-link-image"><svg width="%d" height="%d" viewBox="0 0 %d %d"></svg></span>`, w, h, w, h, W, H, W, H)
	}
	row := func(href, term, rest string) string {
		return `<li class="globalnav-searchresults-list-item" role="listitem"><a class="globalnav-searchresults-list-link" href="` + href + `"><span class="globalnav-searchresults-list-icon">` + images(9, 16, 13, 25) + `</span><span class="globalnav-searchresults-list-text"><span class="globalnav-searchresults-list-text-highlight">` + term + `</span>` + rest + `</span></a></li>`
	}
	section := func(title string, rows ...string) string {
		return `<div class="globalnav-searchresults-container"><h2 class="globalnav-searchresults-header">` + title + `</h2><ul class="globalnav-searchresults-list" role="list">` + strings.Join(rows, "") + `</ul></div>`
	}
	COPY := `<div>` + section("Suggested Links", row("#iphone-18-pro", "iPh", "one 18 Pro"), row("#iphone-air", "iPh", "one Air")) + section("Suggested Searches", row("#search-iphone", "iph", "one"), row("#search-iphone-case", "iph", "one case")) + `</div>`
	INTO := FLY + " .globalnav-searchresults"
	LINKS := ".globalnav-searchresults-container:first-child"
	SEARCHES := ".globalnav-searchresults-container:last-child"
	s := search()
	s.Probes = append(pick(s.Probes, "nav", "nav.phone"), []pagespec.Probe{
		{Name: "search", Set: SET, Select: "#globalnav .globalnav-search", Props: []string{"display"}},
		{Name: "current", Select: FLY + " .globalnav-searchresults-current", Props: []string{"display"}},
		{Name: "field", Select: FLY + " .globalnav-searchfield", Props: slices.Concat(SPACE, []string{"height"}), Box: true, Position: true},
		{Name: "results", Select: FLY + " .globalnav-searchresults", Props: SPACE, Box: true, Position: true},
		{Name: "links", Inject: COPY, Into: INTO, Select: LINKS, Props: SPACE, Box: true, Position: true},
		{Name: "searches", Inject: COPY, Into: INTO, Select: SEARCHES, Props: SPACE, Box: true, Position: true},
		{Name: "links.header", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-header", Props: slices.Concat(TYPE, SPACE, []string{"opacity"}), Box: true, Position: true},
		{Name: "searches.header", Inject: COPY, Into: INTO, Select: SEARCHES + " .globalnav-searchresults-header", Props: slices.Concat(TYPE, SPACE, []string{"opacity"}), Box: true, Position: true},
		{Name: "links.list", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list", Props: slices.Concat(SPACE, []string{"list-style-type"}), Box: true, Position: true},
		{Name: "links.item", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list-item:first-child", Props: slices.Concat(SPACE, []string{"display", "opacity"}), Box: true, Position: true},
		{Name: "links.item2", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list-item:last-child", Props: SPACE, Box: true, Position: true},
		{Name: "searches.item", Inject: COPY, Into: INTO, Select: SEARCHES + " .globalnav-searchresults-list-item:first-child", Props: SPACE, Box: true, Position: true},
		{Name: "searches.item2", Inject: COPY, Into: INTO, Select: SEARCHES + " .globalnav-searchresults-list-item:last-child", Props: SPACE, Box: true, Position: true},
		{Name: "link", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list-link", Props: slices.Concat(TYPE, SPACE, []string{"display", "align-items", "text-decoration-line", "background-color", "border-top-left-radius"}, MOTION), Box: true, Position: true},
		{Name: "text", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list-text", Props: TYPE, Box: true, Position: true},
		{Name: "highlight", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list-text-highlight", Props: slices.Concat(TYPE, []string{"background-color", "display"}, MOTION), Box: true, Position: true},
		{Name: "links.icon", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list-icon", Props: slices.Concat(SPACE, []string{"display", "fill"}), Box: true, Position: true},
		{Name: "searches.icon", Inject: COPY, Into: INTO, Select: SEARCHES + " .globalnav-searchresults-list-icon", Props: slices.Concat(SPACE, []string{"display", "fill"}), Box: true, Position: true},
		{Name: "searches.image", Inject: COPY, Into: INTO, Select: SEARCHES + " .globalnav-image-regular svg", Props: []string{"width", "height", "fill", "display"}, Box: true, Position: true, Widths: WIDE},
		{Name: "searches.image-compact", Inject: COPY, Into: INTO, Select: SEARCHES + " .globalnav-image-compact svg", Props: []string{"width", "height", "fill", "display"}, Box: true, Position: true, Widths: PHONE},
		// @external: hovered and focused
		{Name: "link.states", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list-link", Props: slices.Concat([]string{"color", "background-color"}, FOCUS), States: []string{"hover", "focus-visible"}},
		{Name: "highlight.states", Inject: COPY, Into: INTO, Select: LINKS + " .globalnav-searchresults-list-text-highlight", StateOn: LINKS + " .globalnav-searchresults-list-link", Props: []string{"color"}, States: []string{"hover", "focus-visible"}},
		{Name: "searches.icon-states", Inject: COPY, Into: INTO, Select: SEARCHES + " .globalnav-searchresults-list-icon", StateOn: SEARCHES + " .globalnav-searchresults-list-link", Props: []string{"fill"}, States: []string{"hover", "focus-visible"}},
	}...)
	return s
}
