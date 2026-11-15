package reference

import (
	"fmt"
	"strconv"

	"github.com/at-least/ciderui/internal/ojson"
)

// @invariant: a record's members in the order the probe writes them; reading one
// fails on a key or a type this order does not hold, so nothing a record carries
// is dropped on its way to Go
type reader struct {
	at  string
	err error
}

func (r *reader) fail(format string, args ...any) {
	if r.err == nil {
		r.err = fmt.Errorf("%s: %s", r.at, fmt.Sprintf(format, args...))
	}
}

// @invariant: o's members, each key one of keys, the keys every listed with a ! present
func (r *reader) object(v ojson.Value, keys ...string) map[string]ojson.Value {
	o, ok := v.(ojson.Object)
	if !ok {
		r.fail("%T, not an object", v)
		return nil
	}
	allowed := map[string]bool{}
	for _, k := range keys {
		allowed[k] = true
	}
	out := map[string]ojson.Value{}
	for _, m := range o {
		if _, required := allowed["!"+m.Key]; !required && !allowed[m.Key] {
			r.fail("unknown key %q", m.Key)
			return nil
		}
		if _, twice := out[m.Key]; twice {
			r.fail("key %q twice", m.Key)
			return nil
		}
		out[m.Key] = m.Value
	}
	for _, k := range keys {
		if k[0] == '!' && out[k[1:]] == nil {
			r.fail("no %q", k[1:])
			return nil
		}
	}
	return out
}

func (r *reader) text(v ojson.Value) string {
	s, ok := v.(ojson.String)
	if !ok {
		r.fail("%T, not a string", v)
	}
	return string(s)
}

func (r *reader) number(v ojson.Value) float64 {
	n, ok := v.(ojson.Number)
	if !ok {
		r.fail("%T, not a number", v)
		return 0
	}
	f, err := strconv.ParseFloat(string(n), 64)
	if err != nil {
		r.fail("%v", err)
	}
	return f
}

func (r *reader) integer(v ojson.Value) int {
	n, ok := v.(ojson.Number)
	if !ok {
		r.fail("%T, not a number", v)
		return 0
	}
	i, err := strconv.Atoi(string(n))
	if err != nil {
		r.fail("%v", err)
	}
	return i
}

func (r *reader) array(v ojson.Value) ojson.Array {
	a, ok := v.(ojson.Array)
	if !ok {
		r.fail("%T, not an array", v)
	}
	return a
}

func (r *reader) styles(v ojson.Value) []Style {
	o, ok := v.(ojson.Object)
	if !ok {
		r.fail("styles: %T, not an object", v)
		return nil
	}
	var out []Style
	for _, m := range o {
		if s, ok := m.Value.(ojson.String); ok {
			out = append(out, Style{Name: m.Key, Text: string(s), Image: nil})
			continue
		}
		img := r.object(m.Value, "!viewBox", "!shapes")
		if img == nil {
			return nil
		}
		image := &Image{ViewBox: r.text(img["viewBox"]), Shapes: nil}
		for _, s := range r.array(img["shapes"]) {
			sh := r.object(s, "!shape", "!fill", "!stroke")
			if sh == nil {
				return nil
			}
			image.Shapes = append(image.Shapes, Shape{Shape: r.text(sh["shape"]), Fill: r.text(sh["fill"]), Stroke: r.text(sh["stroke"])})
		}
		out = append(out, Style{Name: m.Key, Text: "", Image: image})
	}
	return out
}

