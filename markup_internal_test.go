package ciderui

import (
	"fmt"
	"maps"
	"strings"
	"testing"
)

func TestClassesInReadsSelectorsAndUtilitiesNotComments(t *testing.T) {
	classes := classesIn(`/* .cui-ghost { } */ .cui-real { color: red; } @media (width >= 1px) { .cui-inner { color: red; } } @utility cui-type { color: red; } @utility other { color: red; }`)
	want := map[string]bool{"cui-real": true, "cui-inner": true, "cui-type": true}
	if !maps.Equal(classes, want) {
		t.Errorf("classes read %v, want %v", classes, want)
	}
}

// @invariant: each way a selector's compound names a class and a data-*, on one line
func compounds(css string) string {
	var lines []string
	for _, c := range compoundsIn(css) {
		line := "." + strings.Join(c.classes, ".")
		for _, a := range c.attributes {
			not := ""
			if a.negated {
				not = "!"
			}
			line += fmt.Sprintf("[%s%s=%s,%v]", not, a.name, a.value, a.hasValue)
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, " ")
}

func TestCompoundsInReadsEachElementsClassesAndDataAttributes(t *testing.T) {
	for _, c := range []struct{ css, want string }{
		{`.a[data-x="1"] { }`, `.a[data-x=1,true]`},
		{`.a[data-x] { }`, `.a[data-x=,false]`},
		{`.a[title="1"] { }`, `.a`},
		{`.a[data-x="1"], .b[data-y="2"] { }`, `.a[data-x=1,true] .b[data-y=2,true]`},
		{`.a[data-x="1"].b { }`, `.a.b[data-x=1,true]`},
		{`.a:hover[data-x="1"]::before { }`, `.a[data-x=1,true]`},
		{`.a:nth-child(2n + 1)[data-x="1"] { }`, `.a[data-x=1,true]`},
		{`.a:not([data-x]) { }`, `.a[!data-x=,false]`},
		{`.a:not(.b)[data-x="1"] { }`, `.a[data-x=1,true]`},
		{`:is(.a, .b)[data-x="1"] { }`, `.a[data-x=1,true] .b[data-x=1,true]`},
		{`:where(.a, .b[data-x="1"]) { }`, `.a .b[data-x=1,true]`},
		// @invariant: the subject's first: a compound of another element comes after
		{`.a[data-x="1"] .b[data-y="2"] { }`, `.b[data-y=2,true] .a[data-x=1,true]`},
		{`.a[data-x="1"] > .b { }`, `.b .a[data-x=1,true]`},
		{`.a[data-x="1"] + .b { }`, `.b .a[data-x=1,true]`},
		{`.a[data-x="1"] ~ .b { }`, `.b .a[data-x=1,true]`},
		{`.a[data-x="1"]>.b { }`, `.b .a[data-x=1,true]`},
		{`.a[data-x="1"]+.b { }`, `.b .a[data-x=1,true]`},
		{`.a[data-x="1"]~.b { }`, `.b .a[data-x=1,true]`},
		{`.a:has(.b[data-x="1"]) { }`, `.a .b[data-x=1,true]`},
		{`.a { &[data-x="1"] { } }`, `.a .a[data-x=1,true]`},
		{`.a, .b { &[data-x="1"] { } }`, `.a .b .a[data-x=1,true] .b[data-x=1,true]`},
		{`.a { .b[data-x="1"] { } }`, `.a .b[data-x=1,true] .a`},
		{`@media (width >= 1px) { .a[data-x="1"] { color: red; } }`, `.a[data-x=1,true]`},
		{`.a { @media (width >= 1px) { &[data-x="1"] { } } }`, `.a .a[data-x=1,true]`},
		{`.a { margin: .5rem; &[data-x="1"] { } }`, `.a .a[data-x=1,true]`},
		{`.a { margin: .5rem } .b[data-x="1"] { }`, `.a .b[data-x=1,true]`},
		{`.a-b1C[data-x="1"] { }`, `.a-b1C[data-x=1,true]`},
	} {
		if got := compounds(c.css); got != c.want {
			t.Errorf("%s\n read %s\n want %s", c.css, got, c.want)
		}
	}
}

// a stylesheet that does not parse names itself: an unclosed comment in one
// of forty would otherwise be found by bisection
func TestAParseFailureNamesTheStylesheet(t *testing.T) {
	defer func() {
		got := fmt.Sprint(recover())
		if want := "kit.css: an unclosed comment"; got != want {
			t.Errorf("panicked with %q, want %q", got, want)
		}
	}()
	parsed("kit.css", func() map[string]bool { return classesIn("/*") })
}

// nothing else would say so: rules would go on reading the rest as selectors
func TestAnUnclosedParenPanics(t *testing.T) {
	defer func() {
		if got, want := fmt.Sprint(recover()), "an unclosed ( in .a:not(.b"; got != want {
			t.Errorf("panicked with %q, want %q", got, want)
		}
	}()
	compoundsIn(`.a:not(.b { }`)
}
