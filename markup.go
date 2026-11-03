package ciderui

import (
	"fmt"
	"slices"
	"strings"
	"sync"
)

var checkedAttributes = []string{"data-variant", "data-size", "data-color"}

type value struct{ class, attribute, value string }

// classes: every cui-* name anywhere in the stylesheets, comments included: a
// part with no rule of its own (cui-post-content) is named only in a comment.
// values: each checked attribute value a rule selects, with the last .cui-*
// class before it in the selector.
var words = sync.OnceValues(func() (map[string]bool, []value) {
	classes := map[string]bool{}
	var values []value
	for _, path := range CSSFiles() {
		text := CSS(path)
		for at := 0; ; {
			i := strings.Index(text[at:], "cui-")
			if i < 0 {
				break
			}
			i += at
			at = i + 4
			if i > 0 && isNameByte(text[i-1]) {
				continue
			}
			classes[name(text[i:])] = true
		}
		for _, selector := range selectors(text) {
			for at := 0; ; {
				i := strings.Index(selector[at:], "[data-")
				if i < 0 {
					break
				}
				i += at
				at = i + 1
				inside, _, ok := strings.Cut(selector[i+1:], "]")
				if !ok {
					continue
				}
				attribute, v, ok := strings.Cut(inside, "=")
				if !ok || !slices.Contains(checkedAttributes, attribute) {
					continue
				}
				c := strings.LastIndex(selector[:i], ".cui-")
				if c < 0 {
					panic(fmt.Sprintf("%s: %s=%s with no cui-* class before it", selector, attribute, v))
				}
				values = append(values, value{name(selector[c+1:]), attribute, strings.Trim(v, `"`)})
			}
		}
	}
	return classes, values
})

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

// every style rule's selector in css, nesting resolved
func selectors(css string) []string {
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
	var out []string
	stack := [][]string{nil}
	var head strings.Builder
	for _, c := range text.String() {
		switch c {
		case '{':
			parents := stack[len(stack)-1]
			prelude := strings.TrimSpace(head.String())
			head.Reset()
			if strings.HasPrefix(prelude, "@") {
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
			out = append(out, full...)
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
	return out
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
// one styles nothing), and on an element with a cui-* class, a
// checkedAttributes value no rule selects (it renders the default, silently).
// For a site's tests, on each page it renders.
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
			if !slices.Contains(checkedAttributes, a.name) {
				continue
			}
			selected := slices.ContainsFunc(values, func(v value) bool {
				return v.attribute == a.name && v.value == a.value && slices.Contains(ours, v.class)
			})
			if selected {
				continue
			}
			var known []string
			for _, c := range ours {
				var vs []string
				for _, v := range values {
					if v.class == c && v.attribute == a.name && !slices.Contains(vs, v.value) {
						vs = append(vs, v.value)
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
