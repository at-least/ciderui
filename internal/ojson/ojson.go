// Package ojson is JSON as a page's JavaScript builds it: objects keep their keys'
// order, numbers their literal text, and Marshal writes what JSON.stringify(v, null, 1)
// writes, so a record the Node probes wrote is rewritten byte for byte.
package ojson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Value is one JSON value: Object, Array, String, Number, Bool or Null. A nil Value is
// JavaScript's undefined: Marshal leaves an object's member of it out and writes it
// as null in an array, as JSON.stringify does.
type Value interface{ value() }

// Member is one key of an Object and its value.
type Member struct {
	Key   string
	Value Value
}

// Object is a JSON object, its members in order.
type Object []Member

// Array is a JSON array.
type Array []Value

// String is a JSON string.
type String string

// Number is a JSON number's literal text.
type Number string

// Bool is a JSON boolean.
type Bool bool

// Null is JSON's null.
type Null struct{}

func (Object) value() {}
func (Array) value()  {}
func (String) value() {}
func (Number) value() {}
func (Bool) value()   {}
func (Null) value()   {}

// Float is v as JavaScript writes a number: the shortest text that reads back as v,
// with an exponent below 1e-6 and from 1e21.
func Float(v float64) Number {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		panic(fmt.Sprintf("ojson: %v has no JSON form", v))
	}
	if v == 0 {
		return "0"
	}
	if a := math.Abs(v); a >= 1e21 || a < 1e-6 {
		s := strconv.FormatFloat(v, 'e', -1, 64)
		mantissa, exp, _ := strings.Cut(s, "e")
		sign := exp[0]
		exp = strings.TrimLeft(exp[1:], "0")
		return Number(mantissa + "e" + string(sign) + exp)
	}
	return Number(strconv.FormatFloat(v, 'f', -1, 64))
}

// Get is the value of key, or nil.
func (o Object) Get(key string) Value {
	for _, m := range o {
		if m.Key == key {
			return m.Value
		}
	}
	return nil
}

// With is o and members spread over it, as {...o, ...members}: a key o has keeps its
// place and takes the new value, a new key goes last.
func (o Object) With(members ...Member) Object {
	out := append(Object(nil), o...)
	for _, m := range members {
		i := -1
		for j, have := range out {
			if have.Key == m.Key {
				i = j
			}
		}
		if i >= 0 {
			out[i].Value = m.Value
		} else {
			out = append(out, m)
		}
	}
	return out
}

// Parse reads one JSON value, keeping its objects' order and its numbers' text.
func Parse(raw []byte) (Value, error) {
	d := json.NewDecoder(bytes.NewReader(raw))
	d.UseNumber()
	v, err := parse(d)
	if err != nil {
		return nil, fmt.Errorf("ojson: %w", err)
	}
	if _, err := d.Token(); !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("ojson: more after the value")
	}
	return v, nil
}

func parse(d *json.Decoder) (Value, error) {
	t, err := d.Token()
	if err != nil {
		return nil, err
	}
	switch t := t.(type) {
	case json.Delim:
		switch t {
		case '{':
			o := Object{}
			for d.More() {
				k, err := d.Token()
				if err != nil {
					return nil, err
				}
				key, ok := k.(string)
				if !ok {
					return nil, fmt.Errorf("%v where a key belongs", k)
				}
				v, err := parse(d)
				if err != nil {
					return nil, err
				}
				o = append(o, Member{key, v})
			}
			if _, err := d.Token(); err != nil {
				return nil, err
			}
			return o, nil
		case '[':
			a := Array{}
			for d.More() {
				v, err := parse(d)
				if err != nil {
					return nil, err
				}
				a = append(a, v)
			}
			if _, err := d.Token(); err != nil {
				return nil, err
			}
			return a, nil
		}
		return nil, fmt.Errorf("%v where a value belongs", t)
	case string:
		return String(t), nil
	case json.Number:
		return Number(t), nil
	case bool:
		return Bool(t), nil
	case nil:
		return Null{}, nil
	}
	return nil, fmt.Errorf("%v where a value belongs", t)
}

// Marshal is JSON.stringify(v, null, 1).
func Marshal(v Value) []byte {
	var b bytes.Buffer
	write(&b, v, 0)
	return b.Bytes()
}

func write(b *bytes.Buffer, v Value, depth int) {
	indent := func(d int) {
		b.WriteByte('\n')
		b.WriteString(strings.Repeat(" ", d))
	}
	switch v := v.(type) {
	case Object:
		if len(v) == 0 {
			b.WriteString("{}")
			return
		}
		var defined Object
		for _, m := range v {
			if m.Value != nil {
				defined = append(defined, m)
			}
		}
		if len(defined) == 0 {
			b.WriteString("{}")
			return
		}
		b.WriteByte('{')
		for i, m := range defined {
			if i > 0 {
				b.WriteByte(',')
			}
			indent(depth + 1)
			quote(b, m.Key)
			b.WriteString(": ")
			write(b, m.Value, depth+1)
		}
		indent(depth)
		b.WriteByte('}')
	case Array:
		if len(v) == 0 {
			b.WriteString("[]")
			return
		}
		b.WriteByte('[')
		for i, e := range v {
			if i > 0 {
				b.WriteByte(',')
			}
			indent(depth + 1)
			write(b, e, depth+1)
		}
		indent(depth)
		b.WriteByte(']')
	case String:
		quote(b, string(v))
	case Number:
		b.WriteString(string(v))
	case Bool:
		b.WriteString(strconv.FormatBool(bool(v)))
	case Null, nil:
		b.WriteString("null")
	default:
		panic(fmt.Sprintf("ojson: %T is no Value", v))
	}
}

// @external: JSON.stringify's escapes: the quote, the backslash, and the control
// characters (\b \f \n \r \t by name, the rest as \u00xx); nothing else
func quote(b *bytes.Buffer, s string) {
	if !utf8.ValidString(s) {
		panic(fmt.Sprintf("ojson: %q is not UTF-8", s))
	}
	b.WriteByte('"')
	for _, r := range s {
		switch r {
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		case '\b':
			b.WriteString(`\b`)
		case '\f':
			b.WriteString(`\f`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		default:
			if r < 0x20 {
				fmt.Fprintf(b, `\u%04x`, r)
			} else {
				b.WriteRune(r)
			}
		}
	}
	b.WriteByte('"')
}
