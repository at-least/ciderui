package ciderui

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
	"sync"
)

// @invariant: selected is a value of a data-* attribute a rule selects on an
// element with class, anyValue for a bare [data-x]. A bare :not([data-x]) has
// neither: it tells the bare attribute from none, and a value there would draw no
// rule of its own. A rule naming no cui-* class of the attribute's element holds
// nothing to a class (.cui-console-player-end > [data-control="replay"], where the
// player's other data-control values are selected by no rule), so it adds none.
type selected struct {
	class, attribute, value string
	anyValue                bool
}

type vocabulary struct {
	classes map[string]bool
	values  []selected
}

// @external: ARIA's values for each stateful aria-* attribute the parts carry,
// not what a rule selects: a value no rule draws (aria-checked="true", the
// part's default look) is the page's own; a value ARIA lacks is a typo
// assistive technology reads wrong
var ariaValues = map[string][]string{
	"aria-checked":  {"false", "mixed", "true"},
	"aria-current":  {"date", "false", "location", "page", "step", "time", "true"},
	"aria-disabled": {"false", "true"},
	"aria-expanded": {"false", "true"},
	"aria-invalid":  {"false", "grammar", "spelling", "true"},
	"aria-pressed":  {"false", "mixed", "true"},
	"aria-selected": {"false", "true"},
	"aria-sort":     {"ascending", "descending", "none", "other"},
}

// @invariant: a fragment links no stylesheet to tell its own set, so it is
// held to both sets'
const bothSets = ""

func vocabularyOf(set string) vocabulary {
	inAnotherSet := func(path string) bool {
		return set != bothSets && strings.HasPrefix(path, "components/") && !strings.HasPrefix(path, "components/"+set+"/")
	}
	// @trap: not every class a selector here names: base.css's touch target names both
	// sets' parts, and that rule alone does not draw them
	foreign := map[string]bool{}
	for _, path := range cssFiles() {
		if inAnotherSet(path) {
			maps.Copy(foreign, parsed(path, func() map[string]bool { return classesIn(css(path)) }))
		}
	}
	v := vocabulary{classes: map[string]bool{}, values: nil}
	for _, path := range cssFiles() {
		if inAnotherSet(path) {
			continue
		}
		text := css(path)
		for class := range parsed(path, func() map[string]bool { return classesIn(text) }) {
			if !foreign[class] {
				v.classes[class] = true
			}
		}
		for _, compound := range parsed(path, func() []compound { return compoundsIn(text) }) {
			for _, class := range compound.classes {
				if !strings.HasPrefix(class, "cui-") {
					continue
				}
				for _, a := range compound.attributes {
					a.class = class
					v.values = append(v.values, a)
				}
			}
		}
	}
	return v
}

func parsed[T any](path string, parse func() T) T {
	defer func() {
		if r := recover(); r != nil {
			panic(path + ": " + fmt.Sprint(r))
		}
	}()
	return parse()
}

var vocabularies = sync.OnceValue(func() map[string]vocabulary {
	return map[string]vocabulary{bothSets: vocabularyOf(bothSets), "site": vocabularyOf("site"), "console": vocabularyOf("console")}
})

func linkedSet(tags []startTag) string {
	for _, tag := range tags {
		if !strings.EqualFold(tag.name, "link") {
			continue
		}
		switch tag.attribute("href") {
		case SiteStylesheet():
			return "site"
		case ConsoleStylesheet():
			return "console"
		}
	}
	return bothSets
}

