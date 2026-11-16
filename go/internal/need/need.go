// Package need panics when a part lacks a field it needs, naming the field
// Go-path style: "<pkg>.<Type>.<Field>[i] needs a/an X".
package need

import "strconv"

// Field panics when value is "".
//
// @trap: what carries its own article ("an ID"): a vowel rule would write
// "an URL"
func Field(part, what, value string) {
	if value == "" {
		panic(part + " needs " + what)
	}
}

// Some panics when values is empty.
func Some[T any](part, what string, values []T) {
	if len(values) == 0 {
		panic(part + " needs " + what)
	}
}

// At names a part's i-th item ("console.ConsoleTabs.Tabs[2]").
//
// @invariant: a site builds items in a loop from data, and a panic naming the
// part alone left it to bisect that data.
func At(path string, i int) string {
	return path + "[" + strconv.Itoa(i) + "]"
}
