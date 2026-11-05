package site

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Date is the date an [Article], [Post] or [Story] shows; nil for none.
type Date struct {
	// Datetime: machine-readable (2026-10-19).
	Datetime string
	Text     string
}

func (d *Date) node(class string) g.Node {
	if d == nil {
		return nil
	}
	if d.Datetime == "" {
		panic("site.Date needs a Datetime")
	}
	if d.Text == "" {
		panic("site.Date needs a Text")
	}
	return h.Time(h.Class(class), h.DateTime(d.Datetime), g.Text(d.Text))
}