func classesIn(css string) map[string]bool {
	classes := map[string]bool{}
	styleRules, utilities := rules(css)
	for _, selector := range styleRules {
		for i := range len(selector) {
			if class := name(selector[i+1:]); selector[i] == '.' && strings.HasPrefix(class, "cui-") {
				classes[class] = true
			}
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

type compound struct {
	classes []string
	// @invariant: each without its class, which the compound's classes give
	attributes []selected
}

// @invariant: subjects returns the ways selector's rightmost compound matches
// its element, and appends every other compound in it, each of another element,
// to others. A :is or :where is the same element, one way per argument: not the
// last class before an attribute, since :where(.a, .b[data-x]) gives .a no
// data-x. A :not is the same element for its attributes alone; a :has is other
// elements.
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
					ways[w].classes = append(ways[w].classes, class)
				}
				i += 1 + len(class)
			case c == '[':
				end := strings.IndexByte(part[i:], ']')
				inside := part[i+1 : i+end]
				if strings.HasPrefix(inside, "data-") {
					attribute, value, hasValue := strings.Cut(inside, "=")
					a := selected{class: "", attribute: attribute, value: strings.Trim(value, `"`), anyValue: !hasValue && !negated}
					for w := range ways {
						ways[w].attributes = append(ways[w].attributes, a)
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
	return text[:indexFunc(text, func(c byte) bool { return !isNameByte(c) })]
}

// @invariant: rules is every style rule's selector in css, nesting resolved,
// and every @utility's name, with its comments gone.
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

// CheckMarkup returns a line for each problem in html, nil when there is none.
// A cui-* class the stylesheet the page links does not hold, either set's for a
// fragment (a misspelt one, or the other set's, styles nothing). A color written
// by its slot (text-(--cui-c-…)) that the set names by a utility. On an element
// with a cui-* class: a value of a data-* attribute some rule selects that no
// rule selects for that element's classes (it renders the default, silently),
// and a value of a stateful aria-* attribute ARIA holds none of (wrong for
// assistive technology, drawn by nothing). On any element, an attribute written
// twice (the browser keeps the first, silently). On a whole page, a reference to
// an id the page lacks (a skip link to no #main moves nothing); a fragment's
// may name the page around it. A data-* or an aria-* no rule selects and every
// other aria-* are the page's own. For a site's tests, on each page it renders;
// site.Page and console.ConsolePage check themselves under go test.
func CheckMarkup(html string) []string {
	tags := startTags(html)
	set := linkedSet(tags)
	known := vocabularies()[set]
	sheet := "ciderui"
	if set != bothSets {
		sheet = "the " + set + " stylesheet this page links"
	}
	slotSets := []string{set}
	if set == bothSets {
		slotSets = []string{"site", "console"}
	}
	var problems []string
	for _, tag := range tags {
		on := tag.on()
		var ours []string
		for _, c := range strings.Fields(tag.attribute("class")) {
			if strings.HasPrefix(c, "cui-") {
				ours = append(ours, c)
				continue
			}
			if utility := slotUtility(c, slotSets); utility != "" {
				problems = append(problems, fmt.Sprintf("class %q on %s: write %s, the set's name for this color", c, on, utility))
			}
		}
		for _, c := range ours {
			if known.classes[c] {
				continue
			}
			if other := otherSet(set); other != "" && vocabularies()[other].classes[c] {
				problems = append(problems, fmt.Sprintf("class %q on %s: a %s class; this page links the %s stylesheet", c, on, other, set))
			} else {
				hint := partClasses(c, known.classes)
				if hint == "" {
					hint = nearest(c, known.classes)
				}
				problems = append(problems, fmt.Sprintf("class %q on %s: %s has no such class%s", c, on, sheet, hint))
			}
		}
		var seen []string
		for _, a := range tag.attributes {
			if slices.Contains(seen, a.name) {
				problems = append(problems, fmt.Sprintf("%s on %s: written twice, and the browser keeps the first", a.written, on))
			}
			seen = append(seen, a.name)
		}
		if len(ours) == 0 {
			continue
		}
		problems = append(problems, dataValueProblems(tag, ours, known.values)...)
		for _, a := range tag.attributes {
			legal, isState := ariaValues[a.name]
			if !isState || slices.Contains(legal, a.value) {
				continue
			}
			problems = append(problems, fmt.Sprintf("%s=%q on %s: not a value of %s (%s)", a.written, a.value, on, a.name, strings.Join(legal, ", ")))
		}
	}
	if isPage(tags) {
		problems = append(problems, missingIDs(tags)...)
	}
	return problems
}

func otherSet(set string) string {
	switch set {
	case "site":
		return "console"
	case "console":
		return "site"
	}
	return ""
}

func partClasses(class string, known map[string]bool) string {
	rest := strings.TrimPrefix(class, "cui-")
	prefix := "cui-"
	if after, ok := strings.CutPrefix(rest, "console-"); ok {
		prefix, rest = "cui-console-", after
	}
	block := prefix + strings.SplitN(rest, "-", 2)[0]
	var part []string
	for k := range known {
		if k == block || strings.HasPrefix(k, block+"-") {
			part = append(part, k)
		}
	}
	slices.Sort(part)
	switch len(part) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("; the only %s class is %s", block, part[0])
	}
	return fmt.Sprintf("; the %s classes are %s", block, inWords(part, "and"))
}

func inWords(items []string, conjunction string) string {
	last := items[len(items)-1]
	if len(items) == 1 {
		return last
	}
	return strings.Join(items[:len(items)-1], ", ") + " " + conjunction + " " + last
}

func valuesTaken(class, attribute string, values []selected) []string {
	var vs []string
	for _, v := range values {
		if v.class != class || v.attribute != attribute {
			continue
		}
		if q := strconv.Quote(v.value); !slices.Contains(vs, q) {
			vs = append(vs, q)
		}
	}
	slices.Sort(vs)
	return vs
}

func attributesTaken(classes []string, values []selected) string {
	var takes []string
	for _, c := range classes {
		var attributes []string
		for _, v := range values {
			if v.class == c && !slices.Contains(attributes, v.attribute) {
				attributes = append(attributes, v.attribute)
			}
		}
		slices.Sort(attributes)
		for _, attribute := range attributes {
			if slices.ContainsFunc(values, func(v selected) bool { return v.class == c && v.attribute == attribute && v.anyValue }) {
				takes = append(takes, fmt.Sprintf("%s takes %s", c, attribute))
				continue
			}
			takes = append(takes, fmt.Sprintf("%s takes %s %s", c, attribute, inWords(valuesTaken(c, attribute, values), "or")))
		}
	}
	if len(takes) == 0 {
		return ""
	}
	return "; " + strings.Join(takes, "; ")
}

func dataValueProblems(tag startTag, ours []string, values []selected) []string {
	var problems []string
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
		var takes []string
		for _, c := range ours {
			vs := valuesTaken(c, a.name, values)
			if len(vs) == 0 {
				continue
			}
			takes = append(takes, fmt.Sprintf("%s takes %s %s", c, a.name, inWords(vs, "or")))
		}
		var allowed string
		switch len(takes) {
		case 0:
			allowed = "none of its classes takes " + a.name + attributesTaken(ours, values)
		case 1:
			allowed = takes[0] + ", or none for its default"
		default:
			allowed = strings.Join(takes, "; ") + "; or none for their defaults"
		}
		problems = append(problems, fmt.Sprintf("%s=%q on %s: %s", a.written, a.value, tag.on(), allowed))
	}
	return problems
}

func isPage(tags []startTag) bool {
	return slices.ContainsFunc(tags, func(t startTag) bool { return strings.EqualFold(t.name, "html") })
}

// @external: the attributes that name ids of the same page; an id holds no
// space, so each is read as a list (aria-describedby's and <output for>'s are)
var idReferences = []string{"popovertarget", "commandfor", "for", "aria-labelledby", "aria-describedby", "aria-controls"}

func missingIDs(tags []startTag) []string {
	ids := map[string]bool{}
	for _, tag := range tags {
		for _, a := range tag.attributes {
			if a.name == "id" {
				ids[a.value] = true
			}
		}
	}
	var problems []string
	for _, tag := range tags {
		for _, a := range tag.attributes {
			var named []string
			if slices.Contains(idReferences, a.name) {
				named = strings.Fields(a.value)
			} else if fragment, ok := strings.CutPrefix(a.value, "#"); ok && a.name == "href" && fragment != "" {
				named = []string{fragment}
			}
			for _, id := range named {
				if !ids[id] {
					problems = append(problems, fmt.Sprintf("%s=%q on %s: this page has no id %q", a.written, a.value, tag.on(), id))
				}
			}
		}
	}
	return problems
}

// @external: the browser reads an attribute's name in lower case (CLASS is the class);
// written is the page's own spelling, for the messages
type attribute struct{ name, written, value string }

type startTag struct {
	name       string
	attributes []attribute
}

func (t startTag) on() string {
	if class := t.attribute("class"); class != "" {
		return fmt.Sprintf("<%s class=%q>", t.name, class)
	}
	return "<" + t.name + ">"
}

func (t startTag) attribute(name string) string {
	for _, a := range t.attributes {
		if a.name == name {
			return a.value
		}
	}
	return ""
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
			a := attribute{name: strings.ToLower(rest[:end]), written: rest[:end]}
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