func (r *reader) probe(name string, v ojson.Value) Probe {
	p := Probe{Name: name, Missing: false, HTML: "", Styles: nil, States: nil, Box: nil, Keyframes: nil, Animations: nil, Matches: nil, Sections: nil}
	if v == (ojson.Null{}) {
		p.Missing = true
		return p
	}
	o := r.object(v, "!html", "!styles", "keyframes", "sections", "matches", "box", "animations", "states")
	if o == nil {
		return p
	}
	p.HTML = r.text(o["html"])
	p.Styles = r.styles(o["styles"])
	if kf, ok := o["keyframes"]; ok {
		rules, ok := kf.(ojson.Object)
		if !ok {
			r.fail("keyframes: %T", kf)
		}
		p.Keyframes = []Keyframes{}
		for _, m := range rules {
			if m.Value == (ojson.Null{}) {
				r.fail("@keyframes %s: not found", m.Key)
			}
			k := Keyframes{Name: m.Key, Frames: []Keyframe{}}
			for _, f := range r.array(m.Value) {
				frame := r.object(f, "!at", "!style")
				k.Frames = append(k.Frames, Keyframe{At: r.text(frame["at"]), Style: r.text(frame["style"])})
			}
			p.Keyframes = append(p.Keyframes, k)
		}
	}
	if s, ok := o["sections"]; ok {
		p.Sections = []Section{}
		for _, e := range r.array(s) {
			pair := r.array(e)
			if len(pair) != 2 {
				r.fail("a section of %d values", len(pair))
				continue
			}
			p.Sections = append(p.Sections, Section{Name: r.text(pair[0]), Rows: r.integer(pair[1])})
		}
	}
	if m, ok := o["matches"]; ok {
		matches, ok := m.(ojson.Object)
		if !ok {
			r.fail("matches: %T", m)
		}
		p.Matches = []Match{}
		for _, e := range matches {
			b, ok := e.Value.(ojson.Bool)
			if !ok {
				r.fail("matches %s: %T", e.Key, e.Value)
			}
			p.Matches = append(p.Matches, Match{Selector: e.Key, Matches: bool(b)})
		}
	}
	if b, ok := o["box"]; ok {
		box := r.object(b, "!width", "!height", "top", "left")
		p.Box = &Box{Width: r.number(box["width"]), Height: r.number(box["height"]), At: nil}
		_, top := box["top"]
		_, left := box["left"]
		if top != left {
			r.fail("box: a top without a left, or a left without a top")
		}
		if top {
			p.Box.At = &Position{Top: r.number(box["top"]), Left: r.number(box["left"])}
		}
	}
	if a, ok := o["animations"]; ok {
		p.Animations = []Animation{}
		for _, e := range r.array(a) {
			an := r.object(e, "!kind", "!what", "!on", "!duration", "!delay", "direction", "!easing", "!from", "!to", "seen")
			anim := Animation{Kind: r.text(an["kind"]), What: r.text(an["what"]), On: r.text(an["on"]), Duration: r.number(an["duration"]), Delay: r.number(an["delay"]), Easing: r.text(an["easing"]), From: r.text(an["from"]), To: r.text(an["to"]), Observed: nil}
			_, direction := an["direction"]
			_, seen := an["seen"]
			if direction != seen {
				r.fail("an animation with a direction or a seen, not both")
			}
			if seen {
				anim.Observed = &Observed{Direction: r.text(an["direction"]), Seen: r.integer(an["seen"])}
			}
			p.Animations = append(p.Animations, anim)
		}
	}
	if s, ok := o["states"]; ok {
		states, ok := s.(ojson.Object)
		if !ok {
			r.fail("states: %T", s)
		}
		p.States = []State{}
		for _, m := range states {
			p.States = append(p.States, State{Name: m.Key, Styles: r.styles(m.Value)})
		}
	}
	return p
}

// FromJSON is the page named name from the record the probe built.
func FromJSON(name string, v ojson.Value) (Page, error) {
	r := &reader{at: name, err: nil}
	page := Page{Name: name, URL: "", FinalURL: "", Browser: "", Date: "", Widths: nil}
	o := r.object(v, "!url", "!finalUrl", "!browser", "!date", "!widths")
	if o == nil {
		return page, r.err
	}
	page.URL, page.FinalURL, page.Browser, page.Date = r.text(o["url"]), r.text(o["finalUrl"]), r.text(o["browser"]), r.text(o["date"])
	widths, ok := o["widths"].(ojson.Object)
	if !ok {
		return page, fmt.Errorf("%s: widths: %T", name, o["widths"])
	}
	for _, w := range widths {
		width, err := strconv.Atoi(w.Key)
		if err != nil {
			return page, fmt.Errorf("%s: width %q: %w", name, w.Key, err)
		}
		probes, ok := w.Value.(ojson.Object)
		if !ok {
			return page, fmt.Errorf("%s: width %d: %T", name, width, w.Value)
		}
		at := Width{Width: width, Probes: nil}
		for _, m := range probes {
			r.at = fmt.Sprintf("%s at %dpx, %s", name, width, m.Key)
			at.Probes = append(at.Probes, r.probe(m.Key, m.Value))
		}
		page.Widths = append(page.Widths, at)
	}
	return page, r.err
}

