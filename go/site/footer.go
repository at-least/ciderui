package site

import (
	"io"

	"github.com/at-least/ciderui/go/icon"
	"github.com/at-least/ciderui/go/internal/class"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type FooterSection struct {
	ID    string
	Title string
	Links []Link
}

type Footer struct {
	// Notes: <p>s and <ol>s.
	Notes     g.Node
	Directory [][]FooterSection
	Shop      g.Node
	// Copyright: text, or a link to the company.
	Copyright g.Node
	Links     []Link
	Locale    *Link
	Class     string
}

func (f Footer) Render(w io.Writer) error {
	if f.Copyright == nil {
		panic("site.Footer needs a Copyright")
	}
	for _, column := range f.Directory {
		for _, section := range column {
			if section.ID == "" {
				panic("site.FooterSection needs an ID")
			}
			if section.Title == "" {
				panic("site.FooterSection needs a Title")
			}
		}
	}
	var links g.Group
	for n, link := range f.Links {
		// a space between them, as apple.com's: its inline items' gap is the
		// space's too
		if n > 0 {
			links = append(links, g.Text(" "))
		}
		links = append(links, h.Li(h.A(h.Href(link.Href), g.Text(link.Text))))
	}
	return h.Footer(class.Root("cui-footer", f.Class),
		h.Div(h.Class("cui-footer-content"),
			g.If(f.Notes != nil, h.Section(h.Class("cui-footer-notes"), h.Aria("label", "Footnotes"), f.Notes)),
			g.If(len(f.Directory) > 0, h.Nav(h.Class("cui-footer-directory"), h.Aria("label", "Directory"),
				g.Map(f.Directory, func(column []FooterSection) g.Node {
					return h.Div(h.Class("cui-footer-column"), g.Map(column, func(section FooterSection) g.Node {
						return h.Div(h.Class("cui-footer-section"),
							h.Input(h.Class("cui-footer-toggle"), h.Type("checkbox"), h.ID(section.ID)),
							h.H3(h.Class("cui-footer-title"), h.Label(h.For(section.ID), g.Text(section.Title), icon.FooterChevron)),
							h.Ul(h.Class("cui-footer-list"), g.Map(section.Links, func(link Link) g.Node {
								return h.Li(h.A(h.Href(link.Href), g.Text(link.Text)))
							})),
						)
					}))
				}),
			)),
			h.Div(h.Class("cui-footer-legal"),
				g.If(f.Shop != nil, h.P(h.Class("cui-footer-shop"), f.Shop)),
				h.Div(h.Class("cui-footer-row"),
					h.P(h.Class("cui-footer-copyright"), f.Copyright),
					g.If(len(f.Links) > 0, h.Ul(h.Class("cui-footer-links"), links)),
					g.Iff(f.Locale != nil, func() g.Node {
						return h.A(h.Class("cui-footer-locale"), h.Href(f.Locale.Href), g.Text(f.Locale.Text))
					}),
				),
			),
		),
	).Render(w)
}
