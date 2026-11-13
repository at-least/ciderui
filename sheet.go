package ciderui

import (
	"fmt"
	"strings"
	"sync"
)

// @invariant: a utility named here reads a --cui-set-* that each set's
// stylesheet points at one of its own --cui-c-* slots; a utility of
// --color-* takes any color prefix (border-divider, text-link), one of
// --text-color-* only text-, one of --background-color-* only bg-
var themeUtilities = []struct{ variable, prefix string }{
	{"--text-color-", "text"},
	{"--background-color-", "bg"},
	{"--color-", ""},
}

// @invariant: slotNames[set]["prefix --slot"] is the utility the set names
// that slot by under that prefix ("text --cui-c-text-2" is text-secondary)
var slotNames = sync.OnceValue(func() map[string]map[string]string {
	theme := themeNames()
	out := map[string]map[string]string{}
	for _, set := range []string{"site", "console"} {
		named := map[string]string{}
		for setVariable, slot := range setSlots(string(built(set).plain)) {
			for _, t := range theme[setVariable] {
				named[t.prefix+" "+slot] = t.name
			}
		}
		out[set] = named
	}
	return out
})

type themeName struct{ prefix, name string }

func themeNames() map[string][]themeName {
	names := map[string][]themeName{}
	for line := range strings.Lines(css("theme.css")) {
		variable, value, ok := strings.Cut(strings.TrimSpace(line), ":")
		setVariable, isSet := strings.CutPrefix(strings.TrimSpace(value), "var(")
		if !ok || !isSet {
			continue
		}
		setVariable = strings.TrimSuffix(setVariable, ");")
		for _, t := range themeUtilities {
			name, ok := strings.CutPrefix(variable, t.variable)
			if !ok {
				continue
			}
			prefixes := []string{t.prefix}
			if t.prefix == "" {
				prefixes = []string{"text", "bg", "border", "outline", "decoration", "fill", "stroke", "ring", "divide", "accent", "caret"}
			}
			for _, p := range prefixes {
				names[setVariable] = append(names[setVariable], themeName{prefix: p, name: p + "-" + name})
			}
			break
		}
	}
	return names
}

func setSlots(built string) map[string]string {
	slots := map[string]string{}
	rest := built
	for {
		i := strings.Index(rest, "--cui-set-")
		if i < 0 {
			return slots
		}
		rest = rest[i:]
		variable, value, ok := strings.Cut(rest[:indexFunc(rest, func(c byte) bool { return c == ';' || c == '}' })], ":")
		rest = rest[len("--cui-set-"):]
		if slot, isSlot := strings.CutPrefix(value, "var("); ok && isSlot {
			slots[variable] = strings.TrimSuffix(slot, ")")
		}
	}
}

// @invariant: slotUtility keeps class's variants and opacity around the
// utility (md:border-(--cui-c-divider)/50 is md:border-divider/50); "" when no
// utility names the slot
func slotUtility(class string, sets []string) string {
	open := strings.Index(class, "-(--cui-c-")
	if open < 0 {
		return ""
	}
	end := strings.IndexByte(class[open:], ')')
	if end < 0 {
		return ""
	}
	variants, prefix := "", class[:open]
	if colon := strings.LastIndexByte(prefix, ':'); colon >= 0 {
		variants, prefix = prefix[:colon+1], prefix[colon+1:]
	}
	slot := class[open+2 : open+end]
	for _, set := range sets {
		if name := slotNames()[set][prefix+" "+slot]; name != "" {
			return variants + name + class[open+end+1:]
		}
	}
	return ""
}

// @invariant: further than this, a cui- class is the page's own name more
// often than a typo of the kit's
const nearestEdits = 2

func nearest(class string, known map[string]bool) string {
	best, bestDistance := "", nearestEdits+1
	for k := range known {
		if d := editDistance(class, k); d < bestDistance || d == bestDistance && best != "" && k < best {
			best, bestDistance = k, d
		}
	}
	if best == "" {
		return ""
	}
	return fmt.Sprintf("; did you mean %s", best)
}

func editDistance(a, b string) int {
	previous := make([]int, len(b)+1)
	for j := range previous {
		previous[j] = j
	}
	for i := 1; i <= len(a); i++ {
		current := make([]int, len(b)+1)
		current[0] = i
		for j := 1; j <= len(b); j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			current[j] = min(previous[j]+1, current[j-1]+1, previous[j-1]+cost)
		}
		previous = current
	}
	return previous[len(b)]
}
