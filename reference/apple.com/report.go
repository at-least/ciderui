package main

import (
	"bytes"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf16"

	"github.com/at-least/ciderui/internal/ojson"
	"github.com/at-least/ciderui/reference/apple.com/results"
)

// @external: JavaScript's \s: the ASCII spaces and Unicode's, which Go's \s leaves out
var space = regexp.MustCompile(`[\t\n\v\f\r \x{a0}\x{1680}\x{2000}-\x{200a}\x{2028}\x{2029}\x{202f}\x{205f}\x{3000}\x{feff}]+`)

// @invariant: a value as the Node report printed it: a string as itself, undefined and an
// object as JavaScript's template literal writes them, an object never === another
type shown struct {
	text   string
	object bool
}

func (v shown) String() string { return v.text }

var undefined = shown{text: "undefined", object: false}

func report(name, prefix string) error {
	page, ok := results.Pages[name]
	if !ok {
		return fmt.Errorf("results/ holds no record %s", name)
	}
	// @invariant: the record as its JSON, which the report read before the records were
	// Go: what it prints is the Node report's
	r := page.JSON()
	var widths []string
	at := map[string]ojson.Object{}
	for _, m := range r.Get("widths").(ojson.Object) {
		widths = append(widths, m.Key)
		at[m.Key] = m.Value.(ojson.Object)
	}
	fmt.Printf("%s: %s %s %s\n", name, page.FinalURL, page.Browser, page.Date)
	var probes []string
	seen := map[string]bool{}
	for _, w := range widths {
		for _, m := range at[w] {
			if !seen[m.Key] && strings.HasPrefix(m.Key, prefix) {
				seen[m.Key] = true
				probes = append(probes, m.Key)
			}
		}
	}
	// @invariant: a probe's entry at a width, nil where it is null or not read there
	entry := func(w, p string) ojson.Object {
		o, _ := at[w].Get(p).(ojson.Object)
		return o
	}
	// @external: a typography class that reads as the unclassed paragraph at every width is
	// one the page does not define
	control := func(p string) bool {
		if !strings.HasPrefix(p, "type.") || p == "type.none" {
			return false
		}
		for _, w := range widths {
			none := entry(w, "type.none")
			if none == nil {
				return false
			}
			mine := entry(w, p)
			if mine == nil || compact(mine.Get("styles")) != compact(none.Get("styles")) {
				return false
			}
		}
		return true
	}
	for _, p := range probes {
		if control(p) {
			fmt.Printf("\n%s  NOT DEFINED HERE: reads as type.none at every width\n", p)
			continue
		}
		head := "null"
		if first := entry(widths[0], p); first != nil {
			if h, ok := first.Get("html").(ojson.String); ok {
				head = string(h)
			}
		}
		fmt.Printf("\n%s  %s\n", p, utf16Prefix(space.ReplaceAllString(head, " "), 150))
		var keys []string
		has := map[string]bool{}
		add := func(k string) {
			if !has[k] {
				has[k] = true
				keys = append(keys, k)
			}
		}
		for _, w := range widths {
			v := entry(w, p)
			if v == nil {
				continue
			}
			styles, _ := v.Get("styles").(ojson.Object)
			for _, m := range styles {
				add(m.Key)
			}
			if v.Get("box") != nil {
				add("box.width")
				add("box.height")
			}
			// @external: a state's value only where it differs from rest, and never its transition
			// (the harness turns transitions off while it forces a state)
			states, _ := v.Get("states").(ojson.Object)
			for _, s := range states {
				st, _ := s.Value.(ojson.Object)
				for _, m := range st {
					k := m.Key
					if strings.HasPrefix(k, "transition") {
						continue
					}
					for _, w2 := range widths {
						x := entry(w2, p)
						if x == nil {
							continue
						}
						xs, _ := x.Get("states").(ojson.Object)
						xst, _ := xs.Get(s.Key).(ojson.Object)
						xstyles, _ := x.Get("styles").(ojson.Object)
						if strictlyDiffer(xst.Get(k), xstyles.Get(k)) {
							add(":" + s.Key + " " + k)
							break
						}
					}
				}
			}
		}
		for _, k := range keys {
			vals := make([]shown, len(widths))
			for i, w := range widths {
				v := entry(w, p)
				switch {
				case v == nil:
					vals[i] = shown{text: "null", object: false}
				case strings.HasPrefix(k, "box."):
					box, ok := v.Get("box").(ojson.Object)
					if !ok {
						return fmt.Errorf("%s at %s: no box to read %s", p, w, k)
					}
					n, ok := box.Get(k[4:]).(ojson.Number)
					if !ok {
						return fmt.Errorf("%s at %s: %s is no number", p, w, k)
					}
					s, err := toFixed3(string(n))
					if err != nil {
						return err
					}
					vals[i] = shown{text: s, object: false}
				case strings.HasPrefix(k, ":"):
					state, prop, _ := strings.Cut(k[1:], " ")
					states, _ := v.Get("states").(ojson.Object)
					st, _ := states.Get(state).(ojson.Object)
					vals[i] = templated(st.Get(prop), "-")
				default:
					styles, _ := v.Get("styles").(ojson.Object)
					// @invariant: a picture read as what it draws (`image`) prints as that
					switch x := styles.Get(k).(type) {
					case nil:
						vals[i] = undefined
					case ojson.Object, ojson.Array, ojson.Null:
						vals[i] = shown{text: compact(x), object: false}
					default:
						vals[i] = templated(x, "")
					}
				}
			}
			same := true
			for i, v := range vals {
				if i > 0 && (v.object || vals[0].object || v.text != vals[0].text) {
					same = false
				}
			}
			line := ""
			if same {
				line = "all " + vals[0].text
			} else {
				var parts []string
				for i, w := range widths {
					parts = append(parts, w+":"+vals[i].text)
				}
				line = strings.Join(parts, "  ")
			}
			fmt.Printf("  %-28s %s\n", k, line)
		}
	}
	return nil
}

