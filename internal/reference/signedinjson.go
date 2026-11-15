package reference

import (
	"fmt"

	"github.com/at-least/ciderui/internal/ojson"
)

func (r *reader) digest(v ojson.Value) [32]byte {
	var d [32]byte
	a := r.array(v)
	if len(a) != len(d) {
		r.fail("a digest of %d bytes", len(a))
		return d
	}
	for i, e := range a {
		b := r.integer(e)
		if b < 0 || b > 255 {
			r.fail("a digest's byte %d", b)
		}
		d[i] = byte(b)
	}
	return d
}

func (r *reader) signedInProbe(name string, v ojson.Value) SignedInProbe {
	p := SignedInProbe{Name: name, HTML: "", Styles: nil, Box: nil}
	o := r.object(v, "!html", "!styles", "box")
	if o == nil {
		return p
	}
	p.HTML = r.text(o["html"])
	styles, ok := o["styles"].(ojson.Object)
	if !ok {
		r.fail("styles: %T, not an object", o["styles"])
		return p
	}
	p.Styles = []Style{}
	for _, m := range styles {
		p.Styles = append(p.Styles, Style{Name: m.Key, Text: r.text(m.Value), Image: nil})
	}
	if b, ok := o["box"]; ok {
		box := r.object(b, "!width", "!height")
		p.Box = &Box{Width: r.number(box["width"]), Height: r.number(box["height"]), At: nil}
	}
	return p
}

// SignedInFromJSON is the page named name from the record a signed-in site's page
// handed back (its result and its sha256, one object).
func SignedInFromJSON(name string, v ojson.Value) (SignedInPage, error) {
	r := &reader{at: name, err: nil}
	page := SignedInPage{Name: name, Path: "", Browser: "", Date: "", Window: Window{Width: 0, Height: 0, PixelRatio: 0}, Probes: nil, Tokens: nil, ProbesDigest: [32]byte{}, ReaderDigest: [32]byte{}, SHA256: [32]byte{}}
	o := r.object(v, "!path", "!window", "!browser", "!date", "probes", "tokens", "!probesDigest", "!readerDigest", "!sha256")
	if o == nil {
		return page, r.err
	}
	page.Path, page.Browser, page.Date = r.text(o["path"]), r.text(o["browser"]), r.text(o["date"])
	window := r.array(o["window"])
	if len(window) != 3 {
		return page, fmt.Errorf("%s: a window of %d values", name, len(window))
	}
	page.Window = Window{Width: r.number(window[0]), Height: r.number(window[1]), PixelRatio: r.number(window[2])}
	page.ProbesDigest, page.ReaderDigest, page.SHA256 = r.digest(o["probesDigest"]), r.digest(o["readerDigest"]), r.digest(o["sha256"])
	probes, hasProbes := o["probes"]
	tokens, hasTokens := o["tokens"]
	if hasProbes == hasTokens {
		return page, fmt.Errorf("%s: probes or tokens, one of them", name)
	}
	if hasProbes {
		list, ok := probes.(ojson.Object)
		if !ok {
			return page, fmt.Errorf("%s: probes: %T", name, probes)
		}
		page.Probes = []SignedInProbe{}
		for _, m := range list {
			r.at = name + ", " + m.Key
			page.Probes = append(page.Probes, r.signedInProbe(m.Key, m.Value))
		}
	}
	if hasTokens {
		list, ok := tokens.(ojson.Object)
		if !ok {
			return page, fmt.Errorf("%s: tokens: %T", name, tokens)
		}
		page.Tokens = []Token{}
		for _, m := range list {
			r.at = name + ", " + m.Key
			page.Tokens = append(page.Tokens, Token{Name: m.Key, Value: r.text(m.Value)})
		}
	}
	return page, r.err
}

func digestJSON(d [32]byte) ojson.Array {
	a := ojson.Array{}
	for _, b := range d {
		a = append(a, ojson.Float(float64(b)))
	}
	return a
}

// JSON is the page as the reader hands it back: its result's members in the order
// the reader writes them, then sha256.
func (p SignedInPage) JSON() ojson.Object {
	o := ojson.Object{
		{Key: "path", Value: ojson.String(p.Path)},
		{Key: "window", Value: ojson.Array{ojson.Float(p.Window.Width), ojson.Float(p.Window.Height), ojson.Float(p.Window.PixelRatio)}},
		{Key: "browser", Value: ojson.String(p.Browser)},
		{Key: "date", Value: ojson.String(p.Date)},
	}
	if p.Tokens != nil {
		tokens := ojson.Object{}
		for _, t := range p.Tokens {
			tokens = append(tokens, ojson.Member{Key: t.Name, Value: ojson.String(t.Value)})
		}
		o = append(o, ojson.Member{Key: "tokens", Value: tokens})
	} else {
		probes := ojson.Object{}
		for _, pr := range p.Probes {
			styles := ojson.Object{}
			for _, s := range pr.Styles {
				styles = append(styles, ojson.Member{Key: s.Name, Value: ojson.String(s.Text)})
			}
			probe := ojson.Object{{Key: "html", Value: ojson.String(pr.HTML)}, {Key: "styles", Value: styles}}
			if pr.Box != nil {
				probe = append(probe, ojson.Member{Key: "box", Value: ojson.Object{{Key: "width", Value: ojson.Float(pr.Box.Width)}, {Key: "height", Value: ojson.Float(pr.Box.Height)}}})
			}
			probes = append(probes, ojson.Member{Key: pr.Name, Value: probe})
		}
		o = append(o, ojson.Member{Key: "probes", Value: probes})
	}
	return append(o,
		ojson.Member{Key: "probesDigest", Value: digestJSON(p.ProbesDigest)},
		ojson.Member{Key: "readerDigest", Value: digestJSON(p.ReaderDigest)},
		ojson.Member{Key: "sha256", Value: digestJSON(p.SHA256)},
	)
}
