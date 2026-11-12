package site

import (
	"io"
	"strconv"

	"github.com/at-least/ciderui/go/internal/glyph"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// FooterSection is one column of the footer's directory: a title over its
// links, folded under its title to 833px.
type FooterSection struct {
	Title string
	Links []FooterLink
}

type FooterLink struct {
	Label string
	Href  string
}

type Footer struct {
	// Directory: nil for none.
	Directory []FooterSection
	// DirectoryAriaLabel: "" with no Directory.
	DirectoryAriaLabel string
	// Copyright: text, or a link to the company.
	Copyright g.Node
}

func (f Footer) Render(w io.Writer) error {
	if f.Copyright == nil {
		panic("site.Footer needs a Copyright")
	}
	if (len(f.Directory) > 0) != (f.DirectoryAriaLabel != "") {
		panic("site.Footer has a Directory or a DirectoryAriaLabel without the other")
	}
	return h.Footer(h.Class("cui-footer"),
		h.Div(h.Class("cui-footer-content"),
			g.If(len(f.Directory) > 0, h.Nav(h.Class("cui-footer-directory"), h.Aria("label", f.DirectoryAriaLabel),
				sections(f.Directory),
			)),
			h.Div(h.Class("cui-footer-legal"),
				h.P(h.Class("cui-footer-copyright"), f.Copyright),
			),
		),
	).Render(w)
}

func sections(directory []FooterSection) g.Group {
	var nodes g.Group
	for i, section := range directory {
		if section.Title == "" || len(section.Links) == 0 {
			panic("site.FooterSection needs a Title and Links")
		}
		// @invariant: the label's for; one footer a page, as one global nav
		id := "cui-footer-section-" + strconv.Itoa(i+1)
		nodes = append(nodes, h.Div(h.Class("cui-footer-section"),
			h.Input(h.Class("cui-footer-toggle"), h.Type("checkbox"), h.ID(id)),
			h.H3(h.Class("cui-footer-title"), h.Label(h.For(id), g.Text(section.Title), glyph.FooterChevron)),
			h.Ul(h.Class("cui-footer-list"), g.Map(section.Links, func(link FooterLink) g.Node {
				if link.Label == "" || link.Href == "" {
					panic("site.FooterSection's Links need their Label and Href")
				}
				return h.Li(h.A(h.Href(link.Href), g.Text(link.Label)))
			})),
		))
	}
	return nodes
}