func styles(list []Style) ojson.Object {
	out := ojson.Object{}
	for _, s := range list {
		if s.Image == nil {
			out = append(out, ojson.Member{Key: s.Name, Value: ojson.String(s.Text)})
			continue
		}
		shapes := ojson.Array{}
		for _, sh := range s.Image.Shapes {
			shapes = append(shapes, ojson.Object{{Key: "shape", Value: ojson.String(sh.Shape)}, {Key: "fill", Value: ojson.String(sh.Fill)}, {Key: "stroke", Value: ojson.String(sh.Stroke)}})
		}
		out = append(out, ojson.Member{Key: s.Name, Value: ojson.Object{{Key: "viewBox", Value: ojson.String(s.Image.ViewBox)}, {Key: "shapes", Value: shapes}}})
	}
	return out
}

func (p Probe) json() ojson.Value {
	if p.Missing {
		return ojson.Null{}
	}
	o := ojson.Object{{Key: "html", Value: ojson.String(p.HTML)}, {Key: "styles", Value: styles(p.Styles)}}
	if p.Keyframes != nil {
		rules := ojson.Object{}
		for _, k := range p.Keyframes {
			frames := ojson.Array{}
			for _, f := range k.Frames {
				frames = append(frames, ojson.Object{{Key: "at", Value: ojson.String(f.At)}, {Key: "style", Value: ojson.String(f.Style)}})
			}
			rules = append(rules, ojson.Member{Key: k.Name, Value: frames})
		}
		o = append(o, ojson.Member{Key: "keyframes", Value: rules})
	}
	if p.Sections != nil {
		sections := ojson.Array{}
		for _, s := range p.Sections {
			sections = append(sections, ojson.Array{ojson.String(s.Name), ojson.Number(strconv.Itoa(s.Rows))})
		}
		o = append(o, ojson.Member{Key: "sections", Value: sections})
	}
	if p.Matches != nil {
		matches := ojson.Object{}
		for _, m := range p.Matches {
			matches = append(matches, ojson.Member{Key: m.Selector, Value: ojson.Bool(m.Matches)})
		}
		o = append(o, ojson.Member{Key: "matches", Value: matches})
	}
	if p.Box != nil {
		box := ojson.Object{{Key: "width", Value: ojson.Float(p.Box.Width)}, {Key: "height", Value: ojson.Float(p.Box.Height)}}
		if p.Box.At != nil {
			box = append(box, ojson.Member{Key: "top", Value: ojson.Float(p.Box.At.Top)}, ojson.Member{Key: "left", Value: ojson.Float(p.Box.At.Left)})
		}
		o = append(o, ojson.Member{Key: "box", Value: box})
	}
	if p.Animations != nil {
		list := ojson.Array{}
		for _, a := range p.Animations {
			an := ojson.Object{{Key: "kind", Value: ojson.String(a.Kind)}, {Key: "what", Value: ojson.String(a.What)}, {Key: "on", Value: ojson.String(a.On)}, {Key: "duration", Value: ojson.Float(a.Duration)}, {Key: "delay", Value: ojson.Float(a.Delay)}}
			if a.Observed != nil {
				an = append(an, ojson.Member{Key: "direction", Value: ojson.String(a.Observed.Direction)})
			}
			an = append(an, ojson.Member{Key: "easing", Value: ojson.String(a.Easing)}, ojson.Member{Key: "from", Value: ojson.String(a.From)}, ojson.Member{Key: "to", Value: ojson.String(a.To)})
			if a.Observed != nil {
				an = append(an, ojson.Member{Key: "seen", Value: ojson.Number(strconv.Itoa(a.Observed.Seen))})
			}
			list = append(list, an)
		}
		o = append(o, ojson.Member{Key: "animations", Value: list})
	}
	if p.States != nil {
		states := ojson.Object{}
		for _, s := range p.States {
			states = append(states, ojson.Member{Key: s.Name, Value: styles(s.Styles)})
		}
		o = append(o, ojson.Member{Key: "states", Value: states})
	}
	return o
}

// JSON is the page as the probe's record: the members in the order it writes them.
func (p Page) JSON() ojson.Object {
	widths := ojson.Object{}
	for _, w := range p.Widths {
		probes := ojson.Object{}
		for _, pr := range w.Probes {
			probes = append(probes, ojson.Member{Key: pr.Name, Value: pr.json()})
		}
		widths = append(widths, ojson.Member{Key: strconv.Itoa(w.Width), Value: probes})
	}
	return ojson.Object{
		{Key: "url", Value: ojson.String(p.URL)},
		{Key: "finalUrl", Value: ojson.String(p.FinalURL)},
		{Key: "browser", Value: ojson.String(p.Browser)},
		{Key: "date", Value: ojson.String(p.Date)},
		{Key: "widths", Value: widths},
	}
}
