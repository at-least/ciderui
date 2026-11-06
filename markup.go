package ciderui

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"sync"
)

// selected is a value of a data-* attribute a rule selects on an element with
// class, anyValue for a bare [data-x]. A bare :not([data-x]) has neither: it
// tells the bare attribute from none, and a value there would draw no rule of
// its own. A rule naming no cui-* class of the attribute's element holds nothing
// to a class (.cui-console-player-end > [data-control="replay"], where the
// player's other data-control values are selected by no rule), so it adds none.
type selected struct {
	class, attribute, value string
	anyValue                bool
}

var words = sync.OnceValues(func() (map[string]bool, []selected) {
	classes := map[string]bool{}
	var values []selected
	for _, path := range cssFiles() {
		text := css(path)
		for class := range classesIn(text) {
			classes[class] = true
		}
		for _, compound := range compoundsIn(text) {
			for _, class := range compound.classes {
				if !strings.HasPrefix(class, "cui-") {
					continue
				}
				for _, a := range compound.attributes {
					values = append(values, selected{class: class, attribute: a.name, value: a.value, anyValue: !a.hasValue && !a.negated})
				}
			}
		}
	}
	return classes, values
})

var selectorClass = regexp.MustCompile(`\.(cui-[a-zA-Z0-9-]+)`)

// classesIn is every class a selector of css names and every @utility of it.
func classesIn(css string) map[string]bool {
	classes := map[string]bool{}
	styleRules, utilities := rules(css)
	for _, selector := range styleRules {
		for _, match := range selectorClass.FindAllStringSubmatch(selector, -1) {
			classes[match[1]] = true
		}
	}
	for _, utility := range utilities {
		if strings.HasPrefix(utility, "cui-") {
			classes[utility] = true
		}
	}
	return classes
}

func compoundsIn(css string) []compound {
	var own, others []compound
	styleRules, _ := rules(css)
	for _, selector := range styleRules {
		own = append(own, subjects(selector, false, &others)...)
	}
	return append(own, others...)
}

// selectorAttribute is a [data-*] of a selector.
type selectorAttribute struct {
	name, value       string
	hasValue, negated bool
}

// compound is one way a compound selector matches one element.
type compound struct {
	classes    []string
	attributes []selectorAttribute
}

// subjects returns the ways selector's rightmost compound matches its element,
// and appends every other compound in it, each of another element, to others.
// A :is or :where is the same element, one way per argument: not the last class
// before an attribute, since :where(.a, .b[data-x]) gives .a no data-x. A :not
// is the same element for its attributes alone; a :has is other elements.
func subjects(selector string, negated bool, others *[]compound) []compound {
	var all []compound
	for _, part := range split(selector) {
		ways := []compound{{classes: nil, attributes: nil}}
		for i := 0; i < len(part); {
			switch c := part[i]; {
			case c == ' ' || c == '>' || c == '+' || c == '~':
				for _, way := range ways {
					if len(way.classes)+len(way.attributes) > 0 {
						*others = append(*others, way)
					}
				}
				ways = []compound{{classes: nil, attributes: nil}}
				i++
			case c == '.':
				class := name(part[i+1:])
				for w := range ways {
					ways[w].classes = append(slices.Clip(ways[w].classes), class)
				}
				i += 1 + len(class)
			case c == '[':
				end := strings.IndexByte(part[i:], ']')
				inside := part[i+1 : i+end]
				if strings.HasPrefix(inside, "data-") {
					attribute, value, hasValue := strings.Cut(inside, "=")
					a := selectorAttribute{name: attribute, value: strings.Trim(value, `"`), hasValue: hasValue, negated: negated}
					for w := range ways {
						ways[w].attributes = append(slices.Clip(ways[w].attributes), a)
					}
				}
				i += end + 1
			case c == ':':
				i++
				for i < len(part) && part[i] == ':' {
					i++
				}
				pseudo := name(part[i:])
				i += len(pseudo)
				if i == len(part) || part[i] != '(' {
					continue
				}
				end := closingParen(part, i)
				arguments := part[i+1 : end]
				i = end + 1
				switch pseudo {
				case "is", "where":
					ways = alongside(ways, subjects(arguments, negated, others), true)
				case "not":
					ways = alongside(ways, subjects(arguments, true, others), false)
				case "has":
					*others = append(*others, subjects(arguments, false, others)...)
				}
			default:
				i++
			}
		}
		all = append(all, ways...)
	}
	return all
}