// @invariant: v as a template literal writes it, missing (undefined or null, the ?? of
// Node report) as orElse; "" for orElse writes undefined and null as themselves
func templated(v ojson.Value, orElse string) shown {
	switch x := v.(type) {
	case nil:
		if orElse != "" {
			return shown{text: orElse, object: false}
		}
		return undefined
	case ojson.Null:
		if orElse != "" {
			return shown{text: orElse, object: false}
		}
		return shown{text: "null", object: false}
	case ojson.String:
		return shown{text: string(x), object: false}
	case ojson.Number:
		return shown{text: string(x), object: false}
	case ojson.Bool:
		return shown{text: strconv.FormatBool(bool(x)), object: false}
	}
	return shown{text: "[object Object]", object: true}
}

// @invariant: JavaScript's a !== b on two parsed JSON values: an object or array is
// never === another
func strictlyDiffer(a, b ojson.Value) bool {
	switch x := a.(type) {
	case nil:
		return b != nil
	case ojson.Null:
		_, same := b.(ojson.Null)
		return !same
	case ojson.String:
		y, ok := b.(ojson.String)
		return !ok || x != y
	case ojson.Bool:
		y, ok := b.(ojson.Bool)
		return !ok || x != y
	case ojson.Number:
		y, ok := b.(ojson.Number)
		if !ok {
			return true
		}
		fx, errx := strconv.ParseFloat(string(x), 64)
		fy, erry := strconv.ParseFloat(string(y), 64)
		return errx != nil || erry != nil || fx != fy
	}
	return true
}

// @external: String(+x.toFixed(3)): the integer nearest x·1000, a tie the larger in
// magnitude (Go's FormatFloat rounds a tie to even), then the shortest number
func toFixed3(literal string) (string, error) {
	x, err := strconv.ParseFloat(literal, 64)
	if err != nil {
		return "", err
	}
	r := new(big.Rat).SetFloat64(x)
	negative := r.Sign() < 0
	r.Abs(r)
	r.Mul(r, big.NewRat(1000, 1))
	r.Add(r, big.NewRat(1, 2))
	n := new(big.Int).Quo(r.Num(), r.Denom())
	f, _ := new(big.Rat).SetFrac(n, big.NewInt(1000)).Float64()
	if negative {
		f = -f
	}
	return string(ojson.Float(f)), nil
}

// @external: String.prototype.slice's prefix: n UTF-16 code units, not runes
func utf16Prefix(s string, n int) string {
	units := 0
	for i, r := range s {
		units += len(utf16.Encode([]rune{r}))
		if units > n {
			return s[:i]
		}
	}
	return s
}

// @invariant: JSON.stringify(v) with no indent; ojson writes only JSON.stringify(v,
// null, 1), so a leaf is its text and the structure is written here
func compact(v ojson.Value) string {
	var b bytes.Buffer
	var write func(v ojson.Value)
	write = func(v ojson.Value) {
		switch x := v.(type) {
		case nil:
			b.WriteString("undefined")
		case ojson.Object:
			b.WriteByte('{')
			first := true
			for _, m := range x {
				if m.Value == nil {
					continue
				}
				if !first {
					b.WriteByte(',')
				}
				first = false
				b.Write(ojson.Marshal(ojson.String(m.Key)))
				b.WriteByte(':')
				write(m.Value)
			}
			b.WriteByte('}')
		case ojson.Array:
			b.WriteByte('[')
			for i, e := range x {
				if i > 0 {
					b.WriteByte(',')
				}
				if e == nil {
					b.WriteString("null")
				} else {
					write(e)
				}
			}
			b.WriteByte(']')
		default:
			b.Write(ojson.Marshal(x))
		}
	}
	write(v)
	return b.String()
}