func closingParen(selector string, open int) int {
	depth := 0
	for i := open; i < len(selector); i++ {
		switch selector[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	panic("an unclosed ( in " + selector)
}

// alongside is each way with each way of a pseudo-class's argument.
func alongside(ways, arguments []compound, withClasses bool) []compound {
	var out []compound
	for _, way := range ways {
		for _, argument := range arguments {
			merged := compound{classes: slices.Clone(way.classes), attributes: slices.Concat(way.attributes, argument.attributes)}
			if withClasses {
				merged.classes = append(merged.classes, argument.classes...)
			}
			out = append(out, merged)
		}
	}
	return out
}

func isNameByte(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '-'
}

func name(text string) string {
	end := 0
	for end < len(text) && isNameByte(text[end]) {
		end++
	}
	return text[:end]
}

// rules is every style rule's selector in css, nesting resolved, and every
// @utility's name, with its comments gone.
func rules(css string) (selectors, utilities []string) {
	var text strings.Builder
	rest := css
	for {
		i := strings.Index(rest, "/*")
		if i < 0 {
			break
		}
		text.WriteString(rest[:i])
		end := strings.Index(rest[i:], "*/")
		if end < 0 {
			panic("an unclosed comment")
		}
		rest = rest[i+end+2:]
	}
	text.WriteString(rest)
	stack := [][]string{nil}
	var head strings.Builder
	for _, c := range text.String() {
		switch c {
		case '{':
			parents := stack[len(stack)-1]
			prelude := strings.TrimSpace(head.String())
			head.Reset()
			if strings.HasPrefix(prelude, "@") {
				if utility, ok := strings.CutPrefix(prelude, "@utility "); ok {
					utilities = append(utilities, utility)
				}
				stack = append(stack, parents)
				continue
			}
			var full []string
			for _, part := range split(prelude) {
				if len(parents) == 0 {
					full = append(full, part)
					continue
				}
				for _, parent := range parents {
					if strings.Contains(part, "&") {
						full = append(full, strings.ReplaceAll(part, "&", parent))
					} else {
						full = append(full, parent+" "+part)
					}
				}
			}
			selectors = append(selectors, full...)
			stack = append(stack, full)
		case '}':
			stack = stack[:len(stack)-1]
			head.Reset()
		case ';':
			head.Reset()
		default:
			head.WriteRune(c)
		}
	}
	return selectors, utilities
}

func split(selector string) []string {
	var parts []string
	depth, start := 0, 0
	for i, c := range selector {
		switch {
		case c == '(' || c == '[':
			depth++
		case c == ')' || c == ']':
			depth--
		case c == ',' && depth == 0:
			parts = append(parts, strings.TrimSpace(selector[start:i]))
			start = i + 1
		}
	}
	return append(parts, strings.TrimSpace(selector[start:]))
}

// CheckMarkup returns a line for each word of ciderui's in html that the kit's
// stylesheets do not know, nil when there is none: a cui-* class (a misspelt
// one styles nothing), and on an element with a cui-* class, a value of a data-*
// attribute some rule selects that no rule selects for that element's classes
// (it renders the default, silently). A data-* no rule selects is the page's
// own. For a site's tests, on each page it renders.
func CheckMarkup(html string) []string {
	classes, values := words()
	var problems []string
	for _, tag := range startTags(html) {
		class := ""
		for _, a := range tag.attributes {
			if a.name == "class" {
				class = a.value
				break
			}
		}
		var ours []string
		for _, c := range strings.Fields(class) {
			if strings.HasPrefix(c, "cui-") {
				ours = append(ours, c)
			}
		}
		for _, c := range ours {
			if !classes[c] {
				problems = append(problems, fmt.Sprintf("class %q on <%s class=%q>: ciderui has no such class", c, tag.name, class))
			}
		}
		if len(ours) == 0 {
			continue
		}
		for _, a := range tag.attributes {
			if !slices.ContainsFunc(values, func(v selected) bool { return v.attribute == a.name }) {
				continue
			}
			ok := slices.ContainsFunc(values, func(v selected) bool {
				return v.attribute == a.name && slices.Contains(ours, v.class) && (v.anyValue || v.value == a.value)
			})
			if ok {
				continue
			}
			var known []string
			for _, c := range ours {
				var vs []string
				for _, v := range values {
					if v.class == c && v.attribute == a.name {
						shown := v.value
						if shown == "" {
							shown = `""`
						}
						if !slices.Contains(vs, shown) {
							vs = append(vs, shown)
						}
					}
				}
				slices.Sort(vs)
				list := "none"
				if len(vs) > 0 {
					list = strings.Join(vs, ", ")
				}
				known = append(known, c+"'s: "+list)
			}
			problems = append(problems, fmt.Sprintf("%s=%q on <%s class=%q>: no rule of ciderui's selects it (%s)", a.name, a.value, tag.name, class, strings.Join(known, "; ")))
		}
	}
	return problems
}

type attribute struct{ name, value string }

type startTag struct {
	name       string
	attributes []attribute
}

func isSpace(c byte) bool { return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f' }

func indexFunc(s string, stop func(byte) bool) int {
	for i := 0; i < len(s); i++ {
		if stop(s[i]) {
			return i
		}
	}
	return len(s)
}

func startTags(html string) []startTag {
	var tags []startTag
	rest := html
	for {
		i := strings.IndexByte(rest, '<')
		if i < 0 {
			return tags
		}
		rest = rest[i+1:]
		if after, ok := strings.CutPrefix(rest, "!--"); ok {
			end := strings.Index(after, "-->")
			if end < 0 {
				rest = ""
			} else {
				rest = after[end+3:]
			}
			continue
		}
		if rest == "" || !(rest[0] >= 'a' && rest[0] <= 'z' || rest[0] >= 'A' && rest[0] <= 'Z') {
			continue
		}
		end := indexFunc(rest, func(c byte) bool { return isSpace(c) || c == '>' || c == '/' })
		tag := startTag{name: rest[:end]}
		rest = rest[end:]
		for {
			rest = rest[indexFunc(rest, func(c byte) bool { return !isSpace(c) && c != '/' }):]
			if rest == "" {
				break
			}
			if rest[0] == '>' {
				rest = rest[1:]
				break
			}
			end := indexFunc(rest, func(c byte) bool { return isSpace(c) || c == '=' || c == '>' || c == '/' })
			a := attribute{name: rest[:end]}
			rest = rest[end:]
			if after, ok := strings.CutPrefix(rest, "="); ok {
				if after != "" && (after[0] == '"' || after[0] == '\'') {
					close := strings.IndexByte(after[1:], after[0])
					if close < 0 {
						a.value, rest = after[1:], ""
					} else {
						a.value, rest = after[1:close+1], after[close+2:]
					}
				} else {
					end := indexFunc(after, func(c byte) bool { return isSpace(c) || c == '>' })
					a.value, rest = after[:end], after[end:]
				}
			}
			tag.attributes = append(tag.attributes, a)
		}
		if strings.EqualFold(tag.name, "script") || strings.EqualFold(tag.name, "style") {
			close := strings.Index(rest, "</"+tag.name)
			if close < 0 {
				rest = ""
			} else {
				rest = rest[close:]
			}
		}
		tags = append(tags, tag)
	}
}
